# go-sozluk
go library for the sozluk dictionary

## usage

```
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
