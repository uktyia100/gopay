package douyin

type TextAntiResponse struct {
	ErrId     string         `json:"error_id,omitempty"`
	ErrCode   int            `json:"code,omitempty"`
	ErrMsg    string         `json:"message,omitempty"`
	Exception string         `json:"exception,omitempty"`
	LogId     string         `json:"log_id,omitempty"`
	Data      []TextAntiData `json:"data,omitempty"`
}

type TextAntiData struct {
	Msg      string        `json:"msg,omitempty"`
	Code     int           `json:"code,omitempty"`
	TaskId   string        `json:"task_id,omitempty"`
	Predicts []TextPredict `json:"predicts,omitempty"`
}

type TextPredict struct {
	Prob      int         `json:"prob,omitempty"`
	Hit       bool        `json:"hit,omitempty"`
	Target    interface{} `json:"target,omitempty"`
	ModelName string      `json:"model_name,omitempty"`
}

type CensorImgResponse struct {
	ErrCode   int       `json:"error,omitempty"`
	ErrMsg    string    `json:"message,omitempty"`
	Predicts  []Predict `json:"predicts,omitempty"`
	ModelName string    `json:"model_name,omitempty"`
	Hit       bool      `json:"hit,omitempty"`
}

type Predict struct {
	ModelName string `json:"model_name,omitempty"`
	Hit       bool   `json:"hit,omitempty"`
}
