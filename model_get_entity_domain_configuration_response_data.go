/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetEntityDomainConfigurationResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEntityDomainConfigurationResponseData{}

// GetEntityDomainConfigurationResponseData struct for GetEntityDomainConfigurationResponseData
type GetEntityDomainConfigurationResponseData struct {
	NormalizeUsernames bool `json:"normalizeUsernames"`
	NormalizeHostnames bool `json:"normalizeHostnames"`
	DefaultNormalizedDomain *string `json:"defaultNormalizedDomain,omitempty"`
	DomainMappings []GetEntityDomainConfigurationResponseDataDomainMappingsInner `json:"domainMappings"`
	AdDomainNormalizationEnabled bool `json:"adDomainNormalizationEnabled"`
	FqdnNormalizationEnabled bool `json:"fqdnNormalizationEnabled"`
	AwsNormalizationEnabled bool `json:"awsNormalizationEnabled"`
}

type _GetEntityDomainConfigurationResponseData GetEntityDomainConfigurationResponseData

// NewGetEntityDomainConfigurationResponseData instantiates a new GetEntityDomainConfigurationResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEntityDomainConfigurationResponseData(normalizeUsernames bool, normalizeHostnames bool, domainMappings []GetEntityDomainConfigurationResponseDataDomainMappingsInner, adDomainNormalizationEnabled bool, fqdnNormalizationEnabled bool, awsNormalizationEnabled bool) *GetEntityDomainConfigurationResponseData {
	this := GetEntityDomainConfigurationResponseData{}
	this.NormalizeUsernames = normalizeUsernames
	this.NormalizeHostnames = normalizeHostnames
	this.DomainMappings = domainMappings
	this.AdDomainNormalizationEnabled = adDomainNormalizationEnabled
	this.FqdnNormalizationEnabled = fqdnNormalizationEnabled
	this.AwsNormalizationEnabled = awsNormalizationEnabled
	return &this
}

// NewGetEntityDomainConfigurationResponseDataWithDefaults instantiates a new GetEntityDomainConfigurationResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEntityDomainConfigurationResponseDataWithDefaults() *GetEntityDomainConfigurationResponseData {
	this := GetEntityDomainConfigurationResponseData{}
	return &this
}

// GetNormalizeUsernames returns the NormalizeUsernames field value
func (o *GetEntityDomainConfigurationResponseData) GetNormalizeUsernames() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NormalizeUsernames
}

// GetNormalizeUsernamesOk returns a tuple with the NormalizeUsernames field value
// and a boolean to check if the value has been set.
func (o *GetEntityDomainConfigurationResponseData) GetNormalizeUsernamesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NormalizeUsernames, true
}

// SetNormalizeUsernames sets field value
func (o *GetEntityDomainConfigurationResponseData) SetNormalizeUsernames(v bool) {
	o.NormalizeUsernames = v
}

// GetNormalizeHostnames returns the NormalizeHostnames field value
func (o *GetEntityDomainConfigurationResponseData) GetNormalizeHostnames() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NormalizeHostnames
}

// GetNormalizeHostnamesOk returns a tuple with the NormalizeHostnames field value
// and a boolean to check if the value has been set.
func (o *GetEntityDomainConfigurationResponseData) GetNormalizeHostnamesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NormalizeHostnames, true
}

// SetNormalizeHostnames sets field value
func (o *GetEntityDomainConfigurationResponseData) SetNormalizeHostnames(v bool) {
	o.NormalizeHostnames = v
}

// GetDefaultNormalizedDomain returns the DefaultNormalizedDomain field value if set, zero value otherwise.
func (o *GetEntityDomainConfigurationResponseData) GetDefaultNormalizedDomain() string {
	if o == nil || IsNil(o.DefaultNormalizedDomain) {
		var ret string
		return ret
	}
	return *o.DefaultNormalizedDomain
}

// GetDefaultNormalizedDomainOk returns a tuple with the DefaultNormalizedDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEntityDomainConfigurationResponseData) GetDefaultNormalizedDomainOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultNormalizedDomain) {
		return nil, false
	}
	return o.DefaultNormalizedDomain, true
}

// HasDefaultNormalizedDomain returns a boolean if a field has been set.
func (o *GetEntityDomainConfigurationResponseData) HasDefaultNormalizedDomain() bool {
	if o != nil && !IsNil(o.DefaultNormalizedDomain) {
		return true
	}

	return false
}

// SetDefaultNormalizedDomain gets a reference to the given string and assigns it to the DefaultNormalizedDomain field.
func (o *GetEntityDomainConfigurationResponseData) SetDefaultNormalizedDomain(v string) {
	o.DefaultNormalizedDomain = &v
}

// GetDomainMappings returns the DomainMappings field value
func (o *GetEntityDomainConfigurationResponseData) GetDomainMappings() []GetEntityDomainConfigurationResponseDataDomainMappingsInner {
	if o == nil {
		var ret []GetEntityDomainConfigurationResponseDataDomainMappingsInner
		return ret
	}

	return o.DomainMappings
}

// GetDomainMappingsOk returns a tuple with the DomainMappings field value
// and a boolean to check if the value has been set.
func (o *GetEntityDomainConfigurationResponseData) GetDomainMappingsOk() ([]GetEntityDomainConfigurationResponseDataDomainMappingsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.DomainMappings, true
}

// SetDomainMappings sets field value
func (o *GetEntityDomainConfigurationResponseData) SetDomainMappings(v []GetEntityDomainConfigurationResponseDataDomainMappingsInner) {
	o.DomainMappings = v
}

// GetAdDomainNormalizationEnabled returns the AdDomainNormalizationEnabled field value
func (o *GetEntityDomainConfigurationResponseData) GetAdDomainNormalizationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AdDomainNormalizationEnabled
}

// GetAdDomainNormalizationEnabledOk returns a tuple with the AdDomainNormalizationEnabled field value
// and a boolean to check if the value has been set.
func (o *GetEntityDomainConfigurationResponseData) GetAdDomainNormalizationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdDomainNormalizationEnabled, true
}

// SetAdDomainNormalizationEnabled sets field value
func (o *GetEntityDomainConfigurationResponseData) SetAdDomainNormalizationEnabled(v bool) {
	o.AdDomainNormalizationEnabled = v
}

// GetFqdnNormalizationEnabled returns the FqdnNormalizationEnabled field value
func (o *GetEntityDomainConfigurationResponseData) GetFqdnNormalizationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FqdnNormalizationEnabled
}

// GetFqdnNormalizationEnabledOk returns a tuple with the FqdnNormalizationEnabled field value
// and a boolean to check if the value has been set.
func (o *GetEntityDomainConfigurationResponseData) GetFqdnNormalizationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FqdnNormalizationEnabled, true
}

// SetFqdnNormalizationEnabled sets field value
func (o *GetEntityDomainConfigurationResponseData) SetFqdnNormalizationEnabled(v bool) {
	o.FqdnNormalizationEnabled = v
}

// GetAwsNormalizationEnabled returns the AwsNormalizationEnabled field value
func (o *GetEntityDomainConfigurationResponseData) GetAwsNormalizationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AwsNormalizationEnabled
}

// GetAwsNormalizationEnabledOk returns a tuple with the AwsNormalizationEnabled field value
// and a boolean to check if the value has been set.
func (o *GetEntityDomainConfigurationResponseData) GetAwsNormalizationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsNormalizationEnabled, true
}

// SetAwsNormalizationEnabled sets field value
func (o *GetEntityDomainConfigurationResponseData) SetAwsNormalizationEnabled(v bool) {
	o.AwsNormalizationEnabled = v
}

func (o GetEntityDomainConfigurationResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEntityDomainConfigurationResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["normalizeUsernames"] = o.NormalizeUsernames
	toSerialize["normalizeHostnames"] = o.NormalizeHostnames
	if !IsNil(o.DefaultNormalizedDomain) {
		toSerialize["defaultNormalizedDomain"] = o.DefaultNormalizedDomain
	}
	toSerialize["domainMappings"] = o.DomainMappings
	toSerialize["adDomainNormalizationEnabled"] = o.AdDomainNormalizationEnabled
	toSerialize["fqdnNormalizationEnabled"] = o.FqdnNormalizationEnabled
	toSerialize["awsNormalizationEnabled"] = o.AwsNormalizationEnabled
	return toSerialize, nil
}

func (o *GetEntityDomainConfigurationResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"normalizeUsernames",
		"normalizeHostnames",
		"domainMappings",
		"adDomainNormalizationEnabled",
		"fqdnNormalizationEnabled",
		"awsNormalizationEnabled",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetEntityDomainConfigurationResponseData := _GetEntityDomainConfigurationResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEntityDomainConfigurationResponseData)

	if err != nil {
		return err
	}

	*o = GetEntityDomainConfigurationResponseData(varGetEntityDomainConfigurationResponseData)

	return err
}

type NullableGetEntityDomainConfigurationResponseData struct {
	value *GetEntityDomainConfigurationResponseData
	isSet bool
}

func (v NullableGetEntityDomainConfigurationResponseData) Get() *GetEntityDomainConfigurationResponseData {
	return v.value
}

func (v *NullableGetEntityDomainConfigurationResponseData) Set(val *GetEntityDomainConfigurationResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEntityDomainConfigurationResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEntityDomainConfigurationResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEntityDomainConfigurationResponseData(val *GetEntityDomainConfigurationResponseData) *NullableGetEntityDomainConfigurationResponseData {
	return &NullableGetEntityDomainConfigurationResponseData{value: val, isSet: true}
}

func (v NullableGetEntityDomainConfigurationResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEntityDomainConfigurationResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

