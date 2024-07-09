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

// checks if the GetEntitiesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEntitiesResponse{}

// GetEntitiesResponse struct for GetEntitiesResponse
type GetEntitiesResponse struct {
	Data GetEntitiesResponseData `json:"data"`
}

type _GetEntitiesResponse GetEntitiesResponse

// NewGetEntitiesResponse instantiates a new GetEntitiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEntitiesResponse(data GetEntitiesResponseData) *GetEntitiesResponse {
	this := GetEntitiesResponse{}
	this.Data = data
	return &this
}

// NewGetEntitiesResponseWithDefaults instantiates a new GetEntitiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEntitiesResponseWithDefaults() *GetEntitiesResponse {
	this := GetEntitiesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetEntitiesResponse) GetData() GetEntitiesResponseData {
	if o == nil {
		var ret GetEntitiesResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponse) GetDataOk() (*GetEntitiesResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetEntitiesResponse) SetData(v GetEntitiesResponseData) {
	o.Data = v
}

func (o GetEntitiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEntitiesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetEntitiesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetEntitiesResponse := _GetEntitiesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEntitiesResponse)

	if err != nil {
		return err
	}

	*o = GetEntitiesResponse(varGetEntitiesResponse)

	return err
}

type NullableGetEntitiesResponse struct {
	value *GetEntitiesResponse
	isSet bool
}

func (v NullableGetEntitiesResponse) Get() *GetEntitiesResponse {
	return v.value
}

func (v *NullableGetEntitiesResponse) Set(val *GetEntitiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEntitiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEntitiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEntitiesResponse(val *GetEntitiesResponse) *NullableGetEntitiesResponse {
	return &NullableGetEntitiesResponse{value: val, isSet: true}
}

func (v NullableGetEntitiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEntitiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


