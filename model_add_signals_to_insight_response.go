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

// checks if the AddSignalsToInsightResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSignalsToInsightResponse{}

// AddSignalsToInsightResponse struct for AddSignalsToInsightResponse
type AddSignalsToInsightResponse struct {
	Data Insight `json:"data"`
}

type _AddSignalsToInsightResponse AddSignalsToInsightResponse

// NewAddSignalsToInsightResponse instantiates a new AddSignalsToInsightResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSignalsToInsightResponse(data Insight) *AddSignalsToInsightResponse {
	this := AddSignalsToInsightResponse{}
	this.Data = data
	return &this
}

// NewAddSignalsToInsightResponseWithDefaults instantiates a new AddSignalsToInsightResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSignalsToInsightResponseWithDefaults() *AddSignalsToInsightResponse {
	this := AddSignalsToInsightResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AddSignalsToInsightResponse) GetData() Insight {
	if o == nil {
		var ret Insight
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddSignalsToInsightResponse) GetDataOk() (*Insight, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AddSignalsToInsightResponse) SetData(v Insight) {
	o.Data = v
}

func (o AddSignalsToInsightResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSignalsToInsightResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AddSignalsToInsightResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAddSignalsToInsightResponse := _AddSignalsToInsightResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddSignalsToInsightResponse)

	if err != nil {
		return err
	}

	*o = AddSignalsToInsightResponse(varAddSignalsToInsightResponse)

	return err
}

type NullableAddSignalsToInsightResponse struct {
	value *AddSignalsToInsightResponse
	isSet bool
}

func (v NullableAddSignalsToInsightResponse) Get() *AddSignalsToInsightResponse {
	return v.value
}

func (v *NullableAddSignalsToInsightResponse) Set(val *AddSignalsToInsightResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSignalsToInsightResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSignalsToInsightResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSignalsToInsightResponse(val *AddSignalsToInsightResponse) *NullableAddSignalsToInsightResponse {
	return &NullableAddSignalsToInsightResponse{value: val, isSet: true}
}

func (v NullableAddSignalsToInsightResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSignalsToInsightResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


