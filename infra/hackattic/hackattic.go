package hackattic

import (
	"github.com/go-resty/resty/v2"
)

var (
	hackattic *resty.Client
)

func InitalizeHackattic() {
	baseUrl := "https://hackattic.com/challenges"

	hackattic = resty.
		New().
		SetBaseURL(baseUrl).
		SetHeader("Content-Type", "application/json")
}
