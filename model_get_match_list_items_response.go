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

// checks if the GetMatchListItemsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMatchListItemsResponse{}

// GetMatchListItemsResponse struct for GetMatchListItemsResponse
type GetMatchListItemsResponse struct {
	Data GetMatchListItemsResponseData `json:"data"`
}

type _GetMatchListItemsResponse GetMatchListItemsResponse

// NewGetMatchListItemsResponse instantiates a new GetMatchListItemsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMatchListItemsResponse(data GetMatchListItemsResponseData) *GetMatchListItemsResponse {
	this := GetMatchListItemsResponse{}
	this.Data = data
	return &this
}

// NewGetMatchListItemsResponseWithDefaults instantiates a new GetMatchListItemsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMatchListItemsResponseWithDefaults() *GetMatchListItemsResponse {
	this := GetMatchListItemsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetMatchListItemsResponse) GetData() GetMatchListItemsResponseData {
	if o == nil {
		var ret GetMatchListItemsResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetMatchListItemsResponse) GetDataOk() (*GetMatchListItemsResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetMatchListItemsResponse) SetData(v GetMatchListItemsResponseData) {
	o.Data = v
}

func (o GetMatchListItemsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMatchListItemsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetMatchListItemsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetMatchListItemsResponse := _GetMatchListItemsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetMatchListItemsResponse)

	if err != nil {
		return err
	}

	*o = GetMatchListItemsResponse(varGetMatchListItemsResponse)

	return err
}

type NullableGetMatchListItemsResponse struct {
	value *GetMatchListItemsResponse
	isSet bool
}

func (v NullableGetMatchListItemsResponse) Get() *GetMatchListItemsResponse {
	return v.value
}

func (v *NullableGetMatchListItemsResponse) Set(val *GetMatchListItemsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMatchListItemsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMatchListItemsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMatchListItemsResponse(val *GetMatchListItemsResponse) *NullableGetMatchListItemsResponse {
	return &NullableGetMatchListItemsResponse{value: val, isSet: true}
}

func (v NullableGetMatchListItemsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMatchListItemsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


