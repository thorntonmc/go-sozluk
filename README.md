[![test](https://github.com/thorntonmc/go-sozluk/actions/workflows/test.yml/badge.svg)](https://github.com/thorntonmc/go-sozluk/actions/workflows/test.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/thorntonmc/go-sozluk)](https://goreportcard.com/report/github.com/thorntonmc/go-sozluk) [![Go Reference](https://pkg.go.dev/badge/github.com/thorntonmc/go-sozluk.svg)](https://pkg.go.dev/github.com/thorntonmc/go-sozluk)
# go-sozluk
go library for the [sözlük of the Turkish Republic](https://sozluk.gov.tr/)

## usage

```go
c := s.NewClient()
k, err := c.Ara("yer")

if err != nil {
	log.Fatalln(err)
}

for _, v := range k {
	fmt.Printf("madde: %v\nBirlesikler: %v\n", v.Madde, v.Birlesikler)
}

```

## endpoints used

### ara

https://sozluk.gov.tr/gts?ara=${word}

### oneri

https://sozluk.gov.tr/oneri?soz=${word}
