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

// checks if the GetEntityEnrichmentsResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEntityEnrichmentsResponseData{}

// GetEntityEnrichmentsResponseData The details for a single Entity (i.e. user, hostname, IP address, etc.)
type GetEntityEnrichmentsResponseData struct {
	Enrichments []Enrichment `json:"enrichments"`
}

type _GetEntityEnrichmentsResponseData GetEntityEnrichmentsResponseData

// NewGetEntityEnrichmentsResponseData instantiates a new GetEntityEnrichmentsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEntityEnrichmentsResponseData(enrichments []Enrichment) *GetEntityEnrichmentsResponseData {
	this := GetEntityEnrichmentsResponseData{}
	this.Enrichments = enrichments
	return &this
}

// NewGetEntityEnrichmentsResponseDataWithDefaults instantiates a new GetEntityEnrichmentsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEntityEnrichmentsResponseDataWithDefaults() *GetEntityEnrichmentsResponseData {
	this := GetEntityEnrichmentsResponseData{}
	return &this
}

// GetEnrichments returns the Enrichments field value
func (o *GetEntityEnrichmentsResponseData) GetEnrichments() []Enrichment {
	if o == nil {
		var ret []Enrichment
		return ret
	}

	return o.Enrichments
}

// GetEnrichmentsOk returns a tuple with the Enrichments field value
// and a boolean to check if the value has been set.
func (o *GetEntityEnrichmentsResponseData) GetEnrichmentsOk() ([]Enrichment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enrichments, true
}

// SetEnrichments sets field value
func (o *GetEntityEnrichmentsResponseData) SetEnrichments(v []Enrichment) {
	o.Enrichments = v
}

func (o GetEntityEnrichmentsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEntityEnrichmentsResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enrichments"] = o.Enrichments
	return toSerialize, nil
}

func (o *GetEntityEnrichmentsResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enrichments",
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

	varGetEntityEnrichmentsResponseData := _GetEntityEnrichmentsResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEntityEnrichmentsResponseData)

	if err != nil {
		return err
	}

	*o = GetEntityEnrichmentsResponseData(varGetEntityEnrichmentsResponseData)

	return err
}

type NullableGetEntityEnrichmentsResponseData struct {
	value *GetEntityEnrichmentsResponseData
	isSet bool
}

func (v NullableGetEntityEnrichmentsResponseData) Get() *GetEntityEnrichmentsResponseData {
	return v.value
}

func (v *NullableGetEntityEnrichmentsResponseData) Set(val *GetEntityEnrichmentsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEntityEnrichmentsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEntityEnrichmentsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEntityEnrichmentsResponseData(val *GetEntityEnrichmentsResponseData) *NullableGetEntityEnrichmentsResponseData {
	return &NullableGetEntityEnrichmentsResponseData{value: val, isSet: true}
}

func (v NullableGetEntityEnrichmentsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEntityEnrichmentsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


