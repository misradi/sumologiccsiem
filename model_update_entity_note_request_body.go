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

// checks if the UpdateEntityNoteRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEntityNoteRequestBody{}

// UpdateEntityNoteRequestBody struct for UpdateEntityNoteRequestBody
type UpdateEntityNoteRequestBody struct {
	Note string `json:"note"`
}

type _UpdateEntityNoteRequestBody UpdateEntityNoteRequestBody

// NewUpdateEntityNoteRequestBody instantiates a new UpdateEntityNoteRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEntityNoteRequestBody(note string) *UpdateEntityNoteRequestBody {
	this := UpdateEntityNoteRequestBody{}
	this.Note = note
	return &this
}

// NewUpdateEntityNoteRequestBodyWithDefaults instantiates a new UpdateEntityNoteRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEntityNoteRequestBodyWithDefaults() *UpdateEntityNoteRequestBody {
	this := UpdateEntityNoteRequestBody{}
	return &this
}

// GetNote returns the Note field value
func (o *UpdateEntityNoteRequestBody) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *UpdateEntityNoteRequestBody) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *UpdateEntityNoteRequestBody) SetNote(v string) {
	o.Note = v
}

func (o UpdateEntityNoteRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEntityNoteRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["note"] = o.Note
	return toSerialize, nil
}

func (o *UpdateEntityNoteRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"note",
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

	varUpdateEntityNoteRequestBody := _UpdateEntityNoteRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateEntityNoteRequestBody)

	if err != nil {
		return err
	}

	*o = UpdateEntityNoteRequestBody(varUpdateEntityNoteRequestBody)

	return err
}

type NullableUpdateEntityNoteRequestBody struct {
	value *UpdateEntityNoteRequestBody
	isSet bool
}

func (v NullableUpdateEntityNoteRequestBody) Get() *UpdateEntityNoteRequestBody {
	return v.value
}

func (v *NullableUpdateEntityNoteRequestBody) Set(val *UpdateEntityNoteRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEntityNoteRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEntityNoteRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEntityNoteRequestBody(val *UpdateEntityNoteRequestBody) *NullableUpdateEntityNoteRequestBody {
	return &NullableUpdateEntityNoteRequestBody{value: val, isSet: true}
}

func (v NullableUpdateEntityNoteRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEntityNoteRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


