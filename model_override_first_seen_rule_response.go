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

// checks if the OverrideFirstSeenRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverrideFirstSeenRuleResponse{}

// OverrideFirstSeenRuleResponse struct for OverrideFirstSeenRuleResponse
type OverrideFirstSeenRuleResponse struct {
	Data FirstSeenRule `json:"data"`
}

type _OverrideFirstSeenRuleResponse OverrideFirstSeenRuleResponse

// NewOverrideFirstSeenRuleResponse instantiates a new OverrideFirstSeenRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideFirstSeenRuleResponse(data FirstSeenRule) *OverrideFirstSeenRuleResponse {
	this := OverrideFirstSeenRuleResponse{}
	this.Data = data
	return &this
}

// NewOverrideFirstSeenRuleResponseWithDefaults instantiates a new OverrideFirstSeenRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideFirstSeenRuleResponseWithDefaults() *OverrideFirstSeenRuleResponse {
	this := OverrideFirstSeenRuleResponse{}
	return &this
}

// GetData returns the Data field value
func (o *OverrideFirstSeenRuleResponse) GetData() FirstSeenRule {
	if o == nil {
		var ret FirstSeenRule
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OverrideFirstSeenRuleResponse) GetDataOk() (*FirstSeenRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OverrideFirstSeenRuleResponse) SetData(v FirstSeenRule) {
	o.Data = v
}

func (o OverrideFirstSeenRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideFirstSeenRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *OverrideFirstSeenRuleResponse) UnmarshalJSON(data []byte) (err error) {
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

	varOverrideFirstSeenRuleResponse := _OverrideFirstSeenRuleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOverrideFirstSeenRuleResponse)

	if err != nil {
		return err
	}

	*o = OverrideFirstSeenRuleResponse(varOverrideFirstSeenRuleResponse)

	return err
}

type NullableOverrideFirstSeenRuleResponse struct {
	value *OverrideFirstSeenRuleResponse
	isSet bool
}

func (v NullableOverrideFirstSeenRuleResponse) Get() *OverrideFirstSeenRuleResponse {
	return v.value
}

func (v *NullableOverrideFirstSeenRuleResponse) Set(val *OverrideFirstSeenRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideFirstSeenRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideFirstSeenRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideFirstSeenRuleResponse(val *OverrideFirstSeenRuleResponse) *NullableOverrideFirstSeenRuleResponse {
	return &NullableOverrideFirstSeenRuleResponse{value: val, isSet: true}
}

func (v NullableOverrideFirstSeenRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideFirstSeenRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


