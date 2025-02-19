/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// CloudLdapKeystore Response with keystore information
type CloudLdapKeystore struct {
	Type *string `json:"type,omitempty"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	Subject *string `json:"subject,omitempty"`
	FileName *string `json:"fileName,omitempty"`
}

// NewCloudLdapKeystore instantiates a new CloudLdapKeystore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudLdapKeystore() *CloudLdapKeystore {
	this := CloudLdapKeystore{}
	return &this
}

// NewCloudLdapKeystoreWithDefaults instantiates a new CloudLdapKeystore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudLdapKeystoreWithDefaults() *CloudLdapKeystore {
	this := CloudLdapKeystore{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CloudLdapKeystore) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapKeystore) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CloudLdapKeystore) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CloudLdapKeystore) SetType(v string) {
	o.Type = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *CloudLdapKeystore) GetExpirationDate() time.Time {
	if o == nil || o.ExpirationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapKeystore) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *CloudLdapKeystore) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *CloudLdapKeystore) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CloudLdapKeystore) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapKeystore) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CloudLdapKeystore) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *CloudLdapKeystore) SetSubject(v string) {
	o.Subject = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *CloudLdapKeystore) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapKeystore) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *CloudLdapKeystore) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *CloudLdapKeystore) SetFileName(v string) {
	o.FileName = &v
}

func (o CloudLdapKeystore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ExpirationDate != nil {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.FileName != nil {
		toSerialize["fileName"] = o.FileName
	}
	return json.Marshal(toSerialize)
}

type NullableCloudLdapKeystore struct {
	value *CloudLdapKeystore
	isSet bool
}

func (v NullableCloudLdapKeystore) Get() *CloudLdapKeystore {
	return v.value
}

func (v *NullableCloudLdapKeystore) Set(val *CloudLdapKeystore) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudLdapKeystore) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudLdapKeystore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudLdapKeystore(val *CloudLdapKeystore) *NullableCloudLdapKeystore {
	return &NullableCloudLdapKeystore{value: val, isSet: true}
}

func (v NullableCloudLdapKeystore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudLdapKeystore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


