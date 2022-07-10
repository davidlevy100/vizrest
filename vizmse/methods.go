package vizmse

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

func CreateClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}

	return httpClient
}

func Get(client *http.Client, url string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}

func GetBase(client *http.Client, url string) (map[string]string, error) {

	body, err := Get(client, url)
	if err != nil {
		return nil, err
	}

	var s Service

	err = xml.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}

	result := map[string]string{}

	for _, thisElement := range s.Workspace.Collection {
		result[thisElement.Categories.Category.Term] = thisElement.Href

	}

	return result, nil

}
