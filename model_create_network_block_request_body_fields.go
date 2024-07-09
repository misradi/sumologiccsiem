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

// checks if the CreateNetworkBlockRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkBlockRequestBodyFields{}

// CreateNetworkBlockRequestBodyFields struct for CreateNetworkBlockRequestBodyFields
type CreateNetworkBlockRequestBodyFields struct {
	AddressBlock string `json:"addressBlock"`
	Label *string `json:"label,omitempty"`
	Internal *bool `json:"internal,omitempty"`
	SuppressesSignals *bool `json:"suppressesSignals,omitempty"`
}

type _CreateNetworkBlockRequestBodyFields CreateNetworkBlockRequestBodyFields

// NewCreateNetworkBlockRequestBodyFields instantiates a new CreateNetworkBlockRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkBlockRequestBodyFields(addressBlock string) *CreateNetworkBlockRequestBodyFields {
	this := CreateNetworkBlockRequestBodyFields{}
	this.AddressBlock = addressBlock
	return &this
}

// NewCreateNetworkBlockRequestBodyFieldsWithDefaults instantiates a new CreateNetworkBlockRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkBlockRequestBodyFieldsWithDefaults() *CreateNetworkBlockRequestBodyFields {
	this := CreateNetworkBlockRequestBodyFields{}
	return &this
}

// GetAddressBlock returns the AddressBlock field value
func (o *CreateNetworkBlockRequestBodyFields) GetAddressBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressBlock
}

// GetAddressBlockOk returns a tuple with the AddressBlock field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkBlockRequestBodyFields) GetAddressBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressBlock, true
}

// SetAddressBlock sets field value
func (o *CreateNetworkBlockRequestBodyFields) SetAddressBlock(v string) {
	o.AddressBlock = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CreateNetworkBlockRequestBodyFields) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkBlockRequestBodyFields) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CreateNetworkBlockRequestBodyFields) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CreateNetworkBlockRequestBodyFields) SetLabel(v string) {
	o.Label = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *CreateNetworkBlockRequestBodyFields) GetInternal() bool {
	if o == nil || IsNil(o.Internal) {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkBlockRequestBodyFields) GetInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *CreateNetworkBlockRequestBodyFields) HasInternal() bool {
	if o != nil && !IsNil(o.Internal) {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *CreateNetworkBlockRequestBodyFields) SetInternal(v bool) {
	o.Internal = &v
}

// GetSuppressesSignals returns the SuppressesSignals field value if set, zero value otherwise.
func (o *CreateNetworkBlockRequestBodyFields) GetSuppressesSignals() bool {
	if o == nil || IsNil(o.SuppressesSignals) {
		var ret bool
		return ret
	}
	return *o.SuppressesSignals
}

// GetSuppressesSignalsOk returns a tuple with the SuppressesSignals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkBlockRequestBodyFields) GetSuppressesSignalsOk() (*bool, bool) {
	if o == nil || IsNil(o.SuppressesSignals) {
		return nil, false
	}
	return o.SuppressesSignals, true
}

// HasSuppressesSignals returns a boolean if a field has been set.
func (o *CreateNetworkBlockRequestBodyFields) HasSuppressesSignals() bool {
	if o != nil && !IsNil(o.SuppressesSignals) {
		return true
	}

	return false
}

// SetSuppressesSignals gets a reference to the given bool and assigns it to the SuppressesSignals field.
func (o *CreateNetworkBlockRequestBodyFields) SetSuppressesSignals(v bool) {
	o.SuppressesSignals = &v
}

func (o CreateNetworkBlockRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkBlockRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addressBlock"] = o.AddressBlock
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Internal) {
		toSerialize["internal"] = o.Internal
	}
	if !IsNil(o.SuppressesSignals) {
		toSerialize["suppressesSignals"] = o.SuppressesSignals
	}
	return toSerialize, nil
}

func (o *CreateNetworkBlockRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addressBlock",
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

	varCreateNetworkBlockRequestBodyFields := _CreateNetworkBlockRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNetworkBlockRequestBodyFields)

	if err != nil {
		return err
	}

	*o = CreateNetworkBlockRequestBodyFields(varCreateNetworkBlockRequestBodyFields)

	return err
}

type NullableCreateNetworkBlockRequestBodyFields struct {
	value *CreateNetworkBlockRequestBodyFields
	isSet bool
}

func (v NullableCreateNetworkBlockRequestBodyFields) Get() *CreateNetworkBlockRequestBodyFields {
	return v.value
}

func (v *NullableCreateNetworkBlockRequestBodyFields) Set(val *CreateNetworkBlockRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkBlockRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkBlockRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkBlockRequestBodyFields(val *CreateNetworkBlockRequestBodyFields) *NullableCreateNetworkBlockRequestBodyFields {
	return &NullableCreateNetworkBlockRequestBodyFields{value: val, isSet: true}
}

func (v NullableCreateNetworkBlockRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkBlockRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


