package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	yamlData, err := readYAML("./ex2/sitemap.yml")
	if err != nil {
		log.Fatal(err)
	}
	parsedMap := buildMap(yamlData)
	fmt.Println(parsedMap)

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	mHandler := mapHandler(parsedMap, mux)

	err = http.ListenAndServe(":8080", mHandler)
	if err != nil {
		log.Fatal(err)
	}
}

func mapHandler(pathMap map[string]string, handle http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wanted := r.URL.Path

		fmt.Println(wanted)
		if _, ok := pathMap[wanted]; ok {
			fmt.Println(pathMap[wanted])
			http.Redirect(w, r, pathMap[wanted], http.StatusFound)
			return
		}
		handle.ServeHTTP(w, r)
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Heelo there, wanted to go to ")
}

func readYAML(path string) (map[interface{}]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func buildMap(in map[interface{}]interface{}) map[string]string {
	rm := make(map[string]string, len(in))
	for k, v := range in {
		rm[fmt.Sprintf("%v", k)] = fmt.Sprintf("%v", v)
	}
	return rm
}
