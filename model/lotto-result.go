package model

type LottoResult struct {
	BnusNo         int    `json:"bnusNo"`
	FirstAccumamnt int    `json:"firstAccumamnt"`
	FirstWinamnt   int    `json:"firstWinamnt"`
	ReturnValue    string `json:"returnValue"`
	TotSellamnt    int64  `json:"totSellamnt"`
	DrwtNo1        int    `json:"drwtNo1"`
	DrwtNo2        int    `json:"drwtNo2"`
	DrwtNo3        int    `json:"drwtNo3"`
	DrwtNo4        int    `json:"drwtNo4"`
	DrwtNo5        int    `json:"drwtNo5"`
	DrwtNo6        int    `json:"drwtNo6"`
	DrwNoDate      string `json:"drwNoDate"`
	DrwNo          int    `json:"drwNo"`
	FirstPrzwnerCo int    `json:"firstPrzwnerCo"`
}