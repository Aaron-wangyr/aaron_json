package aaronjson

import (
	"testing"
)

func TestErrorConstants(t *testing.T) {
	// Test that all error constants are defined and not nil
	if ErrEmptyData == nil {
		t.Error("ErrEmptyData should not be nil")
	}
	if ErrInvalidJsonFormat == nil {
		t.Error("ErrInvalidJsonFormat should not be nil")
	}
	if ErrIncorrectOperation == nil {
		t.Error("ErrIncorrectOperation should not be nil")
	}
	if ErrIncorrectOperationWithMsg == nil {
		t.Error("ErrIncorrectOperationWithMsg should not be nil")
	}
	if ErrKeyNotFound == nil {
		t.Error("ErrKeyNotFound should not be nil")
	}
	if ErrIndexOutOfBounds == nil {
		t.Error("ErrIndexOutOfBounds should not be nil")
	}
	if ErrNilValueAppend == nil {
		t.Error("ErrNilValueAppend should not be nil")
	}
	if ErrNilValueRemove == nil {
		t.Error("ErrNilValueRemove should not be nil")
	}
	if ErrUnmarshalNilInterface == nil {
		t.Error("ErrUnmarshalNilInterface should not be nil")
	}
	if ErrUnmarshalTargetNotPointer == nil {
		t.Error("ErrUnmarshalTargetNotPointer should not be nil")
	}
	if ErrUnmarshalTargetNotSettable == nil {
		t.Error("ErrUnmarshalTargetNotSettable should not be nil")
	}
	if ErrUnmarshalTargetTypeMismatch == nil {
		t.Error("ErrUnmarshalTargetTypeMismatch should not be nil")
	}
}

func TestErrorMessages(t *testing.T) {
	// Test that error messages are meaningful
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "ErrEmptyData",
			err:  ErrEmptyData,
			want: "empty data",
		},
		{
			name: "ErrInvalidJsonFormat",
			err:  ErrInvalidJsonFormat,
			want: "invalid JSON format",
		},
		{
			name: "ErrIncorrectOperation",
			err:  ErrIncorrectOperation,
			want: "incorrect operation",
		},
		{
			name: "ErrKeyNotFound",
			err:  ErrKeyNotFound,
			want: "key not found",
		},
		{
			name: "ErrIndexOutOfBounds",
			err:  ErrIndexOutOfBounds,
			want: "index out of bounds",
		},
		{
			name: "ErrNilValueAppend",
			err:  ErrNilValueAppend,
			want: "cannot append nil value to array",
		},
		{
			name: "ErrNilValueRemove",
			err:  ErrNilValueRemove,
			want: "cannot remove nil value from array",
		},
		{
			name: "ErrUnmarshalNilInterface",
			err:  ErrUnmarshalNilInterface,
			want: "cannot unmarshal into nil interface",
		},
		{
			name: "ErrUnmarshalTargetNotPointer",
			err:  ErrUnmarshalTargetNotPointer,
			want: "unmarshal target must be a pointer",
		},
		{
			name: "ErrUnmarshalTargetNotSettable",
			err:  ErrUnmarshalTargetNotSettable,
			want: "unmarshal target cannot be set",
		},
		{
			name: "ErrUnmarshalTargetTypeMismatch",
			err:  ErrUnmarshalTargetTypeMismatch,
			want: "unmarshal target type mismatch",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.err.Error()
			if got != tt.want {
				t.Errorf("Error message = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorUsage(t *testing.T) {
	// Test that errors are actually used in the codebase
	
	// Test ErrIndexOutOfBounds usage in JsonArray
	arr := NewJsonArray()
	_, err := arr.Index(0)
	if err != ErrIndexOutOfBounds {
		t.Errorf("Expected ErrIndexOutOfBounds, got %v", err)
	}
	
	_, err = arr.SetByIndex(0, NewJsonString("test"))
	if err != ErrIndexOutOfBounds {
		t.Errorf("Expected ErrIndexOutOfBounds, got %v", err)
	}
	
	_, err = arr.RemoveByIndex(0)
	if err != ErrIndexOutOfBounds {
		t.Errorf("Expected ErrIndexOutOfBounds, got %v", err)
	}
	
	// Test ErrNilValueAppend usage in JsonArray
	_, err = arr.Append(nil)
	if err != ErrNilValueAppend {
		t.Errorf("Expected ErrNilValueAppend, got %v", err)
	}
	
	// Test unmarshal errors
	str := NewJsonString("test")
	
	// Test ErrUnmarshalNilInterface
	err = str.Unmarshal(nil)
	if err == nil {
		t.Error("Expected error when unmarshaling to nil")
	}
	
	// Test ErrUnmarshalTargetNotPointer
	var target string
	err = str.Unmarshal(target)
	if err == nil {
		t.Error("Expected error when unmarshaling to non-pointer")
	}
}
