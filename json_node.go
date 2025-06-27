package aaronjson

import (
	"fmt"
)

// Base node structure for JSON data
type jsonNode struct {
	nodeType JsonType
	data     interface{}
}

// JsonData 接口实现 - 基础方法
func (n *jsonNode) String() string {
	switch n.nodeType {
	case TypeString:
		if str, ok := n.data.(string); ok {
			return str
		}
	case TypeNumber:
		if num, ok := n.data.(float64); ok {
			return fmt.Sprintf("%g", num)
		}
	case TypeBool:
		if b, ok := n.data.(bool); ok {
			return fmt.Sprintf("%t", b)
		}
	case TypeNull:
		return "null"
	}
	return ""
}

func (n *jsonNode) Type() JsonType {
	return n.nodeType
}

func (n *jsonNode) Get(key string) (JsonData, error) {
	return nil, fmt.Errorf("cannot use Get on this type, expected TypeObject but got %v", n.nodeType)
}

func (n *jsonNode) Index(i int) (JsonData, error) {
	return nil, fmt.Errorf("cannot use Index on this type, expected TypeArray but got %v", n.nodeType)
}

// 类型转换方法
func (n *jsonNode) AsString() (string, error) {
	return "", fmt.Errorf("cannot convert to string, expected TypeString but got %v", n.nodeType)
}

func (n *jsonNode) AsInt() (int, error) {
	return 0, fmt.Errorf("cannot convert to int, expected TypeNumber but got %v", n.nodeType)
}

func (n *jsonNode) AsFloat() (float64, error) {
	return 0, fmt.Errorf("cannot convert to float64, expected TypeNumber but got %v", n.nodeType)
}

func (n *jsonNode) AsBool() (bool, error) {
	return false, fmt.Errorf("cannot convert to bool, expected TypeBool but got %v", n.nodeType)
}

func (n *jsonNode) AsObject() (JsonObject, error) {
	return nil, fmt.Errorf("cannot convert to object, expected TypeObject but got %v", n.nodeType)
}

func (n *jsonNode) AsArray() (JsonArray, error) {
	return nil, fmt.Errorf("cannot convert to array, expected TypeArray but got %v", n.nodeType)
}

func (n *jsonNode) IsNull() bool {
	return n.nodeType == TypeNull
}

func (n *jsonNode) IsString() bool {
	return n.nodeType == TypeString
}

func (n *jsonNode) IsNumber() bool {
	return n.nodeType == TypeNumber
}

func (n *jsonNode) IsBool() bool {
	return n.nodeType == TypeBool
}

func (n *jsonNode) IsObject() bool {
	return n.nodeType == TypeObject
}

func (n *jsonNode) IsArray() bool {
	return n.nodeType == TypeArray
}

func (n *jsonNode) Length() (int, error) {
	return 0, fmt.Errorf("cannot get length, expected TypeArray or TypeObject but got %v", n.nodeType)
}

func (n *jsonNode) Keys() ([]string, error) {
	return nil, fmt.Errorf("cannot get keys, expected TypeObject but got %v", n.nodeType)
}

func (n *jsonNode) Set(key string, value JsonData) (JsonData, error) {
	return nil, fmt.Errorf("cannot set key on this type, expected TypeObject but got %v", n.nodeType)
}

func (n *jsonNode) SetByIndex(index int, value JsonData) (JsonData, error) {
	return nil, fmt.Errorf("cannot set index on this type, expected TypeArray but got %v", n.nodeType)
}

func (n *jsonNode) Append(value JsonData) (JsonData, error) {
	return nil, fmt.Errorf("cannot append to this type, expected TypeArray but got %v", n.nodeType)
}

func (n *jsonNode) Remove(key string) (JsonData, error) {
	return nil, fmt.Errorf("cannot remove key from this type, expected TypeObject but got %v", n.nodeType)
}

func (n *jsonNode) RemoveByIndex(index int) (JsonData, error) {
	return nil, fmt.Errorf("cannot remove index from this type, expected TypeArray but got %v", n.nodeType)
}

// Unmarshal will unmarshal the JSON data into the provided interface.
func (n *jsonNode) Unmarshal(v interface{}) error {
	return fmt.Errorf("cannot unmarshal into interface, expected TypeObject or TypeArray but got %v", n.nodeType)
}

// UnmarshalTo 是 Unmarshal 的别名，提供更好的语义
func (n *jsonNode) UnmarshalTo(v interface{}) error {
	return n.Unmarshal(v)
}
