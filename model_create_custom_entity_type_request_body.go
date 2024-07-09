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

// checks if the CreateCustomEntityTypeRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomEntityTypeRequestBody{}

// CreateCustomEntityTypeRequestBody struct for CreateCustomEntityTypeRequestBody
type CreateCustomEntityTypeRequestBody struct {
	Fields CreateCustomEntityTypeRequestBodyFields `json:"fields"`
}

type _CreateCustomEntityTypeRequestBody CreateCustomEntityTypeRequestBody

// NewCreateCustomEntityTypeRequestBody instantiates a new CreateCustomEntityTypeRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomEntityTypeRequestBody(fields CreateCustomEntityTypeRequestBodyFields) *CreateCustomEntityTypeRequestBody {
	this := CreateCustomEntityTypeRequestBody{}
	this.Fields = fields
	return &this
}

// NewCreateCustomEntityTypeRequestBodyWithDefaults instantiates a new CreateCustomEntityTypeRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomEntityTypeRequestBodyWithDefaults() *CreateCustomEntityTypeRequestBody {
	this := CreateCustomEntityTypeRequestBody{}
	return &this
}

// GetFields returns the Fields field value
func (o *CreateCustomEntityTypeRequestBody) GetFields() CreateCustomEntityTypeRequestBodyFields {
	if o == nil {
		var ret CreateCustomEntityTypeRequestBodyFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *CreateCustomEntityTypeRequestBody) GetFieldsOk() (*CreateCustomEntityTypeRequestBodyFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *CreateCustomEntityTypeRequestBody) SetFields(v CreateCustomEntityTypeRequestBodyFields) {
	o.Fields = v
}

func (o CreateCustomEntityTypeRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomEntityTypeRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *CreateCustomEntityTypeRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varCreateCustomEntityTypeRequestBody := _CreateCustomEntityTypeRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateCustomEntityTypeRequestBody)

	if err != nil {
		return err
	}

	*o = CreateCustomEntityTypeRequestBody(varCreateCustomEntityTypeRequestBody)

	return err
}

type NullableCreateCustomEntityTypeRequestBody struct {
	value *CreateCustomEntityTypeRequestBody
	isSet bool
}

func (v NullableCreateCustomEntityTypeRequestBody) Get() *CreateCustomEntityTypeRequestBody {
	return v.value
}

func (v *NullableCreateCustomEntityTypeRequestBody) Set(val *CreateCustomEntityTypeRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomEntityTypeRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomEntityTypeRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomEntityTypeRequestBody(val *CreateCustomEntityTypeRequestBody) *NullableCreateCustomEntityTypeRequestBody {
	return &NullableCreateCustomEntityTypeRequestBody{value: val, isSet: true}
}

func (v NullableCreateCustomEntityTypeRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomEntityTypeRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


