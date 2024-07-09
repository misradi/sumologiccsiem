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

// checks if the UpdateCustomMatchListColumnRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCustomMatchListColumnRequestBodyFields{}

// UpdateCustomMatchListColumnRequestBodyFields struct for UpdateCustomMatchListColumnRequestBodyFields
type UpdateCustomMatchListColumnRequestBodyFields struct {
	Name string `json:"name"`
	Fields []string `json:"fields"`
}

type _UpdateCustomMatchListColumnRequestBodyFields UpdateCustomMatchListColumnRequestBodyFields

// NewUpdateCustomMatchListColumnRequestBodyFields instantiates a new UpdateCustomMatchListColumnRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCustomMatchListColumnRequestBodyFields(name string, fields []string) *UpdateCustomMatchListColumnRequestBodyFields {
	this := UpdateCustomMatchListColumnRequestBodyFields{}
	this.Name = name
	this.Fields = fields
	return &this
}

// NewUpdateCustomMatchListColumnRequestBodyFieldsWithDefaults instantiates a new UpdateCustomMatchListColumnRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCustomMatchListColumnRequestBodyFieldsWithDefaults() *UpdateCustomMatchListColumnRequestBodyFields {
	this := UpdateCustomMatchListColumnRequestBodyFields{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateCustomMatchListColumnRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateCustomMatchListColumnRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateCustomMatchListColumnRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetFields returns the Fields field value
func (o *UpdateCustomMatchListColumnRequestBodyFields) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *UpdateCustomMatchListColumnRequestBodyFields) GetFieldsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *UpdateCustomMatchListColumnRequestBodyFields) SetFields(v []string) {
	o.Fields = v
}

func (o UpdateCustomMatchListColumnRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCustomMatchListColumnRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *UpdateCustomMatchListColumnRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateCustomMatchListColumnRequestBodyFields := _UpdateCustomMatchListColumnRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateCustomMatchListColumnRequestBodyFields)

	if err != nil {
		return err
	}

	*o = UpdateCustomMatchListColumnRequestBodyFields(varUpdateCustomMatchListColumnRequestBodyFields)

	return err
}

type NullableUpdateCustomMatchListColumnRequestBodyFields struct {
	value *UpdateCustomMatchListColumnRequestBodyFields
	isSet bool
}

func (v NullableUpdateCustomMatchListColumnRequestBodyFields) Get() *UpdateCustomMatchListColumnRequestBodyFields {
	return v.value
}

func (v *NullableUpdateCustomMatchListColumnRequestBodyFields) Set(val *UpdateCustomMatchListColumnRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCustomMatchListColumnRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCustomMatchListColumnRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCustomMatchListColumnRequestBodyFields(val *UpdateCustomMatchListColumnRequestBodyFields) *NullableUpdateCustomMatchListColumnRequestBodyFields {
	return &NullableUpdateCustomMatchListColumnRequestBodyFields{value: val, isSet: true}
}

func (v NullableUpdateCustomMatchListColumnRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCustomMatchListColumnRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


