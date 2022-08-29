package jottingjwts

type jwtSecretDto struct {
	Secret string `json:"jwt_secret"`
}

type AppUrlDto struct {
	AppUrl string `json:"app_url"`
}

type FinalSolution struct {
	Solution string `json:"solution"`
}
