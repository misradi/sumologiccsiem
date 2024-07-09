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

// checks if the CreateLogMappingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLogMappingResponse{}

// CreateLogMappingResponse struct for CreateLogMappingResponse
type CreateLogMappingResponse struct {
	Data LogMapping `json:"data"`
}

type _CreateLogMappingResponse CreateLogMappingResponse

// NewCreateLogMappingResponse instantiates a new CreateLogMappingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLogMappingResponse(data LogMapping) *CreateLogMappingResponse {
	this := CreateLogMappingResponse{}
	this.Data = data
	return &this
}

// NewCreateLogMappingResponseWithDefaults instantiates a new CreateLogMappingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLogMappingResponseWithDefaults() *CreateLogMappingResponse {
	this := CreateLogMappingResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreateLogMappingResponse) GetData() LogMapping {
	if o == nil {
		var ret LogMapping
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateLogMappingResponse) GetDataOk() (*LogMapping, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateLogMappingResponse) SetData(v LogMapping) {
	o.Data = v
}

func (o CreateLogMappingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLogMappingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CreateLogMappingResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCreateLogMappingResponse := _CreateLogMappingResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateLogMappingResponse)

	if err != nil {
		return err
	}

	*o = CreateLogMappingResponse(varCreateLogMappingResponse)

	return err
}

type NullableCreateLogMappingResponse struct {
	value *CreateLogMappingResponse
	isSet bool
}

func (v NullableCreateLogMappingResponse) Get() *CreateLogMappingResponse {
	return v.value
}

func (v *NullableCreateLogMappingResponse) Set(val *CreateLogMappingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLogMappingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLogMappingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLogMappingResponse(val *CreateLogMappingResponse) *NullableCreateLogMappingResponse {
	return &NullableCreateLogMappingResponse{value: val, isSet: true}
}

func (v NullableCreateLogMappingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLogMappingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

