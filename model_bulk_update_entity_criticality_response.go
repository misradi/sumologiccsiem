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

// checks if the BulkUpdateEntityCriticalityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkUpdateEntityCriticalityResponse{}

// BulkUpdateEntityCriticalityResponse struct for BulkUpdateEntityCriticalityResponse
type BulkUpdateEntityCriticalityResponse struct {
	Data []BulkUpdateEntityCriticalityResponseDataInner `json:"data"`
}

type _BulkUpdateEntityCriticalityResponse BulkUpdateEntityCriticalityResponse

// NewBulkUpdateEntityCriticalityResponse instantiates a new BulkUpdateEntityCriticalityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateEntityCriticalityResponse(data []BulkUpdateEntityCriticalityResponseDataInner) *BulkUpdateEntityCriticalityResponse {
	this := BulkUpdateEntityCriticalityResponse{}
	this.Data = data
	return &this
}

// NewBulkUpdateEntityCriticalityResponseWithDefaults instantiates a new BulkUpdateEntityCriticalityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateEntityCriticalityResponseWithDefaults() *BulkUpdateEntityCriticalityResponse {
	this := BulkUpdateEntityCriticalityResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BulkUpdateEntityCriticalityResponse) GetData() []BulkUpdateEntityCriticalityResponseDataInner {
	if o == nil {
		var ret []BulkUpdateEntityCriticalityResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BulkUpdateEntityCriticalityResponse) GetDataOk() ([]BulkUpdateEntityCriticalityResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BulkUpdateEntityCriticalityResponse) SetData(v []BulkUpdateEntityCriticalityResponseDataInner) {
	o.Data = v
}

func (o BulkUpdateEntityCriticalityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkUpdateEntityCriticalityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *BulkUpdateEntityCriticalityResponse) UnmarshalJSON(data []byte) (err error) {
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

	varBulkUpdateEntityCriticalityResponse := _BulkUpdateEntityCriticalityResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBulkUpdateEntityCriticalityResponse)

	if err != nil {
		return err
	}

	*o = BulkUpdateEntityCriticalityResponse(varBulkUpdateEntityCriticalityResponse)

	return err
}

type NullableBulkUpdateEntityCriticalityResponse struct {
	value *BulkUpdateEntityCriticalityResponse
	isSet bool
}

func (v NullableBulkUpdateEntityCriticalityResponse) Get() *BulkUpdateEntityCriticalityResponse {
	return v.value
}

func (v *NullableBulkUpdateEntityCriticalityResponse) Set(val *BulkUpdateEntityCriticalityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkUpdateEntityCriticalityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkUpdateEntityCriticalityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkUpdateEntityCriticalityResponse(val *BulkUpdateEntityCriticalityResponse) *NullableBulkUpdateEntityCriticalityResponse {
	return &NullableBulkUpdateEntityCriticalityResponse{value: val, isSet: true}
}

func (v NullableBulkUpdateEntityCriticalityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkUpdateEntityCriticalityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

