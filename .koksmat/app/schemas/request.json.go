package schemas

type Request struct {
	Accesstoken string `json:"accesstoken"`
	Request     struct {
		Whatever string `json:"whatever"`
	} `json:"request"`
}
