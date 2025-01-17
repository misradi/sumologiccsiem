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

// checks if the CreateCustomInsightRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomInsightRequestBodyFields{}

// CreateCustomInsightRequestBodyFields struct for CreateCustomInsightRequestBodyFields
type CreateCustomInsightRequestBodyFields struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Severity string `json:"severity"`
	Ordered bool `json:"ordered"`
	Enabled bool `json:"enabled"`
	Tags []string `json:"tags"`
	DynamicSeverity []CreateCustomInsightRequestBodyFieldsDynamicSeverityInner `json:"dynamicSeverity,omitempty"`
	RuleIds []string `json:"ruleIds,omitempty"`
	SignalNames []string `json:"signalNames,omitempty"`
}

type _CreateCustomInsightRequestBodyFields CreateCustomInsightRequestBodyFields

// NewCreateCustomInsightRequestBodyFields instantiates a new CreateCustomInsightRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomInsightRequestBodyFields(name string, description string, severity string, ordered bool, enabled bool, tags []string) *CreateCustomInsightRequestBodyFields {
	this := CreateCustomInsightRequestBodyFields{}
	this.Name = name
	this.Description = description
	this.Severity = severity
	this.Ordered = ordered
	this.Enabled = enabled
	this.Tags = tags
	return &this
}

// NewCreateCustomInsightRequestBodyFieldsWithDefaults instantiates a new CreateCustomInsightRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomInsightRequestBodyFieldsWithDefaults() *CreateCustomInsightRequestBodyFields {
	this := CreateCustomInsightRequestBodyFields{}
	return &this
}

// GetName returns the Name field value
func (o *CreateCustomInsightRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCustomInsightRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCustomInsightRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *CreateCustomInsightRequestBodyFields) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateCustomInsightRequestBodyFields) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateCustomInsightRequestBodyFields) SetDescription(v string) {
	o.Description = v
}

// GetSeverity returns the Severity field value
func (o *CreateCustomInsightRequestBodyFields) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *CreateCustomInsightRequestBodyFields) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *CreateCustomInsightRequestBodyFields) SetSeverity(v string) {
	o.Severity = v
}

// GetOrdered returns the Ordered field value
func (o *CreateCustomInsightRequestBodyFields) GetOrdered() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ordered
}

// GetOrderedOk returns a tuple with the Ordered field value
// and a boolean to check if the value has been set.
func (o *CreateCustomInsightRequestBodyFields) GetOrderedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ordered, true
}

// SetOrdered sets field value
func (o *CreateCustomInsightRequestBodyFields) SetOrdered(v bool) {
	o.Ordered = v
}

// GetEnabled returns the Enabled field value
func (o *CreateCustomInsightRequestBodyFields) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateCustomInsightRequestBodyFields) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateCustomInsightRequestBodyFields) SetEnabled(v bool) {
	o.Enabled = v
}

// GetTags returns the Tags field value
func (o *CreateCustomInsightRequestBodyFields) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *CreateCustomInsightRequestBodyFields) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *CreateCustomInsightRequestBodyFields) SetTags(v []string) {
	o.Tags = v
}

// GetDynamicSeverity returns the DynamicSeverity field value if set, zero value otherwise.
func (o *CreateCustomInsightRequestBodyFields) GetDynamicSeverity() []CreateCustomInsightRequestBodyFieldsDynamicSeverityInner {
	if o == nil || IsNil(o.DynamicSeverity) {
		var ret []CreateCustomInsightRequestBodyFieldsDynamicSeverityInner
		return ret
	}
	return o.DynamicSeverity
}

// GetDynamicSeverityOk returns a tuple with the DynamicSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomInsightRequestBodyFields) GetDynamicSeverityOk() ([]CreateCustomInsightRequestBodyFieldsDynamicSeverityInner, bool) {
	if o == nil || IsNil(o.DynamicSeverity) {
		return nil, false
	}
	return o.DynamicSeverity, true
}

// HasDynamicSeverity returns a boolean if a field has been set.
func (o *CreateCustomInsightRequestBodyFields) HasDynamicSeverity() bool {
	if o != nil && !IsNil(o.DynamicSeverity) {
		return true
	}

	return false
}

// SetDynamicSeverity gets a reference to the given []CreateCustomInsightRequestBodyFieldsDynamicSeverityInner and assigns it to the DynamicSeverity field.
func (o *CreateCustomInsightRequestBodyFields) SetDynamicSeverity(v []CreateCustomInsightRequestBodyFieldsDynamicSeverityInner) {
	o.DynamicSeverity = v
}

// GetRuleIds returns the RuleIds field value if set, zero value otherwise.
func (o *CreateCustomInsightRequestBodyFields) GetRuleIds() []string {
	if o == nil || IsNil(o.RuleIds) {
		var ret []string
		return ret
	}
	return o.RuleIds
}

// GetRuleIdsOk returns a tuple with the RuleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomInsightRequestBodyFields) GetRuleIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RuleIds) {
		return nil, false
	}
	return o.RuleIds, true
}

// HasRuleIds returns a boolean if a field has been set.
func (o *CreateCustomInsightRequestBodyFields) HasRuleIds() bool {
	if o != nil && !IsNil(o.RuleIds) {
		return true
	}

	return false
}

// SetRuleIds gets a reference to the given []string and assigns it to the RuleIds field.
func (o *CreateCustomInsightRequestBodyFields) SetRuleIds(v []string) {
	o.RuleIds = v
}

// GetSignalNames returns the SignalNames field value if set, zero value otherwise.
func (o *CreateCustomInsightRequestBodyFields) GetSignalNames() []string {
	if o == nil || IsNil(o.SignalNames) {
		var ret []string
		return ret
	}
	return o.SignalNames
}

// GetSignalNamesOk returns a tuple with the SignalNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomInsightRequestBodyFields) GetSignalNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.SignalNames) {
		return nil, false
	}
	return o.SignalNames, true
}

// HasSignalNames returns a boolean if a field has been set.
func (o *CreateCustomInsightRequestBodyFields) HasSignalNames() bool {
	if o != nil && !IsNil(o.SignalNames) {
		return true
	}

	return false
}

// SetSignalNames gets a reference to the given []string and assigns it to the SignalNames field.
func (o *CreateCustomInsightRequestBodyFields) SetSignalNames(v []string) {
	o.SignalNames = v
}

func (o CreateCustomInsightRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomInsightRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["severity"] = o.Severity
	toSerialize["ordered"] = o.Ordered
	toSerialize["enabled"] = o.Enabled
	toSerialize["tags"] = o.Tags
	if !IsNil(o.DynamicSeverity) {
		toSerialize["dynamicSeverity"] = o.DynamicSeverity
	}
	if !IsNil(o.RuleIds) {
		toSerialize["ruleIds"] = o.RuleIds
	}
	if !IsNil(o.SignalNames) {
		toSerialize["signalNames"] = o.SignalNames
	}
	return toSerialize, nil
}

func (o *CreateCustomInsightRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"severity",
		"ordered",
		"enabled",
		"tags",
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

	varCreateCustomInsightRequestBodyFields := _CreateCustomInsightRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateCustomInsightRequestBodyFields)

	if err != nil {
		return err
	}

	*o = CreateCustomInsightRequestBodyFields(varCreateCustomInsightRequestBodyFields)

	return err
}

type NullableCreateCustomInsightRequestBodyFields struct {
	value *CreateCustomInsightRequestBodyFields
	isSet bool
}

func (v NullableCreateCustomInsightRequestBodyFields) Get() *CreateCustomInsightRequestBodyFields {
	return v.value
}

func (v *NullableCreateCustomInsightRequestBodyFields) Set(val *CreateCustomInsightRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomInsightRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomInsightRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomInsightRequestBodyFields(val *CreateCustomInsightRequestBodyFields) *NullableCreateCustomInsightRequestBodyFields {
	return &NullableCreateCustomInsightRequestBodyFields{value: val, isSet: true}
}

func (v NullableCreateCustomInsightRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomInsightRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


