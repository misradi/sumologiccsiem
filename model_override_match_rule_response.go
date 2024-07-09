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

// checks if the OverrideMatchRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverrideMatchRuleResponse{}

// OverrideMatchRuleResponse struct for OverrideMatchRuleResponse
type OverrideMatchRuleResponse struct {
	Data MatchRule `json:"data"`
}

type _OverrideMatchRuleResponse OverrideMatchRuleResponse

// NewOverrideMatchRuleResponse instantiates a new OverrideMatchRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideMatchRuleResponse(data MatchRule) *OverrideMatchRuleResponse {
	this := OverrideMatchRuleResponse{}
	this.Data = data
	return &this
}

// NewOverrideMatchRuleResponseWithDefaults instantiates a new OverrideMatchRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideMatchRuleResponseWithDefaults() *OverrideMatchRuleResponse {
	this := OverrideMatchRuleResponse{}
	return &this
}

// GetData returns the Data field value
func (o *OverrideMatchRuleResponse) GetData() MatchRule {
	if o == nil {
		var ret MatchRule
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OverrideMatchRuleResponse) GetDataOk() (*MatchRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OverrideMatchRuleResponse) SetData(v MatchRule) {
	o.Data = v
}

func (o OverrideMatchRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideMatchRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *OverrideMatchRuleResponse) UnmarshalJSON(data []byte) (err error) {
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

	varOverrideMatchRuleResponse := _OverrideMatchRuleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOverrideMatchRuleResponse)

	if err != nil {
		return err
	}

	*o = OverrideMatchRuleResponse(varOverrideMatchRuleResponse)

	return err
}

type NullableOverrideMatchRuleResponse struct {
	value *OverrideMatchRuleResponse
	isSet bool
}

func (v NullableOverrideMatchRuleResponse) Get() *OverrideMatchRuleResponse {
	return v.value
}

func (v *NullableOverrideMatchRuleResponse) Set(val *OverrideMatchRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideMatchRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideMatchRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideMatchRuleResponse(val *OverrideMatchRuleResponse) *NullableOverrideMatchRuleResponse {
	return &NullableOverrideMatchRuleResponse{value: val, isSet: true}
}

func (v NullableOverrideMatchRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideMatchRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


