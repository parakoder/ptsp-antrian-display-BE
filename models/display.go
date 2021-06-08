package models

type DisplayAntrian struct {
	Loket string `json:"loket"`
	Antrian string `json:"antrian"`
}

type ResponseDisplayAntrian struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data []DisplayAntrian `json:"data"`
}

type ResponseTextBerjalan struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data []string `json:"data"`
}

