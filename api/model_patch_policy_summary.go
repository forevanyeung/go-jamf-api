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

// PatchPolicySummary struct for PatchPolicySummary
type PatchPolicySummary struct {
	PolicyId *int32 `json:"policyId,omitempty"`
	PolicyName *string `json:"policyName,omitempty"`
	IsPolicyEnabled *bool `json:"isPolicyEnabled,omitempty"`
	PolicyTargetVersion *string `json:"policyTargetVersion,omitempty"`
	PolicyDeploymentMethod *string `json:"policyDeploymentMethod,omitempty"`
	SoftwareTitle *string `json:"softwareTitle,omitempty"`
	SoftwareTitleConfigurationId *int32 `json:"softwareTitleConfigurationId,omitempty"`
	Pending *int32 `json:"pending,omitempty"`
	Completed *int32 `json:"completed,omitempty"`
	Deferred *int32 `json:"deferred,omitempty"`
	Failed *int32 `json:"failed,omitempty"`
}

// NewPatchPolicySummary instantiates a new PatchPolicySummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchPolicySummary() *PatchPolicySummary {
	this := PatchPolicySummary{}
	return &this
}

// NewPatchPolicySummaryWithDefaults instantiates a new PatchPolicySummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchPolicySummaryWithDefaults() *PatchPolicySummary {
	this := PatchPolicySummary{}
	return &this
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *PatchPolicySummary) GetPolicyId() int32 {
	if o == nil || o.PolicyId == nil {
		var ret int32
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicySummary) GetPolicyIdOk() (*int32, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *PatchPolicySummary) HasPolicyId() bool {
	if o != nil && o.PolicyId != nil {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given int32 and assigns it to the PolicyId field.
func (o *PatchPolicySummary) SetPolicyId(v int32) {
	o.PolicyId = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *PatchPolicySummary) GetPolicyName() string {
	if o == nil || o.PolicyName == nil {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicySummary) GetPolicyNameOk() (*string, bool) {
	if o == nil || o.PolicyName == nil {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *PatchPolicySummary) HasPolicyName() bool {
	if o != nil && o.PolicyName != nil {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *PatchPolicySummary) SetPolicyName(v string) {
	o.PolicyName = &v
}

// GetIsPolicyEnabled returns the IsPolicyEnabled field value if set, zero value otherwise.
func (o *PatchPolicySummary) GetIsPolicyEnabled() bool {
	if o == nil || o.IsPolicyEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsPolicyEnabled
}

// GetIsPolicyEnabledOk returns a tuple with the IsPolicyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicySummary) GetIsPolicyEnabledOk() (*bool, bool) {
	if o == nil || o.IsPolicyEnabled == nil {
		return nil, false
	}
	return o.IsPolicyEnabled, true
}

// HasIsPolicyEnabled returns a boolean if a field has been set.
func (o *PatchPolicySummary) HasIsPolicyEnabled() bool {
	if o != nil && o.IsPolicyEnabled != nil {
		return true
	}

	return false
}

// SetIsPolicyEnabled gets a reference to the given bool and assigns it to the IsPolicyEnabled field.
func (o *PatchPolicySummary) SetIsPolicyEnabled(v bool) {
	o.IsPolicyEnabled = &v
}

// GetPolicyTargetVersion returns the PolicyTargetVersion field value if set, zero value otherwise.
func (o *PatchPolicySummary) GetPolicyTargetVersion() string {
	if o == nil || o.PolicyTargetVersion == nil {
		var ret string
		return ret
	}
	return *o.PolicyTargetVersion
}

// GetPolicyTargetVersionOk returns a tuple with the PolicyTargetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicySummary) GetPolicyTargetVersionOk() (*string, bool) {
	if o == nil || o.PolicyTargetVersion == nil {
		return nil, false
	}
	return o.PolicyTargetVersion, true
}

// HasPolicyTargetVersion returns a boolean if a field has been set.
func (o *PatchPolicySummary) HasPolicyTargetVersion() bool {
	if o != nil && o.PolicyTargetVersion != nil {
		return true
	}

	return false
}

// SetPolicyTargetVersion gets a reference to the given string and assigns it to the PolicyTargetVersion field.
func (o *PatchPolicySummary) SetPolicyTargetVersion(v string) {
	o.PolicyTargetVersion = &v
}

// GetPolicyDeploymentMethod returns the PolicyDeploymentMethod field value if set, zero value otherwise.
func (o *PatchPolicySummary) GetPolicyDeploymentMethod() string {
	if o == nil || o.PolicyDeploymentMethod == nil {
		var ret string
		return ret
	}
	return *o.PolicyDeploymentMethod
}

// GetPolicyDeploymentMethodOk returns a tuple with the PolicyDeploymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicySummary) GetPolicyDeploymentMethodOk() (*string, bool) {
	if o == nil || o.PolicyDeploymentMethod == nil {
		return nil, false
	}
	return o.PolicyDeploymentMethod, true
}

// HasPolicyDeploymentMethod returns a boolean if a field has been set.
func (o *PatchPolicySummary) HasPolicyDeploymentMethod() bool {
	if o != nil && o.PolicyDeploymentMethod != nil {
		return true
	}

	return false
}

// SetPolicyDeploymentMethod gets a reference to the given string and assigns it to the PolicyDeploymentMethod field.
func (o *PatchPolicySummary) SetPolicyDeploymentMethod(v string) {
	o.PolicyDeploymentMethod = &v
}

// GetSoftwareTitle returns the SoftwareTitle field value if set, zero value otherwise.
func (o *PatchPolicySummary) GetSoftwareTitle() string {
	if o == nil || o.SoftwareTitle == nil {
		var ret string
		return ret
	}
	return *o.SoftwareTitle
}

// GetSoftwareTitleOk returns a tuple with the SoftwareTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicySummary) GetSoftwareTitleOk() (*string, bool) {
	if o == nil || o.SoftwareTitle == nil {
		return nil, false
	}
	return o.SoftwareTitle, true
}

// HasSoftwareTitle returns a boolean if a field has been set.
func (o *PatchPolicySummary) HasSoftwareTitle() bool {
	if o != nil && o.SoftwareTitle != nil {
		return true
	}

	return false
}

// SetSoftwareTitle gets a reference to the given string and assigns it to the SoftwareTitle field.
func (o *PatchPolicySummary) SetSoftwareTitle(v string) {
	o.SoftwareTitle = &v
}

// GetSoftwareTitleConfigurationId returns the SoftwareTitleConfigurationId field value if set, zero value otherwise.
func (o *PatchPolicySummary) GetSoftwareTitleConfigurationId() int32 {
	if o == nil || o.SoftwareTitleConfigurationId == nil {
		var ret int32
		return ret
	}
	return *o.SoftwareTitleConfigurationId
}

// GetSoftwareTitleConfigurationIdOk returns a tuple with the SoftwareTitleConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicySummary) GetSoftwareTitleConfigurationIdOk() (*int32, bool) {
	if o == nil || o.SoftwareTitleConfigurationId == nil {
		return nil, false
	}
	return o.SoftwareTitleConfigurationId, true
}

// HasSoftwareTitleConfigurationId returns a boolean if a field has been set.
func (o *PatchPolicySummary) HasSoftwareTitleConfigurationId() bool {
	if o != nil && o.SoftwareTitleConfigurationId != nil {
		return true
	}

	return false
}

// SetSoftwareTitleConfigurationId gets a reference to the given int32 and assigns it to the SoftwareTitleConfigurationId field.
func (o *PatchPolicySummary) SetSoftwareTitleConfigurationId(v int32) {
	o.SoftwareTitleConfigurationId = &v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *PatchPolicySummary) GetPending() int32 {
	if o == nil || o.Pending == nil {
		var ret int32
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicySummary) GetPendingOk() (*int32, bool) {
	if o == nil || o.Pending == nil {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *PatchPolicySummary) HasPending() bool {
	if o != nil && o.Pending != nil {
		return true
	}

	return false
}

// SetPending gets a reference to the given int32 and assigns it to the Pending field.
func (o *PatchPolicySummary) SetPending(v int32) {
	o.Pending = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *PatchPolicySummary) GetCompleted() int32 {
	if o == nil || o.Completed == nil {
		var ret int32
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicySummary) GetCompletedOk() (*int32, bool) {
	if o == nil || o.Completed == nil {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *PatchPolicySummary) HasCompleted() bool {
	if o != nil && o.Completed != nil {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given int32 and assigns it to the Completed field.
func (o *PatchPolicySummary) SetCompleted(v int32) {
	o.Completed = &v
}

// GetDeferred returns the Deferred field value if set, zero value otherwise.
func (o *PatchPolicySummary) GetDeferred() int32 {
	if o == nil || o.Deferred == nil {
		var ret int32
		return ret
	}
	return *o.Deferred
}

// GetDeferredOk returns a tuple with the Deferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicySummary) GetDeferredOk() (*int32, bool) {
	if o == nil || o.Deferred == nil {
		return nil, false
	}
	return o.Deferred, true
}

// HasDeferred returns a boolean if a field has been set.
func (o *PatchPolicySummary) HasDeferred() bool {
	if o != nil && o.Deferred != nil {
		return true
	}

	return false
}

// SetDeferred gets a reference to the given int32 and assigns it to the Deferred field.
func (o *PatchPolicySummary) SetDeferred(v int32) {
	o.Deferred = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *PatchPolicySummary) GetFailed() int32 {
	if o == nil || o.Failed == nil {
		var ret int32
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicySummary) GetFailedOk() (*int32, bool) {
	if o == nil || o.Failed == nil {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *PatchPolicySummary) HasFailed() bool {
	if o != nil && o.Failed != nil {
		return true
	}

	return false
}

// SetFailed gets a reference to the given int32 and assigns it to the Failed field.
func (o *PatchPolicySummary) SetFailed(v int32) {
	o.Failed = &v
}

func (o PatchPolicySummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PolicyId != nil {
		toSerialize["policyId"] = o.PolicyId
	}
	if o.PolicyName != nil {
		toSerialize["policyName"] = o.PolicyName
	}
	if o.IsPolicyEnabled != nil {
		toSerialize["isPolicyEnabled"] = o.IsPolicyEnabled
	}
	if o.PolicyTargetVersion != nil {
		toSerialize["policyTargetVersion"] = o.PolicyTargetVersion
	}
	if o.PolicyDeploymentMethod != nil {
		toSerialize["policyDeploymentMethod"] = o.PolicyDeploymentMethod
	}
	if o.SoftwareTitle != nil {
		toSerialize["softwareTitle"] = o.SoftwareTitle
	}
	if o.SoftwareTitleConfigurationId != nil {
		toSerialize["softwareTitleConfigurationId"] = o.SoftwareTitleConfigurationId
	}
	if o.Pending != nil {
		toSerialize["pending"] = o.Pending
	}
	if o.Completed != nil {
		toSerialize["completed"] = o.Completed
	}
	if o.Deferred != nil {
		toSerialize["deferred"] = o.Deferred
	}
	if o.Failed != nil {
		toSerialize["failed"] = o.Failed
	}
	return json.Marshal(toSerialize)
}

type NullablePatchPolicySummary struct {
	value *PatchPolicySummary
	isSet bool
}

func (v NullablePatchPolicySummary) Get() *PatchPolicySummary {
	return v.value
}

func (v *NullablePatchPolicySummary) Set(val *PatchPolicySummary) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchPolicySummary) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchPolicySummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchPolicySummary(val *PatchPolicySummary) *NullablePatchPolicySummary {
	return &NullablePatchPolicySummary{value: val, isSet: true}
}

func (v NullablePatchPolicySummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchPolicySummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


