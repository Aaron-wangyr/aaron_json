package aaronjson

// Base interface for all JSON values
type JsonValue interface {
	String() string
	Type() JsonType
}

// Base interface for JSON data types
type JsonData interface {
	JsonValue

	Get(key string) (JsonData,error)
	Index(i int) (JsonData, error)

	AsString() (string, error)
	AsInt() (int, error)
	AsFloat() (float64, error)
	AsBool() (bool, error)
	AsObject() (JsonObject, error)
	AsArray() (JsonArray, error)

	IsNull() bool
	IsString() bool
	IsNumber() bool
	IsBool() bool
	IsObject() bool
	IsArray() bool

	// Collection methods
	Length() (int, error)
	Keys() ([]string, error)

	// Modification methods
	Set(key string, value JsonData) (JsonData, error)
	SetByIndex(index int, value JsonData) (JsonData, error)
	Append(value JsonData) (JsonData, error)
	Remove(key string) (JsonData, error)
	RemoveByIndex(index int) (JsonData, error)
	
	// Unmarshal methods
	Unmarshal(v interface{}) error
	// UnmarshalTo unmarshals the JSON data into the provided interface.
	UnmarshalTo(v interface{}) error

	PrettyString() string
}

// JsonType represents the type of JSON data
type JsonType int

const (
	TypeNull = iota
	TypeBool
	TypeNumber
	TypeString
	TypeArray
	TypeObject
)