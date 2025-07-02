package aaronjson

import (
	"testing"
)

func TestJsonNodeInterface(t *testing.T) {
	node := &jsonNode{}
	
	// Test that jsonNode implements JsonValue interface
	var _ JsonValue = node
	
	// Test default implementations return errors
	_, err := node.Get("key")
	if err == nil {
		t.Error("jsonNode.Get() should return error")
	}
	
	_, err = node.GetMap()
	if err == nil {
		t.Error("jsonNode.GetMap() should return error")
	}
	
	_, err = node.GetSlice()
	if err == nil {
		t.Error("jsonNode.GetSlice() should return error")
	}
	
	_, err = node.AsString()
	if err == nil {
		t.Error("jsonNode.AsString() should return error")
	}
	
	_, err = node.AsInt()
	if err == nil {
		t.Error("jsonNode.AsInt() should return error")
	}
	
	_, err = node.AsFloat()
	if err == nil {
		t.Error("jsonNode.AsFloat() should return error")
	}
	
	_, err = node.AsBool()
	if err == nil {
		t.Error("jsonNode.AsBool() should return error")
	}
	
	_, err = node.AsObject()
	if err == nil {
		t.Error("jsonNode.AsObject() should return error")
	}
	
	_, err = node.AsArray()
	if err == nil {
		t.Error("jsonNode.AsArray() should return error")
	}
	
	err = node.Unmarshal(&struct{}{})
	if err == nil {
		t.Error("jsonNode.Unmarshal() should return error")
	}
}

func TestJsonNodeTypeChecks(t *testing.T) {
	node := &jsonNode{}
	
	// All type checks should return false for base jsonNode
	if node.IsNull() {
		t.Error("jsonNode.IsNull() should return false")
	}
	if node.IsString() {
		t.Error("jsonNode.IsString() should return false")
	}
	if node.IsInt() {
		t.Error("jsonNode.IsInt() should return false")
	}
	if node.IsFloat() {
		t.Error("jsonNode.IsFloat() should return false")
	}
	if node.IsBool() {
		t.Error("jsonNode.IsBool() should return false")
	}
	if node.IsObject() {
		t.Error("jsonNode.IsObject() should return false")
	}
	if node.IsArray() {
		t.Error("jsonNode.IsArray() should return false")
	}
}

func TestJsonNodeString(t *testing.T) {
	node := &jsonNode{}
	
	// Test String() method
	str := node.String()
	if str != "" {
		t.Errorf("jsonNode.String() = %v, want empty string", str)
	}
	
	// Test PrettyString() method
	prettyStr := node.PrettyString()
	if prettyStr != "" {
		t.Errorf("jsonNode.PrettyString() = %v, want empty string", prettyStr)
	}
}

func TestJsonNodeInheritance(t *testing.T) {
	// Test that concrete types properly inherit from jsonNode
	
	// JsonString should override some methods
	str := NewJsonString("test")
	if !str.IsString() {
		t.Error("JsonString should override IsString() to return true")
	}
	if str.IsInt() {
		t.Error("JsonString should inherit IsInt() returning false")
	}
	
	// JsonInt should override some methods
	num := NewJsonInt(42)
	if !num.IsInt() {
		t.Error("JsonInt should override IsInt() to return true")
	}
	if num.IsString() {
		t.Error("JsonInt should inherit IsString() returning false")
	}
	
	// JsonBool should override some methods
	boolean := NewJsonBool(true)
	if !boolean.IsBool() {
		t.Error("JsonBool should override IsBool() to return true")
	}
	if boolean.IsString() {
		t.Error("JsonBool should inherit IsString() returning false")
	}
	
	// JsonNull should override some methods
	null := NewJsonNull()
	if !null.IsNull() {
		t.Error("JsonNull should override IsNull() to return true")
	}
	if null.IsString() {
		t.Error("JsonNull should inherit IsString() returning false")
	}
	
	// JsonObject should override some methods
	obj := NewJsonObject()
	if !obj.IsObject() {
		t.Error("JsonObject should override IsObject() to return true")
	}
	if obj.IsString() {
		t.Error("JsonObject should inherit IsString() returning false")
	}
	
	// JsonArray should override some methods
	arr := NewJsonArray()
	if !arr.IsArray() {
		t.Error("JsonArray should override IsArray() to return true")
	}
	if arr.IsString() {
		t.Error("JsonArray should inherit IsString() returning false")
	}
}
