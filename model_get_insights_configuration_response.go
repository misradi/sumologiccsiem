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

// checks if the GetInsightsConfigurationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsightsConfigurationResponse{}

// GetInsightsConfigurationResponse struct for GetInsightsConfigurationResponse
type GetInsightsConfigurationResponse struct {
	Data GetInsightsConfigurationResponseData `json:"data"`
}

type _GetInsightsConfigurationResponse GetInsightsConfigurationResponse

// NewGetInsightsConfigurationResponse instantiates a new GetInsightsConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsightsConfigurationResponse(data GetInsightsConfigurationResponseData) *GetInsightsConfigurationResponse {
	this := GetInsightsConfigurationResponse{}
	this.Data = data
	return &this
}

// NewGetInsightsConfigurationResponseWithDefaults instantiates a new GetInsightsConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsightsConfigurationResponseWithDefaults() *GetInsightsConfigurationResponse {
	this := GetInsightsConfigurationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetInsightsConfigurationResponse) GetData() GetInsightsConfigurationResponseData {
	if o == nil {
		var ret GetInsightsConfigurationResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetInsightsConfigurationResponse) GetDataOk() (*GetInsightsConfigurationResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetInsightsConfigurationResponse) SetData(v GetInsightsConfigurationResponseData) {
	o.Data = v
}

func (o GetInsightsConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsightsConfigurationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetInsightsConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetInsightsConfigurationResponse := _GetInsightsConfigurationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetInsightsConfigurationResponse)

	if err != nil {
		return err
	}

	*o = GetInsightsConfigurationResponse(varGetInsightsConfigurationResponse)

	return err
}

type NullableGetInsightsConfigurationResponse struct {
	value *GetInsightsConfigurationResponse
	isSet bool
}

func (v NullableGetInsightsConfigurationResponse) Get() *GetInsightsConfigurationResponse {
	return v.value
}

func (v *NullableGetInsightsConfigurationResponse) Set(val *GetInsightsConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsightsConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsightsConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsightsConfigurationResponse(val *GetInsightsConfigurationResponse) *NullableGetInsightsConfigurationResponse {
	return &NullableGetInsightsConfigurationResponse{value: val, isSet: true}
}

func (v NullableGetInsightsConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsightsConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


