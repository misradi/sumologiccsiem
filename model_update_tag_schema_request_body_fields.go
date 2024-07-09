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

// checks if the UpdateTagSchemaRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTagSchemaRequestBodyFields{}

// UpdateTagSchemaRequestBodyFields struct for UpdateTagSchemaRequestBodyFields
type UpdateTagSchemaRequestBodyFields struct {
	Key string `json:"key"`
	Label string `json:"label"`
	ContentTypes []string `json:"contentTypes,omitempty"`
	Freeform bool `json:"freeform"`
	ValueOptions []UpdateTagSchemaRequestBodyFieldsValueOptionsInner `json:"valueOptions,omitempty"`
}

type _UpdateTagSchemaRequestBodyFields UpdateTagSchemaRequestBodyFields

// NewUpdateTagSchemaRequestBodyFields instantiates a new UpdateTagSchemaRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTagSchemaRequestBodyFields(key string, label string, freeform bool) *UpdateTagSchemaRequestBodyFields {
	this := UpdateTagSchemaRequestBodyFields{}
	this.Key = key
	this.Label = label
	this.Freeform = freeform
	return &this
}

// NewUpdateTagSchemaRequestBodyFieldsWithDefaults instantiates a new UpdateTagSchemaRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTagSchemaRequestBodyFieldsWithDefaults() *UpdateTagSchemaRequestBodyFields {
	this := UpdateTagSchemaRequestBodyFields{}
	return &this
}

// GetKey returns the Key field value
func (o *UpdateTagSchemaRequestBodyFields) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *UpdateTagSchemaRequestBodyFields) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *UpdateTagSchemaRequestBodyFields) SetKey(v string) {
	o.Key = v
}

// GetLabel returns the Label field value
func (o *UpdateTagSchemaRequestBodyFields) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *UpdateTagSchemaRequestBodyFields) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *UpdateTagSchemaRequestBodyFields) SetLabel(v string) {
	o.Label = v
}

// GetContentTypes returns the ContentTypes field value if set, zero value otherwise.
func (o *UpdateTagSchemaRequestBodyFields) GetContentTypes() []string {
	if o == nil || IsNil(o.ContentTypes) {
		var ret []string
		return ret
	}
	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTagSchemaRequestBodyFields) GetContentTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ContentTypes) {
		return nil, false
	}
	return o.ContentTypes, true
}

// HasContentTypes returns a boolean if a field has been set.
func (o *UpdateTagSchemaRequestBodyFields) HasContentTypes() bool {
	if o != nil && !IsNil(o.ContentTypes) {
		return true
	}

	return false
}

// SetContentTypes gets a reference to the given []string and assigns it to the ContentTypes field.
func (o *UpdateTagSchemaRequestBodyFields) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetFreeform returns the Freeform field value
func (o *UpdateTagSchemaRequestBodyFields) GetFreeform() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Freeform
}

// GetFreeformOk returns a tuple with the Freeform field value
// and a boolean to check if the value has been set.
func (o *UpdateTagSchemaRequestBodyFields) GetFreeformOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Freeform, true
}

// SetFreeform sets field value
func (o *UpdateTagSchemaRequestBodyFields) SetFreeform(v bool) {
	o.Freeform = v
}

// GetValueOptions returns the ValueOptions field value if set, zero value otherwise.
func (o *UpdateTagSchemaRequestBodyFields) GetValueOptions() []UpdateTagSchemaRequestBodyFieldsValueOptionsInner {
	if o == nil || IsNil(o.ValueOptions) {
		var ret []UpdateTagSchemaRequestBodyFieldsValueOptionsInner
		return ret
	}
	return o.ValueOptions
}

// GetValueOptionsOk returns a tuple with the ValueOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTagSchemaRequestBodyFields) GetValueOptionsOk() ([]UpdateTagSchemaRequestBodyFieldsValueOptionsInner, bool) {
	if o == nil || IsNil(o.ValueOptions) {
		return nil, false
	}
	return o.ValueOptions, true
}

// HasValueOptions returns a boolean if a field has been set.
func (o *UpdateTagSchemaRequestBodyFields) HasValueOptions() bool {
	if o != nil && !IsNil(o.ValueOptions) {
		return true
	}

	return false
}

// SetValueOptions gets a reference to the given []UpdateTagSchemaRequestBodyFieldsValueOptionsInner and assigns it to the ValueOptions field.
func (o *UpdateTagSchemaRequestBodyFields) SetValueOptions(v []UpdateTagSchemaRequestBodyFieldsValueOptionsInner) {
	o.ValueOptions = v
}

func (o UpdateTagSchemaRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTagSchemaRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["label"] = o.Label
	if !IsNil(o.ContentTypes) {
		toSerialize["contentTypes"] = o.ContentTypes
	}
	toSerialize["freeform"] = o.Freeform
	if !IsNil(o.ValueOptions) {
		toSerialize["valueOptions"] = o.ValueOptions
	}
	return toSerialize, nil
}

func (o *UpdateTagSchemaRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"label",
		"freeform",
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

	varUpdateTagSchemaRequestBodyFields := _UpdateTagSchemaRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateTagSchemaRequestBodyFields)

	if err != nil {
		return err
	}

	*o = UpdateTagSchemaRequestBodyFields(varUpdateTagSchemaRequestBodyFields)

	return err
}

type NullableUpdateTagSchemaRequestBodyFields struct {
	value *UpdateTagSchemaRequestBodyFields
	isSet bool
}

func (v NullableUpdateTagSchemaRequestBodyFields) Get() *UpdateTagSchemaRequestBodyFields {
	return v.value
}

func (v *NullableUpdateTagSchemaRequestBodyFields) Set(val *UpdateTagSchemaRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTagSchemaRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTagSchemaRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTagSchemaRequestBodyFields(val *UpdateTagSchemaRequestBodyFields) *NullableUpdateTagSchemaRequestBodyFields {
	return &NullableUpdateTagSchemaRequestBodyFields{value: val, isSet: true}
}

func (v NullableUpdateTagSchemaRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTagSchemaRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


