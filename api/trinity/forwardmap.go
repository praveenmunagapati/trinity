package trinity

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/clownpriest/trinity/core/token"
)

var (

	// ErrForwardMapKeyDoesntExist means that a key did not resolve to an FMValue
	ErrForwardMapKeyDoesntExist = errors.New("forward map key doesn't exist")
)

/*
NewForwardMap returns an empty ForwardMap
*/
func NewForwardMap() ForwardMap {
	return ForwardMap{TheMap: make(map[string]*FMValue)}
}

/*
Add appends a location int32 to the FMValue's Locations
list
*/
func (m *FMValue) Add(location uint32) {
	m.Locations = append(m.Locations, location)
}

/*
Get returns the FMValue associated to a kew in the forward map
*/
func (m *ForwardMap) Get(key string) (*FMValue, error) {
	result, exists := m.GetTheMap()[key]
	if !exists {
		return &FMValue{}, ErrForwardMapKeyDoesntExist
	}
	return result, nil
}

/*
Put adds a token/location pair to the map
*/
func (m *ForwardMap) Put(token string, location uint32) {
	theMap := m.GetTheMap()
	value, exists := theMap[token]
	if !exists {
		theMap[token] = &FMValue{Locations: []uint32{location}}
		return
	}
	value.Add(location)
}

/*
Delete removes an entry from the map
*/
func (m *ForwardMap) Delete(token string) {
	delete(m.GetTheMap(), token)
}

/*
AddList dumps a list of tokens into the Map
*/
func (m *ForwardMap) AddList(list token.List) {
	theMap := m.GetTheMap()
	for _, token := range list.Tokens {
		value, exists := theMap[string(token.Value)]
		if !exists {
			theMap[string(token.Value)] = &FMValue{Locations: []uint32{token.Location}}
			continue
		}
		value.Add(token.Location)
	}
	m.DeleteArticles()
}

/*
DeleteArticles removes tokens which are English articles
*/
func (m *ForwardMap) DeleteArticles() {
	m.Delete("A")
	m.Delete("AN")
	m.Delete("THE")
}

/*
SaveJSON serializes a ForwardMap to a JSON file
*/
func (m *ForwardMap) SaveJSON(path string) error {
	file, err := os.Create(path)
	if err == nil {
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		encoder.Encode(m.GetTheMap())
	}
	file.Close()
	return err
}

func (m *ForwardMap) SaveProtoToFIle(path string) error {
	file, _ := os.Create(path)
	data, _ := m.Marshal()
	file.Write(data)
	return nil
}
