package aaronjson

// Base interface for JSON data types
type JsonValue interface {
	Get(key ...string) (JsonValue, error)
	GetMap() (map[string]JsonValue, error)
	GetSlice() ([]JsonValue, error)

	AsString() (string, error)
	AsInt() (int, error)
	AsFloat() (float64, error)
	AsBool() (bool, error)
	AsObject() (*JsonObject, error)
	AsArray() (*JsonArray, error)

	IsNull() bool
	IsString() bool
	IsFloat() bool
	IsInt() bool
	IsBool() bool
	IsObject() bool
	IsArray() bool

	Unmarshal(v interface{}) error

	String() string
	PrettyString() string
}
