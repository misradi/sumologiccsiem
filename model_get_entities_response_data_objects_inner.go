/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the GetEntitiesResponseDataObjectsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEntitiesResponseDataObjectsInner{}

// GetEntitiesResponseDataObjectsInner struct for GetEntitiesResponseDataObjectsInner
type GetEntitiesResponseDataObjectsInner struct {
	Id string `json:"id"`
	ActivityScore int32 `json:"activityScore"`
	LastSeen *time.Time `json:"lastSeen,omitempty"`
	FirstSeen *time.Time `json:"firstSeen,omitempty"`
	IsSuppressed bool `json:"isSuppressed"`
	Name string `json:"name"`
	Inventory []GetEntitiesResponseDataObjectsInnerInventoryInner `json:"inventory"`
	Tags []string `json:"tags"`
	Criticality *string `json:"criticality,omitempty"`
	Reputation *string `json:"reputation,omitempty"`
	RecentSignalSeverity int32 `json:"recentSignalSeverity"`
	EntityType string `json:"entityType"`
	SensorZone *string `json:"sensorZone,omitempty"`
	Value string `json:"value"`
	Hostname *string `json:"hostname,omitempty"`
	MacAddress *string `json:"macAddress,omitempty"`
}

type _GetEntitiesResponseDataObjectsInner GetEntitiesResponseDataObjectsInner

// NewGetEntitiesResponseDataObjectsInner instantiates a new GetEntitiesResponseDataObjectsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEntitiesResponseDataObjectsInner(id string, activityScore int32, isSuppressed bool, name string, inventory []GetEntitiesResponseDataObjectsInnerInventoryInner, tags []string, recentSignalSeverity int32, entityType string, value string) *GetEntitiesResponseDataObjectsInner {
	this := GetEntitiesResponseDataObjectsInner{}
	this.Id = id
	this.ActivityScore = activityScore
	this.IsSuppressed = isSuppressed
	this.Name = name
	this.Inventory = inventory
	this.Tags = tags
	this.RecentSignalSeverity = recentSignalSeverity
	this.EntityType = entityType
	this.Value = value
	return &this
}

// NewGetEntitiesResponseDataObjectsInnerWithDefaults instantiates a new GetEntitiesResponseDataObjectsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEntitiesResponseDataObjectsInnerWithDefaults() *GetEntitiesResponseDataObjectsInner {
	this := GetEntitiesResponseDataObjectsInner{}
	return &this
}

// GetId returns the Id field value
func (o *GetEntitiesResponseDataObjectsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseDataObjectsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetEntitiesResponseDataObjectsInner) SetId(v string) {
	o.Id = v
}

// GetActivityScore returns the ActivityScore field value
func (o *GetEntitiesResponseDataObjectsInner) GetActivityScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ActivityScore
}

// GetActivityScoreOk returns a tuple with the ActivityScore field value
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseDataObjectsInner) GetActivityScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivityScore, true
}

// SetActivityScore sets field value
func (o *GetEntitiesResponseDataObjectsInner) SetActivityScore(v int32) {
	o.ActivityScore = v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *GetEntitiesResponseDataObjectsInner) GetLastSeen() time.Time {
	if o == nil || IsNil(o.LastSeen) {
		var ret time.Time
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseDataObjectsInner) GetLastSeenOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastSeen) {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *GetEntitiesResponseDataObjectsInner) HasLastSeen() bool {
	if o != nil && !IsNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given time.Time and assigns it to the LastSeen field.
func (o *GetEntitiesResponseDataObjectsInner) SetLastSeen(v time.Time) {
	o.LastSeen = &v
}

// GetFirstSeen returns the FirstSeen field value if set, zero value otherwise.
func (o *GetEntitiesResponseDataObjectsInner) GetFirstSeen() time.Time {
	if o == nil || IsNil(o.FirstSeen) {
		var ret time.Time
		return ret
	}
	return *o.FirstSeen
}

// GetFirstSeenOk returns a tuple with the FirstSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseDataObjectsInner) GetFirstSeenOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FirstSeen) {
		return nil, false
	}
	return o.FirstSeen, true
}

// HasFirstSeen returns a boolean if a field has been set.
func (o *GetEntitiesResponseDataObjectsInner) HasFirstSeen() bool {
	if o != nil && !IsNil(o.FirstSeen) {
		return true
	}

	return false
}

// SetFirstSeen gets a reference to the given time.Time and assigns it to the FirstSeen field.
func (o *GetEntitiesResponseDataObjectsInner) SetFirstSeen(v time.Time) {
	o.FirstSeen = &v
}

// GetIsSuppressed returns the IsSuppressed field value
func (o *GetEntitiesResponseDataObjectsInner) GetIsSuppressed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSuppressed
}

// GetIsSuppressedOk returns a tuple with the IsSuppressed field value
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseDataObjectsInner) GetIsSuppressedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSuppressed, true
}

// SetIsSuppressed sets field value
func (o *GetEntitiesResponseDataObjectsInner) SetIsSuppressed(v bool) {
	o.IsSuppressed = v
}

// GetName returns the Name field value
func (o *GetEntitiesResponseDataObjectsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseDataObjectsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetEntitiesResponseDataObjectsInner) SetName(v string) {
	o.Name = v
}

// GetInventory returns the Inventory field value
func (o *GetEntitiesResponseDataObjectsInner) GetInventory() []GetEntitiesResponseDataObjectsInnerInventoryInner {
	if o == nil {
		var ret []GetEntitiesResponseDataObjectsInnerInventoryInner
		return ret
	}

	return o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseDataObjectsInner) GetInventoryOk() ([]GetEntitiesResponseDataObjectsInnerInventoryInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inventory, true
}

// SetInventory sets field value
func (o *GetEntitiesResponseDataObjectsInner) SetInventory(v []GetEntitiesResponseDataObjectsInnerInventoryInner) {
	o.Inventory = v
}

// GetTags returns the Tags field value
func (o *GetEntitiesResponseDataObjectsInner) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseDataObjectsInner) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *GetEntitiesResponseDataObjectsInner) SetTags(v []string) {
	o.Tags = v
}

// GetCriticality returns the Criticality field value if set, zero value otherwise.
func (o *GetEntitiesResponseDataObjectsInner) GetCriticality() string {
	if o == nil || IsNil(o.Criticality) {
		var ret string
		return ret
	}
	return *o.Criticality
}

// GetCriticalityOk returns a tuple with the Criticality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseDataObjectsInner) GetCriticalityOk() (*string, bool) {
	if o == nil || IsNil(o.Criticality) {
		return nil, false
	}
	return o.Criticality, true
}

// HasCriticality returns a boolean if a field has been set.
func (o *GetEntitiesResponseDataObjectsInner) HasCriticality() bool {
	if o != nil && !IsNil(o.Criticality) {
		return true
	}

	return false
}

// SetCriticality gets a reference to the given string and assigns it to the Criticality field.
func (o *GetEntitiesResponseDataObjectsInner) SetCriticality(v string) {
	o.Criticality = &v
}

// GetReputation returns the Reputation field value if set, zero value otherwise.
func (o *GetEntitiesResponseDataObjectsInner) GetReputation() string {
	if o == nil || IsNil(o.Reputation) {
		var ret string
		return ret
	}
	return *o.Reputation
}

// GetReputationOk returns a tuple with the Reputation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseDataObjectsInner) GetReputationOk() (*string, bool) {
	if o == nil || IsNil(o.Reputation) {
		return nil, false
	}
	return o.Reputation, true
}

// HasReputation returns a boolean if a field has been set.
func (o *GetEntitiesResponseDataObjectsInner) HasReputation() bool {
	if o != nil && !IsNil(o.Reputation) {
		return true
	}

	return false
}

// SetReputation gets a reference to the given string and assigns it to the Reputation field.
func (o *GetEntitiesResponseDataObjectsInner) SetReputation(v string) {
	o.Reputation = &v
}

// GetRecentSignalSeverity returns the RecentSignalSeverity field value
func (o *GetEntitiesResponseDataObjectsInner) GetRecentSignalSeverity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RecentSignalSeverity
}

// GetRecentSignalSeverityOk returns a tuple with the RecentSignalSeverity field value
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseDataObjectsInner) GetRecentSignalSeverityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecentSignalSeverity, true
}

// SetRecentSignalSeverity sets field value
func (o *GetEntitiesResponseDataObjectsInner) SetRecentSignalSeverity(v int32) {
	o.RecentSignalSeverity = v
}

// GetEntityType returns the EntityType field value
func (o *GetEntitiesResponseDataObjectsInner) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseDataObjectsInner) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *GetEntitiesResponseDataObjectsInner) SetEntityType(v string) {
	o.EntityType = v
}

// GetSensorZone returns the SensorZone field value if set, zero value otherwise.
func (o *GetEntitiesResponseDataObjectsInner) GetSensorZone() string {
	if o == nil || IsNil(o.SensorZone) {
		var ret string
		return ret
	}
	return *o.SensorZone
}

// GetSensorZoneOk returns a tuple with the SensorZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseDataObjectsInner) GetSensorZoneOk() (*string, bool) {
	if o == nil || IsNil(o.SensorZone) {
		return nil, false
	}
	return o.SensorZone, true
}

// HasSensorZone returns a boolean if a field has been set.
func (o *GetEntitiesResponseDataObjectsInner) HasSensorZone() bool {
	if o != nil && !IsNil(o.SensorZone) {
		return true
	}

	return false
}

// SetSensorZone gets a reference to the given string and assigns it to the SensorZone field.
func (o *GetEntitiesResponseDataObjectsInner) SetSensorZone(v string) {
	o.SensorZone = &v
}

// GetValue returns the Value field value
func (o *GetEntitiesResponseDataObjectsInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseDataObjectsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetEntitiesResponseDataObjectsInner) SetValue(v string) {
	o.Value = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *GetEntitiesResponseDataObjectsInner) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseDataObjectsInner) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *GetEntitiesResponseDataObjectsInner) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *GetEntitiesResponseDataObjectsInner) SetHostname(v string) {
	o.Hostname = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *GetEntitiesResponseDataObjectsInner) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEntitiesResponseDataObjectsInner) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *GetEntitiesResponseDataObjectsInner) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *GetEntitiesResponseDataObjectsInner) SetMacAddress(v string) {
	o.MacAddress = &v
}

func (o GetEntitiesResponseDataObjectsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEntitiesResponseDataObjectsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["activityScore"] = o.ActivityScore
	if !IsNil(o.LastSeen) {
		toSerialize["lastSeen"] = o.LastSeen
	}
	if !IsNil(o.FirstSeen) {
		toSerialize["firstSeen"] = o.FirstSeen
	}
	toSerialize["isSuppressed"] = o.IsSuppressed
	toSerialize["name"] = o.Name
	toSerialize["inventory"] = o.Inventory
	toSerialize["tags"] = o.Tags
	if !IsNil(o.Criticality) {
		toSerialize["criticality"] = o.Criticality
	}
	if !IsNil(o.Reputation) {
		toSerialize["reputation"] = o.Reputation
	}
	toSerialize["recentSignalSeverity"] = o.RecentSignalSeverity
	toSerialize["entityType"] = o.EntityType
	if !IsNil(o.SensorZone) {
		toSerialize["sensorZone"] = o.SensorZone
	}
	toSerialize["value"] = o.Value
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.MacAddress) {
		toSerialize["macAddress"] = o.MacAddress
	}
	return toSerialize, nil
}

func (o *GetEntitiesResponseDataObjectsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"activityScore",
		"isSuppressed",
		"name",
		"inventory",
		"tags",
		"recentSignalSeverity",
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

	varGetEntitiesResponseDataObjectsInner := _GetEntitiesResponseDataObjectsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEntitiesResponseDataObjectsInner)

	if err != nil {
		return err
	}

	*o = GetEntitiesResponseDataObjectsInner(varGetEntitiesResponseDataObjectsInner)

	return err
}

type NullableGetEntitiesResponseDataObjectsInner struct {
	value *GetEntitiesResponseDataObjectsInner
	isSet bool
}

func (v NullableGetEntitiesResponseDataObjectsInner) Get() *GetEntitiesResponseDataObjectsInner {
	return v.value
}

func (v *NullableGetEntitiesResponseDataObjectsInner) Set(val *GetEntitiesResponseDataObjectsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEntitiesResponseDataObjectsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEntitiesResponseDataObjectsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEntitiesResponseDataObjectsInner(val *GetEntitiesResponseDataObjectsInner) *NullableGetEntitiesResponseDataObjectsInner {
	return &NullableGetEntitiesResponseDataObjectsInner{value: val, isSet: true}
}

func (v NullableGetEntitiesResponseDataObjectsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEntitiesResponseDataObjectsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

