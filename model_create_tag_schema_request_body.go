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

// checks if the CreateTagSchemaRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTagSchemaRequestBody{}

// CreateTagSchemaRequestBody struct for CreateTagSchemaRequestBody
type CreateTagSchemaRequestBody struct {
	Fields UpdateTagSchemaRequestBodyFields `json:"fields"`
}

type _CreateTagSchemaRequestBody CreateTagSchemaRequestBody

// NewCreateTagSchemaRequestBody instantiates a new CreateTagSchemaRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTagSchemaRequestBody(fields UpdateTagSchemaRequestBodyFields) *CreateTagSchemaRequestBody {
	this := CreateTagSchemaRequestBody{}
	this.Fields = fields
	return &this
}

// NewCreateTagSchemaRequestBodyWithDefaults instantiates a new CreateTagSchemaRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTagSchemaRequestBodyWithDefaults() *CreateTagSchemaRequestBody {
	this := CreateTagSchemaRequestBody{}
	return &this
}

// GetFields returns the Fields field value
func (o *CreateTagSchemaRequestBody) GetFields() UpdateTagSchemaRequestBodyFields {
	if o == nil {
		var ret UpdateTagSchemaRequestBodyFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *CreateTagSchemaRequestBody) GetFieldsOk() (*UpdateTagSchemaRequestBodyFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *CreateTagSchemaRequestBody) SetFields(v UpdateTagSchemaRequestBodyFields) {
	o.Fields = v
}

func (o CreateTagSchemaRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTagSchemaRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *CreateTagSchemaRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varCreateTagSchemaRequestBody := _CreateTagSchemaRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTagSchemaRequestBody)

	if err != nil {
		return err
	}

	*o = CreateTagSchemaRequestBody(varCreateTagSchemaRequestBody)

	return err
}

type NullableCreateTagSchemaRequestBody struct {
	value *CreateTagSchemaRequestBody
	isSet bool
}

func (v NullableCreateTagSchemaRequestBody) Get() *CreateTagSchemaRequestBody {
	return v.value
}

func (v *NullableCreateTagSchemaRequestBody) Set(val *CreateTagSchemaRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTagSchemaRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTagSchemaRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTagSchemaRequestBody(val *CreateTagSchemaRequestBody) *NullableCreateTagSchemaRequestBody {
	return &NullableCreateTagSchemaRequestBody{value: val, isSet: true}
}

func (v NullableCreateTagSchemaRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTagSchemaRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


