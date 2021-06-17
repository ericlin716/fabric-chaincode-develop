package transaction

type TransferInput struct {
	Recipient string `json:"recipient"`
	Amount    int    `json:"amount"`
}

type TransferFromInput struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value int    `json:"value"`
}
