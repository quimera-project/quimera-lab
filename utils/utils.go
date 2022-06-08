package utils

type Unit struct {
	Text  string `json:"text"`
	Level string `json:"level"`
}

type Table struct {
	Header []string `json:"header"`
	Body   []any    `json:"body"`
}
