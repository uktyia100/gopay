package douyin

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
