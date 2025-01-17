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

// checks if the GetInsightsResponseDataObjectsInnerArtifactsInnerSignal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsightsResponseDataObjectsInnerArtifactsInnerSignal{}

// GetInsightsResponseDataObjectsInnerArtifactsInnerSignal struct for GetInsightsResponseDataObjectsInnerArtifactsInnerSignal
type GetInsightsResponseDataObjectsInnerArtifactsInnerSignal struct {
	Id string `json:"id"`
	Name string `json:"name"`
	// Timestamp of first log record for this signal.
	Timestamp time.Time `json:"timestamp"`
}

type _GetInsightsResponseDataObjectsInnerArtifactsInnerSignal GetInsightsResponseDataObjectsInnerArtifactsInnerSignal

// NewGetInsightsResponseDataObjectsInnerArtifactsInnerSignal instantiates a new GetInsightsResponseDataObjectsInnerArtifactsInnerSignal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsightsResponseDataObjectsInnerArtifactsInnerSignal(id string, name string, timestamp time.Time) *GetInsightsResponseDataObjectsInnerArtifactsInnerSignal {
	this := GetInsightsResponseDataObjectsInnerArtifactsInnerSignal{}
	this.Id = id
	this.Name = name
	this.Timestamp = timestamp
	return &this
}

// NewGetInsightsResponseDataObjectsInnerArtifactsInnerSignalWithDefaults instantiates a new GetInsightsResponseDataObjectsInnerArtifactsInnerSignal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsightsResponseDataObjectsInnerArtifactsInnerSignalWithDefaults() *GetInsightsResponseDataObjectsInnerArtifactsInnerSignal {
	this := GetInsightsResponseDataObjectsInnerArtifactsInnerSignal{}
	return &this
}

// GetId returns the Id field value
func (o *GetInsightsResponseDataObjectsInnerArtifactsInnerSignal) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerArtifactsInnerSignal) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetInsightsResponseDataObjectsInnerArtifactsInnerSignal) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GetInsightsResponseDataObjectsInnerArtifactsInnerSignal) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerArtifactsInnerSignal) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetInsightsResponseDataObjectsInnerArtifactsInnerSignal) SetName(v string) {
	o.Name = v
}

// GetTimestamp returns the Timestamp field value
func (o *GetInsightsResponseDataObjectsInnerArtifactsInnerSignal) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerArtifactsInnerSignal) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *GetInsightsResponseDataObjectsInnerArtifactsInnerSignal) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o GetInsightsResponseDataObjectsInnerArtifactsInnerSignal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsightsResponseDataObjectsInnerArtifactsInnerSignal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

func (o *GetInsightsResponseDataObjectsInnerArtifactsInnerSignal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"timestamp",
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

	varGetInsightsResponseDataObjectsInnerArtifactsInnerSignal := _GetInsightsResponseDataObjectsInnerArtifactsInnerSignal{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetInsightsResponseDataObjectsInnerArtifactsInnerSignal)

	if err != nil {
		return err
	}

	*o = GetInsightsResponseDataObjectsInnerArtifactsInnerSignal(varGetInsightsResponseDataObjectsInnerArtifactsInnerSignal)

	return err
}

type NullableGetInsightsResponseDataObjectsInnerArtifactsInnerSignal struct {
	value *GetInsightsResponseDataObjectsInnerArtifactsInnerSignal
	isSet bool
}

func (v NullableGetInsightsResponseDataObjectsInnerArtifactsInnerSignal) Get() *GetInsightsResponseDataObjectsInnerArtifactsInnerSignal {
	return v.value
}

func (v *NullableGetInsightsResponseDataObjectsInnerArtifactsInnerSignal) Set(val *GetInsightsResponseDataObjectsInnerArtifactsInnerSignal) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsightsResponseDataObjectsInnerArtifactsInnerSignal) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsightsResponseDataObjectsInnerArtifactsInnerSignal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsightsResponseDataObjectsInnerArtifactsInnerSignal(val *GetInsightsResponseDataObjectsInnerArtifactsInnerSignal) *NullableGetInsightsResponseDataObjectsInnerArtifactsInnerSignal {
	return &NullableGetInsightsResponseDataObjectsInnerArtifactsInnerSignal{value: val, isSet: true}
}

func (v NullableGetInsightsResponseDataObjectsInnerArtifactsInnerSignal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsightsResponseDataObjectsInnerArtifactsInnerSignal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


