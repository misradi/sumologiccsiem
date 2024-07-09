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

// checks if the GetTagSchemasResponseDataInnerValueOptionObjectsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTagSchemasResponseDataInnerValueOptionObjectsInner{}

// GetTagSchemasResponseDataInnerValueOptionObjectsInner struct for GetTagSchemasResponseDataInnerValueOptionObjectsInner
type GetTagSchemasResponseDataInnerValueOptionObjectsInner struct {
	Value string `json:"value"`
	Link *string `json:"link,omitempty"`
	Label *string `json:"label,omitempty"`
}

type _GetTagSchemasResponseDataInnerValueOptionObjectsInner GetTagSchemasResponseDataInnerValueOptionObjectsInner

// NewGetTagSchemasResponseDataInnerValueOptionObjectsInner instantiates a new GetTagSchemasResponseDataInnerValueOptionObjectsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTagSchemasResponseDataInnerValueOptionObjectsInner(value string) *GetTagSchemasResponseDataInnerValueOptionObjectsInner {
	this := GetTagSchemasResponseDataInnerValueOptionObjectsInner{}
	this.Value = value
	return &this
}

// NewGetTagSchemasResponseDataInnerValueOptionObjectsInnerWithDefaults instantiates a new GetTagSchemasResponseDataInnerValueOptionObjectsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTagSchemasResponseDataInnerValueOptionObjectsInnerWithDefaults() *GetTagSchemasResponseDataInnerValueOptionObjectsInner {
	this := GetTagSchemasResponseDataInnerValueOptionObjectsInner{}
	return &this
}

// GetValue returns the Value field value
func (o *GetTagSchemasResponseDataInnerValueOptionObjectsInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetTagSchemasResponseDataInnerValueOptionObjectsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetTagSchemasResponseDataInnerValueOptionObjectsInner) SetValue(v string) {
	o.Value = v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *GetTagSchemasResponseDataInnerValueOptionObjectsInner) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTagSchemasResponseDataInnerValueOptionObjectsInner) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *GetTagSchemasResponseDataInnerValueOptionObjectsInner) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *GetTagSchemasResponseDataInnerValueOptionObjectsInner) SetLink(v string) {
	o.Link = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetTagSchemasResponseDataInnerValueOptionObjectsInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTagSchemasResponseDataInnerValueOptionObjectsInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetTagSchemasResponseDataInnerValueOptionObjectsInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetTagSchemasResponseDataInnerValueOptionObjectsInner) SetLabel(v string) {
	o.Label = &v
}

func (o GetTagSchemasResponseDataInnerValueOptionObjectsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTagSchemasResponseDataInnerValueOptionObjectsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

func (o *GetTagSchemasResponseDataInnerValueOptionObjectsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varGetTagSchemasResponseDataInnerValueOptionObjectsInner := _GetTagSchemasResponseDataInnerValueOptionObjectsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTagSchemasResponseDataInnerValueOptionObjectsInner)

	if err != nil {
		return err
	}

	*o = GetTagSchemasResponseDataInnerValueOptionObjectsInner(varGetTagSchemasResponseDataInnerValueOptionObjectsInner)

	return err
}

type NullableGetTagSchemasResponseDataInnerValueOptionObjectsInner struct {
	value *GetTagSchemasResponseDataInnerValueOptionObjectsInner
	isSet bool
}

func (v NullableGetTagSchemasResponseDataInnerValueOptionObjectsInner) Get() *GetTagSchemasResponseDataInnerValueOptionObjectsInner {
	return v.value
}

func (v *NullableGetTagSchemasResponseDataInnerValueOptionObjectsInner) Set(val *GetTagSchemasResponseDataInnerValueOptionObjectsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTagSchemasResponseDataInnerValueOptionObjectsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTagSchemasResponseDataInnerValueOptionObjectsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTagSchemasResponseDataInnerValueOptionObjectsInner(val *GetTagSchemasResponseDataInnerValueOptionObjectsInner) *NullableGetTagSchemasResponseDataInnerValueOptionObjectsInner {
	return &NullableGetTagSchemasResponseDataInnerValueOptionObjectsInner{value: val, isSet: true}
}

func (v NullableGetTagSchemasResponseDataInnerValueOptionObjectsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTagSchemasResponseDataInnerValueOptionObjectsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


