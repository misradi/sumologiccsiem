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

// checks if the OverrideTemplatedMatchRuleRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverrideTemplatedMatchRuleRequestBody{}

// OverrideTemplatedMatchRuleRequestBody struct for OverrideTemplatedMatchRuleRequestBody
type OverrideTemplatedMatchRuleRequestBody struct {
	Fields OverrideTemplatedMatchRuleRequestBodyFields `json:"fields"`
}

type _OverrideTemplatedMatchRuleRequestBody OverrideTemplatedMatchRuleRequestBody

// NewOverrideTemplatedMatchRuleRequestBody instantiates a new OverrideTemplatedMatchRuleRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideTemplatedMatchRuleRequestBody(fields OverrideTemplatedMatchRuleRequestBodyFields) *OverrideTemplatedMatchRuleRequestBody {
	this := OverrideTemplatedMatchRuleRequestBody{}
	this.Fields = fields
	return &this
}

// NewOverrideTemplatedMatchRuleRequestBodyWithDefaults instantiates a new OverrideTemplatedMatchRuleRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideTemplatedMatchRuleRequestBodyWithDefaults() *OverrideTemplatedMatchRuleRequestBody {
	this := OverrideTemplatedMatchRuleRequestBody{}
	return &this
}

// GetFields returns the Fields field value
func (o *OverrideTemplatedMatchRuleRequestBody) GetFields() OverrideTemplatedMatchRuleRequestBodyFields {
	if o == nil {
		var ret OverrideTemplatedMatchRuleRequestBodyFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *OverrideTemplatedMatchRuleRequestBody) GetFieldsOk() (*OverrideTemplatedMatchRuleRequestBodyFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *OverrideTemplatedMatchRuleRequestBody) SetFields(v OverrideTemplatedMatchRuleRequestBodyFields) {
	o.Fields = v
}

func (o OverrideTemplatedMatchRuleRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideTemplatedMatchRuleRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *OverrideTemplatedMatchRuleRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varOverrideTemplatedMatchRuleRequestBody := _OverrideTemplatedMatchRuleRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOverrideTemplatedMatchRuleRequestBody)

	if err != nil {
		return err
	}

	*o = OverrideTemplatedMatchRuleRequestBody(varOverrideTemplatedMatchRuleRequestBody)

	return err
}

type NullableOverrideTemplatedMatchRuleRequestBody struct {
	value *OverrideTemplatedMatchRuleRequestBody
	isSet bool
}

func (v NullableOverrideTemplatedMatchRuleRequestBody) Get() *OverrideTemplatedMatchRuleRequestBody {
	return v.value
}

func (v *NullableOverrideTemplatedMatchRuleRequestBody) Set(val *OverrideTemplatedMatchRuleRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideTemplatedMatchRuleRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideTemplatedMatchRuleRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideTemplatedMatchRuleRequestBody(val *OverrideTemplatedMatchRuleRequestBody) *NullableOverrideTemplatedMatchRuleRequestBody {
	return &NullableOverrideTemplatedMatchRuleRequestBody{value: val, isSet: true}
}

func (v NullableOverrideTemplatedMatchRuleRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideTemplatedMatchRuleRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


