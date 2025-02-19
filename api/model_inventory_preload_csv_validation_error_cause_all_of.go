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

// InventoryPreloadCsvValidationErrorCauseAllOf struct for InventoryPreloadCsvValidationErrorCauseAllOf
type InventoryPreloadCsvValidationErrorCauseAllOf struct {
	Value *string `json:"value,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`
	Line *int32 `json:"line,omitempty"`
	FieldSize *int32 `json:"fieldSize,omitempty"`
	DeviceType *string `json:"deviceType,omitempty"`
}

// NewInventoryPreloadCsvValidationErrorCauseAllOf instantiates a new InventoryPreloadCsvValidationErrorCauseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryPreloadCsvValidationErrorCauseAllOf() *InventoryPreloadCsvValidationErrorCauseAllOf {
	this := InventoryPreloadCsvValidationErrorCauseAllOf{}
	return &this
}

// NewInventoryPreloadCsvValidationErrorCauseAllOfWithDefaults instantiates a new InventoryPreloadCsvValidationErrorCauseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryPreloadCsvValidationErrorCauseAllOfWithDefaults() *InventoryPreloadCsvValidationErrorCauseAllOf {
	this := InventoryPreloadCsvValidationErrorCauseAllOf{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) SetValue(v string) {
	o.Value = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) GetLine() int32 {
	if o == nil || o.Line == nil {
		var ret int32
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) GetLineOk() (*int32, bool) {
	if o == nil || o.Line == nil {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) HasLine() bool {
	if o != nil && o.Line != nil {
		return true
	}

	return false
}

// SetLine gets a reference to the given int32 and assigns it to the Line field.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) SetLine(v int32) {
	o.Line = &v
}

// GetFieldSize returns the FieldSize field value if set, zero value otherwise.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) GetFieldSize() int32 {
	if o == nil || o.FieldSize == nil {
		var ret int32
		return ret
	}
	return *o.FieldSize
}

// GetFieldSizeOk returns a tuple with the FieldSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) GetFieldSizeOk() (*int32, bool) {
	if o == nil || o.FieldSize == nil {
		return nil, false
	}
	return o.FieldSize, true
}

// HasFieldSize returns a boolean if a field has been set.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) HasFieldSize() bool {
	if o != nil && o.FieldSize != nil {
		return true
	}

	return false
}

// SetFieldSize gets a reference to the given int32 and assigns it to the FieldSize field.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) SetFieldSize(v int32) {
	o.FieldSize = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) GetDeviceTypeOk() (*string, bool) {
	if o == nil || o.DeviceType == nil {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) HasDeviceType() bool {
	if o != nil && o.DeviceType != nil {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *InventoryPreloadCsvValidationErrorCauseAllOf) SetDeviceType(v string) {
	o.DeviceType = &v
}

func (o InventoryPreloadCsvValidationErrorCauseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.SerialNumber != nil {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if o.Line != nil {
		toSerialize["line"] = o.Line
	}
	if o.FieldSize != nil {
		toSerialize["fieldSize"] = o.FieldSize
	}
	if o.DeviceType != nil {
		toSerialize["deviceType"] = o.DeviceType
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryPreloadCsvValidationErrorCauseAllOf struct {
	value *InventoryPreloadCsvValidationErrorCauseAllOf
	isSet bool
}

func (v NullableInventoryPreloadCsvValidationErrorCauseAllOf) Get() *InventoryPreloadCsvValidationErrorCauseAllOf {
	return v.value
}

func (v *NullableInventoryPreloadCsvValidationErrorCauseAllOf) Set(val *InventoryPreloadCsvValidationErrorCauseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryPreloadCsvValidationErrorCauseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryPreloadCsvValidationErrorCauseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryPreloadCsvValidationErrorCauseAllOf(val *InventoryPreloadCsvValidationErrorCauseAllOf) *NullableInventoryPreloadCsvValidationErrorCauseAllOf {
	return &NullableInventoryPreloadCsvValidationErrorCauseAllOf{value: val, isSet: true}
}

func (v NullableInventoryPreloadCsvValidationErrorCauseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryPreloadCsvValidationErrorCauseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


