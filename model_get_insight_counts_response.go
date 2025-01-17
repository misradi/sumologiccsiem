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

// checks if the GetInsightCountsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsightCountsResponse{}

// GetInsightCountsResponse struct for GetInsightCountsResponse
type GetInsightCountsResponse struct {
	Data []GetInsightCountsResponseDataInner `json:"data"`
}

type _GetInsightCountsResponse GetInsightCountsResponse

// NewGetInsightCountsResponse instantiates a new GetInsightCountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsightCountsResponse(data []GetInsightCountsResponseDataInner) *GetInsightCountsResponse {
	this := GetInsightCountsResponse{}
	this.Data = data
	return &this
}

// NewGetInsightCountsResponseWithDefaults instantiates a new GetInsightCountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsightCountsResponseWithDefaults() *GetInsightCountsResponse {
	this := GetInsightCountsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetInsightCountsResponse) GetData() []GetInsightCountsResponseDataInner {
	if o == nil {
		var ret []GetInsightCountsResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetInsightCountsResponse) GetDataOk() ([]GetInsightCountsResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetInsightCountsResponse) SetData(v []GetInsightCountsResponseDataInner) {
	o.Data = v
}

func (o GetInsightCountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsightCountsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetInsightCountsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetInsightCountsResponse := _GetInsightCountsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetInsightCountsResponse)

	if err != nil {
		return err
	}

	*o = GetInsightCountsResponse(varGetInsightCountsResponse)

	return err
}

type NullableGetInsightCountsResponse struct {
	value *GetInsightCountsResponse
	isSet bool
}

func (v NullableGetInsightCountsResponse) Get() *GetInsightCountsResponse {
	return v.value
}

func (v *NullableGetInsightCountsResponse) Set(val *GetInsightCountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsightCountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsightCountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsightCountsResponse(val *GetInsightCountsResponse) *NullableGetInsightCountsResponse {
	return &NullableGetInsightCountsResponse{value: val, isSet: true}
}

func (v NullableGetInsightCountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsightCountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


