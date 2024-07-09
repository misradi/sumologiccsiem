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

// checks if the CreateInsightResolutionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInsightResolutionResponse{}

// CreateInsightResolutionResponse struct for CreateInsightResolutionResponse
type CreateInsightResolutionResponse struct {
	Data InsightResolution `json:"data"`
}

type _CreateInsightResolutionResponse CreateInsightResolutionResponse

// NewCreateInsightResolutionResponse instantiates a new CreateInsightResolutionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInsightResolutionResponse(data InsightResolution) *CreateInsightResolutionResponse {
	this := CreateInsightResolutionResponse{}
	this.Data = data
	return &this
}

// NewCreateInsightResolutionResponseWithDefaults instantiates a new CreateInsightResolutionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInsightResolutionResponseWithDefaults() *CreateInsightResolutionResponse {
	this := CreateInsightResolutionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreateInsightResolutionResponse) GetData() InsightResolution {
	if o == nil {
		var ret InsightResolution
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateInsightResolutionResponse) GetDataOk() (*InsightResolution, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateInsightResolutionResponse) SetData(v InsightResolution) {
	o.Data = v
}

func (o CreateInsightResolutionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInsightResolutionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CreateInsightResolutionResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCreateInsightResolutionResponse := _CreateInsightResolutionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateInsightResolutionResponse)

	if err != nil {
		return err
	}

	*o = CreateInsightResolutionResponse(varCreateInsightResolutionResponse)

	return err
}

type NullableCreateInsightResolutionResponse struct {
	value *CreateInsightResolutionResponse
	isSet bool
}

func (v NullableCreateInsightResolutionResponse) Get() *CreateInsightResolutionResponse {
	return v.value
}

func (v *NullableCreateInsightResolutionResponse) Set(val *CreateInsightResolutionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInsightResolutionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInsightResolutionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInsightResolutionResponse(val *CreateInsightResolutionResponse) *NullableCreateInsightResolutionResponse {
	return &NullableCreateInsightResolutionResponse{value: val, isSet: true}
}

func (v NullableCreateInsightResolutionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInsightResolutionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


