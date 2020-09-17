package dcdn

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

// UpdateDcdnIpaDomain invokes the dcdn.UpdateDcdnIpaDomain API synchronously
// api document: https://help.aliyun.com/api/dcdn/updatedcdnipadomain.html
func (client *Client) UpdateDcdnIpaDomain(request *UpdateDcdnIpaDomainRequest) (response *UpdateDcdnIpaDomainResponse, err error) {
	response = CreateUpdateDcdnIpaDomainResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDcdnIpaDomainWithChan invokes the dcdn.UpdateDcdnIpaDomain API asynchronously
// api document: https://help.aliyun.com/api/dcdn/updatedcdnipadomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateDcdnIpaDomainWithChan(request *UpdateDcdnIpaDomainRequest) (<-chan *UpdateDcdnIpaDomainResponse, <-chan error) {
	responseChan := make(chan *UpdateDcdnIpaDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDcdnIpaDomain(request)
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

// UpdateDcdnIpaDomainWithCallback invokes the dcdn.UpdateDcdnIpaDomain API asynchronously
// api document: https://help.aliyun.com/api/dcdn/updatedcdnipadomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateDcdnIpaDomainWithCallback(request *UpdateDcdnIpaDomainRequest, callback func(response *UpdateDcdnIpaDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDcdnIpaDomainResponse
		var err error
		defer close(result)
		response, err = client.UpdateDcdnIpaDomain(request)
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

// UpdateDcdnIpaDomainRequest is the request struct for api UpdateDcdnIpaDomain
type UpdateDcdnIpaDomainRequest struct {
	*requests.RpcRequest
	Sources         string           `position:"Query" name:"Sources"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
	TopLevelDomain  string           `position:"Query" name:"TopLevelDomain"`
	DomainName      string           `position:"Query" name:"DomainName"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
}

// UpdateDcdnIpaDomainResponse is the response struct for api UpdateDcdnIpaDomain
type UpdateDcdnIpaDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateDcdnIpaDomainRequest creates a request to invoke UpdateDcdnIpaDomain API
func CreateUpdateDcdnIpaDomainRequest() (request *UpdateDcdnIpaDomainRequest) {
	request = &UpdateDcdnIpaDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "UpdateDcdnIpaDomain", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateDcdnIpaDomainResponse creates a response to parse from UpdateDcdnIpaDomain response
func CreateUpdateDcdnIpaDomainResponse() (response *UpdateDcdnIpaDomainResponse) {
	response = &UpdateDcdnIpaDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
