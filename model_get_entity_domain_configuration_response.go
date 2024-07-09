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

// checks if the GetEntityDomainConfigurationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEntityDomainConfigurationResponse{}

// GetEntityDomainConfigurationResponse struct for GetEntityDomainConfigurationResponse
type GetEntityDomainConfigurationResponse struct {
	Data GetEntityDomainConfigurationResponseData `json:"data"`
}

type _GetEntityDomainConfigurationResponse GetEntityDomainConfigurationResponse

// NewGetEntityDomainConfigurationResponse instantiates a new GetEntityDomainConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEntityDomainConfigurationResponse(data GetEntityDomainConfigurationResponseData) *GetEntityDomainConfigurationResponse {
	this := GetEntityDomainConfigurationResponse{}
	this.Data = data
	return &this
}

// NewGetEntityDomainConfigurationResponseWithDefaults instantiates a new GetEntityDomainConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEntityDomainConfigurationResponseWithDefaults() *GetEntityDomainConfigurationResponse {
	this := GetEntityDomainConfigurationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetEntityDomainConfigurationResponse) GetData() GetEntityDomainConfigurationResponseData {
	if o == nil {
		var ret GetEntityDomainConfigurationResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetEntityDomainConfigurationResponse) GetDataOk() (*GetEntityDomainConfigurationResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetEntityDomainConfigurationResponse) SetData(v GetEntityDomainConfigurationResponseData) {
	o.Data = v
}

func (o GetEntityDomainConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEntityDomainConfigurationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetEntityDomainConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetEntityDomainConfigurationResponse := _GetEntityDomainConfigurationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEntityDomainConfigurationResponse)

	if err != nil {
		return err
	}

	*o = GetEntityDomainConfigurationResponse(varGetEntityDomainConfigurationResponse)

	return err
}

type NullableGetEntityDomainConfigurationResponse struct {
	value *GetEntityDomainConfigurationResponse
	isSet bool
}

func (v NullableGetEntityDomainConfigurationResponse) Get() *GetEntityDomainConfigurationResponse {
	return v.value
}

func (v *NullableGetEntityDomainConfigurationResponse) Set(val *GetEntityDomainConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEntityDomainConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEntityDomainConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEntityDomainConfigurationResponse(val *GetEntityDomainConfigurationResponse) *NullableGetEntityDomainConfigurationResponse {
	return &NullableGetEntityDomainConfigurationResponse{value: val, isSet: true}
}

func (v NullableGetEntityDomainConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEntityDomainConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


