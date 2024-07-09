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

// checks if the GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner{}

// GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner struct for GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner
type GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Expression string `json:"expression"`
	Enabled bool `json:"enabled"`
	IsGlobal bool `json:"isGlobal"`
}

type _GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner

// NewGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner instantiates a new GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner(id string, name string, description string, expression string, enabled bool, isGlobal bool) *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner {
	this := GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Expression = expression
	this.Enabled = enabled
	this.IsGlobal = isGlobal
	return &this
}

// NewGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInnerWithDefaults instantiates a new GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInnerWithDefaults() *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner {
	this := GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner{}
	return &this
}

// GetId returns the Id field value
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) SetDescription(v string) {
	o.Description = v
}

// GetExpression returns the Expression field value
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) SetExpression(v string) {
	o.Expression = v
}

// GetEnabled returns the Enabled field value
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIsGlobal returns the IsGlobal field value
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) GetIsGlobal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsGlobal
}

// GetIsGlobalOk returns a tuple with the IsGlobal field value
// and a boolean to check if the value has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) GetIsGlobalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsGlobal, true
}

// SetIsGlobal sets field value
func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) SetIsGlobal(v bool) {
	o.IsGlobal = v
}

func (o GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["expression"] = o.Expression
	toSerialize["enabled"] = o.Enabled
	toSerialize["isGlobal"] = o.IsGlobal
	return toSerialize, nil
}

func (o *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
		"expression",
		"enabled",
		"isGlobal",
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

	varGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner := _GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner)

	if err != nil {
		return err
	}

	*o = GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner(varGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner)

	return err
}

type NullableGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner struct {
	value *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner
	isSet bool
}

func (v NullableGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) Get() *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner {
	return v.value
}

func (v *NullableGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) Set(val *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner(val *GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) *NullableGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner {
	return &NullableGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner{value: val, isSet: true}
}

func (v NullableGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


