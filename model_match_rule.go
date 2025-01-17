/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologiccsiem

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the MatchRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MatchRule{}

// MatchRule struct for MatchRule
type MatchRule struct {
	Category *string `json:"category,omitempty"`
	ContentType string `json:"contentType"`
	Created time.Time `json:"created"`
	CreatedBy string `json:"createdBy"`
	Deleted bool `json:"deleted"`
	Enabled bool `json:"enabled"`
	EntitySelectors []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner `json:"entitySelectors"`
	Id string `json:"id"`
	IsPrototype bool `json:"isPrototype"`
	LastUpdated time.Time `json:"lastUpdated"`
	LastUpdatedBy string `json:"lastUpdatedBy"`
	Name string `json:"name"`
	ParentJaskId *string `json:"parentJaskId,omitempty"`
	Status GetRulesResponseDataObjectsInnerOneOfStatus `json:"status"`
	RuleId int32 `json:"ruleId"`
	RuleSource string `json:"ruleSource"`
	RuleType string `json:"ruleType"`
	// The number of Signals generated by this Rule in the past 7 days
	SignalCount07d int32 `json:"signalCount07d"`
	// The number of Signals generated by this Rule in the past 24 hours
	SignalCount24h int32 `json:"signalCount24h"`
	SummaryExpression string `json:"summaryExpression"`
	Tags []string `json:"tags"`
	HasOverride bool `json:"hasOverride"`
	TagsOverride *GetRulesResponseDataObjectsInnerOneOfTagsOverride `json:"tagsOverride,omitempty"`
	IsPrototypeOverride *GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride `json:"isPrototypeOverride,omitempty"`
	EntitySelectorsOverride *GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride `json:"entitySelectorsOverride,omitempty"`
	NameOverride *GetRulesResponseDataObjectsInnerOneOfNameOverride `json:"nameOverride,omitempty"`
	SummaryExpressionOverride *GetRulesResponseDataObjectsInnerOneOfNameOverride `json:"summaryExpressionOverride,omitempty"`
	AssetField string `json:"assetField"`
	// The description to be used on the generated Signal
	Description string `json:"description"`
	Expression string `json:"expression"`
	Score int32 `json:"score"`
	Stream string `json:"stream"`
}

type _MatchRule MatchRule

// NewMatchRule instantiates a new MatchRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchRule(contentType string, created time.Time, createdBy string, deleted bool, enabled bool, entitySelectors []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner, id string, isPrototype bool, lastUpdated time.Time, lastUpdatedBy string, name string, status GetRulesResponseDataObjectsInnerOneOfStatus, ruleId int32, ruleSource string, ruleType string, signalCount07d int32, signalCount24h int32, summaryExpression string, tags []string, hasOverride bool, assetField string, description string, expression string, score int32, stream string) *MatchRule {
	this := MatchRule{}
	this.ContentType = contentType
	this.Created = created
	this.CreatedBy = createdBy
	this.Deleted = deleted
	this.Enabled = enabled
	this.EntitySelectors = entitySelectors
	this.Id = id
	this.IsPrototype = isPrototype
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Name = name
	this.Status = status
	this.RuleId = ruleId
	this.RuleSource = ruleSource
	this.RuleType = ruleType
	this.SignalCount07d = signalCount07d
	this.SignalCount24h = signalCount24h
	this.SummaryExpression = summaryExpression
	this.Tags = tags
	this.HasOverride = hasOverride
	this.AssetField = assetField
	this.Description = description
	this.Expression = expression
	this.Score = score
	this.Stream = stream
	return &this
}

// NewMatchRuleWithDefaults instantiates a new MatchRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchRuleWithDefaults() *MatchRule {
	this := MatchRule{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *MatchRule) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchRule) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *MatchRule) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *MatchRule) SetCategory(v string) {
	o.Category = &v
}

// GetContentType returns the ContentType field value
func (o *MatchRule) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *MatchRule) SetContentType(v string) {
	o.ContentType = v
}

// GetCreated returns the Created field value
func (o *MatchRule) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *MatchRule) SetCreated(v time.Time) {
	o.Created = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *MatchRule) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *MatchRule) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetDeleted returns the Deleted field value
func (o *MatchRule) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *MatchRule) SetDeleted(v bool) {
	o.Deleted = v
}

// GetEnabled returns the Enabled field value
func (o *MatchRule) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *MatchRule) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEntitySelectors returns the EntitySelectors field value
func (o *MatchRule) GetEntitySelectors() []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner {
	if o == nil {
		var ret []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner
		return ret
	}

	return o.EntitySelectors
}

// GetEntitySelectorsOk returns a tuple with the EntitySelectors field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetEntitySelectorsOk() ([]GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntitySelectors, true
}

// SetEntitySelectors sets field value
func (o *MatchRule) SetEntitySelectors(v []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner) {
	o.EntitySelectors = v
}

// GetId returns the Id field value
func (o *MatchRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MatchRule) SetId(v string) {
	o.Id = v
}

// GetIsPrototype returns the IsPrototype field value
func (o *MatchRule) GetIsPrototype() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPrototype
}

// GetIsPrototypeOk returns a tuple with the IsPrototype field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetIsPrototypeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrototype, true
}

// SetIsPrototype sets field value
func (o *MatchRule) SetIsPrototype(v bool) {
	o.IsPrototype = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *MatchRule) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *MatchRule) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *MatchRule) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *MatchRule) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetName returns the Name field value
func (o *MatchRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MatchRule) SetName(v string) {
	o.Name = v
}

// GetParentJaskId returns the ParentJaskId field value if set, zero value otherwise.
func (o *MatchRule) GetParentJaskId() string {
	if o == nil || IsNil(o.ParentJaskId) {
		var ret string
		return ret
	}
	return *o.ParentJaskId
}

// GetParentJaskIdOk returns a tuple with the ParentJaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchRule) GetParentJaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentJaskId) {
		return nil, false
	}
	return o.ParentJaskId, true
}

// HasParentJaskId returns a boolean if a field has been set.
func (o *MatchRule) HasParentJaskId() bool {
	if o != nil && !IsNil(o.ParentJaskId) {
		return true
	}

	return false
}

// SetParentJaskId gets a reference to the given string and assigns it to the ParentJaskId field.
func (o *MatchRule) SetParentJaskId(v string) {
	o.ParentJaskId = &v
}

// GetStatus returns the Status field value
func (o *MatchRule) GetStatus() GetRulesResponseDataObjectsInnerOneOfStatus {
	if o == nil {
		var ret GetRulesResponseDataObjectsInnerOneOfStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetStatusOk() (*GetRulesResponseDataObjectsInnerOneOfStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *MatchRule) SetStatus(v GetRulesResponseDataObjectsInnerOneOfStatus) {
	o.Status = v
}

// GetRuleId returns the RuleId field value
func (o *MatchRule) GetRuleId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetRuleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value
func (o *MatchRule) SetRuleId(v int32) {
	o.RuleId = v
}

// GetRuleSource returns the RuleSource field value
func (o *MatchRule) GetRuleSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleSource
}

// GetRuleSourceOk returns a tuple with the RuleSource field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetRuleSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleSource, true
}

// SetRuleSource sets field value
func (o *MatchRule) SetRuleSource(v string) {
	o.RuleSource = v
}

// GetRuleType returns the RuleType field value
func (o *MatchRule) GetRuleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetRuleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleType, true
}

// SetRuleType sets field value
func (o *MatchRule) SetRuleType(v string) {
	o.RuleType = v
}

// GetSignalCount07d returns the SignalCount07d field value
func (o *MatchRule) GetSignalCount07d() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SignalCount07d
}

// GetSignalCount07dOk returns a tuple with the SignalCount07d field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetSignalCount07dOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalCount07d, true
}

// SetSignalCount07d sets field value
func (o *MatchRule) SetSignalCount07d(v int32) {
	o.SignalCount07d = v
}

// GetSignalCount24h returns the SignalCount24h field value
func (o *MatchRule) GetSignalCount24h() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SignalCount24h
}

// GetSignalCount24hOk returns a tuple with the SignalCount24h field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetSignalCount24hOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalCount24h, true
}

// SetSignalCount24h sets field value
func (o *MatchRule) SetSignalCount24h(v int32) {
	o.SignalCount24h = v
}

// GetSummaryExpression returns the SummaryExpression field value
func (o *MatchRule) GetSummaryExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SummaryExpression
}

// GetSummaryExpressionOk returns a tuple with the SummaryExpression field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetSummaryExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SummaryExpression, true
}

// SetSummaryExpression sets field value
func (o *MatchRule) SetSummaryExpression(v string) {
	o.SummaryExpression = v
}

// GetTags returns the Tags field value
func (o *MatchRule) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *MatchRule) SetTags(v []string) {
	o.Tags = v
}

// GetHasOverride returns the HasOverride field value
func (o *MatchRule) GetHasOverride() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasOverride
}

// GetHasOverrideOk returns a tuple with the HasOverride field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetHasOverrideOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasOverride, true
}

// SetHasOverride sets field value
func (o *MatchRule) SetHasOverride(v bool) {
	o.HasOverride = v
}

// GetTagsOverride returns the TagsOverride field value if set, zero value otherwise.
func (o *MatchRule) GetTagsOverride() GetRulesResponseDataObjectsInnerOneOfTagsOverride {
	if o == nil || IsNil(o.TagsOverride) {
		var ret GetRulesResponseDataObjectsInnerOneOfTagsOverride
		return ret
	}
	return *o.TagsOverride
}

// GetTagsOverrideOk returns a tuple with the TagsOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchRule) GetTagsOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfTagsOverride, bool) {
	if o == nil || IsNil(o.TagsOverride) {
		return nil, false
	}
	return o.TagsOverride, true
}

// HasTagsOverride returns a boolean if a field has been set.
func (o *MatchRule) HasTagsOverride() bool {
	if o != nil && !IsNil(o.TagsOverride) {
		return true
	}

	return false
}

// SetTagsOverride gets a reference to the given GetRulesResponseDataObjectsInnerOneOfTagsOverride and assigns it to the TagsOverride field.
func (o *MatchRule) SetTagsOverride(v GetRulesResponseDataObjectsInnerOneOfTagsOverride) {
	o.TagsOverride = &v
}

// GetIsPrototypeOverride returns the IsPrototypeOverride field value if set, zero value otherwise.
func (o *MatchRule) GetIsPrototypeOverride() GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride {
	if o == nil || IsNil(o.IsPrototypeOverride) {
		var ret GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride
		return ret
	}
	return *o.IsPrototypeOverride
}

// GetIsPrototypeOverrideOk returns a tuple with the IsPrototypeOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchRule) GetIsPrototypeOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride, bool) {
	if o == nil || IsNil(o.IsPrototypeOverride) {
		return nil, false
	}
	return o.IsPrototypeOverride, true
}

// HasIsPrototypeOverride returns a boolean if a field has been set.
func (o *MatchRule) HasIsPrototypeOverride() bool {
	if o != nil && !IsNil(o.IsPrototypeOverride) {
		return true
	}

	return false
}

// SetIsPrototypeOverride gets a reference to the given GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride and assigns it to the IsPrototypeOverride field.
func (o *MatchRule) SetIsPrototypeOverride(v GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride) {
	o.IsPrototypeOverride = &v
}

// GetEntitySelectorsOverride returns the EntitySelectorsOverride field value if set, zero value otherwise.
func (o *MatchRule) GetEntitySelectorsOverride() GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride {
	if o == nil || IsNil(o.EntitySelectorsOverride) {
		var ret GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride
		return ret
	}
	return *o.EntitySelectorsOverride
}

// GetEntitySelectorsOverrideOk returns a tuple with the EntitySelectorsOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchRule) GetEntitySelectorsOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride, bool) {
	if o == nil || IsNil(o.EntitySelectorsOverride) {
		return nil, false
	}
	return o.EntitySelectorsOverride, true
}

// HasEntitySelectorsOverride returns a boolean if a field has been set.
func (o *MatchRule) HasEntitySelectorsOverride() bool {
	if o != nil && !IsNil(o.EntitySelectorsOverride) {
		return true
	}

	return false
}

// SetEntitySelectorsOverride gets a reference to the given GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride and assigns it to the EntitySelectorsOverride field.
func (o *MatchRule) SetEntitySelectorsOverride(v GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) {
	o.EntitySelectorsOverride = &v
}

// GetNameOverride returns the NameOverride field value if set, zero value otherwise.
func (o *MatchRule) GetNameOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride {
	if o == nil || IsNil(o.NameOverride) {
		var ret GetRulesResponseDataObjectsInnerOneOfNameOverride
		return ret
	}
	return *o.NameOverride
}

// GetNameOverrideOk returns a tuple with the NameOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchRule) GetNameOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool) {
	if o == nil || IsNil(o.NameOverride) {
		return nil, false
	}
	return o.NameOverride, true
}

// HasNameOverride returns a boolean if a field has been set.
func (o *MatchRule) HasNameOverride() bool {
	if o != nil && !IsNil(o.NameOverride) {
		return true
	}

	return false
}

// SetNameOverride gets a reference to the given GetRulesResponseDataObjectsInnerOneOfNameOverride and assigns it to the NameOverride field.
func (o *MatchRule) SetNameOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride) {
	o.NameOverride = &v
}

// GetSummaryExpressionOverride returns the SummaryExpressionOverride field value if set, zero value otherwise.
func (o *MatchRule) GetSummaryExpressionOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride {
	if o == nil || IsNil(o.SummaryExpressionOverride) {
		var ret GetRulesResponseDataObjectsInnerOneOfNameOverride
		return ret
	}
	return *o.SummaryExpressionOverride
}

// GetSummaryExpressionOverrideOk returns a tuple with the SummaryExpressionOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchRule) GetSummaryExpressionOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool) {
	if o == nil || IsNil(o.SummaryExpressionOverride) {
		return nil, false
	}
	return o.SummaryExpressionOverride, true
}

// HasSummaryExpressionOverride returns a boolean if a field has been set.
func (o *MatchRule) HasSummaryExpressionOverride() bool {
	if o != nil && !IsNil(o.SummaryExpressionOverride) {
		return true
	}

	return false
}

// SetSummaryExpressionOverride gets a reference to the given GetRulesResponseDataObjectsInnerOneOfNameOverride and assigns it to the SummaryExpressionOverride field.
func (o *MatchRule) SetSummaryExpressionOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride) {
	o.SummaryExpressionOverride = &v
}

// GetAssetField returns the AssetField field value
func (o *MatchRule) GetAssetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetField
}

// GetAssetFieldOk returns a tuple with the AssetField field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetAssetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetField, true
}

// SetAssetField sets field value
func (o *MatchRule) SetAssetField(v string) {
	o.AssetField = v
}

// GetDescription returns the Description field value
func (o *MatchRule) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *MatchRule) SetDescription(v string) {
	o.Description = v
}

// GetExpression returns the Expression field value
func (o *MatchRule) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *MatchRule) SetExpression(v string) {
	o.Expression = v
}

// GetScore returns the Score field value
func (o *MatchRule) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *MatchRule) SetScore(v int32) {
	o.Score = v
}

// GetStream returns the Stream field value
func (o *MatchRule) GetStream() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stream
}

// GetStreamOk returns a tuple with the Stream field value
// and a boolean to check if the value has been set.
func (o *MatchRule) GetStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stream, true
}

// SetStream sets field value
func (o *MatchRule) SetStream(v string) {
	o.Stream = v
}

func (o MatchRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatchRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	toSerialize["contentType"] = o.ContentType
	toSerialize["created"] = o.Created
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["deleted"] = o.Deleted
	toSerialize["enabled"] = o.Enabled
	toSerialize["entitySelectors"] = o.EntitySelectors
	toSerialize["id"] = o.Id
	toSerialize["isPrototype"] = o.IsPrototype
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	toSerialize["name"] = o.Name
	if !IsNil(o.ParentJaskId) {
		toSerialize["parentJaskId"] = o.ParentJaskId
	}
	toSerialize["status"] = o.Status
	toSerialize["ruleId"] = o.RuleId
	toSerialize["ruleSource"] = o.RuleSource
	toSerialize["ruleType"] = o.RuleType
	toSerialize["signalCount07d"] = o.SignalCount07d
	toSerialize["signalCount24h"] = o.SignalCount24h
	toSerialize["summaryExpression"] = o.SummaryExpression
	toSerialize["tags"] = o.Tags
	toSerialize["hasOverride"] = o.HasOverride
	if !IsNil(o.TagsOverride) {
		toSerialize["tagsOverride"] = o.TagsOverride
	}
	if !IsNil(o.IsPrototypeOverride) {
		toSerialize["isPrototypeOverride"] = o.IsPrototypeOverride
	}
	if !IsNil(o.EntitySelectorsOverride) {
		toSerialize["entitySelectorsOverride"] = o.EntitySelectorsOverride
	}
	if !IsNil(o.NameOverride) {
		toSerialize["nameOverride"] = o.NameOverride
	}
	if !IsNil(o.SummaryExpressionOverride) {
		toSerialize["summaryExpressionOverride"] = o.SummaryExpressionOverride
	}
	toSerialize["assetField"] = o.AssetField
	toSerialize["description"] = o.Description
	toSerialize["expression"] = o.Expression
	toSerialize["score"] = o.Score
	toSerialize["stream"] = o.Stream
	return toSerialize, nil
}

func (o *MatchRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contentType",
		"created",
		"createdBy",
		"deleted",
		"enabled",
		"entitySelectors",
		"id",
		"isPrototype",
		"lastUpdated",
		"lastUpdatedBy",
		"name",
		"status",
		"ruleId",
		"ruleSource",
		"ruleType",
		"signalCount07d",
		"signalCount24h",
		"summaryExpression",
		"tags",
		"hasOverride",
		"assetField",
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

	varMatchRule := _MatchRule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMatchRule)

	if err != nil {
		return err
	}

	*o = MatchRule(varMatchRule)

	return err
}

type NullableMatchRule struct {
	value *MatchRule
	isSet bool
}

func (v NullableMatchRule) Get() *MatchRule {
	return v.value
}

func (v *NullableMatchRule) Set(val *MatchRule) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchRule) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchRule(val *MatchRule) *NullableMatchRule {
	return &NullableMatchRule{value: val, isSet: true}
}

func (v NullableMatchRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


