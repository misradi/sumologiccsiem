/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologiccsiem

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OverrideThresholdRuleRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverrideThresholdRuleRequestBody{}

// OverrideThresholdRuleRequestBody struct for OverrideThresholdRuleRequestBody
type OverrideThresholdRuleRequestBody struct {
	Fields OverrideThresholdRuleRequestBodyFields `json:"fields"`
}

type _OverrideThresholdRuleRequestBody OverrideThresholdRuleRequestBody

// NewOverrideThresholdRuleRequestBody instantiates a new OverrideThresholdRuleRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideThresholdRuleRequestBody(fields OverrideThresholdRuleRequestBodyFields) *OverrideThresholdRuleRequestBody {
	this := OverrideThresholdRuleRequestBody{}
	this.Fields = fields
	return &this
}

// NewOverrideThresholdRuleRequestBodyWithDefaults instantiates a new OverrideThresholdRuleRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideThresholdRuleRequestBodyWithDefaults() *OverrideThresholdRuleRequestBody {
	this := OverrideThresholdRuleRequestBody{}
	return &this
}

// GetFields returns the Fields field value
func (o *OverrideThresholdRuleRequestBody) GetFields() OverrideThresholdRuleRequestBodyFields {
	if o == nil {
		var ret OverrideThresholdRuleRequestBodyFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *OverrideThresholdRuleRequestBody) GetFieldsOk() (*OverrideThresholdRuleRequestBodyFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *OverrideThresholdRuleRequestBody) SetFields(v OverrideThresholdRuleRequestBodyFields) {
	o.Fields = v
}

func (o OverrideThresholdRuleRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideThresholdRuleRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *OverrideThresholdRuleRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varOverrideThresholdRuleRequestBody := _OverrideThresholdRuleRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOverrideThresholdRuleRequestBody)

	if err != nil {
		return err
	}

	*o = OverrideThresholdRuleRequestBody(varOverrideThresholdRuleRequestBody)

	return err
}

type NullableOverrideThresholdRuleRequestBody struct {
	value *OverrideThresholdRuleRequestBody
	isSet bool
}

func (v NullableOverrideThresholdRuleRequestBody) Get() *OverrideThresholdRuleRequestBody {
	return v.value
}

func (v *NullableOverrideThresholdRuleRequestBody) Set(val *OverrideThresholdRuleRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideThresholdRuleRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideThresholdRuleRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideThresholdRuleRequestBody(val *OverrideThresholdRuleRequestBody) *NullableOverrideThresholdRuleRequestBody {
	return &NullableOverrideThresholdRuleRequestBody{value: val, isSet: true}
}

func (v NullableOverrideThresholdRuleRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideThresholdRuleRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


