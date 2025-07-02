package aaronjson

import "errors"

var (
	ErrEmptyData                 = errors.New("empty data")
	ErrInvalidJsonFormat         = errors.New("invalid JSON format")
	ErrIncorrectOperation        = errors.New("incorrect operation")
	ErrIncorrectOperationWithMsg = errors.New("incorrect operation: %s")
	ErrKeyNotFound               = errors.New("key not found")
	ErrIndexOutOfBounds          = errors.New("index out of bounds")
	ErrNilValueAppend            = errors.New("cannot append nil value to array")
	ErrNilValueRemove            = errors.New("cannot remove nil value from array")

	ErrUnmarshalNilInterface       = errors.New("cannot unmarshal into nil interface")
	ErrUnmarshalTargetNotPointer   = errors.New("unmarshal target must be a pointer")
	ErrUnmarshalTargetNotSettable  = errors.New("unmarshal target cannot be set")
	ErrUnmarshalTargetTypeMismatch = errors.New("unmarshal target type mismatch")
)
