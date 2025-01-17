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

// checks if the ExecuteAutomationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecuteAutomationResponse{}

// ExecuteAutomationResponse struct for ExecuteAutomationResponse
type ExecuteAutomationResponse struct {
	Data ExecuteAutomationResponseData `json:"data"`
}

type _ExecuteAutomationResponse ExecuteAutomationResponse

// NewExecuteAutomationResponse instantiates a new ExecuteAutomationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteAutomationResponse(data ExecuteAutomationResponseData) *ExecuteAutomationResponse {
	this := ExecuteAutomationResponse{}
	this.Data = data
	return &this
}

// NewExecuteAutomationResponseWithDefaults instantiates a new ExecuteAutomationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteAutomationResponseWithDefaults() *ExecuteAutomationResponse {
	this := ExecuteAutomationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ExecuteAutomationResponse) GetData() ExecuteAutomationResponseData {
	if o == nil {
		var ret ExecuteAutomationResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExecuteAutomationResponse) GetDataOk() (*ExecuteAutomationResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ExecuteAutomationResponse) SetData(v ExecuteAutomationResponseData) {
	o.Data = v
}

func (o ExecuteAutomationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecuteAutomationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ExecuteAutomationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varExecuteAutomationResponse := _ExecuteAutomationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExecuteAutomationResponse)

	if err != nil {
		return err
	}

	*o = ExecuteAutomationResponse(varExecuteAutomationResponse)

	return err
}

type NullableExecuteAutomationResponse struct {
	value *ExecuteAutomationResponse
	isSet bool
}

func (v NullableExecuteAutomationResponse) Get() *ExecuteAutomationResponse {
	return v.value
}

func (v *NullableExecuteAutomationResponse) Set(val *ExecuteAutomationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteAutomationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteAutomationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteAutomationResponse(val *ExecuteAutomationResponse) *NullableExecuteAutomationResponse {
	return &NullableExecuteAutomationResponse{value: val, isSet: true}
}

func (v NullableExecuteAutomationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteAutomationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


