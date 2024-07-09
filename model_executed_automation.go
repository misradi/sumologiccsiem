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

// checks if the ExecutedAutomation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutedAutomation{}

// ExecutedAutomation struct for ExecutedAutomation
type ExecutedAutomation struct {
	ExecutionId string `json:"executionId"`
	CseResourceType string `json:"cseResourceType"`
	CseResourceId string `json:"cseResourceId"`
	Status string `json:"status"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
	PlaybookId string `json:"playbookId"`
	PlaybookName string `json:"playbookName"`
	PlaybookDescription *string `json:"playbookDescription,omitempty"`
}

type _ExecutedAutomation ExecutedAutomation

// NewExecutedAutomation instantiates a new ExecutedAutomation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutedAutomation(executionId string, cseResourceType string, cseResourceId string, status string, playbookId string, playbookName string) *ExecutedAutomation {
	this := ExecutedAutomation{}
	this.ExecutionId = executionId
	this.CseResourceType = cseResourceType
	this.CseResourceId = cseResourceId
	this.Status = status
	this.PlaybookId = playbookId
	this.PlaybookName = playbookName
	return &this
}

// NewExecutedAutomationWithDefaults instantiates a new ExecutedAutomation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutedAutomationWithDefaults() *ExecutedAutomation {
	this := ExecutedAutomation{}
	return &this
}

// GetExecutionId returns the ExecutionId field value
func (o *ExecutedAutomation) GetExecutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value
// and a boolean to check if the value has been set.
func (o *ExecutedAutomation) GetExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionId, true
}

// SetExecutionId sets field value
func (o *ExecutedAutomation) SetExecutionId(v string) {
	o.ExecutionId = v
}

// GetCseResourceType returns the CseResourceType field value
func (o *ExecutedAutomation) GetCseResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CseResourceType
}

// GetCseResourceTypeOk returns a tuple with the CseResourceType field value
// and a boolean to check if the value has been set.
func (o *ExecutedAutomation) GetCseResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CseResourceType, true
}

// SetCseResourceType sets field value
func (o *ExecutedAutomation) SetCseResourceType(v string) {
	o.CseResourceType = v
}

// GetCseResourceId returns the CseResourceId field value
func (o *ExecutedAutomation) GetCseResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CseResourceId
}

// GetCseResourceIdOk returns a tuple with the CseResourceId field value
// and a boolean to check if the value has been set.
func (o *ExecutedAutomation) GetCseResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CseResourceId, true
}

// SetCseResourceId sets field value
func (o *ExecutedAutomation) SetCseResourceId(v string) {
	o.CseResourceId = v
}

// GetStatus returns the Status field value
func (o *ExecutedAutomation) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExecutedAutomation) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ExecutedAutomation) SetStatus(v string) {
	o.Status = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ExecutedAutomation) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutedAutomation) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ExecutedAutomation) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *ExecutedAutomation) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ExecutedAutomation) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutedAutomation) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ExecutedAutomation) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *ExecutedAutomation) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetPlaybookId returns the PlaybookId field value
func (o *ExecutedAutomation) GetPlaybookId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlaybookId
}

// GetPlaybookIdOk returns a tuple with the PlaybookId field value
// and a boolean to check if the value has been set.
func (o *ExecutedAutomation) GetPlaybookIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlaybookId, true
}

// SetPlaybookId sets field value
func (o *ExecutedAutomation) SetPlaybookId(v string) {
	o.PlaybookId = v
}

// GetPlaybookName returns the PlaybookName field value
func (o *ExecutedAutomation) GetPlaybookName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlaybookName
}

// GetPlaybookNameOk returns a tuple with the PlaybookName field value
// and a boolean to check if the value has been set.
func (o *ExecutedAutomation) GetPlaybookNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlaybookName, true
}

// SetPlaybookName sets field value
func (o *ExecutedAutomation) SetPlaybookName(v string) {
	o.PlaybookName = v
}

// GetPlaybookDescription returns the PlaybookDescription field value if set, zero value otherwise.
func (o *ExecutedAutomation) GetPlaybookDescription() string {
	if o == nil || IsNil(o.PlaybookDescription) {
		var ret string
		return ret
	}
	return *o.PlaybookDescription
}

// GetPlaybookDescriptionOk returns a tuple with the PlaybookDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutedAutomation) GetPlaybookDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PlaybookDescription) {
		return nil, false
	}
	return o.PlaybookDescription, true
}

// HasPlaybookDescription returns a boolean if a field has been set.
func (o *ExecutedAutomation) HasPlaybookDescription() bool {
	if o != nil && !IsNil(o.PlaybookDescription) {
		return true
	}

	return false
}

// SetPlaybookDescription gets a reference to the given string and assigns it to the PlaybookDescription field.
func (o *ExecutedAutomation) SetPlaybookDescription(v string) {
	o.PlaybookDescription = &v
}

func (o ExecutedAutomation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutedAutomation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["executionId"] = o.ExecutionId
	toSerialize["cseResourceType"] = o.CseResourceType
	toSerialize["cseResourceId"] = o.CseResourceId
	toSerialize["status"] = o.Status
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["playbookId"] = o.PlaybookId
	toSerialize["playbookName"] = o.PlaybookName
	if !IsNil(o.PlaybookDescription) {
		toSerialize["playbookDescription"] = o.PlaybookDescription
	}
	return toSerialize, nil
}

func (o *ExecutedAutomation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"executionId",
		"cseResourceType",
		"cseResourceId",
		"status",
		"playbookId",
		"playbookName",
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

	varExecutedAutomation := _ExecutedAutomation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExecutedAutomation)

	if err != nil {
		return err
	}

	*o = ExecutedAutomation(varExecutedAutomation)

	return err
}

type NullableExecutedAutomation struct {
	value *ExecutedAutomation
	isSet bool
}

func (v NullableExecutedAutomation) Get() *ExecutedAutomation {
	return v.value
}

func (v *NullableExecutedAutomation) Set(val *ExecutedAutomation) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutedAutomation) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutedAutomation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutedAutomation(val *ExecutedAutomation) *NullableExecutedAutomation {
	return &NullableExecutedAutomation{value: val, isSet: true}
}

func (v NullableExecutedAutomation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutedAutomation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


