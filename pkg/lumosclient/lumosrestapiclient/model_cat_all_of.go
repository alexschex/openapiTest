/*
Sample API

Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML.

API version: 0.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CatAllOf struct for CatAllOf
type CatAllOf struct {
	Hunts bool `json:"hunts"`
}

// NewCatAllOf instantiates a new CatAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatAllOf(hunts bool) *CatAllOf {
	this := CatAllOf{}
	this.Hunts = hunts
	return &this
}

// NewCatAllOfWithDefaults instantiates a new CatAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatAllOfWithDefaults() *CatAllOf {
	this := CatAllOf{}
	return &this
}

// GetHunts returns the Hunts field value
func (o *CatAllOf) GetHunts() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hunts
}

// GetHuntsOk returns a tuple with the Hunts field value
// and a boolean to check if the value has been set.
func (o *CatAllOf) GetHuntsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hunts, true
}

// SetHunts sets field value
func (o *CatAllOf) SetHunts(v bool) {
	o.Hunts = v
}

func (o CatAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hunts"] = o.Hunts
	}
	return json.Marshal(toSerialize)
}

type NullableCatAllOf struct {
	value *CatAllOf
	isSet bool
}

func (v NullableCatAllOf) Get() *CatAllOf {
	return v.value
}

func (v *NullableCatAllOf) Set(val *CatAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCatAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCatAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatAllOf(val *CatAllOf) *NullableCatAllOf {
	return &NullableCatAllOf{value: val, isSet: true}
}

func (v NullableCatAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


