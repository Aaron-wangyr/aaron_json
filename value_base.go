package aaronjson

type JsonValue struct {
}

func (v *JsonValue) parse(data []byte,pos int) (JsonData,int,error){
	// This method is a placeholder and should be implemented in derived types.
	// It should parse the JSON data and return the appropriate JsonObject.
	return nil, pos, nil
}

func (v *JsonValue) UnmarshalJson(obj interface{}) error {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) Get(key string) (JsonData, error) {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) GetString(key string) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) Set(key string, value JsonData) error {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) Remove(key string) error {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) Keys() ([]string, error) {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) Values() ([]JsonData, error) {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) String() (string, error) {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) PrettyString() (string, error) {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) GetMap() (map[string]JsonData, error) {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) GetArray() ([]JsonData, error) {
	panic("not implemented") // TODO: Implement
}

func (v *JsonValue) Contains(key string) bool {
	panic("not implemented") // TODO: Implement
}


