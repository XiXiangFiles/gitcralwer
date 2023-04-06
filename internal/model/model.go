package model

type Commitment struct {
	ShortHash string `json:"short_hash"`
	Message   string `json:"message"`
	Author    string `json:"author"`
}

type ReduceCommitment struct {
	Ticket string `json:"ticket"`
	Author string `json:"author"`
}
