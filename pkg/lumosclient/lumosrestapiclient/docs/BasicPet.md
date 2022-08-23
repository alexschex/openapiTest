# BasicPet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of the pet | 
**Type** | [**PetTypes**](PetTypes.md) |  | 
**Age** | **int32** | age of the pet | 

## Methods

### NewBasicPet

`func NewBasicPet(name string, type_ PetTypes, age int32, ) *BasicPet`

NewBasicPet instantiates a new BasicPet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicPetWithDefaults

`func NewBasicPetWithDefaults() *BasicPet`

NewBasicPetWithDefaults instantiates a new BasicPet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BasicPet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicPet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicPet) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BasicPet) GetType() PetTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BasicPet) GetTypeOk() (*PetTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BasicPet) SetType(v PetTypes)`

SetType sets Type field to given value.


### GetAge

`func (o *BasicPet) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *BasicPet) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *BasicPet) SetAge(v int32)`

SetAge sets Age field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


