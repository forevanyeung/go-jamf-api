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

// SsoKeystore struct for SsoKeystore
type SsoKeystore struct {
	SsoKeystoreParse
	Keys []CertificateKey `json:"keys,omitempty"`
	Key string `json:"key"`
	Password string `json:"password"`
	Type string `json:"type"`
	KeystoreSetupType *string `json:"keystoreSetupType,omitempty"`
}

// NewSsoKeystore instantiates a new SsoKeystore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoKeystore(key string, password string, type_ string, keystorePassword string, keystoreFile string, keystoreFileName string) *SsoKeystore {
	this := SsoKeystore{}
	this.KeystorePassword = keystorePassword
	this.KeystoreFile = keystoreFile
	this.KeystoreFileName = keystoreFileName
	this.Key = key
	this.Password = password
	this.Type = type_
	return &this
}

// NewSsoKeystoreWithDefaults instantiates a new SsoKeystore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoKeystoreWithDefaults() *SsoKeystore {
	this := SsoKeystore{}
	var key string = " "
	this.Key = key
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *SsoKeystore) GetKeys() []CertificateKey {
	if o == nil || o.Keys == nil {
		var ret []CertificateKey
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoKeystore) GetKeysOk() ([]CertificateKey, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *SsoKeystore) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []CertificateKey and assigns it to the Keys field.
func (o *SsoKeystore) SetKeys(v []CertificateKey) {
	o.Keys = v
}

// GetKey returns the Key field value
func (o *SsoKeystore) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SsoKeystore) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SsoKeystore) SetKey(v string) {
	o.Key = v
}

// GetPassword returns the Password field value
func (o *SsoKeystore) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SsoKeystore) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SsoKeystore) SetPassword(v string) {
	o.Password = v
}

// GetType returns the Type field value
func (o *SsoKeystore) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SsoKeystore) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SsoKeystore) SetType(v string) {
	o.Type = v
}

// GetKeystoreSetupType returns the KeystoreSetupType field value if set, zero value otherwise.
func (o *SsoKeystore) GetKeystoreSetupType() string {
	if o == nil || o.KeystoreSetupType == nil {
		var ret string
		return ret
	}
	return *o.KeystoreSetupType
}

// GetKeystoreSetupTypeOk returns a tuple with the KeystoreSetupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoKeystore) GetKeystoreSetupTypeOk() (*string, bool) {
	if o == nil || o.KeystoreSetupType == nil {
		return nil, false
	}
	return o.KeystoreSetupType, true
}

// HasKeystoreSetupType returns a boolean if a field has been set.
func (o *SsoKeystore) HasKeystoreSetupType() bool {
	if o != nil && o.KeystoreSetupType != nil {
		return true
	}

	return false
}

// SetKeystoreSetupType gets a reference to the given string and assigns it to the KeystoreSetupType field.
func (o *SsoKeystore) SetKeystoreSetupType(v string) {
	o.KeystoreSetupType = &v
}

func (o SsoKeystore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSsoKeystoreParse, errSsoKeystoreParse := json.Marshal(o.SsoKeystoreParse)
	if errSsoKeystoreParse != nil {
		return []byte{}, errSsoKeystoreParse
	}
	errSsoKeystoreParse = json.Unmarshal([]byte(serializedSsoKeystoreParse), &toSerialize)
	if errSsoKeystoreParse != nil {
		return []byte{}, errSsoKeystoreParse
	}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.KeystoreSetupType != nil {
		toSerialize["keystoreSetupType"] = o.KeystoreSetupType
	}
	return json.Marshal(toSerialize)
}

type NullableSsoKeystore struct {
	value *SsoKeystore
	isSet bool
}

func (v NullableSsoKeystore) Get() *SsoKeystore {
	return v.value
}

func (v *NullableSsoKeystore) Set(val *SsoKeystore) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoKeystore) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoKeystore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoKeystore(val *SsoKeystore) *NullableSsoKeystore {
	return &NullableSsoKeystore{value: val, isSet: true}
}

func (v NullableSsoKeystore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoKeystore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


