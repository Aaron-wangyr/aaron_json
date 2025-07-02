package aaronjson

import (
	"testing"
)

func TestNewJsonArray(t *testing.T) {
	arr := NewJsonArray()
	if arr == nil {
		t.Error("NewJsonArray() returned nil")
	}
	if !arr.IsArray() {
		t.Error("NewJsonArray() should return an array")
	}
	if arr.String() != "[]" {
		t.Errorf("NewJsonArray().String() = %v, want []", arr.String())
	}
}

func TestJsonArrayAppend(t *testing.T) {
	arr := NewJsonArray()
	
	// Test appending valid value
	val, err := arr.Append(NewJsonString("hello"))
	if err != nil {
		t.Errorf("Append() error = %v", err)
	}
	if val == nil {
		t.Error("Append() returned nil value")
	}
	
	// Test appending nil value (should fail)
	_, err = arr.Append(nil)
	if err == nil {
		t.Error("Append() should return error when appending nil value")
	}
	
	// Test the array contains the value
	if arr.String() != `["hello"]` {
		t.Errorf("Array string representation = %v, want [\"hello\"]", arr.String())
	}
}

func TestJsonArrayIndex(t *testing.T) {
	arr := NewJsonArray()
	_, _ = arr.Append(NewJsonString("first"))
	_, _ = arr.Append(NewJsonString("second"))
	_, _ = arr.Append(NewJsonInt(42))
	
	// Test getting valid index
	val, err := arr.Index(0)
	if err != nil {
		t.Errorf("Index() error = %v", err)
	}
	if val.String() != "first" {
		t.Errorf("Index(0) = %v, want first", val.String())
	}
	
	// Test getting index out of bounds
	_, err = arr.Index(10)
	if err == nil {
		t.Error("Index() should return error for out of bounds index")
	}
	
	// Test negative index
	_, err = arr.Index(-1)
	if err == nil {
		t.Error("Index() should return error for negative index")
	}
}

func TestJsonArraySetByIndex(t *testing.T) {
	arr := NewJsonArray()
	_, _ = arr.Append(NewJsonString("first"))
	_, _ = arr.Append(NewJsonString("second"))
	
	// Test setting valid index
	val, err := arr.SetByIndex(0, NewJsonString("updated"))
	if err != nil {
		t.Errorf("SetByIndex() error = %v", err)
	}
	if val.String() != "updated" {
		t.Errorf("SetByIndex() = %v, want updated", val.String())
	}
	
	// Verify the change
	val, _ = arr.Index(0)
	if val.String() != "updated" {
		t.Errorf("Index(0) after SetByIndex = %v, want updated", val.String())
	}
	
	// Test setting index out of bounds
	_, err = arr.SetByIndex(10, NewJsonString("invalid"))
	if err == nil {
		t.Error("SetByIndex() should return error for out of bounds index")
	}
}

func TestJsonArrayRemoveByIndex(t *testing.T) {
	arr := NewJsonArray()
	_, _ = arr.Append(NewJsonString("first"))
	_, _ = arr.Append(NewJsonString("second"))
	_, _ = arr.Append(NewJsonString("third"))
	
	// Test removing valid index
	val, err := arr.RemoveByIndex(1)
	if err != nil {
		t.Errorf("RemoveByIndex() error = %v", err)
	}
	if val.String() != "second" {
		t.Errorf("RemoveByIndex(1) = %v, want second", val.String())
	}
	
	// Verify the array length
	length, _ := arr.Length()
	if length != 2 {
		t.Errorf("Array length after remove = %v, want 2", length)
	}
	
	// Verify remaining elements
	val, _ = arr.Index(1)
	if val.String() != "third" {
		t.Errorf("Index(1) after remove = %v, want third", val.String())
	}
	
	// Test removing index out of bounds
	_, err = arr.RemoveByIndex(10)
	if err == nil {
		t.Error("RemoveByIndex() should return error for out of bounds index")
	}
}

func TestJsonArrayLength(t *testing.T) {
	arr := NewJsonArray()
	
	// Test empty array length
	length, err := arr.Length()
	if err != nil {
		t.Errorf("Length() error = %v", err)
	}
	if length != 0 {
		t.Errorf("Length() = %v, want 0", length)
	}
	
	// Test after adding elements
	_, _ = arr.Append(NewJsonString("item1"))
	_, _ = arr.Append(NewJsonString("item2"))
	_, _ = arr.Append(NewJsonString("item3"))
	
	length, err = arr.Length()
	if err != nil {
		t.Errorf("Length() error = %v", err)
	}
	if length != 3 {
		t.Errorf("Length() = %v, want 3", length)
	}
}

func TestJsonArrayGetSlice(t *testing.T) {
	arr := NewJsonArray()
	_, _ = arr.Append(NewJsonString("item1"))
	_, _ = arr.Append(NewJsonString("item2"))
	
	// Test getting slice
	slice, err := arr.GetSlice()
	if err != nil {
		t.Errorf("GetSlice() error = %v", err)
	}
	if len(slice) != 2 {
		t.Errorf("GetSlice() length = %v, want 2", len(slice))
	}
	if slice[0].String() != "item1" {
		t.Errorf("GetSlice()[0] = %v, want item1", slice[0].String())
	}
}

func TestJsonArrayUnmarshal(t *testing.T) {
	arr := NewJsonArray()
	_, _ = arr.Append(NewJsonString("hello"))
	_, _ = arr.Append(NewJsonString("world"))
	_, _ = arr.Append(NewJsonInt(42))
	
	// Test unmarshaling to slice
	var slice []interface{}
	err := arr.Unmarshal(&slice)
	if err != nil {
		t.Errorf("Unmarshal() error = %v", err)
	}
	if len(slice) != 3 {
		t.Errorf("Unmarshaled slice length = %v, want 3", len(slice))
	}
	
	// Test unmarshaling to string slice
	strArr := NewJsonArray()
	_, _ = strArr.Append(NewJsonString("hello"))
	_, _ = strArr.Append(NewJsonString("world"))
	
	var strSlice []string
	err = strArr.Unmarshal(&strSlice)
	if err != nil {
		t.Errorf("Unmarshal() to string slice error = %v", err)
	}
	if len(strSlice) != 2 {
		t.Errorf("String slice length = %v, want 2", len(strSlice))
	}
	if strSlice[0] != "hello" {
		t.Errorf("String slice[0] = %v, want hello", strSlice[0])
	}
	
	// Test unmarshaling to array
	var fixedArray [2]string
	err = strArr.Unmarshal(&fixedArray)
	if err != nil {
		t.Errorf("Unmarshal() to array error = %v", err)
	}
	if fixedArray[0] != "hello" {
		t.Errorf("Array[0] = %v, want hello", fixedArray[0])
	}
	
	// Test unmarshaling to array with insufficient capacity
	var smallArray [1]string
	err = strArr.Unmarshal(&smallArray)
	if err == nil {
		t.Error("Unmarshal() should return error when array capacity is insufficient")
	}
	
	// Test unmarshaling to non-pointer (should fail)
	var badTarget []string
	err = arr.Unmarshal(badTarget)
	if err == nil {
		t.Error("Unmarshal() should return error for non-pointer target")
	}
	
	// Test unmarshaling to nil (should fail)
	err = arr.Unmarshal(nil)
	if err == nil {
		t.Error("Unmarshal() should return error for nil target")
	}
}

func TestJsonArrayString(t *testing.T) {
	arr := NewJsonArray()
	
	// Test empty array
	if arr.String() != "[]" {
		t.Errorf("Empty array String() = %v, want []", arr.String())
	}
	
	// Test array with elements
	_, _ = arr.Append(NewJsonString("hello"))
	_, _ = arr.Append(NewJsonInt(42))
	_, _ = arr.Append(NewJsonBool(true))
	
	expected := "[\"hello\", 42.000000, true]"
	if arr.String() != expected {
		t.Errorf("Array String() = %v, want %v", arr.String(), expected)
	}
}

func TestJsonArrayPrettyString(t *testing.T) {
	arr := NewJsonArray()
	_, _ = arr.Append(NewJsonString("hello"))
	_, _ = arr.Append(NewJsonInt(42))
	
	pretty := arr.PrettyString()
	if pretty == "" {
		t.Error("PrettyString() returned empty string")
	}
	
	// Test empty array pretty string
	emptyArr := NewJsonArray()
	if emptyArr.PrettyString() != "[]" {
		t.Errorf("Empty array PrettyString() = %v, want []", emptyArr.PrettyString())
	}
}

func TestJsonArrayNestedStructures(t *testing.T) {
	// Create nested array structure
	innerArr := NewJsonArray()
	_, _ = innerArr.Append(NewJsonString("nested"))
	_, _ = innerArr.Append(NewJsonInt(123))
	
	outerArr := NewJsonArray()
	_, _ = outerArr.Append(innerArr)
	_, _ = outerArr.Append(NewJsonString("outer"))
	
	// Test the structure
	length, _ := outerArr.Length()
	if length != 2 {
		t.Errorf("Outer array length = %v, want 2", length)
	}
	
	// Get inner array
	inner, err := outerArr.Index(0)
	if err != nil {
		t.Errorf("Getting inner array error = %v", err)
	}
	if !inner.IsArray() {
		t.Error("First element should be an array")
	}
}
