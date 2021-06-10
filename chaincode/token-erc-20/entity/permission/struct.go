package permission

type ApproveInput struct {
	Spender string `json:"spender"`
	Value   int    `json:"value"`
}

type AllowanceInput struct {
	Owner   string `json:"owner"`
	Spender string `json:"spender"`
}
