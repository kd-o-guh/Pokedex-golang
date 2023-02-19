package manipulador

import (
	"fmt"
	"net/http"
	"pokedex-golang/model"
	"strconv"
	"text/template"

	"github.com/tidwall/gjson"
)

var (
	templates = template.Must(template.ParseFiles("html/index.html"))
)

func Render(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	if r.Method == "GET" {
		index := model.Data{}

		index.Name = "Bulbasaur"
		index.Number = 1
		index.StyleHeight = "style=&#34 height: 7% &#34"
		index.Image = "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/versions/generation-v/black-white/animated/1.gif"

		if err := templates.Execute(w, index); err != nil {
			http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
			fmt.Println("[Ola] Erro na execucao do modelo: ", err.Error())
		}

	} else {

		r.ParseForm()

		pokemon := r.Form["search"]

		jsonData := GetPokemon(pokemon[0])

		name := gjson.Get(jsonData, "name")

		id := gjson.Get(jsonData, "id")

		height := gjson.Get(jsonData, "height")

		image := gjson.Get(jsonData, "sprites.versions.generation-viii.icons.front_default")

		fmt.Println("pokemon:", name)
		fmt.Println("teste:")

		stylestring := "style= height:" + strconv.FormatFloat(height.Num, 'f', -1, 32) + "%"

		index := model.Data{}
		index.Name = name.Str
		index.StyleHeight = stylestring
		fmt.Println(index.StyleHeight)
		index.Number = float64(id.Int())
		index.Image = image.Str

		if err := templates.Execute(w, index); err != nil {
			http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
			fmt.Println("[Ola] Erro na execucao do modelo: ", err.Error())
		}

	}

}
