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

// ComputerGeneralUpdate struct for ComputerGeneralUpdate
type ComputerGeneralUpdate struct {
	Name *string `json:"name,omitempty"`
	LastIpAddress *string `json:"lastIpAddress,omitempty"`
	Barcode1 *string `json:"barcode1,omitempty"`
	Barcode2 *string `json:"barcode2,omitempty"`
	AssetTag *string `json:"assetTag,omitempty"`
	ExtensionAttributes []ComputerExtensionAttribute `json:"extensionAttributes,omitempty"`
}

// NewComputerGeneralUpdate instantiates a new ComputerGeneralUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerGeneralUpdate() *ComputerGeneralUpdate {
	this := ComputerGeneralUpdate{}
	return &this
}

// NewComputerGeneralUpdateWithDefaults instantiates a new ComputerGeneralUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerGeneralUpdateWithDefaults() *ComputerGeneralUpdate {
	this := ComputerGeneralUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputerGeneralUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneralUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputerGeneralUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputerGeneralUpdate) SetName(v string) {
	o.Name = &v
}

// GetLastIpAddress returns the LastIpAddress field value if set, zero value otherwise.
func (o *ComputerGeneralUpdate) GetLastIpAddress() string {
	if o == nil || o.LastIpAddress == nil {
		var ret string
		return ret
	}
	return *o.LastIpAddress
}

// GetLastIpAddressOk returns a tuple with the LastIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneralUpdate) GetLastIpAddressOk() (*string, bool) {
	if o == nil || o.LastIpAddress == nil {
		return nil, false
	}
	return o.LastIpAddress, true
}

// HasLastIpAddress returns a boolean if a field has been set.
func (o *ComputerGeneralUpdate) HasLastIpAddress() bool {
	if o != nil && o.LastIpAddress != nil {
		return true
	}

	return false
}

// SetLastIpAddress gets a reference to the given string and assigns it to the LastIpAddress field.
func (o *ComputerGeneralUpdate) SetLastIpAddress(v string) {
	o.LastIpAddress = &v
}

// GetBarcode1 returns the Barcode1 field value if set, zero value otherwise.
func (o *ComputerGeneralUpdate) GetBarcode1() string {
	if o == nil || o.Barcode1 == nil {
		var ret string
		return ret
	}
	return *o.Barcode1
}

// GetBarcode1Ok returns a tuple with the Barcode1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneralUpdate) GetBarcode1Ok() (*string, bool) {
	if o == nil || o.Barcode1 == nil {
		return nil, false
	}
	return o.Barcode1, true
}

// HasBarcode1 returns a boolean if a field has been set.
func (o *ComputerGeneralUpdate) HasBarcode1() bool {
	if o != nil && o.Barcode1 != nil {
		return true
	}

	return false
}

// SetBarcode1 gets a reference to the given string and assigns it to the Barcode1 field.
func (o *ComputerGeneralUpdate) SetBarcode1(v string) {
	o.Barcode1 = &v
}

// GetBarcode2 returns the Barcode2 field value if set, zero value otherwise.
func (o *ComputerGeneralUpdate) GetBarcode2() string {
	if o == nil || o.Barcode2 == nil {
		var ret string
		return ret
	}
	return *o.Barcode2
}

// GetBarcode2Ok returns a tuple with the Barcode2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneralUpdate) GetBarcode2Ok() (*string, bool) {
	if o == nil || o.Barcode2 == nil {
		return nil, false
	}
	return o.Barcode2, true
}

// HasBarcode2 returns a boolean if a field has been set.
func (o *ComputerGeneralUpdate) HasBarcode2() bool {
	if o != nil && o.Barcode2 != nil {
		return true
	}

	return false
}

// SetBarcode2 gets a reference to the given string and assigns it to the Barcode2 field.
func (o *ComputerGeneralUpdate) SetBarcode2(v string) {
	o.Barcode2 = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise.
func (o *ComputerGeneralUpdate) GetAssetTag() string {
	if o == nil || o.AssetTag == nil {
		var ret string
		return ret
	}
	return *o.AssetTag
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneralUpdate) GetAssetTagOk() (*string, bool) {
	if o == nil || o.AssetTag == nil {
		return nil, false
	}
	return o.AssetTag, true
}

// HasAssetTag returns a boolean if a field has been set.
func (o *ComputerGeneralUpdate) HasAssetTag() bool {
	if o != nil && o.AssetTag != nil {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given string and assigns it to the AssetTag field.
func (o *ComputerGeneralUpdate) SetAssetTag(v string) {
	o.AssetTag = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *ComputerGeneralUpdate) GetExtensionAttributes() []ComputerExtensionAttribute {
	if o == nil || o.ExtensionAttributes == nil {
		var ret []ComputerExtensionAttribute
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGeneralUpdate) GetExtensionAttributesOk() ([]ComputerExtensionAttribute, bool) {
	if o == nil || o.ExtensionAttributes == nil {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *ComputerGeneralUpdate) HasExtensionAttributes() bool {
	if o != nil && o.ExtensionAttributes != nil {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given []ComputerExtensionAttribute and assigns it to the ExtensionAttributes field.
func (o *ComputerGeneralUpdate) SetExtensionAttributes(v []ComputerExtensionAttribute) {
	o.ExtensionAttributes = v
}

func (o ComputerGeneralUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.LastIpAddress != nil {
		toSerialize["lastIpAddress"] = o.LastIpAddress
	}
	if o.Barcode1 != nil {
		toSerialize["barcode1"] = o.Barcode1
	}
	if o.Barcode2 != nil {
		toSerialize["barcode2"] = o.Barcode2
	}
	if o.AssetTag != nil {
		toSerialize["assetTag"] = o.AssetTag
	}
	if o.ExtensionAttributes != nil {
		toSerialize["extensionAttributes"] = o.ExtensionAttributes
	}
	return json.Marshal(toSerialize)
}

type NullableComputerGeneralUpdate struct {
	value *ComputerGeneralUpdate
	isSet bool
}

func (v NullableComputerGeneralUpdate) Get() *ComputerGeneralUpdate {
	return v.value
}

func (v *NullableComputerGeneralUpdate) Set(val *ComputerGeneralUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerGeneralUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerGeneralUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerGeneralUpdate(val *ComputerGeneralUpdate) *NullableComputerGeneralUpdate {
	return &NullableComputerGeneralUpdate{value: val, isSet: true}
}

func (v NullableComputerGeneralUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerGeneralUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


