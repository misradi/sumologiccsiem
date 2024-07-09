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

// checks if the UpdateEntityValueRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEntityValueRequestBody{}

// UpdateEntityValueRequestBody struct for UpdateEntityValueRequestBody
type UpdateEntityValueRequestBody struct {
	Fields CreateEntityValueRequestBodyFields `json:"fields"`
}

type _UpdateEntityValueRequestBody UpdateEntityValueRequestBody

// NewUpdateEntityValueRequestBody instantiates a new UpdateEntityValueRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEntityValueRequestBody(fields CreateEntityValueRequestBodyFields) *UpdateEntityValueRequestBody {
	this := UpdateEntityValueRequestBody{}
	this.Fields = fields
	return &this
}

// NewUpdateEntityValueRequestBodyWithDefaults instantiates a new UpdateEntityValueRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEntityValueRequestBodyWithDefaults() *UpdateEntityValueRequestBody {
	this := UpdateEntityValueRequestBody{}
	return &this
}

// GetFields returns the Fields field value
func (o *UpdateEntityValueRequestBody) GetFields() CreateEntityValueRequestBodyFields {
	if o == nil {
		var ret CreateEntityValueRequestBodyFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *UpdateEntityValueRequestBody) GetFieldsOk() (*CreateEntityValueRequestBodyFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *UpdateEntityValueRequestBody) SetFields(v CreateEntityValueRequestBodyFields) {
	o.Fields = v
}

func (o UpdateEntityValueRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEntityValueRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *UpdateEntityValueRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateEntityValueRequestBody := _UpdateEntityValueRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateEntityValueRequestBody)

	if err != nil {
		return err
	}

	*o = UpdateEntityValueRequestBody(varUpdateEntityValueRequestBody)

	return err
}

type NullableUpdateEntityValueRequestBody struct {
	value *UpdateEntityValueRequestBody
	isSet bool
}

func (v NullableUpdateEntityValueRequestBody) Get() *UpdateEntityValueRequestBody {
	return v.value
}

func (v *NullableUpdateEntityValueRequestBody) Set(val *UpdateEntityValueRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEntityValueRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEntityValueRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEntityValueRequestBody(val *UpdateEntityValueRequestBody) *NullableUpdateEntityValueRequestBody {
	return &NullableUpdateEntityValueRequestBody{value: val, isSet: true}
}

func (v NullableUpdateEntityValueRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEntityValueRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


