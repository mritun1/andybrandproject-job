package models

// -------------------------------
// FOR OUTPUT OF JSON DATA
// -------------------------------
type OutPut struct {
	Code  int   `json:"code"`
	Count int64 `json:"count"`
	Data  any   `json:"data"`
}

// -------------------------------
// STANDARD FORMAT FOR OUTPUT JSON
// -------------------------------
type StandardOutPut struct {
	Code  int   `json:"code"`
	Count int64 `json:"count"`
}

// -------------------------------
// JSON OUTPUT FOR LOGGING
// -------------------------------
type LogOutPut struct {
	Code  int    `json:"code"`
	Token string `json:"token"`
	Data  any    `json:"data"`
}
