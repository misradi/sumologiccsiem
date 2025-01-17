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

// checks if the UpdateEntityCriticalityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEntityCriticalityResponse{}

// UpdateEntityCriticalityResponse struct for UpdateEntityCriticalityResponse
type UpdateEntityCriticalityResponse struct {
	Data UpdateEntityCriticalityResponseData `json:"data"`
}

type _UpdateEntityCriticalityResponse UpdateEntityCriticalityResponse

// NewUpdateEntityCriticalityResponse instantiates a new UpdateEntityCriticalityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEntityCriticalityResponse(data UpdateEntityCriticalityResponseData) *UpdateEntityCriticalityResponse {
	this := UpdateEntityCriticalityResponse{}
	this.Data = data
	return &this
}

// NewUpdateEntityCriticalityResponseWithDefaults instantiates a new UpdateEntityCriticalityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEntityCriticalityResponseWithDefaults() *UpdateEntityCriticalityResponse {
	this := UpdateEntityCriticalityResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateEntityCriticalityResponse) GetData() UpdateEntityCriticalityResponseData {
	if o == nil {
		var ret UpdateEntityCriticalityResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateEntityCriticalityResponse) GetDataOk() (*UpdateEntityCriticalityResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateEntityCriticalityResponse) SetData(v UpdateEntityCriticalityResponseData) {
	o.Data = v
}

func (o UpdateEntityCriticalityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEntityCriticalityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateEntityCriticalityResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateEntityCriticalityResponse := _UpdateEntityCriticalityResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateEntityCriticalityResponse)

	if err != nil {
		return err
	}

	*o = UpdateEntityCriticalityResponse(varUpdateEntityCriticalityResponse)

	return err
}

type NullableUpdateEntityCriticalityResponse struct {
	value *UpdateEntityCriticalityResponse
	isSet bool
}

func (v NullableUpdateEntityCriticalityResponse) Get() *UpdateEntityCriticalityResponse {
	return v.value
}

func (v *NullableUpdateEntityCriticalityResponse) Set(val *UpdateEntityCriticalityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEntityCriticalityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEntityCriticalityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEntityCriticalityResponse(val *UpdateEntityCriticalityResponse) *NullableUpdateEntityCriticalityResponse {
	return &NullableUpdateEntityCriticalityResponse{value: val, isSet: true}
}

func (v NullableUpdateEntityCriticalityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEntityCriticalityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


