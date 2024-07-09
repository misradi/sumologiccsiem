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

// checks if the OverrideTemplatedMatchRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverrideTemplatedMatchRuleResponse{}

// OverrideTemplatedMatchRuleResponse struct for OverrideTemplatedMatchRuleResponse
type OverrideTemplatedMatchRuleResponse struct {
	Data TemplatedMatchRule `json:"data"`
}

type _OverrideTemplatedMatchRuleResponse OverrideTemplatedMatchRuleResponse

// NewOverrideTemplatedMatchRuleResponse instantiates a new OverrideTemplatedMatchRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideTemplatedMatchRuleResponse(data TemplatedMatchRule) *OverrideTemplatedMatchRuleResponse {
	this := OverrideTemplatedMatchRuleResponse{}
	this.Data = data
	return &this
}

// NewOverrideTemplatedMatchRuleResponseWithDefaults instantiates a new OverrideTemplatedMatchRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideTemplatedMatchRuleResponseWithDefaults() *OverrideTemplatedMatchRuleResponse {
	this := OverrideTemplatedMatchRuleResponse{}
	return &this
}

// GetData returns the Data field value
func (o *OverrideTemplatedMatchRuleResponse) GetData() TemplatedMatchRule {
	if o == nil {
		var ret TemplatedMatchRule
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OverrideTemplatedMatchRuleResponse) GetDataOk() (*TemplatedMatchRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OverrideTemplatedMatchRuleResponse) SetData(v TemplatedMatchRule) {
	o.Data = v
}

func (o OverrideTemplatedMatchRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideTemplatedMatchRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *OverrideTemplatedMatchRuleResponse) UnmarshalJSON(data []byte) (err error) {
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

	varOverrideTemplatedMatchRuleResponse := _OverrideTemplatedMatchRuleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOverrideTemplatedMatchRuleResponse)

	if err != nil {
		return err
	}

	*o = OverrideTemplatedMatchRuleResponse(varOverrideTemplatedMatchRuleResponse)

	return err
}

type NullableOverrideTemplatedMatchRuleResponse struct {
	value *OverrideTemplatedMatchRuleResponse
	isSet bool
}

func (v NullableOverrideTemplatedMatchRuleResponse) Get() *OverrideTemplatedMatchRuleResponse {
	return v.value
}

func (v *NullableOverrideTemplatedMatchRuleResponse) Set(val *OverrideTemplatedMatchRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideTemplatedMatchRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideTemplatedMatchRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideTemplatedMatchRuleResponse(val *OverrideTemplatedMatchRuleResponse) *NullableOverrideTemplatedMatchRuleResponse {
	return &NullableOverrideTemplatedMatchRuleResponse{value: val, isSet: true}
}

func (v NullableOverrideTemplatedMatchRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideTemplatedMatchRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


