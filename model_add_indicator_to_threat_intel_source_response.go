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

// checks if the AddIndicatorToThreatIntelSourceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddIndicatorToThreatIntelSourceResponse{}

// AddIndicatorToThreatIntelSourceResponse struct for AddIndicatorToThreatIntelSourceResponse
type AddIndicatorToThreatIntelSourceResponse struct {
	Data DeleteContextActionResponseData `json:"data"`
}

type _AddIndicatorToThreatIntelSourceResponse AddIndicatorToThreatIntelSourceResponse

// NewAddIndicatorToThreatIntelSourceResponse instantiates a new AddIndicatorToThreatIntelSourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddIndicatorToThreatIntelSourceResponse(data DeleteContextActionResponseData) *AddIndicatorToThreatIntelSourceResponse {
	this := AddIndicatorToThreatIntelSourceResponse{}
	this.Data = data
	return &this
}

// NewAddIndicatorToThreatIntelSourceResponseWithDefaults instantiates a new AddIndicatorToThreatIntelSourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddIndicatorToThreatIntelSourceResponseWithDefaults() *AddIndicatorToThreatIntelSourceResponse {
	this := AddIndicatorToThreatIntelSourceResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AddIndicatorToThreatIntelSourceResponse) GetData() DeleteContextActionResponseData {
	if o == nil {
		var ret DeleteContextActionResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddIndicatorToThreatIntelSourceResponse) GetDataOk() (*DeleteContextActionResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AddIndicatorToThreatIntelSourceResponse) SetData(v DeleteContextActionResponseData) {
	o.Data = v
}

func (o AddIndicatorToThreatIntelSourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddIndicatorToThreatIntelSourceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AddIndicatorToThreatIntelSourceResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAddIndicatorToThreatIntelSourceResponse := _AddIndicatorToThreatIntelSourceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddIndicatorToThreatIntelSourceResponse)

	if err != nil {
		return err
	}

	*o = AddIndicatorToThreatIntelSourceResponse(varAddIndicatorToThreatIntelSourceResponse)

	return err
}

type NullableAddIndicatorToThreatIntelSourceResponse struct {
	value *AddIndicatorToThreatIntelSourceResponse
	isSet bool
}

func (v NullableAddIndicatorToThreatIntelSourceResponse) Get() *AddIndicatorToThreatIntelSourceResponse {
	return v.value
}

func (v *NullableAddIndicatorToThreatIntelSourceResponse) Set(val *AddIndicatorToThreatIntelSourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddIndicatorToThreatIntelSourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddIndicatorToThreatIntelSourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddIndicatorToThreatIntelSourceResponse(val *AddIndicatorToThreatIntelSourceResponse) *NullableAddIndicatorToThreatIntelSourceResponse {
	return &NullableAddIndicatorToThreatIntelSourceResponse{value: val, isSet: true}
}

func (v NullableAddIndicatorToThreatIntelSourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddIndicatorToThreatIntelSourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


