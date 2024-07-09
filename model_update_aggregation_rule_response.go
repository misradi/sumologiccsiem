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

// checks if the UpdateAggregationRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAggregationRuleResponse{}

// UpdateAggregationRuleResponse struct for UpdateAggregationRuleResponse
type UpdateAggregationRuleResponse struct {
	Data AggregationRule `json:"data"`
}

type _UpdateAggregationRuleResponse UpdateAggregationRuleResponse

// NewUpdateAggregationRuleResponse instantiates a new UpdateAggregationRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAggregationRuleResponse(data AggregationRule) *UpdateAggregationRuleResponse {
	this := UpdateAggregationRuleResponse{}
	this.Data = data
	return &this
}

// NewUpdateAggregationRuleResponseWithDefaults instantiates a new UpdateAggregationRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAggregationRuleResponseWithDefaults() *UpdateAggregationRuleResponse {
	this := UpdateAggregationRuleResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateAggregationRuleResponse) GetData() AggregationRule {
	if o == nil {
		var ret AggregationRule
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateAggregationRuleResponse) GetDataOk() (*AggregationRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateAggregationRuleResponse) SetData(v AggregationRule) {
	o.Data = v
}

func (o UpdateAggregationRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAggregationRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateAggregationRuleResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateAggregationRuleResponse := _UpdateAggregationRuleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateAggregationRuleResponse)

	if err != nil {
		return err
	}

	*o = UpdateAggregationRuleResponse(varUpdateAggregationRuleResponse)

	return err
}

type NullableUpdateAggregationRuleResponse struct {
	value *UpdateAggregationRuleResponse
	isSet bool
}

func (v NullableUpdateAggregationRuleResponse) Get() *UpdateAggregationRuleResponse {
	return v.value
}

func (v *NullableUpdateAggregationRuleResponse) Set(val *UpdateAggregationRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAggregationRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAggregationRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAggregationRuleResponse(val *UpdateAggregationRuleResponse) *NullableUpdateAggregationRuleResponse {
	return &NullableUpdateAggregationRuleResponse{value: val, isSet: true}
}

func (v NullableUpdateAggregationRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAggregationRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


