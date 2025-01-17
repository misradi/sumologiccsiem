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

// checks if the GetCustomMatchListColumnsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCustomMatchListColumnsResponse{}

// GetCustomMatchListColumnsResponse struct for GetCustomMatchListColumnsResponse
type GetCustomMatchListColumnsResponse struct {
	Data []CustomMatchListColumn `json:"data"`
}

type _GetCustomMatchListColumnsResponse GetCustomMatchListColumnsResponse

// NewGetCustomMatchListColumnsResponse instantiates a new GetCustomMatchListColumnsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCustomMatchListColumnsResponse(data []CustomMatchListColumn) *GetCustomMatchListColumnsResponse {
	this := GetCustomMatchListColumnsResponse{}
	this.Data = data
	return &this
}

// NewGetCustomMatchListColumnsResponseWithDefaults instantiates a new GetCustomMatchListColumnsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCustomMatchListColumnsResponseWithDefaults() *GetCustomMatchListColumnsResponse {
	this := GetCustomMatchListColumnsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetCustomMatchListColumnsResponse) GetData() []CustomMatchListColumn {
	if o == nil {
		var ret []CustomMatchListColumn
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCustomMatchListColumnsResponse) GetDataOk() ([]CustomMatchListColumn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetCustomMatchListColumnsResponse) SetData(v []CustomMatchListColumn) {
	o.Data = v
}

func (o GetCustomMatchListColumnsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCustomMatchListColumnsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetCustomMatchListColumnsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetCustomMatchListColumnsResponse := _GetCustomMatchListColumnsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCustomMatchListColumnsResponse)

	if err != nil {
		return err
	}

	*o = GetCustomMatchListColumnsResponse(varGetCustomMatchListColumnsResponse)

	return err
}

type NullableGetCustomMatchListColumnsResponse struct {
	value *GetCustomMatchListColumnsResponse
	isSet bool
}

func (v NullableGetCustomMatchListColumnsResponse) Get() *GetCustomMatchListColumnsResponse {
	return v.value
}

func (v *NullableGetCustomMatchListColumnsResponse) Set(val *GetCustomMatchListColumnsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCustomMatchListColumnsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCustomMatchListColumnsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCustomMatchListColumnsResponse(val *GetCustomMatchListColumnsResponse) *NullableGetCustomMatchListColumnsResponse {
	return &NullableGetCustomMatchListColumnsResponse{value: val, isSet: true}
}

func (v NullableGetCustomMatchListColumnsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCustomMatchListColumnsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


