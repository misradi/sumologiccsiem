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

// checks if the UpdateCustomEntityTypeRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCustomEntityTypeRequestBodyFields{}

// UpdateCustomEntityTypeRequestBodyFields struct for UpdateCustomEntityTypeRequestBodyFields
type UpdateCustomEntityTypeRequestBodyFields struct {
	// Human friend and unique name. Examples: \"Ip Address\", \"Username\", \"Mac Address\".
	Name string `json:"name"`
	// Record schema fields. Examples: \"file_hash_md5\", \"file_hash_sha1\".
	Fields []string `json:"fields"`
}

type _UpdateCustomEntityTypeRequestBodyFields UpdateCustomEntityTypeRequestBodyFields

// NewUpdateCustomEntityTypeRequestBodyFields instantiates a new UpdateCustomEntityTypeRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCustomEntityTypeRequestBodyFields(name string, fields []string) *UpdateCustomEntityTypeRequestBodyFields {
	this := UpdateCustomEntityTypeRequestBodyFields{}
	this.Name = name
	this.Fields = fields
	return &this
}

// NewUpdateCustomEntityTypeRequestBodyFieldsWithDefaults instantiates a new UpdateCustomEntityTypeRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCustomEntityTypeRequestBodyFieldsWithDefaults() *UpdateCustomEntityTypeRequestBodyFields {
	this := UpdateCustomEntityTypeRequestBodyFields{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateCustomEntityTypeRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateCustomEntityTypeRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateCustomEntityTypeRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetFields returns the Fields field value
func (o *UpdateCustomEntityTypeRequestBodyFields) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *UpdateCustomEntityTypeRequestBodyFields) GetFieldsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *UpdateCustomEntityTypeRequestBodyFields) SetFields(v []string) {
	o.Fields = v
}

func (o UpdateCustomEntityTypeRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCustomEntityTypeRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *UpdateCustomEntityTypeRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varUpdateCustomEntityTypeRequestBodyFields := _UpdateCustomEntityTypeRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateCustomEntityTypeRequestBodyFields)

	if err != nil {
		return err
	}

	*o = UpdateCustomEntityTypeRequestBodyFields(varUpdateCustomEntityTypeRequestBodyFields)

	return err
}

type NullableUpdateCustomEntityTypeRequestBodyFields struct {
	value *UpdateCustomEntityTypeRequestBodyFields
	isSet bool
}

func (v NullableUpdateCustomEntityTypeRequestBodyFields) Get() *UpdateCustomEntityTypeRequestBodyFields {
	return v.value
}

func (v *NullableUpdateCustomEntityTypeRequestBodyFields) Set(val *UpdateCustomEntityTypeRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCustomEntityTypeRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCustomEntityTypeRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCustomEntityTypeRequestBodyFields(val *UpdateCustomEntityTypeRequestBodyFields) *NullableUpdateCustomEntityTypeRequestBodyFields {
	return &NullableUpdateCustomEntityTypeRequestBodyFields{value: val, isSet: true}
}

func (v NullableUpdateCustomEntityTypeRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCustomEntityTypeRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


