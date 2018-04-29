package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/eItems", eItems).Methods("POST")
	fmt.Println("eItems is Ready")
	log.Fatal(http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), r))
	//log.Fatal(http.ListenAndServe("0.0.0.0:5000", r))

}

func eItems(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received request to get epTest details")
	var request WebHookRequest
	_ = json.NewDecoder(r.Body).Decode(&request)
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	fmt.Println(err)
	var speech = ""
	var displayText = ""

	intentName := request.Result.Metadata.IntentName

	if intentName == "myStatus" {
		speech = "There are 2 items in preparation. 1 item is in review with the office. Total amount available is 500k.2 items are pending for approval. Available balance 300k. For Item 1, report is due for submission next week."
		displayText = "There are 2 items in preparation. 1 item is in review with the office. Total amount available is 500k. 2 items are pending for approval. Available balance 300k. For Item 1, report is due for submission next week."
	} else if intentName == "anyNewUpdates" {
		speech = "grants.gov published 20 new opportunities and updated 40 opportunities."
		displayText = "grants.gov published 20 new opportunities and updated 40 opportunities."
	} else if intentName == "getNIHOpp" {
		speech = "70 opportunities posted"
		displayText = "70 opportunities posted"
	} else if intentName == "createProposal" {
		speech = "Not integrated with creation App"
		displayText = "Not integrated with creation App"
	}

	hookResp := WebHookResp{
		speech,
		displayText,
	}

	json.NewEncoder(w).Encode(hookResp)
}
