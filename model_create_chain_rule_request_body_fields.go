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

// checks if the CreateChainRuleRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateChainRuleRequestBodyFields{}

// CreateChainRuleRequestBodyFields struct for CreateChainRuleRequestBodyFields
type CreateChainRuleRequestBodyFields struct {
	AssetField *string `json:"assetField,omitempty"`
	Category *string `json:"category,omitempty"`
	Enabled bool `json:"enabled"`
	EntitySelectors []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner `json:"entitySelectors,omitempty"`
	IsPrototype *bool `json:"isPrototype,omitempty"`
	Name string `json:"name"`
	ParentJaskId *string `json:"parentJaskId,omitempty"`
	SummaryExpression *string `json:"summaryExpression,omitempty"`
	Tags []string `json:"tags,omitempty"`
	TuningExpressionIds []string `json:"tuningExpressionIds,omitempty"`
	// The description to be used on the generated Signal
	Description string `json:"description"`
	ExpressionsAndLimits []CreateChainRuleRequestBodyFieldsExpressionsAndLimitsInner `json:"expressionsAndLimits"`
	GroupByFields []string `json:"groupByFields,omitempty"`
	Ordered *bool `json:"ordered,omitempty"`
	Score int32 `json:"score"`
	Stream string `json:"stream"`
	WindowSize string `json:"windowSize"`
	// Can be used instead of windowSize.
	WindowSizeMilliseconds *string `json:"windowSizeMilliseconds,omitempty"`
}

type _CreateChainRuleRequestBodyFields CreateChainRuleRequestBodyFields

// NewCreateChainRuleRequestBodyFields instantiates a new CreateChainRuleRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateChainRuleRequestBodyFields(enabled bool, name string, description string, expressionsAndLimits []CreateChainRuleRequestBodyFieldsExpressionsAndLimitsInner, score int32, stream string, windowSize string) *CreateChainRuleRequestBodyFields {
	this := CreateChainRuleRequestBodyFields{}
	this.Enabled = enabled
	this.Name = name
	this.Description = description
	this.ExpressionsAndLimits = expressionsAndLimits
	this.Score = score
	this.Stream = stream
	this.WindowSize = windowSize
	return &this
}

// NewCreateChainRuleRequestBodyFieldsWithDefaults instantiates a new CreateChainRuleRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateChainRuleRequestBodyFieldsWithDefaults() *CreateChainRuleRequestBodyFields {
	this := CreateChainRuleRequestBodyFields{}
	return &this
}

// GetAssetField returns the AssetField field value if set, zero value otherwise.
func (o *CreateChainRuleRequestBodyFields) GetAssetField() string {
	if o == nil || IsNil(o.AssetField) {
		var ret string
		return ret
	}
	return *o.AssetField
}

// GetAssetFieldOk returns a tuple with the AssetField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetAssetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.AssetField) {
		return nil, false
	}
	return o.AssetField, true
}

// HasAssetField returns a boolean if a field has been set.
func (o *CreateChainRuleRequestBodyFields) HasAssetField() bool {
	if o != nil && !IsNil(o.AssetField) {
		return true
	}

	return false
}

// SetAssetField gets a reference to the given string and assigns it to the AssetField field.
func (o *CreateChainRuleRequestBodyFields) SetAssetField(v string) {
	o.AssetField = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CreateChainRuleRequestBodyFields) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CreateChainRuleRequestBodyFields) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *CreateChainRuleRequestBodyFields) SetCategory(v string) {
	o.Category = &v
}

// GetEnabled returns the Enabled field value
func (o *CreateChainRuleRequestBodyFields) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateChainRuleRequestBodyFields) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEntitySelectors returns the EntitySelectors field value if set, zero value otherwise.
func (o *CreateChainRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner {
	if o == nil || IsNil(o.EntitySelectors) {
		var ret []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner
		return ret
	}
	return o.EntitySelectors
}

// GetEntitySelectorsOk returns a tuple with the EntitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetEntitySelectorsOk() ([]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool) {
	if o == nil || IsNil(o.EntitySelectors) {
		return nil, false
	}
	return o.EntitySelectors, true
}

// HasEntitySelectors returns a boolean if a field has been set.
func (o *CreateChainRuleRequestBodyFields) HasEntitySelectors() bool {
	if o != nil && !IsNil(o.EntitySelectors) {
		return true
	}

	return false
}

// SetEntitySelectors gets a reference to the given []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner and assigns it to the EntitySelectors field.
func (o *CreateChainRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner) {
	o.EntitySelectors = v
}

// GetIsPrototype returns the IsPrototype field value if set, zero value otherwise.
func (o *CreateChainRuleRequestBodyFields) GetIsPrototype() bool {
	if o == nil || IsNil(o.IsPrototype) {
		var ret bool
		return ret
	}
	return *o.IsPrototype
}

// GetIsPrototypeOk returns a tuple with the IsPrototype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrototype) {
		return nil, false
	}
	return o.IsPrototype, true
}

// HasIsPrototype returns a boolean if a field has been set.
func (o *CreateChainRuleRequestBodyFields) HasIsPrototype() bool {
	if o != nil && !IsNil(o.IsPrototype) {
		return true
	}

	return false
}

// SetIsPrototype gets a reference to the given bool and assigns it to the IsPrototype field.
func (o *CreateChainRuleRequestBodyFields) SetIsPrototype(v bool) {
	o.IsPrototype = &v
}

// GetName returns the Name field value
func (o *CreateChainRuleRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateChainRuleRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetParentJaskId returns the ParentJaskId field value if set, zero value otherwise.
func (o *CreateChainRuleRequestBodyFields) GetParentJaskId() string {
	if o == nil || IsNil(o.ParentJaskId) {
		var ret string
		return ret
	}
	return *o.ParentJaskId
}

// GetParentJaskIdOk returns a tuple with the ParentJaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetParentJaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentJaskId) {
		return nil, false
	}
	return o.ParentJaskId, true
}

// HasParentJaskId returns a boolean if a field has been set.
func (o *CreateChainRuleRequestBodyFields) HasParentJaskId() bool {
	if o != nil && !IsNil(o.ParentJaskId) {
		return true
	}

	return false
}

// SetParentJaskId gets a reference to the given string and assigns it to the ParentJaskId field.
func (o *CreateChainRuleRequestBodyFields) SetParentJaskId(v string) {
	o.ParentJaskId = &v
}

// GetSummaryExpression returns the SummaryExpression field value if set, zero value otherwise.
func (o *CreateChainRuleRequestBodyFields) GetSummaryExpression() string {
	if o == nil || IsNil(o.SummaryExpression) {
		var ret string
		return ret
	}
	return *o.SummaryExpression
}

// GetSummaryExpressionOk returns a tuple with the SummaryExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.SummaryExpression) {
		return nil, false
	}
	return o.SummaryExpression, true
}

// HasSummaryExpression returns a boolean if a field has been set.
func (o *CreateChainRuleRequestBodyFields) HasSummaryExpression() bool {
	if o != nil && !IsNil(o.SummaryExpression) {
		return true
	}

	return false
}

// SetSummaryExpression gets a reference to the given string and assigns it to the SummaryExpression field.
func (o *CreateChainRuleRequestBodyFields) SetSummaryExpression(v string) {
	o.SummaryExpression = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateChainRuleRequestBodyFields) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateChainRuleRequestBodyFields) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateChainRuleRequestBodyFields) SetTags(v []string) {
	o.Tags = v
}

// GetTuningExpressionIds returns the TuningExpressionIds field value if set, zero value otherwise.
func (o *CreateChainRuleRequestBodyFields) GetTuningExpressionIds() []string {
	if o == nil || IsNil(o.TuningExpressionIds) {
		var ret []string
		return ret
	}
	return o.TuningExpressionIds
}

// GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetTuningExpressionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TuningExpressionIds) {
		return nil, false
	}
	return o.TuningExpressionIds, true
}

// HasTuningExpressionIds returns a boolean if a field has been set.
func (o *CreateChainRuleRequestBodyFields) HasTuningExpressionIds() bool {
	if o != nil && !IsNil(o.TuningExpressionIds) {
		return true
	}

	return false
}

// SetTuningExpressionIds gets a reference to the given []string and assigns it to the TuningExpressionIds field.
func (o *CreateChainRuleRequestBodyFields) SetTuningExpressionIds(v []string) {
	o.TuningExpressionIds = v
}

// GetDescription returns the Description field value
func (o *CreateChainRuleRequestBodyFields) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateChainRuleRequestBodyFields) SetDescription(v string) {
	o.Description = v
}

// GetExpressionsAndLimits returns the ExpressionsAndLimits field value
func (o *CreateChainRuleRequestBodyFields) GetExpressionsAndLimits() []CreateChainRuleRequestBodyFieldsExpressionsAndLimitsInner {
	if o == nil {
		var ret []CreateChainRuleRequestBodyFieldsExpressionsAndLimitsInner
		return ret
	}

	return o.ExpressionsAndLimits
}

// GetExpressionsAndLimitsOk returns a tuple with the ExpressionsAndLimits field value
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetExpressionsAndLimitsOk() ([]CreateChainRuleRequestBodyFieldsExpressionsAndLimitsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpressionsAndLimits, true
}

// SetExpressionsAndLimits sets field value
func (o *CreateChainRuleRequestBodyFields) SetExpressionsAndLimits(v []CreateChainRuleRequestBodyFieldsExpressionsAndLimitsInner) {
	o.ExpressionsAndLimits = v
}

// GetGroupByFields returns the GroupByFields field value if set, zero value otherwise.
func (o *CreateChainRuleRequestBodyFields) GetGroupByFields() []string {
	if o == nil || IsNil(o.GroupByFields) {
		var ret []string
		return ret
	}
	return o.GroupByFields
}

// GetGroupByFieldsOk returns a tuple with the GroupByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetGroupByFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupByFields) {
		return nil, false
	}
	return o.GroupByFields, true
}

// HasGroupByFields returns a boolean if a field has been set.
func (o *CreateChainRuleRequestBodyFields) HasGroupByFields() bool {
	if o != nil && !IsNil(o.GroupByFields) {
		return true
	}

	return false
}

// SetGroupByFields gets a reference to the given []string and assigns it to the GroupByFields field.
func (o *CreateChainRuleRequestBodyFields) SetGroupByFields(v []string) {
	o.GroupByFields = v
}

// GetOrdered returns the Ordered field value if set, zero value otherwise.
func (o *CreateChainRuleRequestBodyFields) GetOrdered() bool {
	if o == nil || IsNil(o.Ordered) {
		var ret bool
		return ret
	}
	return *o.Ordered
}

// GetOrderedOk returns a tuple with the Ordered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetOrderedOk() (*bool, bool) {
	if o == nil || IsNil(o.Ordered) {
		return nil, false
	}
	return o.Ordered, true
}

// HasOrdered returns a boolean if a field has been set.
func (o *CreateChainRuleRequestBodyFields) HasOrdered() bool {
	if o != nil && !IsNil(o.Ordered) {
		return true
	}

	return false
}

// SetOrdered gets a reference to the given bool and assigns it to the Ordered field.
func (o *CreateChainRuleRequestBodyFields) SetOrdered(v bool) {
	o.Ordered = &v
}

// GetScore returns the Score field value
func (o *CreateChainRuleRequestBodyFields) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *CreateChainRuleRequestBodyFields) SetScore(v int32) {
	o.Score = v
}

// GetStream returns the Stream field value
func (o *CreateChainRuleRequestBodyFields) GetStream() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stream
}

// GetStreamOk returns a tuple with the Stream field value
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stream, true
}

// SetStream sets field value
func (o *CreateChainRuleRequestBodyFields) SetStream(v string) {
	o.Stream = v
}

// GetWindowSize returns the WindowSize field value
func (o *CreateChainRuleRequestBodyFields) GetWindowSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WindowSize
}

// GetWindowSizeOk returns a tuple with the WindowSize field value
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetWindowSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowSize, true
}

// SetWindowSize sets field value
func (o *CreateChainRuleRequestBodyFields) SetWindowSize(v string) {
	o.WindowSize = v
}

// GetWindowSizeMilliseconds returns the WindowSizeMilliseconds field value if set, zero value otherwise.
func (o *CreateChainRuleRequestBodyFields) GetWindowSizeMilliseconds() string {
	if o == nil || IsNil(o.WindowSizeMilliseconds) {
		var ret string
		return ret
	}
	return *o.WindowSizeMilliseconds
}

// GetWindowSizeMillisecondsOk returns a tuple with the WindowSizeMilliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChainRuleRequestBodyFields) GetWindowSizeMillisecondsOk() (*string, bool) {
	if o == nil || IsNil(o.WindowSizeMilliseconds) {
		return nil, false
	}
	return o.WindowSizeMilliseconds, true
}

// HasWindowSizeMilliseconds returns a boolean if a field has been set.
func (o *CreateChainRuleRequestBodyFields) HasWindowSizeMilliseconds() bool {
	if o != nil && !IsNil(o.WindowSizeMilliseconds) {
		return true
	}

	return false
}

// SetWindowSizeMilliseconds gets a reference to the given string and assigns it to the WindowSizeMilliseconds field.
func (o *CreateChainRuleRequestBodyFields) SetWindowSizeMilliseconds(v string) {
	o.WindowSizeMilliseconds = &v
}

func (o CreateChainRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateChainRuleRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetField) {
		toSerialize["assetField"] = o.AssetField
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.EntitySelectors) {
		toSerialize["entitySelectors"] = o.EntitySelectors
	}
	if !IsNil(o.IsPrototype) {
		toSerialize["isPrototype"] = o.IsPrototype
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.ParentJaskId) {
		toSerialize["parentJaskId"] = o.ParentJaskId
	}
	if !IsNil(o.SummaryExpression) {
		toSerialize["summaryExpression"] = o.SummaryExpression
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TuningExpressionIds) {
		toSerialize["tuningExpressionIds"] = o.TuningExpressionIds
	}
	toSerialize["description"] = o.Description
	toSerialize["expressionsAndLimits"] = o.ExpressionsAndLimits
	if !IsNil(o.GroupByFields) {
		toSerialize["groupByFields"] = o.GroupByFields
	}
	if !IsNil(o.Ordered) {
		toSerialize["ordered"] = o.Ordered
	}
	toSerialize["score"] = o.Score
	toSerialize["stream"] = o.Stream
	toSerialize["windowSize"] = o.WindowSize
	if !IsNil(o.WindowSizeMilliseconds) {
		toSerialize["windowSizeMilliseconds"] = o.WindowSizeMilliseconds
	}
	return toSerialize, nil
}

func (o *CreateChainRuleRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"name",
		"description",
		"expressionsAndLimits",
		"score",
		"stream",
		"windowSize",
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

	varCreateChainRuleRequestBodyFields := _CreateChainRuleRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateChainRuleRequestBodyFields)

	if err != nil {
		return err
	}

	*o = CreateChainRuleRequestBodyFields(varCreateChainRuleRequestBodyFields)

	return err
}

type NullableCreateChainRuleRequestBodyFields struct {
	value *CreateChainRuleRequestBodyFields
	isSet bool
}

func (v NullableCreateChainRuleRequestBodyFields) Get() *CreateChainRuleRequestBodyFields {
	return v.value
}

func (v *NullableCreateChainRuleRequestBodyFields) Set(val *CreateChainRuleRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChainRuleRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChainRuleRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChainRuleRequestBodyFields(val *CreateChainRuleRequestBodyFields) *NullableCreateChainRuleRequestBodyFields {
	return &NullableCreateChainRuleRequestBodyFields{value: val, isSet: true}
}

func (v NullableCreateChainRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChainRuleRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


