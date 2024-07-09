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

// checks if the CreateInventoryRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInventoryRequestBody{}

// CreateInventoryRequestBody struct for CreateInventoryRequestBody
type CreateInventoryRequestBody struct {
	Fields CreateInventoryRequestBodyFields `json:"fields"`
}

type _CreateInventoryRequestBody CreateInventoryRequestBody

// NewCreateInventoryRequestBody instantiates a new CreateInventoryRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInventoryRequestBody(fields CreateInventoryRequestBodyFields) *CreateInventoryRequestBody {
	this := CreateInventoryRequestBody{}
	this.Fields = fields
	return &this
}

// NewCreateInventoryRequestBodyWithDefaults instantiates a new CreateInventoryRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInventoryRequestBodyWithDefaults() *CreateInventoryRequestBody {
	this := CreateInventoryRequestBody{}
	return &this
}

// GetFields returns the Fields field value
func (o *CreateInventoryRequestBody) GetFields() CreateInventoryRequestBodyFields {
	if o == nil {
		var ret CreateInventoryRequestBodyFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *CreateInventoryRequestBody) GetFieldsOk() (*CreateInventoryRequestBodyFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *CreateInventoryRequestBody) SetFields(v CreateInventoryRequestBodyFields) {
	o.Fields = v
}

func (o CreateInventoryRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInventoryRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *CreateInventoryRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varCreateInventoryRequestBody := _CreateInventoryRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateInventoryRequestBody)

	if err != nil {
		return err
	}

	*o = CreateInventoryRequestBody(varCreateInventoryRequestBody)

	return err
}

type NullableCreateInventoryRequestBody struct {
	value *CreateInventoryRequestBody
	isSet bool
}

func (v NullableCreateInventoryRequestBody) Get() *CreateInventoryRequestBody {
	return v.value
}

func (v *NullableCreateInventoryRequestBody) Set(val *CreateInventoryRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInventoryRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInventoryRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInventoryRequestBody(val *CreateInventoryRequestBody) *NullableCreateInventoryRequestBody {
	return &NullableCreateInventoryRequestBody{value: val, isSet: true}
}

func (v NullableCreateInventoryRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInventoryRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

