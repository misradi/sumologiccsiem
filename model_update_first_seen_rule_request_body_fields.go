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

// checks if the UpdateFirstSeenRuleRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFirstSeenRuleRequestBodyFields{}

// UpdateFirstSeenRuleRequestBodyFields struct for UpdateFirstSeenRuleRequestBodyFields
type UpdateFirstSeenRuleRequestBodyFields struct {
	AssetField *string `json:"assetField,omitempty"`
	Category *string `json:"category,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	EntitySelectors []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner `json:"entitySelectors,omitempty"`
	IsPrototype *bool `json:"isPrototype,omitempty"`
	Name string `json:"name"`
	ParentJaskId *string `json:"parentJaskId,omitempty"`
	SummaryExpression *string `json:"summaryExpression,omitempty"`
	Tags []string `json:"tags,omitempty"`
	TuningExpressionIds []string `json:"tuningExpressionIds,omitempty"`
	// The description to be used on the generated Signal
	DescriptionExpression string `json:"descriptionExpression"`
	NameExpression string `json:"nameExpression"`
	FilterExpression string `json:"filterExpression"`
	// The value(s) to build an expression from - can be used instead of valueExpression
	ValueFields []string `json:"valueFields,omitempty"`
	ValueExpression *string `json:"valueExpression,omitempty"`
	GroupByFields []string `json:"groupByFields"`
	Score int32 `json:"score"`
	Version int32 `json:"version"`
	// milliseconds
	BaselineWindowSize string `json:"baselineWindowSize"`
	// milliseconds
	RetentionWindowSize string `json:"retentionWindowSize"`
	BaselineType *string `json:"baselineType,omitempty"`
}

type _UpdateFirstSeenRuleRequestBodyFields UpdateFirstSeenRuleRequestBodyFields

// NewUpdateFirstSeenRuleRequestBodyFields instantiates a new UpdateFirstSeenRuleRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFirstSeenRuleRequestBodyFields(name string, descriptionExpression string, nameExpression string, filterExpression string, groupByFields []string, score int32, version int32, baselineWindowSize string, retentionWindowSize string) *UpdateFirstSeenRuleRequestBodyFields {
	this := UpdateFirstSeenRuleRequestBodyFields{}
	this.Name = name
	this.DescriptionExpression = descriptionExpression
	this.NameExpression = nameExpression
	this.FilterExpression = filterExpression
	this.GroupByFields = groupByFields
	this.Score = score
	this.Version = version
	this.BaselineWindowSize = baselineWindowSize
	this.RetentionWindowSize = retentionWindowSize
	return &this
}

// NewUpdateFirstSeenRuleRequestBodyFieldsWithDefaults instantiates a new UpdateFirstSeenRuleRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFirstSeenRuleRequestBodyFieldsWithDefaults() *UpdateFirstSeenRuleRequestBodyFields {
	this := UpdateFirstSeenRuleRequestBodyFields{}
	return &this
}

// GetAssetField returns the AssetField field value if set, zero value otherwise.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetAssetField() string {
	if o == nil || IsNil(o.AssetField) {
		var ret string
		return ret
	}
	return *o.AssetField
}

// GetAssetFieldOk returns a tuple with the AssetField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetAssetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.AssetField) {
		return nil, false
	}
	return o.AssetField, true
}

// HasAssetField returns a boolean if a field has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) HasAssetField() bool {
	if o != nil && !IsNil(o.AssetField) {
		return true
	}

	return false
}

// SetAssetField gets a reference to the given string and assigns it to the AssetField field.
func (o *UpdateFirstSeenRuleRequestBodyFields) SetAssetField(v string) {
	o.AssetField = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *UpdateFirstSeenRuleRequestBodyFields) SetCategory(v string) {
	o.Category = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateFirstSeenRuleRequestBodyFields) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEntitySelectors returns the EntitySelectors field value if set, zero value otherwise.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner {
	if o == nil || IsNil(o.EntitySelectors) {
		var ret []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner
		return ret
	}
	return o.EntitySelectors
}

// GetEntitySelectorsOk returns a tuple with the EntitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetEntitySelectorsOk() ([]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool) {
	if o == nil || IsNil(o.EntitySelectors) {
		return nil, false
	}
	return o.EntitySelectors, true
}

// HasEntitySelectors returns a boolean if a field has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) HasEntitySelectors() bool {
	if o != nil && !IsNil(o.EntitySelectors) {
		return true
	}

	return false
}

// SetEntitySelectors gets a reference to the given []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner and assigns it to the EntitySelectors field.
func (o *UpdateFirstSeenRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner) {
	o.EntitySelectors = v
}

// GetIsPrototype returns the IsPrototype field value if set, zero value otherwise.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetIsPrototype() bool {
	if o == nil || IsNil(o.IsPrototype) {
		var ret bool
		return ret
	}
	return *o.IsPrototype
}

// GetIsPrototypeOk returns a tuple with the IsPrototype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrototype) {
		return nil, false
	}
	return o.IsPrototype, true
}

// HasIsPrototype returns a boolean if a field has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) HasIsPrototype() bool {
	if o != nil && !IsNil(o.IsPrototype) {
		return true
	}

	return false
}

// SetIsPrototype gets a reference to the given bool and assigns it to the IsPrototype field.
func (o *UpdateFirstSeenRuleRequestBodyFields) SetIsPrototype(v bool) {
	o.IsPrototype = &v
}

// GetName returns the Name field value
func (o *UpdateFirstSeenRuleRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateFirstSeenRuleRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetParentJaskId returns the ParentJaskId field value if set, zero value otherwise.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetParentJaskId() string {
	if o == nil || IsNil(o.ParentJaskId) {
		var ret string
		return ret
	}
	return *o.ParentJaskId
}

// GetParentJaskIdOk returns a tuple with the ParentJaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetParentJaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentJaskId) {
		return nil, false
	}
	return o.ParentJaskId, true
}

// HasParentJaskId returns a boolean if a field has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) HasParentJaskId() bool {
	if o != nil && !IsNil(o.ParentJaskId) {
		return true
	}

	return false
}

// SetParentJaskId gets a reference to the given string and assigns it to the ParentJaskId field.
func (o *UpdateFirstSeenRuleRequestBodyFields) SetParentJaskId(v string) {
	o.ParentJaskId = &v
}

// GetSummaryExpression returns the SummaryExpression field value if set, zero value otherwise.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetSummaryExpression() string {
	if o == nil || IsNil(o.SummaryExpression) {
		var ret string
		return ret
	}
	return *o.SummaryExpression
}

// GetSummaryExpressionOk returns a tuple with the SummaryExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.SummaryExpression) {
		return nil, false
	}
	return o.SummaryExpression, true
}

// HasSummaryExpression returns a boolean if a field has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) HasSummaryExpression() bool {
	if o != nil && !IsNil(o.SummaryExpression) {
		return true
	}

	return false
}

// SetSummaryExpression gets a reference to the given string and assigns it to the SummaryExpression field.
func (o *UpdateFirstSeenRuleRequestBodyFields) SetSummaryExpression(v string) {
	o.SummaryExpression = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateFirstSeenRuleRequestBodyFields) SetTags(v []string) {
	o.Tags = v
}

// GetTuningExpressionIds returns the TuningExpressionIds field value if set, zero value otherwise.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetTuningExpressionIds() []string {
	if o == nil || IsNil(o.TuningExpressionIds) {
		var ret []string
		return ret
	}
	return o.TuningExpressionIds
}

// GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetTuningExpressionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TuningExpressionIds) {
		return nil, false
	}
	return o.TuningExpressionIds, true
}

// HasTuningExpressionIds returns a boolean if a field has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) HasTuningExpressionIds() bool {
	if o != nil && !IsNil(o.TuningExpressionIds) {
		return true
	}

	return false
}

// SetTuningExpressionIds gets a reference to the given []string and assigns it to the TuningExpressionIds field.
func (o *UpdateFirstSeenRuleRequestBodyFields) SetTuningExpressionIds(v []string) {
	o.TuningExpressionIds = v
}

// GetDescriptionExpression returns the DescriptionExpression field value
func (o *UpdateFirstSeenRuleRequestBodyFields) GetDescriptionExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DescriptionExpression
}

// GetDescriptionExpressionOk returns a tuple with the DescriptionExpression field value
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetDescriptionExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DescriptionExpression, true
}

// SetDescriptionExpression sets field value
func (o *UpdateFirstSeenRuleRequestBodyFields) SetDescriptionExpression(v string) {
	o.DescriptionExpression = v
}

// GetNameExpression returns the NameExpression field value
func (o *UpdateFirstSeenRuleRequestBodyFields) GetNameExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameExpression
}

// GetNameExpressionOk returns a tuple with the NameExpression field value
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetNameExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameExpression, true
}

// SetNameExpression sets field value
func (o *UpdateFirstSeenRuleRequestBodyFields) SetNameExpression(v string) {
	o.NameExpression = v
}

// GetFilterExpression returns the FilterExpression field value
func (o *UpdateFirstSeenRuleRequestBodyFields) GetFilterExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterExpression
}

// GetFilterExpressionOk returns a tuple with the FilterExpression field value
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetFilterExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterExpression, true
}

// SetFilterExpression sets field value
func (o *UpdateFirstSeenRuleRequestBodyFields) SetFilterExpression(v string) {
	o.FilterExpression = v
}

// GetValueFields returns the ValueFields field value if set, zero value otherwise.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetValueFields() []string {
	if o == nil || IsNil(o.ValueFields) {
		var ret []string
		return ret
	}
	return o.ValueFields
}

// GetValueFieldsOk returns a tuple with the ValueFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetValueFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.ValueFields) {
		return nil, false
	}
	return o.ValueFields, true
}

// HasValueFields returns a boolean if a field has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) HasValueFields() bool {
	if o != nil && !IsNil(o.ValueFields) {
		return true
	}

	return false
}

// SetValueFields gets a reference to the given []string and assigns it to the ValueFields field.
func (o *UpdateFirstSeenRuleRequestBodyFields) SetValueFields(v []string) {
	o.ValueFields = v
}

// GetValueExpression returns the ValueExpression field value if set, zero value otherwise.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetValueExpression() string {
	if o == nil || IsNil(o.ValueExpression) {
		var ret string
		return ret
	}
	return *o.ValueExpression
}

// GetValueExpressionOk returns a tuple with the ValueExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetValueExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.ValueExpression) {
		return nil, false
	}
	return o.ValueExpression, true
}

// HasValueExpression returns a boolean if a field has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) HasValueExpression() bool {
	if o != nil && !IsNil(o.ValueExpression) {
		return true
	}

	return false
}

// SetValueExpression gets a reference to the given string and assigns it to the ValueExpression field.
func (o *UpdateFirstSeenRuleRequestBodyFields) SetValueExpression(v string) {
	o.ValueExpression = &v
}

// GetGroupByFields returns the GroupByFields field value
func (o *UpdateFirstSeenRuleRequestBodyFields) GetGroupByFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GroupByFields
}

// GetGroupByFieldsOk returns a tuple with the GroupByFields field value
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetGroupByFieldsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupByFields, true
}

// SetGroupByFields sets field value
func (o *UpdateFirstSeenRuleRequestBodyFields) SetGroupByFields(v []string) {
	o.GroupByFields = v
}

// GetScore returns the Score field value
func (o *UpdateFirstSeenRuleRequestBodyFields) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *UpdateFirstSeenRuleRequestBodyFields) SetScore(v int32) {
	o.Score = v
}

// GetVersion returns the Version field value
func (o *UpdateFirstSeenRuleRequestBodyFields) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *UpdateFirstSeenRuleRequestBodyFields) SetVersion(v int32) {
	o.Version = v
}

// GetBaselineWindowSize returns the BaselineWindowSize field value
func (o *UpdateFirstSeenRuleRequestBodyFields) GetBaselineWindowSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaselineWindowSize
}

// GetBaselineWindowSizeOk returns a tuple with the BaselineWindowSize field value
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetBaselineWindowSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaselineWindowSize, true
}

// SetBaselineWindowSize sets field value
func (o *UpdateFirstSeenRuleRequestBodyFields) SetBaselineWindowSize(v string) {
	o.BaselineWindowSize = v
}

// GetRetentionWindowSize returns the RetentionWindowSize field value
func (o *UpdateFirstSeenRuleRequestBodyFields) GetRetentionWindowSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RetentionWindowSize
}

// GetRetentionWindowSizeOk returns a tuple with the RetentionWindowSize field value
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetRetentionWindowSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionWindowSize, true
}

// SetRetentionWindowSize sets field value
func (o *UpdateFirstSeenRuleRequestBodyFields) SetRetentionWindowSize(v string) {
	o.RetentionWindowSize = v
}

// GetBaselineType returns the BaselineType field value if set, zero value otherwise.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetBaselineType() string {
	if o == nil || IsNil(o.BaselineType) {
		var ret string
		return ret
	}
	return *o.BaselineType
}

// GetBaselineTypeOk returns a tuple with the BaselineType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) GetBaselineTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BaselineType) {
		return nil, false
	}
	return o.BaselineType, true
}

// HasBaselineType returns a boolean if a field has been set.
func (o *UpdateFirstSeenRuleRequestBodyFields) HasBaselineType() bool {
	if o != nil && !IsNil(o.BaselineType) {
		return true
	}

	return false
}

// SetBaselineType gets a reference to the given string and assigns it to the BaselineType field.
func (o *UpdateFirstSeenRuleRequestBodyFields) SetBaselineType(v string) {
	o.BaselineType = &v
}

func (o UpdateFirstSeenRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFirstSeenRuleRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetField) {
		toSerialize["assetField"] = o.AssetField
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
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
	toSerialize["descriptionExpression"] = o.DescriptionExpression
	toSerialize["nameExpression"] = o.NameExpression
	toSerialize["filterExpression"] = o.FilterExpression
	if !IsNil(o.ValueFields) {
		toSerialize["valueFields"] = o.ValueFields
	}
	if !IsNil(o.ValueExpression) {
		toSerialize["valueExpression"] = o.ValueExpression
	}
	toSerialize["groupByFields"] = o.GroupByFields
	toSerialize["score"] = o.Score
	toSerialize["version"] = o.Version
	toSerialize["baselineWindowSize"] = o.BaselineWindowSize
	toSerialize["retentionWindowSize"] = o.RetentionWindowSize
	if !IsNil(o.BaselineType) {
		toSerialize["baselineType"] = o.BaselineType
	}
	return toSerialize, nil
}

func (o *UpdateFirstSeenRuleRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"descriptionExpression",
		"nameExpression",
		"filterExpression",
		"groupByFields",
		"score",
		"version",
		"baselineWindowSize",
		"retentionWindowSize",
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

	varUpdateFirstSeenRuleRequestBodyFields := _UpdateFirstSeenRuleRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateFirstSeenRuleRequestBodyFields)

	if err != nil {
		return err
	}

	*o = UpdateFirstSeenRuleRequestBodyFields(varUpdateFirstSeenRuleRequestBodyFields)

	return err
}

type NullableUpdateFirstSeenRuleRequestBodyFields struct {
	value *UpdateFirstSeenRuleRequestBodyFields
	isSet bool
}

func (v NullableUpdateFirstSeenRuleRequestBodyFields) Get() *UpdateFirstSeenRuleRequestBodyFields {
	return v.value
}

func (v *NullableUpdateFirstSeenRuleRequestBodyFields) Set(val *UpdateFirstSeenRuleRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFirstSeenRuleRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFirstSeenRuleRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFirstSeenRuleRequestBodyFields(val *UpdateFirstSeenRuleRequestBodyFields) *NullableUpdateFirstSeenRuleRequestBodyFields {
	return &NullableUpdateFirstSeenRuleRequestBodyFields{value: val, isSet: true}
}

func (v NullableUpdateFirstSeenRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFirstSeenRuleRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


