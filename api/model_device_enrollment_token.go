/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DeviceEnrollmentToken struct for DeviceEnrollmentToken
type DeviceEnrollmentToken struct {
	// Optional name of the token to be saved, if no name is provided one will be auto-generated
	TokenFileName *string `json:"tokenFileName,omitempty"`
	// The base 64 encoded token
	EncodedToken *string `json:"encodedToken,omitempty"`
}

// NewDeviceEnrollmentToken instantiates a new DeviceEnrollmentToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceEnrollmentToken() *DeviceEnrollmentToken {
	this := DeviceEnrollmentToken{}
	return &this
}

// NewDeviceEnrollmentTokenWithDefaults instantiates a new DeviceEnrollmentToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceEnrollmentTokenWithDefaults() *DeviceEnrollmentToken {
	this := DeviceEnrollmentToken{}
	return &this
}

// GetTokenFileName returns the TokenFileName field value if set, zero value otherwise.
func (o *DeviceEnrollmentToken) GetTokenFileName() string {
	if o == nil || o.TokenFileName == nil {
		var ret string
		return ret
	}
	return *o.TokenFileName
}

// GetTokenFileNameOk returns a tuple with the TokenFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentToken) GetTokenFileNameOk() (*string, bool) {
	if o == nil || o.TokenFileName == nil {
		return nil, false
	}
	return o.TokenFileName, true
}

// HasTokenFileName returns a boolean if a field has been set.
func (o *DeviceEnrollmentToken) HasTokenFileName() bool {
	if o != nil && o.TokenFileName != nil {
		return true
	}

	return false
}

// SetTokenFileName gets a reference to the given string and assigns it to the TokenFileName field.
func (o *DeviceEnrollmentToken) SetTokenFileName(v string) {
	o.TokenFileName = &v
}

// GetEncodedToken returns the EncodedToken field value if set, zero value otherwise.
func (o *DeviceEnrollmentToken) GetEncodedToken() string {
	if o == nil || o.EncodedToken == nil {
		var ret string
		return ret
	}
	return *o.EncodedToken
}

// GetEncodedTokenOk returns a tuple with the EncodedToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentToken) GetEncodedTokenOk() (*string, bool) {
	if o == nil || o.EncodedToken == nil {
		return nil, false
	}
	return o.EncodedToken, true
}

// HasEncodedToken returns a boolean if a field has been set.
func (o *DeviceEnrollmentToken) HasEncodedToken() bool {
	if o != nil && o.EncodedToken != nil {
		return true
	}

	return false
}

// SetEncodedToken gets a reference to the given string and assigns it to the EncodedToken field.
func (o *DeviceEnrollmentToken) SetEncodedToken(v string) {
	o.EncodedToken = &v
}

func (o DeviceEnrollmentToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TokenFileName != nil {
		toSerialize["tokenFileName"] = o.TokenFileName
	}
	if o.EncodedToken != nil {
		toSerialize["encodedToken"] = o.EncodedToken
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceEnrollmentToken struct {
	value *DeviceEnrollmentToken
	isSet bool
}

func (v NullableDeviceEnrollmentToken) Get() *DeviceEnrollmentToken {
	return v.value
}

func (v *NullableDeviceEnrollmentToken) Set(val *DeviceEnrollmentToken) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceEnrollmentToken) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceEnrollmentToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceEnrollmentToken(val *DeviceEnrollmentToken) *NullableDeviceEnrollmentToken {
	return &NullableDeviceEnrollmentToken{value: val, isSet: true}
}

func (v NullableDeviceEnrollmentToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceEnrollmentToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


