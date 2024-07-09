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

// checks if the OverrideFirstSeenRuleRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverrideFirstSeenRuleRequestBody{}

// OverrideFirstSeenRuleRequestBody struct for OverrideFirstSeenRuleRequestBody
type OverrideFirstSeenRuleRequestBody struct {
	Fields OverrideFirstSeenRuleRequestBodyFields `json:"fields"`
}

type _OverrideFirstSeenRuleRequestBody OverrideFirstSeenRuleRequestBody

// NewOverrideFirstSeenRuleRequestBody instantiates a new OverrideFirstSeenRuleRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideFirstSeenRuleRequestBody(fields OverrideFirstSeenRuleRequestBodyFields) *OverrideFirstSeenRuleRequestBody {
	this := OverrideFirstSeenRuleRequestBody{}
	this.Fields = fields
	return &this
}

// NewOverrideFirstSeenRuleRequestBodyWithDefaults instantiates a new OverrideFirstSeenRuleRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideFirstSeenRuleRequestBodyWithDefaults() *OverrideFirstSeenRuleRequestBody {
	this := OverrideFirstSeenRuleRequestBody{}
	return &this
}

// GetFields returns the Fields field value
func (o *OverrideFirstSeenRuleRequestBody) GetFields() OverrideFirstSeenRuleRequestBodyFields {
	if o == nil {
		var ret OverrideFirstSeenRuleRequestBodyFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *OverrideFirstSeenRuleRequestBody) GetFieldsOk() (*OverrideFirstSeenRuleRequestBodyFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *OverrideFirstSeenRuleRequestBody) SetFields(v OverrideFirstSeenRuleRequestBodyFields) {
	o.Fields = v
}

func (o OverrideFirstSeenRuleRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideFirstSeenRuleRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *OverrideFirstSeenRuleRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fields",
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

	varOverrideFirstSeenRuleRequestBody := _OverrideFirstSeenRuleRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOverrideFirstSeenRuleRequestBody)

	if err != nil {
		return err
	}

	*o = OverrideFirstSeenRuleRequestBody(varOverrideFirstSeenRuleRequestBody)

	return err
}

type NullableOverrideFirstSeenRuleRequestBody struct {
	value *OverrideFirstSeenRuleRequestBody
	isSet bool
}

func (v NullableOverrideFirstSeenRuleRequestBody) Get() *OverrideFirstSeenRuleRequestBody {
	return v.value
}

func (v *NullableOverrideFirstSeenRuleRequestBody) Set(val *OverrideFirstSeenRuleRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideFirstSeenRuleRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideFirstSeenRuleRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideFirstSeenRuleRequestBody(val *OverrideFirstSeenRuleRequestBody) *NullableOverrideFirstSeenRuleRequestBody {
	return &NullableOverrideFirstSeenRuleRequestBody{value: val, isSet: true}
}

func (v NullableOverrideFirstSeenRuleRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideFirstSeenRuleRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

