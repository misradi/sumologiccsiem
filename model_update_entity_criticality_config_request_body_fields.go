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

// checks if the UpdateEntityCriticalityConfigRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEntityCriticalityConfigRequestBodyFields{}

// UpdateEntityCriticalityConfigRequestBodyFields struct for UpdateEntityCriticalityConfigRequestBodyFields
type UpdateEntityCriticalityConfigRequestBodyFields struct {
	// Algebraic expression representing this entity's criticality. Examples: \"severity * 2\", \"severity - 5\", \"severity / 3\"
	SeverityExpression string `json:"severityExpression"`
}

type _UpdateEntityCriticalityConfigRequestBodyFields UpdateEntityCriticalityConfigRequestBodyFields

// NewUpdateEntityCriticalityConfigRequestBodyFields instantiates a new UpdateEntityCriticalityConfigRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEntityCriticalityConfigRequestBodyFields(severityExpression string) *UpdateEntityCriticalityConfigRequestBodyFields {
	this := UpdateEntityCriticalityConfigRequestBodyFields{}
	this.SeverityExpression = severityExpression
	return &this
}

// NewUpdateEntityCriticalityConfigRequestBodyFieldsWithDefaults instantiates a new UpdateEntityCriticalityConfigRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEntityCriticalityConfigRequestBodyFieldsWithDefaults() *UpdateEntityCriticalityConfigRequestBodyFields {
	this := UpdateEntityCriticalityConfigRequestBodyFields{}
	return &this
}

// GetSeverityExpression returns the SeverityExpression field value
func (o *UpdateEntityCriticalityConfigRequestBodyFields) GetSeverityExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeverityExpression
}

// GetSeverityExpressionOk returns a tuple with the SeverityExpression field value
// and a boolean to check if the value has been set.
func (o *UpdateEntityCriticalityConfigRequestBodyFields) GetSeverityExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeverityExpression, true
}

// SetSeverityExpression sets field value
func (o *UpdateEntityCriticalityConfigRequestBodyFields) SetSeverityExpression(v string) {
	o.SeverityExpression = v
}

func (o UpdateEntityCriticalityConfigRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEntityCriticalityConfigRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["severityExpression"] = o.SeverityExpression
	return toSerialize, nil
}

func (o *UpdateEntityCriticalityConfigRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"severityExpression",
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

	varUpdateEntityCriticalityConfigRequestBodyFields := _UpdateEntityCriticalityConfigRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateEntityCriticalityConfigRequestBodyFields)

	if err != nil {
		return err
	}

	*o = UpdateEntityCriticalityConfigRequestBodyFields(varUpdateEntityCriticalityConfigRequestBodyFields)

	return err
}

type NullableUpdateEntityCriticalityConfigRequestBodyFields struct {
	value *UpdateEntityCriticalityConfigRequestBodyFields
	isSet bool
}

func (v NullableUpdateEntityCriticalityConfigRequestBodyFields) Get() *UpdateEntityCriticalityConfigRequestBodyFields {
	return v.value
}

func (v *NullableUpdateEntityCriticalityConfigRequestBodyFields) Set(val *UpdateEntityCriticalityConfigRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEntityCriticalityConfigRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEntityCriticalityConfigRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEntityCriticalityConfigRequestBodyFields(val *UpdateEntityCriticalityConfigRequestBodyFields) *NullableUpdateEntityCriticalityConfigRequestBodyFields {
	return &NullableUpdateEntityCriticalityConfigRequestBodyFields{value: val, isSet: true}
}

func (v NullableUpdateEntityCriticalityConfigRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEntityCriticalityConfigRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


