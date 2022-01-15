package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var kitapKelime = Kelime{
	Madde:       "kitap",
	Birlesikler: "kitap açacağı, kitap cebi, kitap dolabı, kitap düşkünü, kitap ehli, kitabevi, kitap fuarı, kitap kurdu, kitap sarayı, kitapsever, ana kitap, beyaz kitap, ehlikitap, hesap kitap, kara kaplı kitap, yardımcı kitap, yasak kitap, adres kitabı, baş ucu kitabı, boyama kitabı, cep kitabı, el kitabı, okuma kitabı, şiir kitabı",
}

func newSozlukServer(t *testing.T) *httptest.Server {
	t.Helper()
	respWri, err := json.Marshal(kitapKelime)

	assert.NoError(t, err)

	sozlukServer := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				w.Write([]byte(respWri))
			},
		),
	)

	return sozlukServer
}

func TestGetAra(t *testing.T) {
	search := "kitappp"
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
	assert.Equal(t, k.Madde, search)
}
