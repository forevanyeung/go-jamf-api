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
	"reflect"
)


type ClientCheckInApi interface {

	/*
	V2CheckInGet Get Client Check-In settings 

	Gets `Client Check-In` object.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV2CheckInGetRequest

	Deprecated
	*/
	V2CheckInGet(ctx context.Context) ApiV2CheckInGetRequest

	// V2CheckInGetExecute executes the request
	//  @return ClientCheckInV2
	// Deprecated
	V2CheckInGetExecute(r ApiV2CheckInGetRequest) (*ClientCheckInV2, *http.Response, error)

	/*
	V2CheckInHistoryGet Get Client Check-In history object 

	Gets Client Check-In history object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV2CheckInHistoryGetRequest

	Deprecated
	*/
	V2CheckInHistoryGet(ctx context.Context) ApiV2CheckInHistoryGetRequest

	// V2CheckInHistoryGetExecute executes the request
	//  @return HistorySearchResultsV1
	// Deprecated
	V2CheckInHistoryGetExecute(r ApiV2CheckInHistoryGetRequest) (*HistorySearchResultsV1, *http.Response, error)

	/*
	V2CheckInHistoryPost Add a Note to Client Check-In History 

	Adds Client Check-In history object notes


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV2CheckInHistoryPostRequest

	Deprecated
	*/
	V2CheckInHistoryPost(ctx context.Context) ApiV2CheckInHistoryPostRequest

	// V2CheckInHistoryPostExecute executes the request
	//  @return HrefResponse
	// Deprecated
	V2CheckInHistoryPostExecute(r ApiV2CheckInHistoryPostRequest) (*HrefResponse, *http.Response, error)

	/*
	V2CheckInPut Update Client Check-In object 

	Update Client Check-In object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV2CheckInPutRequest

	Deprecated
	*/
	V2CheckInPut(ctx context.Context) ApiV2CheckInPutRequest

	// V2CheckInPutExecute executes the request
	//  @return ClientCheckInV2
	// Deprecated
	V2CheckInPutExecute(r ApiV2CheckInPutRequest) (*ClientCheckInV2, *http.Response, error)

	/*
	V3CheckInGet Get Client Check-In settings 

	Gets `Client Check-In` object.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV3CheckInGetRequest
	*/
	V3CheckInGet(ctx context.Context) ApiV3CheckInGetRequest

	// V3CheckInGetExecute executes the request
	//  @return ClientCheckInV3
	V3CheckInGetExecute(r ApiV3CheckInGetRequest) (*ClientCheckInV3, *http.Response, error)

	/*
	V3CheckInHistoryGet Get Client Check-In history object 

	Gets Client Check-In history object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV3CheckInHistoryGetRequest
	*/
	V3CheckInHistoryGet(ctx context.Context) ApiV3CheckInHistoryGetRequest

	// V3CheckInHistoryGetExecute executes the request
	//  @return HistorySearchResultsV1
	V3CheckInHistoryGetExecute(r ApiV3CheckInHistoryGetRequest) (*HistorySearchResultsV1, *http.Response, error)

	/*
	V3CheckInHistoryPost Add a Note to Client Check-In History 

	Adds Client Check-In history object notes


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV3CheckInHistoryPostRequest
	*/
	V3CheckInHistoryPost(ctx context.Context) ApiV3CheckInHistoryPostRequest

	// V3CheckInHistoryPostExecute executes the request
	//  @return HrefResponse
	V3CheckInHistoryPostExecute(r ApiV3CheckInHistoryPostRequest) (*HrefResponse, *http.Response, error)

	/*
	V3CheckInPut Update Client Check-In object 

	Update Client Check-In object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV3CheckInPutRequest
	*/
	V3CheckInPut(ctx context.Context) ApiV3CheckInPutRequest

	// V3CheckInPutExecute executes the request
	//  @return ClientCheckInV3
	V3CheckInPutExecute(r ApiV3CheckInPutRequest) (*ClientCheckInV3, *http.Response, error)
}

// ClientCheckInApiService ClientCheckInApi service
type ClientCheckInApiService service

type ApiV2CheckInGetRequest struct {
	ctx context.Context
	ApiService ClientCheckInApi
}

func (r ApiV2CheckInGetRequest) Execute() (*ClientCheckInV2, *http.Response, error) {
	return r.ApiService.V2CheckInGetExecute(r)
}

/*
V2CheckInGet Get Client Check-In settings 

Gets `Client Check-In` object.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV2CheckInGetRequest

Deprecated
*/
func (a *ClientCheckInApiService) V2CheckInGet(ctx context.Context) ApiV2CheckInGetRequest {
	return ApiV2CheckInGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ClientCheckInV2
// Deprecated
func (a *ClientCheckInApiService) V2CheckInGetExecute(r ApiV2CheckInGetRequest) (*ClientCheckInV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClientCheckInV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCheckInApiService.V2CheckInGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/check-in"

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

type ApiV2CheckInHistoryGetRequest struct {
	ctx context.Context
	ApiService ClientCheckInApi
	page *int32
	pageSize *int32
	sort *[]string
	filter *string
}

func (r ApiV2CheckInHistoryGetRequest) Page(page int32) ApiV2CheckInHistoryGetRequest {
	r.page = &page
	return r
}

func (r ApiV2CheckInHistoryGetRequest) PageSize(pageSize int32) ApiV2CheckInHistoryGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,username:asc 
func (r ApiV2CheckInHistoryGetRequest) Sort(sort []string) ApiV2CheckInHistoryGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15
func (r ApiV2CheckInHistoryGetRequest) Filter(filter string) ApiV2CheckInHistoryGetRequest {
	r.filter = &filter
	return r
}

func (r ApiV2CheckInHistoryGetRequest) Execute() (*HistorySearchResultsV1, *http.Response, error) {
	return r.ApiService.V2CheckInHistoryGetExecute(r)
}

/*
V2CheckInHistoryGet Get Client Check-In history object 

Gets Client Check-In history object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV2CheckInHistoryGetRequest

Deprecated
*/
func (a *ClientCheckInApiService) V2CheckInHistoryGet(ctx context.Context) ApiV2CheckInHistoryGetRequest {
	return ApiV2CheckInHistoryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HistorySearchResultsV1
// Deprecated
func (a *ClientCheckInApiService) V2CheckInHistoryGetExecute(r ApiV2CheckInHistoryGetRequest) (*HistorySearchResultsV1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistorySearchResultsV1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCheckInApiService.V2CheckInHistoryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/check-in/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page-size", parameterToString(*r.pageSize, ""))
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("sort", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("sort", parameterToString(t, "multi"))
		}
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
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

type ApiV2CheckInHistoryPostRequest struct {
	ctx context.Context
	ApiService ClientCheckInApi
	objectHistoryNote *ObjectHistoryNote
}

// history notes to create
func (r ApiV2CheckInHistoryPostRequest) ObjectHistoryNote(objectHistoryNote ObjectHistoryNote) ApiV2CheckInHistoryPostRequest {
	r.objectHistoryNote = &objectHistoryNote
	return r
}

func (r ApiV2CheckInHistoryPostRequest) Execute() (*HrefResponse, *http.Response, error) {
	return r.ApiService.V2CheckInHistoryPostExecute(r)
}

/*
V2CheckInHistoryPost Add a Note to Client Check-In History 

Adds Client Check-In history object notes


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV2CheckInHistoryPostRequest

Deprecated
*/
func (a *ClientCheckInApiService) V2CheckInHistoryPost(ctx context.Context) ApiV2CheckInHistoryPostRequest {
	return ApiV2CheckInHistoryPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HrefResponse
// Deprecated
func (a *ClientCheckInApiService) V2CheckInHistoryPostExecute(r ApiV2CheckInHistoryPostRequest) (*HrefResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HrefResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCheckInApiService.V2CheckInHistoryPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/check-in/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.objectHistoryNote == nil {
		return localVarReturnValue, nil, reportError("objectHistoryNote is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.objectHistoryNote
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
		if localVarHTTPResponse.StatusCode == 503 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiV2CheckInPutRequest struct {
	ctx context.Context
	ApiService ClientCheckInApi
	clientCheckInV2 *ClientCheckInV2
}

// Client Check-In object to update
func (r ApiV2CheckInPutRequest) ClientCheckInV2(clientCheckInV2 ClientCheckInV2) ApiV2CheckInPutRequest {
	r.clientCheckInV2 = &clientCheckInV2
	return r
}

func (r ApiV2CheckInPutRequest) Execute() (*ClientCheckInV2, *http.Response, error) {
	return r.ApiService.V2CheckInPutExecute(r)
}

/*
V2CheckInPut Update Client Check-In object 

Update Client Check-In object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV2CheckInPutRequest

Deprecated
*/
func (a *ClientCheckInApiService) V2CheckInPut(ctx context.Context) ApiV2CheckInPutRequest {
	return ApiV2CheckInPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ClientCheckInV2
// Deprecated
func (a *ClientCheckInApiService) V2CheckInPutExecute(r ApiV2CheckInPutRequest) (*ClientCheckInV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClientCheckInV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCheckInApiService.V2CheckInPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/check-in"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientCheckInV2 == nil {
		return localVarReturnValue, nil, reportError("clientCheckInV2 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.clientCheckInV2
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

type ApiV3CheckInGetRequest struct {
	ctx context.Context
	ApiService ClientCheckInApi
}

func (r ApiV3CheckInGetRequest) Execute() (*ClientCheckInV3, *http.Response, error) {
	return r.ApiService.V3CheckInGetExecute(r)
}

/*
V3CheckInGet Get Client Check-In settings 

Gets `Client Check-In` object.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV3CheckInGetRequest
*/
func (a *ClientCheckInApiService) V3CheckInGet(ctx context.Context) ApiV3CheckInGetRequest {
	return ApiV3CheckInGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ClientCheckInV3
func (a *ClientCheckInApiService) V3CheckInGetExecute(r ApiV3CheckInGetRequest) (*ClientCheckInV3, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClientCheckInV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCheckInApiService.V3CheckInGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/check-in"

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

type ApiV3CheckInHistoryGetRequest struct {
	ctx context.Context
	ApiService ClientCheckInApi
	page *int32
	pageSize *int32
	sort *[]string
	filter *string
}

func (r ApiV3CheckInHistoryGetRequest) Page(page int32) ApiV3CheckInHistoryGetRequest {
	r.page = &page
	return r
}

func (r ApiV3CheckInHistoryGetRequest) PageSize(pageSize int32) ApiV3CheckInHistoryGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,username:asc 
func (r ApiV3CheckInHistoryGetRequest) Sort(sort []string) ApiV3CheckInHistoryGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15
func (r ApiV3CheckInHistoryGetRequest) Filter(filter string) ApiV3CheckInHistoryGetRequest {
	r.filter = &filter
	return r
}

func (r ApiV3CheckInHistoryGetRequest) Execute() (*HistorySearchResultsV1, *http.Response, error) {
	return r.ApiService.V3CheckInHistoryGetExecute(r)
}

/*
V3CheckInHistoryGet Get Client Check-In history object 

Gets Client Check-In history object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV3CheckInHistoryGetRequest
*/
func (a *ClientCheckInApiService) V3CheckInHistoryGet(ctx context.Context) ApiV3CheckInHistoryGetRequest {
	return ApiV3CheckInHistoryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HistorySearchResultsV1
func (a *ClientCheckInApiService) V3CheckInHistoryGetExecute(r ApiV3CheckInHistoryGetRequest) (*HistorySearchResultsV1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistorySearchResultsV1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCheckInApiService.V3CheckInHistoryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/check-in/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page-size", parameterToString(*r.pageSize, ""))
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("sort", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("sort", parameterToString(t, "multi"))
		}
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
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

type ApiV3CheckInHistoryPostRequest struct {
	ctx context.Context
	ApiService ClientCheckInApi
	objectHistoryNote *ObjectHistoryNote
}

// history notes to create
func (r ApiV3CheckInHistoryPostRequest) ObjectHistoryNote(objectHistoryNote ObjectHistoryNote) ApiV3CheckInHistoryPostRequest {
	r.objectHistoryNote = &objectHistoryNote
	return r
}

func (r ApiV3CheckInHistoryPostRequest) Execute() (*HrefResponse, *http.Response, error) {
	return r.ApiService.V3CheckInHistoryPostExecute(r)
}

/*
V3CheckInHistoryPost Add a Note to Client Check-In History 

Adds Client Check-In history object notes


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV3CheckInHistoryPostRequest
*/
func (a *ClientCheckInApiService) V3CheckInHistoryPost(ctx context.Context) ApiV3CheckInHistoryPostRequest {
	return ApiV3CheckInHistoryPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HrefResponse
func (a *ClientCheckInApiService) V3CheckInHistoryPostExecute(r ApiV3CheckInHistoryPostRequest) (*HrefResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HrefResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCheckInApiService.V3CheckInHistoryPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/check-in/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.objectHistoryNote == nil {
		return localVarReturnValue, nil, reportError("objectHistoryNote is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.objectHistoryNote
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
		if localVarHTTPResponse.StatusCode == 503 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiV3CheckInPutRequest struct {
	ctx context.Context
	ApiService ClientCheckInApi
	clientCheckInV3 *ClientCheckInV3
}

// Client Check-In object to update
func (r ApiV3CheckInPutRequest) ClientCheckInV3(clientCheckInV3 ClientCheckInV3) ApiV3CheckInPutRequest {
	r.clientCheckInV3 = &clientCheckInV3
	return r
}

func (r ApiV3CheckInPutRequest) Execute() (*ClientCheckInV3, *http.Response, error) {
	return r.ApiService.V3CheckInPutExecute(r)
}

/*
V3CheckInPut Update Client Check-In object 

Update Client Check-In object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV3CheckInPutRequest
*/
func (a *ClientCheckInApiService) V3CheckInPut(ctx context.Context) ApiV3CheckInPutRequest {
	return ApiV3CheckInPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ClientCheckInV3
func (a *ClientCheckInApiService) V3CheckInPutExecute(r ApiV3CheckInPutRequest) (*ClientCheckInV3, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClientCheckInV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCheckInApiService.V3CheckInPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/check-in"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientCheckInV3 == nil {
		return localVarReturnValue, nil, reportError("clientCheckInV3 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.clientCheckInV3
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
