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

// checks if the CreateLogMappingRequestBodyFieldsFieldsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLogMappingRequestBodyFieldsFieldsInner{}

// CreateLogMappingRequestBodyFieldsFieldsInner struct for CreateLogMappingRequestBodyFieldsFieldsInner
type CreateLogMappingRequestBodyFieldsFieldsInner struct {
	Name string `json:"name"`
	Value *string `json:"value,omitempty"`
	ValueType *string `json:"valueType,omitempty"`
	SkippedValues []string `json:"skippedValues,omitempty"`
	Lookup []CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner `json:"lookup,omitempty"`
	DefaultValue *string `json:"defaultValue,omitempty"`
	Format *string `json:"format,omitempty"`
	CaseInsensitive *bool `json:"caseInsensitive,omitempty"`
	AlternateValues []string `json:"alternateValues,omitempty"`
	TimeZone *string `json:"timeZone,omitempty"`
	SplitDelimiter *string `json:"splitDelimiter,omitempty"`
	SplitIndex *string `json:"splitIndex,omitempty"`
	FieldJoin []string `json:"fieldJoin,omitempty"`
	JoinDelimiter *string `json:"joinDelimiter,omitempty"`
	FormatParameters []string `json:"formatParameters,omitempty"`
}

type _CreateLogMappingRequestBodyFieldsFieldsInner CreateLogMappingRequestBodyFieldsFieldsInner

// NewCreateLogMappingRequestBodyFieldsFieldsInner instantiates a new CreateLogMappingRequestBodyFieldsFieldsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLogMappingRequestBodyFieldsFieldsInner(name string) *CreateLogMappingRequestBodyFieldsFieldsInner {
	this := CreateLogMappingRequestBodyFieldsFieldsInner{}
	this.Name = name
	return &this
}

// NewCreateLogMappingRequestBodyFieldsFieldsInnerWithDefaults instantiates a new CreateLogMappingRequestBodyFieldsFieldsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLogMappingRequestBodyFieldsFieldsInnerWithDefaults() *CreateLogMappingRequestBodyFieldsFieldsInner {
	this := CreateLogMappingRequestBodyFieldsFieldsInner{}
	return &this
}

// GetName returns the Name field value
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetValue(v string) {
	o.Value = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetValueType() string {
	if o == nil || IsNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetValueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetValueType(v string) {
	o.ValueType = &v
}

// GetSkippedValues returns the SkippedValues field value if set, zero value otherwise.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetSkippedValues() []string {
	if o == nil || IsNil(o.SkippedValues) {
		var ret []string
		return ret
	}
	return o.SkippedValues
}

// GetSkippedValuesOk returns a tuple with the SkippedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetSkippedValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.SkippedValues) {
		return nil, false
	}
	return o.SkippedValues, true
}

// HasSkippedValues returns a boolean if a field has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasSkippedValues() bool {
	if o != nil && !IsNil(o.SkippedValues) {
		return true
	}

	return false
}

// SetSkippedValues gets a reference to the given []string and assigns it to the SkippedValues field.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetSkippedValues(v []string) {
	o.SkippedValues = v
}

// GetLookup returns the Lookup field value if set, zero value otherwise.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetLookup() []CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner {
	if o == nil || IsNil(o.Lookup) {
		var ret []CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner
		return ret
	}
	return o.Lookup
}

// GetLookupOk returns a tuple with the Lookup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetLookupOk() ([]CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner, bool) {
	if o == nil || IsNil(o.Lookup) {
		return nil, false
	}
	return o.Lookup, true
}

// HasLookup returns a boolean if a field has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasLookup() bool {
	if o != nil && !IsNil(o.Lookup) {
		return true
	}

	return false
}

// SetLookup gets a reference to the given []CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner and assigns it to the Lookup field.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetLookup(v []CreateLogMappingRequestBodyFieldsFieldsInnerLookupInner) {
	o.Lookup = v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetFormat(v string) {
	o.Format = &v
}

// GetCaseInsensitive returns the CaseInsensitive field value if set, zero value otherwise.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetCaseInsensitive() bool {
	if o == nil || IsNil(o.CaseInsensitive) {
		var ret bool
		return ret
	}
	return *o.CaseInsensitive
}

// GetCaseInsensitiveOk returns a tuple with the CaseInsensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetCaseInsensitiveOk() (*bool, bool) {
	if o == nil || IsNil(o.CaseInsensitive) {
		return nil, false
	}
	return o.CaseInsensitive, true
}

// HasCaseInsensitive returns a boolean if a field has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasCaseInsensitive() bool {
	if o != nil && !IsNil(o.CaseInsensitive) {
		return true
	}

	return false
}

// SetCaseInsensitive gets a reference to the given bool and assigns it to the CaseInsensitive field.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetCaseInsensitive(v bool) {
	o.CaseInsensitive = &v
}

// GetAlternateValues returns the AlternateValues field value if set, zero value otherwise.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetAlternateValues() []string {
	if o == nil || IsNil(o.AlternateValues) {
		var ret []string
		return ret
	}
	return o.AlternateValues
}

// GetAlternateValuesOk returns a tuple with the AlternateValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetAlternateValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.AlternateValues) {
		return nil, false
	}
	return o.AlternateValues, true
}

// HasAlternateValues returns a boolean if a field has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasAlternateValues() bool {
	if o != nil && !IsNil(o.AlternateValues) {
		return true
	}

	return false
}

// SetAlternateValues gets a reference to the given []string and assigns it to the AlternateValues field.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetAlternateValues(v []string) {
	o.AlternateValues = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetSplitDelimiter returns the SplitDelimiter field value if set, zero value otherwise.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetSplitDelimiter() string {
	if o == nil || IsNil(o.SplitDelimiter) {
		var ret string
		return ret
	}
	return *o.SplitDelimiter
}

// GetSplitDelimiterOk returns a tuple with the SplitDelimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetSplitDelimiterOk() (*string, bool) {
	if o == nil || IsNil(o.SplitDelimiter) {
		return nil, false
	}
	return o.SplitDelimiter, true
}

// HasSplitDelimiter returns a boolean if a field has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasSplitDelimiter() bool {
	if o != nil && !IsNil(o.SplitDelimiter) {
		return true
	}

	return false
}

// SetSplitDelimiter gets a reference to the given string and assigns it to the SplitDelimiter field.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetSplitDelimiter(v string) {
	o.SplitDelimiter = &v
}

// GetSplitIndex returns the SplitIndex field value if set, zero value otherwise.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetSplitIndex() string {
	if o == nil || IsNil(o.SplitIndex) {
		var ret string
		return ret
	}
	return *o.SplitIndex
}

// GetSplitIndexOk returns a tuple with the SplitIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetSplitIndexOk() (*string, bool) {
	if o == nil || IsNil(o.SplitIndex) {
		return nil, false
	}
	return o.SplitIndex, true
}

// HasSplitIndex returns a boolean if a field has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasSplitIndex() bool {
	if o != nil && !IsNil(o.SplitIndex) {
		return true
	}

	return false
}

// SetSplitIndex gets a reference to the given string and assigns it to the SplitIndex field.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetSplitIndex(v string) {
	o.SplitIndex = &v
}

// GetFieldJoin returns the FieldJoin field value if set, zero value otherwise.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetFieldJoin() []string {
	if o == nil || IsNil(o.FieldJoin) {
		var ret []string
		return ret
	}
	return o.FieldJoin
}

// GetFieldJoinOk returns a tuple with the FieldJoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetFieldJoinOk() ([]string, bool) {
	if o == nil || IsNil(o.FieldJoin) {
		return nil, false
	}
	return o.FieldJoin, true
}

// HasFieldJoin returns a boolean if a field has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasFieldJoin() bool {
	if o != nil && !IsNil(o.FieldJoin) {
		return true
	}

	return false
}

// SetFieldJoin gets a reference to the given []string and assigns it to the FieldJoin field.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetFieldJoin(v []string) {
	o.FieldJoin = v
}

// GetJoinDelimiter returns the JoinDelimiter field value if set, zero value otherwise.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetJoinDelimiter() string {
	if o == nil || IsNil(o.JoinDelimiter) {
		var ret string
		return ret
	}
	return *o.JoinDelimiter
}

// GetJoinDelimiterOk returns a tuple with the JoinDelimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetJoinDelimiterOk() (*string, bool) {
	if o == nil || IsNil(o.JoinDelimiter) {
		return nil, false
	}
	return o.JoinDelimiter, true
}

// HasJoinDelimiter returns a boolean if a field has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasJoinDelimiter() bool {
	if o != nil && !IsNil(o.JoinDelimiter) {
		return true
	}

	return false
}

// SetJoinDelimiter gets a reference to the given string and assigns it to the JoinDelimiter field.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetJoinDelimiter(v string) {
	o.JoinDelimiter = &v
}

// GetFormatParameters returns the FormatParameters field value if set, zero value otherwise.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetFormatParameters() []string {
	if o == nil || IsNil(o.FormatParameters) {
		var ret []string
		return ret
	}
	return o.FormatParameters
}

// GetFormatParametersOk returns a tuple with the FormatParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) GetFormatParametersOk() ([]string, bool) {
	if o == nil || IsNil(o.FormatParameters) {
		return nil, false
	}
	return o.FormatParameters, true
}

// HasFormatParameters returns a boolean if a field has been set.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) HasFormatParameters() bool {
	if o != nil && !IsNil(o.FormatParameters) {
		return true
	}

	return false
}

// SetFormatParameters gets a reference to the given []string and assigns it to the FormatParameters field.
func (o *CreateLogMappingRequestBodyFieldsFieldsInner) SetFormatParameters(v []string) {
	o.FormatParameters = v
}

func (o CreateLogMappingRequestBodyFieldsFieldsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLogMappingRequestBodyFieldsFieldsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ValueType) {
		toSerialize["valueType"] = o.ValueType
	}
	if !IsNil(o.SkippedValues) {
		toSerialize["skippedValues"] = o.SkippedValues
	}
	if !IsNil(o.Lookup) {
		toSerialize["lookup"] = o.Lookup
	}
	if !IsNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.CaseInsensitive) {
		toSerialize["caseInsensitive"] = o.CaseInsensitive
	}
	if !IsNil(o.AlternateValues) {
		toSerialize["alternateValues"] = o.AlternateValues
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.SplitDelimiter) {
		toSerialize["splitDelimiter"] = o.SplitDelimiter
	}
	if !IsNil(o.SplitIndex) {
		toSerialize["splitIndex"] = o.SplitIndex
	}
	if !IsNil(o.FieldJoin) {
		toSerialize["fieldJoin"] = o.FieldJoin
	}
	if !IsNil(o.JoinDelimiter) {
		toSerialize["joinDelimiter"] = o.JoinDelimiter
	}
	if !IsNil(o.FormatParameters) {
		toSerialize["formatParameters"] = o.FormatParameters
	}
	return toSerialize, nil
}

func (o *CreateLogMappingRequestBodyFieldsFieldsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateLogMappingRequestBodyFieldsFieldsInner := _CreateLogMappingRequestBodyFieldsFieldsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateLogMappingRequestBodyFieldsFieldsInner)

	if err != nil {
		return err
	}

	*o = CreateLogMappingRequestBodyFieldsFieldsInner(varCreateLogMappingRequestBodyFieldsFieldsInner)

	return err
}

type NullableCreateLogMappingRequestBodyFieldsFieldsInner struct {
	value *CreateLogMappingRequestBodyFieldsFieldsInner
	isSet bool
}

func (v NullableCreateLogMappingRequestBodyFieldsFieldsInner) Get() *CreateLogMappingRequestBodyFieldsFieldsInner {
	return v.value
}

func (v *NullableCreateLogMappingRequestBodyFieldsFieldsInner) Set(val *CreateLogMappingRequestBodyFieldsFieldsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLogMappingRequestBodyFieldsFieldsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLogMappingRequestBodyFieldsFieldsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLogMappingRequestBodyFieldsFieldsInner(val *CreateLogMappingRequestBodyFieldsFieldsInner) *NullableCreateLogMappingRequestBodyFieldsFieldsInner {
	return &NullableCreateLogMappingRequestBodyFieldsFieldsInner{value: val, isSet: true}
}

func (v NullableCreateLogMappingRequestBodyFieldsFieldsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLogMappingRequestBodyFieldsFieldsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


