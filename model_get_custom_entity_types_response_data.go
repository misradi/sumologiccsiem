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

// checks if the GetCustomEntityTypesResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCustomEntityTypesResponseData{}

// GetCustomEntityTypesResponseData struct for GetCustomEntityTypesResponseData
type GetCustomEntityTypesResponseData struct {
	HasNextPage bool `json:"hasNextPage"`
	Total int32 `json:"total"`
	Objects []CustomEntityType `json:"objects"`
}

type _GetCustomEntityTypesResponseData GetCustomEntityTypesResponseData

// NewGetCustomEntityTypesResponseData instantiates a new GetCustomEntityTypesResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCustomEntityTypesResponseData(hasNextPage bool, total int32, objects []CustomEntityType) *GetCustomEntityTypesResponseData {
	this := GetCustomEntityTypesResponseData{}
	this.HasNextPage = hasNextPage
	this.Total = total
	this.Objects = objects
	return &this
}

// NewGetCustomEntityTypesResponseDataWithDefaults instantiates a new GetCustomEntityTypesResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCustomEntityTypesResponseDataWithDefaults() *GetCustomEntityTypesResponseData {
	this := GetCustomEntityTypesResponseData{}
	return &this
}

// GetHasNextPage returns the HasNextPage field value
func (o *GetCustomEntityTypesResponseData) GetHasNextPage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNextPage
}

// GetHasNextPageOk returns a tuple with the HasNextPage field value
// and a boolean to check if the value has been set.
func (o *GetCustomEntityTypesResponseData) GetHasNextPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasNextPage, true
}

// SetHasNextPage sets field value
func (o *GetCustomEntityTypesResponseData) SetHasNextPage(v bool) {
	o.HasNextPage = v
}

// GetTotal returns the Total field value
func (o *GetCustomEntityTypesResponseData) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GetCustomEntityTypesResponseData) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GetCustomEntityTypesResponseData) SetTotal(v int32) {
	o.Total = v
}

// GetObjects returns the Objects field value
func (o *GetCustomEntityTypesResponseData) GetObjects() []CustomEntityType {
	if o == nil {
		var ret []CustomEntityType
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *GetCustomEntityTypesResponseData) GetObjectsOk() ([]CustomEntityType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *GetCustomEntityTypesResponseData) SetObjects(v []CustomEntityType) {
	o.Objects = v
}

func (o GetCustomEntityTypesResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCustomEntityTypesResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hasNextPage"] = o.HasNextPage
	toSerialize["total"] = o.Total
	toSerialize["objects"] = o.Objects
	return toSerialize, nil
}

func (o *GetCustomEntityTypesResponseData) UnmarshalJSON(data []byte) (err error) {
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

	varGetCustomEntityTypesResponseData := _GetCustomEntityTypesResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCustomEntityTypesResponseData)

	if err != nil {
		return err
	}

	*o = GetCustomEntityTypesResponseData(varGetCustomEntityTypesResponseData)

	return err
}

type NullableGetCustomEntityTypesResponseData struct {
	value *GetCustomEntityTypesResponseData
	isSet bool
}

func (v NullableGetCustomEntityTypesResponseData) Get() *GetCustomEntityTypesResponseData {
	return v.value
}

func (v *NullableGetCustomEntityTypesResponseData) Set(val *GetCustomEntityTypesResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCustomEntityTypesResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCustomEntityTypesResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCustomEntityTypesResponseData(val *GetCustomEntityTypesResponseData) *NullableGetCustomEntityTypesResponseData {
	return &NullableGetCustomEntityTypesResponseData{value: val, isSet: true}
}

func (v NullableGetCustomEntityTypesResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCustomEntityTypesResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


