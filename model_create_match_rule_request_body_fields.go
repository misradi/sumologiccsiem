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

// checks if the CreateMatchRuleRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMatchRuleRequestBodyFields{}

// CreateMatchRuleRequestBodyFields struct for CreateMatchRuleRequestBodyFields
type CreateMatchRuleRequestBodyFields struct {
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
	Expression string `json:"expression"`
	Score int32 `json:"score"`
	Stream string `json:"stream"`
}

type _CreateMatchRuleRequestBodyFields CreateMatchRuleRequestBodyFields

// NewCreateMatchRuleRequestBodyFields instantiates a new CreateMatchRuleRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMatchRuleRequestBodyFields(enabled bool, name string, description string, expression string, score int32, stream string) *CreateMatchRuleRequestBodyFields {
	this := CreateMatchRuleRequestBodyFields{}
	this.Enabled = enabled
	this.Name = name
	this.Description = description
	this.Expression = expression
	this.Score = score
	this.Stream = stream
	return &this
}

// NewCreateMatchRuleRequestBodyFieldsWithDefaults instantiates a new CreateMatchRuleRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMatchRuleRequestBodyFieldsWithDefaults() *CreateMatchRuleRequestBodyFields {
	this := CreateMatchRuleRequestBodyFields{}
	return &this
}

// GetAssetField returns the AssetField field value if set, zero value otherwise.
func (o *CreateMatchRuleRequestBodyFields) GetAssetField() string {
	if o == nil || IsNil(o.AssetField) {
		var ret string
		return ret
	}
	return *o.AssetField
}

// GetAssetFieldOk returns a tuple with the AssetField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMatchRuleRequestBodyFields) GetAssetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.AssetField) {
		return nil, false
	}
	return o.AssetField, true
}

// HasAssetField returns a boolean if a field has been set.
func (o *CreateMatchRuleRequestBodyFields) HasAssetField() bool {
	if o != nil && !IsNil(o.AssetField) {
		return true
	}

	return false
}

// SetAssetField gets a reference to the given string and assigns it to the AssetField field.
func (o *CreateMatchRuleRequestBodyFields) SetAssetField(v string) {
	o.AssetField = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CreateMatchRuleRequestBodyFields) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMatchRuleRequestBodyFields) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CreateMatchRuleRequestBodyFields) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *CreateMatchRuleRequestBodyFields) SetCategory(v string) {
	o.Category = &v
}

// GetEnabled returns the Enabled field value
func (o *CreateMatchRuleRequestBodyFields) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateMatchRuleRequestBodyFields) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateMatchRuleRequestBodyFields) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEntitySelectors returns the EntitySelectors field value if set, zero value otherwise.
func (o *CreateMatchRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner {
	if o == nil || IsNil(o.EntitySelectors) {
		var ret []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner
		return ret
	}
	return o.EntitySelectors
}

// GetEntitySelectorsOk returns a tuple with the EntitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMatchRuleRequestBodyFields) GetEntitySelectorsOk() ([]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool) {
	if o == nil || IsNil(o.EntitySelectors) {
		return nil, false
	}
	return o.EntitySelectors, true
}

// HasEntitySelectors returns a boolean if a field has been set.
func (o *CreateMatchRuleRequestBodyFields) HasEntitySelectors() bool {
	if o != nil && !IsNil(o.EntitySelectors) {
		return true
	}

	return false
}

// SetEntitySelectors gets a reference to the given []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner and assigns it to the EntitySelectors field.
func (o *CreateMatchRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner) {
	o.EntitySelectors = v
}

// GetIsPrototype returns the IsPrototype field value if set, zero value otherwise.
func (o *CreateMatchRuleRequestBodyFields) GetIsPrototype() bool {
	if o == nil || IsNil(o.IsPrototype) {
		var ret bool
		return ret
	}
	return *o.IsPrototype
}

// GetIsPrototypeOk returns a tuple with the IsPrototype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMatchRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrototype) {
		return nil, false
	}
	return o.IsPrototype, true
}

// HasIsPrototype returns a boolean if a field has been set.
func (o *CreateMatchRuleRequestBodyFields) HasIsPrototype() bool {
	if o != nil && !IsNil(o.IsPrototype) {
		return true
	}

	return false
}

// SetIsPrototype gets a reference to the given bool and assigns it to the IsPrototype field.
func (o *CreateMatchRuleRequestBodyFields) SetIsPrototype(v bool) {
	o.IsPrototype = &v
}

// GetName returns the Name field value
func (o *CreateMatchRuleRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateMatchRuleRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateMatchRuleRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetParentJaskId returns the ParentJaskId field value if set, zero value otherwise.
func (o *CreateMatchRuleRequestBodyFields) GetParentJaskId() string {
	if o == nil || IsNil(o.ParentJaskId) {
		var ret string
		return ret
	}
	return *o.ParentJaskId
}

// GetParentJaskIdOk returns a tuple with the ParentJaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMatchRuleRequestBodyFields) GetParentJaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentJaskId) {
		return nil, false
	}
	return o.ParentJaskId, true
}

// HasParentJaskId returns a boolean if a field has been set.
func (o *CreateMatchRuleRequestBodyFields) HasParentJaskId() bool {
	if o != nil && !IsNil(o.ParentJaskId) {
		return true
	}

	return false
}

// SetParentJaskId gets a reference to the given string and assigns it to the ParentJaskId field.
func (o *CreateMatchRuleRequestBodyFields) SetParentJaskId(v string) {
	o.ParentJaskId = &v
}

// GetSummaryExpression returns the SummaryExpression field value if set, zero value otherwise.
func (o *CreateMatchRuleRequestBodyFields) GetSummaryExpression() string {
	if o == nil || IsNil(o.SummaryExpression) {
		var ret string
		return ret
	}
	return *o.SummaryExpression
}

// GetSummaryExpressionOk returns a tuple with the SummaryExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMatchRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.SummaryExpression) {
		return nil, false
	}
	return o.SummaryExpression, true
}

// HasSummaryExpression returns a boolean if a field has been set.
func (o *CreateMatchRuleRequestBodyFields) HasSummaryExpression() bool {
	if o != nil && !IsNil(o.SummaryExpression) {
		return true
	}

	return false
}

// SetSummaryExpression gets a reference to the given string and assigns it to the SummaryExpression field.
func (o *CreateMatchRuleRequestBodyFields) SetSummaryExpression(v string) {
	o.SummaryExpression = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateMatchRuleRequestBodyFields) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMatchRuleRequestBodyFields) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateMatchRuleRequestBodyFields) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateMatchRuleRequestBodyFields) SetTags(v []string) {
	o.Tags = v
}

// GetTuningExpressionIds returns the TuningExpressionIds field value if set, zero value otherwise.
func (o *CreateMatchRuleRequestBodyFields) GetTuningExpressionIds() []string {
	if o == nil || IsNil(o.TuningExpressionIds) {
		var ret []string
		return ret
	}
	return o.TuningExpressionIds
}

// GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMatchRuleRequestBodyFields) GetTuningExpressionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TuningExpressionIds) {
		return nil, false
	}
	return o.TuningExpressionIds, true
}

// HasTuningExpressionIds returns a boolean if a field has been set.
func (o *CreateMatchRuleRequestBodyFields) HasTuningExpressionIds() bool {
	if o != nil && !IsNil(o.TuningExpressionIds) {
		return true
	}

	return false
}

// SetTuningExpressionIds gets a reference to the given []string and assigns it to the TuningExpressionIds field.
func (o *CreateMatchRuleRequestBodyFields) SetTuningExpressionIds(v []string) {
	o.TuningExpressionIds = v
}

// GetDescription returns the Description field value
func (o *CreateMatchRuleRequestBodyFields) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateMatchRuleRequestBodyFields) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateMatchRuleRequestBodyFields) SetDescription(v string) {
	o.Description = v
}

// GetExpression returns the Expression field value
func (o *CreateMatchRuleRequestBodyFields) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *CreateMatchRuleRequestBodyFields) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *CreateMatchRuleRequestBodyFields) SetExpression(v string) {
	o.Expression = v
}

// GetScore returns the Score field value
func (o *CreateMatchRuleRequestBodyFields) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *CreateMatchRuleRequestBodyFields) GetScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *CreateMatchRuleRequestBodyFields) SetScore(v int32) {
	o.Score = v
}

// GetStream returns the Stream field value
func (o *CreateMatchRuleRequestBodyFields) GetStream() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stream
}

// GetStreamOk returns a tuple with the Stream field value
// and a boolean to check if the value has been set.
func (o *CreateMatchRuleRequestBodyFields) GetStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stream, true
}

// SetStream sets field value
func (o *CreateMatchRuleRequestBodyFields) SetStream(v string) {
	o.Stream = v
}

func (o CreateMatchRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMatchRuleRequestBodyFields) ToMap() (map[string]interface{}, error) {
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
	toSerialize["expression"] = o.Expression
	toSerialize["score"] = o.Score
	toSerialize["stream"] = o.Stream
	return toSerialize, nil
}

func (o *CreateMatchRuleRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"name",
		"description",
		"expression",
		"score",
		"stream",
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

	varCreateMatchRuleRequestBodyFields := _CreateMatchRuleRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateMatchRuleRequestBodyFields)

	if err != nil {
		return err
	}

	*o = CreateMatchRuleRequestBodyFields(varCreateMatchRuleRequestBodyFields)

	return err
}

type NullableCreateMatchRuleRequestBodyFields struct {
	value *CreateMatchRuleRequestBodyFields
	isSet bool
}

func (v NullableCreateMatchRuleRequestBodyFields) Get() *CreateMatchRuleRequestBodyFields {
	return v.value
}

func (v *NullableCreateMatchRuleRequestBodyFields) Set(val *CreateMatchRuleRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMatchRuleRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMatchRuleRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMatchRuleRequestBodyFields(val *CreateMatchRuleRequestBodyFields) *NullableCreateMatchRuleRequestBodyFields {
	return &NullableCreateMatchRuleRequestBodyFields{value: val, isSet: true}
}

func (v NullableCreateMatchRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMatchRuleRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


