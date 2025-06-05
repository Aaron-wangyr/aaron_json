package aaronjson

type JsonData interface {
	parse(data []byte,pos int) (JsonData,int, error)
	UnmarshalJson(obj interface{}) error
	Get(key string) (JsonData, error)
	GetString(key string) (string, error)
	Set(key string, value JsonData) error
	Remove(key string) error
	Keys() ([]string, error)
	Values() ([]JsonData, error)
	String() (string, error)
	PrettyString() (string, error)
	GetMap() (map[string]JsonData, error)
	GetArray() ([]JsonData, error)
	Contains(key string) bool
}
