package hello

import(
	"encoding/json"
	"net/http"
	"github.com/ONSdigital/dp-frontend-models/model/hello"
	"github.com/ONSdigital/go-ns/log"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	model := hello.Page{
		Greeting: "Hello from Controller",
	}

	data, err := json.Marshal(model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Uh oh..."))
		log.ErrorR(req, err, nil)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
