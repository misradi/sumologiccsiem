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

// checks if the GetThreatIntelIndicatorsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetThreatIntelIndicatorsResponse{}

// GetThreatIntelIndicatorsResponse struct for GetThreatIntelIndicatorsResponse
type GetThreatIntelIndicatorsResponse struct {
	Data GetThreatIntelIndicatorsResponseData `json:"data"`
}

type _GetThreatIntelIndicatorsResponse GetThreatIntelIndicatorsResponse

// NewGetThreatIntelIndicatorsResponse instantiates a new GetThreatIntelIndicatorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetThreatIntelIndicatorsResponse(data GetThreatIntelIndicatorsResponseData) *GetThreatIntelIndicatorsResponse {
	this := GetThreatIntelIndicatorsResponse{}
	this.Data = data
	return &this
}

// NewGetThreatIntelIndicatorsResponseWithDefaults instantiates a new GetThreatIntelIndicatorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetThreatIntelIndicatorsResponseWithDefaults() *GetThreatIntelIndicatorsResponse {
	this := GetThreatIntelIndicatorsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetThreatIntelIndicatorsResponse) GetData() GetThreatIntelIndicatorsResponseData {
	if o == nil {
		var ret GetThreatIntelIndicatorsResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetThreatIntelIndicatorsResponse) GetDataOk() (*GetThreatIntelIndicatorsResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetThreatIntelIndicatorsResponse) SetData(v GetThreatIntelIndicatorsResponseData) {
	o.Data = v
}

func (o GetThreatIntelIndicatorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetThreatIntelIndicatorsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetThreatIntelIndicatorsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetThreatIntelIndicatorsResponse := _GetThreatIntelIndicatorsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetThreatIntelIndicatorsResponse)

	if err != nil {
		return err
	}

	*o = GetThreatIntelIndicatorsResponse(varGetThreatIntelIndicatorsResponse)

	return err
}

type NullableGetThreatIntelIndicatorsResponse struct {
	value *GetThreatIntelIndicatorsResponse
	isSet bool
}

func (v NullableGetThreatIntelIndicatorsResponse) Get() *GetThreatIntelIndicatorsResponse {
	return v.value
}

func (v *NullableGetThreatIntelIndicatorsResponse) Set(val *GetThreatIntelIndicatorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetThreatIntelIndicatorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetThreatIntelIndicatorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetThreatIntelIndicatorsResponse(val *GetThreatIntelIndicatorsResponse) *NullableGetThreatIntelIndicatorsResponse {
	return &NullableGetThreatIntelIndicatorsResponse{value: val, isSet: true}
}

func (v NullableGetThreatIntelIndicatorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetThreatIntelIndicatorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


