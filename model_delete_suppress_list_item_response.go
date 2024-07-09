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

// checks if the DeleteSuppressListItemResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSuppressListItemResponse{}

// DeleteSuppressListItemResponse struct for DeleteSuppressListItemResponse
type DeleteSuppressListItemResponse struct {
	Data DeleteContextActionResponseData `json:"data"`
}

type _DeleteSuppressListItemResponse DeleteSuppressListItemResponse

// NewDeleteSuppressListItemResponse instantiates a new DeleteSuppressListItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSuppressListItemResponse(data DeleteContextActionResponseData) *DeleteSuppressListItemResponse {
	this := DeleteSuppressListItemResponse{}
	this.Data = data
	return &this
}

// NewDeleteSuppressListItemResponseWithDefaults instantiates a new DeleteSuppressListItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSuppressListItemResponseWithDefaults() *DeleteSuppressListItemResponse {
	this := DeleteSuppressListItemResponse{}
	return &this
}

// GetData returns the Data field value
func (o *DeleteSuppressListItemResponse) GetData() DeleteContextActionResponseData {
	if o == nil {
		var ret DeleteContextActionResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeleteSuppressListItemResponse) GetDataOk() (*DeleteContextActionResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeleteSuppressListItemResponse) SetData(v DeleteContextActionResponseData) {
	o.Data = v
}

func (o DeleteSuppressListItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteSuppressListItemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *DeleteSuppressListItemResponse) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteSuppressListItemResponse := _DeleteSuppressListItemResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteSuppressListItemResponse)

	if err != nil {
		return err
	}

	*o = DeleteSuppressListItemResponse(varDeleteSuppressListItemResponse)

	return err
}

type NullableDeleteSuppressListItemResponse struct {
	value *DeleteSuppressListItemResponse
	isSet bool
}

func (v NullableDeleteSuppressListItemResponse) Get() *DeleteSuppressListItemResponse {
	return v.value
}

func (v *NullableDeleteSuppressListItemResponse) Set(val *DeleteSuppressListItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSuppressListItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSuppressListItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSuppressListItemResponse(val *DeleteSuppressListItemResponse) *NullableDeleteSuppressListItemResponse {
	return &NullableDeleteSuppressListItemResponse{value: val, isSet: true}
}

func (v NullableDeleteSuppressListItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSuppressListItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


