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

// checks if the UpdateLogMappingEnabledResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateLogMappingEnabledResponse{}

// UpdateLogMappingEnabledResponse struct for UpdateLogMappingEnabledResponse
type UpdateLogMappingEnabledResponse struct {
	Data LogMapping `json:"data"`
}

type _UpdateLogMappingEnabledResponse UpdateLogMappingEnabledResponse

// NewUpdateLogMappingEnabledResponse instantiates a new UpdateLogMappingEnabledResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLogMappingEnabledResponse(data LogMapping) *UpdateLogMappingEnabledResponse {
	this := UpdateLogMappingEnabledResponse{}
	this.Data = data
	return &this
}

// NewUpdateLogMappingEnabledResponseWithDefaults instantiates a new UpdateLogMappingEnabledResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLogMappingEnabledResponseWithDefaults() *UpdateLogMappingEnabledResponse {
	this := UpdateLogMappingEnabledResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateLogMappingEnabledResponse) GetData() LogMapping {
	if o == nil {
		var ret LogMapping
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateLogMappingEnabledResponse) GetDataOk() (*LogMapping, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateLogMappingEnabledResponse) SetData(v LogMapping) {
	o.Data = v
}

func (o UpdateLogMappingEnabledResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateLogMappingEnabledResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateLogMappingEnabledResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varUpdateLogMappingEnabledResponse := _UpdateLogMappingEnabledResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateLogMappingEnabledResponse)

	if err != nil {
		return err
	}

	*o = UpdateLogMappingEnabledResponse(varUpdateLogMappingEnabledResponse)

	return err
}

type NullableUpdateLogMappingEnabledResponse struct {
	value *UpdateLogMappingEnabledResponse
	isSet bool
}

func (v NullableUpdateLogMappingEnabledResponse) Get() *UpdateLogMappingEnabledResponse {
	return v.value
}

func (v *NullableUpdateLogMappingEnabledResponse) Set(val *UpdateLogMappingEnabledResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLogMappingEnabledResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLogMappingEnabledResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLogMappingEnabledResponse(val *UpdateLogMappingEnabledResponse) *NullableUpdateLogMappingEnabledResponse {
	return &NullableUpdateLogMappingEnabledResponse{value: val, isSet: true}
}

func (v NullableUpdateLogMappingEnabledResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLogMappingEnabledResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

