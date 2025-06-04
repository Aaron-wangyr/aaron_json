package aaronjson


type JsonValue struct {
}

func (v *JsonValue) parse(data []byte)  error {
	// This method is a placeholder and should be implemented in derived types.
	// It should parse the JSON data and return the appropriate JsonObject.
	return nil
}

func (v *JsonValue) UnmarshalJson(obj interface{}) error {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) Get(key string) (JsonObject, error) {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) GetString(key string) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) Set(key string, value JsonObject) error {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) Remove(key string) error {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) Keys() ([]string, error) {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) Values() ([]JsonObject, error) {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) String() (string, error) {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) PrettyString() (string, error) {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) GetMap() (map[string]JsonObject, error) {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) GetArray() ([]JsonObject, error) {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) Contains(key string) bool {
	panic("not implemented") // TODO: Implement
}


type JsonMap struct {
	JsonValue
	data map[string]JsonObject

	nodeId int
}

func (m *JsonMap) parse(data []byte)  error {
	
	return nil
}

type JsonArray struct {
	JsonValue
	data []JsonObject
}

func (m *JsonArray) parse(data []byte)  error {
	
	return nil
}

type JsonString struct {
	JsonValue
	data string
}

func (s *JsonString) String() (string, error) {
	if s == nil {
		return "", nil
	}
	return s.data, nil
}

type JsonNumber struct {
	JsonValue
	data float64
}

type JSONBool struct {
	JsonValue
	data bool
}

var JsonNull = &JsonValue{}
var JSONTrue = &JsonValue{}
var JSONFalse = &JsonValue{}

