/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-权限中心(BlueKing-IAM) available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package pap

import (
	"github.com/TencentBlueKing/gopkg/errorx"

	"iam/pkg/cacheimpls"
	"iam/pkg/service"
	"iam/pkg/service/types"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE -package=mock

// DepartmentCTL ...
const DepartmentCTL = "DepartmentCTL"

type DepartmentController interface {
	ListPagingSubjectDepartment(limit, offset int64) ([]SubjectDepartment, error)
	BulkCreateSubjectDepartments(subjectDepartments []SubjectDepartment) error
	BulkUpdateSubjectDepartments(subjectDepartments []SubjectDepartment) error
	BulkDeleteSubjectDepartments(subjectIDs []string) error
}

type departmentController struct {
	service service.DepartmentService

	subjectService service.SubjectService
}

func NewDepartmentController() DepartmentController {
	return &departmentController{
		service: service.NewDepartmentService(),

		subjectService: service.NewSubjectService(),
	}
}

// ListPagingSubjectDepartment ...
func (c *departmentController) ListPagingSubjectDepartment(limit, offset int64) ([]SubjectDepartment, error) {
	errorWrapf := errorx.NewLayerFunctionErrorWrapf(DepartmentCTL, "ListPagingSubjectDepartment")
	svcSubjectDepartments, err := c.service.ListPagingSubjectDepartment(limit, offset)
	if err != nil {
		return nil, errorWrapf(err, "service.ListPagingSubjectDepartment limit=`%d` offset=`%d` fail", limit, offset)
	}

	pks := make([]int64, 0, len(svcSubjectDepartments)*5)
	for _, svcSubjectDepartment := range svcSubjectDepartments {
		pks = append(pks, svcSubjectDepartment.SubjectPK)
		pks = append(pks, svcSubjectDepartment.DepartmentPKs...)
	}

	subjectMap := make(map[int64]types.Subject, len(pks))
	for _, pk := range pks {
		// TODO 优化批量查询缓存
		subject, err := cacheimpls.GetSubjectByPK(pk)
		if err != nil {
			return nil, errorWrapf(err, "cacheimpls.GetSubjectByPK pk=`%d` fail", pk)
		}
		subjectMap[pk] = subject
	}

	subjectDepartments := make([]SubjectDepartment, 0, len(svcSubjectDepartments))
	for _, svcSubjectDepartment := range svcSubjectDepartments {
		subjectID := subjectMap[svcSubjectDepartment.SubjectPK].ID
		departmentIDs := make([]string, 0, len(svcSubjectDepartment.DepartmentPKs))
		for _, depPK := range svcSubjectDepartment.DepartmentPKs {
			departmentIDs = append(departmentIDs, subjectMap[depPK].ID)
		}

		subjectDepartments = append(subjectDepartments, SubjectDepartment{
			SubjectID:     subjectID,
			DepartmentIDs: departmentIDs,
		})
	}

	return subjectDepartments, nil
}

// BulkCreateSubjectDepartments ...
func (c *departmentController) BulkCreateSubjectDepartments(subjectDepartments []SubjectDepartment) error {
	errorWrapf := errorx.NewLayerFunctionErrorWrapf(DepartmentCTL, "BulkCreateSubjectDepartments")
	serviceSubjectDepartments, err := convertToServiceSubjectDepartments(subjectDepartments)
	if err != nil {
		return errorWrapf(err, "convertToServiceSubjectDepartments subjectDepartments=`%+v` fail", subjectDepartments)
	}

	err = c.service.BulkCreateSubjectDepartments(serviceSubjectDepartments)
	if err != nil {
		return errorWrapf(err, "BulkCreateSubjectDepartments subjectDepartments=`%+v` fail", subjectDepartments)
	}

	return nil
}

// BulkUpdateSubjectDepartments ...
func (c *departmentController) BulkUpdateSubjectDepartments(subjectDepartments []SubjectDepartment) error {
	errorWrapf := errorx.NewLayerFunctionErrorWrapf(DepartmentCTL, "BulkUpdateSubjectDepartments")
	serviceSubjectDepartments, err := convertToServiceSubjectDepartments(subjectDepartments)
	if err != nil {
		return errorWrapf(err, "convertToServiceSubjectDepartments subjectDepartments=`%+v` fail", subjectDepartments)
	}

	err = c.service.BulkUpdateSubjectDepartments(serviceSubjectDepartments)
	if err != nil {
		return errorWrapf(err, "BulkUpdateSubjectDepartments subjectDepartments=`%+v` fail", subjectDepartments)
	}

	subjectPKs := make([]int64, 0, len(serviceSubjectDepartments))
	for _, svcSubjectDepartment := range serviceSubjectDepartments {
		subjectPKs = append(subjectPKs, svcSubjectDepartment.SubjectPK)
	}

	// delete from cache
	cacheimpls.BatchDeleteSubjectCache(subjectPKs)

	return nil
}

// BulkDeleteSubjectDepartments ...
func (c *departmentController) BulkDeleteSubjectDepartments(subjectIDs []string) error {
	errorWrapf := errorx.NewLayerFunctionErrorWrapf(DepartmentCTL, "BulkDeleteSubjectDepartments")
	subjects := make([]types.Subject, 0, len(subjectIDs))
	for _, subjectID := range subjectIDs {
		subjects = append(subjects, types.Subject{
			Type: types.UserType,
			ID:   subjectID,
		})
	}

	subjectPKs, err := c.subjectService.ListPKsBySubjects(subjects)
	if err != nil {
		return errorWrapf(err, "ListPKsBySubjects subjects=`%+v` fail", subjects)
	}

	err = c.service.BulkDeleteSubjectDepartments(subjectPKs)
	if err != nil {
		return errorWrapf(err, "BulkDeleteSubjectDepartments subjectIDs=`%s` fail", subjectIDs)
	}

	// delete from cache
	cacheimpls.BatchDeleteSubjectCache(subjectPKs)

	return nil
}

func convertToServiceSubjectDepartments(subjectDepartments []SubjectDepartment) ([]types.SubjectDepartment, error) {
	// TODO 改进批量查询缓存
	serviceSubjectDepartments := make([]types.SubjectDepartment, 0, len(subjectDepartments))
	for _, subjectDepartment := range subjectDepartments {
		subjectPK, err := cacheimpls.GetSubjectPK(types.UserType, subjectDepartment.SubjectID)
		if err != nil {
			return nil, err
		}

		departmentPKs := make([]int64, 0, len(subjectDepartment.DepartmentIDs))
		for _, departmentID := range subjectDepartment.DepartmentIDs {
			departmentPK, err := cacheimpls.GetSubjectPK(types.DepartmentType, departmentID)
			if err != nil {
				return nil, err
			}
			departmentPKs = append(departmentPKs, departmentPK)
		}

		serviceSubjectDepartments = append(serviceSubjectDepartments, types.SubjectDepartment{
			SubjectPK:     subjectPK,
			DepartmentPKs: departmentPKs,
		})
	}

	return serviceSubjectDepartments, nil
}
