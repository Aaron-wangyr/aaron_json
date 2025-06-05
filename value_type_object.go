package aaronjson


type JsonObject struct {
	JsonValue
	sortedKeys []string
	data map[string]JsonData

	nodeId int
}

func (m *JsonObject) parse(data []byte,pos int) (JsonData, int, error) {
	return nil, pos, nil
}
