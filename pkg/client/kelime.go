package client

// Kelime is an entry in the sozluk
type Kelime struct {
	Madde       string `json:"madde"`
	Birlesikler string `json:"birlesikler"`
}
