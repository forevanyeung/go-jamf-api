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

// VolumePurchasingSubscriptions struct for VolumePurchasingSubscriptions
type VolumePurchasingSubscriptions struct {
	TotalCount *int32 `json:"totalCount,omitempty"`
	Results []VolumePurchasingSubscription `json:"results,omitempty"`
}

// NewVolumePurchasingSubscriptions instantiates a new VolumePurchasingSubscriptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumePurchasingSubscriptions() *VolumePurchasingSubscriptions {
	this := VolumePurchasingSubscriptions{}
	return &this
}

// NewVolumePurchasingSubscriptionsWithDefaults instantiates a new VolumePurchasingSubscriptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumePurchasingSubscriptionsWithDefaults() *VolumePurchasingSubscriptions {
	this := VolumePurchasingSubscriptions{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *VolumePurchasingSubscriptions) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscriptions) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *VolumePurchasingSubscriptions) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *VolumePurchasingSubscriptions) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *VolumePurchasingSubscriptions) GetResults() []VolumePurchasingSubscription {
	if o == nil || o.Results == nil {
		var ret []VolumePurchasingSubscription
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingSubscriptions) GetResultsOk() ([]VolumePurchasingSubscription, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *VolumePurchasingSubscriptions) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []VolumePurchasingSubscription and assigns it to the Results field.
func (o *VolumePurchasingSubscriptions) SetResults(v []VolumePurchasingSubscription) {
	o.Results = v
}

func (o VolumePurchasingSubscriptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableVolumePurchasingSubscriptions struct {
	value *VolumePurchasingSubscriptions
	isSet bool
}

func (v NullableVolumePurchasingSubscriptions) Get() *VolumePurchasingSubscriptions {
	return v.value
}

func (v *NullableVolumePurchasingSubscriptions) Set(val *VolumePurchasingSubscriptions) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumePurchasingSubscriptions) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumePurchasingSubscriptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumePurchasingSubscriptions(val *VolumePurchasingSubscriptions) *NullableVolumePurchasingSubscriptions {
	return &NullableVolumePurchasingSubscriptions{value: val, isSet: true}
}

func (v NullableVolumePurchasingSubscriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumePurchasingSubscriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


