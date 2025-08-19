package hello

import "fmt"

var helloByLang = map[string]string{
	"pt": "Oi",
	"en": "Hello",
	"es": "Hola",
	"fr": "Salut",
	"de": "Hallo",
	"it": "Ciao",
	"nl": "Hallo",
	"ru": "Привет",
	"ja": "こんにちは",
	"zh": "你好",
	"ar": "مرحبا،",
	"tr": "Merhaba",
}

const (
	defaultGreet = "👋"
	defaultName  = "Anonymous"
)

func helloWithName(lang, name string) string {
	if name == "" {
		name = defaultName
	}

	if lang == "" {
		lang = defaultGreet
	}

	hello, ok := helloByLang[lang]

	if !ok {
		return fmt.Sprintf("%s, %s", defaultGreet, name)
	}

	return fmt.Sprintf("%s, %s", hello, name)
}
