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

// checks if the OverrideMatchRuleRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverrideMatchRuleRequestBodyFields{}

// OverrideMatchRuleRequestBodyFields struct for OverrideMatchRuleRequestBodyFields
type OverrideMatchRuleRequestBodyFields struct {
	TuningExpressionIds []string `json:"tuningExpressionIds"`
	IsPrototype *bool `json:"isPrototype,omitempty"`
	EntitySelectors []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner `json:"entitySelectors,omitempty"`
	Name *string `json:"name,omitempty"`
	RuleDescription *string `json:"ruleDescription,omitempty"`
	SummaryExpression *string `json:"summaryExpression,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

type _OverrideMatchRuleRequestBodyFields OverrideMatchRuleRequestBodyFields

// NewOverrideMatchRuleRequestBodyFields instantiates a new OverrideMatchRuleRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideMatchRuleRequestBodyFields(tuningExpressionIds []string) *OverrideMatchRuleRequestBodyFields {
	this := OverrideMatchRuleRequestBodyFields{}
	this.TuningExpressionIds = tuningExpressionIds
	return &this
}

// NewOverrideMatchRuleRequestBodyFieldsWithDefaults instantiates a new OverrideMatchRuleRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideMatchRuleRequestBodyFieldsWithDefaults() *OverrideMatchRuleRequestBodyFields {
	this := OverrideMatchRuleRequestBodyFields{}
	return &this
}

// GetTuningExpressionIds returns the TuningExpressionIds field value
func (o *OverrideMatchRuleRequestBodyFields) GetTuningExpressionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TuningExpressionIds
}

// GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field value
// and a boolean to check if the value has been set.
func (o *OverrideMatchRuleRequestBodyFields) GetTuningExpressionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TuningExpressionIds, true
}

// SetTuningExpressionIds sets field value
func (o *OverrideMatchRuleRequestBodyFields) SetTuningExpressionIds(v []string) {
	o.TuningExpressionIds = v
}

// GetIsPrototype returns the IsPrototype field value if set, zero value otherwise.
func (o *OverrideMatchRuleRequestBodyFields) GetIsPrototype() bool {
	if o == nil || IsNil(o.IsPrototype) {
		var ret bool
		return ret
	}
	return *o.IsPrototype
}

// GetIsPrototypeOk returns a tuple with the IsPrototype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideMatchRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrototype) {
		return nil, false
	}
	return o.IsPrototype, true
}

// HasIsPrototype returns a boolean if a field has been set.
func (o *OverrideMatchRuleRequestBodyFields) HasIsPrototype() bool {
	if o != nil && !IsNil(o.IsPrototype) {
		return true
	}

	return false
}

// SetIsPrototype gets a reference to the given bool and assigns it to the IsPrototype field.
func (o *OverrideMatchRuleRequestBodyFields) SetIsPrototype(v bool) {
	o.IsPrototype = &v
}

// GetEntitySelectors returns the EntitySelectors field value if set, zero value otherwise.
func (o *OverrideMatchRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner {
	if o == nil || IsNil(o.EntitySelectors) {
		var ret []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner
		return ret
	}
	return o.EntitySelectors
}

// GetEntitySelectorsOk returns a tuple with the EntitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideMatchRuleRequestBodyFields) GetEntitySelectorsOk() ([]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool) {
	if o == nil || IsNil(o.EntitySelectors) {
		return nil, false
	}
	return o.EntitySelectors, true
}

// HasEntitySelectors returns a boolean if a field has been set.
func (o *OverrideMatchRuleRequestBodyFields) HasEntitySelectors() bool {
	if o != nil && !IsNil(o.EntitySelectors) {
		return true
	}

	return false
}

// SetEntitySelectors gets a reference to the given []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner and assigns it to the EntitySelectors field.
func (o *OverrideMatchRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner) {
	o.EntitySelectors = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OverrideMatchRuleRequestBodyFields) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideMatchRuleRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OverrideMatchRuleRequestBodyFields) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OverrideMatchRuleRequestBodyFields) SetName(v string) {
	o.Name = &v
}

// GetRuleDescription returns the RuleDescription field value if set, zero value otherwise.
func (o *OverrideMatchRuleRequestBodyFields) GetRuleDescription() string {
	if o == nil || IsNil(o.RuleDescription) {
		var ret string
		return ret
	}
	return *o.RuleDescription
}

// GetRuleDescriptionOk returns a tuple with the RuleDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideMatchRuleRequestBodyFields) GetRuleDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.RuleDescription) {
		return nil, false
	}
	return o.RuleDescription, true
}

// HasRuleDescription returns a boolean if a field has been set.
func (o *OverrideMatchRuleRequestBodyFields) HasRuleDescription() bool {
	if o != nil && !IsNil(o.RuleDescription) {
		return true
	}

	return false
}

// SetRuleDescription gets a reference to the given string and assigns it to the RuleDescription field.
func (o *OverrideMatchRuleRequestBodyFields) SetRuleDescription(v string) {
	o.RuleDescription = &v
}

// GetSummaryExpression returns the SummaryExpression field value if set, zero value otherwise.
func (o *OverrideMatchRuleRequestBodyFields) GetSummaryExpression() string {
	if o == nil || IsNil(o.SummaryExpression) {
		var ret string
		return ret
	}
	return *o.SummaryExpression
}

// GetSummaryExpressionOk returns a tuple with the SummaryExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideMatchRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.SummaryExpression) {
		return nil, false
	}
	return o.SummaryExpression, true
}

// HasSummaryExpression returns a boolean if a field has been set.
func (o *OverrideMatchRuleRequestBodyFields) HasSummaryExpression() bool {
	if o != nil && !IsNil(o.SummaryExpression) {
		return true
	}

	return false
}

// SetSummaryExpression gets a reference to the given string and assigns it to the SummaryExpression field.
func (o *OverrideMatchRuleRequestBodyFields) SetSummaryExpression(v string) {
	o.SummaryExpression = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *OverrideMatchRuleRequestBodyFields) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideMatchRuleRequestBodyFields) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *OverrideMatchRuleRequestBodyFields) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *OverrideMatchRuleRequestBodyFields) SetTags(v []string) {
	o.Tags = v
}

func (o OverrideMatchRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideMatchRuleRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tuningExpressionIds"] = o.TuningExpressionIds
	if !IsNil(o.IsPrototype) {
		toSerialize["isPrototype"] = o.IsPrototype
	}
	if !IsNil(o.EntitySelectors) {
		toSerialize["entitySelectors"] = o.EntitySelectors
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RuleDescription) {
		toSerialize["ruleDescription"] = o.RuleDescription
	}
	if !IsNil(o.SummaryExpression) {
		toSerialize["summaryExpression"] = o.SummaryExpression
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

func (o *OverrideMatchRuleRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tuningExpressionIds",
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

	varOverrideMatchRuleRequestBodyFields := _OverrideMatchRuleRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOverrideMatchRuleRequestBodyFields)

	if err != nil {
		return err
	}

	*o = OverrideMatchRuleRequestBodyFields(varOverrideMatchRuleRequestBodyFields)

	return err
}

type NullableOverrideMatchRuleRequestBodyFields struct {
	value *OverrideMatchRuleRequestBodyFields
	isSet bool
}

func (v NullableOverrideMatchRuleRequestBodyFields) Get() *OverrideMatchRuleRequestBodyFields {
	return v.value
}

func (v *NullableOverrideMatchRuleRequestBodyFields) Set(val *OverrideMatchRuleRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideMatchRuleRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideMatchRuleRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideMatchRuleRequestBodyFields(val *OverrideMatchRuleRequestBodyFields) *NullableOverrideMatchRuleRequestBodyFields {
	return &NullableOverrideMatchRuleRequestBodyFields{value: val, isSet: true}
}

func (v NullableOverrideMatchRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideMatchRuleRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


