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

// checks if the CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner{}

// CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner struct for CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner
type CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type _CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner

// NewCreateLogMappingRequestBodyFieldsFieldsInnerLookupInner instantiates a new CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLogMappingRequestBodyFieldsFieldsInnerLookupInner(key string, value string) *CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner {
	this := CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner{}
	this.Key = key
	this.Value = value
	return &this
}

// NewCreateLogMappingRequestBodyFieldsFieldsInnerLookupInnerWithDefaults instantiates a new CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLogMappingRequestBodyFieldsFieldsInnerLookupInnerWithDefaults() *CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner {
	this := CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner{}
	return &this
}

// GetKey returns the Key field value
func (o *CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) SetValue(v string) {
	o.Value = v
}

func (o CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"value",
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

	varCreateLogMappingRequestBodyFieldsFieldsInnerLookupInner := _CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateLogMappingRequestBodyFieldsFieldsInnerLookupInner)

	if err != nil {
		return err
	}

	*o = CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner(varCreateLogMappingRequestBodyFieldsFieldsInnerLookupInner)

	return err
}

type NullableCreateLogMappingRequestBodyFieldsFieldsInnerLookupInner struct {
	value *CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner
	isSet bool
}

func (v NullableCreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) Get() *CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner {
	return v.value
}

func (v *NullableCreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) Set(val *CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLogMappingRequestBodyFieldsFieldsInnerLookupInner(val *CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) *NullableCreateLogMappingRequestBodyFieldsFieldsInnerLookupInner {
	return &NullableCreateLogMappingRequestBodyFieldsFieldsInnerLookupInner{value: val, isSet: true}
}

func (v NullableCreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


