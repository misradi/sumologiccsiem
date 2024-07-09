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

// checks if the CreateTagSchemaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTagSchemaResponse{}

// CreateTagSchemaResponse struct for CreateTagSchemaResponse
type CreateTagSchemaResponse struct {
	Data GetTagSchemasResponseDataInner `json:"data"`
}

type _CreateTagSchemaResponse CreateTagSchemaResponse

// NewCreateTagSchemaResponse instantiates a new CreateTagSchemaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTagSchemaResponse(data GetTagSchemasResponseDataInner) *CreateTagSchemaResponse {
	this := CreateTagSchemaResponse{}
	this.Data = data
	return &this
}

// NewCreateTagSchemaResponseWithDefaults instantiates a new CreateTagSchemaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTagSchemaResponseWithDefaults() *CreateTagSchemaResponse {
	this := CreateTagSchemaResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreateTagSchemaResponse) GetData() GetTagSchemasResponseDataInner {
	if o == nil {
		var ret GetTagSchemasResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateTagSchemaResponse) GetDataOk() (*GetTagSchemasResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateTagSchemaResponse) SetData(v GetTagSchemasResponseDataInner) {
	o.Data = v
}

func (o CreateTagSchemaResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTagSchemaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CreateTagSchemaResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCreateTagSchemaResponse := _CreateTagSchemaResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTagSchemaResponse)

	if err != nil {
		return err
	}

	*o = CreateTagSchemaResponse(varCreateTagSchemaResponse)

	return err
}

type NullableCreateTagSchemaResponse struct {
	value *CreateTagSchemaResponse
	isSet bool
}

func (v NullableCreateTagSchemaResponse) Get() *CreateTagSchemaResponse {
	return v.value
}

func (v *NullableCreateTagSchemaResponse) Set(val *CreateTagSchemaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTagSchemaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTagSchemaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTagSchemaResponse(val *CreateTagSchemaResponse) *NullableCreateTagSchemaResponse {
	return &NullableCreateTagSchemaResponse{value: val, isSet: true}
}

func (v NullableCreateTagSchemaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTagSchemaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


