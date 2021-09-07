/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-权限中心(BlueKing-IAM) available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package evaluation

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"iam/pkg/abac/pdp/condition"
	pdptypes "iam/pkg/abac/pdp/types"
	"iam/pkg/abac/types"
	"iam/pkg/cache/impls"
)

/*
求值逻辑, 包括:

对Policy的condition求值
*/

// EvalPolicies 计算是否满足
func EvalPolicies(ctx *pdptypes.ExprContext, policies []types.AuthPolicy) (isPass bool, policyID int64, err error) {
	for _, policy := range policies {
		isPass, err = evalPolicy(ctx, policy)
		if err != nil {
			log.Debugf("pdp evalPolicy: ctx=`%+v`, policy=`%+v`, error=`%s`", ctx, policy, err)
		}

		if isPass {
			log.Debugf("pdp evalPolicy: ctx=`%+v`, policy=`%+v`, pass", ctx, policy)
			return isPass, policy.ID, err
		}
	}
	// TODO: 如果一条报错, 整体结果如何???
	return false, -1, nil
}

// evalPolicy 计算单个policy是否满足
func evalPolicy(ctx *pdptypes.ExprContext, policy types.AuthPolicy) (bool, error) {
	// action 不关联资源类型时, 直接返回true
	if ctx.Action.WithoutResourceType() {
		log.Debugf("pdp evalPolicy WithoutResourceType action: %s %s", ctx.System, ctx.Action.ID)
		return true, nil
	}

	// TODO: 如果请求中没有相关的资源信息, 此时执行会怎么样?????
	//if ctx.Resource == nil {
	//	return false, fmt.Errorf("evalPolicy action: %s get resource nil", ctx.Action.ID)
	//}

	cond, err := impls.GetUnmarshalledResourceExpression(policy.Expression, policy.ExpressionSignature)
	if err != nil {
		log.Debugf("pdp evalPolicy policy id: %d expression: %s format error: %v",
			policy.ID, policy.Expression, err)

		return false, err
	}

	isPass := cond.Eval(ctx)
	return isPass, err
}

// PartialEvalPolicies 筛选check pass的policies
func PartialEvalPolicies(ctx *pdptypes.ExprContext, policies []types.AuthPolicy) ([]condition.Condition, []int64, error) {
	passConditions := make([]condition.Condition, 0, len(policies))

	passedPolicyIDs := make([]int64, 0, len(policies))
	for _, policy := range policies {
		isPass, condition, err := partialEvalPolicy(ctx, policy)
		if err != nil {
			// TODO: 一条报错怎么处理?????
			log.Debugf("pdp PartialEvalPoliciesy policy: %+v ctx: %+v error: %s", policy, ctx, err)
		}

		if condition != nil {
			passConditions = append(passConditions, condition)
		}

		if isPass {
			passedPolicyIDs = append(passedPolicyIDs, policy.ID)
		}
	}

	return passConditions, passedPolicyIDs, nil
}

func partialEvalPolicy(ctx *pdptypes.ExprContext, policy types.AuthPolicy) (bool, condition.Condition, error) {
	// action 不关联资源类型时, 直接返回true
	if ctx.Action.WithoutResourceType() {
		log.Debugf("pdp evalPolicy WithoutResourceType action: %s %s", ctx.System, ctx.Action.ID)
		return true, condition.NewAnyCondition(), nil
	}

	cond, err := impls.GetUnmarshalledResourceExpression(policy.Expression, policy.ExpressionSignature)
	if err != nil {
		log.Debugf("pdp evalPolicy policy id: %d expression: %s format error: %v",
			policy.ID, policy.Expression, err)
		return false, nil, err
	}

	switch cond.GetName() {
	case "AND", "OR":
		ok, c := cond.(condition.LogicalCondition).PartialEval(ctx)
		return ok, c, nil
	case "Any":
		return true, condition.NewAnyCondition(), nil
	default:
		key := cond.GetKeys()[0]
		dotIdx := strings.LastIndexByte(key, '.')
		if dotIdx == -1 {
			// TODO: return error
			panic("should contains dot in key")
		}
		_type := key[:dotIdx]
		if ctx.HasKey(_type) {
			if cond.Eval(ctx) {
				return true, condition.NewAnyCondition(), nil
			} else {
				return false, nil, nil
			}
		} else {
			// TODO: 这里该怎么处理, 如果传进来一个不相干的resource
			return false, cond, nil
		}
	}
}
