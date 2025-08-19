package hello

import "testing"

func TestHelloWithName(t *testing.T) {
	// Quando há muitos casos de test, é normal colocar em uma lista de structs
	tests := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "Portuguese greet",
			got:  helloWithName("pt", "João"),
			want: "Oi, João",
		},
		{
			name: "English greet",
			got:  helloWithName("en", "João"),
			want: "Hello, João",
		},
		{
			name: "Spanish greet",
			got:  helloWithName("es", "João"),
			want: "Hola, João",
		},
		{
			name: "French greet",
			got:  helloWithName("fr", "João"),
			want: "Salut, João",
		},
		{
			name: "German greet",
			got:  helloWithName("de", "João"),
			want: "Hallo, João",
		},
		{
			name: "Italian greet",
			got:  helloWithName("it", "João"),
			want: "Ciao, João",
		},
		{
			name: "Dutch greet",
			got:  helloWithName("nl", "João"),
			want: "Hallo, João",
		},
		{
			name: "Russian greet",
			got:  helloWithName("ru", "João"),
			want: "Привет, João",
		},
		{
			name: "Japanese greet",
			got:  helloWithName("ja", "João"),
			want: "こんにちは, João",
		},
		{
			name: "Chinese greet",
			got:  helloWithName("zh", "João"),
			want: "你好, João",
		},
		{
			name: "Arabic greet",
			got:  helloWithName("ar", "João"),
			want: "مرحبا،, João",
		},
		{
			name: "Turkish greet",
			got:  helloWithName("tr", "João"),
			want: "Merhaba, João",
		},
		{
			name: "Empty string name",
			got:  helloWithName("pt", ""),
			want: "Oi, Anonymous",
		},
		{
			name: "Empty language",
			got:  helloWithName("", "João"),
			want: "👋, João",
		},
		{
			name: "Invalid language",
			got:  helloWithName("www", "João"),
			want: "👋, João",
		},
		{
			name: "Empty all",
			got:  helloWithName("", ""),
			want: "👋, Anonymous",
		},
	}

	for _, tc := range tests {

		/*
		*  O `t.Run` executará um subtest dentro do teste principal,
		*  com isso é mais fácil de identificar o que está sendo testado por conta do `name` presente em cada subtest
		* */
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Errorf("got: %q | want: %q\n", tc.got, tc.want)
			}
		})
	}
}
