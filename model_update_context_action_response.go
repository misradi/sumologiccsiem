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

// checks if the UpdateContextActionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateContextActionResponse{}

// UpdateContextActionResponse struct for UpdateContextActionResponse
type UpdateContextActionResponse struct {
	Data ContextAction `json:"data"`
}

type _UpdateContextActionResponse UpdateContextActionResponse

// NewUpdateContextActionResponse instantiates a new UpdateContextActionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateContextActionResponse(data ContextAction) *UpdateContextActionResponse {
	this := UpdateContextActionResponse{}
	this.Data = data
	return &this
}

// NewUpdateContextActionResponseWithDefaults instantiates a new UpdateContextActionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateContextActionResponseWithDefaults() *UpdateContextActionResponse {
	this := UpdateContextActionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateContextActionResponse) GetData() ContextAction {
	if o == nil {
		var ret ContextAction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateContextActionResponse) GetDataOk() (*ContextAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateContextActionResponse) SetData(v ContextAction) {
	o.Data = v
}

func (o UpdateContextActionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateContextActionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateContextActionResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateContextActionResponse := _UpdateContextActionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateContextActionResponse)

	if err != nil {
		return err
	}

	*o = UpdateContextActionResponse(varUpdateContextActionResponse)

	return err
}

type NullableUpdateContextActionResponse struct {
	value *UpdateContextActionResponse
	isSet bool
}

func (v NullableUpdateContextActionResponse) Get() *UpdateContextActionResponse {
	return v.value
}

func (v *NullableUpdateContextActionResponse) Set(val *UpdateContextActionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateContextActionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateContextActionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateContextActionResponse(val *UpdateContextActionResponse) *NullableUpdateContextActionResponse {
	return &NullableUpdateContextActionResponse{value: val, isSet: true}
}

func (v NullableUpdateContextActionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateContextActionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


