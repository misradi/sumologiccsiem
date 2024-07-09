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

// checks if the GetSuppressListsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSuppressListsResponse{}

// GetSuppressListsResponse struct for GetSuppressListsResponse
type GetSuppressListsResponse struct {
	Data GetSuppressListsResponseData `json:"data"`
}

type _GetSuppressListsResponse GetSuppressListsResponse

// NewGetSuppressListsResponse instantiates a new GetSuppressListsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSuppressListsResponse(data GetSuppressListsResponseData) *GetSuppressListsResponse {
	this := GetSuppressListsResponse{}
	this.Data = data
	return &this
}

// NewGetSuppressListsResponseWithDefaults instantiates a new GetSuppressListsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSuppressListsResponseWithDefaults() *GetSuppressListsResponse {
	this := GetSuppressListsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetSuppressListsResponse) GetData() GetSuppressListsResponseData {
	if o == nil {
		var ret GetSuppressListsResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetSuppressListsResponse) GetDataOk() (*GetSuppressListsResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetSuppressListsResponse) SetData(v GetSuppressListsResponseData) {
	o.Data = v
}

func (o GetSuppressListsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSuppressListsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetSuppressListsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetSuppressListsResponse := _GetSuppressListsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSuppressListsResponse)

	if err != nil {
		return err
	}

	*o = GetSuppressListsResponse(varGetSuppressListsResponse)

	return err
}

type NullableGetSuppressListsResponse struct {
	value *GetSuppressListsResponse
	isSet bool
}

func (v NullableGetSuppressListsResponse) Get() *GetSuppressListsResponse {
	return v.value
}

func (v *NullableGetSuppressListsResponse) Set(val *GetSuppressListsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSuppressListsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSuppressListsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSuppressListsResponse(val *GetSuppressListsResponse) *NullableGetSuppressListsResponse {
	return &NullableGetSuppressListsResponse{value: val, isSet: true}
}

func (v NullableGetSuppressListsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSuppressListsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


