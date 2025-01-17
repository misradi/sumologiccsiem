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

// checks if the GetSuppressListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSuppressListResponse{}

// GetSuppressListResponse struct for GetSuppressListResponse
type GetSuppressListResponse struct {
	Data SuppressList `json:"data"`
}

type _GetSuppressListResponse GetSuppressListResponse

// NewGetSuppressListResponse instantiates a new GetSuppressListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSuppressListResponse(data SuppressList) *GetSuppressListResponse {
	this := GetSuppressListResponse{}
	this.Data = data
	return &this
}

// NewGetSuppressListResponseWithDefaults instantiates a new GetSuppressListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSuppressListResponseWithDefaults() *GetSuppressListResponse {
	this := GetSuppressListResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetSuppressListResponse) GetData() SuppressList {
	if o == nil {
		var ret SuppressList
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetSuppressListResponse) GetDataOk() (*SuppressList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetSuppressListResponse) SetData(v SuppressList) {
	o.Data = v
}

func (o GetSuppressListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSuppressListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetSuppressListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetSuppressListResponse := _GetSuppressListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSuppressListResponse)

	if err != nil {
		return err
	}

	*o = GetSuppressListResponse(varGetSuppressListResponse)

	return err
}

type NullableGetSuppressListResponse struct {
	value *GetSuppressListResponse
	isSet bool
}

func (v NullableGetSuppressListResponse) Get() *GetSuppressListResponse {
	return v.value
}

func (v *NullableGetSuppressListResponse) Set(val *GetSuppressListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSuppressListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSuppressListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSuppressListResponse(val *GetSuppressListResponse) *NullableGetSuppressListResponse {
	return &NullableGetSuppressListResponse{value: val, isSet: true}
}

func (v NullableGetSuppressListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSuppressListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


