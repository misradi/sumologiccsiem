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

// checks if the UpdateEntitySuppressedRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEntitySuppressedRequestBody{}

// UpdateEntitySuppressedRequestBody struct for UpdateEntitySuppressedRequestBody
type UpdateEntitySuppressedRequestBody struct {
	Suppressed bool `json:"suppressed"`
}

type _UpdateEntitySuppressedRequestBody UpdateEntitySuppressedRequestBody

// NewUpdateEntitySuppressedRequestBody instantiates a new UpdateEntitySuppressedRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEntitySuppressedRequestBody(suppressed bool) *UpdateEntitySuppressedRequestBody {
	this := UpdateEntitySuppressedRequestBody{}
	this.Suppressed = suppressed
	return &this
}

// NewUpdateEntitySuppressedRequestBodyWithDefaults instantiates a new UpdateEntitySuppressedRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEntitySuppressedRequestBodyWithDefaults() *UpdateEntitySuppressedRequestBody {
	this := UpdateEntitySuppressedRequestBody{}
	return &this
}

// GetSuppressed returns the Suppressed field value
func (o *UpdateEntitySuppressedRequestBody) GetSuppressed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value
// and a boolean to check if the value has been set.
func (o *UpdateEntitySuppressedRequestBody) GetSuppressedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suppressed, true
}

// SetSuppressed sets field value
func (o *UpdateEntitySuppressedRequestBody) SetSuppressed(v bool) {
	o.Suppressed = v
}

func (o UpdateEntitySuppressedRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEntitySuppressedRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["suppressed"] = o.Suppressed
	return toSerialize, nil
}

func (o *UpdateEntitySuppressedRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"suppressed",
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

	varUpdateEntitySuppressedRequestBody := _UpdateEntitySuppressedRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateEntitySuppressedRequestBody)

	if err != nil {
		return err
	}

	*o = UpdateEntitySuppressedRequestBody(varUpdateEntitySuppressedRequestBody)

	return err
}

type NullableUpdateEntitySuppressedRequestBody struct {
	value *UpdateEntitySuppressedRequestBody
	isSet bool
}

func (v NullableUpdateEntitySuppressedRequestBody) Get() *UpdateEntitySuppressedRequestBody {
	return v.value
}

func (v *NullableUpdateEntitySuppressedRequestBody) Set(val *UpdateEntitySuppressedRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEntitySuppressedRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEntitySuppressedRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEntitySuppressedRequestBody(val *UpdateEntitySuppressedRequestBody) *NullableUpdateEntitySuppressedRequestBody {
	return &NullableUpdateEntitySuppressedRequestBody{value: val, isSet: true}
}

func (v NullableUpdateEntitySuppressedRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEntitySuppressedRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


