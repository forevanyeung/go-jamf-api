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

// GetMobileDevicePrestage struct for GetMobileDevicePrestage
type GetMobileDevicePrestage struct {
	DisplayName string `json:"displayName"`
	IsMandatory bool `json:"isMandatory"`
	IsMdmRemovable bool `json:"isMdmRemovable"`
	SupportPhoneNumber string `json:"supportPhoneNumber"`
	SupportEmailAddress string `json:"supportEmailAddress"`
	Department string `json:"department"`
	IsDefaultPrestage bool `json:"isDefaultPrestage"`
	EnrollmentSiteId int32 `json:"enrollmentSiteId"`
	IsKeepExistingSiteMembership bool `json:"isKeepExistingSiteMembership"`
	IsKeepExistingLocationInformation bool `json:"isKeepExistingLocationInformation"`
	IsRequireAuthentication bool `json:"isRequireAuthentication"`
	AuthenticationPrompt string `json:"authenticationPrompt"`
	IsPreventActivationLock bool `json:"isPreventActivationLock"`
	IsEnableDeviceBasedActivationLock bool `json:"isEnableDeviceBasedActivationLock"`
	DeviceEnrollmentProgramInstanceId int32 `json:"deviceEnrollmentProgramInstanceId"`
	SkipSetupItems *map[string]bool `json:"skipSetupItems,omitempty"`
	LocationInformation LocationInformation `json:"locationInformation"`
	PurchasingInformation PrestagePurchasingInformation `json:"purchasingInformation"`
	// The Base64 encoded PEM Certificate
	AnchorCertificates []string `json:"anchorCertificates,omitempty"`
	EnrollmentCustomizationId *int32 `json:"enrollmentCustomizationId,omitempty"`
	IsAllowPairing bool `json:"isAllowPairing"`
	IsMultiUser bool `json:"isMultiUser"`
	IsSupervised bool `json:"isSupervised"`
	MaximumSharedAccounts int32 `json:"maximumSharedAccounts"`
	IsAutoAdvanceSetup bool `json:"isAutoAdvanceSetup"`
	IsConfigureDeviceBeforeSetupAssistant bool `json:"isConfigureDeviceBeforeSetupAssistant"`
	Language *string `json:"language,omitempty"`
	Region *string `json:"region,omitempty"`
	Names *MobileDevicePrestageNames `json:"names,omitempty"`
	Id *int32 `json:"id,omitempty"`
	ProfileUUID *string `json:"profileUUID,omitempty"`
	SiteId *int32 `json:"siteId,omitempty"`
	VersionLock *int32 `json:"versionLock,omitempty"`
}

// NewGetMobileDevicePrestage instantiates a new GetMobileDevicePrestage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMobileDevicePrestage(displayName string, isMandatory bool, isMdmRemovable bool, supportPhoneNumber string, supportEmailAddress string, department string, isDefaultPrestage bool, enrollmentSiteId int32, isKeepExistingSiteMembership bool, isKeepExistingLocationInformation bool, isRequireAuthentication bool, authenticationPrompt string, isPreventActivationLock bool, isEnableDeviceBasedActivationLock bool, deviceEnrollmentProgramInstanceId int32, locationInformation LocationInformation, purchasingInformation PrestagePurchasingInformation, isAllowPairing bool, isMultiUser bool, isSupervised bool, maximumSharedAccounts int32, isAutoAdvanceSetup bool, isConfigureDeviceBeforeSetupAssistant bool) *GetMobileDevicePrestage {
	this := GetMobileDevicePrestage{}
	this.DisplayName = displayName
	this.IsMandatory = isMandatory
	this.IsMdmRemovable = isMdmRemovable
	this.SupportPhoneNumber = supportPhoneNumber
	this.SupportEmailAddress = supportEmailAddress
	this.Department = department
	this.IsDefaultPrestage = isDefaultPrestage
	this.EnrollmentSiteId = enrollmentSiteId
	this.IsKeepExistingSiteMembership = isKeepExistingSiteMembership
	this.IsKeepExistingLocationInformation = isKeepExistingLocationInformation
	this.IsRequireAuthentication = isRequireAuthentication
	this.AuthenticationPrompt = authenticationPrompt
	this.IsPreventActivationLock = isPreventActivationLock
	this.IsEnableDeviceBasedActivationLock = isEnableDeviceBasedActivationLock
	this.DeviceEnrollmentProgramInstanceId = deviceEnrollmentProgramInstanceId
	this.LocationInformation = locationInformation
	this.PurchasingInformation = purchasingInformation
	this.IsAllowPairing = isAllowPairing
	this.IsMultiUser = isMultiUser
	this.IsSupervised = isSupervised
	this.MaximumSharedAccounts = maximumSharedAccounts
	this.IsAutoAdvanceSetup = isAutoAdvanceSetup
	this.IsConfigureDeviceBeforeSetupAssistant = isConfigureDeviceBeforeSetupAssistant
	return &this
}

// NewGetMobileDevicePrestageWithDefaults instantiates a new GetMobileDevicePrestage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMobileDevicePrestageWithDefaults() *GetMobileDevicePrestage {
	this := GetMobileDevicePrestage{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *GetMobileDevicePrestage) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *GetMobileDevicePrestage) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetIsMandatory returns the IsMandatory field value
func (o *GetMobileDevicePrestage) GetIsMandatory() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMandatory
}

// GetIsMandatoryOk returns a tuple with the IsMandatory field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetIsMandatoryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMandatory, true
}

// SetIsMandatory sets field value
func (o *GetMobileDevicePrestage) SetIsMandatory(v bool) {
	o.IsMandatory = v
}

// GetIsMdmRemovable returns the IsMdmRemovable field value
func (o *GetMobileDevicePrestage) GetIsMdmRemovable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMdmRemovable
}

// GetIsMdmRemovableOk returns a tuple with the IsMdmRemovable field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetIsMdmRemovableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMdmRemovable, true
}

// SetIsMdmRemovable sets field value
func (o *GetMobileDevicePrestage) SetIsMdmRemovable(v bool) {
	o.IsMdmRemovable = v
}

// GetSupportPhoneNumber returns the SupportPhoneNumber field value
func (o *GetMobileDevicePrestage) GetSupportPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportPhoneNumber
}

// GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetSupportPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportPhoneNumber, true
}

// SetSupportPhoneNumber sets field value
func (o *GetMobileDevicePrestage) SetSupportPhoneNumber(v string) {
	o.SupportPhoneNumber = v
}

// GetSupportEmailAddress returns the SupportEmailAddress field value
func (o *GetMobileDevicePrestage) GetSupportEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportEmailAddress
}

// GetSupportEmailAddressOk returns a tuple with the SupportEmailAddress field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetSupportEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportEmailAddress, true
}

// SetSupportEmailAddress sets field value
func (o *GetMobileDevicePrestage) SetSupportEmailAddress(v string) {
	o.SupportEmailAddress = v
}

// GetDepartment returns the Department field value
func (o *GetMobileDevicePrestage) GetDepartment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Department
}

// GetDepartmentOk returns a tuple with the Department field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetDepartmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Department, true
}

// SetDepartment sets field value
func (o *GetMobileDevicePrestage) SetDepartment(v string) {
	o.Department = v
}

// GetIsDefaultPrestage returns the IsDefaultPrestage field value
func (o *GetMobileDevicePrestage) GetIsDefaultPrestage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefaultPrestage
}

// GetIsDefaultPrestageOk returns a tuple with the IsDefaultPrestage field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetIsDefaultPrestageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefaultPrestage, true
}

// SetIsDefaultPrestage sets field value
func (o *GetMobileDevicePrestage) SetIsDefaultPrestage(v bool) {
	o.IsDefaultPrestage = v
}

// GetEnrollmentSiteId returns the EnrollmentSiteId field value
func (o *GetMobileDevicePrestage) GetEnrollmentSiteId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EnrollmentSiteId
}

// GetEnrollmentSiteIdOk returns a tuple with the EnrollmentSiteId field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetEnrollmentSiteIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnrollmentSiteId, true
}

// SetEnrollmentSiteId sets field value
func (o *GetMobileDevicePrestage) SetEnrollmentSiteId(v int32) {
	o.EnrollmentSiteId = v
}

// GetIsKeepExistingSiteMembership returns the IsKeepExistingSiteMembership field value
func (o *GetMobileDevicePrestage) GetIsKeepExistingSiteMembership() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsKeepExistingSiteMembership
}

// GetIsKeepExistingSiteMembershipOk returns a tuple with the IsKeepExistingSiteMembership field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetIsKeepExistingSiteMembershipOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsKeepExistingSiteMembership, true
}

// SetIsKeepExistingSiteMembership sets field value
func (o *GetMobileDevicePrestage) SetIsKeepExistingSiteMembership(v bool) {
	o.IsKeepExistingSiteMembership = v
}

// GetIsKeepExistingLocationInformation returns the IsKeepExistingLocationInformation field value
func (o *GetMobileDevicePrestage) GetIsKeepExistingLocationInformation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsKeepExistingLocationInformation
}

// GetIsKeepExistingLocationInformationOk returns a tuple with the IsKeepExistingLocationInformation field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetIsKeepExistingLocationInformationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsKeepExistingLocationInformation, true
}

// SetIsKeepExistingLocationInformation sets field value
func (o *GetMobileDevicePrestage) SetIsKeepExistingLocationInformation(v bool) {
	o.IsKeepExistingLocationInformation = v
}

// GetIsRequireAuthentication returns the IsRequireAuthentication field value
func (o *GetMobileDevicePrestage) GetIsRequireAuthentication() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRequireAuthentication
}

// GetIsRequireAuthenticationOk returns a tuple with the IsRequireAuthentication field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetIsRequireAuthenticationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRequireAuthentication, true
}

// SetIsRequireAuthentication sets field value
func (o *GetMobileDevicePrestage) SetIsRequireAuthentication(v bool) {
	o.IsRequireAuthentication = v
}

// GetAuthenticationPrompt returns the AuthenticationPrompt field value
func (o *GetMobileDevicePrestage) GetAuthenticationPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationPrompt
}

// GetAuthenticationPromptOk returns a tuple with the AuthenticationPrompt field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetAuthenticationPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationPrompt, true
}

// SetAuthenticationPrompt sets field value
func (o *GetMobileDevicePrestage) SetAuthenticationPrompt(v string) {
	o.AuthenticationPrompt = v
}

// GetIsPreventActivationLock returns the IsPreventActivationLock field value
func (o *GetMobileDevicePrestage) GetIsPreventActivationLock() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPreventActivationLock
}

// GetIsPreventActivationLockOk returns a tuple with the IsPreventActivationLock field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetIsPreventActivationLockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPreventActivationLock, true
}

// SetIsPreventActivationLock sets field value
func (o *GetMobileDevicePrestage) SetIsPreventActivationLock(v bool) {
	o.IsPreventActivationLock = v
}

// GetIsEnableDeviceBasedActivationLock returns the IsEnableDeviceBasedActivationLock field value
func (o *GetMobileDevicePrestage) GetIsEnableDeviceBasedActivationLock() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnableDeviceBasedActivationLock
}

// GetIsEnableDeviceBasedActivationLockOk returns a tuple with the IsEnableDeviceBasedActivationLock field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetIsEnableDeviceBasedActivationLockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnableDeviceBasedActivationLock, true
}

// SetIsEnableDeviceBasedActivationLock sets field value
func (o *GetMobileDevicePrestage) SetIsEnableDeviceBasedActivationLock(v bool) {
	o.IsEnableDeviceBasedActivationLock = v
}

// GetDeviceEnrollmentProgramInstanceId returns the DeviceEnrollmentProgramInstanceId field value
func (o *GetMobileDevicePrestage) GetDeviceEnrollmentProgramInstanceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DeviceEnrollmentProgramInstanceId
}

// GetDeviceEnrollmentProgramInstanceIdOk returns a tuple with the DeviceEnrollmentProgramInstanceId field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetDeviceEnrollmentProgramInstanceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceEnrollmentProgramInstanceId, true
}

// SetDeviceEnrollmentProgramInstanceId sets field value
func (o *GetMobileDevicePrestage) SetDeviceEnrollmentProgramInstanceId(v int32) {
	o.DeviceEnrollmentProgramInstanceId = v
}

// GetSkipSetupItems returns the SkipSetupItems field value if set, zero value otherwise.
func (o *GetMobileDevicePrestage) GetSkipSetupItems() map[string]bool {
	if o == nil || o.SkipSetupItems == nil {
		var ret map[string]bool
		return ret
	}
	return *o.SkipSetupItems
}

// GetSkipSetupItemsOk returns a tuple with the SkipSetupItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetSkipSetupItemsOk() (*map[string]bool, bool) {
	if o == nil || o.SkipSetupItems == nil {
		return nil, false
	}
	return o.SkipSetupItems, true
}

// HasSkipSetupItems returns a boolean if a field has been set.
func (o *GetMobileDevicePrestage) HasSkipSetupItems() bool {
	if o != nil && o.SkipSetupItems != nil {
		return true
	}

	return false
}

// SetSkipSetupItems gets a reference to the given map[string]bool and assigns it to the SkipSetupItems field.
func (o *GetMobileDevicePrestage) SetSkipSetupItems(v map[string]bool) {
	o.SkipSetupItems = &v
}

// GetLocationInformation returns the LocationInformation field value
func (o *GetMobileDevicePrestage) GetLocationInformation() LocationInformation {
	if o == nil {
		var ret LocationInformation
		return ret
	}

	return o.LocationInformation
}

// GetLocationInformationOk returns a tuple with the LocationInformation field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetLocationInformationOk() (*LocationInformation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationInformation, true
}

// SetLocationInformation sets field value
func (o *GetMobileDevicePrestage) SetLocationInformation(v LocationInformation) {
	o.LocationInformation = v
}

// GetPurchasingInformation returns the PurchasingInformation field value
func (o *GetMobileDevicePrestage) GetPurchasingInformation() PrestagePurchasingInformation {
	if o == nil {
		var ret PrestagePurchasingInformation
		return ret
	}

	return o.PurchasingInformation
}

// GetPurchasingInformationOk returns a tuple with the PurchasingInformation field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetPurchasingInformationOk() (*PrestagePurchasingInformation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchasingInformation, true
}

// SetPurchasingInformation sets field value
func (o *GetMobileDevicePrestage) SetPurchasingInformation(v PrestagePurchasingInformation) {
	o.PurchasingInformation = v
}

// GetAnchorCertificates returns the AnchorCertificates field value if set, zero value otherwise.
func (o *GetMobileDevicePrestage) GetAnchorCertificates() []string {
	if o == nil || o.AnchorCertificates == nil {
		var ret []string
		return ret
	}
	return o.AnchorCertificates
}

// GetAnchorCertificatesOk returns a tuple with the AnchorCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetAnchorCertificatesOk() ([]string, bool) {
	if o == nil || o.AnchorCertificates == nil {
		return nil, false
	}
	return o.AnchorCertificates, true
}

// HasAnchorCertificates returns a boolean if a field has been set.
func (o *GetMobileDevicePrestage) HasAnchorCertificates() bool {
	if o != nil && o.AnchorCertificates != nil {
		return true
	}

	return false
}

// SetAnchorCertificates gets a reference to the given []string and assigns it to the AnchorCertificates field.
func (o *GetMobileDevicePrestage) SetAnchorCertificates(v []string) {
	o.AnchorCertificates = v
}

// GetEnrollmentCustomizationId returns the EnrollmentCustomizationId field value if set, zero value otherwise.
func (o *GetMobileDevicePrestage) GetEnrollmentCustomizationId() int32 {
	if o == nil || o.EnrollmentCustomizationId == nil {
		var ret int32
		return ret
	}
	return *o.EnrollmentCustomizationId
}

// GetEnrollmentCustomizationIdOk returns a tuple with the EnrollmentCustomizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetEnrollmentCustomizationIdOk() (*int32, bool) {
	if o == nil || o.EnrollmentCustomizationId == nil {
		return nil, false
	}
	return o.EnrollmentCustomizationId, true
}

// HasEnrollmentCustomizationId returns a boolean if a field has been set.
func (o *GetMobileDevicePrestage) HasEnrollmentCustomizationId() bool {
	if o != nil && o.EnrollmentCustomizationId != nil {
		return true
	}

	return false
}

// SetEnrollmentCustomizationId gets a reference to the given int32 and assigns it to the EnrollmentCustomizationId field.
func (o *GetMobileDevicePrestage) SetEnrollmentCustomizationId(v int32) {
	o.EnrollmentCustomizationId = &v
}

// GetIsAllowPairing returns the IsAllowPairing field value
func (o *GetMobileDevicePrestage) GetIsAllowPairing() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAllowPairing
}

// GetIsAllowPairingOk returns a tuple with the IsAllowPairing field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetIsAllowPairingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAllowPairing, true
}

// SetIsAllowPairing sets field value
func (o *GetMobileDevicePrestage) SetIsAllowPairing(v bool) {
	o.IsAllowPairing = v
}

// GetIsMultiUser returns the IsMultiUser field value
func (o *GetMobileDevicePrestage) GetIsMultiUser() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMultiUser
}

// GetIsMultiUserOk returns a tuple with the IsMultiUser field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetIsMultiUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMultiUser, true
}

// SetIsMultiUser sets field value
func (o *GetMobileDevicePrestage) SetIsMultiUser(v bool) {
	o.IsMultiUser = v
}

// GetIsSupervised returns the IsSupervised field value
func (o *GetMobileDevicePrestage) GetIsSupervised() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSupervised
}

// GetIsSupervisedOk returns a tuple with the IsSupervised field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetIsSupervisedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSupervised, true
}

// SetIsSupervised sets field value
func (o *GetMobileDevicePrestage) SetIsSupervised(v bool) {
	o.IsSupervised = v
}

// GetMaximumSharedAccounts returns the MaximumSharedAccounts field value
func (o *GetMobileDevicePrestage) GetMaximumSharedAccounts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaximumSharedAccounts
}

// GetMaximumSharedAccountsOk returns a tuple with the MaximumSharedAccounts field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetMaximumSharedAccountsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaximumSharedAccounts, true
}

// SetMaximumSharedAccounts sets field value
func (o *GetMobileDevicePrestage) SetMaximumSharedAccounts(v int32) {
	o.MaximumSharedAccounts = v
}

// GetIsAutoAdvanceSetup returns the IsAutoAdvanceSetup field value
func (o *GetMobileDevicePrestage) GetIsAutoAdvanceSetup() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAutoAdvanceSetup
}

// GetIsAutoAdvanceSetupOk returns a tuple with the IsAutoAdvanceSetup field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetIsAutoAdvanceSetupOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAutoAdvanceSetup, true
}

// SetIsAutoAdvanceSetup sets field value
func (o *GetMobileDevicePrestage) SetIsAutoAdvanceSetup(v bool) {
	o.IsAutoAdvanceSetup = v
}

// GetIsConfigureDeviceBeforeSetupAssistant returns the IsConfigureDeviceBeforeSetupAssistant field value
func (o *GetMobileDevicePrestage) GetIsConfigureDeviceBeforeSetupAssistant() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsConfigureDeviceBeforeSetupAssistant
}

// GetIsConfigureDeviceBeforeSetupAssistantOk returns a tuple with the IsConfigureDeviceBeforeSetupAssistant field value
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetIsConfigureDeviceBeforeSetupAssistantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsConfigureDeviceBeforeSetupAssistant, true
}

// SetIsConfigureDeviceBeforeSetupAssistant sets field value
func (o *GetMobileDevicePrestage) SetIsConfigureDeviceBeforeSetupAssistant(v bool) {
	o.IsConfigureDeviceBeforeSetupAssistant = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *GetMobileDevicePrestage) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *GetMobileDevicePrestage) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *GetMobileDevicePrestage) SetLanguage(v string) {
	o.Language = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GetMobileDevicePrestage) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GetMobileDevicePrestage) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *GetMobileDevicePrestage) SetRegion(v string) {
	o.Region = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *GetMobileDevicePrestage) GetNames() MobileDevicePrestageNames {
	if o == nil || o.Names == nil {
		var ret MobileDevicePrestageNames
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetNamesOk() (*MobileDevicePrestageNames, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *GetMobileDevicePrestage) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given MobileDevicePrestageNames and assigns it to the Names field.
func (o *GetMobileDevicePrestage) SetNames(v MobileDevicePrestageNames) {
	o.Names = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetMobileDevicePrestage) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetMobileDevicePrestage) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetMobileDevicePrestage) SetId(v int32) {
	o.Id = &v
}

// GetProfileUUID returns the ProfileUUID field value if set, zero value otherwise.
func (o *GetMobileDevicePrestage) GetProfileUUID() string {
	if o == nil || o.ProfileUUID == nil {
		var ret string
		return ret
	}
	return *o.ProfileUUID
}

// GetProfileUUIDOk returns a tuple with the ProfileUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetProfileUUIDOk() (*string, bool) {
	if o == nil || o.ProfileUUID == nil {
		return nil, false
	}
	return o.ProfileUUID, true
}

// HasProfileUUID returns a boolean if a field has been set.
func (o *GetMobileDevicePrestage) HasProfileUUID() bool {
	if o != nil && o.ProfileUUID != nil {
		return true
	}

	return false
}

// SetProfileUUID gets a reference to the given string and assigns it to the ProfileUUID field.
func (o *GetMobileDevicePrestage) SetProfileUUID(v string) {
	o.ProfileUUID = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *GetMobileDevicePrestage) GetSiteId() int32 {
	if o == nil || o.SiteId == nil {
		var ret int32
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetSiteIdOk() (*int32, bool) {
	if o == nil || o.SiteId == nil {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *GetMobileDevicePrestage) HasSiteId() bool {
	if o != nil && o.SiteId != nil {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given int32 and assigns it to the SiteId field.
func (o *GetMobileDevicePrestage) SetSiteId(v int32) {
	o.SiteId = &v
}

// GetVersionLock returns the VersionLock field value if set, zero value otherwise.
func (o *GetMobileDevicePrestage) GetVersionLock() int32 {
	if o == nil || o.VersionLock == nil {
		var ret int32
		return ret
	}
	return *o.VersionLock
}

// GetVersionLockOk returns a tuple with the VersionLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMobileDevicePrestage) GetVersionLockOk() (*int32, bool) {
	if o == nil || o.VersionLock == nil {
		return nil, false
	}
	return o.VersionLock, true
}

// HasVersionLock returns a boolean if a field has been set.
func (o *GetMobileDevicePrestage) HasVersionLock() bool {
	if o != nil && o.VersionLock != nil {
		return true
	}

	return false
}

// SetVersionLock gets a reference to the given int32 and assigns it to the VersionLock field.
func (o *GetMobileDevicePrestage) SetVersionLock(v int32) {
	o.VersionLock = &v
}

func (o GetMobileDevicePrestage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["isMandatory"] = o.IsMandatory
	}
	if true {
		toSerialize["isMdmRemovable"] = o.IsMdmRemovable
	}
	if true {
		toSerialize["supportPhoneNumber"] = o.SupportPhoneNumber
	}
	if true {
		toSerialize["supportEmailAddress"] = o.SupportEmailAddress
	}
	if true {
		toSerialize["department"] = o.Department
	}
	if true {
		toSerialize["isDefaultPrestage"] = o.IsDefaultPrestage
	}
	if true {
		toSerialize["enrollmentSiteId"] = o.EnrollmentSiteId
	}
	if true {
		toSerialize["isKeepExistingSiteMembership"] = o.IsKeepExistingSiteMembership
	}
	if true {
		toSerialize["isKeepExistingLocationInformation"] = o.IsKeepExistingLocationInformation
	}
	if true {
		toSerialize["isRequireAuthentication"] = o.IsRequireAuthentication
	}
	if true {
		toSerialize["authenticationPrompt"] = o.AuthenticationPrompt
	}
	if true {
		toSerialize["isPreventActivationLock"] = o.IsPreventActivationLock
	}
	if true {
		toSerialize["isEnableDeviceBasedActivationLock"] = o.IsEnableDeviceBasedActivationLock
	}
	if true {
		toSerialize["deviceEnrollmentProgramInstanceId"] = o.DeviceEnrollmentProgramInstanceId
	}
	if o.SkipSetupItems != nil {
		toSerialize["skipSetupItems"] = o.SkipSetupItems
	}
	if true {
		toSerialize["locationInformation"] = o.LocationInformation
	}
	if true {
		toSerialize["purchasingInformation"] = o.PurchasingInformation
	}
	if o.AnchorCertificates != nil {
		toSerialize["anchorCertificates"] = o.AnchorCertificates
	}
	if o.EnrollmentCustomizationId != nil {
		toSerialize["enrollmentCustomizationId"] = o.EnrollmentCustomizationId
	}
	if true {
		toSerialize["isAllowPairing"] = o.IsAllowPairing
	}
	if true {
		toSerialize["isMultiUser"] = o.IsMultiUser
	}
	if true {
		toSerialize["isSupervised"] = o.IsSupervised
	}
	if true {
		toSerialize["maximumSharedAccounts"] = o.MaximumSharedAccounts
	}
	if true {
		toSerialize["isAutoAdvanceSetup"] = o.IsAutoAdvanceSetup
	}
	if true {
		toSerialize["isConfigureDeviceBeforeSetupAssistant"] = o.IsConfigureDeviceBeforeSetupAssistant
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ProfileUUID != nil {
		toSerialize["profileUUID"] = o.ProfileUUID
	}
	if o.SiteId != nil {
		toSerialize["siteId"] = o.SiteId
	}
	if o.VersionLock != nil {
		toSerialize["versionLock"] = o.VersionLock
	}
	return json.Marshal(toSerialize)
}

type NullableGetMobileDevicePrestage struct {
	value *GetMobileDevicePrestage
	isSet bool
}

func (v NullableGetMobileDevicePrestage) Get() *GetMobileDevicePrestage {
	return v.value
}

func (v *NullableGetMobileDevicePrestage) Set(val *GetMobileDevicePrestage) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMobileDevicePrestage) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMobileDevicePrestage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMobileDevicePrestage(val *GetMobileDevicePrestage) *NullableGetMobileDevicePrestage {
	return &NullableGetMobileDevicePrestage{value: val, isSet: true}
}

func (v NullableGetMobileDevicePrestage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMobileDevicePrestage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


