/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


type JamfProInformationPreviewApi interface {

	/*
	V1JamfProInformationGet Get basic information about the Jamf Pro Server 

	Preview version of the endpoint. There may still be some breaking changes in the future.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1JamfProInformationGetRequest
	*/
	V1JamfProInformationGet(ctx context.Context) ApiV1JamfProInformationGetRequest

	// V1JamfProInformationGetExecute executes the request
	//  @return JamfProInformation
	V1JamfProInformationGetExecute(r ApiV1JamfProInformationGetRequest) (*JamfProInformation, *http.Response, error)
}

// JamfProInformationPreviewApiService JamfProInformationPreviewApi service
type JamfProInformationPreviewApiService service

type ApiV1JamfProInformationGetRequest struct {
	ctx context.Context
	ApiService JamfProInformationPreviewApi
}

func (r ApiV1JamfProInformationGetRequest) Execute() (*JamfProInformation, *http.Response, error) {
	return r.ApiService.V1JamfProInformationGetExecute(r)
}

/*
V1JamfProInformationGet Get basic information about the Jamf Pro Server 

Preview version of the endpoint. There may still be some breaking changes in the future.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1JamfProInformationGetRequest
*/
func (a *JamfProInformationPreviewApiService) V1JamfProInformationGet(ctx context.Context) ApiV1JamfProInformationGetRequest {
	return ApiV1JamfProInformationGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return JamfProInformation
func (a *JamfProInformationPreviewApiService) V1JamfProInformationGetExecute(r ApiV1JamfProInformationGetRequest) (*JamfProInformation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JamfProInformation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfProInformationPreviewApiService.V1JamfProInformationGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jamf-pro-information"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
