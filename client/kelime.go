package client

// Kelime is an entry in the sozluk
type Kelime struct {
	MaddeId  string `json:"madde_id"`
	Kac      string `json:"kac"`
	KelimeNo string `json:"kelime_no"`
	Cesit    string `json:"cesit"`
	AnlamGor string `json:"anlam_gor"`

	// Madde represents the entry name, e.g. the word name
	Madde string `json:"madde"`

	CesitSay string `json:"cesit_say"`
	AnlamSay string `json:"anlam_say"`

	// Taki represents the conjucation of the word when referring to a specific object
	Taki string `json:"taki"`

	LisanKodu string `json:"lisan_kodu"`

	// Birlesikler are the "compounds" of the Kelime
	Birlesikler string `json:"birlesikler"`

	// Lisan represents the language of origin for the Kelime
	Lisan string `json:"lisan"`

	// anamlarListe represents the list of meanings a kelime has
	AnlamlarListe []Anlam `json:"anlamlarListe"`
}

// Anlam represents one meaning of a Kelime.
type Anlam struct {
	Anlam string `json:"anlam"`

	// Examples represents the list of example sentences for the Anlam
	Ornekler []Ornekler `json:"orneklerListe"`
}

// Ornekler represent the list of examples for an Anlam
type Ornekler struct {
	// Ornek represents the example sentence for an Anlam
	Ornek string `json:"ornek"`
}
