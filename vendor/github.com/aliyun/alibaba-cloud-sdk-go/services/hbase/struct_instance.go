package hbase

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

// Instance is a nested struct in hbase response
type Instance struct {
	InstanceId         string `json:"InstanceId" xml:"InstanceId"`
	InstanceName       string `json:"InstanceName" xml:"InstanceName"`
	Status             string `json:"Status" xml:"Status"`
	PayType            string `json:"PayType" xml:"PayType"`
	CreatedTime        string `json:"CreatedTime" xml:"CreatedTime"`
	ExpireTime         string `json:"ExpireTime" xml:"ExpireTime"`
	MajorVersion       string `json:"MajorVersion" xml:"MajorVersion"`
	Engine             string `json:"Engine" xml:"Engine"`
	IsHa               bool   `json:"IsHa" xml:"IsHa"`
	NetworkType        string `json:"NetworkType" xml:"NetworkType"`
	VpcId              string `json:"VpcId" xml:"VpcId"`
	VswitchId          string `json:"VswitchId" xml:"VswitchId"`
	MasterInstanceType string `json:"MasterInstanceType" xml:"MasterInstanceType"`
	MasterNodeCount    int    `json:"MasterNodeCount" xml:"MasterNodeCount"`
	MasterDiskType     string `json:"MasterDiskType" xml:"MasterDiskType"`
	MasterDiskSize     int    `json:"MasterDiskSize" xml:"MasterDiskSize"`
	CoreInstanceType   string `json:"CoreInstanceType" xml:"CoreInstanceType"`
	CoreNodeCount      int    `json:"CoreNodeCount" xml:"CoreNodeCount"`
	CoreDiskType       string `json:"CoreDiskType" xml:"CoreDiskType"`
	CoreDiskSize       int    `json:"CoreDiskSize" xml:"CoreDiskSize"`
	RegionId           string `json:"RegionId" xml:"RegionId"`
	ZoneId             string `json:"ZoneId" xml:"ZoneId"`
	ColdStorageStatus  string `json:"ColdStorageStatus" xml:"ColdStorageStatus"`
	BackupStatus       string `json:"BackupStatus" xml:"BackupStatus"`
}
