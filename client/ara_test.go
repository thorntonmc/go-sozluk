package client

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const tKitapString = "kitap"
const tMassaString = "massa"
const kitapResponse = `[{"madde_id":"27874","kac":"0","kelime_no":"30668","cesit":"0","anlam_gor":"0","on_taki":null,"madde":"kitap","cesit_say":"6","anlam_say":"3","taki":"bı","cogul_mu":"0","ozel_mi":"0","lisan_kodu":"11","lisan":"Arapça kitāb","telaffuz":null,"birlesikler":"kitap açacağı, kitap cebi, kitap dolabı, kitap düşkünü, kitap ehli, kitabevi, kitap fuarı, kitap kurdu, kitap sarayı, kitapsever, ana kitap, beyaz kitap, ehlikitap, hesap kitap, kara kaplı kitap, yardımcı kitap, yasak kitap, adres kitabı, baş ucu kitabı, boyama kitabı, cep kitabı, el kitabı, okuma kitabı, şiir kitabı","font":null,"madde_duz":"kitap","gosterim_tarihi":null,"anlamlarListe":[{"anlam_id":"52009","madde_id":"27874","anlam_sira":"1","fiil":"0","tipkes":"0","anlam":"Ciltli veya ciltsiz olarak bir araya getirilmiş, basılı veya yazılı kâğıt yaprakların bütünü","gos":"0","orneklerListe":[{"ornek_id":"9889","anlam_id":"52009","ornek_sira":"1","ornek":"Ama ben, bir kitap üzerine bir fikir edinmek istedim mi o kitabı kendim okurum.","kac":"1","yazar_id":"42","yazar":[{"yazar_id":"42","tam_adi":"Nurullah Ataç","kisa_adi":"N. Ataç","ekno":"173"}]}],"ozelliklerListe":[{"ozellik_id":"19","tur":"3","tam_adi":"isim","kisa_adi":"a.","ekno":"30"}]},{"anlam_id":"52010","madde_id":"27874","anlam_sira":"2","fiil":"0","tipkes":"0","anlam":"Herhangi bir konuda yazılmış eser","gos":"0","orneklerListe":[{"ornek_id":"9890","anlam_id":"52010","ornek_sira":"1","ornek":"Acaba bir edebiyat kitabında hazır bir tarif bulamaz mıyız?","kac":"1","yazar_id":"4","yazar":[{"yazar_id":"4","tam_adi":"Falih Rıfkı Atay","kisa_adi":"F. R. Atay","ekno":"129"}]}]},{"anlam_id":"52011","madde_id":"27874","anlam_sira":"3","fiil":"0","tipkes":"0","anlam":"Kutsal kitap","gos":"0"}],"atasozu":[{"madde_id":"27877","madde":"kitaba (veya kitabına) uydurmak","on_taki":null},{"madde_id":"27876","madde":"kitaba el basmak","on_taki":null},{"madde_id":"27878","madde":"kitabı kapamak","on_taki":null},{"madde_id":"27879","madde":"kitabında yer almamak","on_taki":null},{"madde_id":"27875","madde":"kitap (veya kitaplar) devirmek (veya devretmek)","on_taki":null},{"madde_id":"27880","madde":"kitapta yeri olmak","on_taki":null}]}]`
const massaResponse = `{"error":"Sonuç bulunamadı"}`

var kitapKelime = Kelime{
	Madde:       "kitap",
	Birlesikler: "kitap açacağı, kitap cebi, kitap dolabı, kitap düşkünü, kitap ehli, kitabevi, kitap fuarı, kitap kurdu, kitap sarayı, kitapsever, ana kitap, beyaz kitap, ehlikitap, hesap kitap, kara kaplı kitap, yardımcı kitap, yasak kitap, adres kitabı, baş ucu kitabı, boyama kitabı, cep kitabı, el kitabı, okuma kitabı, şiir kitabı",
}

func newSozlukServer(t *testing.T) *httptest.Server {
	t.Helper()

	sozlukServer := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				values := r.URL.Query()
				ara := values.Get("ara")
				assert.NotEmpty(t, ara)

				if ara == tKitapString {
					w.WriteHeader(200)
					w.Write([]byte(kitapResponse))
				}
				if ara == tMassaString {
					w.WriteHeader(200)
					w.Write([]byte(massaResponse))
				}
			},
		),
	)

	return sozlukServer
}

func TestGetAra(t *testing.T) {
	search := []struct {
		str    string
		expect string
		err    bool
	}{
		{"kitap", "kitap", false},
		{"massa", "", true},
	}

	s := newSozlukServer(t)
	defer s.Close()

	c := NewClient(OptionDebug(true), OptionEndpoint(s.URL))

	for _, i := range search {
		k, e := c.Ara(i.str)
		if e != nil {
			log.Println(e)
		}
		assert.Equal(t, i.err, e != nil, "expected an error to be returned for", i.str)
		if i.expect == "" {
			assert.Empty(t, k, "expected empty array for", i.str)
		} else {
			assert.NotEmpty(t, k)
			assert.NotEmpty(t, k[0])
			assert.NotEmpty(t, k[0].Madde, "returned word is null")
			assert.Equal(t, i.expect, k[0].Madde, "did not receive expected word")
		}
	}
}
