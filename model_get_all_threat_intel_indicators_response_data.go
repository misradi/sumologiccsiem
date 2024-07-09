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

// checks if the GetAllThreatIntelIndicatorsResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllThreatIntelIndicatorsResponseData{}

// GetAllThreatIntelIndicatorsResponseData struct for GetAllThreatIntelIndicatorsResponseData
type GetAllThreatIntelIndicatorsResponseData struct {
	NextPageToken *string `json:"nextPageToken,omitempty"`
	Objects []GetThreatIntelIndicatorResponseData `json:"objects"`
}

type _GetAllThreatIntelIndicatorsResponseData GetAllThreatIntelIndicatorsResponseData

// NewGetAllThreatIntelIndicatorsResponseData instantiates a new GetAllThreatIntelIndicatorsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllThreatIntelIndicatorsResponseData(objects []GetThreatIntelIndicatorResponseData) *GetAllThreatIntelIndicatorsResponseData {
	this := GetAllThreatIntelIndicatorsResponseData{}
	this.Objects = objects
	return &this
}

// NewGetAllThreatIntelIndicatorsResponseDataWithDefaults instantiates a new GetAllThreatIntelIndicatorsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllThreatIntelIndicatorsResponseDataWithDefaults() *GetAllThreatIntelIndicatorsResponseData {
	this := GetAllThreatIntelIndicatorsResponseData{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *GetAllThreatIntelIndicatorsResponseData) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllThreatIntelIndicatorsResponseData) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *GetAllThreatIntelIndicatorsResponseData) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *GetAllThreatIntelIndicatorsResponseData) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetObjects returns the Objects field value
func (o *GetAllThreatIntelIndicatorsResponseData) GetObjects() []GetThreatIntelIndicatorResponseData {
	if o == nil {
		var ret []GetThreatIntelIndicatorResponseData
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *GetAllThreatIntelIndicatorsResponseData) GetObjectsOk() ([]GetThreatIntelIndicatorResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *GetAllThreatIntelIndicatorsResponseData) SetObjects(v []GetThreatIntelIndicatorResponseData) {
	o.Objects = v
}

func (o GetAllThreatIntelIndicatorsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllThreatIntelIndicatorsResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	toSerialize["objects"] = o.Objects
	return toSerialize, nil
}

func (o *GetAllThreatIntelIndicatorsResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varGetAllThreatIntelIndicatorsResponseData := _GetAllThreatIntelIndicatorsResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAllThreatIntelIndicatorsResponseData)

	if err != nil {
		return err
	}

	*o = GetAllThreatIntelIndicatorsResponseData(varGetAllThreatIntelIndicatorsResponseData)

	return err
}

type NullableGetAllThreatIntelIndicatorsResponseData struct {
	value *GetAllThreatIntelIndicatorsResponseData
	isSet bool
}

func (v NullableGetAllThreatIntelIndicatorsResponseData) Get() *GetAllThreatIntelIndicatorsResponseData {
	return v.value
}

func (v *NullableGetAllThreatIntelIndicatorsResponseData) Set(val *GetAllThreatIntelIndicatorsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllThreatIntelIndicatorsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllThreatIntelIndicatorsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllThreatIntelIndicatorsResponseData(val *GetAllThreatIntelIndicatorsResponseData) *NullableGetAllThreatIntelIndicatorsResponseData {
	return &NullableGetAllThreatIntelIndicatorsResponseData{value: val, isSet: true}
}

func (v NullableGetAllThreatIntelIndicatorsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllThreatIntelIndicatorsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


