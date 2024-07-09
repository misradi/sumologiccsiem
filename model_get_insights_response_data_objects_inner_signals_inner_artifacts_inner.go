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

// checks if the GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner{}

// GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner struct for GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner
type GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner struct {
	Name string `json:"name"`
	Value string `json:"value"`
	RecordUid string `json:"recordUid"`
	RecordStream string `json:"recordStream"`
}

type _GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner

// NewGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner instantiates a new GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner(name string, value string, recordUid string, recordStream string) *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner {
	this := GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner{}
	this.Name = name
	this.Value = value
	this.RecordUid = recordUid
	this.RecordStream = recordStream
	return &this
}

// NewGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInnerWithDefaults instantiates a new GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInnerWithDefaults() *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner {
	this := GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner{}
	return &this
}

// GetName returns the Name field value
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) SetValue(v string) {
	o.Value = v
}

// GetRecordUid returns the RecordUid field value
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) GetRecordUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordUid
}

// GetRecordUidOk returns a tuple with the RecordUid field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) GetRecordUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordUid, true
}

// SetRecordUid sets field value
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) SetRecordUid(v string) {
	o.RecordUid = v
}

// GetRecordStream returns the RecordStream field value
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) GetRecordStream() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordStream
}

// GetRecordStreamOk returns a tuple with the RecordStream field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) GetRecordStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordStream, true
}

// SetRecordStream sets field value
func (o *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) SetRecordStream(v string) {
	o.RecordStream = v
}

func (o GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	toSerialize["recordUid"] = o.RecordUid
	toSerialize["recordStream"] = o.RecordStream
	return toSerialize, nil
}

func (o *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
		"recordUid",
		"recordStream",
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

	varGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner := _GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner)

	if err != nil {
		return err
	}

	*o = GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner(varGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner)

	return err
}

type NullableGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner struct {
	value *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner
	isSet bool
}

func (v NullableGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) Get() *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner {
	return v.value
}

func (v *NullableGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) Set(val *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner(val *GetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) *NullableGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner {
	return &NullableGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner{value: val, isSet: true}
}

func (v NullableGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsightsResponseDataObjectsInnerSignalsInnerArtifactsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

