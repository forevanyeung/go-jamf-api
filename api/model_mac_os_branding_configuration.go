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

// MacOsBrandingConfiguration struct for MacOsBrandingConfiguration
type MacOsBrandingConfiguration struct {
	Id *string `json:"id,omitempty"`
	ApplicationName *string `json:"applicationName,omitempty"`
	BrandingName *string `json:"brandingName,omitempty"`
	BrandingNameSecondary *string `json:"brandingNameSecondary,omitempty"`
	IconId *int32 `json:"iconId,omitempty"`
	BrandingHeaderImageId *int32 `json:"brandingHeaderImageId,omitempty"`
}

// NewMacOsBrandingConfiguration instantiates a new MacOsBrandingConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMacOsBrandingConfiguration() *MacOsBrandingConfiguration {
	this := MacOsBrandingConfiguration{}
	return &this
}

// NewMacOsBrandingConfigurationWithDefaults instantiates a new MacOsBrandingConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMacOsBrandingConfigurationWithDefaults() *MacOsBrandingConfiguration {
	this := MacOsBrandingConfiguration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MacOsBrandingConfiguration) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacOsBrandingConfiguration) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MacOsBrandingConfiguration) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MacOsBrandingConfiguration) SetId(v string) {
	o.Id = &v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *MacOsBrandingConfiguration) GetApplicationName() string {
	if o == nil || o.ApplicationName == nil {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacOsBrandingConfiguration) GetApplicationNameOk() (*string, bool) {
	if o == nil || o.ApplicationName == nil {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *MacOsBrandingConfiguration) HasApplicationName() bool {
	if o != nil && o.ApplicationName != nil {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *MacOsBrandingConfiguration) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetBrandingName returns the BrandingName field value if set, zero value otherwise.
func (o *MacOsBrandingConfiguration) GetBrandingName() string {
	if o == nil || o.BrandingName == nil {
		var ret string
		return ret
	}
	return *o.BrandingName
}

// GetBrandingNameOk returns a tuple with the BrandingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacOsBrandingConfiguration) GetBrandingNameOk() (*string, bool) {
	if o == nil || o.BrandingName == nil {
		return nil, false
	}
	return o.BrandingName, true
}

// HasBrandingName returns a boolean if a field has been set.
func (o *MacOsBrandingConfiguration) HasBrandingName() bool {
	if o != nil && o.BrandingName != nil {
		return true
	}

	return false
}

// SetBrandingName gets a reference to the given string and assigns it to the BrandingName field.
func (o *MacOsBrandingConfiguration) SetBrandingName(v string) {
	o.BrandingName = &v
}

// GetBrandingNameSecondary returns the BrandingNameSecondary field value if set, zero value otherwise.
func (o *MacOsBrandingConfiguration) GetBrandingNameSecondary() string {
	if o == nil || o.BrandingNameSecondary == nil {
		var ret string
		return ret
	}
	return *o.BrandingNameSecondary
}

// GetBrandingNameSecondaryOk returns a tuple with the BrandingNameSecondary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacOsBrandingConfiguration) GetBrandingNameSecondaryOk() (*string, bool) {
	if o == nil || o.BrandingNameSecondary == nil {
		return nil, false
	}
	return o.BrandingNameSecondary, true
}

// HasBrandingNameSecondary returns a boolean if a field has been set.
func (o *MacOsBrandingConfiguration) HasBrandingNameSecondary() bool {
	if o != nil && o.BrandingNameSecondary != nil {
		return true
	}

	return false
}

// SetBrandingNameSecondary gets a reference to the given string and assigns it to the BrandingNameSecondary field.
func (o *MacOsBrandingConfiguration) SetBrandingNameSecondary(v string) {
	o.BrandingNameSecondary = &v
}

// GetIconId returns the IconId field value if set, zero value otherwise.
func (o *MacOsBrandingConfiguration) GetIconId() int32 {
	if o == nil || o.IconId == nil {
		var ret int32
		return ret
	}
	return *o.IconId
}

// GetIconIdOk returns a tuple with the IconId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacOsBrandingConfiguration) GetIconIdOk() (*int32, bool) {
	if o == nil || o.IconId == nil {
		return nil, false
	}
	return o.IconId, true
}

// HasIconId returns a boolean if a field has been set.
func (o *MacOsBrandingConfiguration) HasIconId() bool {
	if o != nil && o.IconId != nil {
		return true
	}

	return false
}

// SetIconId gets a reference to the given int32 and assigns it to the IconId field.
func (o *MacOsBrandingConfiguration) SetIconId(v int32) {
	o.IconId = &v
}

// GetBrandingHeaderImageId returns the BrandingHeaderImageId field value if set, zero value otherwise.
func (o *MacOsBrandingConfiguration) GetBrandingHeaderImageId() int32 {
	if o == nil || o.BrandingHeaderImageId == nil {
		var ret int32
		return ret
	}
	return *o.BrandingHeaderImageId
}

// GetBrandingHeaderImageIdOk returns a tuple with the BrandingHeaderImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacOsBrandingConfiguration) GetBrandingHeaderImageIdOk() (*int32, bool) {
	if o == nil || o.BrandingHeaderImageId == nil {
		return nil, false
	}
	return o.BrandingHeaderImageId, true
}

// HasBrandingHeaderImageId returns a boolean if a field has been set.
func (o *MacOsBrandingConfiguration) HasBrandingHeaderImageId() bool {
	if o != nil && o.BrandingHeaderImageId != nil {
		return true
	}

	return false
}

// SetBrandingHeaderImageId gets a reference to the given int32 and assigns it to the BrandingHeaderImageId field.
func (o *MacOsBrandingConfiguration) SetBrandingHeaderImageId(v int32) {
	o.BrandingHeaderImageId = &v
}

func (o MacOsBrandingConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ApplicationName != nil {
		toSerialize["applicationName"] = o.ApplicationName
	}
	if o.BrandingName != nil {
		toSerialize["brandingName"] = o.BrandingName
	}
	if o.BrandingNameSecondary != nil {
		toSerialize["brandingNameSecondary"] = o.BrandingNameSecondary
	}
	if o.IconId != nil {
		toSerialize["iconId"] = o.IconId
	}
	if o.BrandingHeaderImageId != nil {
		toSerialize["brandingHeaderImageId"] = o.BrandingHeaderImageId
	}
	return json.Marshal(toSerialize)
}

type NullableMacOsBrandingConfiguration struct {
	value *MacOsBrandingConfiguration
	isSet bool
}

func (v NullableMacOsBrandingConfiguration) Get() *MacOsBrandingConfiguration {
	return v.value
}

func (v *NullableMacOsBrandingConfiguration) Set(val *MacOsBrandingConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableMacOsBrandingConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableMacOsBrandingConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacOsBrandingConfiguration(val *MacOsBrandingConfiguration) *NullableMacOsBrandingConfiguration {
	return &NullableMacOsBrandingConfiguration{value: val, isSet: true}
}

func (v NullableMacOsBrandingConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacOsBrandingConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


