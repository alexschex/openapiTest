# Cat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of the pet | 
**Type** | [**PetTypes**](PetTypes.md) |  | 
**Age** | **int32** | age of the pet | 
**Hunts** | **bool** |  | 

## Methods

### NewCat

`func NewCat(name string, type_ PetTypes, age int32, hunts bool, ) *Cat`

NewCat instantiates a new Cat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatWithDefaults

`func NewCatWithDefaults() *Cat`

NewCatWithDefaults instantiates a new Cat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Cat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cat) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Cat) GetType() PetTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Cat) GetTypeOk() (*PetTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Cat) SetType(v PetTypes)`

SetType sets Type field to given value.


### GetAge

`func (o *Cat) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Cat) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Cat) SetAge(v int32)`

SetAge sets Age field to given value.


### GetHunts

`func (o *Cat) GetHunts() bool`

GetHunts returns the Hunts field if non-nil, zero value otherwise.

### GetHuntsOk

`func (o *Cat) GetHuntsOk() (*bool, bool)`

GetHuntsOk returns a tuple with the Hunts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHunts

`func (o *Cat) SetHunts(v bool)`

SetHunts sets Hunts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


