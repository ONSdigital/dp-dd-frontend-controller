package hello

import(
	"net/http"
	"github.com/ONSdigital/dp-frontend-models/model/hello"
	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/dp-dd-frontend-controller/renderer"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	// Generate model
	model := hello.Page{
		Greeting: "Hello from Controller",
	}

	// Now call renderer to render view
	view, err := renderer.Render(model, "hello")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to marshall model"))
		log.ErrorR(req, err, nil)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(view)
}
