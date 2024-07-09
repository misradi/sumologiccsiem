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

// checks if the UpdateInsightConfigurationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInsightConfigurationResponse{}

// UpdateInsightConfigurationResponse struct for UpdateInsightConfigurationResponse
type UpdateInsightConfigurationResponse struct {
	Data GetInsightsConfigurationResponseData `json:"data"`
}

type _UpdateInsightConfigurationResponse UpdateInsightConfigurationResponse

// NewUpdateInsightConfigurationResponse instantiates a new UpdateInsightConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInsightConfigurationResponse(data GetInsightsConfigurationResponseData) *UpdateInsightConfigurationResponse {
	this := UpdateInsightConfigurationResponse{}
	this.Data = data
	return &this
}

// NewUpdateInsightConfigurationResponseWithDefaults instantiates a new UpdateInsightConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInsightConfigurationResponseWithDefaults() *UpdateInsightConfigurationResponse {
	this := UpdateInsightConfigurationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateInsightConfigurationResponse) GetData() GetInsightsConfigurationResponseData {
	if o == nil {
		var ret GetInsightsConfigurationResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateInsightConfigurationResponse) GetDataOk() (*GetInsightsConfigurationResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateInsightConfigurationResponse) SetData(v GetInsightsConfigurationResponseData) {
	o.Data = v
}

func (o UpdateInsightConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateInsightConfigurationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateInsightConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateInsightConfigurationResponse := _UpdateInsightConfigurationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateInsightConfigurationResponse)

	if err != nil {
		return err
	}

	*o = UpdateInsightConfigurationResponse(varUpdateInsightConfigurationResponse)

	return err
}

type NullableUpdateInsightConfigurationResponse struct {
	value *UpdateInsightConfigurationResponse
	isSet bool
}

func (v NullableUpdateInsightConfigurationResponse) Get() *UpdateInsightConfigurationResponse {
	return v.value
}

func (v *NullableUpdateInsightConfigurationResponse) Set(val *UpdateInsightConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInsightConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInsightConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInsightConfigurationResponse(val *UpdateInsightConfigurationResponse) *NullableUpdateInsightConfigurationResponse {
	return &NullableUpdateInsightConfigurationResponse{value: val, isSet: true}
}

func (v NullableUpdateInsightConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateInsightConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


