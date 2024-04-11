package schemas

type SignedRequest struct {
	Hash    string `json:"hash"`
	Request struct {
		Whatever string `json:"whatever"`
	} `json:"request"`
	Upn string `json:"upn"`
}
