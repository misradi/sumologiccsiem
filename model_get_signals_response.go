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

// checks if the GetSignalsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSignalsResponse{}

// GetSignalsResponse struct for GetSignalsResponse
type GetSignalsResponse struct {
	Data GetSignalsResponseData `json:"data"`
}

type _GetSignalsResponse GetSignalsResponse

// NewGetSignalsResponse instantiates a new GetSignalsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignalsResponse(data GetSignalsResponseData) *GetSignalsResponse {
	this := GetSignalsResponse{}
	this.Data = data
	return &this
}

// NewGetSignalsResponseWithDefaults instantiates a new GetSignalsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignalsResponseWithDefaults() *GetSignalsResponse {
	this := GetSignalsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetSignalsResponse) GetData() GetSignalsResponseData {
	if o == nil {
		var ret GetSignalsResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetSignalsResponse) GetDataOk() (*GetSignalsResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetSignalsResponse) SetData(v GetSignalsResponseData) {
	o.Data = v
}

func (o GetSignalsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSignalsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetSignalsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetSignalsResponse := _GetSignalsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSignalsResponse)

	if err != nil {
		return err
	}

	*o = GetSignalsResponse(varGetSignalsResponse)

	return err
}

type NullableGetSignalsResponse struct {
	value *GetSignalsResponse
	isSet bool
}

func (v NullableGetSignalsResponse) Get() *GetSignalsResponse {
	return v.value
}

func (v *NullableGetSignalsResponse) Set(val *GetSignalsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignalsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignalsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignalsResponse(val *GetSignalsResponse) *NullableGetSignalsResponse {
	return &NullableGetSignalsResponse{value: val, isSet: true}
}

func (v NullableGetSignalsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignalsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


