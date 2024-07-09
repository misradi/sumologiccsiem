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

// checks if the CreateMatchRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMatchRuleResponse{}

// CreateMatchRuleResponse struct for CreateMatchRuleResponse
type CreateMatchRuleResponse struct {
	Data MatchRule `json:"data"`
}

type _CreateMatchRuleResponse CreateMatchRuleResponse

// NewCreateMatchRuleResponse instantiates a new CreateMatchRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMatchRuleResponse(data MatchRule) *CreateMatchRuleResponse {
	this := CreateMatchRuleResponse{}
	this.Data = data
	return &this
}

// NewCreateMatchRuleResponseWithDefaults instantiates a new CreateMatchRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMatchRuleResponseWithDefaults() *CreateMatchRuleResponse {
	this := CreateMatchRuleResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreateMatchRuleResponse) GetData() MatchRule {
	if o == nil {
		var ret MatchRule
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateMatchRuleResponse) GetDataOk() (*MatchRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateMatchRuleResponse) SetData(v MatchRule) {
	o.Data = v
}

func (o CreateMatchRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMatchRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CreateMatchRuleResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCreateMatchRuleResponse := _CreateMatchRuleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateMatchRuleResponse)

	if err != nil {
		return err
	}

	*o = CreateMatchRuleResponse(varCreateMatchRuleResponse)

	return err
}

type NullableCreateMatchRuleResponse struct {
	value *CreateMatchRuleResponse
	isSet bool
}

func (v NullableCreateMatchRuleResponse) Get() *CreateMatchRuleResponse {
	return v.value
}

func (v *NullableCreateMatchRuleResponse) Set(val *CreateMatchRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMatchRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMatchRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMatchRuleResponse(val *CreateMatchRuleResponse) *NullableCreateMatchRuleResponse {
	return &NullableCreateMatchRuleResponse{value: val, isSet: true}
}

func (v NullableCreateMatchRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMatchRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


