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

// checks if the SuppressListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuppressListItem{}

// SuppressListItem struct for SuppressListItem
type SuppressListItem struct {
	Id string `json:"id"`
	Value string `json:"value"`
	Active bool `json:"active"`
	Expiration *time.Time `json:"expiration,omitempty"`
	Meta *GetThreatIntelIndicatorResponseDataMeta `json:"meta,omitempty"`
	ListName string `json:"listName"`
}

type _SuppressListItem SuppressListItem

// NewSuppressListItem instantiates a new SuppressListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuppressListItem(id string, value string, active bool, listName string) *SuppressListItem {
	this := SuppressListItem{}
	this.Id = id
	this.Value = value
	this.Active = active
	this.ListName = listName
	return &this
}

// NewSuppressListItemWithDefaults instantiates a new SuppressListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuppressListItemWithDefaults() *SuppressListItem {
	this := SuppressListItem{}
	return &this
}

// GetId returns the Id field value
func (o *SuppressListItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SuppressListItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SuppressListItem) SetId(v string) {
	o.Id = v
}

// GetValue returns the Value field value
func (o *SuppressListItem) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SuppressListItem) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SuppressListItem) SetValue(v string) {
	o.Value = v
}

// GetActive returns the Active field value
func (o *SuppressListItem) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *SuppressListItem) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *SuppressListItem) SetActive(v bool) {
	o.Active = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *SuppressListItem) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration) {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuppressListItem) GetExpirationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *SuppressListItem) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *SuppressListItem) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SuppressListItem) GetMeta() GetThreatIntelIndicatorResponseDataMeta {
	if o == nil || IsNil(o.Meta) {
		var ret GetThreatIntelIndicatorResponseDataMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuppressListItem) GetMetaOk() (*GetThreatIntelIndicatorResponseDataMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SuppressListItem) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given GetThreatIntelIndicatorResponseDataMeta and assigns it to the Meta field.
func (o *SuppressListItem) SetMeta(v GetThreatIntelIndicatorResponseDataMeta) {
	o.Meta = &v
}

// GetListName returns the ListName field value
func (o *SuppressListItem) GetListName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListName
}

// GetListNameOk returns a tuple with the ListName field value
// and a boolean to check if the value has been set.
func (o *SuppressListItem) GetListNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListName, true
}

// SetListName sets field value
func (o *SuppressListItem) SetListName(v string) {
	o.ListName = v
}

func (o SuppressListItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuppressListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["value"] = o.Value
	toSerialize["active"] = o.Active
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	toSerialize["listName"] = o.ListName
	return toSerialize, nil
}

func (o *SuppressListItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"value",
		"active",
		"listName",
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

	varSuppressListItem := _SuppressListItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSuppressListItem)

	if err != nil {
		return err
	}

	*o = SuppressListItem(varSuppressListItem)

	return err
}

type NullableSuppressListItem struct {
	value *SuppressListItem
	isSet bool
}

func (v NullableSuppressListItem) Get() *SuppressListItem {
	return v.value
}

func (v *NullableSuppressListItem) Set(val *SuppressListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSuppressListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSuppressListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuppressListItem(val *SuppressListItem) *NullableSuppressListItem {
	return &NullableSuppressListItem{value: val, isSet: true}
}

func (v NullableSuppressListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuppressListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


