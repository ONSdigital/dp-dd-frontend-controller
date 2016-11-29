package renderer

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/ONSdigital/dp-dd-frontend-controller/config"
	"encoding/json"
	"github.com/ONSdigital/go-ns/log"
	"bytes"
)

// Call front-end renderer to render the given model with the given template
func Render(model interface{}, template string) (renderedView []byte, err error) {
	body, err := json.Marshal(model)
	if err != nil {
		log.Error(err, nil)
		return
	}

	request, err := http.NewRequest("POST", config.RendererURL + "/" + template, bytes.NewReader(body))
	if err != nil {
		log.Error(err, nil)
		return
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.ErrorR(request, err, nil)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("Handler.handler: unexpected status code: %d", res.StatusCode)
		return
	}

	renderedView, err = ioutil.ReadAll(res.Body)

	return
}
