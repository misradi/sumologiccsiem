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

// checks if the GetEntitiesResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEntitiesResponseData{}

// GetEntitiesResponseData struct for GetEntitiesResponseData
type GetEntitiesResponseData struct {
	HasNextPage bool `json:"hasNextPage"`
	Total int32 `json:"total"`
	Objects []GetEntitiesResponseDataObjectsInner `json:"objects"`
}

type _GetEntitiesResponseData GetEntitiesResponseData

// NewGetEntitiesResponseData instantiates a new GetEntitiesResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEntitiesResponseData(hasNextPage bool, total int32, objects []GetEntitiesResponseDataObjectsInner) *GetEntitiesResponseData {
	this := GetEntitiesResponseData{}
	this.HasNextPage = hasNextPage
	this.Total = total
	this.Objects = objects
	return &this
}

// NewGetEntitiesResponseDataWithDefaults instantiates a new GetEntitiesResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEntitiesResponseDataWithDefaults() *GetEntitiesResponseData {
	this := GetEntitiesResponseData{}
	return &this
}

// GetHasNextPage returns the HasNextPage field value
func (o *GetEntitiesResponseData) GetHasNextPage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNextPage
}

// GetHasNextPageOk returns a tuple with the HasNextPage field value
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseData) GetHasNextPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasNextPage, true
}

// SetHasNextPage sets field value
func (o *GetEntitiesResponseData) SetHasNextPage(v bool) {
	o.HasNextPage = v
}

// GetTotal returns the Total field value
func (o *GetEntitiesResponseData) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseData) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GetEntitiesResponseData) SetTotal(v int32) {
	o.Total = v
}

// GetObjects returns the Objects field value
func (o *GetEntitiesResponseData) GetObjects() []GetEntitiesResponseDataObjectsInner {
	if o == nil {
		var ret []GetEntitiesResponseDataObjectsInner
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseData) GetObjectsOk() ([]GetEntitiesResponseDataObjectsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *GetEntitiesResponseData) SetObjects(v []GetEntitiesResponseDataObjectsInner) {
	o.Objects = v
}

func (o GetEntitiesResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEntitiesResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hasNextPage"] = o.HasNextPage
	toSerialize["total"] = o.Total
	toSerialize["objects"] = o.Objects
	return toSerialize, nil
}

func (o *GetEntitiesResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hasNextPage",
		"total",
		"objects",
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

	varGetEntitiesResponseData := _GetEntitiesResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEntitiesResponseData)

	if err != nil {
		return err
	}

	*o = GetEntitiesResponseData(varGetEntitiesResponseData)

	return err
}

type NullableGetEntitiesResponseData struct {
	value *GetEntitiesResponseData
	isSet bool
}

func (v NullableGetEntitiesResponseData) Get() *GetEntitiesResponseData {
	return v.value
}

func (v *NullableGetEntitiesResponseData) Set(val *GetEntitiesResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEntitiesResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEntitiesResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEntitiesResponseData(val *GetEntitiesResponseData) *NullableGetEntitiesResponseData {
	return &NullableGetEntitiesResponseData{value: val, isSet: true}
}

func (v NullableGetEntitiesResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEntitiesResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

