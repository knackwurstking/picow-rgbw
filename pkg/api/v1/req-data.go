package v1

type ReqPutDevice struct {
	Addr string `json:"addr"`
	RGBW [4]int `json:"rgbw"`
}
