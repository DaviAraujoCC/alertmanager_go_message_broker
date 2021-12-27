package controller

import (
	"api/models"
	"api/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	apiUrl   = os.Getenv("URL")
	endpoint = os.Getenv("ENDPOINT")
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var webhook structs.Webhook
	decoder.Decode(&webhook)
	endpoints := models.FindAllEndpoints()

	if len(endpoints) == 0 {
		fmt.Fprintf(w, "No endpoints found")
	} else {
		for _, endpoint := range endpoints {
			if endpoint.Alertname == "default" || strings.Contains(webhook.Alerts.Labels.Alertname, endpoint.Alertname) {
				log.Printf("Sending alert to endpoint: %v", endpoint.Endpoint)

				go func(url, endpoint string) {

					payload := structs.Payload_test{
						Alertname:   "test",
						Instance:    "localhost",
						Job:         "test",
						Status:      "firing",
						Description: "description test",
					}

					json, _ := json.Marshal(payload)
					url = url + "/" + endpoint

					// Insert logic here
					log.Println(string(json))

					// req, err := http.NewRequest("POST", url, strings.NewReader(string(json)))
					// if err == nil {
					// 	log.Println(err)
					// }
					// req.Header.Set("Content-Type", "application/json")

					// client := &http.Client{}
					// _, _ = client.Do(req)

				}(endpoint.Url, endpoint.Endpoint)

			}
		}
	}
}

func SenderHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var webhook structs.Webhook
	decoder.Decode(&webhook)
	endpoints := models.FindAllEndpoints()

	if len(endpoints) == 0 {
		fmt.Fprintf(w, "No endpoints found")
	} else {
		for _, endpoint := range endpoints {
			if endpoint.Alertname == "default" || strings.Contains(webhook.Alerts.Labels.Alertname, endpoint.Alertname) {
				log.Printf("Sending alert to endpoint: %v", endpoint.Endpoint)

				go func(url, endpoint string) {

					payload := structs.Payload_test{
						Alertname:   webhook.Alerts.Labels.Alertname,
						Instance:    webhook.Alerts.Labels.Instance,
						Job:         webhook.Alerts.Labels.Job,
						Status:      webhook.Alerts.Status,
						Description: webhook.Alerts.Annotations.Description,
					}

					json, _ := json.Marshal(payload)
					url = url + "/" + endpoint

					// Insert logic here

					log.Println(string(json))

					// req, err := http.NewRequest("POST", url, strings.NewReader(string(json)))
					// if err == nil {
					// 	log.Println(err)
					// }
					// req.Header.Set("Content-Type", "application/json")

					// client := &http.Client{}
					// _, _ = client.Do(req)

				}(endpoint.Url, endpoint.Endpoint)

			}
		}
	}
}

func EndpointHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var e structs.RequestCreateEndpoint

	err := decoder.Decode(&e)
	if err != nil {
		log.Println(err)
	}

	if e.Url != "" && e.Endpoint != "" {
		if e.Action == "add" {
			switch {
			case r.Method == "POST":
				if e.Alertname == "" {
					e.Alertname = "default"
				}
				if models.InsertEndpoint(e.Url, e.Endpoint, e.Alertname) {
					fmt.Fprintf(w, "Endpoint added successfully")
				} else {
					fmt.Fprintf(w, "Error adding endpoint, check if endpoint already exists")
				}
			default:
				fmt.Fprintf(w, "Method not allowed")
			}
		} else if e.Action == "delete" {
			switch {
			case r.Method == "POST":
				if models.DeleteEndpoint(e.ID) {
					fmt.Fprintf(w, "Endpoint deleted successfully")
				} else {
					fmt.Fprintf(w, "Error deleting endpoint, check if endpoint id already exists")
				}
			default:
				fmt.Fprintf(w, "Method not allowed")
			}
		}
	} else {
		fmt.Fprintf(w, "Bad request")
	}

}

func EndpointsHandler(w http.ResponseWriter, r *http.Request) {
	endpoints := models.FindAllEndpoints()

	if len(endpoints) == 0 {
		fmt.Fprintf(w, "No endpoints found")
	} else {
		for i, endpoint := range endpoints {
			var bodyText strings.Builder

			bodyText.WriteString(fmt.Sprintf("ID: %v \n", i+1) +
				fmt.Sprintf("Url: %v \n", endpoint.Url) +
				fmt.Sprintf("Endpoint: %v \n", endpoint.Endpoint) +
				fmt.Sprintf("Alertname: %v \n", endpoint.Alertname) +
				fmt.Sprintf(" \n"))

			fmt.Fprintf(w, bodyText.String())
		}

	}
}
