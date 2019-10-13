package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"

	/*only for debug*/
	//"github.com/davecgh/go-spew/spew"
)

type response struct {
	ID			string `json:"ID"`
	Response	string `json:"Response"`
}

type allResponse []response

var responseNotFound = response {
	ID: "",
	Response: "Response not found with specific ID",
}

var responseAlreadyExists = response {
	ID: "",
	Response: "Response already exist",
}

var responses = allResponse {
	{
		ID: "1",
		Response: "First response",
	},
}

func createResponse(w http.ResponseWriter, r *http.Request) {
	var newResponse response
	var hasResponse bool = false

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newResponse)
	
	for _, response := range responses {
		if newResponse.ID == response.ID {
			hasResponse = true
		}
	}

	if hasResponse == true {
		json.NewEncoder(w).Encode(responseAlreadyExists)
		return
	}

	responses = append(responses, newResponse)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newResponse)
}

func getOneResponse(w http.ResponseWriter, r *http.Request) {
	responseID := mux.Vars(r)["id"]

	if (responses == nil) {
		fmt.Fprintf(w, "Pleases, first create a new response")
	}

	if (responseID == "" && len(responseID) <= 0) {
		fmt.Fprintf(w, "Please, you must provided responseID")
	}

	var selectedResponse response

	for _, response := range responses {
		if response.ID == responseID {
			selectedResponse = response
			break
		}
	}

	if selectedResponse != (response{}) {
		json.NewEncoder(w).Encode(selectedResponse)
		return
	}

	json.NewEncoder(w).Encode(responseNotFound)
}

func getAllResponses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(responses)
}

func updateResponse(w http.ResponseWriter, r *http.Request) {
	responseID := mux.Vars(r)["id"]
	var updatedData response
	var responseUpdated bool = false

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "error on patch")
	}

	if (responses == nil) {
		fmt.Fprintf(w, "Pleases, first create a new response")
	}

	if (responseID == "" && len(responseID) <= 0) {
		fmt.Fprintf(w, "Please, you must provided responseID")
	}

	json.Unmarshal(reqBody, &updatedData)

	if (updatedData.Response == "" && len(updatedData.Response) <= 0) {
		fmt.Fprintf(w, "Please, you must provided response data")
	}

	for i, response := range responses {
		if responseID == response.ID {
			responses[i].Response = updatedData.Response
			json.NewEncoder(w).Encode(responses[i])
			responseUpdated = true
			break
		}
	}

	if responseUpdated == false {
		json.NewEncoder(w).Encode(responseNotFound)
	}
}

func deleteOneResponse(w http.ResponseWriter, r *http.Request) {
	responseID := mux.Vars(r)["id"]
	var responseExist bool = false

	if (responses == nil) {
		fmt.Fprintf(w, "Pleases, first create a new response")
	}

	if (responseID == "" && len(responseID) <= 0) {
		fmt.Fprintf(w, "Please, you must provided responseID")
	}

	for i, response := range responses {
		if responseID == response.ID {
			responses = append(responses[:i], responses[i+1:]...)
			json.NewEncoder(w).Encode("Response deleted")
			responseExist = true
			break
		}
	}

	if responseExist == false {
		json.NewEncoder(w).Encode(responseNotFound)
	}
}

func deleteAllResponses(w http.ResponseWriter, r *http.Request) {
	responses = allResponse{}
	json.NewEncoder(w).Encode(responses)	
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home, sweet home.")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/response", createResponse).Methods("POST")
	router.HandleFunc("/responses", getAllResponses).Methods("GET")
	router.HandleFunc("/response/{id}", getOneResponse).Methods("GET")
	router.HandleFunc("/response/{id}", updateResponse).Methods("PATCH")
	router.HandleFunc("/responses", deleteAllResponses).Methods("DELETE")
	router.HandleFunc("/response/{id}", deleteOneResponse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}