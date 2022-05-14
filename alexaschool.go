package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/thefriendlycoder/schoolscraper"
)

type Connection struct{}

// Helper function that generates an Alexa response describing a school
// closure based on a specific school identified by the name and district
func checkClosureStatus(districtName string, schoolName string) alexa.Response {
	// Load raw HTTML data from the school website
	resp, err := http.Get(schoolscraper.ScheduleURL)
	if err != nil {
		return alexa.NewSimpleResponse("Failed", "Failed to load school data from URL")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return alexa.NewSimpleResponse("Failed", "Failed to read school data from URL")
	}

	// Scrape the HTML content to extract school property data
	school, err := schoolscraper.GetSchoolStatus(string(body), districtName, schoolName)
	if err != nil {
		return alexa.NewSimpleResponse("Failed", "Failed to parse school data from URL")
	}

	// Generate an appropriate Alexa response depending on the school status
	var message string
	if school.IsOpen() {
		message = "School is open"
	} else {
		message = "School is closed"
	}

	if school.AllBusesOnTime() {
		message += " and all buses are running on time"
	} else {
		message += " and there are some late buses"
	}
	return alexa.NewSimpleResponse("School Status", message)
}

// Callback for handling Alexa requests triggered by "intents" defined
// by the Alexa app interface. See the JSON files under the
// ./skill-package/interactionModels folder for details
func (connection Connection) IntentDispatcher(ctx context.Context, request alexa.Request) (alexa.Response, error) {
	district := "FREDERICTON"
	school := "École Sainte-Anne"
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "SchoolOpenIntent", "":
		response = checkClosureStatus(district, school)
	default:
		msg := fmt.Sprintf("Unrecognized intent %v", request.Body.Intent.Name)
		response = alexa.NewSimpleResponse("Unknown Request", msg)
	}
	return response, nil
}

// Entrypoint function
func main() {
	connection := Connection{}
	lambda.Start(connection.IntentDispatcher)
}
