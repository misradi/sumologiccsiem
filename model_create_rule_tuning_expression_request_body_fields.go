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

// checks if the CreateRuleTuningExpressionRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRuleTuningExpressionRequestBodyFields{}

// CreateRuleTuningExpressionRequestBodyFields struct for CreateRuleTuningExpressionRequestBodyFields
type CreateRuleTuningExpressionRequestBodyFields struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Expression string `json:"expression"`
	Enabled bool `json:"enabled"`
	IsGlobal bool `json:"isGlobal"`
	Exclude bool `json:"exclude"`
	RuleIds []string `json:"ruleIds"`
}

type _CreateRuleTuningExpressionRequestBodyFields CreateRuleTuningExpressionRequestBodyFields

// NewCreateRuleTuningExpressionRequestBodyFields instantiates a new CreateRuleTuningExpressionRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRuleTuningExpressionRequestBodyFields(name string, description string, expression string, enabled bool, isGlobal bool, exclude bool, ruleIds []string) *CreateRuleTuningExpressionRequestBodyFields {
	this := CreateRuleTuningExpressionRequestBodyFields{}
	this.Name = name
	this.Description = description
	this.Expression = expression
	this.Enabled = enabled
	this.IsGlobal = isGlobal
	this.Exclude = exclude
	this.RuleIds = ruleIds
	return &this
}

// NewCreateRuleTuningExpressionRequestBodyFieldsWithDefaults instantiates a new CreateRuleTuningExpressionRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRuleTuningExpressionRequestBodyFieldsWithDefaults() *CreateRuleTuningExpressionRequestBodyFields {
	this := CreateRuleTuningExpressionRequestBodyFields{}
	return &this
}

// GetName returns the Name field value
func (o *CreateRuleTuningExpressionRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRuleTuningExpressionRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRuleTuningExpressionRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *CreateRuleTuningExpressionRequestBodyFields) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateRuleTuningExpressionRequestBodyFields) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateRuleTuningExpressionRequestBodyFields) SetDescription(v string) {
	o.Description = v
}

// GetExpression returns the Expression field value
func (o *CreateRuleTuningExpressionRequestBodyFields) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *CreateRuleTuningExpressionRequestBodyFields) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *CreateRuleTuningExpressionRequestBodyFields) SetExpression(v string) {
	o.Expression = v
}

// GetEnabled returns the Enabled field value
func (o *CreateRuleTuningExpressionRequestBodyFields) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateRuleTuningExpressionRequestBodyFields) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateRuleTuningExpressionRequestBodyFields) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIsGlobal returns the IsGlobal field value
func (o *CreateRuleTuningExpressionRequestBodyFields) GetIsGlobal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsGlobal
}

// GetIsGlobalOk returns a tuple with the IsGlobal field value
// and a boolean to check if the value has been set.
func (o *CreateRuleTuningExpressionRequestBodyFields) GetIsGlobalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsGlobal, true
}

// SetIsGlobal sets field value
func (o *CreateRuleTuningExpressionRequestBodyFields) SetIsGlobal(v bool) {
	o.IsGlobal = v
}

// GetExclude returns the Exclude field value
func (o *CreateRuleTuningExpressionRequestBodyFields) GetExclude() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value
// and a boolean to check if the value has been set.
func (o *CreateRuleTuningExpressionRequestBodyFields) GetExcludeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exclude, true
}

// SetExclude sets field value
func (o *CreateRuleTuningExpressionRequestBodyFields) SetExclude(v bool) {
	o.Exclude = v
}

// GetRuleIds returns the RuleIds field value
func (o *CreateRuleTuningExpressionRequestBodyFields) GetRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RuleIds
}

// GetRuleIdsOk returns a tuple with the RuleIds field value
// and a boolean to check if the value has been set.
func (o *CreateRuleTuningExpressionRequestBodyFields) GetRuleIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleIds, true
}

// SetRuleIds sets field value
func (o *CreateRuleTuningExpressionRequestBodyFields) SetRuleIds(v []string) {
	o.RuleIds = v
}

func (o CreateRuleTuningExpressionRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRuleTuningExpressionRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["expression"] = o.Expression
	toSerialize["enabled"] = o.Enabled
	toSerialize["isGlobal"] = o.IsGlobal
	toSerialize["exclude"] = o.Exclude
	toSerialize["ruleIds"] = o.RuleIds
	return toSerialize, nil
}

func (o *CreateRuleTuningExpressionRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"expression",
		"enabled",
		"isGlobal",
		"exclude",
		"ruleIds",
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

	varCreateRuleTuningExpressionRequestBodyFields := _CreateRuleTuningExpressionRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRuleTuningExpressionRequestBodyFields)

	if err != nil {
		return err
	}

	*o = CreateRuleTuningExpressionRequestBodyFields(varCreateRuleTuningExpressionRequestBodyFields)

	return err
}

type NullableCreateRuleTuningExpressionRequestBodyFields struct {
	value *CreateRuleTuningExpressionRequestBodyFields
	isSet bool
}

func (v NullableCreateRuleTuningExpressionRequestBodyFields) Get() *CreateRuleTuningExpressionRequestBodyFields {
	return v.value
}

func (v *NullableCreateRuleTuningExpressionRequestBodyFields) Set(val *CreateRuleTuningExpressionRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRuleTuningExpressionRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRuleTuningExpressionRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRuleTuningExpressionRequestBodyFields(val *CreateRuleTuningExpressionRequestBodyFields) *NullableCreateRuleTuningExpressionRequestBodyFields {
	return &NullableCreateRuleTuningExpressionRequestBodyFields{value: val, isSet: true}
}

func (v NullableCreateRuleTuningExpressionRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRuleTuningExpressionRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

