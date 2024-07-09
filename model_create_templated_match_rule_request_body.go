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

// checks if the CreateTemplatedMatchRuleRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTemplatedMatchRuleRequestBody{}

// CreateTemplatedMatchRuleRequestBody struct for CreateTemplatedMatchRuleRequestBody
type CreateTemplatedMatchRuleRequestBody struct {
	Fields CreateTemplatedMatchRuleRequestBodyFields `json:"fields"`
}

type _CreateTemplatedMatchRuleRequestBody CreateTemplatedMatchRuleRequestBody

// NewCreateTemplatedMatchRuleRequestBody instantiates a new CreateTemplatedMatchRuleRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTemplatedMatchRuleRequestBody(fields CreateTemplatedMatchRuleRequestBodyFields) *CreateTemplatedMatchRuleRequestBody {
	this := CreateTemplatedMatchRuleRequestBody{}
	this.Fields = fields
	return &this
}

// NewCreateTemplatedMatchRuleRequestBodyWithDefaults instantiates a new CreateTemplatedMatchRuleRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTemplatedMatchRuleRequestBodyWithDefaults() *CreateTemplatedMatchRuleRequestBody {
	this := CreateTemplatedMatchRuleRequestBody{}
	return &this
}

// GetFields returns the Fields field value
func (o *CreateTemplatedMatchRuleRequestBody) GetFields() CreateTemplatedMatchRuleRequestBodyFields {
	if o == nil {
		var ret CreateTemplatedMatchRuleRequestBodyFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *CreateTemplatedMatchRuleRequestBody) GetFieldsOk() (*CreateTemplatedMatchRuleRequestBodyFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *CreateTemplatedMatchRuleRequestBody) SetFields(v CreateTemplatedMatchRuleRequestBodyFields) {
	o.Fields = v
}

func (o CreateTemplatedMatchRuleRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTemplatedMatchRuleRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *CreateTemplatedMatchRuleRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varCreateTemplatedMatchRuleRequestBody := _CreateTemplatedMatchRuleRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTemplatedMatchRuleRequestBody)

	if err != nil {
		return err
	}

	*o = CreateTemplatedMatchRuleRequestBody(varCreateTemplatedMatchRuleRequestBody)

	return err
}

type NullableCreateTemplatedMatchRuleRequestBody struct {
	value *CreateTemplatedMatchRuleRequestBody
	isSet bool
}

func (v NullableCreateTemplatedMatchRuleRequestBody) Get() *CreateTemplatedMatchRuleRequestBody {
	return v.value
}

func (v *NullableCreateTemplatedMatchRuleRequestBody) Set(val *CreateTemplatedMatchRuleRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTemplatedMatchRuleRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTemplatedMatchRuleRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTemplatedMatchRuleRequestBody(val *CreateTemplatedMatchRuleRequestBody) *NullableCreateTemplatedMatchRuleRequestBody {
	return &NullableCreateTemplatedMatchRuleRequestBody{value: val, isSet: true}
}

func (v NullableCreateTemplatedMatchRuleRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTemplatedMatchRuleRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


