package aaronjson

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