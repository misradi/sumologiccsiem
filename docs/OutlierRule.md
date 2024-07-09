# OutlierRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**ContentType** | **string** |  | 
**Created** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**Deleted** | **bool** |  | 
**Enabled** | **bool** |  | 
**EntitySelectors** | [**[]GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner**](GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner.md) |  | 
**Id** | **string** |  | 
**IsPrototype** | **bool** |  | 
**LastUpdated** | **time.Time** |  | 
**LastUpdatedBy** | **string** |  | 
**Name** | **string** |  | 
**ParentJaskId** | Pointer to **string** |  | [optional] 
**Status** | [**GetRulesResponseDataObjectsInnerOneOfStatus**](GetRulesResponseDataObjectsInnerOneOfStatus.md) |  | 
**RuleId** | **int32** |  | 
**RuleSource** | **string** |  | 
**RuleType** | **string** |  | 
**SignalCount07d** | **int32** | The number of Signals generated by this Rule in the past 7 days | 
**SignalCount24h** | **int32** | The number of Signals generated by this Rule in the past 24 hours | 
**SummaryExpression** | **string** |  | 
**Tags** | **[]string** |  | 
**HasOverride** | **bool** |  | 
**TagsOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfTagsOverride**](GetRulesResponseDataObjectsInnerOneOfTagsOverride.md) |  | [optional] 
**IsPrototypeOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride**](GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride.md) |  | [optional] 
**EntitySelectorsOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride**](GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride.md) |  | [optional] 
**NameOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfNameOverride**](GetRulesResponseDataObjectsInnerOneOfNameOverride.md) |  | [optional] 
**SummaryExpressionOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfNameOverride**](GetRulesResponseDataObjectsInnerOneOfNameOverride.md) |  | [optional] 
**DescriptionExpression** | **string** | The description to be used on the generated Signal | 
**NameExpression** | **string** |  | 
**MatchExpression** | **string** |  | 
**GroupByFields** | **[]string** |  | 
**WindowSize** | **int32** | milliseconds | 
**WindowSizeName** | **string** |  | 
**Score** | **int32** |  | 
**Version** | **int32** |  | 
**BaselineWindowSize** | **string** | milliseconds | 
**RetentionWindowSize** | **string** | milliseconds | 
**BaselineLastStarted** | Pointer to **time.Time** |  | [optional] 
**FloorValue** | **int32** |  | 
**DeviationThreshold** | **int32** |  | 
**AggregationFunctions** | [**[]GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner**](GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner.md) |  | 
**BaselineWindowSizeOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride**](GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride.md) |  | [optional] 
**DescriptionExpressionOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfNameOverride**](GetRulesResponseDataObjectsInnerOneOfNameOverride.md) |  | [optional] 
**GroupByFieldsOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfTagsOverride**](GetRulesResponseDataObjectsInnerOneOfTagsOverride.md) |  | [optional] 
**NameExpressionOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfNameOverride**](GetRulesResponseDataObjectsInnerOneOfNameOverride.md) |  | [optional] 
**RetentionWindowSizeOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride**](GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride.md) |  | [optional] 
**FloorValueOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOf1ScoreOverride**](GetRulesResponseDataObjectsInnerOneOf1ScoreOverride.md) |  | [optional] 
**DeviationThresholdOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOf1ScoreOverride**](GetRulesResponseDataObjectsInnerOneOf1ScoreOverride.md) |  | [optional] 
**ScoreOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOf1ScoreOverride**](GetRulesResponseDataObjectsInnerOneOf1ScoreOverride.md) |  | [optional] 
**WindowSizeOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOf1WindowSizeOverride**](GetRulesResponseDataObjectsInnerOneOf1WindowSizeOverride.md) |  | [optional] 

## Methods

### NewOutlierRule

`func NewOutlierRule(contentType string, created time.Time, createdBy string, deleted bool, enabled bool, entitySelectors []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner, id string, isPrototype bool, lastUpdated time.Time, lastUpdatedBy string, name string, status GetRulesResponseDataObjectsInnerOneOfStatus, ruleId int32, ruleSource string, ruleType string, signalCount07d int32, signalCount24h int32, summaryExpression string, tags []string, hasOverride bool, descriptionExpression string, nameExpression string, matchExpression string, groupByFields []string, windowSize int32, windowSizeName string, score int32, version int32, baselineWindowSize string, retentionWindowSize string, floorValue int32, deviationThreshold int32, aggregationFunctions []GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner, ) *OutlierRule`

NewOutlierRule instantiates a new OutlierRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutlierRuleWithDefaults

`func NewOutlierRuleWithDefaults() *OutlierRule`

NewOutlierRuleWithDefaults instantiates a new OutlierRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *OutlierRule) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OutlierRule) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OutlierRule) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *OutlierRule) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetContentType

`func (o *OutlierRule) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *OutlierRule) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *OutlierRule) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetCreated

`func (o *OutlierRule) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OutlierRule) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OutlierRule) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCreatedBy

`func (o *OutlierRule) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OutlierRule) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OutlierRule) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDeleted

`func (o *OutlierRule) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *OutlierRule) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *OutlierRule) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetEnabled

`func (o *OutlierRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OutlierRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OutlierRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEntitySelectors

`func (o *OutlierRule) GetEntitySelectors() []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner`

GetEntitySelectors returns the EntitySelectors field if non-nil, zero value otherwise.

### GetEntitySelectorsOk

`func (o *OutlierRule) GetEntitySelectorsOk() (*[]GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner, bool)`

GetEntitySelectorsOk returns a tuple with the EntitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectors

`func (o *OutlierRule) SetEntitySelectors(v []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner)`

SetEntitySelectors sets EntitySelectors field to given value.


### GetId

`func (o *OutlierRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutlierRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutlierRule) SetId(v string)`

SetId sets Id field to given value.


### GetIsPrototype

`func (o *OutlierRule) GetIsPrototype() bool`

GetIsPrototype returns the IsPrototype field if non-nil, zero value otherwise.

### GetIsPrototypeOk

`func (o *OutlierRule) GetIsPrototypeOk() (*bool, bool)`

GetIsPrototypeOk returns a tuple with the IsPrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototype

`func (o *OutlierRule) SetIsPrototype(v bool)`

SetIsPrototype sets IsPrototype field to given value.


### GetLastUpdated

`func (o *OutlierRule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *OutlierRule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *OutlierRule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *OutlierRule) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *OutlierRule) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *OutlierRule) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetName

`func (o *OutlierRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutlierRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutlierRule) SetName(v string)`

SetName sets Name field to given value.


### GetParentJaskId

`func (o *OutlierRule) GetParentJaskId() string`

GetParentJaskId returns the ParentJaskId field if non-nil, zero value otherwise.

### GetParentJaskIdOk

`func (o *OutlierRule) GetParentJaskIdOk() (*string, bool)`

GetParentJaskIdOk returns a tuple with the ParentJaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJaskId

`func (o *OutlierRule) SetParentJaskId(v string)`

SetParentJaskId sets ParentJaskId field to given value.

### HasParentJaskId

`func (o *OutlierRule) HasParentJaskId() bool`

HasParentJaskId returns a boolean if a field has been set.

### GetStatus

`func (o *OutlierRule) GetStatus() GetRulesResponseDataObjectsInnerOneOfStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OutlierRule) GetStatusOk() (*GetRulesResponseDataObjectsInnerOneOfStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OutlierRule) SetStatus(v GetRulesResponseDataObjectsInnerOneOfStatus)`

SetStatus sets Status field to given value.


### GetRuleId

`func (o *OutlierRule) GetRuleId() int32`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *OutlierRule) GetRuleIdOk() (*int32, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *OutlierRule) SetRuleId(v int32)`

SetRuleId sets RuleId field to given value.


### GetRuleSource

`func (o *OutlierRule) GetRuleSource() string`

GetRuleSource returns the RuleSource field if non-nil, zero value otherwise.

### GetRuleSourceOk

`func (o *OutlierRule) GetRuleSourceOk() (*string, bool)`

GetRuleSourceOk returns a tuple with the RuleSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSource

`func (o *OutlierRule) SetRuleSource(v string)`

SetRuleSource sets RuleSource field to given value.


### GetRuleType

`func (o *OutlierRule) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *OutlierRule) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *OutlierRule) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetSignalCount07d

`func (o *OutlierRule) GetSignalCount07d() int32`

GetSignalCount07d returns the SignalCount07d field if non-nil, zero value otherwise.

### GetSignalCount07dOk

`func (o *OutlierRule) GetSignalCount07dOk() (*int32, bool)`

GetSignalCount07dOk returns a tuple with the SignalCount07d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalCount07d

`func (o *OutlierRule) SetSignalCount07d(v int32)`

SetSignalCount07d sets SignalCount07d field to given value.


### GetSignalCount24h

`func (o *OutlierRule) GetSignalCount24h() int32`

GetSignalCount24h returns the SignalCount24h field if non-nil, zero value otherwise.

### GetSignalCount24hOk

`func (o *OutlierRule) GetSignalCount24hOk() (*int32, bool)`

GetSignalCount24hOk returns a tuple with the SignalCount24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalCount24h

`func (o *OutlierRule) SetSignalCount24h(v int32)`

SetSignalCount24h sets SignalCount24h field to given value.


### GetSummaryExpression

`func (o *OutlierRule) GetSummaryExpression() string`

GetSummaryExpression returns the SummaryExpression field if non-nil, zero value otherwise.

### GetSummaryExpressionOk

`func (o *OutlierRule) GetSummaryExpressionOk() (*string, bool)`

GetSummaryExpressionOk returns a tuple with the SummaryExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpression

`func (o *OutlierRule) SetSummaryExpression(v string)`

SetSummaryExpression sets SummaryExpression field to given value.


### GetTags

`func (o *OutlierRule) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OutlierRule) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OutlierRule) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetHasOverride

`func (o *OutlierRule) GetHasOverride() bool`

GetHasOverride returns the HasOverride field if non-nil, zero value otherwise.

### GetHasOverrideOk

`func (o *OutlierRule) GetHasOverrideOk() (*bool, bool)`

GetHasOverrideOk returns a tuple with the HasOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOverride

`func (o *OutlierRule) SetHasOverride(v bool)`

SetHasOverride sets HasOverride field to given value.


### GetTagsOverride

`func (o *OutlierRule) GetTagsOverride() GetRulesResponseDataObjectsInnerOneOfTagsOverride`

GetTagsOverride returns the TagsOverride field if non-nil, zero value otherwise.

### GetTagsOverrideOk

`func (o *OutlierRule) GetTagsOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfTagsOverride, bool)`

GetTagsOverrideOk returns a tuple with the TagsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsOverride

`func (o *OutlierRule) SetTagsOverride(v GetRulesResponseDataObjectsInnerOneOfTagsOverride)`

SetTagsOverride sets TagsOverride field to given value.

### HasTagsOverride

`func (o *OutlierRule) HasTagsOverride() bool`

HasTagsOverride returns a boolean if a field has been set.

### GetIsPrototypeOverride

`func (o *OutlierRule) GetIsPrototypeOverride() GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride`

GetIsPrototypeOverride returns the IsPrototypeOverride field if non-nil, zero value otherwise.

### GetIsPrototypeOverrideOk

`func (o *OutlierRule) GetIsPrototypeOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride, bool)`

GetIsPrototypeOverrideOk returns a tuple with the IsPrototypeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototypeOverride

`func (o *OutlierRule) SetIsPrototypeOverride(v GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride)`

SetIsPrototypeOverride sets IsPrototypeOverride field to given value.

### HasIsPrototypeOverride

`func (o *OutlierRule) HasIsPrototypeOverride() bool`

HasIsPrototypeOverride returns a boolean if a field has been set.

### GetEntitySelectorsOverride

`func (o *OutlierRule) GetEntitySelectorsOverride() GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride`

GetEntitySelectorsOverride returns the EntitySelectorsOverride field if non-nil, zero value otherwise.

### GetEntitySelectorsOverrideOk

`func (o *OutlierRule) GetEntitySelectorsOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride, bool)`

GetEntitySelectorsOverrideOk returns a tuple with the EntitySelectorsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectorsOverride

`func (o *OutlierRule) SetEntitySelectorsOverride(v GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride)`

SetEntitySelectorsOverride sets EntitySelectorsOverride field to given value.

### HasEntitySelectorsOverride

`func (o *OutlierRule) HasEntitySelectorsOverride() bool`

HasEntitySelectorsOverride returns a boolean if a field has been set.

### GetNameOverride

`func (o *OutlierRule) GetNameOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride`

GetNameOverride returns the NameOverride field if non-nil, zero value otherwise.

### GetNameOverrideOk

`func (o *OutlierRule) GetNameOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool)`

GetNameOverrideOk returns a tuple with the NameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOverride

`func (o *OutlierRule) SetNameOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride)`

SetNameOverride sets NameOverride field to given value.

### HasNameOverride

`func (o *OutlierRule) HasNameOverride() bool`

HasNameOverride returns a boolean if a field has been set.

### GetSummaryExpressionOverride

`func (o *OutlierRule) GetSummaryExpressionOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride`

GetSummaryExpressionOverride returns the SummaryExpressionOverride field if non-nil, zero value otherwise.

### GetSummaryExpressionOverrideOk

`func (o *OutlierRule) GetSummaryExpressionOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool)`

GetSummaryExpressionOverrideOk returns a tuple with the SummaryExpressionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpressionOverride

`func (o *OutlierRule) SetSummaryExpressionOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride)`

SetSummaryExpressionOverride sets SummaryExpressionOverride field to given value.

### HasSummaryExpressionOverride

`func (o *OutlierRule) HasSummaryExpressionOverride() bool`

HasSummaryExpressionOverride returns a boolean if a field has been set.

### GetDescriptionExpression

`func (o *OutlierRule) GetDescriptionExpression() string`

GetDescriptionExpression returns the DescriptionExpression field if non-nil, zero value otherwise.

### GetDescriptionExpressionOk

`func (o *OutlierRule) GetDescriptionExpressionOk() (*string, bool)`

GetDescriptionExpressionOk returns a tuple with the DescriptionExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionExpression

`func (o *OutlierRule) SetDescriptionExpression(v string)`

SetDescriptionExpression sets DescriptionExpression field to given value.


### GetNameExpression

`func (o *OutlierRule) GetNameExpression() string`

GetNameExpression returns the NameExpression field if non-nil, zero value otherwise.

### GetNameExpressionOk

`func (o *OutlierRule) GetNameExpressionOk() (*string, bool)`

GetNameExpressionOk returns a tuple with the NameExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExpression

`func (o *OutlierRule) SetNameExpression(v string)`

SetNameExpression sets NameExpression field to given value.


### GetMatchExpression

`func (o *OutlierRule) GetMatchExpression() string`

GetMatchExpression returns the MatchExpression field if non-nil, zero value otherwise.

### GetMatchExpressionOk

`func (o *OutlierRule) GetMatchExpressionOk() (*string, bool)`

GetMatchExpressionOk returns a tuple with the MatchExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchExpression

`func (o *OutlierRule) SetMatchExpression(v string)`

SetMatchExpression sets MatchExpression field to given value.


### GetGroupByFields

`func (o *OutlierRule) GetGroupByFields() []string`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *OutlierRule) GetGroupByFieldsOk() (*[]string, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *OutlierRule) SetGroupByFields(v []string)`

SetGroupByFields sets GroupByFields field to given value.


### GetWindowSize

`func (o *OutlierRule) GetWindowSize() int32`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *OutlierRule) GetWindowSizeOk() (*int32, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *OutlierRule) SetWindowSize(v int32)`

SetWindowSize sets WindowSize field to given value.


### GetWindowSizeName

`func (o *OutlierRule) GetWindowSizeName() string`

GetWindowSizeName returns the WindowSizeName field if non-nil, zero value otherwise.

### GetWindowSizeNameOk

`func (o *OutlierRule) GetWindowSizeNameOk() (*string, bool)`

GetWindowSizeNameOk returns a tuple with the WindowSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSizeName

`func (o *OutlierRule) SetWindowSizeName(v string)`

SetWindowSizeName sets WindowSizeName field to given value.


### GetScore

`func (o *OutlierRule) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *OutlierRule) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *OutlierRule) SetScore(v int32)`

SetScore sets Score field to given value.


### GetVersion

`func (o *OutlierRule) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OutlierRule) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OutlierRule) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetBaselineWindowSize

`func (o *OutlierRule) GetBaselineWindowSize() string`

GetBaselineWindowSize returns the BaselineWindowSize field if non-nil, zero value otherwise.

### GetBaselineWindowSizeOk

`func (o *OutlierRule) GetBaselineWindowSizeOk() (*string, bool)`

GetBaselineWindowSizeOk returns a tuple with the BaselineWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineWindowSize

`func (o *OutlierRule) SetBaselineWindowSize(v string)`

SetBaselineWindowSize sets BaselineWindowSize field to given value.


### GetRetentionWindowSize

`func (o *OutlierRule) GetRetentionWindowSize() string`

GetRetentionWindowSize returns the RetentionWindowSize field if non-nil, zero value otherwise.

### GetRetentionWindowSizeOk

`func (o *OutlierRule) GetRetentionWindowSizeOk() (*string, bool)`

GetRetentionWindowSizeOk returns a tuple with the RetentionWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionWindowSize

`func (o *OutlierRule) SetRetentionWindowSize(v string)`

SetRetentionWindowSize sets RetentionWindowSize field to given value.


### GetBaselineLastStarted

`func (o *OutlierRule) GetBaselineLastStarted() time.Time`

GetBaselineLastStarted returns the BaselineLastStarted field if non-nil, zero value otherwise.

### GetBaselineLastStartedOk

`func (o *OutlierRule) GetBaselineLastStartedOk() (*time.Time, bool)`

GetBaselineLastStartedOk returns a tuple with the BaselineLastStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineLastStarted

`func (o *OutlierRule) SetBaselineLastStarted(v time.Time)`

SetBaselineLastStarted sets BaselineLastStarted field to given value.

### HasBaselineLastStarted

`func (o *OutlierRule) HasBaselineLastStarted() bool`

HasBaselineLastStarted returns a boolean if a field has been set.

### GetFloorValue

`func (o *OutlierRule) GetFloorValue() int32`

GetFloorValue returns the FloorValue field if non-nil, zero value otherwise.

### GetFloorValueOk

`func (o *OutlierRule) GetFloorValueOk() (*int32, bool)`

GetFloorValueOk returns a tuple with the FloorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorValue

`func (o *OutlierRule) SetFloorValue(v int32)`

SetFloorValue sets FloorValue field to given value.


### GetDeviationThreshold

`func (o *OutlierRule) GetDeviationThreshold() int32`

GetDeviationThreshold returns the DeviationThreshold field if non-nil, zero value otherwise.

### GetDeviationThresholdOk

`func (o *OutlierRule) GetDeviationThresholdOk() (*int32, bool)`

GetDeviationThresholdOk returns a tuple with the DeviationThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviationThreshold

`func (o *OutlierRule) SetDeviationThreshold(v int32)`

SetDeviationThreshold sets DeviationThreshold field to given value.


### GetAggregationFunctions

`func (o *OutlierRule) GetAggregationFunctions() []GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner`

GetAggregationFunctions returns the AggregationFunctions field if non-nil, zero value otherwise.

### GetAggregationFunctionsOk

`func (o *OutlierRule) GetAggregationFunctionsOk() (*[]GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner, bool)`

GetAggregationFunctionsOk returns a tuple with the AggregationFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationFunctions

`func (o *OutlierRule) SetAggregationFunctions(v []GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner)`

SetAggregationFunctions sets AggregationFunctions field to given value.


### GetBaselineWindowSizeOverride

`func (o *OutlierRule) GetBaselineWindowSizeOverride() GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride`

GetBaselineWindowSizeOverride returns the BaselineWindowSizeOverride field if non-nil, zero value otherwise.

### GetBaselineWindowSizeOverrideOk

`func (o *OutlierRule) GetBaselineWindowSizeOverrideOk() (*GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride, bool)`

GetBaselineWindowSizeOverrideOk returns a tuple with the BaselineWindowSizeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineWindowSizeOverride

`func (o *OutlierRule) SetBaselineWindowSizeOverride(v GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride)`

SetBaselineWindowSizeOverride sets BaselineWindowSizeOverride field to given value.

### HasBaselineWindowSizeOverride

`func (o *OutlierRule) HasBaselineWindowSizeOverride() bool`

HasBaselineWindowSizeOverride returns a boolean if a field has been set.

### GetDescriptionExpressionOverride

`func (o *OutlierRule) GetDescriptionExpressionOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride`

GetDescriptionExpressionOverride returns the DescriptionExpressionOverride field if non-nil, zero value otherwise.

### GetDescriptionExpressionOverrideOk

`func (o *OutlierRule) GetDescriptionExpressionOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool)`

GetDescriptionExpressionOverrideOk returns a tuple with the DescriptionExpressionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionExpressionOverride

`func (o *OutlierRule) SetDescriptionExpressionOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride)`

SetDescriptionExpressionOverride sets DescriptionExpressionOverride field to given value.

### HasDescriptionExpressionOverride

`func (o *OutlierRule) HasDescriptionExpressionOverride() bool`

HasDescriptionExpressionOverride returns a boolean if a field has been set.

### GetGroupByFieldsOverride

`func (o *OutlierRule) GetGroupByFieldsOverride() GetRulesResponseDataObjectsInnerOneOfTagsOverride`

GetGroupByFieldsOverride returns the GroupByFieldsOverride field if non-nil, zero value otherwise.

### GetGroupByFieldsOverrideOk

`func (o *OutlierRule) GetGroupByFieldsOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfTagsOverride, bool)`

GetGroupByFieldsOverrideOk returns a tuple with the GroupByFieldsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFieldsOverride

`func (o *OutlierRule) SetGroupByFieldsOverride(v GetRulesResponseDataObjectsInnerOneOfTagsOverride)`

SetGroupByFieldsOverride sets GroupByFieldsOverride field to given value.

### HasGroupByFieldsOverride

`func (o *OutlierRule) HasGroupByFieldsOverride() bool`

HasGroupByFieldsOverride returns a boolean if a field has been set.

### GetNameExpressionOverride

`func (o *OutlierRule) GetNameExpressionOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride`

GetNameExpressionOverride returns the NameExpressionOverride field if non-nil, zero value otherwise.

### GetNameExpressionOverrideOk

`func (o *OutlierRule) GetNameExpressionOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool)`

GetNameExpressionOverrideOk returns a tuple with the NameExpressionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExpressionOverride

`func (o *OutlierRule) SetNameExpressionOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride)`

SetNameExpressionOverride sets NameExpressionOverride field to given value.

### HasNameExpressionOverride

`func (o *OutlierRule) HasNameExpressionOverride() bool`

HasNameExpressionOverride returns a boolean if a field has been set.

### GetRetentionWindowSizeOverride

`func (o *OutlierRule) GetRetentionWindowSizeOverride() GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride`

GetRetentionWindowSizeOverride returns the RetentionWindowSizeOverride field if non-nil, zero value otherwise.

### GetRetentionWindowSizeOverrideOk

`func (o *OutlierRule) GetRetentionWindowSizeOverrideOk() (*GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride, bool)`

GetRetentionWindowSizeOverrideOk returns a tuple with the RetentionWindowSizeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionWindowSizeOverride

`func (o *OutlierRule) SetRetentionWindowSizeOverride(v GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride)`

SetRetentionWindowSizeOverride sets RetentionWindowSizeOverride field to given value.

### HasRetentionWindowSizeOverride

`func (o *OutlierRule) HasRetentionWindowSizeOverride() bool`

HasRetentionWindowSizeOverride returns a boolean if a field has been set.

### GetFloorValueOverride

`func (o *OutlierRule) GetFloorValueOverride() GetRulesResponseDataObjectsInnerOneOf1ScoreOverride`

GetFloorValueOverride returns the FloorValueOverride field if non-nil, zero value otherwise.

### GetFloorValueOverrideOk

`func (o *OutlierRule) GetFloorValueOverrideOk() (*GetRulesResponseDataObjectsInnerOneOf1ScoreOverride, bool)`

GetFloorValueOverrideOk returns a tuple with the FloorValueOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorValueOverride

`func (o *OutlierRule) SetFloorValueOverride(v GetRulesResponseDataObjectsInnerOneOf1ScoreOverride)`

SetFloorValueOverride sets FloorValueOverride field to given value.

### HasFloorValueOverride

`func (o *OutlierRule) HasFloorValueOverride() bool`

HasFloorValueOverride returns a boolean if a field has been set.

### GetDeviationThresholdOverride

`func (o *OutlierRule) GetDeviationThresholdOverride() GetRulesResponseDataObjectsInnerOneOf1ScoreOverride`

GetDeviationThresholdOverride returns the DeviationThresholdOverride field if non-nil, zero value otherwise.

### GetDeviationThresholdOverrideOk

`func (o *OutlierRule) GetDeviationThresholdOverrideOk() (*GetRulesResponseDataObjectsInnerOneOf1ScoreOverride, bool)`

GetDeviationThresholdOverrideOk returns a tuple with the DeviationThresholdOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviationThresholdOverride

`func (o *OutlierRule) SetDeviationThresholdOverride(v GetRulesResponseDataObjectsInnerOneOf1ScoreOverride)`

SetDeviationThresholdOverride sets DeviationThresholdOverride field to given value.

### HasDeviationThresholdOverride

`func (o *OutlierRule) HasDeviationThresholdOverride() bool`

HasDeviationThresholdOverride returns a boolean if a field has been set.

### GetScoreOverride

`func (o *OutlierRule) GetScoreOverride() GetRulesResponseDataObjectsInnerOneOf1ScoreOverride`

GetScoreOverride returns the ScoreOverride field if non-nil, zero value otherwise.

### GetScoreOverrideOk

`func (o *OutlierRule) GetScoreOverrideOk() (*GetRulesResponseDataObjectsInnerOneOf1ScoreOverride, bool)`

GetScoreOverrideOk returns a tuple with the ScoreOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreOverride

`func (o *OutlierRule) SetScoreOverride(v GetRulesResponseDataObjectsInnerOneOf1ScoreOverride)`

SetScoreOverride sets ScoreOverride field to given value.

### HasScoreOverride

`func (o *OutlierRule) HasScoreOverride() bool`

HasScoreOverride returns a boolean if a field has been set.

### GetWindowSizeOverride

`func (o *OutlierRule) GetWindowSizeOverride() GetRulesResponseDataObjectsInnerOneOf1WindowSizeOverride`

GetWindowSizeOverride returns the WindowSizeOverride field if non-nil, zero value otherwise.

### GetWindowSizeOverrideOk

`func (o *OutlierRule) GetWindowSizeOverrideOk() (*GetRulesResponseDataObjectsInnerOneOf1WindowSizeOverride, bool)`

GetWindowSizeOverrideOk returns a tuple with the WindowSizeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSizeOverride

`func (o *OutlierRule) SetWindowSizeOverride(v GetRulesResponseDataObjectsInnerOneOf1WindowSizeOverride)`

SetWindowSizeOverride sets WindowSizeOverride field to given value.

### HasWindowSizeOverride

`func (o *OutlierRule) HasWindowSizeOverride() bool`

HasWindowSizeOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

