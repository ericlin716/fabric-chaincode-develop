package event

// event provides an organized struct for emitting events
type Event struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value int    `json:"value"`
}
