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

// Cat struct for Cat
type Cat struct {
	// name of the pet
	Name string `json:"name"`
	Type PetTypes `json:"type"`
	// age of the pet
	Age int32 `json:"age"`
	Hunts bool `json:"hunts"`
}

// NewCat instantiates a new Cat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCat(name string, type_ PetTypes, age int32, hunts bool) *Cat {
	this := Cat{}
	this.Name = name
	this.Type = type_
	this.Age = age
	this.Hunts = hunts
	return &this
}

// NewCatWithDefaults instantiates a new Cat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatWithDefaults() *Cat {
	this := Cat{}
	return &this
}

// GetName returns the Name field value
func (o *Cat) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Cat) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Cat) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *Cat) GetType() PetTypes {
	if o == nil {
		var ret PetTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Cat) GetTypeOk() (*PetTypes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Cat) SetType(v PetTypes) {
	o.Type = v
}

// GetAge returns the Age field value
func (o *Cat) GetAge() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Age
}

// GetAgeOk returns a tuple with the Age field value
// and a boolean to check if the value has been set.
func (o *Cat) GetAgeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Age, true
}

// SetAge sets field value
func (o *Cat) SetAge(v int32) {
	o.Age = v
}

// GetHunts returns the Hunts field value
func (o *Cat) GetHunts() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hunts
}

// GetHuntsOk returns a tuple with the Hunts field value
// and a boolean to check if the value has been set.
func (o *Cat) GetHuntsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hunts, true
}

// SetHunts sets field value
func (o *Cat) SetHunts(v bool) {
	o.Hunts = v
}

func (o Cat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["age"] = o.Age
	}
	if true {
		toSerialize["hunts"] = o.Hunts
	}
	return json.Marshal(toSerialize)
}

type NullableCat struct {
	value *Cat
	isSet bool
}

func (v NullableCat) Get() *Cat {
	return v.value
}

func (v *NullableCat) Set(val *Cat) {
	v.value = val
	v.isSet = true
}

func (v NullableCat) IsSet() bool {
	return v.isSet
}

func (v *NullableCat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCat(val *Cat) *NullableCat {
	return &NullableCat{value: val, isSet: true}
}

func (v NullableCat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


