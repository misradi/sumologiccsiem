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

// checks if the GetInsightResolutionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsightResolutionResponse{}

// GetInsightResolutionResponse struct for GetInsightResolutionResponse
type GetInsightResolutionResponse struct {
	Data InsightResolution `json:"data"`
}

type _GetInsightResolutionResponse GetInsightResolutionResponse

// NewGetInsightResolutionResponse instantiates a new GetInsightResolutionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsightResolutionResponse(data InsightResolution) *GetInsightResolutionResponse {
	this := GetInsightResolutionResponse{}
	this.Data = data
	return &this
}

// NewGetInsightResolutionResponseWithDefaults instantiates a new GetInsightResolutionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsightResolutionResponseWithDefaults() *GetInsightResolutionResponse {
	this := GetInsightResolutionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetInsightResolutionResponse) GetData() InsightResolution {
	if o == nil {
		var ret InsightResolution
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetInsightResolutionResponse) GetDataOk() (*InsightResolution, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetInsightResolutionResponse) SetData(v InsightResolution) {
	o.Data = v
}

func (o GetInsightResolutionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsightResolutionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetInsightResolutionResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetInsightResolutionResponse := _GetInsightResolutionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetInsightResolutionResponse)

	if err != nil {
		return err
	}

	*o = GetInsightResolutionResponse(varGetInsightResolutionResponse)

	return err
}

type NullableGetInsightResolutionResponse struct {
	value *GetInsightResolutionResponse
	isSet bool
}

func (v NullableGetInsightResolutionResponse) Get() *GetInsightResolutionResponse {
	return v.value
}

func (v *NullableGetInsightResolutionResponse) Set(val *GetInsightResolutionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsightResolutionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsightResolutionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsightResolutionResponse(val *GetInsightResolutionResponse) *NullableGetInsightResolutionResponse {
	return &NullableGetInsightResolutionResponse{value: val, isSet: true}
}

func (v NullableGetInsightResolutionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsightResolutionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


