/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spacetradersapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the RepairTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepairTransaction{}

// RepairTransaction Result of a repair transaction.
type RepairTransaction struct {
	// The symbol of the waypoint.
	WaypointSymbol string `json:"waypointSymbol"`
	// The symbol of the ship.
	ShipSymbol string `json:"shipSymbol"`
	// The total price of the transaction.
	TotalPrice int32 `json:"totalPrice"`
	// The timestamp of the transaction.
	Timestamp time.Time `json:"timestamp"`
}

type _RepairTransaction RepairTransaction

// NewRepairTransaction instantiates a new RepairTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepairTransaction(waypointSymbol string, shipSymbol string, totalPrice int32, timestamp time.Time) *RepairTransaction {
	this := RepairTransaction{}
	this.WaypointSymbol = waypointSymbol
	this.ShipSymbol = shipSymbol
	this.TotalPrice = totalPrice
	this.Timestamp = timestamp
	return &this
}

// NewRepairTransactionWithDefaults instantiates a new RepairTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepairTransactionWithDefaults() *RepairTransaction {
	this := RepairTransaction{}
	return &this
}

// GetWaypointSymbol returns the WaypointSymbol field value
func (o *RepairTransaction) GetWaypointSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WaypointSymbol
}

// GetWaypointSymbolOk returns a tuple with the WaypointSymbol field value
// and a boolean to check if the value has been set.
func (o *RepairTransaction) GetWaypointSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaypointSymbol, true
}

// SetWaypointSymbol sets field value
func (o *RepairTransaction) SetWaypointSymbol(v string) {
	o.WaypointSymbol = v
}

// GetShipSymbol returns the ShipSymbol field value
func (o *RepairTransaction) GetShipSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipSymbol
}

// GetShipSymbolOk returns a tuple with the ShipSymbol field value
// and a boolean to check if the value has been set.
func (o *RepairTransaction) GetShipSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipSymbol, true
}

// SetShipSymbol sets field value
func (o *RepairTransaction) SetShipSymbol(v string) {
	o.ShipSymbol = v
}

// GetTotalPrice returns the TotalPrice field value
func (o *RepairTransaction) GetTotalPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value
// and a boolean to check if the value has been set.
func (o *RepairTransaction) GetTotalPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPrice, true
}

// SetTotalPrice sets field value
func (o *RepairTransaction) SetTotalPrice(v int32) {
	o.TotalPrice = v
}

// GetTimestamp returns the Timestamp field value
func (o *RepairTransaction) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *RepairTransaction) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *RepairTransaction) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o RepairTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepairTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["waypointSymbol"] = o.WaypointSymbol
	toSerialize["shipSymbol"] = o.ShipSymbol
	toSerialize["totalPrice"] = o.TotalPrice
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

func (o *RepairTransaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"waypointSymbol",
		"shipSymbol",
		"totalPrice",
		"timestamp",
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

	varRepairTransaction := _RepairTransaction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRepairTransaction)

	if err != nil {
		return err
	}

	*o = RepairTransaction(varRepairTransaction)

	return err
}

type NullableRepairTransaction struct {
	value *RepairTransaction
	isSet bool
}

func (v NullableRepairTransaction) Get() *RepairTransaction {
	return v.value
}

func (v *NullableRepairTransaction) Set(val *RepairTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableRepairTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableRepairTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepairTransaction(val *RepairTransaction) *NullableRepairTransaction {
	return &NullableRepairTransaction{value: val, isSet: true}
}

func (v NullableRepairTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepairTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


