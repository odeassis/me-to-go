package main

import "testing"

func TestHello(t *testing.T) {

	checkCorrectMessage := func(t *testing.T, result string, expect string) {
		/*
			t.Helper () é necessário para dizermos ao conjunto de testes que este é
			um método auxiliar. Ao fazer isso, quando o teste falhar, o número da linha
			relatada estará em nossa chamada de função, e não dentro do nosso auxiliar de teste.
			Isso ajudará outras pessoas a entenderem melhor o que está acontecendo quando falhamos em nosso teste.
		*/
		t.Helper()

		if result != expect {
			t.Errorf("result '%s', expect '%s'", result, expect)
		}
	}

	t.Run("say hello for people", func(t *testing.T) {
		result := hello("Francisco")
		expect := "Hello, Francisco"

		checkCorrectMessage(t, result, expect)
	})

	t.Run("say 'Hello, word' if hello paramater is empty", func(t *testing.T) {
		result := hello("")
		expect := "Hello, word"

		checkCorrectMessage(t, result, expect)
	})
}
