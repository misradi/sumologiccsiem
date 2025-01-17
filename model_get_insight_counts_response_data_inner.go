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

// checks if the GetInsightCountsResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsightCountsResponseDataInner{}

// GetInsightCountsResponseDataInner struct for GetInsightCountsResponseDataInner
type GetInsightCountsResponseDataInner struct {
	Timestamp time.Time `json:"timestamp"`
	Value string `json:"value"`
}

type _GetInsightCountsResponseDataInner GetInsightCountsResponseDataInner

// NewGetInsightCountsResponseDataInner instantiates a new GetInsightCountsResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsightCountsResponseDataInner(timestamp time.Time, value string) *GetInsightCountsResponseDataInner {
	this := GetInsightCountsResponseDataInner{}
	this.Timestamp = timestamp
	this.Value = value
	return &this
}

// NewGetInsightCountsResponseDataInnerWithDefaults instantiates a new GetInsightCountsResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsightCountsResponseDataInnerWithDefaults() *GetInsightCountsResponseDataInner {
	this := GetInsightCountsResponseDataInner{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *GetInsightCountsResponseDataInner) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *GetInsightCountsResponseDataInner) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *GetInsightCountsResponseDataInner) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetValue returns the Value field value
func (o *GetInsightCountsResponseDataInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetInsightCountsResponseDataInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetInsightCountsResponseDataInner) SetValue(v string) {
	o.Value = v
}

func (o GetInsightCountsResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsightCountsResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *GetInsightCountsResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"timestamp",
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

	varGetInsightCountsResponseDataInner := _GetInsightCountsResponseDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetInsightCountsResponseDataInner)

	if err != nil {
		return err
	}

	*o = GetInsightCountsResponseDataInner(varGetInsightCountsResponseDataInner)

	return err
}

type NullableGetInsightCountsResponseDataInner struct {
	value *GetInsightCountsResponseDataInner
	isSet bool
}

func (v NullableGetInsightCountsResponseDataInner) Get() *GetInsightCountsResponseDataInner {
	return v.value
}

func (v *NullableGetInsightCountsResponseDataInner) Set(val *GetInsightCountsResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsightCountsResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsightCountsResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsightCountsResponseDataInner(val *GetInsightCountsResponseDataInner) *NullableGetInsightCountsResponseDataInner {
	return &NullableGetInsightCountsResponseDataInner{value: val, isSet: true}
}

func (v NullableGetInsightCountsResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsightCountsResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


