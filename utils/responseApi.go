package utils

type ResponsJson struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type ResponsJsonString struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type ResponsJsonArray struct {
	Error   bool     `json:"error"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

type ResponsJsonStruct struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
