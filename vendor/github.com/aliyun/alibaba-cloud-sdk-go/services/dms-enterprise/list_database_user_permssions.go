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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListDatabaseUserPermssions invokes the dms_enterprise.ListDatabaseUserPermssions API synchronously
func (client *Client) ListDatabaseUserPermssions(request *ListDatabaseUserPermssionsRequest) (response *ListDatabaseUserPermssionsResponse, err error) {
	response = CreateListDatabaseUserPermssionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListDatabaseUserPermssionsWithChan invokes the dms_enterprise.ListDatabaseUserPermssions API asynchronously
func (client *Client) ListDatabaseUserPermssionsWithChan(request *ListDatabaseUserPermssionsRequest) (<-chan *ListDatabaseUserPermssionsResponse, <-chan error) {
	responseChan := make(chan *ListDatabaseUserPermssionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDatabaseUserPermssions(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListDatabaseUserPermssionsWithCallback invokes the dms_enterprise.ListDatabaseUserPermssions API asynchronously
func (client *Client) ListDatabaseUserPermssionsWithCallback(request *ListDatabaseUserPermssionsRequest, callback func(response *ListDatabaseUserPermssionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDatabaseUserPermssionsResponse
		var err error
		defer close(result)
		response, err = client.ListDatabaseUserPermssions(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListDatabaseUserPermssionsRequest is the request struct for api ListDatabaseUserPermssions
type ListDatabaseUserPermssionsRequest struct {
	*requests.RpcRequest
	PermType   string           `position:"Query" name:"PermType"`
	DbId       string           `position:"Query" name:"DbId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Logic      requests.Boolean `position:"Query" name:"Logic"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	UserName   string           `position:"Query" name:"UserName"`
}

// ListDatabaseUserPermssionsResponse is the response struct for api ListDatabaseUserPermssions
type ListDatabaseUserPermssionsResponse struct {
	*responses.BaseResponse
	RequestId       string                                      `json:"RequestId" xml:"RequestId"`
	Success         bool                                        `json:"Success" xml:"Success"`
	ErrorMessage    string                                      `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode       string                                      `json:"ErrorCode" xml:"ErrorCode"`
	TotalCount      int64                                       `json:"TotalCount" xml:"TotalCount"`
	UserPermissions UserPermissionsInListDatabaseUserPermssions `json:"UserPermissions" xml:"UserPermissions"`
}

// CreateListDatabaseUserPermssionsRequest creates a request to invoke ListDatabaseUserPermssions API
func CreateListDatabaseUserPermssionsRequest() (request *ListDatabaseUserPermssionsRequest) {
	request = &ListDatabaseUserPermssionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListDatabaseUserPermssions", "dmsenterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDatabaseUserPermssionsResponse creates a response to parse from ListDatabaseUserPermssions response
func CreateListDatabaseUserPermssionsResponse() (response *ListDatabaseUserPermssionsResponse) {
	response = &ListDatabaseUserPermssionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
