package emr

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

// CancelETLJobRelease invokes the emr.CancelETLJobRelease API synchronously
// api document: https://help.aliyun.com/api/emr/canceletljobrelease.html
func (client *Client) CancelETLJobRelease(request *CancelETLJobReleaseRequest) (response *CancelETLJobReleaseResponse, err error) {
	response = CreateCancelETLJobReleaseResponse()
	err = client.DoAction(request, response)
	return
}

// CancelETLJobReleaseWithChan invokes the emr.CancelETLJobRelease API asynchronously
// api document: https://help.aliyun.com/api/emr/canceletljobrelease.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelETLJobReleaseWithChan(request *CancelETLJobReleaseRequest) (<-chan *CancelETLJobReleaseResponse, <-chan error) {
	responseChan := make(chan *CancelETLJobReleaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelETLJobRelease(request)
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

// CancelETLJobReleaseWithCallback invokes the emr.CancelETLJobRelease API asynchronously
// api document: https://help.aliyun.com/api/emr/canceletljobrelease.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelETLJobReleaseWithCallback(request *CancelETLJobReleaseRequest, callback func(response *CancelETLJobReleaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelETLJobReleaseResponse
		var err error
		defer close(result)
		response, err = client.CancelETLJobRelease(request)
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

// CancelETLJobReleaseRequest is the request struct for api CancelETLJobRelease
type CancelETLJobReleaseRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EtlJobId        string           `position:"Query" name:"EtlJobId"`
	ReleaseId       string           `position:"Query" name:"ReleaseId"`
}

// CancelETLJobReleaseResponse is the response struct for api CancelETLJobRelease
type CancelETLJobReleaseResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelETLJobReleaseRequest creates a request to invoke CancelETLJobRelease API
func CreateCancelETLJobReleaseRequest() (request *CancelETLJobReleaseRequest) {
	request = &CancelETLJobReleaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "CancelETLJobRelease", "emr", "openAPI")
	return
}

// CreateCancelETLJobReleaseResponse creates a response to parse from CancelETLJobRelease response
func CreateCancelETLJobReleaseResponse() (response *CancelETLJobReleaseResponse) {
	response = &CancelETLJobReleaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}