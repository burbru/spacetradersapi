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

// checks if the ShipNavRoute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipNavRoute{}

// ShipNavRoute The routing information for the ship's most recent transit or current location.
type ShipNavRoute struct {
	Destination ShipNavRouteWaypoint `json:"destination"`
	Origin ShipNavRouteWaypoint `json:"origin"`
	// The date time of the ship's departure.
	DepartureTime time.Time `json:"departureTime"`
	// The date time of the ship's arrival. If the ship is in-transit, this is the expected time of arrival.
	Arrival time.Time `json:"arrival"`
}

type _ShipNavRoute ShipNavRoute

// NewShipNavRoute instantiates a new ShipNavRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipNavRoute(destination ShipNavRouteWaypoint, origin ShipNavRouteWaypoint, departureTime time.Time, arrival time.Time) *ShipNavRoute {
	this := ShipNavRoute{}
	this.Destination = destination
	this.Origin = origin
	this.DepartureTime = departureTime
	this.Arrival = arrival
	return &this
}

// NewShipNavRouteWithDefaults instantiates a new ShipNavRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipNavRouteWithDefaults() *ShipNavRoute {
	this := ShipNavRoute{}
	return &this
}

// GetDestination returns the Destination field value
func (o *ShipNavRoute) GetDestination() ShipNavRouteWaypoint {
	if o == nil {
		var ret ShipNavRouteWaypoint
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *ShipNavRoute) GetDestinationOk() (*ShipNavRouteWaypoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *ShipNavRoute) SetDestination(v ShipNavRouteWaypoint) {
	o.Destination = v
}

// GetOrigin returns the Origin field value
func (o *ShipNavRoute) GetOrigin() ShipNavRouteWaypoint {
	if o == nil {
		var ret ShipNavRouteWaypoint
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *ShipNavRoute) GetOriginOk() (*ShipNavRouteWaypoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *ShipNavRoute) SetOrigin(v ShipNavRouteWaypoint) {
	o.Origin = v
}

// GetDepartureTime returns the DepartureTime field value
func (o *ShipNavRoute) GetDepartureTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DepartureTime
}

// GetDepartureTimeOk returns a tuple with the DepartureTime field value
// and a boolean to check if the value has been set.
func (o *ShipNavRoute) GetDepartureTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepartureTime, true
}

// SetDepartureTime sets field value
func (o *ShipNavRoute) SetDepartureTime(v time.Time) {
	o.DepartureTime = v
}

// GetArrival returns the Arrival field value
func (o *ShipNavRoute) GetArrival() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Arrival
}

// GetArrivalOk returns a tuple with the Arrival field value
// and a boolean to check if the value has been set.
func (o *ShipNavRoute) GetArrivalOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arrival, true
}

// SetArrival sets field value
func (o *ShipNavRoute) SetArrival(v time.Time) {
	o.Arrival = v
}

func (o ShipNavRoute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipNavRoute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination"] = o.Destination
	toSerialize["origin"] = o.Origin
	toSerialize["departureTime"] = o.DepartureTime
	toSerialize["arrival"] = o.Arrival
	return toSerialize, nil
}

func (o *ShipNavRoute) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination",
		"origin",
		"departureTime",
		"arrival",
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

	varShipNavRoute := _ShipNavRoute{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipNavRoute)

	if err != nil {
		return err
	}

	*o = ShipNavRoute(varShipNavRoute)

	return err
}

type NullableShipNavRoute struct {
	value *ShipNavRoute
	isSet bool
}

func (v NullableShipNavRoute) Get() *ShipNavRoute {
	return v.value
}

func (v *NullableShipNavRoute) Set(val *ShipNavRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableShipNavRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableShipNavRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipNavRoute(val *ShipNavRoute) *NullableShipNavRoute {
	return &NullableShipNavRoute{value: val, isSet: true}
}

func (v NullableShipNavRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipNavRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


