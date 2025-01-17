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

// checks if the CreateEntityCriticalityConfigResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEntityCriticalityConfigResponse{}

// CreateEntityCriticalityConfigResponse struct for CreateEntityCriticalityConfigResponse
type CreateEntityCriticalityConfigResponse struct {
	Data EntityCriticalityConfig `json:"data"`
}

type _CreateEntityCriticalityConfigResponse CreateEntityCriticalityConfigResponse

// NewCreateEntityCriticalityConfigResponse instantiates a new CreateEntityCriticalityConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEntityCriticalityConfigResponse(data EntityCriticalityConfig) *CreateEntityCriticalityConfigResponse {
	this := CreateEntityCriticalityConfigResponse{}
	this.Data = data
	return &this
}

// NewCreateEntityCriticalityConfigResponseWithDefaults instantiates a new CreateEntityCriticalityConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEntityCriticalityConfigResponseWithDefaults() *CreateEntityCriticalityConfigResponse {
	this := CreateEntityCriticalityConfigResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreateEntityCriticalityConfigResponse) GetData() EntityCriticalityConfig {
	if o == nil {
		var ret EntityCriticalityConfig
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateEntityCriticalityConfigResponse) GetDataOk() (*EntityCriticalityConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateEntityCriticalityConfigResponse) SetData(v EntityCriticalityConfig) {
	o.Data = v
}

func (o CreateEntityCriticalityConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEntityCriticalityConfigResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CreateEntityCriticalityConfigResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCreateEntityCriticalityConfigResponse := _CreateEntityCriticalityConfigResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateEntityCriticalityConfigResponse)

	if err != nil {
		return err
	}

	*o = CreateEntityCriticalityConfigResponse(varCreateEntityCriticalityConfigResponse)

	return err
}

type NullableCreateEntityCriticalityConfigResponse struct {
	value *CreateEntityCriticalityConfigResponse
	isSet bool
}

func (v NullableCreateEntityCriticalityConfigResponse) Get() *CreateEntityCriticalityConfigResponse {
	return v.value
}

func (v *NullableCreateEntityCriticalityConfigResponse) Set(val *CreateEntityCriticalityConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEntityCriticalityConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEntityCriticalityConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEntityCriticalityConfigResponse(val *CreateEntityCriticalityConfigResponse) *NullableCreateEntityCriticalityConfigResponse {
	return &NullableCreateEntityCriticalityConfigResponse{value: val, isSet: true}
}

func (v NullableCreateEntityCriticalityConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEntityCriticalityConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


