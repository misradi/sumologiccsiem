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

// checks if the AddIndicatorToThreatIntelSourceRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddIndicatorToThreatIntelSourceRequestBody{}

// AddIndicatorToThreatIntelSourceRequestBody struct for AddIndicatorToThreatIntelSourceRequestBody
type AddIndicatorToThreatIntelSourceRequestBody struct {
	Indicators []AddItemsToSuppressListRequestBodyItemsInner `json:"indicators"`
}

type _AddIndicatorToThreatIntelSourceRequestBody AddIndicatorToThreatIntelSourceRequestBody

// NewAddIndicatorToThreatIntelSourceRequestBody instantiates a new AddIndicatorToThreatIntelSourceRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddIndicatorToThreatIntelSourceRequestBody(indicators []AddItemsToSuppressListRequestBodyItemsInner) *AddIndicatorToThreatIntelSourceRequestBody {
	this := AddIndicatorToThreatIntelSourceRequestBody{}
	this.Indicators = indicators
	return &this
}

// NewAddIndicatorToThreatIntelSourceRequestBodyWithDefaults instantiates a new AddIndicatorToThreatIntelSourceRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddIndicatorToThreatIntelSourceRequestBodyWithDefaults() *AddIndicatorToThreatIntelSourceRequestBody {
	this := AddIndicatorToThreatIntelSourceRequestBody{}
	return &this
}

// GetIndicators returns the Indicators field value
func (o *AddIndicatorToThreatIntelSourceRequestBody) GetIndicators() []AddItemsToSuppressListRequestBodyItemsInner {
	if o == nil {
		var ret []AddItemsToSuppressListRequestBodyItemsInner
		return ret
	}

	return o.Indicators
}

// GetIndicatorsOk returns a tuple with the Indicators field value
// and a boolean to check if the value has been set.
func (o *AddIndicatorToThreatIntelSourceRequestBody) GetIndicatorsOk() ([]AddItemsToSuppressListRequestBodyItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Indicators, true
}

// SetIndicators sets field value
func (o *AddIndicatorToThreatIntelSourceRequestBody) SetIndicators(v []AddItemsToSuppressListRequestBodyItemsInner) {
	o.Indicators = v
}

func (o AddIndicatorToThreatIntelSourceRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddIndicatorToThreatIntelSourceRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["indicators"] = o.Indicators
	return toSerialize, nil
}

func (o *AddIndicatorToThreatIntelSourceRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"indicators",
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

	varAddIndicatorToThreatIntelSourceRequestBody := _AddIndicatorToThreatIntelSourceRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddIndicatorToThreatIntelSourceRequestBody)

	if err != nil {
		return err
	}

	*o = AddIndicatorToThreatIntelSourceRequestBody(varAddIndicatorToThreatIntelSourceRequestBody)

	return err
}

type NullableAddIndicatorToThreatIntelSourceRequestBody struct {
	value *AddIndicatorToThreatIntelSourceRequestBody
	isSet bool
}

func (v NullableAddIndicatorToThreatIntelSourceRequestBody) Get() *AddIndicatorToThreatIntelSourceRequestBody {
	return v.value
}

func (v *NullableAddIndicatorToThreatIntelSourceRequestBody) Set(val *AddIndicatorToThreatIntelSourceRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAddIndicatorToThreatIntelSourceRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAddIndicatorToThreatIntelSourceRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddIndicatorToThreatIntelSourceRequestBody(val *AddIndicatorToThreatIntelSourceRequestBody) *NullableAddIndicatorToThreatIntelSourceRequestBody {
	return &NullableAddIndicatorToThreatIntelSourceRequestBody{value: val, isSet: true}
}

func (v NullableAddIndicatorToThreatIntelSourceRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddIndicatorToThreatIntelSourceRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


