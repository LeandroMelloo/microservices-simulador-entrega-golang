package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Driver struct {
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Drivers struct {
	Drivers []Driver // slice of driver
}

func GetDrivers() []byte {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		panic(err.Error()) // em caso de erro finaliza o processo
	}

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile) // ReadAll lê os dados
	if err != nil {
		panic(err.Error()) // em caso de erro finaliza o processo
	}

	return data // dats -> retorna os dados para o cliente
}

func ShowDrivers(w http.ResponseWriter, r *http.Request) {
	drivers := GetDrivers() // busca todos os dados do data.json
	w.Write([]byte(drivers))
}

func GetDriversByUuid(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)
	data := GetDrivers()

	var drivers Drivers
	json.Unmarshal(data, &drivers)

	for _, d := range drivers.Drivers {
		if d.Uuid == query["uuid"] {
			driver, _ := json.Marshal(d)
			w.Write([]byte(driver))
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/drivers", ShowDrivers)
	r.HandleFunc("/drivers/{uuid}", GetDriversByUuid)

	http.ListenAndServe(":8089", r)
}
