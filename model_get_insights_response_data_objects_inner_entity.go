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

// checks if the GetInsightsResponseDataObjectsInnerEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsightsResponseDataObjectsInnerEntity{}

// GetInsightsResponseDataObjectsInnerEntity The  primary Entity associated with this Insight
type GetInsightsResponseDataObjectsInnerEntity struct {
	Id string `json:"id"`
	EntityType string `json:"entityType"`
	Name string `json:"name"`
	Value string `json:"value"`
	Hostname *string `json:"hostname,omitempty"`
	MacAddress *string `json:"macAddress,omitempty"`
	SensorZone *string `json:"sensorZone,omitempty"`
}

type _GetInsightsResponseDataObjectsInnerEntity GetInsightsResponseDataObjectsInnerEntity

// NewGetInsightsResponseDataObjectsInnerEntity instantiates a new GetInsightsResponseDataObjectsInnerEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsightsResponseDataObjectsInnerEntity(id string, entityType string, name string, value string) *GetInsightsResponseDataObjectsInnerEntity {
	this := GetInsightsResponseDataObjectsInnerEntity{}
	this.Id = id
	this.EntityType = entityType
	this.Name = name
	this.Value = value
	return &this
}

// NewGetInsightsResponseDataObjectsInnerEntityWithDefaults instantiates a new GetInsightsResponseDataObjectsInnerEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsightsResponseDataObjectsInnerEntityWithDefaults() *GetInsightsResponseDataObjectsInnerEntity {
	this := GetInsightsResponseDataObjectsInnerEntity{}
	return &this
}

// GetId returns the Id field value
func (o *GetInsightsResponseDataObjectsInnerEntity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerEntity) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetInsightsResponseDataObjectsInnerEntity) SetId(v string) {
	o.Id = v
}

// GetEntityType returns the EntityType field value
func (o *GetInsightsResponseDataObjectsInnerEntity) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerEntity) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *GetInsightsResponseDataObjectsInnerEntity) SetEntityType(v string) {
	o.EntityType = v
}

// GetName returns the Name field value
func (o *GetInsightsResponseDataObjectsInnerEntity) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerEntity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetInsightsResponseDataObjectsInnerEntity) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *GetInsightsResponseDataObjectsInnerEntity) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerEntity) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetInsightsResponseDataObjectsInnerEntity) SetValue(v string) {
	o.Value = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *GetInsightsResponseDataObjectsInnerEntity) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerEntity) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *GetInsightsResponseDataObjectsInnerEntity) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *GetInsightsResponseDataObjectsInnerEntity) SetHostname(v string) {
	o.Hostname = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *GetInsightsResponseDataObjectsInnerEntity) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerEntity) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *GetInsightsResponseDataObjectsInnerEntity) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *GetInsightsResponseDataObjectsInnerEntity) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetSensorZone returns the SensorZone field value if set, zero value otherwise.
func (o *GetInsightsResponseDataObjectsInnerEntity) GetSensorZone() string {
	if o == nil || IsNil(o.SensorZone) {
		var ret string
		return ret
	}
	return *o.SensorZone
}

// GetSensorZoneOk returns a tuple with the SensorZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInsightsResponseDataObjectsInnerEntity) GetSensorZoneOk() (*string, bool) {
	if o == nil || IsNil(o.SensorZone) {
		return nil, false
	}
	return o.SensorZone, true
}

// HasSensorZone returns a boolean if a field has been set.
func (o *GetInsightsResponseDataObjectsInnerEntity) HasSensorZone() bool {
	if o != nil && !IsNil(o.SensorZone) {
		return true
	}

	return false
}

// SetSensorZone gets a reference to the given string and assigns it to the SensorZone field.
func (o *GetInsightsResponseDataObjectsInnerEntity) SetSensorZone(v string) {
	o.SensorZone = &v
}

func (o GetInsightsResponseDataObjectsInnerEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsightsResponseDataObjectsInnerEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["entityType"] = o.EntityType
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.MacAddress) {
		toSerialize["macAddress"] = o.MacAddress
	}
	if !IsNil(o.SensorZone) {
		toSerialize["sensorZone"] = o.SensorZone
	}
	return toSerialize, nil
}

func (o *GetInsightsResponseDataObjectsInnerEntity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"entityType",
		"name",
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

	varGetInsightsResponseDataObjectsInnerEntity := _GetInsightsResponseDataObjectsInnerEntity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetInsightsResponseDataObjectsInnerEntity)

	if err != nil {
		return err
	}

	*o = GetInsightsResponseDataObjectsInnerEntity(varGetInsightsResponseDataObjectsInnerEntity)

	return err
}

type NullableGetInsightsResponseDataObjectsInnerEntity struct {
	value *GetInsightsResponseDataObjectsInnerEntity
	isSet bool
}

func (v NullableGetInsightsResponseDataObjectsInnerEntity) Get() *GetInsightsResponseDataObjectsInnerEntity {
	return v.value
}

func (v *NullableGetInsightsResponseDataObjectsInnerEntity) Set(val *GetInsightsResponseDataObjectsInnerEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsightsResponseDataObjectsInnerEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsightsResponseDataObjectsInnerEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsightsResponseDataObjectsInnerEntity(val *GetInsightsResponseDataObjectsInnerEntity) *NullableGetInsightsResponseDataObjectsInnerEntity {
	return &NullableGetInsightsResponseDataObjectsInnerEntity{value: val, isSet: true}
}

func (v NullableGetInsightsResponseDataObjectsInnerEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsightsResponseDataObjectsInnerEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


