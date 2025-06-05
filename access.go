package aaronjson


func Marshal(obj interface{}) JsonData {
	return nil
}

func Parse(data []byte) (JsonData, error) {
	return parseJsonByte(data)
}

func ParseString(data string) (JsonData, error) {	
	return Parse([]byte(data))
}

