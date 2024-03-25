package book

import (
	"github.com/go-chi/render"
	"github.com/samurenkoroma/lha/internal/adapters"
	"net/http"
)

func List(writer http.ResponseWriter, request *http.Request) {
	render.JSON(writer, request, adapters.Books)
}

func ById(writer http.ResponseWriter, request *http.Request) {

}
func Create(writer http.ResponseWriter, request *http.Request) {

}
