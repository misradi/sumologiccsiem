/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologiccsiem

import (
	"encoding/json"
)

// checks if the GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride{}

// GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride struct for GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride
type GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride struct {
	Original []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner `json:"original,omitempty"`
	Override []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner `json:"override,omitempty"`
}

// NewGetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride instantiates a new GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride() *GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride {
	this := GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride{}
	return &this
}

// NewGetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverrideWithDefaults instantiates a new GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverrideWithDefaults() *GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride {
	this := GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride{}
	return &this
}

// GetOriginal returns the Original field value if set, zero value otherwise.
func (o *GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) GetOriginal() []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner {
	if o == nil || IsNil(o.Original) {
		var ret []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner
		return ret
	}
	return o.Original
}

// GetOriginalOk returns a tuple with the Original field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) GetOriginalOk() ([]GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner, bool) {
	if o == nil || IsNil(o.Original) {
		return nil, false
	}
	return o.Original, true
}

// HasOriginal returns a boolean if a field has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) HasOriginal() bool {
	if o != nil && !IsNil(o.Original) {
		return true
	}

	return false
}

// SetOriginal gets a reference to the given []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner and assigns it to the Original field.
func (o *GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) SetOriginal(v []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner) {
	o.Original = v
}

// GetOverride returns the Override field value if set, zero value otherwise.
func (o *GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) GetOverride() []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner {
	if o == nil || IsNil(o.Override) {
		var ret []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner
		return ret
	}
	return o.Override
}

// GetOverrideOk returns a tuple with the Override field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) GetOverrideOk() ([]GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner, bool) {
	if o == nil || IsNil(o.Override) {
		return nil, false
	}
	return o.Override, true
}

// HasOverride returns a boolean if a field has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) HasOverride() bool {
	if o != nil && !IsNil(o.Override) {
		return true
	}

	return false
}

// SetOverride gets a reference to the given []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner and assigns it to the Override field.
func (o *GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) SetOverride(v []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner) {
	o.Override = v
}

func (o GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Original) {
		toSerialize["original"] = o.Original
	}
	if !IsNil(o.Override) {
		toSerialize["override"] = o.Override
	}
	return toSerialize, nil
}

type NullableGetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride struct {
	value *GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride
	isSet bool
}

func (v NullableGetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) Get() *GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride {
	return v.value
}

func (v *NullableGetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) Set(val *GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride(val *GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) *NullableGetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride {
	return &NullableGetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride{value: val, isSet: true}
}

func (v NullableGetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


