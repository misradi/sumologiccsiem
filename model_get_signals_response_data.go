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

// checks if the GetSignalsResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSignalsResponseData{}

// GetSignalsResponseData struct for GetSignalsResponseData
type GetSignalsResponseData struct {
	HasNextPage bool `json:"hasNextPage"`
	Total int32 `json:"total"`
	Objects []Signal `json:"objects"`
}

type _GetSignalsResponseData GetSignalsResponseData

// NewGetSignalsResponseData instantiates a new GetSignalsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignalsResponseData(hasNextPage bool, total int32, objects []Signal) *GetSignalsResponseData {
	this := GetSignalsResponseData{}
	this.HasNextPage = hasNextPage
	this.Total = total
	this.Objects = objects
	return &this
}

// NewGetSignalsResponseDataWithDefaults instantiates a new GetSignalsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignalsResponseDataWithDefaults() *GetSignalsResponseData {
	this := GetSignalsResponseData{}
	return &this
}

// GetHasNextPage returns the HasNextPage field value
func (o *GetSignalsResponseData) GetHasNextPage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNextPage
}

// GetHasNextPageOk returns a tuple with the HasNextPage field value
// and a boolean to check if the value has been set.
func (o *GetSignalsResponseData) GetHasNextPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasNextPage, true
}

// SetHasNextPage sets field value
func (o *GetSignalsResponseData) SetHasNextPage(v bool) {
	o.HasNextPage = v
}

// GetTotal returns the Total field value
func (o *GetSignalsResponseData) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GetSignalsResponseData) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GetSignalsResponseData) SetTotal(v int32) {
	o.Total = v
}

// GetObjects returns the Objects field value
func (o *GetSignalsResponseData) GetObjects() []Signal {
	if o == nil {
		var ret []Signal
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *GetSignalsResponseData) GetObjectsOk() ([]Signal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *GetSignalsResponseData) SetObjects(v []Signal) {
	o.Objects = v
}

func (o GetSignalsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSignalsResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hasNextPage"] = o.HasNextPage
	toSerialize["total"] = o.Total
	toSerialize["objects"] = o.Objects
	return toSerialize, nil
}

func (o *GetSignalsResponseData) UnmarshalJSON(data []byte) (err error) {
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

	varGetSignalsResponseData := _GetSignalsResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSignalsResponseData)

	if err != nil {
		return err
	}

	*o = GetSignalsResponseData(varGetSignalsResponseData)

	return err
}

type NullableGetSignalsResponseData struct {
	value *GetSignalsResponseData
	isSet bool
}

func (v NullableGetSignalsResponseData) Get() *GetSignalsResponseData {
	return v.value
}

func (v *NullableGetSignalsResponseData) Set(val *GetSignalsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignalsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignalsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignalsResponseData(val *GetSignalsResponseData) *NullableGetSignalsResponseData {
	return &NullableGetSignalsResponseData{value: val, isSet: true}
}

func (v NullableGetSignalsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignalsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


