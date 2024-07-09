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

// checks if the CreateMatchListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMatchListResponse{}

// CreateMatchListResponse struct for CreateMatchListResponse
type CreateMatchListResponse struct {
	Data MatchList `json:"data"`
}

type _CreateMatchListResponse CreateMatchListResponse

// NewCreateMatchListResponse instantiates a new CreateMatchListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMatchListResponse(data MatchList) *CreateMatchListResponse {
	this := CreateMatchListResponse{}
	this.Data = data
	return &this
}

// NewCreateMatchListResponseWithDefaults instantiates a new CreateMatchListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMatchListResponseWithDefaults() *CreateMatchListResponse {
	this := CreateMatchListResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreateMatchListResponse) GetData() MatchList {
	if o == nil {
		var ret MatchList
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateMatchListResponse) GetDataOk() (*MatchList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateMatchListResponse) SetData(v MatchList) {
	o.Data = v
}

func (o CreateMatchListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMatchListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CreateMatchListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCreateMatchListResponse := _CreateMatchListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateMatchListResponse)

	if err != nil {
		return err
	}

	*o = CreateMatchListResponse(varCreateMatchListResponse)

	return err
}

type NullableCreateMatchListResponse struct {
	value *CreateMatchListResponse
	isSet bool
}

func (v NullableCreateMatchListResponse) Get() *CreateMatchListResponse {
	return v.value
}

func (v *NullableCreateMatchListResponse) Set(val *CreateMatchListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMatchListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMatchListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMatchListResponse(val *CreateMatchListResponse) *NullableCreateMatchListResponse {
	return &NullableCreateMatchListResponse{value: val, isSet: true}
}

func (v NullableCreateMatchListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMatchListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


