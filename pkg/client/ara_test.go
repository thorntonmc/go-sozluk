package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const kitapResponse = `[{"madde_id":"27874","kac":"0","kelime_no":"30668","cesit":"0","anlam_gor":"0","on_taki":null,"madde":"kitap","cesit_say":"6","anlam_say":"3","taki":"bı","cogul_mu":"0","ozel_mi":"0","lisan_kodu":"11","lisan":"Arapça kitāb","telaffuz":null,"birlesikler":"kitap açacağı, kitap cebi, kitap dolabı, kitap düşkünü, kitap ehli, kitabevi, kitap fuarı, kitap kurdu, kitap sarayı, kitapsever, ana kitap, beyaz kitap, ehlikitap, hesap kitap, kara kaplı kitap, yardımcı kitap, yasak kitap, adres kitabı, baş ucu kitabı, boyama kitabı, cep kitabı, el kitabı, okuma kitabı, şiir kitabı","font":null,"madde_duz":"kitap","gosterim_tarihi":null,"anlamlarListe":[{"anlam_id":"52009","madde_id":"27874","anlam_sira":"1","fiil":"0","tipkes":"0","anlam":"Ciltli veya ciltsiz olarak bir araya getirilmiş, basılı veya yazılı kâğıt yaprakların bütünü","gos":"0","orneklerListe":[{"ornek_id":"9889","anlam_id":"52009","ornek_sira":"1","ornek":"Ama ben, bir kitap üzerine bir fikir edinmek istedim mi o kitabı kendim okurum.","kac":"1","yazar_id":"42","yazar":[{"yazar_id":"42","tam_adi":"Nurullah Ataç","kisa_adi":"N. Ataç","ekno":"173"}]}],"ozelliklerListe":[{"ozellik_id":"19","tur":"3","tam_adi":"isim","kisa_adi":"a.","ekno":"30"}]},{"anlam_id":"52010","madde_id":"27874","anlam_sira":"2","fiil":"0","tipkes":"0","anlam":"Herhangi bir konuda yazılmış eser","gos":"0","orneklerListe":[{"ornek_id":"9890","anlam_id":"52010","ornek_sira":"1","ornek":"Acaba bir edebiyat kitabında hazır bir tarif bulamaz mıyız?","kac":"1","yazar_id":"4","yazar":[{"yazar_id":"4","tam_adi":"Falih Rıfkı Atay","kisa_adi":"F. R. Atay","ekno":"129"}]}]},{"anlam_id":"52011","madde_id":"27874","anlam_sira":"3","fiil":"0","tipkes":"0","anlam":"Kutsal kitap","gos":"0"}],"atasozu":[{"madde_id":"27877","madde":"kitaba (veya kitabına) uydurmak","on_taki":null},{"madde_id":"27876","madde":"kitaba el basmak","on_taki":null},{"madde_id":"27878","madde":"kitabı kapamak","on_taki":null},{"madde_id":"27879","madde":"kitabında yer almamak","on_taki":null},{"madde_id":"27875","madde":"kitap (veya kitaplar) devirmek (veya devretmek)","on_taki":null},{"madde_id":"27880","madde":"kitapta yeri olmak","on_taki":null}]}]`

var kitapKelime = Kelime{
	Madde:       "kitap",
	Birlesikler: "kitap açacağı, kitap cebi, kitap dolabı, kitap düşkünü, kitap ehli, kitabevi, kitap fuarı, kitap kurdu, kitap sarayı, kitapsever, ana kitap, beyaz kitap, ehlikitap, hesap kitap, kara kaplı kitap, yardımcı kitap, yasak kitap, adres kitabı, baş ucu kitabı, boyama kitabı, cep kitabı, el kitabı, okuma kitabı, şiir kitabı",
}

func newSozlukServer(t *testing.T) *httptest.Server {
	t.Helper()
	sozlukServer := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				w.Write([]byte(kitapResponse))
			},
		),
	)

	return sozlukServer
}

func TestGetAra(t *testing.T) {
	search := "kitap"
	s := newSozlukServer(t)
	defer s.Close()

	c := &Client{
		&http.Client{},
		logrus.Logger{},
		false,
		s.URL,
	}

	k, e := c.Ara(search)

	assert.NoError(t, e)
	assert.Equal(t, k[0].Madde, search)
}
