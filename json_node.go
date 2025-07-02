package aaronjson

import (
	"fmt"
)

type jsonNode struct {
}

func (n *jsonNode) Get(key ...string) (JsonValue, error) {
	return nil, fmt.Errorf("cannot get key %v from %s", key, n.String())
}

func (n *jsonNode) GetMap() (map[string]JsonValue, error) {
	return nil, fmt.Errorf("cannot get map from %s", n.String())
}

func (n *jsonNode) GetSlice() ([]JsonValue, error) {
	return nil, fmt.Errorf("cannot get slice from %s", n.String())
}

func (n *jsonNode) String() string {
	return ""
}

func (n *jsonNode) PrettyString() string {
	return n.String()
}

func (n *jsonNode) AsString() (string, error) {
	return "", fmt.Errorf("cannot convert %s to string", n.String())
}

func (n *jsonNode) AsInt() (int, error) {
	return 0, fmt.Errorf("cannot convert %s to int", n.String())
}

func (n *jsonNode) AsFloat() (float64, error) {
	return 0, fmt.Errorf("cannot convert %s to float", n.String())
}

func (n *jsonNode) AsBool() (bool, error) {
	return false, fmt.Errorf("cannot convert %s to bool", n.String())
}

func (n *jsonNode) AsObject() (*JsonObject, error) {
	return nil, fmt.Errorf("cannot convert %s to object", n.String())
}

func (n *jsonNode) AsArray() (*JsonArray, error) {
	return nil, fmt.Errorf("cannot convert %s to array", n.String())
}

func (n *jsonNode) IsNull() bool {
	return false
}

func (n *jsonNode) IsString() bool {
	return false
}

func (n *jsonNode) IsInt() bool {
	return false
}

func (node *jsonNode) IsFloat() bool {
	return false
}

func (n *jsonNode) IsBool() bool {
	return false
}

func (n *jsonNode) IsObject() bool {
	return false
}

func (n *jsonNode) IsArray() bool {
	return false
}

// Unmarshal will unmarshal the JSON data into the provided interface.
func (n *jsonNode) Unmarshal(v interface{}) error {
	return fmt.Errorf("cannot unmarshal %s into %T", n.String(), v)
}
