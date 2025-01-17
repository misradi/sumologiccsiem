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

// checks if the UpdateEntitySuppressedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEntitySuppressedResponse{}

// UpdateEntitySuppressedResponse struct for UpdateEntitySuppressedResponse
type UpdateEntitySuppressedResponse struct {
	Data UpdateEntityCriticalityResponseData `json:"data"`
}

type _UpdateEntitySuppressedResponse UpdateEntitySuppressedResponse

// NewUpdateEntitySuppressedResponse instantiates a new UpdateEntitySuppressedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEntitySuppressedResponse(data UpdateEntityCriticalityResponseData) *UpdateEntitySuppressedResponse {
	this := UpdateEntitySuppressedResponse{}
	this.Data = data
	return &this
}

// NewUpdateEntitySuppressedResponseWithDefaults instantiates a new UpdateEntitySuppressedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEntitySuppressedResponseWithDefaults() *UpdateEntitySuppressedResponse {
	this := UpdateEntitySuppressedResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateEntitySuppressedResponse) GetData() UpdateEntityCriticalityResponseData {
	if o == nil {
		var ret UpdateEntityCriticalityResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateEntitySuppressedResponse) GetDataOk() (*UpdateEntityCriticalityResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateEntitySuppressedResponse) SetData(v UpdateEntityCriticalityResponseData) {
	o.Data = v
}

func (o UpdateEntitySuppressedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEntitySuppressedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateEntitySuppressedResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateEntitySuppressedResponse := _UpdateEntitySuppressedResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateEntitySuppressedResponse)

	if err != nil {
		return err
	}

	*o = UpdateEntitySuppressedResponse(varUpdateEntitySuppressedResponse)

	return err
}

type NullableUpdateEntitySuppressedResponse struct {
	value *UpdateEntitySuppressedResponse
	isSet bool
}

func (v NullableUpdateEntitySuppressedResponse) Get() *UpdateEntitySuppressedResponse {
	return v.value
}

func (v *NullableUpdateEntitySuppressedResponse) Set(val *UpdateEntitySuppressedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEntitySuppressedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEntitySuppressedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEntitySuppressedResponse(val *UpdateEntitySuppressedResponse) *NullableUpdateEntitySuppressedResponse {
	return &NullableUpdateEntitySuppressedResponse{value: val, isSet: true}
}

func (v NullableUpdateEntitySuppressedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEntitySuppressedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


