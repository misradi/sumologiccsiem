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

// checks if the CreateThreatIntelSourceRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateThreatIntelSourceRequestBodyFields{}

// CreateThreatIntelSourceRequestBodyFields struct for CreateThreatIntelSourceRequestBodyFields
type CreateThreatIntelSourceRequestBodyFields struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
}

type _CreateThreatIntelSourceRequestBodyFields CreateThreatIntelSourceRequestBodyFields

// NewCreateThreatIntelSourceRequestBodyFields instantiates a new CreateThreatIntelSourceRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateThreatIntelSourceRequestBodyFields(name string) *CreateThreatIntelSourceRequestBodyFields {
	this := CreateThreatIntelSourceRequestBodyFields{}
	this.Name = name
	return &this
}

// NewCreateThreatIntelSourceRequestBodyFieldsWithDefaults instantiates a new CreateThreatIntelSourceRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateThreatIntelSourceRequestBodyFieldsWithDefaults() *CreateThreatIntelSourceRequestBodyFields {
	this := CreateThreatIntelSourceRequestBodyFields{}
	return &this
}

// GetName returns the Name field value
func (o *CreateThreatIntelSourceRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateThreatIntelSourceRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateThreatIntelSourceRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateThreatIntelSourceRequestBodyFields) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThreatIntelSourceRequestBodyFields) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateThreatIntelSourceRequestBodyFields) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateThreatIntelSourceRequestBodyFields) SetDescription(v string) {
	o.Description = &v
}

func (o CreateThreatIntelSourceRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateThreatIntelSourceRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *CreateThreatIntelSourceRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateThreatIntelSourceRequestBodyFields := _CreateThreatIntelSourceRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateThreatIntelSourceRequestBodyFields)

	if err != nil {
		return err
	}

	*o = CreateThreatIntelSourceRequestBodyFields(varCreateThreatIntelSourceRequestBodyFields)

	return err
}

type NullableCreateThreatIntelSourceRequestBodyFields struct {
	value *CreateThreatIntelSourceRequestBodyFields
	isSet bool
}

func (v NullableCreateThreatIntelSourceRequestBodyFields) Get() *CreateThreatIntelSourceRequestBodyFields {
	return v.value
}

func (v *NullableCreateThreatIntelSourceRequestBodyFields) Set(val *CreateThreatIntelSourceRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateThreatIntelSourceRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateThreatIntelSourceRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateThreatIntelSourceRequestBodyFields(val *CreateThreatIntelSourceRequestBodyFields) *NullableCreateThreatIntelSourceRequestBodyFields {
	return &NullableCreateThreatIntelSourceRequestBodyFields{value: val, isSet: true}
}

func (v NullableCreateThreatIntelSourceRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateThreatIntelSourceRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

