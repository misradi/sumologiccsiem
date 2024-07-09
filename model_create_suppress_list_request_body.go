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

// checks if the CreateSuppressListRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSuppressListRequestBody{}

// CreateSuppressListRequestBody struct for CreateSuppressListRequestBody
type CreateSuppressListRequestBody struct {
	Fields CreateSuppressListRequestBodyFields `json:"fields"`
}

type _CreateSuppressListRequestBody CreateSuppressListRequestBody

// NewCreateSuppressListRequestBody instantiates a new CreateSuppressListRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSuppressListRequestBody(fields CreateSuppressListRequestBodyFields) *CreateSuppressListRequestBody {
	this := CreateSuppressListRequestBody{}
	this.Fields = fields
	return &this
}

// NewCreateSuppressListRequestBodyWithDefaults instantiates a new CreateSuppressListRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSuppressListRequestBodyWithDefaults() *CreateSuppressListRequestBody {
	this := CreateSuppressListRequestBody{}
	return &this
}

// GetFields returns the Fields field value
func (o *CreateSuppressListRequestBody) GetFields() CreateSuppressListRequestBodyFields {
	if o == nil {
		var ret CreateSuppressListRequestBodyFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *CreateSuppressListRequestBody) GetFieldsOk() (*CreateSuppressListRequestBodyFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *CreateSuppressListRequestBody) SetFields(v CreateSuppressListRequestBodyFields) {
	o.Fields = v
}

func (o CreateSuppressListRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSuppressListRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *CreateSuppressListRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fields",
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

	varCreateSuppressListRequestBody := _CreateSuppressListRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateSuppressListRequestBody)

	if err != nil {
		return err
	}

	*o = CreateSuppressListRequestBody(varCreateSuppressListRequestBody)

	return err
}

type NullableCreateSuppressListRequestBody struct {
	value *CreateSuppressListRequestBody
	isSet bool
}

func (v NullableCreateSuppressListRequestBody) Get() *CreateSuppressListRequestBody {
	return v.value
}

func (v *NullableCreateSuppressListRequestBody) Set(val *CreateSuppressListRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSuppressListRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSuppressListRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSuppressListRequestBody(val *CreateSuppressListRequestBody) *NullableCreateSuppressListRequestBody {
	return &NullableCreateSuppressListRequestBody{value: val, isSet: true}
}

func (v NullableCreateSuppressListRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSuppressListRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


