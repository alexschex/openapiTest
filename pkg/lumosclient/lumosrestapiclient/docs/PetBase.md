# PetBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of the pet | 
**Type** | [**PetTypes**](PetTypes.md) |  | 
**Age** | **int32** | age of the pet | 

## Methods

### NewPetBase

`func NewPetBase(name string, type_ PetTypes, age int32, ) *PetBase`

NewPetBase instantiates a new PetBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPetBaseWithDefaults

`func NewPetBaseWithDefaults() *PetBase`

NewPetBaseWithDefaults instantiates a new PetBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PetBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PetBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PetBase) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *PetBase) GetType() PetTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PetBase) GetTypeOk() (*PetTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PetBase) SetType(v PetTypes)`

SetType sets Type field to given value.


### GetAge

`func (o *PetBase) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *PetBase) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *PetBase) SetAge(v int32)`

SetAge sets Age field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


