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

// checks if the GetContextActionsResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetContextActionsResponseData{}

// GetContextActionsResponseData struct for GetContextActionsResponseData
type GetContextActionsResponseData struct {
	Total int32 `json:"total"`
	HasNextPage bool `json:"hasNextPage"`
	Objects []ContextAction `json:"objects"`
}

type _GetContextActionsResponseData GetContextActionsResponseData

// NewGetContextActionsResponseData instantiates a new GetContextActionsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetContextActionsResponseData(total int32, hasNextPage bool, objects []ContextAction) *GetContextActionsResponseData {
	this := GetContextActionsResponseData{}
	this.Total = total
	this.HasNextPage = hasNextPage
	this.Objects = objects
	return &this
}

// NewGetContextActionsResponseDataWithDefaults instantiates a new GetContextActionsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetContextActionsResponseDataWithDefaults() *GetContextActionsResponseData {
	this := GetContextActionsResponseData{}
	return &this
}

// GetTotal returns the Total field value
func (o *GetContextActionsResponseData) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GetContextActionsResponseData) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GetContextActionsResponseData) SetTotal(v int32) {
	o.Total = v
}

// GetHasNextPage returns the HasNextPage field value
func (o *GetContextActionsResponseData) GetHasNextPage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNextPage
}

// GetHasNextPageOk returns a tuple with the HasNextPage field value
// and a boolean to check if the value has been set.
func (o *GetContextActionsResponseData) GetHasNextPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasNextPage, true
}

// SetHasNextPage sets field value
func (o *GetContextActionsResponseData) SetHasNextPage(v bool) {
	o.HasNextPage = v
}

// GetObjects returns the Objects field value
func (o *GetContextActionsResponseData) GetObjects() []ContextAction {
	if o == nil {
		var ret []ContextAction
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *GetContextActionsResponseData) GetObjectsOk() ([]ContextAction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *GetContextActionsResponseData) SetObjects(v []ContextAction) {
	o.Objects = v
}

func (o GetContextActionsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetContextActionsResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	toSerialize["hasNextPage"] = o.HasNextPage
	toSerialize["objects"] = o.Objects
	return toSerialize, nil
}

func (o *GetContextActionsResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total",
		"hasNextPage",
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

	varGetContextActionsResponseData := _GetContextActionsResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetContextActionsResponseData)

	if err != nil {
		return err
	}

	*o = GetContextActionsResponseData(varGetContextActionsResponseData)

	return err
}

type NullableGetContextActionsResponseData struct {
	value *GetContextActionsResponseData
	isSet bool
}

func (v NullableGetContextActionsResponseData) Get() *GetContextActionsResponseData {
	return v.value
}

func (v *NullableGetContextActionsResponseData) Set(val *GetContextActionsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetContextActionsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetContextActionsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetContextActionsResponseData(val *GetContextActionsResponseData) *NullableGetContextActionsResponseData {
	return &NullableGetContextActionsResponseData{value: val, isSet: true}
}

func (v NullableGetContextActionsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetContextActionsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


