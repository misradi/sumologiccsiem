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

// checks if the UpdateAutomationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAutomationResponse{}

// UpdateAutomationResponse struct for UpdateAutomationResponse
type UpdateAutomationResponse struct {
	Data Automation `json:"data"`
}

type _UpdateAutomationResponse UpdateAutomationResponse

// NewUpdateAutomationResponse instantiates a new UpdateAutomationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAutomationResponse(data Automation) *UpdateAutomationResponse {
	this := UpdateAutomationResponse{}
	this.Data = data
	return &this
}

// NewUpdateAutomationResponseWithDefaults instantiates a new UpdateAutomationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAutomationResponseWithDefaults() *UpdateAutomationResponse {
	this := UpdateAutomationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateAutomationResponse) GetData() Automation {
	if o == nil {
		var ret Automation
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateAutomationResponse) GetDataOk() (*Automation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateAutomationResponse) SetData(v Automation) {
	o.Data = v
}

func (o UpdateAutomationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAutomationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateAutomationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateAutomationResponse := _UpdateAutomationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateAutomationResponse)

	if err != nil {
		return err
	}

	*o = UpdateAutomationResponse(varUpdateAutomationResponse)

	return err
}

type NullableUpdateAutomationResponse struct {
	value *UpdateAutomationResponse
	isSet bool
}

func (v NullableUpdateAutomationResponse) Get() *UpdateAutomationResponse {
	return v.value
}

func (v *NullableUpdateAutomationResponse) Set(val *UpdateAutomationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAutomationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAutomationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAutomationResponse(val *UpdateAutomationResponse) *NullableUpdateAutomationResponse {
	return &NullableUpdateAutomationResponse{value: val, isSet: true}
}

func (v NullableUpdateAutomationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAutomationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


