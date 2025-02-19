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


type CacheSettingsApi interface {

	/*
	V1CacheSettingsGet Get Cache Settings 

	gets cache settings

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1CacheSettingsGetRequest
	*/
	V1CacheSettingsGet(ctx context.Context) ApiV1CacheSettingsGetRequest

	// V1CacheSettingsGetExecute executes the request
	//  @return CacheSettings
	V1CacheSettingsGetExecute(r ApiV1CacheSettingsGetRequest) (*CacheSettings, *http.Response, error)

	/*
	V1CacheSettingsPut Update Cache Settings 

	updates cache settings


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1CacheSettingsPutRequest
	*/
	V1CacheSettingsPut(ctx context.Context) ApiV1CacheSettingsPutRequest

	// V1CacheSettingsPutExecute executes the request
	//  @return CacheSettings
	V1CacheSettingsPutExecute(r ApiV1CacheSettingsPutRequest) (*CacheSettings, *http.Response, error)
}

// CacheSettingsApiService CacheSettingsApi service
type CacheSettingsApiService service

type ApiV1CacheSettingsGetRequest struct {
	ctx context.Context
	ApiService CacheSettingsApi
}

func (r ApiV1CacheSettingsGetRequest) Execute() (*CacheSettings, *http.Response, error) {
	return r.ApiService.V1CacheSettingsGetExecute(r)
}

/*
V1CacheSettingsGet Get Cache Settings 

gets cache settings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1CacheSettingsGetRequest
*/
func (a *CacheSettingsApiService) V1CacheSettingsGet(ctx context.Context) ApiV1CacheSettingsGetRequest {
	return ApiV1CacheSettingsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CacheSettings
func (a *CacheSettingsApiService) V1CacheSettingsGetExecute(r ApiV1CacheSettingsGetRequest) (*CacheSettings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CacheSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CacheSettingsApiService.V1CacheSettingsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cache-settings"

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

type ApiV1CacheSettingsPutRequest struct {
	ctx context.Context
	ApiService CacheSettingsApi
	cacheSettings *CacheSettings
}

func (r ApiV1CacheSettingsPutRequest) CacheSettings(cacheSettings CacheSettings) ApiV1CacheSettingsPutRequest {
	r.cacheSettings = &cacheSettings
	return r
}

func (r ApiV1CacheSettingsPutRequest) Execute() (*CacheSettings, *http.Response, error) {
	return r.ApiService.V1CacheSettingsPutExecute(r)
}

/*
V1CacheSettingsPut Update Cache Settings 

updates cache settings


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1CacheSettingsPutRequest
*/
func (a *CacheSettingsApiService) V1CacheSettingsPut(ctx context.Context) ApiV1CacheSettingsPutRequest {
	return ApiV1CacheSettingsPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CacheSettings
func (a *CacheSettingsApiService) V1CacheSettingsPutExecute(r ApiV1CacheSettingsPutRequest) (*CacheSettings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CacheSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CacheSettingsApiService.V1CacheSettingsPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cache-settings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cacheSettings == nil {
		return localVarReturnValue, nil, reportError("cacheSettings is required and must be specified")
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
	localVarPostBody = r.cacheSettings
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
		if localVarHTTPResponse.StatusCode == 403 {
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
