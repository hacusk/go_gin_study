package model

// Subway return Subway Operation Info Result
type Subway struct {
	Result Result `json:"result"`
}

// Result return Subway Operation Info Records
type Result struct {
	Records []Records `json:"records"`
}

// Records return Subway Operation Info
type Records struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Jp     string `json:"ja-jp"`
	En     string `json:"en"`
	Note   string `json:"note"`
	Time   string `json:"time"`
	Date   string `json:"date"`
}
