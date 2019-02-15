package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type requestJSON struct {
	Request struct {
		Object struct {
			Metadata struct {
				Labels map[string]string
			}
		}
	}
}

type responseJSON struct {
	Response struct {
		Allowed bool `json:"allowed"`
		Status  struct {
			Metadata struct{} `json:"metadata"`
			Message  string   `json:"message"`
		} `json:"status"`
	} `json:"response"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestJSON := requestJSON{}
		response := responseJSON{}
		responseJSON := []byte{}

		jsonBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal([]byte(jsonBody), &requestJSON)
		labels := requestJSON.Request.Object.Metadata.Labels

		// check if labels map contains 'owner' key
		if owner, ok := labels["owner"]; ok {
			response.Response.Allowed = true
			log.Println("Pod owner: " + owner)
			responseJSON, _ = json.Marshal(response)
		} else {
			response.Response.Allowed = false
			response.Response.Status.Message = "Every pod requires 'owner' label."
			log.Println("Could not find 'owner' labels. Provided labels: ")
			log.Println(labels)
			responseJSON, _ = json.Marshal(response)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	})
	log.Println("Listening on port 443...")
	log.Fatal(http.ListenAndServeTLS(":443", "/etc/webhook/certs/cert.pem", "/etc/webhook/certs/key.pem", nil))
}
