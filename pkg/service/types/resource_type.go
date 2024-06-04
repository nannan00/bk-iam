/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-权限中心(BlueKing-IAM) available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package types

// ResourceType ...
type ResourceType struct {
	AllowEmptyFields

	ID             string                   `json:"id"              structs:"id"`
	Name           string                   `json:"name"            structs:"name"`
	NameEn         string                   `json:"name_en"         structs:"name_en"`
	Description    string                   `json:"description"     structs:"description"`
	DescriptionEn  string                   `json:"description_en"  structs:"description_en"`
	Sensitivity    int64                    `json:"sensitivity"     structs:"sensitivity"`
	Parents        []map[string]interface{} `json:"parents"         structs:"parents"`
	ProviderConfig map[string]interface{}   `json:"provider_config" structs:"provider_config"`
	Version        int64                    `json:"version"         structs:"version"`
}
