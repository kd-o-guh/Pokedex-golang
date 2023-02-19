package manipulador

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetPokemon(pokemon string) string {

	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon/"+pokemon, nil)

	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o Servidor. Erro:", err.Error())
	}

	req.Header.Set("User-Agent", "pokego")

	res, err := cliente.Do(req)

	if err != nil {
		fmt.Println("[main] Erro ao abrir a página do Servidor. Erro:", err.Error())
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return string(body)
}
