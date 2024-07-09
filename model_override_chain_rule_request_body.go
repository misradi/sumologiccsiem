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

// checks if the OverrideChainRuleRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverrideChainRuleRequestBody{}

// OverrideChainRuleRequestBody struct for OverrideChainRuleRequestBody
type OverrideChainRuleRequestBody struct {
	Fields OverrideChainRuleRequestBodyFields `json:"fields"`
}

type _OverrideChainRuleRequestBody OverrideChainRuleRequestBody

// NewOverrideChainRuleRequestBody instantiates a new OverrideChainRuleRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideChainRuleRequestBody(fields OverrideChainRuleRequestBodyFields) *OverrideChainRuleRequestBody {
	this := OverrideChainRuleRequestBody{}
	this.Fields = fields
	return &this
}

// NewOverrideChainRuleRequestBodyWithDefaults instantiates a new OverrideChainRuleRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideChainRuleRequestBodyWithDefaults() *OverrideChainRuleRequestBody {
	this := OverrideChainRuleRequestBody{}
	return &this
}

// GetFields returns the Fields field value
func (o *OverrideChainRuleRequestBody) GetFields() OverrideChainRuleRequestBodyFields {
	if o == nil {
		var ret OverrideChainRuleRequestBodyFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *OverrideChainRuleRequestBody) GetFieldsOk() (*OverrideChainRuleRequestBodyFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *OverrideChainRuleRequestBody) SetFields(v OverrideChainRuleRequestBodyFields) {
	o.Fields = v
}

func (o OverrideChainRuleRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideChainRuleRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *OverrideChainRuleRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varOverrideChainRuleRequestBody := _OverrideChainRuleRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOverrideChainRuleRequestBody)

	if err != nil {
		return err
	}

	*o = OverrideChainRuleRequestBody(varOverrideChainRuleRequestBody)

	return err
}

type NullableOverrideChainRuleRequestBody struct {
	value *OverrideChainRuleRequestBody
	isSet bool
}

func (v NullableOverrideChainRuleRequestBody) Get() *OverrideChainRuleRequestBody {
	return v.value
}

func (v *NullableOverrideChainRuleRequestBody) Set(val *OverrideChainRuleRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideChainRuleRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideChainRuleRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideChainRuleRequestBody(val *OverrideChainRuleRequestBody) *NullableOverrideChainRuleRequestBody {
	return &NullableOverrideChainRuleRequestBody{value: val, isSet: true}
}

func (v NullableOverrideChainRuleRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideChainRuleRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


