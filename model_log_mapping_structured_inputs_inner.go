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

// checks if the LogMappingStructuredInputsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogMappingStructuredInputsInner{}

// LogMappingStructuredInputsInner struct for LogMappingStructuredInputsInner
type LogMappingStructuredInputsInner struct {
	Vendor string `json:"vendor"`
	Product string `json:"product"`
	EventIdPattern string `json:"eventIdPattern"`
	LogFormat string `json:"logFormat"`
}

type _LogMappingStructuredInputsInner LogMappingStructuredInputsInner

// NewLogMappingStructuredInputsInner instantiates a new LogMappingStructuredInputsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogMappingStructuredInputsInner(vendor string, product string, eventIdPattern string, logFormat string) *LogMappingStructuredInputsInner {
	this := LogMappingStructuredInputsInner{}
	this.Vendor = vendor
	this.Product = product
	this.EventIdPattern = eventIdPattern
	this.LogFormat = logFormat
	return &this
}

// NewLogMappingStructuredInputsInnerWithDefaults instantiates a new LogMappingStructuredInputsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogMappingStructuredInputsInnerWithDefaults() *LogMappingStructuredInputsInner {
	this := LogMappingStructuredInputsInner{}
	return &this
}

// GetVendor returns the Vendor field value
func (o *LogMappingStructuredInputsInner) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *LogMappingStructuredInputsInner) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *LogMappingStructuredInputsInner) SetVendor(v string) {
	o.Vendor = v
}

// GetProduct returns the Product field value
func (o *LogMappingStructuredInputsInner) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *LogMappingStructuredInputsInner) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *LogMappingStructuredInputsInner) SetProduct(v string) {
	o.Product = v
}

// GetEventIdPattern returns the EventIdPattern field value
func (o *LogMappingStructuredInputsInner) GetEventIdPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventIdPattern
}

// GetEventIdPatternOk returns a tuple with the EventIdPattern field value
// and a boolean to check if the value has been set.
func (o *LogMappingStructuredInputsInner) GetEventIdPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIdPattern, true
}

// SetEventIdPattern sets field value
func (o *LogMappingStructuredInputsInner) SetEventIdPattern(v string) {
	o.EventIdPattern = v
}

// GetLogFormat returns the LogFormat field value
func (o *LogMappingStructuredInputsInner) GetLogFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFormat
}

// GetLogFormatOk returns a tuple with the LogFormat field value
// and a boolean to check if the value has been set.
func (o *LogMappingStructuredInputsInner) GetLogFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogFormat, true
}

// SetLogFormat sets field value
func (o *LogMappingStructuredInputsInner) SetLogFormat(v string) {
	o.LogFormat = v
}

func (o LogMappingStructuredInputsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogMappingStructuredInputsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vendor"] = o.Vendor
	toSerialize["product"] = o.Product
	toSerialize["eventIdPattern"] = o.EventIdPattern
	toSerialize["logFormat"] = o.LogFormat
	return toSerialize, nil
}

func (o *LogMappingStructuredInputsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vendor",
		"product",
		"eventIdPattern",
		"logFormat",
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

	varLogMappingStructuredInputsInner := _LogMappingStructuredInputsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLogMappingStructuredInputsInner)

	if err != nil {
		return err
	}

	*o = LogMappingStructuredInputsInner(varLogMappingStructuredInputsInner)

	return err
}

type NullableLogMappingStructuredInputsInner struct {
	value *LogMappingStructuredInputsInner
	isSet bool
}

func (v NullableLogMappingStructuredInputsInner) Get() *LogMappingStructuredInputsInner {
	return v.value
}

func (v *NullableLogMappingStructuredInputsInner) Set(val *LogMappingStructuredInputsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableLogMappingStructuredInputsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableLogMappingStructuredInputsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogMappingStructuredInputsInner(val *LogMappingStructuredInputsInner) *NullableLogMappingStructuredInputsInner {
	return &NullableLogMappingStructuredInputsInner{value: val, isSet: true}
}

func (v NullableLogMappingStructuredInputsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogMappingStructuredInputsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


