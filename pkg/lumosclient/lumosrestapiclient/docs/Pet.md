# Pet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of the pet | 
**Type** | [**PetTypes**](PetTypes.md) |  | 
**Age** | **int32** | age of the pet | 
**Bark** | **bool** |  | 
**Hunts** | **bool** |  | 

## Methods

### NewPet

`func NewPet(name string, type_ PetTypes, age int32, bark bool, hunts bool, ) *Pet`

NewPet instantiates a new Pet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPetWithDefaults

`func NewPetWithDefaults() *Pet`

NewPetWithDefaults instantiates a new Pet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Pet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pet) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Pet) GetType() PetTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Pet) GetTypeOk() (*PetTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Pet) SetType(v PetTypes)`

SetType sets Type field to given value.


### GetAge

`func (o *Pet) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Pet) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Pet) SetAge(v int32)`

SetAge sets Age field to given value.


### GetBark

`func (o *Pet) GetBark() bool`

GetBark returns the Bark field if non-nil, zero value otherwise.

### GetBarkOk

`func (o *Pet) GetBarkOk() (*bool, bool)`

GetBarkOk returns a tuple with the Bark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBark

`func (o *Pet) SetBark(v bool)`

SetBark sets Bark field to given value.


### GetHunts

`func (o *Pet) GetHunts() bool`

GetHunts returns the Hunts field if non-nil, zero value otherwise.

### GetHuntsOk

`func (o *Pet) GetHuntsOk() (*bool, bool)`

GetHuntsOk returns a tuple with the Hunts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHunts

`func (o *Pet) SetHunts(v bool)`

SetHunts sets Hunts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


