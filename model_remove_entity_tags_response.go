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

// checks if the RemoveEntityTagsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveEntityTagsResponse{}

// RemoveEntityTagsResponse struct for RemoveEntityTagsResponse
type RemoveEntityTagsResponse struct {
	Data UpdateEntityCriticalityResponseData `json:"data"`
}

type _RemoveEntityTagsResponse RemoveEntityTagsResponse

// NewRemoveEntityTagsResponse instantiates a new RemoveEntityTagsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveEntityTagsResponse(data UpdateEntityCriticalityResponseData) *RemoveEntityTagsResponse {
	this := RemoveEntityTagsResponse{}
	this.Data = data
	return &this
}

// NewRemoveEntityTagsResponseWithDefaults instantiates a new RemoveEntityTagsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveEntityTagsResponseWithDefaults() *RemoveEntityTagsResponse {
	this := RemoveEntityTagsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *RemoveEntityTagsResponse) GetData() UpdateEntityCriticalityResponseData {
	if o == nil {
		var ret UpdateEntityCriticalityResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RemoveEntityTagsResponse) GetDataOk() (*UpdateEntityCriticalityResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RemoveEntityTagsResponse) SetData(v UpdateEntityCriticalityResponseData) {
	o.Data = v
}

func (o RemoveEntityTagsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveEntityTagsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *RemoveEntityTagsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varRemoveEntityTagsResponse := _RemoveEntityTagsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRemoveEntityTagsResponse)

	if err != nil {
		return err
	}

	*o = RemoveEntityTagsResponse(varRemoveEntityTagsResponse)

	return err
}

type NullableRemoveEntityTagsResponse struct {
	value *RemoveEntityTagsResponse
	isSet bool
}

func (v NullableRemoveEntityTagsResponse) Get() *RemoveEntityTagsResponse {
	return v.value
}

func (v *NullableRemoveEntityTagsResponse) Set(val *RemoveEntityTagsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveEntityTagsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveEntityTagsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveEntityTagsResponse(val *RemoveEntityTagsResponse) *NullableRemoveEntityTagsResponse {
	return &NullableRemoveEntityTagsResponse{value: val, isSet: true}
}

func (v NullableRemoveEntityTagsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveEntityTagsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


