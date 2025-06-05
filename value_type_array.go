package aaronjson

type JsonArray struct {
	JsonValue
	data []JsonData
}

func (m *JsonArray) parse(data []byte, pos int) (JsonData, int, error) {

	return m, pos, nil
}
