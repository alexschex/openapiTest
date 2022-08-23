# Dog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of the pet | 
**Type** | [**PetTypes**](PetTypes.md) |  | 
**Age** | **int32** | age of the pet | 
**Bark** | **bool** |  | 

## Methods

### NewDog

`func NewDog(name string, type_ PetTypes, age int32, bark bool, ) *Dog`

NewDog instantiates a new Dog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDogWithDefaults

`func NewDogWithDefaults() *Dog`

NewDogWithDefaults instantiates a new Dog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Dog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dog) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Dog) GetType() PetTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Dog) GetTypeOk() (*PetTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Dog) SetType(v PetTypes)`

SetType sets Type field to given value.


### GetAge

`func (o *Dog) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Dog) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Dog) SetAge(v int32)`

SetAge sets Age field to given value.


### GetBark

`func (o *Dog) GetBark() bool`

GetBark returns the Bark field if non-nil, zero value otherwise.

### GetBarkOk

`func (o *Dog) GetBarkOk() (*bool, bool)`

GetBarkOk returns a tuple with the Bark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBark

`func (o *Dog) SetBark(v bool)`

SetBark sets Bark field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


