package aaronjson

import (
	"testing"
)

func TestEscapeString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple string",
			input: "hello",
			want:  "hello",
		},
		{
			name:  "string with quotes",
			input: `he said "hello"`,
			want:  `he said \"hello\"`,
		},
		{
			name:  "string with backslash",
			input: `path\to\file`,
			want:  `path\\to\\file`,
		},
		{
			name:  "string with newline",
			input: "hello\nworld",
			want:  "hello\\nworld",
		},
		{
			name:  "string with tab",
			input: "hello\tworld",
			want:  "hello\\tworld",
		},
		{
			name:  "string with carriage return",
			input: "hello\rworld",
			want:  "hello\\rworld",
		},
		{
			name:  "string with form feed",
			input: "hello\fworld",
			want:  "hello\\fworld",
		},
		{
			name:  "string with backspace",
			input: "hello\bworld",
			want:  "hello\\bworld",
		},
		{
			name:  "string with control character",
			input: "hello\x01world",
			want:  "hello\\u0001world",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "string with multiple escapes",
			input: "line1\nline2\tindented\"quoted\"",
			want:  "line1\\nline2\\tindented\\\"quoted\\\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := escapeString(tt.input)
			if got != tt.want {
				t.Errorf("escapeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkipWhitespace(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		pos  int
		want int
	}{
		{
			name: "no whitespace",
			data: []byte("hello"),
			pos:  0,
			want: 0,
		},
		{
			name: "leading spaces",
			data: []byte("   hello"),
			pos:  0,
			want: 3,
		},
		{
			name: "leading tabs",
			data: []byte("\t\thello"),
			pos:  0,
			want: 2,
		},
		{
			name: "leading newlines",
			data: []byte("\n\nhello"),
			pos:  0,
			want: 2,
		},
		{
			name: "leading carriage returns",
			data: []byte("\r\rhello"),
			pos:  0,
			want: 2,
		},
		{
			name: "mixed whitespace",
			data: []byte(" \t\n\r hello"),
			pos:  0,
			want: 5,
		},
		{
			name: "whitespace in middle",
			data: []byte("hello   world"),
			pos:  5,
			want: 8,
		},
		{
			name: "no whitespace from middle",
			data: []byte("hello world"),
			pos:  5,
			want: 6, // position 5 is a space, should skip to 6
		},
		{
			name: "position at end",
			data: []byte("hello"),
			pos:  5,
			want: 5,
		},
		{
			name: "position beyond end",
			data: []byte("hello"),
			pos:  10,
			want: 10,
		},
		{
			name: "all whitespace",
			data: []byte("   \t\n\r   "),
			pos:  0,
			want: 9, // length is 9, not 10
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := skipWhitespace(tt.data, tt.pos)
			if got != tt.want {
				t.Errorf("skipWhitespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidNumberTerminator(t *testing.T) {
	tests := []struct {
		name  string
		input byte
		want  bool
	}{
		{
			name:  "comma",
			input: ',',
			want:  true,
		},
		{
			name:  "closing bracket",
			input: ']',
			want:  true,
		},
		{
			name:  "closing brace",
			input: '}',
			want:  true,
		},
		{
			name:  "space",
			input: ' ',
			want:  true,
		},
		{
			name:  "tab",
			input: '\t',
			want:  true,
		},
		{
			name:  "newline",
			input: '\n',
			want:  true,
		},
		{
			name:  "carriage return",
			input: '\r',
			want:  true,
		},
		{
			name:  "letter",
			input: 'a',
			want:  false,
		},
		{
			name:  "digit",
			input: '5',
			want:  false,
		},
		{
			name:  "opening bracket",
			input: '[',
			want:  false,
		},
		{
			name:  "opening brace",
			input: '{',
			want:  false,
		},
		{
			name:  "quote",
			input: '"',
			want:  false,
		},
		{
			name:  "colon",
			input: ':',
			want:  false,
		},
		{
			name:  "semicolon",
			input: ';',
			want:  false,
		},
		{
			name:  "period",
			input: '.',
			want:  false,
		},
		{
			name:  "minus",
			input: '-',
			want:  false,
		},
		{
			name:  "plus",
			input: '+',
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidNumberTerminator(tt.input)
			if got != tt.want {
				t.Errorf("isValidNumberTerminator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkipWhitespaceEdgeCases(t *testing.T) {
	// Test empty slice
	data := []byte{}
	pos := skipWhitespace(data, 0)
	if pos != 0 {
		t.Errorf("skipWhitespace on empty slice = %v, want 0", pos)
	}

	// Test starting position beyond data length
	data = []byte("hello")
	pos = skipWhitespace(data, 100)
	if pos != 100 {
		t.Errorf("skipWhitespace with pos beyond length = %v, want 100", pos)
	}

	// Note: We don't test negative position as it's not a valid use case
	// and the function assumes valid input
}

func TestEscapeStringUnicode(t *testing.T) {
	// Test with unicode characters
	input := "Hello 世界"
	got := escapeString(input)
	want := "Hello 世界" // Unicode characters should pass through unchanged
	if got != want {
		t.Errorf("escapeString() with unicode = %v, want %v", got, want)
	}

	// Test with control characters that should be escaped as unicode
	input = "hello\x00\x1fworld"
	got = escapeString(input)
	want = "hello\\u0000\\u001fworld"
	if got != want {
		t.Errorf("escapeString() with control chars = %v, want %v", got, want)
	}
}
