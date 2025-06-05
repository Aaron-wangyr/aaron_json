package aaronjson

type JsonBool struct {
	JsonValue
	data bool
}


var JSONTrue = &JsonValue{}
var JSONFalse = &JsonValue{}