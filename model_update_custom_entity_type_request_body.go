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

// checks if the UpdateCustomEntityTypeRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCustomEntityTypeRequestBody{}

// UpdateCustomEntityTypeRequestBody struct for UpdateCustomEntityTypeRequestBody
type UpdateCustomEntityTypeRequestBody struct {
	Fields UpdateCustomEntityTypeRequestBodyFields `json:"fields"`
}

type _UpdateCustomEntityTypeRequestBody UpdateCustomEntityTypeRequestBody

// NewUpdateCustomEntityTypeRequestBody instantiates a new UpdateCustomEntityTypeRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCustomEntityTypeRequestBody(fields UpdateCustomEntityTypeRequestBodyFields) *UpdateCustomEntityTypeRequestBody {
	this := UpdateCustomEntityTypeRequestBody{}
	this.Fields = fields
	return &this
}

// NewUpdateCustomEntityTypeRequestBodyWithDefaults instantiates a new UpdateCustomEntityTypeRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCustomEntityTypeRequestBodyWithDefaults() *UpdateCustomEntityTypeRequestBody {
	this := UpdateCustomEntityTypeRequestBody{}
	return &this
}

// GetFields returns the Fields field value
func (o *UpdateCustomEntityTypeRequestBody) GetFields() UpdateCustomEntityTypeRequestBodyFields {
	if o == nil {
		var ret UpdateCustomEntityTypeRequestBodyFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *UpdateCustomEntityTypeRequestBody) GetFieldsOk() (*UpdateCustomEntityTypeRequestBodyFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *UpdateCustomEntityTypeRequestBody) SetFields(v UpdateCustomEntityTypeRequestBodyFields) {
	o.Fields = v
}

func (o UpdateCustomEntityTypeRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCustomEntityTypeRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *UpdateCustomEntityTypeRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateCustomEntityTypeRequestBody := _UpdateCustomEntityTypeRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateCustomEntityTypeRequestBody)

	if err != nil {
		return err
	}

	*o = UpdateCustomEntityTypeRequestBody(varUpdateCustomEntityTypeRequestBody)

	return err
}

type NullableUpdateCustomEntityTypeRequestBody struct {
	value *UpdateCustomEntityTypeRequestBody
	isSet bool
}

func (v NullableUpdateCustomEntityTypeRequestBody) Get() *UpdateCustomEntityTypeRequestBody {
	return v.value
}

func (v *NullableUpdateCustomEntityTypeRequestBody) Set(val *UpdateCustomEntityTypeRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCustomEntityTypeRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCustomEntityTypeRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCustomEntityTypeRequestBody(val *UpdateCustomEntityTypeRequestBody) *NullableUpdateCustomEntityTypeRequestBody {
	return &NullableUpdateCustomEntityTypeRequestBody{value: val, isSet: true}
}

func (v NullableUpdateCustomEntityTypeRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCustomEntityTypeRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


