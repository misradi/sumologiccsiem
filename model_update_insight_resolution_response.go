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

// checks if the UpdateInsightResolutionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInsightResolutionResponse{}

// UpdateInsightResolutionResponse struct for UpdateInsightResolutionResponse
type UpdateInsightResolutionResponse struct {
	Data InsightResolution `json:"data"`
}

type _UpdateInsightResolutionResponse UpdateInsightResolutionResponse

// NewUpdateInsightResolutionResponse instantiates a new UpdateInsightResolutionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInsightResolutionResponse(data InsightResolution) *UpdateInsightResolutionResponse {
	this := UpdateInsightResolutionResponse{}
	this.Data = data
	return &this
}

// NewUpdateInsightResolutionResponseWithDefaults instantiates a new UpdateInsightResolutionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInsightResolutionResponseWithDefaults() *UpdateInsightResolutionResponse {
	this := UpdateInsightResolutionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateInsightResolutionResponse) GetData() InsightResolution {
	if o == nil {
		var ret InsightResolution
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateInsightResolutionResponse) GetDataOk() (*InsightResolution, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateInsightResolutionResponse) SetData(v InsightResolution) {
	o.Data = v
}

func (o UpdateInsightResolutionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateInsightResolutionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateInsightResolutionResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateInsightResolutionResponse := _UpdateInsightResolutionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateInsightResolutionResponse)

	if err != nil {
		return err
	}

	*o = UpdateInsightResolutionResponse(varUpdateInsightResolutionResponse)

	return err
}

type NullableUpdateInsightResolutionResponse struct {
	value *UpdateInsightResolutionResponse
	isSet bool
}

func (v NullableUpdateInsightResolutionResponse) Get() *UpdateInsightResolutionResponse {
	return v.value
}

func (v *NullableUpdateInsightResolutionResponse) Set(val *UpdateInsightResolutionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInsightResolutionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInsightResolutionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInsightResolutionResponse(val *UpdateInsightResolutionResponse) *NullableUpdateInsightResolutionResponse {
	return &NullableUpdateInsightResolutionResponse{value: val, isSet: true}
}

func (v NullableUpdateInsightResolutionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateInsightResolutionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


