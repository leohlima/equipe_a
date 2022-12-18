package main

import (
	"fmt"
	"log"
	"net/http"
	"produto/src/config"
	"produto/src/router"
	"produto/webui"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	config.Carregar()

	r := mux.NewRouter()

	n := negroni.New(
		negroni.NewLogger(),
	)

	webui.RegisterUIHandlers(r, n)

	r = router.Gerar(r)

	fmt.Printf("Escutando na porta %d: ", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
