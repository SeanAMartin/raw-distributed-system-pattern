package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
)

func RegisterService(r Registration) error {
	serviceUpdateURL, err := url.Parse(r.ServiceUpdateURL)

	if err != nil {
		return nil
	}

	http.Handle(serviceUpdateURL.Path, &serviceUpdateHandler{})
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err = enc.Encode(r)

	if err != nil {
		return err
	}

	res, err := http.Post(ServicesURL, "application/json", buf)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to register service. Registry Service responded with %v", res.StatusCode)
	}
	return nil
}

type serviceUpdateHandler struct{}

func (suh serviceUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	dec := json.NewDecoder(r.Body)
	var p patch
	err := dec.Decode(&p)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	prov.Update(p)
}

func ShutdownService(serviceURL string) error {
	req, err := http.NewRequest(http.MethodDelete, ServicesURL, bytes.NewBuffer([]byte(serviceURL)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "text/plain")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK || err != nil {
		return fmt.Errorf("Failed to deregister service. Registry Service responded with %v", res.StatusCode)
	}
	return err
}

type providers struct {
	services map[ServiceName][]string
	mutex    *sync.RWMutex
}

var prov = providers{
	services: make(map[ServiceName][]string),
	mutex:    new(sync.RWMutex),
}

func (p *providers) Update(pat patch) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, patchEntry := range pat.Added {
		if _, ok := p.services[patchEntry.Name]; !ok {
			p.services[patchEntry.Name] = make([]string, 0)
		}

		p.services[patchEntry.Name] = append(p.services[patchEntry.Name], patchEntry.URL)
	}

	for _, patchEntry := range pat.Removed {
		if providerURLs, ok := p.services[patchEntry.Name]; ok {
			for i := range providerURLs {
				if providerURLs[i] == patchEntry.URL {
					p.services[patchEntry.Name] = append(providerURLs[:i], providerURLs[i+1:]...)
				}
			}
		}
	}
}

func (p providers) get(name ServiceName) (string, error) {
	providers, ok := p.services[name]
	if !ok {
		return "", fmt.Errorf("No providers available for service %v", name)
	}

	index := int(rand.Float32() * float32(len(providers)))
	return providers[index], nil
}

func GetProvider(name ServiceName) (string, error) {
	return prov.get(name)
}
