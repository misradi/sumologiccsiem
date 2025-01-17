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

// checks if the InsightResolution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InsightResolution{}

// InsightResolution struct for InsightResolution
type InsightResolution struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Source string `json:"source"`
	Parent *InsightResolutionParent `json:"parent,omitempty"`
}

type _InsightResolution InsightResolution

// NewInsightResolution instantiates a new InsightResolution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsightResolution(id int32, name string, description string, source string) *InsightResolution {
	this := InsightResolution{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Source = source
	return &this
}

// NewInsightResolutionWithDefaults instantiates a new InsightResolution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsightResolutionWithDefaults() *InsightResolution {
	this := InsightResolution{}
	return &this
}

// GetId returns the Id field value
func (o *InsightResolution) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InsightResolution) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InsightResolution) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *InsightResolution) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InsightResolution) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InsightResolution) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *InsightResolution) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InsightResolution) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *InsightResolution) SetDescription(v string) {
	o.Description = v
}

// GetSource returns the Source field value
func (o *InsightResolution) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *InsightResolution) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *InsightResolution) SetSource(v string) {
	o.Source = v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *InsightResolution) GetParent() InsightResolutionParent {
	if o == nil || IsNil(o.Parent) {
		var ret InsightResolutionParent
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightResolution) GetParentOk() (*InsightResolutionParent, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *InsightResolution) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given InsightResolutionParent and assigns it to the Parent field.
func (o *InsightResolution) SetParent(v InsightResolutionParent) {
	o.Parent = &v
}

func (o InsightResolution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InsightResolution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["source"] = o.Source
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	return toSerialize, nil
}

func (o *InsightResolution) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
		"source",
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

	varInsightResolution := _InsightResolution{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInsightResolution)

	if err != nil {
		return err
	}

	*o = InsightResolution(varInsightResolution)

	return err
}

type NullableInsightResolution struct {
	value *InsightResolution
	isSet bool
}

func (v NullableInsightResolution) Get() *InsightResolution {
	return v.value
}

func (v *NullableInsightResolution) Set(val *InsightResolution) {
	v.value = val
	v.isSet = true
}

func (v NullableInsightResolution) IsSet() bool {
	return v.isSet
}

func (v *NullableInsightResolution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsightResolution(val *InsightResolution) *NullableInsightResolution {
	return &NullableInsightResolution{value: val, isSet: true}
}

func (v NullableInsightResolution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsightResolution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


