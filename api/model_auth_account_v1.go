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

// AuthAccountV1 struct for AuthAccountV1
type AuthAccountV1 struct {
	Id *string `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
	RealName *string `json:"realName,omitempty"`
	Email *string `json:"email,omitempty"`
	Preferences *AccountPreferencesV1 `json:"preferences,omitempty"`
	MultiSiteAdmin *bool `json:"multiSiteAdmin,omitempty"`
	AccessLevel *string `json:"accessLevel,omitempty"`
	PrivilegeSet *string `json:"privilegeSet,omitempty"`
	PrivilegesBySite *map[string][]string `json:"privilegesBySite,omitempty"`
	GroupIds []string `json:"groupIds,omitempty"`
	CurrentSiteId *string `json:"currentSiteId,omitempty"`
}

// NewAuthAccountV1 instantiates a new AuthAccountV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAccountV1() *AuthAccountV1 {
	this := AuthAccountV1{}
	return &this
}

// NewAuthAccountV1WithDefaults instantiates a new AuthAccountV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAccountV1WithDefaults() *AuthAccountV1 {
	this := AuthAccountV1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthAccountV1) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccountV1) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthAccountV1) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthAccountV1) SetId(v string) {
	o.Id = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *AuthAccountV1) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccountV1) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *AuthAccountV1) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *AuthAccountV1) SetUsername(v string) {
	o.Username = &v
}

// GetRealName returns the RealName field value if set, zero value otherwise.
func (o *AuthAccountV1) GetRealName() string {
	if o == nil || o.RealName == nil {
		var ret string
		return ret
	}
	return *o.RealName
}

// GetRealNameOk returns a tuple with the RealName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccountV1) GetRealNameOk() (*string, bool) {
	if o == nil || o.RealName == nil {
		return nil, false
	}
	return o.RealName, true
}

// HasRealName returns a boolean if a field has been set.
func (o *AuthAccountV1) HasRealName() bool {
	if o != nil && o.RealName != nil {
		return true
	}

	return false
}

// SetRealName gets a reference to the given string and assigns it to the RealName field.
func (o *AuthAccountV1) SetRealName(v string) {
	o.RealName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AuthAccountV1) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccountV1) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AuthAccountV1) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AuthAccountV1) SetEmail(v string) {
	o.Email = &v
}

// GetPreferences returns the Preferences field value if set, zero value otherwise.
func (o *AuthAccountV1) GetPreferences() AccountPreferencesV1 {
	if o == nil || o.Preferences == nil {
		var ret AccountPreferencesV1
		return ret
	}
	return *o.Preferences
}

// GetPreferencesOk returns a tuple with the Preferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccountV1) GetPreferencesOk() (*AccountPreferencesV1, bool) {
	if o == nil || o.Preferences == nil {
		return nil, false
	}
	return o.Preferences, true
}

// HasPreferences returns a boolean if a field has been set.
func (o *AuthAccountV1) HasPreferences() bool {
	if o != nil && o.Preferences != nil {
		return true
	}

	return false
}

// SetPreferences gets a reference to the given AccountPreferencesV1 and assigns it to the Preferences field.
func (o *AuthAccountV1) SetPreferences(v AccountPreferencesV1) {
	o.Preferences = &v
}

// GetMultiSiteAdmin returns the MultiSiteAdmin field value if set, zero value otherwise.
func (o *AuthAccountV1) GetMultiSiteAdmin() bool {
	if o == nil || o.MultiSiteAdmin == nil {
		var ret bool
		return ret
	}
	return *o.MultiSiteAdmin
}

// GetMultiSiteAdminOk returns a tuple with the MultiSiteAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccountV1) GetMultiSiteAdminOk() (*bool, bool) {
	if o == nil || o.MultiSiteAdmin == nil {
		return nil, false
	}
	return o.MultiSiteAdmin, true
}

// HasMultiSiteAdmin returns a boolean if a field has been set.
func (o *AuthAccountV1) HasMultiSiteAdmin() bool {
	if o != nil && o.MultiSiteAdmin != nil {
		return true
	}

	return false
}

// SetMultiSiteAdmin gets a reference to the given bool and assigns it to the MultiSiteAdmin field.
func (o *AuthAccountV1) SetMultiSiteAdmin(v bool) {
	o.MultiSiteAdmin = &v
}

// GetAccessLevel returns the AccessLevel field value if set, zero value otherwise.
func (o *AuthAccountV1) GetAccessLevel() string {
	if o == nil || o.AccessLevel == nil {
		var ret string
		return ret
	}
	return *o.AccessLevel
}

// GetAccessLevelOk returns a tuple with the AccessLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccountV1) GetAccessLevelOk() (*string, bool) {
	if o == nil || o.AccessLevel == nil {
		return nil, false
	}
	return o.AccessLevel, true
}

// HasAccessLevel returns a boolean if a field has been set.
func (o *AuthAccountV1) HasAccessLevel() bool {
	if o != nil && o.AccessLevel != nil {
		return true
	}

	return false
}

// SetAccessLevel gets a reference to the given string and assigns it to the AccessLevel field.
func (o *AuthAccountV1) SetAccessLevel(v string) {
	o.AccessLevel = &v
}

// GetPrivilegeSet returns the PrivilegeSet field value if set, zero value otherwise.
func (o *AuthAccountV1) GetPrivilegeSet() string {
	if o == nil || o.PrivilegeSet == nil {
		var ret string
		return ret
	}
	return *o.PrivilegeSet
}

// GetPrivilegeSetOk returns a tuple with the PrivilegeSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccountV1) GetPrivilegeSetOk() (*string, bool) {
	if o == nil || o.PrivilegeSet == nil {
		return nil, false
	}
	return o.PrivilegeSet, true
}

// HasPrivilegeSet returns a boolean if a field has been set.
func (o *AuthAccountV1) HasPrivilegeSet() bool {
	if o != nil && o.PrivilegeSet != nil {
		return true
	}

	return false
}

// SetPrivilegeSet gets a reference to the given string and assigns it to the PrivilegeSet field.
func (o *AuthAccountV1) SetPrivilegeSet(v string) {
	o.PrivilegeSet = &v
}

// GetPrivilegesBySite returns the PrivilegesBySite field value if set, zero value otherwise.
func (o *AuthAccountV1) GetPrivilegesBySite() map[string][]string {
	if o == nil || o.PrivilegesBySite == nil {
		var ret map[string][]string
		return ret
	}
	return *o.PrivilegesBySite
}

// GetPrivilegesBySiteOk returns a tuple with the PrivilegesBySite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccountV1) GetPrivilegesBySiteOk() (*map[string][]string, bool) {
	if o == nil || o.PrivilegesBySite == nil {
		return nil, false
	}
	return o.PrivilegesBySite, true
}

// HasPrivilegesBySite returns a boolean if a field has been set.
func (o *AuthAccountV1) HasPrivilegesBySite() bool {
	if o != nil && o.PrivilegesBySite != nil {
		return true
	}

	return false
}

// SetPrivilegesBySite gets a reference to the given map[string][]string and assigns it to the PrivilegesBySite field.
func (o *AuthAccountV1) SetPrivilegesBySite(v map[string][]string) {
	o.PrivilegesBySite = &v
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *AuthAccountV1) GetGroupIds() []string {
	if o == nil || o.GroupIds == nil {
		var ret []string
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccountV1) GetGroupIdsOk() ([]string, bool) {
	if o == nil || o.GroupIds == nil {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *AuthAccountV1) HasGroupIds() bool {
	if o != nil && o.GroupIds != nil {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *AuthAccountV1) SetGroupIds(v []string) {
	o.GroupIds = v
}

// GetCurrentSiteId returns the CurrentSiteId field value if set, zero value otherwise.
func (o *AuthAccountV1) GetCurrentSiteId() string {
	if o == nil || o.CurrentSiteId == nil {
		var ret string
		return ret
	}
	return *o.CurrentSiteId
}

// GetCurrentSiteIdOk returns a tuple with the CurrentSiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccountV1) GetCurrentSiteIdOk() (*string, bool) {
	if o == nil || o.CurrentSiteId == nil {
		return nil, false
	}
	return o.CurrentSiteId, true
}

// HasCurrentSiteId returns a boolean if a field has been set.
func (o *AuthAccountV1) HasCurrentSiteId() bool {
	if o != nil && o.CurrentSiteId != nil {
		return true
	}

	return false
}

// SetCurrentSiteId gets a reference to the given string and assigns it to the CurrentSiteId field.
func (o *AuthAccountV1) SetCurrentSiteId(v string) {
	o.CurrentSiteId = &v
}

func (o AuthAccountV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.RealName != nil {
		toSerialize["realName"] = o.RealName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Preferences != nil {
		toSerialize["preferences"] = o.Preferences
	}
	if o.MultiSiteAdmin != nil {
		toSerialize["multiSiteAdmin"] = o.MultiSiteAdmin
	}
	if o.AccessLevel != nil {
		toSerialize["accessLevel"] = o.AccessLevel
	}
	if o.PrivilegeSet != nil {
		toSerialize["privilegeSet"] = o.PrivilegeSet
	}
	if o.PrivilegesBySite != nil {
		toSerialize["privilegesBySite"] = o.PrivilegesBySite
	}
	if o.GroupIds != nil {
		toSerialize["groupIds"] = o.GroupIds
	}
	if o.CurrentSiteId != nil {
		toSerialize["currentSiteId"] = o.CurrentSiteId
	}
	return json.Marshal(toSerialize)
}

type NullableAuthAccountV1 struct {
	value *AuthAccountV1
	isSet bool
}

func (v NullableAuthAccountV1) Get() *AuthAccountV1 {
	return v.value
}

func (v *NullableAuthAccountV1) Set(val *AuthAccountV1) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAccountV1) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAccountV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAccountV1(val *AuthAccountV1) *NullableAuthAccountV1 {
	return &NullableAuthAccountV1{value: val, isSet: true}
}

func (v NullableAuthAccountV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAccountV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


