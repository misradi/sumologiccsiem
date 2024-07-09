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

// checks if the UpdateFirstSeenRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFirstSeenRuleResponse{}

// UpdateFirstSeenRuleResponse struct for UpdateFirstSeenRuleResponse
type UpdateFirstSeenRuleResponse struct {
	Data FirstSeenRule `json:"data"`
}

type _UpdateFirstSeenRuleResponse UpdateFirstSeenRuleResponse

// NewUpdateFirstSeenRuleResponse instantiates a new UpdateFirstSeenRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFirstSeenRuleResponse(data FirstSeenRule) *UpdateFirstSeenRuleResponse {
	this := UpdateFirstSeenRuleResponse{}
	this.Data = data
	return &this
}

// NewUpdateFirstSeenRuleResponseWithDefaults instantiates a new UpdateFirstSeenRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFirstSeenRuleResponseWithDefaults() *UpdateFirstSeenRuleResponse {
	this := UpdateFirstSeenRuleResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateFirstSeenRuleResponse) GetData() FirstSeenRule {
	if o == nil {
		var ret FirstSeenRule
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleResponse) GetDataOk() (*FirstSeenRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateFirstSeenRuleResponse) SetData(v FirstSeenRule) {
	o.Data = v
}

func (o UpdateFirstSeenRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFirstSeenRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateFirstSeenRuleResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateFirstSeenRuleResponse := _UpdateFirstSeenRuleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateFirstSeenRuleResponse)

	if err != nil {
		return err
	}

	*o = UpdateFirstSeenRuleResponse(varUpdateFirstSeenRuleResponse)

	return err
}

type NullableUpdateFirstSeenRuleResponse struct {
	value *UpdateFirstSeenRuleResponse
	isSet bool
}

func (v NullableUpdateFirstSeenRuleResponse) Get() *UpdateFirstSeenRuleResponse {
	return v.value
}

func (v *NullableUpdateFirstSeenRuleResponse) Set(val *UpdateFirstSeenRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFirstSeenRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFirstSeenRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFirstSeenRuleResponse(val *UpdateFirstSeenRuleResponse) *NullableUpdateFirstSeenRuleResponse {
	return &NullableUpdateFirstSeenRuleResponse{value: val, isSet: true}
}

func (v NullableUpdateFirstSeenRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFirstSeenRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

