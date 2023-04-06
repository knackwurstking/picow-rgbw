package v1

import "github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/pico"

type RequestPostDevice struct {
	Addr string       `json:"addr"`
	RGBW [4]pico.Duty `json:"rgbw"`
}
