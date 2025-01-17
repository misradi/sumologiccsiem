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

// checks if the UpdateThreatIntelIndicatorRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateThreatIntelIndicatorRequestBody{}

// UpdateThreatIntelIndicatorRequestBody struct for UpdateThreatIntelIndicatorRequestBody
type UpdateThreatIntelIndicatorRequestBody struct {
	Fields *UpdateSuppressListItemRequestBodyFields `json:"fields,omitempty"`
}

// NewUpdateThreatIntelIndicatorRequestBody instantiates a new UpdateThreatIntelIndicatorRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateThreatIntelIndicatorRequestBody() *UpdateThreatIntelIndicatorRequestBody {
	this := UpdateThreatIntelIndicatorRequestBody{}
	return &this
}

// NewUpdateThreatIntelIndicatorRequestBodyWithDefaults instantiates a new UpdateThreatIntelIndicatorRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateThreatIntelIndicatorRequestBodyWithDefaults() *UpdateThreatIntelIndicatorRequestBody {
	this := UpdateThreatIntelIndicatorRequestBody{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *UpdateThreatIntelIndicatorRequestBody) GetFields() UpdateSuppressListItemRequestBodyFields {
	if o == nil || IsNil(o.Fields) {
		var ret UpdateSuppressListItemRequestBodyFields
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreatIntelIndicatorRequestBody) GetFieldsOk() (*UpdateSuppressListItemRequestBodyFields, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *UpdateThreatIntelIndicatorRequestBody) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given UpdateSuppressListItemRequestBodyFields and assigns it to the Fields field.
func (o *UpdateThreatIntelIndicatorRequestBody) SetFields(v UpdateSuppressListItemRequestBodyFields) {
	o.Fields = &v
}

func (o UpdateThreatIntelIndicatorRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateThreatIntelIndicatorRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	return toSerialize, nil
}

type NullableUpdateThreatIntelIndicatorRequestBody struct {
	value *UpdateThreatIntelIndicatorRequestBody
	isSet bool
}

func (v NullableUpdateThreatIntelIndicatorRequestBody) Get() *UpdateThreatIntelIndicatorRequestBody {
	return v.value
}

func (v *NullableUpdateThreatIntelIndicatorRequestBody) Set(val *UpdateThreatIntelIndicatorRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateThreatIntelIndicatorRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateThreatIntelIndicatorRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateThreatIntelIndicatorRequestBody(val *UpdateThreatIntelIndicatorRequestBody) *NullableUpdateThreatIntelIndicatorRequestBody {
	return &NullableUpdateThreatIntelIndicatorRequestBody{value: val, isSet: true}
}

func (v NullableUpdateThreatIntelIndicatorRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateThreatIntelIndicatorRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


