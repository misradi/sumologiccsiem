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

// checks if the UpdateChainRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateChainRuleResponse{}

// UpdateChainRuleResponse struct for UpdateChainRuleResponse
type UpdateChainRuleResponse struct {
	Data ChainRule `json:"data"`
}

type _UpdateChainRuleResponse UpdateChainRuleResponse

// NewUpdateChainRuleResponse instantiates a new UpdateChainRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateChainRuleResponse(data ChainRule) *UpdateChainRuleResponse {
	this := UpdateChainRuleResponse{}
	this.Data = data
	return &this
}

// NewUpdateChainRuleResponseWithDefaults instantiates a new UpdateChainRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateChainRuleResponseWithDefaults() *UpdateChainRuleResponse {
	this := UpdateChainRuleResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateChainRuleResponse) GetData() ChainRule {
	if o == nil {
		var ret ChainRule
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateChainRuleResponse) GetDataOk() (*ChainRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateChainRuleResponse) SetData(v ChainRule) {
	o.Data = v
}

func (o UpdateChainRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateChainRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateChainRuleResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateChainRuleResponse := _UpdateChainRuleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateChainRuleResponse)

	if err != nil {
		return err
	}

	*o = UpdateChainRuleResponse(varUpdateChainRuleResponse)

	return err
}

type NullableUpdateChainRuleResponse struct {
	value *UpdateChainRuleResponse
	isSet bool
}

func (v NullableUpdateChainRuleResponse) Get() *UpdateChainRuleResponse {
	return v.value
}

func (v *NullableUpdateChainRuleResponse) Set(val *UpdateChainRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateChainRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateChainRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateChainRuleResponse(val *UpdateChainRuleResponse) *NullableUpdateChainRuleResponse {
	return &NullableUpdateChainRuleResponse{value: val, isSet: true}
}

func (v NullableUpdateChainRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateChainRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


