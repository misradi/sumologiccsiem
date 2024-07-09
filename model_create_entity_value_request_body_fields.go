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

// checks if the CreateEntityValueRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEntityValueRequestBodyFields{}

// CreateEntityValueRequestBodyFields struct for CreateEntityValueRequestBodyFields
type CreateEntityValueRequestBodyFields struct {
	Name string `json:"name"`
	ConfigurationType string `json:"configurationType"`
	Description *string `json:"description,omitempty"`
	EntityNamespace *string `json:"entityNamespace,omitempty"`
	EntityType *string `json:"entityType,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
	Suffix *string `json:"suffix,omitempty"`
	NetworkBlock *string `json:"networkBlock,omitempty"`
	Criticality *string `json:"criticality,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Suppressed *bool `json:"suppressed,omitempty"`
}

type _CreateEntityValueRequestBodyFields CreateEntityValueRequestBodyFields

// NewCreateEntityValueRequestBodyFields instantiates a new CreateEntityValueRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEntityValueRequestBodyFields(name string, configurationType string) *CreateEntityValueRequestBodyFields {
	this := CreateEntityValueRequestBodyFields{}
	this.Name = name
	this.ConfigurationType = configurationType
	return &this
}

// NewCreateEntityValueRequestBodyFieldsWithDefaults instantiates a new CreateEntityValueRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEntityValueRequestBodyFieldsWithDefaults() *CreateEntityValueRequestBodyFields {
	this := CreateEntityValueRequestBodyFields{}
	return &this
}

// GetName returns the Name field value
func (o *CreateEntityValueRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateEntityValueRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateEntityValueRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetConfigurationType returns the ConfigurationType field value
func (o *CreateEntityValueRequestBodyFields) GetConfigurationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigurationType
}

// GetConfigurationTypeOk returns a tuple with the ConfigurationType field value
// and a boolean to check if the value has been set.
func (o *CreateEntityValueRequestBodyFields) GetConfigurationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationType, true
}

// SetConfigurationType sets field value
func (o *CreateEntityValueRequestBodyFields) SetConfigurationType(v string) {
	o.ConfigurationType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateEntityValueRequestBodyFields) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEntityValueRequestBodyFields) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateEntityValueRequestBodyFields) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateEntityValueRequestBodyFields) SetDescription(v string) {
	o.Description = &v
}

// GetEntityNamespace returns the EntityNamespace field value if set, zero value otherwise.
func (o *CreateEntityValueRequestBodyFields) GetEntityNamespace() string {
	if o == nil || IsNil(o.EntityNamespace) {
		var ret string
		return ret
	}
	return *o.EntityNamespace
}

// GetEntityNamespaceOk returns a tuple with the EntityNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEntityValueRequestBodyFields) GetEntityNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.EntityNamespace) {
		return nil, false
	}
	return o.EntityNamespace, true
}

// HasEntityNamespace returns a boolean if a field has been set.
func (o *CreateEntityValueRequestBodyFields) HasEntityNamespace() bool {
	if o != nil && !IsNil(o.EntityNamespace) {
		return true
	}

	return false
}

// SetEntityNamespace gets a reference to the given string and assigns it to the EntityNamespace field.
func (o *CreateEntityValueRequestBodyFields) SetEntityNamespace(v string) {
	o.EntityNamespace = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *CreateEntityValueRequestBodyFields) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEntityValueRequestBodyFields) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *CreateEntityValueRequestBodyFields) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *CreateEntityValueRequestBodyFields) SetEntityType(v string) {
	o.EntityType = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *CreateEntityValueRequestBodyFields) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEntityValueRequestBodyFields) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *CreateEntityValueRequestBodyFields) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *CreateEntityValueRequestBodyFields) SetPrefix(v string) {
	o.Prefix = &v
}

// GetSuffix returns the Suffix field value if set, zero value otherwise.
func (o *CreateEntityValueRequestBodyFields) GetSuffix() string {
	if o == nil || IsNil(o.Suffix) {
		var ret string
		return ret
	}
	return *o.Suffix
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEntityValueRequestBodyFields) GetSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.Suffix) {
		return nil, false
	}
	return o.Suffix, true
}

// HasSuffix returns a boolean if a field has been set.
func (o *CreateEntityValueRequestBodyFields) HasSuffix() bool {
	if o != nil && !IsNil(o.Suffix) {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given string and assigns it to the Suffix field.
func (o *CreateEntityValueRequestBodyFields) SetSuffix(v string) {
	o.Suffix = &v
}

// GetNetworkBlock returns the NetworkBlock field value if set, zero value otherwise.
func (o *CreateEntityValueRequestBodyFields) GetNetworkBlock() string {
	if o == nil || IsNil(o.NetworkBlock) {
		var ret string
		return ret
	}
	return *o.NetworkBlock
}

// GetNetworkBlockOk returns a tuple with the NetworkBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEntityValueRequestBodyFields) GetNetworkBlockOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkBlock) {
		return nil, false
	}
	return o.NetworkBlock, true
}

// HasNetworkBlock returns a boolean if a field has been set.
func (o *CreateEntityValueRequestBodyFields) HasNetworkBlock() bool {
	if o != nil && !IsNil(o.NetworkBlock) {
		return true
	}

	return false
}

// SetNetworkBlock gets a reference to the given string and assigns it to the NetworkBlock field.
func (o *CreateEntityValueRequestBodyFields) SetNetworkBlock(v string) {
	o.NetworkBlock = &v
}

// GetCriticality returns the Criticality field value if set, zero value otherwise.
func (o *CreateEntityValueRequestBodyFields) GetCriticality() string {
	if o == nil || IsNil(o.Criticality) {
		var ret string
		return ret
	}
	return *o.Criticality
}

// GetCriticalityOk returns a tuple with the Criticality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEntityValueRequestBodyFields) GetCriticalityOk() (*string, bool) {
	if o == nil || IsNil(o.Criticality) {
		return nil, false
	}
	return o.Criticality, true
}

// HasCriticality returns a boolean if a field has been set.
func (o *CreateEntityValueRequestBodyFields) HasCriticality() bool {
	if o != nil && !IsNil(o.Criticality) {
		return true
	}

	return false
}

// SetCriticality gets a reference to the given string and assigns it to the Criticality field.
func (o *CreateEntityValueRequestBodyFields) SetCriticality(v string) {
	o.Criticality = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateEntityValueRequestBodyFields) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEntityValueRequestBodyFields) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateEntityValueRequestBodyFields) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateEntityValueRequestBodyFields) SetTags(v []string) {
	o.Tags = v
}

// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *CreateEntityValueRequestBodyFields) GetSuppressed() bool {
	if o == nil || IsNil(o.Suppressed) {
		var ret bool
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEntityValueRequestBodyFields) GetSuppressedOk() (*bool, bool) {
	if o == nil || IsNil(o.Suppressed) {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *CreateEntityValueRequestBodyFields) HasSuppressed() bool {
	if o != nil && !IsNil(o.Suppressed) {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given bool and assigns it to the Suppressed field.
func (o *CreateEntityValueRequestBodyFields) SetSuppressed(v bool) {
	o.Suppressed = &v
}

func (o CreateEntityValueRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEntityValueRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["configurationType"] = o.ConfigurationType
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EntityNamespace) {
		toSerialize["entityNamespace"] = o.EntityNamespace
	}
	if !IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Suffix) {
		toSerialize["suffix"] = o.Suffix
	}
	if !IsNil(o.NetworkBlock) {
		toSerialize["networkBlock"] = o.NetworkBlock
	}
	if !IsNil(o.Criticality) {
		toSerialize["criticality"] = o.Criticality
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Suppressed) {
		toSerialize["suppressed"] = o.Suppressed
	}
	return toSerialize, nil
}

func (o *CreateEntityValueRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"configurationType",
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

	varCreateEntityValueRequestBodyFields := _CreateEntityValueRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateEntityValueRequestBodyFields)

	if err != nil {
		return err
	}

	*o = CreateEntityValueRequestBodyFields(varCreateEntityValueRequestBodyFields)

	return err
}

type NullableCreateEntityValueRequestBodyFields struct {
	value *CreateEntityValueRequestBodyFields
	isSet bool
}

func (v NullableCreateEntityValueRequestBodyFields) Get() *CreateEntityValueRequestBodyFields {
	return v.value
}

func (v *NullableCreateEntityValueRequestBodyFields) Set(val *CreateEntityValueRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEntityValueRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEntityValueRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEntityValueRequestBodyFields(val *CreateEntityValueRequestBodyFields) *NullableCreateEntityValueRequestBodyFields {
	return &NullableCreateEntityValueRequestBodyFields{value: val, isSet: true}
}

func (v NullableCreateEntityValueRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEntityValueRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


