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

// checks if the CreateAggregationRuleRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAggregationRuleRequestBodyFields{}

// CreateAggregationRuleRequestBodyFields struct for CreateAggregationRuleRequestBodyFields
type CreateAggregationRuleRequestBodyFields struct {
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
	AggregationFunctions []GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner `json:"aggregationFunctions"`
	DescriptionExpression string `json:"descriptionExpression"`
	GroupByAsset bool `json:"groupByAsset"`
	GroupByFields []string `json:"groupByFields"`
	MatchExpression string `json:"matchExpression"`
	NameExpression string `json:"nameExpression"`
	ScoreMapping GetRulesResponseDataObjectsInnerOneOf2ScoreMapping `json:"scoreMapping"`
	Stream string `json:"stream"`
	TriggerExpression string `json:"triggerExpression"`
	WindowSize string `json:"windowSize"`
	// Can be used instead of windowSize.
	WindowSizeMilliseconds *string `json:"windowSizeMilliseconds,omitempty"`
}

type _CreateAggregationRuleRequestBodyFields CreateAggregationRuleRequestBodyFields

// NewCreateAggregationRuleRequestBodyFields instantiates a new CreateAggregationRuleRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAggregationRuleRequestBodyFields(enabled bool, name string, aggregationFunctions []GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner, descriptionExpression string, groupByAsset bool, groupByFields []string, matchExpression string, nameExpression string, scoreMapping GetRulesResponseDataObjectsInnerOneOf2ScoreMapping, stream string, triggerExpression string, windowSize string) *CreateAggregationRuleRequestBodyFields {
	this := CreateAggregationRuleRequestBodyFields{}
	this.Enabled = enabled
	this.Name = name
	this.AggregationFunctions = aggregationFunctions
	this.DescriptionExpression = descriptionExpression
	this.GroupByAsset = groupByAsset
	this.GroupByFields = groupByFields
	this.MatchExpression = matchExpression
	this.NameExpression = nameExpression
	this.ScoreMapping = scoreMapping
	this.Stream = stream
	this.TriggerExpression = triggerExpression
	this.WindowSize = windowSize
	return &this
}

// NewCreateAggregationRuleRequestBodyFieldsWithDefaults instantiates a new CreateAggregationRuleRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAggregationRuleRequestBodyFieldsWithDefaults() *CreateAggregationRuleRequestBodyFields {
	this := CreateAggregationRuleRequestBodyFields{}
	return &this
}

// GetAssetField returns the AssetField field value if set, zero value otherwise.
func (o *CreateAggregationRuleRequestBodyFields) GetAssetField() string {
	if o == nil || IsNil(o.AssetField) {
		var ret string
		return ret
	}
	return *o.AssetField
}

// GetAssetFieldOk returns a tuple with the AssetField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetAssetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.AssetField) {
		return nil, false
	}
	return o.AssetField, true
}

// HasAssetField returns a boolean if a field has been set.
func (o *CreateAggregationRuleRequestBodyFields) HasAssetField() bool {
	if o != nil && !IsNil(o.AssetField) {
		return true
	}

	return false
}

// SetAssetField gets a reference to the given string and assigns it to the AssetField field.
func (o *CreateAggregationRuleRequestBodyFields) SetAssetField(v string) {
	o.AssetField = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CreateAggregationRuleRequestBodyFields) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CreateAggregationRuleRequestBodyFields) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *CreateAggregationRuleRequestBodyFields) SetCategory(v string) {
	o.Category = &v
}

// GetEnabled returns the Enabled field value
func (o *CreateAggregationRuleRequestBodyFields) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateAggregationRuleRequestBodyFields) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEntitySelectors returns the EntitySelectors field value if set, zero value otherwise.
func (o *CreateAggregationRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner {
	if o == nil || IsNil(o.EntitySelectors) {
		var ret []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner
		return ret
	}
	return o.EntitySelectors
}

// GetEntitySelectorsOk returns a tuple with the EntitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetEntitySelectorsOk() ([]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool) {
	if o == nil || IsNil(o.EntitySelectors) {
		return nil, false
	}
	return o.EntitySelectors, true
}

// HasEntitySelectors returns a boolean if a field has been set.
func (o *CreateAggregationRuleRequestBodyFields) HasEntitySelectors() bool {
	if o != nil && !IsNil(o.EntitySelectors) {
		return true
	}

	return false
}

// SetEntitySelectors gets a reference to the given []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner and assigns it to the EntitySelectors field.
func (o *CreateAggregationRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner) {
	o.EntitySelectors = v
}

// GetIsPrototype returns the IsPrototype field value if set, zero value otherwise.
func (o *CreateAggregationRuleRequestBodyFields) GetIsPrototype() bool {
	if o == nil || IsNil(o.IsPrototype) {
		var ret bool
		return ret
	}
	return *o.IsPrototype
}

// GetIsPrototypeOk returns a tuple with the IsPrototype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrototype) {
		return nil, false
	}
	return o.IsPrototype, true
}

// HasIsPrototype returns a boolean if a field has been set.
func (o *CreateAggregationRuleRequestBodyFields) HasIsPrototype() bool {
	if o != nil && !IsNil(o.IsPrototype) {
		return true
	}

	return false
}

// SetIsPrototype gets a reference to the given bool and assigns it to the IsPrototype field.
func (o *CreateAggregationRuleRequestBodyFields) SetIsPrototype(v bool) {
	o.IsPrototype = &v
}

// GetName returns the Name field value
func (o *CreateAggregationRuleRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAggregationRuleRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetParentJaskId returns the ParentJaskId field value if set, zero value otherwise.
func (o *CreateAggregationRuleRequestBodyFields) GetParentJaskId() string {
	if o == nil || IsNil(o.ParentJaskId) {
		var ret string
		return ret
	}
	return *o.ParentJaskId
}

// GetParentJaskIdOk returns a tuple with the ParentJaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetParentJaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentJaskId) {
		return nil, false
	}
	return o.ParentJaskId, true
}

// HasParentJaskId returns a boolean if a field has been set.
func (o *CreateAggregationRuleRequestBodyFields) HasParentJaskId() bool {
	if o != nil && !IsNil(o.ParentJaskId) {
		return true
	}

	return false
}

// SetParentJaskId gets a reference to the given string and assigns it to the ParentJaskId field.
func (o *CreateAggregationRuleRequestBodyFields) SetParentJaskId(v string) {
	o.ParentJaskId = &v
}

// GetSummaryExpression returns the SummaryExpression field value if set, zero value otherwise.
func (o *CreateAggregationRuleRequestBodyFields) GetSummaryExpression() string {
	if o == nil || IsNil(o.SummaryExpression) {
		var ret string
		return ret
	}
	return *o.SummaryExpression
}

// GetSummaryExpressionOk returns a tuple with the SummaryExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.SummaryExpression) {
		return nil, false
	}
	return o.SummaryExpression, true
}

// HasSummaryExpression returns a boolean if a field has been set.
func (o *CreateAggregationRuleRequestBodyFields) HasSummaryExpression() bool {
	if o != nil && !IsNil(o.SummaryExpression) {
		return true
	}

	return false
}

// SetSummaryExpression gets a reference to the given string and assigns it to the SummaryExpression field.
func (o *CreateAggregationRuleRequestBodyFields) SetSummaryExpression(v string) {
	o.SummaryExpression = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateAggregationRuleRequestBodyFields) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateAggregationRuleRequestBodyFields) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateAggregationRuleRequestBodyFields) SetTags(v []string) {
	o.Tags = v
}

// GetTuningExpressionIds returns the TuningExpressionIds field value if set, zero value otherwise.
func (o *CreateAggregationRuleRequestBodyFields) GetTuningExpressionIds() []string {
	if o == nil || IsNil(o.TuningExpressionIds) {
		var ret []string
		return ret
	}
	return o.TuningExpressionIds
}

// GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetTuningExpressionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TuningExpressionIds) {
		return nil, false
	}
	return o.TuningExpressionIds, true
}

// HasTuningExpressionIds returns a boolean if a field has been set.
func (o *CreateAggregationRuleRequestBodyFields) HasTuningExpressionIds() bool {
	if o != nil && !IsNil(o.TuningExpressionIds) {
		return true
	}

	return false
}

// SetTuningExpressionIds gets a reference to the given []string and assigns it to the TuningExpressionIds field.
func (o *CreateAggregationRuleRequestBodyFields) SetTuningExpressionIds(v []string) {
	o.TuningExpressionIds = v
}

// GetAggregationFunctions returns the AggregationFunctions field value
func (o *CreateAggregationRuleRequestBodyFields) GetAggregationFunctions() []GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner {
	if o == nil {
		var ret []GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner
		return ret
	}

	return o.AggregationFunctions
}

// GetAggregationFunctionsOk returns a tuple with the AggregationFunctions field value
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetAggregationFunctionsOk() ([]GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.AggregationFunctions, true
}

// SetAggregationFunctions sets field value
func (o *CreateAggregationRuleRequestBodyFields) SetAggregationFunctions(v []GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner) {
	o.AggregationFunctions = v
}

// GetDescriptionExpression returns the DescriptionExpression field value
func (o *CreateAggregationRuleRequestBodyFields) GetDescriptionExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DescriptionExpression
}

// GetDescriptionExpressionOk returns a tuple with the DescriptionExpression field value
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetDescriptionExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DescriptionExpression, true
}

// SetDescriptionExpression sets field value
func (o *CreateAggregationRuleRequestBodyFields) SetDescriptionExpression(v string) {
	o.DescriptionExpression = v
}

// GetGroupByAsset returns the GroupByAsset field value
func (o *CreateAggregationRuleRequestBodyFields) GetGroupByAsset() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.GroupByAsset
}

// GetGroupByAssetOk returns a tuple with the GroupByAsset field value
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetGroupByAssetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupByAsset, true
}

// SetGroupByAsset sets field value
func (o *CreateAggregationRuleRequestBodyFields) SetGroupByAsset(v bool) {
	o.GroupByAsset = v
}

// GetGroupByFields returns the GroupByFields field value
func (o *CreateAggregationRuleRequestBodyFields) GetGroupByFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GroupByFields
}

// GetGroupByFieldsOk returns a tuple with the GroupByFields field value
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetGroupByFieldsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupByFields, true
}

// SetGroupByFields sets field value
func (o *CreateAggregationRuleRequestBodyFields) SetGroupByFields(v []string) {
	o.GroupByFields = v
}

// GetMatchExpression returns the MatchExpression field value
func (o *CreateAggregationRuleRequestBodyFields) GetMatchExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MatchExpression
}

// GetMatchExpressionOk returns a tuple with the MatchExpression field value
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetMatchExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchExpression, true
}

// SetMatchExpression sets field value
func (o *CreateAggregationRuleRequestBodyFields) SetMatchExpression(v string) {
	o.MatchExpression = v
}

// GetNameExpression returns the NameExpression field value
func (o *CreateAggregationRuleRequestBodyFields) GetNameExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameExpression
}

// GetNameExpressionOk returns a tuple with the NameExpression field value
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetNameExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameExpression, true
}

// SetNameExpression sets field value
func (o *CreateAggregationRuleRequestBodyFields) SetNameExpression(v string) {
	o.NameExpression = v
}

// GetScoreMapping returns the ScoreMapping field value
func (o *CreateAggregationRuleRequestBodyFields) GetScoreMapping() GetRulesResponseDataObjectsInnerOneOf2ScoreMapping {
	if o == nil {
		var ret GetRulesResponseDataObjectsInnerOneOf2ScoreMapping
		return ret
	}

	return o.ScoreMapping
}

// GetScoreMappingOk returns a tuple with the ScoreMapping field value
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetScoreMappingOk() (*GetRulesResponseDataObjectsInnerOneOf2ScoreMapping, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScoreMapping, true
}

// SetScoreMapping sets field value
func (o *CreateAggregationRuleRequestBodyFields) SetScoreMapping(v GetRulesResponseDataObjectsInnerOneOf2ScoreMapping) {
	o.ScoreMapping = v
}

// GetStream returns the Stream field value
func (o *CreateAggregationRuleRequestBodyFields) GetStream() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stream
}

// GetStreamOk returns a tuple with the Stream field value
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stream, true
}

// SetStream sets field value
func (o *CreateAggregationRuleRequestBodyFields) SetStream(v string) {
	o.Stream = v
}

// GetTriggerExpression returns the TriggerExpression field value
func (o *CreateAggregationRuleRequestBodyFields) GetTriggerExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggerExpression
}

// GetTriggerExpressionOk returns a tuple with the TriggerExpression field value
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetTriggerExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerExpression, true
}

// SetTriggerExpression sets field value
func (o *CreateAggregationRuleRequestBodyFields) SetTriggerExpression(v string) {
	o.TriggerExpression = v
}

// GetWindowSize returns the WindowSize field value
func (o *CreateAggregationRuleRequestBodyFields) GetWindowSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WindowSize
}

// GetWindowSizeOk returns a tuple with the WindowSize field value
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetWindowSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowSize, true
}

// SetWindowSize sets field value
func (o *CreateAggregationRuleRequestBodyFields) SetWindowSize(v string) {
	o.WindowSize = v
}

// GetWindowSizeMilliseconds returns the WindowSizeMilliseconds field value if set, zero value otherwise.
func (o *CreateAggregationRuleRequestBodyFields) GetWindowSizeMilliseconds() string {
	if o == nil || IsNil(o.WindowSizeMilliseconds) {
		var ret string
		return ret
	}
	return *o.WindowSizeMilliseconds
}

// GetWindowSizeMillisecondsOk returns a tuple with the WindowSizeMilliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAggregationRuleRequestBodyFields) GetWindowSizeMillisecondsOk() (*string, bool) {
	if o == nil || IsNil(o.WindowSizeMilliseconds) {
		return nil, false
	}
	return o.WindowSizeMilliseconds, true
}

// HasWindowSizeMilliseconds returns a boolean if a field has been set.
func (o *CreateAggregationRuleRequestBodyFields) HasWindowSizeMilliseconds() bool {
	if o != nil && !IsNil(o.WindowSizeMilliseconds) {
		return true
	}

	return false
}

// SetWindowSizeMilliseconds gets a reference to the given string and assigns it to the WindowSizeMilliseconds field.
func (o *CreateAggregationRuleRequestBodyFields) SetWindowSizeMilliseconds(v string) {
	o.WindowSizeMilliseconds = &v
}

func (o CreateAggregationRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAggregationRuleRequestBodyFields) ToMap() (map[string]interface{}, error) {
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
	toSerialize["aggregationFunctions"] = o.AggregationFunctions
	toSerialize["descriptionExpression"] = o.DescriptionExpression
	toSerialize["groupByAsset"] = o.GroupByAsset
	toSerialize["groupByFields"] = o.GroupByFields
	toSerialize["matchExpression"] = o.MatchExpression
	toSerialize["nameExpression"] = o.NameExpression
	toSerialize["scoreMapping"] = o.ScoreMapping
	toSerialize["stream"] = o.Stream
	toSerialize["triggerExpression"] = o.TriggerExpression
	toSerialize["windowSize"] = o.WindowSize
	if !IsNil(o.WindowSizeMilliseconds) {
		toSerialize["windowSizeMilliseconds"] = o.WindowSizeMilliseconds
	}
	return toSerialize, nil
}

func (o *CreateAggregationRuleRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"name",
		"aggregationFunctions",
		"descriptionExpression",
		"groupByAsset",
		"groupByFields",
		"matchExpression",
		"nameExpression",
		"scoreMapping",
		"stream",
		"triggerExpression",
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

	varCreateAggregationRuleRequestBodyFields := _CreateAggregationRuleRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateAggregationRuleRequestBodyFields)

	if err != nil {
		return err
	}

	*o = CreateAggregationRuleRequestBodyFields(varCreateAggregationRuleRequestBodyFields)

	return err
}

type NullableCreateAggregationRuleRequestBodyFields struct {
	value *CreateAggregationRuleRequestBodyFields
	isSet bool
}

func (v NullableCreateAggregationRuleRequestBodyFields) Get() *CreateAggregationRuleRequestBodyFields {
	return v.value
}

func (v *NullableCreateAggregationRuleRequestBodyFields) Set(val *CreateAggregationRuleRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAggregationRuleRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAggregationRuleRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAggregationRuleRequestBodyFields(val *CreateAggregationRuleRequestBodyFields) *NullableCreateAggregationRuleRequestBodyFields {
	return &NullableCreateAggregationRuleRequestBodyFields{value: val, isSet: true}
}

func (v NullableCreateAggregationRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAggregationRuleRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


