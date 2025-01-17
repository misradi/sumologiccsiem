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

// checks if the GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner{}

// GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner struct for GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner
type GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner struct {
	Id string `json:"id"`
	EntityType string `json:"entityType"`
	Value string `json:"value"`
	SensorZone *string `json:"sensorZone,omitempty"`
}

type _GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner

// NewGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner instantiates a new GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner(id string, entityType string, value string) *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner {
	this := GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner{}
	this.Id = id
	this.EntityType = entityType
	this.Value = value
	return &this
}

// NewGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInnerWithDefaults instantiates a new GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInnerWithDefaults() *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner {
	this := GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner{}
	return &this
}

// GetId returns the Id field value
func (o *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) SetId(v string) {
	o.Id = v
}

// GetEntityType returns the EntityType field value
func (o *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) SetEntityType(v string) {
	o.EntityType = v
}

// GetValue returns the Value field value
func (o *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) SetValue(v string) {
	o.Value = v
}

// GetSensorZone returns the SensorZone field value if set, zero value otherwise.
func (o *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) GetSensorZone() string {
	if o == nil || IsNil(o.SensorZone) {
		var ret string
		return ret
	}
	return *o.SensorZone
}

// GetSensorZoneOk returns a tuple with the SensorZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) GetSensorZoneOk() (*string, bool) {
	if o == nil || IsNil(o.SensorZone) {
		return nil, false
	}
	return o.SensorZone, true
}

// HasSensorZone returns a boolean if a field has been set.
func (o *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) HasSensorZone() bool {
	if o != nil && !IsNil(o.SensorZone) {
		return true
	}

	return false
}

// SetSensorZone gets a reference to the given string and assigns it to the SensorZone field.
func (o *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) SetSensorZone(v string) {
	o.SensorZone = &v
}

func (o GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["entityType"] = o.EntityType
	toSerialize["value"] = o.Value
	if !IsNil(o.SensorZone) {
		toSerialize["sensorZone"] = o.SensorZone
	}
	return toSerialize, nil
}

func (o *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"entityType",
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

	varGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner := _GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner)

	if err != nil {
		return err
	}

	*o = GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner(varGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner)

	return err
}

type NullableGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner struct {
	value *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner
	isSet bool
}

func (v NullableGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) Get() *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner {
	return v.value
}

func (v *NullableGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) Set(val *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner(val *GetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) *NullableGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner {
	return &NullableGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner{value: val, isSet: true}
}

func (v NullableGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsightRelatedEntitiesResponseDataInvolvedEntitiesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


