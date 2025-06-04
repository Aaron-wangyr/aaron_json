package aaronjson

type JsonObject interface{
	parse(data []byte) (error)
	UnmarshalJson(obj interface{}) error
	Get(key string) (JsonObject, error)
	GetString(key string) (string, error)
	Set(key string, value JsonObject) error
	Remove(key string) error
	Keys() ([]string, error)
	Values() ([]JsonObject, error)
	String() (string, error)
	PrettyString() (string, error)
	GetMap() (map[string]JsonObject, error)
	GetArray() ([]JsonObject, error)
	Contains(key string) bool
}

