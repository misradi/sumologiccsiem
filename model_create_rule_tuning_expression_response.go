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

// checks if the CreateRuleTuningExpressionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRuleTuningExpressionResponse{}

// CreateRuleTuningExpressionResponse struct for CreateRuleTuningExpressionResponse
type CreateRuleTuningExpressionResponse struct {
	Data RuleTuningExpression `json:"data"`
}

type _CreateRuleTuningExpressionResponse CreateRuleTuningExpressionResponse

// NewCreateRuleTuningExpressionResponse instantiates a new CreateRuleTuningExpressionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRuleTuningExpressionResponse(data RuleTuningExpression) *CreateRuleTuningExpressionResponse {
	this := CreateRuleTuningExpressionResponse{}
	this.Data = data
	return &this
}

// NewCreateRuleTuningExpressionResponseWithDefaults instantiates a new CreateRuleTuningExpressionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRuleTuningExpressionResponseWithDefaults() *CreateRuleTuningExpressionResponse {
	this := CreateRuleTuningExpressionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreateRuleTuningExpressionResponse) GetData() RuleTuningExpression {
	if o == nil {
		var ret RuleTuningExpression
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateRuleTuningExpressionResponse) GetDataOk() (*RuleTuningExpression, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateRuleTuningExpressionResponse) SetData(v RuleTuningExpression) {
	o.Data = v
}

func (o CreateRuleTuningExpressionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRuleTuningExpressionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CreateRuleTuningExpressionResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCreateRuleTuningExpressionResponse := _CreateRuleTuningExpressionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRuleTuningExpressionResponse)

	if err != nil {
		return err
	}

	*o = CreateRuleTuningExpressionResponse(varCreateRuleTuningExpressionResponse)

	return err
}

type NullableCreateRuleTuningExpressionResponse struct {
	value *CreateRuleTuningExpressionResponse
	isSet bool
}

func (v NullableCreateRuleTuningExpressionResponse) Get() *CreateRuleTuningExpressionResponse {
	return v.value
}

func (v *NullableCreateRuleTuningExpressionResponse) Set(val *CreateRuleTuningExpressionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRuleTuningExpressionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRuleTuningExpressionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRuleTuningExpressionResponse(val *CreateRuleTuningExpressionResponse) *NullableCreateRuleTuningExpressionResponse {
	return &NullableCreateRuleTuningExpressionResponse{value: val, isSet: true}
}

func (v NullableCreateRuleTuningExpressionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRuleTuningExpressionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

