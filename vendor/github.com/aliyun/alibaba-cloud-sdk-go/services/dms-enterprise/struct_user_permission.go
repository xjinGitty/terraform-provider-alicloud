package dms_enterprise

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// UserPermission is a nested struct in dms_enterprise response
type UserPermission struct {
	Logic        bool                                    `json:"Logic" xml:"Logic"`
	SchemaName   string                                  `json:"SchemaName" xml:"SchemaName"`
	TableName    string                                  `json:"TableName" xml:"TableName"`
	TableId      string                                  `json:"TableId" xml:"TableId"`
	UserNickName string                                  `json:"UserNickName" xml:"UserNickName"`
	DbType       string                                  `json:"DbType" xml:"DbType"`
	ColumnName   string                                  `json:"ColumnName" xml:"ColumnName"`
	SearchName   string                                  `json:"SearchName" xml:"SearchName"`
	EnvType      string                                  `json:"EnvType" xml:"EnvType"`
	UserId       string                                  `json:"UserId" xml:"UserId"`
	InstanceId   string                                  `json:"InstanceId" xml:"InstanceId"`
	Alias        string                                  `json:"Alias" xml:"Alias"`
	DsType       string                                  `json:"DsType" xml:"DsType"`
	DbId         string                                  `json:"DbId" xml:"DbId"`
	PermDetails  PermDetailsInListDatabaseUserPermssions `json:"PermDetails" xml:"PermDetails"`
}