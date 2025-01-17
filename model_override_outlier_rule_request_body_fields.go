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

// checks if the OverrideOutlierRuleRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverrideOutlierRuleRequestBodyFields{}

// OverrideOutlierRuleRequestBodyFields struct for OverrideOutlierRuleRequestBodyFields
type OverrideOutlierRuleRequestBodyFields struct {
	TuningExpressionIds []string `json:"tuningExpressionIds"`
	IsPrototype *bool `json:"isPrototype,omitempty"`
	EntitySelectors []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner `json:"entitySelectors,omitempty"`
	Name *string `json:"name,omitempty"`
	RuleDescription *string `json:"ruleDescription,omitempty"`
	SummaryExpression *string `json:"summaryExpression,omitempty"`
	Tags []string `json:"tags,omitempty"`
	NameExpression *string `json:"nameExpression,omitempty"`
	// milliseconds
	FloorValue *string `json:"floorValue,omitempty"`
	GroupByFields []string `json:"groupByFields,omitempty"`
	DescriptionExpression *string `json:"descriptionExpression,omitempty"`
	// milliseconds
	BaselineWindowSize *string `json:"baselineWindowSize,omitempty"`
	// milliseconds
	DeviationThreshold *string `json:"deviationThreshold,omitempty"`
	// milliseconds
	RetentionWindowSize *string `json:"retentionWindowSize,omitempty"`
	Score *int32 `json:"score,omitempty"`
	WindowSize *string `json:"windowSize,omitempty"`
	// Can be used instead of windowSize.
	WindowSizeMilliseconds *string `json:"windowSizeMilliseconds,omitempty"`
}

type _OverrideOutlierRuleRequestBodyFields OverrideOutlierRuleRequestBodyFields

// NewOverrideOutlierRuleRequestBodyFields instantiates a new OverrideOutlierRuleRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideOutlierRuleRequestBodyFields(tuningExpressionIds []string) *OverrideOutlierRuleRequestBodyFields {
	this := OverrideOutlierRuleRequestBodyFields{}
	this.TuningExpressionIds = tuningExpressionIds
	return &this
}

// NewOverrideOutlierRuleRequestBodyFieldsWithDefaults instantiates a new OverrideOutlierRuleRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideOutlierRuleRequestBodyFieldsWithDefaults() *OverrideOutlierRuleRequestBodyFields {
	this := OverrideOutlierRuleRequestBodyFields{}
	return &this
}

// GetTuningExpressionIds returns the TuningExpressionIds field value
func (o *OverrideOutlierRuleRequestBodyFields) GetTuningExpressionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TuningExpressionIds
}

// GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field value
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetTuningExpressionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TuningExpressionIds, true
}

// SetTuningExpressionIds sets field value
func (o *OverrideOutlierRuleRequestBodyFields) SetTuningExpressionIds(v []string) {
	o.TuningExpressionIds = v
}

// GetIsPrototype returns the IsPrototype field value if set, zero value otherwise.
func (o *OverrideOutlierRuleRequestBodyFields) GetIsPrototype() bool {
	if o == nil || IsNil(o.IsPrototype) {
		var ret bool
		return ret
	}
	return *o.IsPrototype
}

// GetIsPrototypeOk returns a tuple with the IsPrototype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrototype) {
		return nil, false
	}
	return o.IsPrototype, true
}

// HasIsPrototype returns a boolean if a field has been set.
func (o *OverrideOutlierRuleRequestBodyFields) HasIsPrototype() bool {
	if o != nil && !IsNil(o.IsPrototype) {
		return true
	}

	return false
}

// SetIsPrototype gets a reference to the given bool and assigns it to the IsPrototype field.
func (o *OverrideOutlierRuleRequestBodyFields) SetIsPrototype(v bool) {
	o.IsPrototype = &v
}

// GetEntitySelectors returns the EntitySelectors field value if set, zero value otherwise.
func (o *OverrideOutlierRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner {
	if o == nil || IsNil(o.EntitySelectors) {
		var ret []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner
		return ret
	}
	return o.EntitySelectors
}

// GetEntitySelectorsOk returns a tuple with the EntitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetEntitySelectorsOk() ([]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool) {
	if o == nil || IsNil(o.EntitySelectors) {
		return nil, false
	}
	return o.EntitySelectors, true
}

// HasEntitySelectors returns a boolean if a field has been set.
func (o *OverrideOutlierRuleRequestBodyFields) HasEntitySelectors() bool {
	if o != nil && !IsNil(o.EntitySelectors) {
		return true
	}

	return false
}

// SetEntitySelectors gets a reference to the given []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner and assigns it to the EntitySelectors field.
func (o *OverrideOutlierRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner) {
	o.EntitySelectors = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OverrideOutlierRuleRequestBodyFields) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OverrideOutlierRuleRequestBodyFields) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OverrideOutlierRuleRequestBodyFields) SetName(v string) {
	o.Name = &v
}

// GetRuleDescription returns the RuleDescription field value if set, zero value otherwise.
func (o *OverrideOutlierRuleRequestBodyFields) GetRuleDescription() string {
	if o == nil || IsNil(o.RuleDescription) {
		var ret string
		return ret
	}
	return *o.RuleDescription
}

// GetRuleDescriptionOk returns a tuple with the RuleDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetRuleDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.RuleDescription) {
		return nil, false
	}
	return o.RuleDescription, true
}

// HasRuleDescription returns a boolean if a field has been set.
func (o *OverrideOutlierRuleRequestBodyFields) HasRuleDescription() bool {
	if o != nil && !IsNil(o.RuleDescription) {
		return true
	}

	return false
}

// SetRuleDescription gets a reference to the given string and assigns it to the RuleDescription field.
func (o *OverrideOutlierRuleRequestBodyFields) SetRuleDescription(v string) {
	o.RuleDescription = &v
}

// GetSummaryExpression returns the SummaryExpression field value if set, zero value otherwise.
func (o *OverrideOutlierRuleRequestBodyFields) GetSummaryExpression() string {
	if o == nil || IsNil(o.SummaryExpression) {
		var ret string
		return ret
	}
	return *o.SummaryExpression
}

// GetSummaryExpressionOk returns a tuple with the SummaryExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.SummaryExpression) {
		return nil, false
	}
	return o.SummaryExpression, true
}

// HasSummaryExpression returns a boolean if a field has been set.
func (o *OverrideOutlierRuleRequestBodyFields) HasSummaryExpression() bool {
	if o != nil && !IsNil(o.SummaryExpression) {
		return true
	}

	return false
}

// SetSummaryExpression gets a reference to the given string and assigns it to the SummaryExpression field.
func (o *OverrideOutlierRuleRequestBodyFields) SetSummaryExpression(v string) {
	o.SummaryExpression = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *OverrideOutlierRuleRequestBodyFields) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *OverrideOutlierRuleRequestBodyFields) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *OverrideOutlierRuleRequestBodyFields) SetTags(v []string) {
	o.Tags = v
}

// GetNameExpression returns the NameExpression field value if set, zero value otherwise.
func (o *OverrideOutlierRuleRequestBodyFields) GetNameExpression() string {
	if o == nil || IsNil(o.NameExpression) {
		var ret string
		return ret
	}
	return *o.NameExpression
}

// GetNameExpressionOk returns a tuple with the NameExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetNameExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.NameExpression) {
		return nil, false
	}
	return o.NameExpression, true
}

// HasNameExpression returns a boolean if a field has been set.
func (o *OverrideOutlierRuleRequestBodyFields) HasNameExpression() bool {
	if o != nil && !IsNil(o.NameExpression) {
		return true
	}

	return false
}

// SetNameExpression gets a reference to the given string and assigns it to the NameExpression field.
func (o *OverrideOutlierRuleRequestBodyFields) SetNameExpression(v string) {
	o.NameExpression = &v
}

// GetFloorValue returns the FloorValue field value if set, zero value otherwise.
func (o *OverrideOutlierRuleRequestBodyFields) GetFloorValue() string {
	if o == nil || IsNil(o.FloorValue) {
		var ret string
		return ret
	}
	return *o.FloorValue
}

// GetFloorValueOk returns a tuple with the FloorValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetFloorValueOk() (*string, bool) {
	if o == nil || IsNil(o.FloorValue) {
		return nil, false
	}
	return o.FloorValue, true
}

// HasFloorValue returns a boolean if a field has been set.
func (o *OverrideOutlierRuleRequestBodyFields) HasFloorValue() bool {
	if o != nil && !IsNil(o.FloorValue) {
		return true
	}

	return false
}

// SetFloorValue gets a reference to the given string and assigns it to the FloorValue field.
func (o *OverrideOutlierRuleRequestBodyFields) SetFloorValue(v string) {
	o.FloorValue = &v
}

// GetGroupByFields returns the GroupByFields field value if set, zero value otherwise.
func (o *OverrideOutlierRuleRequestBodyFields) GetGroupByFields() []string {
	if o == nil || IsNil(o.GroupByFields) {
		var ret []string
		return ret
	}
	return o.GroupByFields
}

// GetGroupByFieldsOk returns a tuple with the GroupByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetGroupByFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupByFields) {
		return nil, false
	}
	return o.GroupByFields, true
}

// HasGroupByFields returns a boolean if a field has been set.
func (o *OverrideOutlierRuleRequestBodyFields) HasGroupByFields() bool {
	if o != nil && !IsNil(o.GroupByFields) {
		return true
	}

	return false
}

// SetGroupByFields gets a reference to the given []string and assigns it to the GroupByFields field.
func (o *OverrideOutlierRuleRequestBodyFields) SetGroupByFields(v []string) {
	o.GroupByFields = v
}

// GetDescriptionExpression returns the DescriptionExpression field value if set, zero value otherwise.
func (o *OverrideOutlierRuleRequestBodyFields) GetDescriptionExpression() string {
	if o == nil || IsNil(o.DescriptionExpression) {
		var ret string
		return ret
	}
	return *o.DescriptionExpression
}

// GetDescriptionExpressionOk returns a tuple with the DescriptionExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetDescriptionExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.DescriptionExpression) {
		return nil, false
	}
	return o.DescriptionExpression, true
}

// HasDescriptionExpression returns a boolean if a field has been set.
func (o *OverrideOutlierRuleRequestBodyFields) HasDescriptionExpression() bool {
	if o != nil && !IsNil(o.DescriptionExpression) {
		return true
	}

	return false
}

// SetDescriptionExpression gets a reference to the given string and assigns it to the DescriptionExpression field.
func (o *OverrideOutlierRuleRequestBodyFields) SetDescriptionExpression(v string) {
	o.DescriptionExpression = &v
}

// GetBaselineWindowSize returns the BaselineWindowSize field value if set, zero value otherwise.
func (o *OverrideOutlierRuleRequestBodyFields) GetBaselineWindowSize() string {
	if o == nil || IsNil(o.BaselineWindowSize) {
		var ret string
		return ret
	}
	return *o.BaselineWindowSize
}

// GetBaselineWindowSizeOk returns a tuple with the BaselineWindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetBaselineWindowSizeOk() (*string, bool) {
	if o == nil || IsNil(o.BaselineWindowSize) {
		return nil, false
	}
	return o.BaselineWindowSize, true
}

// HasBaselineWindowSize returns a boolean if a field has been set.
func (o *OverrideOutlierRuleRequestBodyFields) HasBaselineWindowSize() bool {
	if o != nil && !IsNil(o.BaselineWindowSize) {
		return true
	}

	return false
}

// SetBaselineWindowSize gets a reference to the given string and assigns it to the BaselineWindowSize field.
func (o *OverrideOutlierRuleRequestBodyFields) SetBaselineWindowSize(v string) {
	o.BaselineWindowSize = &v
}

// GetDeviationThreshold returns the DeviationThreshold field value if set, zero value otherwise.
func (o *OverrideOutlierRuleRequestBodyFields) GetDeviationThreshold() string {
	if o == nil || IsNil(o.DeviationThreshold) {
		var ret string
		return ret
	}
	return *o.DeviationThreshold
}

// GetDeviationThresholdOk returns a tuple with the DeviationThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetDeviationThresholdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviationThreshold) {
		return nil, false
	}
	return o.DeviationThreshold, true
}

// HasDeviationThreshold returns a boolean if a field has been set.
func (o *OverrideOutlierRuleRequestBodyFields) HasDeviationThreshold() bool {
	if o != nil && !IsNil(o.DeviationThreshold) {
		return true
	}

	return false
}

// SetDeviationThreshold gets a reference to the given string and assigns it to the DeviationThreshold field.
func (o *OverrideOutlierRuleRequestBodyFields) SetDeviationThreshold(v string) {
	o.DeviationThreshold = &v
}

// GetRetentionWindowSize returns the RetentionWindowSize field value if set, zero value otherwise.
func (o *OverrideOutlierRuleRequestBodyFields) GetRetentionWindowSize() string {
	if o == nil || IsNil(o.RetentionWindowSize) {
		var ret string
		return ret
	}
	return *o.RetentionWindowSize
}

// GetRetentionWindowSizeOk returns a tuple with the RetentionWindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetRetentionWindowSizeOk() (*string, bool) {
	if o == nil || IsNil(o.RetentionWindowSize) {
		return nil, false
	}
	return o.RetentionWindowSize, true
}

// HasRetentionWindowSize returns a boolean if a field has been set.
func (o *OverrideOutlierRuleRequestBodyFields) HasRetentionWindowSize() bool {
	if o != nil && !IsNil(o.RetentionWindowSize) {
		return true
	}

	return false
}

// SetRetentionWindowSize gets a reference to the given string and assigns it to the RetentionWindowSize field.
func (o *OverrideOutlierRuleRequestBodyFields) SetRetentionWindowSize(v string) {
	o.RetentionWindowSize = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *OverrideOutlierRuleRequestBodyFields) GetScore() int32 {
	if o == nil || IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *OverrideOutlierRuleRequestBodyFields) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *OverrideOutlierRuleRequestBodyFields) SetScore(v int32) {
	o.Score = &v
}

// GetWindowSize returns the WindowSize field value if set, zero value otherwise.
func (o *OverrideOutlierRuleRequestBodyFields) GetWindowSize() string {
	if o == nil || IsNil(o.WindowSize) {
		var ret string
		return ret
	}
	return *o.WindowSize
}

// GetWindowSizeOk returns a tuple with the WindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetWindowSizeOk() (*string, bool) {
	if o == nil || IsNil(o.WindowSize) {
		return nil, false
	}
	return o.WindowSize, true
}

// HasWindowSize returns a boolean if a field has been set.
func (o *OverrideOutlierRuleRequestBodyFields) HasWindowSize() bool {
	if o != nil && !IsNil(o.WindowSize) {
		return true
	}

	return false
}

// SetWindowSize gets a reference to the given string and assigns it to the WindowSize field.
func (o *OverrideOutlierRuleRequestBodyFields) SetWindowSize(v string) {
	o.WindowSize = &v
}

// GetWindowSizeMilliseconds returns the WindowSizeMilliseconds field value if set, zero value otherwise.
func (o *OverrideOutlierRuleRequestBodyFields) GetWindowSizeMilliseconds() string {
	if o == nil || IsNil(o.WindowSizeMilliseconds) {
		var ret string
		return ret
	}
	return *o.WindowSizeMilliseconds
}

// GetWindowSizeMillisecondsOk returns a tuple with the WindowSizeMilliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideOutlierRuleRequestBodyFields) GetWindowSizeMillisecondsOk() (*string, bool) {
	if o == nil || IsNil(o.WindowSizeMilliseconds) {
		return nil, false
	}
	return o.WindowSizeMilliseconds, true
}

// HasWindowSizeMilliseconds returns a boolean if a field has been set.
func (o *OverrideOutlierRuleRequestBodyFields) HasWindowSizeMilliseconds() bool {
	if o != nil && !IsNil(o.WindowSizeMilliseconds) {
		return true
	}

	return false
}

// SetWindowSizeMilliseconds gets a reference to the given string and assigns it to the WindowSizeMilliseconds field.
func (o *OverrideOutlierRuleRequestBodyFields) SetWindowSizeMilliseconds(v string) {
	o.WindowSizeMilliseconds = &v
}

func (o OverrideOutlierRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideOutlierRuleRequestBodyFields) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.NameExpression) {
		toSerialize["nameExpression"] = o.NameExpression
	}
	if !IsNil(o.FloorValue) {
		toSerialize["floorValue"] = o.FloorValue
	}
	if !IsNil(o.GroupByFields) {
		toSerialize["groupByFields"] = o.GroupByFields
	}
	if !IsNil(o.DescriptionExpression) {
		toSerialize["descriptionExpression"] = o.DescriptionExpression
	}
	if !IsNil(o.BaselineWindowSize) {
		toSerialize["baselineWindowSize"] = o.BaselineWindowSize
	}
	if !IsNil(o.DeviationThreshold) {
		toSerialize["deviationThreshold"] = o.DeviationThreshold
	}
	if !IsNil(o.RetentionWindowSize) {
		toSerialize["retentionWindowSize"] = o.RetentionWindowSize
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
	return toSerialize, nil
}

func (o *OverrideOutlierRuleRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
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

	varOverrideOutlierRuleRequestBodyFields := _OverrideOutlierRuleRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOverrideOutlierRuleRequestBodyFields)

	if err != nil {
		return err
	}

	*o = OverrideOutlierRuleRequestBodyFields(varOverrideOutlierRuleRequestBodyFields)

	return err
}

type NullableOverrideOutlierRuleRequestBodyFields struct {
	value *OverrideOutlierRuleRequestBodyFields
	isSet bool
}

func (v NullableOverrideOutlierRuleRequestBodyFields) Get() *OverrideOutlierRuleRequestBodyFields {
	return v.value
}

func (v *NullableOverrideOutlierRuleRequestBodyFields) Set(val *OverrideOutlierRuleRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideOutlierRuleRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideOutlierRuleRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideOutlierRuleRequestBodyFields(val *OverrideOutlierRuleRequestBodyFields) *NullableOverrideOutlierRuleRequestBodyFields {
	return &NullableOverrideOutlierRuleRequestBodyFields{value: val, isSet: true}
}

func (v NullableOverrideOutlierRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideOutlierRuleRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


