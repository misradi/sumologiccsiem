/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UpdateInsightAssigneeRequestBodyAssignee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInsightAssigneeRequestBodyAssignee{}

// UpdateInsightAssigneeRequestBodyAssignee The assignee to assign this Insight to.
type UpdateInsightAssigneeRequestBodyAssignee struct {
	// The type of the assignee, either USER or TEAM.
	Type string `json:"type"`
	// The username or team name of the user/team to be assigned.
	Value string `json:"value"`
}

type _UpdateInsightAssigneeRequestBodyAssignee UpdateInsightAssigneeRequestBodyAssignee

// NewUpdateInsightAssigneeRequestBodyAssignee instantiates a new UpdateInsightAssigneeRequestBodyAssignee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInsightAssigneeRequestBodyAssignee(type_ string, value string) *UpdateInsightAssigneeRequestBodyAssignee {
	this := UpdateInsightAssigneeRequestBodyAssignee{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewUpdateInsightAssigneeRequestBodyAssigneeWithDefaults instantiates a new UpdateInsightAssigneeRequestBodyAssignee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInsightAssigneeRequestBodyAssigneeWithDefaults() *UpdateInsightAssigneeRequestBodyAssignee {
	this := UpdateInsightAssigneeRequestBodyAssignee{}
	return &this
}

// GetType returns the Type field value
func (o *UpdateInsightAssigneeRequestBodyAssignee) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateInsightAssigneeRequestBodyAssignee) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateInsightAssigneeRequestBodyAssignee) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *UpdateInsightAssigneeRequestBodyAssignee) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UpdateInsightAssigneeRequestBodyAssignee) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UpdateInsightAssigneeRequestBodyAssignee) SetValue(v string) {
	o.Value = v
}

func (o UpdateInsightAssigneeRequestBodyAssignee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateInsightAssigneeRequestBodyAssignee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *UpdateInsightAssigneeRequestBodyAssignee) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varUpdateInsightAssigneeRequestBodyAssignee := _UpdateInsightAssigneeRequestBodyAssignee{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateInsightAssigneeRequestBodyAssignee)

	if err != nil {
		return err
	}

	*o = UpdateInsightAssigneeRequestBodyAssignee(varUpdateInsightAssigneeRequestBodyAssignee)

	return err
}

type NullableUpdateInsightAssigneeRequestBodyAssignee struct {
	value *UpdateInsightAssigneeRequestBodyAssignee
	isSet bool
}

func (v NullableUpdateInsightAssigneeRequestBodyAssignee) Get() *UpdateInsightAssigneeRequestBodyAssignee {
	return v.value
}

func (v *NullableUpdateInsightAssigneeRequestBodyAssignee) Set(val *UpdateInsightAssigneeRequestBodyAssignee) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInsightAssigneeRequestBodyAssignee) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInsightAssigneeRequestBodyAssignee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInsightAssigneeRequestBodyAssignee(val *UpdateInsightAssigneeRequestBodyAssignee) *NullableUpdateInsightAssigneeRequestBodyAssignee {
	return &NullableUpdateInsightAssigneeRequestBodyAssignee{value: val, isSet: true}
}

func (v NullableUpdateInsightAssigneeRequestBodyAssignee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateInsightAssigneeRequestBodyAssignee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

