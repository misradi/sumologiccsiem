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

// checks if the OverrideAggregationRuleRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverrideAggregationRuleRequestBody{}

// OverrideAggregationRuleRequestBody struct for OverrideAggregationRuleRequestBody
type OverrideAggregationRuleRequestBody struct {
	Fields OverrideAggregationRuleRequestBodyFields `json:"fields"`
}

type _OverrideAggregationRuleRequestBody OverrideAggregationRuleRequestBody

// NewOverrideAggregationRuleRequestBody instantiates a new OverrideAggregationRuleRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideAggregationRuleRequestBody(fields OverrideAggregationRuleRequestBodyFields) *OverrideAggregationRuleRequestBody {
	this := OverrideAggregationRuleRequestBody{}
	this.Fields = fields
	return &this
}

// NewOverrideAggregationRuleRequestBodyWithDefaults instantiates a new OverrideAggregationRuleRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideAggregationRuleRequestBodyWithDefaults() *OverrideAggregationRuleRequestBody {
	this := OverrideAggregationRuleRequestBody{}
	return &this
}

// GetFields returns the Fields field value
func (o *OverrideAggregationRuleRequestBody) GetFields() OverrideAggregationRuleRequestBodyFields {
	if o == nil {
		var ret OverrideAggregationRuleRequestBodyFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *OverrideAggregationRuleRequestBody) GetFieldsOk() (*OverrideAggregationRuleRequestBodyFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *OverrideAggregationRuleRequestBody) SetFields(v OverrideAggregationRuleRequestBodyFields) {
	o.Fields = v
}

func (o OverrideAggregationRuleRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideAggregationRuleRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *OverrideAggregationRuleRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varOverrideAggregationRuleRequestBody := _OverrideAggregationRuleRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOverrideAggregationRuleRequestBody)

	if err != nil {
		return err
	}

	*o = OverrideAggregationRuleRequestBody(varOverrideAggregationRuleRequestBody)

	return err
}

type NullableOverrideAggregationRuleRequestBody struct {
	value *OverrideAggregationRuleRequestBody
	isSet bool
}

func (v NullableOverrideAggregationRuleRequestBody) Get() *OverrideAggregationRuleRequestBody {
	return v.value
}

func (v *NullableOverrideAggregationRuleRequestBody) Set(val *OverrideAggregationRuleRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideAggregationRuleRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideAggregationRuleRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideAggregationRuleRequestBody(val *OverrideAggregationRuleRequestBody) *NullableOverrideAggregationRuleRequestBody {
	return &NullableOverrideAggregationRuleRequestBody{value: val, isSet: true}
}

func (v NullableOverrideAggregationRuleRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideAggregationRuleRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


