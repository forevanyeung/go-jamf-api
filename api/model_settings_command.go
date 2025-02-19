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

// SettingsCommand struct for SettingsCommand
type SettingsCommand struct {
	BootstrapTokenAllowed *bool `json:"bootstrapTokenAllowed,omitempty"`
	Bluetooth *bool `json:"bluetooth,omitempty"`
	AppAnalytics *AppAnalyticsSetting `json:"appAnalytics,omitempty"`
	DiagnosticSubmission *DiagnosticSubmissionSetting `json:"diagnosticSubmission,omitempty"`
	DataRoaming *DataRoamingSetting `json:"dataRoaming,omitempty"`
	VoiceRoaming *VoiceRoamingSetting `json:"voiceRoaming,omitempty"`
	PersonalHotspot *PersonalHotspotSetting `json:"personalHotspot,omitempty"`
	MaximumResidentUsers *int32 `json:"maximumResidentUsers,omitempty"`
	DeviceName *string `json:"deviceName,omitempty"`
	ApplicationAttributes *ApplicationAttributes `json:"applicationAttributes,omitempty"`
	SharedDeviceConfiguration *SharedDeviceConfiguration `json:"sharedDeviceConfiguration,omitempty"`
	ApplicationConfiguration *ApplicationConfiguration `json:"applicationConfiguration,omitempty"`
	TimeZone *string `json:"timeZone,omitempty"`
	SoftwareUpdateSettings *SoftwareUpdateSettings `json:"softwareUpdateSettings,omitempty"`
}

// NewSettingsCommand instantiates a new SettingsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsCommand() *SettingsCommand {
	this := SettingsCommand{}
	return &this
}

// NewSettingsCommandWithDefaults instantiates a new SettingsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsCommandWithDefaults() *SettingsCommand {
	this := SettingsCommand{}
	return &this
}

// GetBootstrapTokenAllowed returns the BootstrapTokenAllowed field value if set, zero value otherwise.
func (o *SettingsCommand) GetBootstrapTokenAllowed() bool {
	if o == nil || o.BootstrapTokenAllowed == nil {
		var ret bool
		return ret
	}
	return *o.BootstrapTokenAllowed
}

// GetBootstrapTokenAllowedOk returns a tuple with the BootstrapTokenAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCommand) GetBootstrapTokenAllowedOk() (*bool, bool) {
	if o == nil || o.BootstrapTokenAllowed == nil {
		return nil, false
	}
	return o.BootstrapTokenAllowed, true
}

// HasBootstrapTokenAllowed returns a boolean if a field has been set.
func (o *SettingsCommand) HasBootstrapTokenAllowed() bool {
	if o != nil && o.BootstrapTokenAllowed != nil {
		return true
	}

	return false
}

// SetBootstrapTokenAllowed gets a reference to the given bool and assigns it to the BootstrapTokenAllowed field.
func (o *SettingsCommand) SetBootstrapTokenAllowed(v bool) {
	o.BootstrapTokenAllowed = &v
}

// GetBluetooth returns the Bluetooth field value if set, zero value otherwise.
func (o *SettingsCommand) GetBluetooth() bool {
	if o == nil || o.Bluetooth == nil {
		var ret bool
		return ret
	}
	return *o.Bluetooth
}

// GetBluetoothOk returns a tuple with the Bluetooth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCommand) GetBluetoothOk() (*bool, bool) {
	if o == nil || o.Bluetooth == nil {
		return nil, false
	}
	return o.Bluetooth, true
}

// HasBluetooth returns a boolean if a field has been set.
func (o *SettingsCommand) HasBluetooth() bool {
	if o != nil && o.Bluetooth != nil {
		return true
	}

	return false
}

// SetBluetooth gets a reference to the given bool and assigns it to the Bluetooth field.
func (o *SettingsCommand) SetBluetooth(v bool) {
	o.Bluetooth = &v
}

// GetAppAnalytics returns the AppAnalytics field value if set, zero value otherwise.
func (o *SettingsCommand) GetAppAnalytics() AppAnalyticsSetting {
	if o == nil || o.AppAnalytics == nil {
		var ret AppAnalyticsSetting
		return ret
	}
	return *o.AppAnalytics
}

// GetAppAnalyticsOk returns a tuple with the AppAnalytics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCommand) GetAppAnalyticsOk() (*AppAnalyticsSetting, bool) {
	if o == nil || o.AppAnalytics == nil {
		return nil, false
	}
	return o.AppAnalytics, true
}

// HasAppAnalytics returns a boolean if a field has been set.
func (o *SettingsCommand) HasAppAnalytics() bool {
	if o != nil && o.AppAnalytics != nil {
		return true
	}

	return false
}

// SetAppAnalytics gets a reference to the given AppAnalyticsSetting and assigns it to the AppAnalytics field.
func (o *SettingsCommand) SetAppAnalytics(v AppAnalyticsSetting) {
	o.AppAnalytics = &v
}

// GetDiagnosticSubmission returns the DiagnosticSubmission field value if set, zero value otherwise.
func (o *SettingsCommand) GetDiagnosticSubmission() DiagnosticSubmissionSetting {
	if o == nil || o.DiagnosticSubmission == nil {
		var ret DiagnosticSubmissionSetting
		return ret
	}
	return *o.DiagnosticSubmission
}

// GetDiagnosticSubmissionOk returns a tuple with the DiagnosticSubmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCommand) GetDiagnosticSubmissionOk() (*DiagnosticSubmissionSetting, bool) {
	if o == nil || o.DiagnosticSubmission == nil {
		return nil, false
	}
	return o.DiagnosticSubmission, true
}

// HasDiagnosticSubmission returns a boolean if a field has been set.
func (o *SettingsCommand) HasDiagnosticSubmission() bool {
	if o != nil && o.DiagnosticSubmission != nil {
		return true
	}

	return false
}

// SetDiagnosticSubmission gets a reference to the given DiagnosticSubmissionSetting and assigns it to the DiagnosticSubmission field.
func (o *SettingsCommand) SetDiagnosticSubmission(v DiagnosticSubmissionSetting) {
	o.DiagnosticSubmission = &v
}

// GetDataRoaming returns the DataRoaming field value if set, zero value otherwise.
func (o *SettingsCommand) GetDataRoaming() DataRoamingSetting {
	if o == nil || o.DataRoaming == nil {
		var ret DataRoamingSetting
		return ret
	}
	return *o.DataRoaming
}

// GetDataRoamingOk returns a tuple with the DataRoaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCommand) GetDataRoamingOk() (*DataRoamingSetting, bool) {
	if o == nil || o.DataRoaming == nil {
		return nil, false
	}
	return o.DataRoaming, true
}

// HasDataRoaming returns a boolean if a field has been set.
func (o *SettingsCommand) HasDataRoaming() bool {
	if o != nil && o.DataRoaming != nil {
		return true
	}

	return false
}

// SetDataRoaming gets a reference to the given DataRoamingSetting and assigns it to the DataRoaming field.
func (o *SettingsCommand) SetDataRoaming(v DataRoamingSetting) {
	o.DataRoaming = &v
}

// GetVoiceRoaming returns the VoiceRoaming field value if set, zero value otherwise.
func (o *SettingsCommand) GetVoiceRoaming() VoiceRoamingSetting {
	if o == nil || o.VoiceRoaming == nil {
		var ret VoiceRoamingSetting
		return ret
	}
	return *o.VoiceRoaming
}

// GetVoiceRoamingOk returns a tuple with the VoiceRoaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCommand) GetVoiceRoamingOk() (*VoiceRoamingSetting, bool) {
	if o == nil || o.VoiceRoaming == nil {
		return nil, false
	}
	return o.VoiceRoaming, true
}

// HasVoiceRoaming returns a boolean if a field has been set.
func (o *SettingsCommand) HasVoiceRoaming() bool {
	if o != nil && o.VoiceRoaming != nil {
		return true
	}

	return false
}

// SetVoiceRoaming gets a reference to the given VoiceRoamingSetting and assigns it to the VoiceRoaming field.
func (o *SettingsCommand) SetVoiceRoaming(v VoiceRoamingSetting) {
	o.VoiceRoaming = &v
}

// GetPersonalHotspot returns the PersonalHotspot field value if set, zero value otherwise.
func (o *SettingsCommand) GetPersonalHotspot() PersonalHotspotSetting {
	if o == nil || o.PersonalHotspot == nil {
		var ret PersonalHotspotSetting
		return ret
	}
	return *o.PersonalHotspot
}

// GetPersonalHotspotOk returns a tuple with the PersonalHotspot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCommand) GetPersonalHotspotOk() (*PersonalHotspotSetting, bool) {
	if o == nil || o.PersonalHotspot == nil {
		return nil, false
	}
	return o.PersonalHotspot, true
}

// HasPersonalHotspot returns a boolean if a field has been set.
func (o *SettingsCommand) HasPersonalHotspot() bool {
	if o != nil && o.PersonalHotspot != nil {
		return true
	}

	return false
}

// SetPersonalHotspot gets a reference to the given PersonalHotspotSetting and assigns it to the PersonalHotspot field.
func (o *SettingsCommand) SetPersonalHotspot(v PersonalHotspotSetting) {
	o.PersonalHotspot = &v
}

// GetMaximumResidentUsers returns the MaximumResidentUsers field value if set, zero value otherwise.
func (o *SettingsCommand) GetMaximumResidentUsers() int32 {
	if o == nil || o.MaximumResidentUsers == nil {
		var ret int32
		return ret
	}
	return *o.MaximumResidentUsers
}

// GetMaximumResidentUsersOk returns a tuple with the MaximumResidentUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCommand) GetMaximumResidentUsersOk() (*int32, bool) {
	if o == nil || o.MaximumResidentUsers == nil {
		return nil, false
	}
	return o.MaximumResidentUsers, true
}

// HasMaximumResidentUsers returns a boolean if a field has been set.
func (o *SettingsCommand) HasMaximumResidentUsers() bool {
	if o != nil && o.MaximumResidentUsers != nil {
		return true
	}

	return false
}

// SetMaximumResidentUsers gets a reference to the given int32 and assigns it to the MaximumResidentUsers field.
func (o *SettingsCommand) SetMaximumResidentUsers(v int32) {
	o.MaximumResidentUsers = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *SettingsCommand) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCommand) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *SettingsCommand) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *SettingsCommand) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetApplicationAttributes returns the ApplicationAttributes field value if set, zero value otherwise.
func (o *SettingsCommand) GetApplicationAttributes() ApplicationAttributes {
	if o == nil || o.ApplicationAttributes == nil {
		var ret ApplicationAttributes
		return ret
	}
	return *o.ApplicationAttributes
}

// GetApplicationAttributesOk returns a tuple with the ApplicationAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCommand) GetApplicationAttributesOk() (*ApplicationAttributes, bool) {
	if o == nil || o.ApplicationAttributes == nil {
		return nil, false
	}
	return o.ApplicationAttributes, true
}

// HasApplicationAttributes returns a boolean if a field has been set.
func (o *SettingsCommand) HasApplicationAttributes() bool {
	if o != nil && o.ApplicationAttributes != nil {
		return true
	}

	return false
}

// SetApplicationAttributes gets a reference to the given ApplicationAttributes and assigns it to the ApplicationAttributes field.
func (o *SettingsCommand) SetApplicationAttributes(v ApplicationAttributes) {
	o.ApplicationAttributes = &v
}

// GetSharedDeviceConfiguration returns the SharedDeviceConfiguration field value if set, zero value otherwise.
func (o *SettingsCommand) GetSharedDeviceConfiguration() SharedDeviceConfiguration {
	if o == nil || o.SharedDeviceConfiguration == nil {
		var ret SharedDeviceConfiguration
		return ret
	}
	return *o.SharedDeviceConfiguration
}

// GetSharedDeviceConfigurationOk returns a tuple with the SharedDeviceConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCommand) GetSharedDeviceConfigurationOk() (*SharedDeviceConfiguration, bool) {
	if o == nil || o.SharedDeviceConfiguration == nil {
		return nil, false
	}
	return o.SharedDeviceConfiguration, true
}

// HasSharedDeviceConfiguration returns a boolean if a field has been set.
func (o *SettingsCommand) HasSharedDeviceConfiguration() bool {
	if o != nil && o.SharedDeviceConfiguration != nil {
		return true
	}

	return false
}

// SetSharedDeviceConfiguration gets a reference to the given SharedDeviceConfiguration and assigns it to the SharedDeviceConfiguration field.
func (o *SettingsCommand) SetSharedDeviceConfiguration(v SharedDeviceConfiguration) {
	o.SharedDeviceConfiguration = &v
}

// GetApplicationConfiguration returns the ApplicationConfiguration field value if set, zero value otherwise.
func (o *SettingsCommand) GetApplicationConfiguration() ApplicationConfiguration {
	if o == nil || o.ApplicationConfiguration == nil {
		var ret ApplicationConfiguration
		return ret
	}
	return *o.ApplicationConfiguration
}

// GetApplicationConfigurationOk returns a tuple with the ApplicationConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCommand) GetApplicationConfigurationOk() (*ApplicationConfiguration, bool) {
	if o == nil || o.ApplicationConfiguration == nil {
		return nil, false
	}
	return o.ApplicationConfiguration, true
}

// HasApplicationConfiguration returns a boolean if a field has been set.
func (o *SettingsCommand) HasApplicationConfiguration() bool {
	if o != nil && o.ApplicationConfiguration != nil {
		return true
	}

	return false
}

// SetApplicationConfiguration gets a reference to the given ApplicationConfiguration and assigns it to the ApplicationConfiguration field.
func (o *SettingsCommand) SetApplicationConfiguration(v ApplicationConfiguration) {
	o.ApplicationConfiguration = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *SettingsCommand) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCommand) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *SettingsCommand) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *SettingsCommand) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetSoftwareUpdateSettings returns the SoftwareUpdateSettings field value if set, zero value otherwise.
func (o *SettingsCommand) GetSoftwareUpdateSettings() SoftwareUpdateSettings {
	if o == nil || o.SoftwareUpdateSettings == nil {
		var ret SoftwareUpdateSettings
		return ret
	}
	return *o.SoftwareUpdateSettings
}

// GetSoftwareUpdateSettingsOk returns a tuple with the SoftwareUpdateSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsCommand) GetSoftwareUpdateSettingsOk() (*SoftwareUpdateSettings, bool) {
	if o == nil || o.SoftwareUpdateSettings == nil {
		return nil, false
	}
	return o.SoftwareUpdateSettings, true
}

// HasSoftwareUpdateSettings returns a boolean if a field has been set.
func (o *SettingsCommand) HasSoftwareUpdateSettings() bool {
	if o != nil && o.SoftwareUpdateSettings != nil {
		return true
	}

	return false
}

// SetSoftwareUpdateSettings gets a reference to the given SoftwareUpdateSettings and assigns it to the SoftwareUpdateSettings field.
func (o *SettingsCommand) SetSoftwareUpdateSettings(v SoftwareUpdateSettings) {
	o.SoftwareUpdateSettings = &v
}

func (o SettingsCommand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BootstrapTokenAllowed != nil {
		toSerialize["bootstrapTokenAllowed"] = o.BootstrapTokenAllowed
	}
	if o.Bluetooth != nil {
		toSerialize["bluetooth"] = o.Bluetooth
	}
	if o.AppAnalytics != nil {
		toSerialize["appAnalytics"] = o.AppAnalytics
	}
	if o.DiagnosticSubmission != nil {
		toSerialize["diagnosticSubmission"] = o.DiagnosticSubmission
	}
	if o.DataRoaming != nil {
		toSerialize["dataRoaming"] = o.DataRoaming
	}
	if o.VoiceRoaming != nil {
		toSerialize["voiceRoaming"] = o.VoiceRoaming
	}
	if o.PersonalHotspot != nil {
		toSerialize["personalHotspot"] = o.PersonalHotspot
	}
	if o.MaximumResidentUsers != nil {
		toSerialize["maximumResidentUsers"] = o.MaximumResidentUsers
	}
	if o.DeviceName != nil {
		toSerialize["deviceName"] = o.DeviceName
	}
	if o.ApplicationAttributes != nil {
		toSerialize["applicationAttributes"] = o.ApplicationAttributes
	}
	if o.SharedDeviceConfiguration != nil {
		toSerialize["sharedDeviceConfiguration"] = o.SharedDeviceConfiguration
	}
	if o.ApplicationConfiguration != nil {
		toSerialize["applicationConfiguration"] = o.ApplicationConfiguration
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}
	if o.SoftwareUpdateSettings != nil {
		toSerialize["softwareUpdateSettings"] = o.SoftwareUpdateSettings
	}
	return json.Marshal(toSerialize)
}

type NullableSettingsCommand struct {
	value *SettingsCommand
	isSet bool
}

func (v NullableSettingsCommand) Get() *SettingsCommand {
	return v.value
}

func (v *NullableSettingsCommand) Set(val *SettingsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsCommand(val *SettingsCommand) *NullableSettingsCommand {
	return &NullableSettingsCommand{value: val, isSet: true}
}

func (v NullableSettingsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


