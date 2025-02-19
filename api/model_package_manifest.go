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

// PackageManifest struct for PackageManifest
type PackageManifest struct {
	Url string `json:"url"`
	Hash string `json:"hash"`
	HashType string `json:"hashType"`
	DisplayImageUrl *string `json:"displayImageUrl,omitempty"`
	FullSizeImageUrl *string `json:"fullSizeImageUrl,omitempty"`
	BundleId string `json:"bundleId"`
	BundleVersion string `json:"bundleVersion"`
	Subtitle *string `json:"subtitle,omitempty"`
	Title string `json:"title"`
	SizeInBytes int32 `json:"sizeInBytes"`
}

// NewPackageManifest instantiates a new PackageManifest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageManifest(url string, hash string, hashType string, bundleId string, bundleVersion string, title string, sizeInBytes int32) *PackageManifest {
	this := PackageManifest{}
	this.Url = url
	this.Hash = hash
	this.HashType = hashType
	this.BundleId = bundleId
	this.BundleVersion = bundleVersion
	this.Title = title
	this.SizeInBytes = sizeInBytes
	return &this
}

// NewPackageManifestWithDefaults instantiates a new PackageManifest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageManifestWithDefaults() *PackageManifest {
	this := PackageManifest{}
	return &this
}

// GetUrl returns the Url field value
func (o *PackageManifest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PackageManifest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PackageManifest) SetUrl(v string) {
	o.Url = v
}

// GetHash returns the Hash field value
func (o *PackageManifest) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *PackageManifest) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *PackageManifest) SetHash(v string) {
	o.Hash = v
}

// GetHashType returns the HashType field value
func (o *PackageManifest) GetHashType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HashType
}

// GetHashTypeOk returns a tuple with the HashType field value
// and a boolean to check if the value has been set.
func (o *PackageManifest) GetHashTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HashType, true
}

// SetHashType sets field value
func (o *PackageManifest) SetHashType(v string) {
	o.HashType = v
}

// GetDisplayImageUrl returns the DisplayImageUrl field value if set, zero value otherwise.
func (o *PackageManifest) GetDisplayImageUrl() string {
	if o == nil || o.DisplayImageUrl == nil {
		var ret string
		return ret
	}
	return *o.DisplayImageUrl
}

// GetDisplayImageUrlOk returns a tuple with the DisplayImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageManifest) GetDisplayImageUrlOk() (*string, bool) {
	if o == nil || o.DisplayImageUrl == nil {
		return nil, false
	}
	return o.DisplayImageUrl, true
}

// HasDisplayImageUrl returns a boolean if a field has been set.
func (o *PackageManifest) HasDisplayImageUrl() bool {
	if o != nil && o.DisplayImageUrl != nil {
		return true
	}

	return false
}

// SetDisplayImageUrl gets a reference to the given string and assigns it to the DisplayImageUrl field.
func (o *PackageManifest) SetDisplayImageUrl(v string) {
	o.DisplayImageUrl = &v
}

// GetFullSizeImageUrl returns the FullSizeImageUrl field value if set, zero value otherwise.
func (o *PackageManifest) GetFullSizeImageUrl() string {
	if o == nil || o.FullSizeImageUrl == nil {
		var ret string
		return ret
	}
	return *o.FullSizeImageUrl
}

// GetFullSizeImageUrlOk returns a tuple with the FullSizeImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageManifest) GetFullSizeImageUrlOk() (*string, bool) {
	if o == nil || o.FullSizeImageUrl == nil {
		return nil, false
	}
	return o.FullSizeImageUrl, true
}

// HasFullSizeImageUrl returns a boolean if a field has been set.
func (o *PackageManifest) HasFullSizeImageUrl() bool {
	if o != nil && o.FullSizeImageUrl != nil {
		return true
	}

	return false
}

// SetFullSizeImageUrl gets a reference to the given string and assigns it to the FullSizeImageUrl field.
func (o *PackageManifest) SetFullSizeImageUrl(v string) {
	o.FullSizeImageUrl = &v
}

// GetBundleId returns the BundleId field value
func (o *PackageManifest) GetBundleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value
// and a boolean to check if the value has been set.
func (o *PackageManifest) GetBundleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BundleId, true
}

// SetBundleId sets field value
func (o *PackageManifest) SetBundleId(v string) {
	o.BundleId = v
}

// GetBundleVersion returns the BundleVersion field value
func (o *PackageManifest) GetBundleVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BundleVersion
}

// GetBundleVersionOk returns a tuple with the BundleVersion field value
// and a boolean to check if the value has been set.
func (o *PackageManifest) GetBundleVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BundleVersion, true
}

// SetBundleVersion sets field value
func (o *PackageManifest) SetBundleVersion(v string) {
	o.BundleVersion = v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *PackageManifest) GetSubtitle() string {
	if o == nil || o.Subtitle == nil {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageManifest) GetSubtitleOk() (*string, bool) {
	if o == nil || o.Subtitle == nil {
		return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *PackageManifest) HasSubtitle() bool {
	if o != nil && o.Subtitle != nil {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *PackageManifest) SetSubtitle(v string) {
	o.Subtitle = &v
}

// GetTitle returns the Title field value
func (o *PackageManifest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PackageManifest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PackageManifest) SetTitle(v string) {
	o.Title = v
}

// GetSizeInBytes returns the SizeInBytes field value
func (o *PackageManifest) GetSizeInBytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SizeInBytes
}

// GetSizeInBytesOk returns a tuple with the SizeInBytes field value
// and a boolean to check if the value has been set.
func (o *PackageManifest) GetSizeInBytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeInBytes, true
}

// SetSizeInBytes sets field value
func (o *PackageManifest) SetSizeInBytes(v int32) {
	o.SizeInBytes = v
}

func (o PackageManifest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["hashType"] = o.HashType
	}
	if o.DisplayImageUrl != nil {
		toSerialize["displayImageUrl"] = o.DisplayImageUrl
	}
	if o.FullSizeImageUrl != nil {
		toSerialize["fullSizeImageUrl"] = o.FullSizeImageUrl
	}
	if true {
		toSerialize["bundleId"] = o.BundleId
	}
	if true {
		toSerialize["bundleVersion"] = o.BundleVersion
	}
	if o.Subtitle != nil {
		toSerialize["subtitle"] = o.Subtitle
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["sizeInBytes"] = o.SizeInBytes
	}
	return json.Marshal(toSerialize)
}

type NullablePackageManifest struct {
	value *PackageManifest
	isSet bool
}

func (v NullablePackageManifest) Get() *PackageManifest {
	return v.value
}

func (v *NullablePackageManifest) Set(val *PackageManifest) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageManifest) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageManifest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageManifest(val *PackageManifest) *NullablePackageManifest {
	return &NullablePackageManifest{value: val, isSet: true}
}

func (v NullablePackageManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageManifest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


