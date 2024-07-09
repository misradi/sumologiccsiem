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

// checks if the GetAllSignalsResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllSignalsResponseData{}

// GetAllSignalsResponseData struct for GetAllSignalsResponseData
type GetAllSignalsResponseData struct {
	NextPageToken *string `json:"nextPageToken,omitempty"`
	Objects []Signal `json:"objects"`
}

type _GetAllSignalsResponseData GetAllSignalsResponseData

// NewGetAllSignalsResponseData instantiates a new GetAllSignalsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllSignalsResponseData(objects []Signal) *GetAllSignalsResponseData {
	this := GetAllSignalsResponseData{}
	this.Objects = objects
	return &this
}

// NewGetAllSignalsResponseDataWithDefaults instantiates a new GetAllSignalsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllSignalsResponseDataWithDefaults() *GetAllSignalsResponseData {
	this := GetAllSignalsResponseData{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *GetAllSignalsResponseData) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSignalsResponseData) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *GetAllSignalsResponseData) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *GetAllSignalsResponseData) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetObjects returns the Objects field value
func (o *GetAllSignalsResponseData) GetObjects() []Signal {
	if o == nil {
		var ret []Signal
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *GetAllSignalsResponseData) GetObjectsOk() ([]Signal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *GetAllSignalsResponseData) SetObjects(v []Signal) {
	o.Objects = v
}

func (o GetAllSignalsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllSignalsResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	toSerialize["objects"] = o.Objects
	return toSerialize, nil
}

func (o *GetAllSignalsResponseData) UnmarshalJSON(data []byte) (err error) {
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

	varGetAllSignalsResponseData := _GetAllSignalsResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAllSignalsResponseData)

	if err != nil {
		return err
	}

	*o = GetAllSignalsResponseData(varGetAllSignalsResponseData)

	return err
}

type NullableGetAllSignalsResponseData struct {
	value *GetAllSignalsResponseData
	isSet bool
}

func (v NullableGetAllSignalsResponseData) Get() *GetAllSignalsResponseData {
	return v.value
}

func (v *NullableGetAllSignalsResponseData) Set(val *GetAllSignalsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllSignalsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllSignalsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllSignalsResponseData(val *GetAllSignalsResponseData) *NullableGetAllSignalsResponseData {
	return &NullableGetAllSignalsResponseData{value: val, isSet: true}
}

func (v NullableGetAllSignalsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllSignalsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


