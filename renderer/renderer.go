package renderer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ONSdigital/dp-dd-frontend-controller/config"
	"github.com/ONSdigital/go-ns/log"
	"io"
	"io/ioutil"
	"net/http"
)

// Render calls the configured front-end renderer service to render the given model with the given template.
func Render(model interface{}, template string) (renderedView []byte, err error) {
	body, err := json.Marshal(model)
	if err != nil {
		log.Error(err, nil)
		return
	}

	request, err := http.NewRequest("POST", config.RendererURL+"/"+template, bytes.NewReader(body))
	if err != nil {
		log.Error(err, nil)
		return
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.ErrorR(request, err, nil)
		return
	}

	defer checkClose(res.Body)

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("renderer.Render: unexpected status code from front-end renderer: %d (template: %s)", res.StatusCode, template)
		return
	}

	renderedView, err = ioutil.ReadAll(res.Body)

	return
}

func checkClose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Error(err, nil)
	}
}
