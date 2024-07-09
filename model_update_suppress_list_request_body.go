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

// checks if the UpdateSuppressListRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSuppressListRequestBody{}

// UpdateSuppressListRequestBody struct for UpdateSuppressListRequestBody
type UpdateSuppressListRequestBody struct {
	Fields UpdateSuppressListRequestBodyFields `json:"fields"`
}

type _UpdateSuppressListRequestBody UpdateSuppressListRequestBody

// NewUpdateSuppressListRequestBody instantiates a new UpdateSuppressListRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSuppressListRequestBody(fields UpdateSuppressListRequestBodyFields) *UpdateSuppressListRequestBody {
	this := UpdateSuppressListRequestBody{}
	this.Fields = fields
	return &this
}

// NewUpdateSuppressListRequestBodyWithDefaults instantiates a new UpdateSuppressListRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSuppressListRequestBodyWithDefaults() *UpdateSuppressListRequestBody {
	this := UpdateSuppressListRequestBody{}
	return &this
}

// GetFields returns the Fields field value
func (o *UpdateSuppressListRequestBody) GetFields() UpdateSuppressListRequestBodyFields {
	if o == nil {
		var ret UpdateSuppressListRequestBodyFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *UpdateSuppressListRequestBody) GetFieldsOk() (*UpdateSuppressListRequestBodyFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *UpdateSuppressListRequestBody) SetFields(v UpdateSuppressListRequestBodyFields) {
	o.Fields = v
}

func (o UpdateSuppressListRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSuppressListRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *UpdateSuppressListRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateSuppressListRequestBody := _UpdateSuppressListRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateSuppressListRequestBody)

	if err != nil {
		return err
	}

	*o = UpdateSuppressListRequestBody(varUpdateSuppressListRequestBody)

	return err
}

type NullableUpdateSuppressListRequestBody struct {
	value *UpdateSuppressListRequestBody
	isSet bool
}

func (v NullableUpdateSuppressListRequestBody) Get() *UpdateSuppressListRequestBody {
	return v.value
}

func (v *NullableUpdateSuppressListRequestBody) Set(val *UpdateSuppressListRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSuppressListRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSuppressListRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSuppressListRequestBody(val *UpdateSuppressListRequestBody) *NullableUpdateSuppressListRequestBody {
	return &NullableUpdateSuppressListRequestBody{value: val, isSet: true}
}

func (v NullableUpdateSuppressListRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSuppressListRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


