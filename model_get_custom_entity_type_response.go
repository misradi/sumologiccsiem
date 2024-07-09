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

// checks if the GetCustomEntityTypeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCustomEntityTypeResponse{}

// GetCustomEntityTypeResponse struct for GetCustomEntityTypeResponse
type GetCustomEntityTypeResponse struct {
	Data CustomEntityType `json:"data"`
}

type _GetCustomEntityTypeResponse GetCustomEntityTypeResponse

// NewGetCustomEntityTypeResponse instantiates a new GetCustomEntityTypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCustomEntityTypeResponse(data CustomEntityType) *GetCustomEntityTypeResponse {
	this := GetCustomEntityTypeResponse{}
	this.Data = data
	return &this
}

// NewGetCustomEntityTypeResponseWithDefaults instantiates a new GetCustomEntityTypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCustomEntityTypeResponseWithDefaults() *GetCustomEntityTypeResponse {
	this := GetCustomEntityTypeResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetCustomEntityTypeResponse) GetData() CustomEntityType {
	if o == nil {
		var ret CustomEntityType
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCustomEntityTypeResponse) GetDataOk() (*CustomEntityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetCustomEntityTypeResponse) SetData(v CustomEntityType) {
	o.Data = v
}

func (o GetCustomEntityTypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCustomEntityTypeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetCustomEntityTypeResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetCustomEntityTypeResponse := _GetCustomEntityTypeResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCustomEntityTypeResponse)

	if err != nil {
		return err
	}

	*o = GetCustomEntityTypeResponse(varGetCustomEntityTypeResponse)

	return err
}

type NullableGetCustomEntityTypeResponse struct {
	value *GetCustomEntityTypeResponse
	isSet bool
}

func (v NullableGetCustomEntityTypeResponse) Get() *GetCustomEntityTypeResponse {
	return v.value
}

func (v *NullableGetCustomEntityTypeResponse) Set(val *GetCustomEntityTypeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCustomEntityTypeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCustomEntityTypeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCustomEntityTypeResponse(val *GetCustomEntityTypeResponse) *NullableGetCustomEntityTypeResponse {
	return &NullableGetCustomEntityTypeResponse{value: val, isSet: true}
}

func (v NullableGetCustomEntityTypeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCustomEntityTypeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


