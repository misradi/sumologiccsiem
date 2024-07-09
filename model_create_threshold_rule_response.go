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

// checks if the CreateThresholdRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateThresholdRuleResponse{}

// CreateThresholdRuleResponse struct for CreateThresholdRuleResponse
type CreateThresholdRuleResponse struct {
	Data ThresholdRule `json:"data"`
}

type _CreateThresholdRuleResponse CreateThresholdRuleResponse

// NewCreateThresholdRuleResponse instantiates a new CreateThresholdRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateThresholdRuleResponse(data ThresholdRule) *CreateThresholdRuleResponse {
	this := CreateThresholdRuleResponse{}
	this.Data = data
	return &this
}

// NewCreateThresholdRuleResponseWithDefaults instantiates a new CreateThresholdRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateThresholdRuleResponseWithDefaults() *CreateThresholdRuleResponse {
	this := CreateThresholdRuleResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreateThresholdRuleResponse) GetData() ThresholdRule {
	if o == nil {
		var ret ThresholdRule
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleResponse) GetDataOk() (*ThresholdRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateThresholdRuleResponse) SetData(v ThresholdRule) {
	o.Data = v
}

func (o CreateThresholdRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateThresholdRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CreateThresholdRuleResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCreateThresholdRuleResponse := _CreateThresholdRuleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateThresholdRuleResponse)

	if err != nil {
		return err
	}

	*o = CreateThresholdRuleResponse(varCreateThresholdRuleResponse)

	return err
}

type NullableCreateThresholdRuleResponse struct {
	value *CreateThresholdRuleResponse
	isSet bool
}

func (v NullableCreateThresholdRuleResponse) Get() *CreateThresholdRuleResponse {
	return v.value
}

func (v *NullableCreateThresholdRuleResponse) Set(val *CreateThresholdRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateThresholdRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateThresholdRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateThresholdRuleResponse(val *CreateThresholdRuleResponse) *NullableCreateThresholdRuleResponse {
	return &NullableCreateThresholdRuleResponse{value: val, isSet: true}
}

func (v NullableCreateThresholdRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateThresholdRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


