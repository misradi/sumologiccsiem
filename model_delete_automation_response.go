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

// checks if the DeleteAutomationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteAutomationResponse{}

// DeleteAutomationResponse struct for DeleteAutomationResponse
type DeleteAutomationResponse struct {
	Data DeleteContextActionResponseData `json:"data"`
}

type _DeleteAutomationResponse DeleteAutomationResponse

// NewDeleteAutomationResponse instantiates a new DeleteAutomationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAutomationResponse(data DeleteContextActionResponseData) *DeleteAutomationResponse {
	this := DeleteAutomationResponse{}
	this.Data = data
	return &this
}

// NewDeleteAutomationResponseWithDefaults instantiates a new DeleteAutomationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAutomationResponseWithDefaults() *DeleteAutomationResponse {
	this := DeleteAutomationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *DeleteAutomationResponse) GetData() DeleteContextActionResponseData {
	if o == nil {
		var ret DeleteContextActionResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeleteAutomationResponse) GetDataOk() (*DeleteContextActionResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeleteAutomationResponse) SetData(v DeleteContextActionResponseData) {
	o.Data = v
}

func (o DeleteAutomationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteAutomationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *DeleteAutomationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteAutomationResponse := _DeleteAutomationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteAutomationResponse)

	if err != nil {
		return err
	}

	*o = DeleteAutomationResponse(varDeleteAutomationResponse)

	return err
}

type NullableDeleteAutomationResponse struct {
	value *DeleteAutomationResponse
	isSet bool
}

func (v NullableDeleteAutomationResponse) Get() *DeleteAutomationResponse {
	return v.value
}

func (v *NullableDeleteAutomationResponse) Set(val *DeleteAutomationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAutomationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAutomationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAutomationResponse(val *DeleteAutomationResponse) *NullableDeleteAutomationResponse {
	return &NullableDeleteAutomationResponse{value: val, isSet: true}
}

func (v NullableDeleteAutomationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAutomationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


