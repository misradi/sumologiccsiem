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

// checks if the GetInsightExecutedAutomationsResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsightExecutedAutomationsResponseData{}

// GetInsightExecutedAutomationsResponseData struct for GetInsightExecutedAutomationsResponseData
type GetInsightExecutedAutomationsResponseData struct {
	ExecutedAutomations []ExecutedAutomation `json:"executedAutomations"`
}

type _GetInsightExecutedAutomationsResponseData GetInsightExecutedAutomationsResponseData

// NewGetInsightExecutedAutomationsResponseData instantiates a new GetInsightExecutedAutomationsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsightExecutedAutomationsResponseData(executedAutomations []ExecutedAutomation) *GetInsightExecutedAutomationsResponseData {
	this := GetInsightExecutedAutomationsResponseData{}
	this.ExecutedAutomations = executedAutomations
	return &this
}

// NewGetInsightExecutedAutomationsResponseDataWithDefaults instantiates a new GetInsightExecutedAutomationsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsightExecutedAutomationsResponseDataWithDefaults() *GetInsightExecutedAutomationsResponseData {
	this := GetInsightExecutedAutomationsResponseData{}
	return &this
}

// GetExecutedAutomations returns the ExecutedAutomations field value
func (o *GetInsightExecutedAutomationsResponseData) GetExecutedAutomations() []ExecutedAutomation {
	if o == nil {
		var ret []ExecutedAutomation
		return ret
	}

	return o.ExecutedAutomations
}

// GetExecutedAutomationsOk returns a tuple with the ExecutedAutomations field value
// and a boolean to check if the value has been set.
func (o *GetInsightExecutedAutomationsResponseData) GetExecutedAutomationsOk() ([]ExecutedAutomation, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutedAutomations, true
}

// SetExecutedAutomations sets field value
func (o *GetInsightExecutedAutomationsResponseData) SetExecutedAutomations(v []ExecutedAutomation) {
	o.ExecutedAutomations = v
}

func (o GetInsightExecutedAutomationsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsightExecutedAutomationsResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["executedAutomations"] = o.ExecutedAutomations
	return toSerialize, nil
}

func (o *GetInsightExecutedAutomationsResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"executedAutomations",
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

	varGetInsightExecutedAutomationsResponseData := _GetInsightExecutedAutomationsResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetInsightExecutedAutomationsResponseData)

	if err != nil {
		return err
	}

	*o = GetInsightExecutedAutomationsResponseData(varGetInsightExecutedAutomationsResponseData)

	return err
}

type NullableGetInsightExecutedAutomationsResponseData struct {
	value *GetInsightExecutedAutomationsResponseData
	isSet bool
}

func (v NullableGetInsightExecutedAutomationsResponseData) Get() *GetInsightExecutedAutomationsResponseData {
	return v.value
}

func (v *NullableGetInsightExecutedAutomationsResponseData) Set(val *GetInsightExecutedAutomationsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsightExecutedAutomationsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsightExecutedAutomationsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsightExecutedAutomationsResponseData(val *GetInsightExecutedAutomationsResponseData) *NullableGetInsightExecutedAutomationsResponseData {
	return &NullableGetInsightExecutedAutomationsResponseData{value: val, isSet: true}
}

func (v NullableGetInsightExecutedAutomationsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsightExecutedAutomationsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


