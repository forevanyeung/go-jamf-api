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

// InventoryPreloadRecordSearchResults struct for InventoryPreloadRecordSearchResults
type InventoryPreloadRecordSearchResults struct {
	TotalCount *int32 `json:"totalCount,omitempty"`
	Results []InventoryPreloadRecord `json:"results,omitempty"`
}

// NewInventoryPreloadRecordSearchResults instantiates a new InventoryPreloadRecordSearchResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryPreloadRecordSearchResults() *InventoryPreloadRecordSearchResults {
	this := InventoryPreloadRecordSearchResults{}
	return &this
}

// NewInventoryPreloadRecordSearchResultsWithDefaults instantiates a new InventoryPreloadRecordSearchResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryPreloadRecordSearchResultsWithDefaults() *InventoryPreloadRecordSearchResults {
	this := InventoryPreloadRecordSearchResults{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *InventoryPreloadRecordSearchResults) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecordSearchResults) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *InventoryPreloadRecordSearchResults) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *InventoryPreloadRecordSearchResults) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *InventoryPreloadRecordSearchResults) GetResults() []InventoryPreloadRecord {
	if o == nil || o.Results == nil {
		var ret []InventoryPreloadRecord
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecordSearchResults) GetResultsOk() ([]InventoryPreloadRecord, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *InventoryPreloadRecordSearchResults) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []InventoryPreloadRecord and assigns it to the Results field.
func (o *InventoryPreloadRecordSearchResults) SetResults(v []InventoryPreloadRecord) {
	o.Results = v
}

func (o InventoryPreloadRecordSearchResults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryPreloadRecordSearchResults struct {
	value *InventoryPreloadRecordSearchResults
	isSet bool
}

func (v NullableInventoryPreloadRecordSearchResults) Get() *InventoryPreloadRecordSearchResults {
	return v.value
}

func (v *NullableInventoryPreloadRecordSearchResults) Set(val *InventoryPreloadRecordSearchResults) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryPreloadRecordSearchResults) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryPreloadRecordSearchResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryPreloadRecordSearchResults(val *InventoryPreloadRecordSearchResults) *NullableInventoryPreloadRecordSearchResults {
	return &NullableInventoryPreloadRecordSearchResults{value: val, isSet: true}
}

func (v NullableInventoryPreloadRecordSearchResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryPreloadRecordSearchResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


