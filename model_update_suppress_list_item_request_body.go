/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologiccsiem

import (
	"encoding/json"
)

// checks if the UpdateSuppressListItemRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSuppressListItemRequestBody{}

// UpdateSuppressListItemRequestBody struct for UpdateSuppressListItemRequestBody
type UpdateSuppressListItemRequestBody struct {
	Fields *UpdateSuppressListItemRequestBodyFields `json:"fields,omitempty"`
}

// NewUpdateSuppressListItemRequestBody instantiates a new UpdateSuppressListItemRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSuppressListItemRequestBody() *UpdateSuppressListItemRequestBody {
	this := UpdateSuppressListItemRequestBody{}
	return &this
}

// NewUpdateSuppressListItemRequestBodyWithDefaults instantiates a new UpdateSuppressListItemRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSuppressListItemRequestBodyWithDefaults() *UpdateSuppressListItemRequestBody {
	this := UpdateSuppressListItemRequestBody{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *UpdateSuppressListItemRequestBody) GetFields() UpdateSuppressListItemRequestBodyFields {
	if o == nil || IsNil(o.Fields) {
		var ret UpdateSuppressListItemRequestBodyFields
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSuppressListItemRequestBody) GetFieldsOk() (*UpdateSuppressListItemRequestBodyFields, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *UpdateSuppressListItemRequestBody) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given UpdateSuppressListItemRequestBodyFields and assigns it to the Fields field.
func (o *UpdateSuppressListItemRequestBody) SetFields(v UpdateSuppressListItemRequestBodyFields) {
	o.Fields = &v
}

func (o UpdateSuppressListItemRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSuppressListItemRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	return toSerialize, nil
}

type NullableUpdateSuppressListItemRequestBody struct {
	value *UpdateSuppressListItemRequestBody
	isSet bool
}

func (v NullableUpdateSuppressListItemRequestBody) Get() *UpdateSuppressListItemRequestBody {
	return v.value
}

func (v *NullableUpdateSuppressListItemRequestBody) Set(val *UpdateSuppressListItemRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSuppressListItemRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSuppressListItemRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSuppressListItemRequestBody(val *UpdateSuppressListItemRequestBody) *NullableUpdateSuppressListItemRequestBody {
	return &NullableUpdateSuppressListItemRequestBody{value: val, isSet: true}
}

func (v NullableUpdateSuppressListItemRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSuppressListItemRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


