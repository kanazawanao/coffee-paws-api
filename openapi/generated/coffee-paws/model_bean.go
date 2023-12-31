/*
coffee-paws api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Bean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bean{}

// Bean struct for Bean
type Bean struct {
	Id string `json:"id"`
	StoreId string `json:"storeId"`
	ProductionArea string `json:"productionArea"`
	PlantationName string `json:"plantationName"`
	Kind string `json:"kind"`
	RoastLevel string `json:"roastLevel"`
	Price string `json:"price"`
}

// NewBean instantiates a new Bean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBean(id string, storeId string, productionArea string, plantationName string, kind string, roastLevel string, price string) *Bean {
	this := Bean{}
	this.Id = id
	this.StoreId = storeId
	this.ProductionArea = productionArea
	this.PlantationName = plantationName
	this.Kind = kind
	this.RoastLevel = roastLevel
	this.Price = price
	return &this
}

// NewBeanWithDefaults instantiates a new Bean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeanWithDefaults() *Bean {
	this := Bean{}
	return &this
}

// GetId returns the Id field value
func (o *Bean) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Bean) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Bean) SetId(v string) {
	o.Id = v
}

// GetStoreId returns the StoreId field value
func (o *Bean) GetStoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value
// and a boolean to check if the value has been set.
func (o *Bean) GetStoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoreId, true
}

// SetStoreId sets field value
func (o *Bean) SetStoreId(v string) {
	o.StoreId = v
}

// GetProductionArea returns the ProductionArea field value
func (o *Bean) GetProductionArea() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductionArea
}

// GetProductionAreaOk returns a tuple with the ProductionArea field value
// and a boolean to check if the value has been set.
func (o *Bean) GetProductionAreaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductionArea, true
}

// SetProductionArea sets field value
func (o *Bean) SetProductionArea(v string) {
	o.ProductionArea = v
}

// GetPlantationName returns the PlantationName field value
func (o *Bean) GetPlantationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlantationName
}

// GetPlantationNameOk returns a tuple with the PlantationName field value
// and a boolean to check if the value has been set.
func (o *Bean) GetPlantationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlantationName, true
}

// SetPlantationName sets field value
func (o *Bean) SetPlantationName(v string) {
	o.PlantationName = v
}

// GetKind returns the Kind field value
func (o *Bean) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *Bean) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *Bean) SetKind(v string) {
	o.Kind = v
}

// GetRoastLevel returns the RoastLevel field value
func (o *Bean) GetRoastLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoastLevel
}

// GetRoastLevelOk returns a tuple with the RoastLevel field value
// and a boolean to check if the value has been set.
func (o *Bean) GetRoastLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoastLevel, true
}

// SetRoastLevel sets field value
func (o *Bean) SetRoastLevel(v string) {
	o.RoastLevel = v
}

// GetPrice returns the Price field value
func (o *Bean) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *Bean) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *Bean) SetPrice(v string) {
	o.Price = v
}

func (o Bean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Bean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["storeId"] = o.StoreId
	toSerialize["productionArea"] = o.ProductionArea
	toSerialize["plantationName"] = o.PlantationName
	toSerialize["kind"] = o.Kind
	toSerialize["roastLevel"] = o.RoastLevel
	toSerialize["price"] = o.Price
	return toSerialize, nil
}

type NullableBean struct {
	value *Bean
	isSet bool
}

func (v NullableBean) Get() *Bean {
	return v.value
}

func (v *NullableBean) Set(val *Bean) {
	v.value = val
	v.isSet = true
}

func (v NullableBean) IsSet() bool {
	return v.isSet
}

func (v *NullableBean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBean(val *Bean) *NullableBean {
	return &NullableBean{value: val, isSet: true}
}

func (v NullableBean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


