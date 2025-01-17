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

// checks if the GetRelatedEntitiesByIdResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRelatedEntitiesByIdResponse{}

// GetRelatedEntitiesByIdResponse struct for GetRelatedEntitiesByIdResponse
type GetRelatedEntitiesByIdResponse struct {
	Data []GetRelatedEntitiesByIdResponseDataInner `json:"data"`
}

type _GetRelatedEntitiesByIdResponse GetRelatedEntitiesByIdResponse

// NewGetRelatedEntitiesByIdResponse instantiates a new GetRelatedEntitiesByIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRelatedEntitiesByIdResponse(data []GetRelatedEntitiesByIdResponseDataInner) *GetRelatedEntitiesByIdResponse {
	this := GetRelatedEntitiesByIdResponse{}
	this.Data = data
	return &this
}

// NewGetRelatedEntitiesByIdResponseWithDefaults instantiates a new GetRelatedEntitiesByIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRelatedEntitiesByIdResponseWithDefaults() *GetRelatedEntitiesByIdResponse {
	this := GetRelatedEntitiesByIdResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetRelatedEntitiesByIdResponse) GetData() []GetRelatedEntitiesByIdResponseDataInner {
	if o == nil {
		var ret []GetRelatedEntitiesByIdResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetRelatedEntitiesByIdResponse) GetDataOk() ([]GetRelatedEntitiesByIdResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetRelatedEntitiesByIdResponse) SetData(v []GetRelatedEntitiesByIdResponseDataInner) {
	o.Data = v
}

func (o GetRelatedEntitiesByIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRelatedEntitiesByIdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetRelatedEntitiesByIdResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetRelatedEntitiesByIdResponse := _GetRelatedEntitiesByIdResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRelatedEntitiesByIdResponse)

	if err != nil {
		return err
	}

	*o = GetRelatedEntitiesByIdResponse(varGetRelatedEntitiesByIdResponse)

	return err
}

type NullableGetRelatedEntitiesByIdResponse struct {
	value *GetRelatedEntitiesByIdResponse
	isSet bool
}

func (v NullableGetRelatedEntitiesByIdResponse) Get() *GetRelatedEntitiesByIdResponse {
	return v.value
}

func (v *NullableGetRelatedEntitiesByIdResponse) Set(val *GetRelatedEntitiesByIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRelatedEntitiesByIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRelatedEntitiesByIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRelatedEntitiesByIdResponse(val *GetRelatedEntitiesByIdResponse) *NullableGetRelatedEntitiesByIdResponse {
	return &NullableGetRelatedEntitiesByIdResponse{value: val, isSet: true}
}

func (v NullableGetRelatedEntitiesByIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRelatedEntitiesByIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


