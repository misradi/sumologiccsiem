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

// checks if the OverrideChainRuleRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverrideChainRuleRequestBodyFields{}

// OverrideChainRuleRequestBodyFields struct for OverrideChainRuleRequestBodyFields
type OverrideChainRuleRequestBodyFields struct {
	TuningExpressionIds []string `json:"tuningExpressionIds"`
	IsPrototype *bool `json:"isPrototype,omitempty"`
	EntitySelectors []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner `json:"entitySelectors,omitempty"`
	Name *string `json:"name,omitempty"`
	RuleDescription *string `json:"ruleDescription,omitempty"`
	SummaryExpression *string `json:"summaryExpression,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Description *string `json:"description,omitempty"`
	GroupByFields []string `json:"groupByFields,omitempty"`
	Score *int32 `json:"score,omitempty"`
	WindowSize *string `json:"windowSize,omitempty"`
	// Can be used instead of windowSize.
	WindowSizeMilliseconds *string `json:"windowSizeMilliseconds,omitempty"`
	Limits []OverrideChainRuleRequestBodyFieldsLimitsInner `json:"limits,omitempty"`
}

type _OverrideChainRuleRequestBodyFields OverrideChainRuleRequestBodyFields

// NewOverrideChainRuleRequestBodyFields instantiates a new OverrideChainRuleRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideChainRuleRequestBodyFields(tuningExpressionIds []string) *OverrideChainRuleRequestBodyFields {
	this := OverrideChainRuleRequestBodyFields{}
	this.TuningExpressionIds = tuningExpressionIds
	return &this
}

// NewOverrideChainRuleRequestBodyFieldsWithDefaults instantiates a new OverrideChainRuleRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideChainRuleRequestBodyFieldsWithDefaults() *OverrideChainRuleRequestBodyFields {
	this := OverrideChainRuleRequestBodyFields{}
	return &this
}

// GetTuningExpressionIds returns the TuningExpressionIds field value
func (o *OverrideChainRuleRequestBodyFields) GetTuningExpressionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TuningExpressionIds
}

// GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field value
// and a boolean to check if the value has been set.
func (o *OverrideChainRuleRequestBodyFields) GetTuningExpressionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TuningExpressionIds, true
}

// SetTuningExpressionIds sets field value
func (o *OverrideChainRuleRequestBodyFields) SetTuningExpressionIds(v []string) {
	o.TuningExpressionIds = v
}

// GetIsPrototype returns the IsPrototype field value if set, zero value otherwise.
func (o *OverrideChainRuleRequestBodyFields) GetIsPrototype() bool {
	if o == nil || IsNil(o.IsPrototype) {
		var ret bool
		return ret
	}
	return *o.IsPrototype
}

// GetIsPrototypeOk returns a tuple with the IsPrototype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideChainRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrototype) {
		return nil, false
	}
	return o.IsPrototype, true
}

// HasIsPrototype returns a boolean if a field has been set.
func (o *OverrideChainRuleRequestBodyFields) HasIsPrototype() bool {
	if o != nil && !IsNil(o.IsPrototype) {
		return true
	}

	return false
}

// SetIsPrototype gets a reference to the given bool and assigns it to the IsPrototype field.
func (o *OverrideChainRuleRequestBodyFields) SetIsPrototype(v bool) {
	o.IsPrototype = &v
}

// GetEntitySelectors returns the EntitySelectors field value if set, zero value otherwise.
func (o *OverrideChainRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner {
	if o == nil || IsNil(o.EntitySelectors) {
		var ret []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner
		return ret
	}
	return o.EntitySelectors
}

// GetEntitySelectorsOk returns a tuple with the EntitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideChainRuleRequestBodyFields) GetEntitySelectorsOk() ([]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool) {
	if o == nil || IsNil(o.EntitySelectors) {
		return nil, false
	}
	return o.EntitySelectors, true
}

// HasEntitySelectors returns a boolean if a field has been set.
func (o *OverrideChainRuleRequestBodyFields) HasEntitySelectors() bool {
	if o != nil && !IsNil(o.EntitySelectors) {
		return true
	}

	return false
}

// SetEntitySelectors gets a reference to the given []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner and assigns it to the EntitySelectors field.
func (o *OverrideChainRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner) {
	o.EntitySelectors = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OverrideChainRuleRequestBodyFields) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideChainRuleRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OverrideChainRuleRequestBodyFields) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OverrideChainRuleRequestBodyFields) SetName(v string) {
	o.Name = &v
}

// GetRuleDescription returns the RuleDescription field value if set, zero value otherwise.
func (o *OverrideChainRuleRequestBodyFields) GetRuleDescription() string {
	if o == nil || IsNil(o.RuleDescription) {
		var ret string
		return ret
	}
	return *o.RuleDescription
}

// GetRuleDescriptionOk returns a tuple with the RuleDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideChainRuleRequestBodyFields) GetRuleDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.RuleDescription) {
		return nil, false
	}
	return o.RuleDescription, true
}

// HasRuleDescription returns a boolean if a field has been set.
func (o *OverrideChainRuleRequestBodyFields) HasRuleDescription() bool {
	if o != nil && !IsNil(o.RuleDescription) {
		return true
	}

	return false
}

// SetRuleDescription gets a reference to the given string and assigns it to the RuleDescription field.
func (o *OverrideChainRuleRequestBodyFields) SetRuleDescription(v string) {
	o.RuleDescription = &v
}

// GetSummaryExpression returns the SummaryExpression field value if set, zero value otherwise.
func (o *OverrideChainRuleRequestBodyFields) GetSummaryExpression() string {
	if o == nil || IsNil(o.SummaryExpression) {
		var ret string
		return ret
	}
	return *o.SummaryExpression
}

// GetSummaryExpressionOk returns a tuple with the SummaryExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideChainRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.SummaryExpression) {
		return nil, false
	}
	return o.SummaryExpression, true
}

// HasSummaryExpression returns a boolean if a field has been set.
func (o *OverrideChainRuleRequestBodyFields) HasSummaryExpression() bool {
	if o != nil && !IsNil(o.SummaryExpression) {
		return true
	}

	return false
}

// SetSummaryExpression gets a reference to the given string and assigns it to the SummaryExpression field.
func (o *OverrideChainRuleRequestBodyFields) SetSummaryExpression(v string) {
	o.SummaryExpression = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *OverrideChainRuleRequestBodyFields) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideChainRuleRequestBodyFields) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *OverrideChainRuleRequestBodyFields) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *OverrideChainRuleRequestBodyFields) SetTags(v []string) {
	o.Tags = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OverrideChainRuleRequestBodyFields) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideChainRuleRequestBodyFields) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OverrideChainRuleRequestBodyFields) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OverrideChainRuleRequestBodyFields) SetDescription(v string) {
	o.Description = &v
}

// GetGroupByFields returns the GroupByFields field value if set, zero value otherwise.
func (o *OverrideChainRuleRequestBodyFields) GetGroupByFields() []string {
	if o == nil || IsNil(o.GroupByFields) {
		var ret []string
		return ret
	}
	return o.GroupByFields
}

// GetGroupByFieldsOk returns a tuple with the GroupByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideChainRuleRequestBodyFields) GetGroupByFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupByFields) {
		return nil, false
	}
	return o.GroupByFields, true
}

// HasGroupByFields returns a boolean if a field has been set.
func (o *OverrideChainRuleRequestBodyFields) HasGroupByFields() bool {
	if o != nil && !IsNil(o.GroupByFields) {
		return true
	}

	return false
}

// SetGroupByFields gets a reference to the given []string and assigns it to the GroupByFields field.
func (o *OverrideChainRuleRequestBodyFields) SetGroupByFields(v []string) {
	o.GroupByFields = v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *OverrideChainRuleRequestBodyFields) GetScore() int32 {
	if o == nil || IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideChainRuleRequestBodyFields) GetScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *OverrideChainRuleRequestBodyFields) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *OverrideChainRuleRequestBodyFields) SetScore(v int32) {
	o.Score = &v
}

// GetWindowSize returns the WindowSize field value if set, zero value otherwise.
func (o *OverrideChainRuleRequestBodyFields) GetWindowSize() string {
	if o == nil || IsNil(o.WindowSize) {
		var ret string
		return ret
	}
	return *o.WindowSize
}

// GetWindowSizeOk returns a tuple with the WindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideChainRuleRequestBodyFields) GetWindowSizeOk() (*string, bool) {
	if o == nil || IsNil(o.WindowSize) {
		return nil, false
	}
	return o.WindowSize, true
}

// HasWindowSize returns a boolean if a field has been set.
func (o *OverrideChainRuleRequestBodyFields) HasWindowSize() bool {
	if o != nil && !IsNil(o.WindowSize) {
		return true
	}

	return false
}

// SetWindowSize gets a reference to the given string and assigns it to the WindowSize field.
func (o *OverrideChainRuleRequestBodyFields) SetWindowSize(v string) {
	o.WindowSize = &v
}

// GetWindowSizeMilliseconds returns the WindowSizeMilliseconds field value if set, zero value otherwise.
func (o *OverrideChainRuleRequestBodyFields) GetWindowSizeMilliseconds() string {
	if o == nil || IsNil(o.WindowSizeMilliseconds) {
		var ret string
		return ret
	}
	return *o.WindowSizeMilliseconds
}

// GetWindowSizeMillisecondsOk returns a tuple with the WindowSizeMilliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideChainRuleRequestBodyFields) GetWindowSizeMillisecondsOk() (*string, bool) {
	if o == nil || IsNil(o.WindowSizeMilliseconds) {
		return nil, false
	}
	return o.WindowSizeMilliseconds, true
}

// HasWindowSizeMilliseconds returns a boolean if a field has been set.
func (o *OverrideChainRuleRequestBodyFields) HasWindowSizeMilliseconds() bool {
	if o != nil && !IsNil(o.WindowSizeMilliseconds) {
		return true
	}

	return false
}

// SetWindowSizeMilliseconds gets a reference to the given string and assigns it to the WindowSizeMilliseconds field.
func (o *OverrideChainRuleRequestBodyFields) SetWindowSizeMilliseconds(v string) {
	o.WindowSizeMilliseconds = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *OverrideChainRuleRequestBodyFields) GetLimits() []OverrideChainRuleRequestBodyFieldsLimitsInner {
	if o == nil || IsNil(o.Limits) {
		var ret []OverrideChainRuleRequestBodyFieldsLimitsInner
		return ret
	}
	return o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideChainRuleRequestBodyFields) GetLimitsOk() ([]OverrideChainRuleRequestBodyFieldsLimitsInner, bool) {
	if o == nil || IsNil(o.Limits) {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *OverrideChainRuleRequestBodyFields) HasLimits() bool {
	if o != nil && !IsNil(o.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given []OverrideChainRuleRequestBodyFieldsLimitsInner and assigns it to the Limits field.
func (o *OverrideChainRuleRequestBodyFields) SetLimits(v []OverrideChainRuleRequestBodyFieldsLimitsInner) {
	o.Limits = v
}

func (o OverrideChainRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideChainRuleRequestBodyFields) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.GroupByFields) {
		toSerialize["groupByFields"] = o.GroupByFields
	}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.WindowSize) {
		toSerialize["windowSize"] = o.WindowSize
	}
	if !IsNil(o.WindowSizeMilliseconds) {
		toSerialize["windowSizeMilliseconds"] = o.WindowSizeMilliseconds
	}
	if !IsNil(o.Limits) {
		toSerialize["limits"] = o.Limits
	}
	return toSerialize, nil
}

func (o *OverrideChainRuleRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
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

	varOverrideChainRuleRequestBodyFields := _OverrideChainRuleRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOverrideChainRuleRequestBodyFields)

	if err != nil {
		return err
	}

	*o = OverrideChainRuleRequestBodyFields(varOverrideChainRuleRequestBodyFields)

	return err
}

type NullableOverrideChainRuleRequestBodyFields struct {
	value *OverrideChainRuleRequestBodyFields
	isSet bool
}

func (v NullableOverrideChainRuleRequestBodyFields) Get() *OverrideChainRuleRequestBodyFields {
	return v.value
}

func (v *NullableOverrideChainRuleRequestBodyFields) Set(val *OverrideChainRuleRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideChainRuleRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideChainRuleRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideChainRuleRequestBodyFields(val *OverrideChainRuleRequestBodyFields) *NullableOverrideChainRuleRequestBodyFields {
	return &NullableOverrideChainRuleRequestBodyFields{value: val, isSet: true}
}

func (v NullableOverrideChainRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideChainRuleRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


