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
	"bytes"
	"fmt"
)

// checks if the Waypoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Waypoint{}

// Waypoint A waypoint is a location that ships can travel to such as a Planet, Moon or Space Station.
type Waypoint struct {
	// The symbol of the waypoint.
	Symbol string `json:"symbol"`
	Type WaypointType `json:"type"`
	// The symbol of the system.
	SystemSymbol string `json:"systemSymbol"`
	// Relative position of the waypoint on the system's x axis. This is not an absolute position in the universe.
	X int32 `json:"x"`
	// Relative position of the waypoint on the system's y axis. This is not an absolute position in the universe.
	Y int32 `json:"y"`
	// Waypoints that orbit this waypoint.
	Orbitals []WaypointOrbital `json:"orbitals"`
	// The symbol of the parent waypoint, if this waypoint is in orbit around another waypoint. Otherwise this value is undefined.
	Orbits *string `json:"orbits,omitempty"`
	Faction *WaypointFaction `json:"faction,omitempty"`
	// The traits of the waypoint.
	Traits []WaypointTrait `json:"traits"`
	// The modifiers of the waypoint.
	Modifiers []WaypointModifier `json:"modifiers,omitempty"`
	Chart *Chart `json:"chart,omitempty"`
	// True if the waypoint is under construction.
	IsUnderConstruction bool `json:"isUnderConstruction"`
}

type _Waypoint Waypoint

// NewWaypoint instantiates a new Waypoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWaypoint(symbol string, type_ WaypointType, systemSymbol string, x int32, y int32, orbitals []WaypointOrbital, traits []WaypointTrait, isUnderConstruction bool) *Waypoint {
	this := Waypoint{}
	this.Symbol = symbol
	this.Type = type_
	this.SystemSymbol = systemSymbol
	this.X = x
	this.Y = y
	this.Orbitals = orbitals
	this.Traits = traits
	this.IsUnderConstruction = isUnderConstruction
	return &this
}

// NewWaypointWithDefaults instantiates a new Waypoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWaypointWithDefaults() *Waypoint {
	this := Waypoint{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *Waypoint) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *Waypoint) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *Waypoint) SetSymbol(v string) {
	o.Symbol = v
}

// GetType returns the Type field value
func (o *Waypoint) GetType() WaypointType {
	if o == nil {
		var ret WaypointType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Waypoint) GetTypeOk() (*WaypointType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Waypoint) SetType(v WaypointType) {
	o.Type = v
}

// GetSystemSymbol returns the SystemSymbol field value
func (o *Waypoint) GetSystemSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemSymbol
}

// GetSystemSymbolOk returns a tuple with the SystemSymbol field value
// and a boolean to check if the value has been set.
func (o *Waypoint) GetSystemSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemSymbol, true
}

// SetSystemSymbol sets field value
func (o *Waypoint) SetSystemSymbol(v string) {
	o.SystemSymbol = v
}

// GetX returns the X field value
func (o *Waypoint) GetX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *Waypoint) GetXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *Waypoint) SetX(v int32) {
	o.X = v
}

// GetY returns the Y field value
func (o *Waypoint) GetY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *Waypoint) GetYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *Waypoint) SetY(v int32) {
	o.Y = v
}

// GetOrbitals returns the Orbitals field value
func (o *Waypoint) GetOrbitals() []WaypointOrbital {
	if o == nil {
		var ret []WaypointOrbital
		return ret
	}

	return o.Orbitals
}

// GetOrbitalsOk returns a tuple with the Orbitals field value
// and a boolean to check if the value has been set.
func (o *Waypoint) GetOrbitalsOk() ([]WaypointOrbital, bool) {
	if o == nil {
		return nil, false
	}
	return o.Orbitals, true
}

// SetOrbitals sets field value
func (o *Waypoint) SetOrbitals(v []WaypointOrbital) {
	o.Orbitals = v
}

// GetOrbits returns the Orbits field value if set, zero value otherwise.
func (o *Waypoint) GetOrbits() string {
	if o == nil || IsNil(o.Orbits) {
		var ret string
		return ret
	}
	return *o.Orbits
}

// GetOrbitsOk returns a tuple with the Orbits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Waypoint) GetOrbitsOk() (*string, bool) {
	if o == nil || IsNil(o.Orbits) {
		return nil, false
	}
	return o.Orbits, true
}

// HasOrbits returns a boolean if a field has been set.
func (o *Waypoint) HasOrbits() bool {
	if o != nil && !IsNil(o.Orbits) {
		return true
	}

	return false
}

// SetOrbits gets a reference to the given string and assigns it to the Orbits field.
func (o *Waypoint) SetOrbits(v string) {
	o.Orbits = &v
}

// GetFaction returns the Faction field value if set, zero value otherwise.
func (o *Waypoint) GetFaction() WaypointFaction {
	if o == nil || IsNil(o.Faction) {
		var ret WaypointFaction
		return ret
	}
	return *o.Faction
}

// GetFactionOk returns a tuple with the Faction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Waypoint) GetFactionOk() (*WaypointFaction, bool) {
	if o == nil || IsNil(o.Faction) {
		return nil, false
	}
	return o.Faction, true
}

// HasFaction returns a boolean if a field has been set.
func (o *Waypoint) HasFaction() bool {
	if o != nil && !IsNil(o.Faction) {
		return true
	}

	return false
}

// SetFaction gets a reference to the given WaypointFaction and assigns it to the Faction field.
func (o *Waypoint) SetFaction(v WaypointFaction) {
	o.Faction = &v
}

// GetTraits returns the Traits field value
func (o *Waypoint) GetTraits() []WaypointTrait {
	if o == nil {
		var ret []WaypointTrait
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
func (o *Waypoint) GetTraitsOk() ([]WaypointTrait, bool) {
	if o == nil {
		return nil, false
	}
	return o.Traits, true
}

// SetTraits sets field value
func (o *Waypoint) SetTraits(v []WaypointTrait) {
	o.Traits = v
}

// GetModifiers returns the Modifiers field value if set, zero value otherwise.
func (o *Waypoint) GetModifiers() []WaypointModifier {
	if o == nil || IsNil(o.Modifiers) {
		var ret []WaypointModifier
		return ret
	}
	return o.Modifiers
}

// GetModifiersOk returns a tuple with the Modifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Waypoint) GetModifiersOk() ([]WaypointModifier, bool) {
	if o == nil || IsNil(o.Modifiers) {
		return nil, false
	}
	return o.Modifiers, true
}

// HasModifiers returns a boolean if a field has been set.
func (o *Waypoint) HasModifiers() bool {
	if o != nil && !IsNil(o.Modifiers) {
		return true
	}

	return false
}

// SetModifiers gets a reference to the given []WaypointModifier and assigns it to the Modifiers field.
func (o *Waypoint) SetModifiers(v []WaypointModifier) {
	o.Modifiers = v
}

// GetChart returns the Chart field value if set, zero value otherwise.
func (o *Waypoint) GetChart() Chart {
	if o == nil || IsNil(o.Chart) {
		var ret Chart
		return ret
	}
	return *o.Chart
}

// GetChartOk returns a tuple with the Chart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Waypoint) GetChartOk() (*Chart, bool) {
	if o == nil || IsNil(o.Chart) {
		return nil, false
	}
	return o.Chart, true
}

// HasChart returns a boolean if a field has been set.
func (o *Waypoint) HasChart() bool {
	if o != nil && !IsNil(o.Chart) {
		return true
	}

	return false
}

// SetChart gets a reference to the given Chart and assigns it to the Chart field.
func (o *Waypoint) SetChart(v Chart) {
	o.Chart = &v
}

// GetIsUnderConstruction returns the IsUnderConstruction field value
func (o *Waypoint) GetIsUnderConstruction() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUnderConstruction
}

// GetIsUnderConstructionOk returns a tuple with the IsUnderConstruction field value
// and a boolean to check if the value has been set.
func (o *Waypoint) GetIsUnderConstructionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUnderConstruction, true
}

// SetIsUnderConstruction sets field value
func (o *Waypoint) SetIsUnderConstruction(v bool) {
	o.IsUnderConstruction = v
}

func (o Waypoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Waypoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["type"] = o.Type
	toSerialize["systemSymbol"] = o.SystemSymbol
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	toSerialize["orbitals"] = o.Orbitals
	if !IsNil(o.Orbits) {
		toSerialize["orbits"] = o.Orbits
	}
	if !IsNil(o.Faction) {
		toSerialize["faction"] = o.Faction
	}
	toSerialize["traits"] = o.Traits
	if !IsNil(o.Modifiers) {
		toSerialize["modifiers"] = o.Modifiers
	}
	if !IsNil(o.Chart) {
		toSerialize["chart"] = o.Chart
	}
	toSerialize["isUnderConstruction"] = o.IsUnderConstruction
	return toSerialize, nil
}

func (o *Waypoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"type",
		"systemSymbol",
		"x",
		"y",
		"orbitals",
		"traits",
		"isUnderConstruction",
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

	varWaypoint := _Waypoint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWaypoint)

	if err != nil {
		return err
	}

	*o = Waypoint(varWaypoint)

	return err
}

type NullableWaypoint struct {
	value *Waypoint
	isSet bool
}

func (v NullableWaypoint) Get() *Waypoint {
	return v.value
}

func (v *NullableWaypoint) Set(val *Waypoint) {
	v.value = val
	v.isSet = true
}

func (v NullableWaypoint) IsSet() bool {
	return v.isSet
}

func (v *NullableWaypoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWaypoint(val *Waypoint) *NullableWaypoint {
	return &NullableWaypoint{value: val, isSet: true}
}

func (v NullableWaypoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWaypoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


