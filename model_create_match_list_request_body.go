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

// checks if the CreateMatchListRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMatchListRequestBody{}

// CreateMatchListRequestBody struct for CreateMatchListRequestBody
type CreateMatchListRequestBody struct {
	Fields CreateSuppressListRequestBodyFields `json:"fields"`
}

type _CreateMatchListRequestBody CreateMatchListRequestBody

// NewCreateMatchListRequestBody instantiates a new CreateMatchListRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMatchListRequestBody(fields CreateSuppressListRequestBodyFields) *CreateMatchListRequestBody {
	this := CreateMatchListRequestBody{}
	this.Fields = fields
	return &this
}

// NewCreateMatchListRequestBodyWithDefaults instantiates a new CreateMatchListRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMatchListRequestBodyWithDefaults() *CreateMatchListRequestBody {
	this := CreateMatchListRequestBody{}
	return &this
}

// GetFields returns the Fields field value
func (o *CreateMatchListRequestBody) GetFields() CreateSuppressListRequestBodyFields {
	if o == nil {
		var ret CreateSuppressListRequestBodyFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *CreateMatchListRequestBody) GetFieldsOk() (*CreateSuppressListRequestBodyFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *CreateMatchListRequestBody) SetFields(v CreateSuppressListRequestBodyFields) {
	o.Fields = v
}

func (o CreateMatchListRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMatchListRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *CreateMatchListRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varCreateMatchListRequestBody := _CreateMatchListRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateMatchListRequestBody)

	if err != nil {
		return err
	}

	*o = CreateMatchListRequestBody(varCreateMatchListRequestBody)

	return err
}

type NullableCreateMatchListRequestBody struct {
	value *CreateMatchListRequestBody
	isSet bool
}

func (v NullableCreateMatchListRequestBody) Get() *CreateMatchListRequestBody {
	return v.value
}

func (v *NullableCreateMatchListRequestBody) Set(val *CreateMatchListRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMatchListRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMatchListRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMatchListRequestBody(val *CreateMatchListRequestBody) *NullableCreateMatchListRequestBody {
	return &NullableCreateMatchListRequestBody{value: val, isSet: true}
}

func (v NullableCreateMatchListRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMatchListRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


