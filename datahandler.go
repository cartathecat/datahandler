/*
Package datahandler for Ports
*/
package datahandler

/*
	Contributors:
		Mick Moriarty
		v0.0.1
*/
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*
Port structure
*/
type Port struct {
	Name        string    `json:"name"`        // Name of port
	City        string    `json:"city"`        // City of port
	Country     string    `json:"country"`     // Country of port
	Alias       []string  `json:"alias"`       // Alias'
	Regions     []string  `json:"regions"`     // Regions
	Coordinates []float64 `json:"coordinates"` // Long/Lat coordinates
	Province    string    `json:"province"`    // Province
	TimeZone    string    `json:"timezone"`    // Time zone
	Unlocs      []string  `json:"unlocs"`      // Unlocs
	Code        string    `json:"code"`        // Code
}

/*
PortID structure
*/
type PortID struct {
	ID string `json:"Id"` // ID of port
}

// Datastore
var result map[string]Port

/*
GetPort function to return a single port
*/
func (key *PortID) GetPort() (Port, error) {

	var (
		portlookup Port
		ok         bool
		err        error
	)

	if portlookup, ok = result[key.ID]; !ok {
		err = fmt.Errorf("Lookup value not found for %s", key.ID)
	}
	return portlookup, err
}

/*
ListOfPorts functions to return all port id's
*/
func ListOfPorts() []PortID {

	resp := make([]PortID, 0)

	// Loop through list of ports ... extract the name (idx)
	for idx := range result {
		portID := PortID{}
		portID.ID = idx
		resp = append(resp, portID)
	}

	return resp
}

/*
ListAllPorts is a function to return a list of ALL ports
*/
func ListAllPorts() map[string]interface{} {

	resp := map[string]interface{}{}

	// Loop through the data ...
	for idx := range result {
		port := result[idx]
		resp[idx] = port
	}

	return resp
}

/*
LoadData from  a JSON file ...
*/
func LoadData() error {

	log.Printf("Loading JSON data file ....")

	//	Open data file
	jsonFile, err := os.Open("ports.json")
	if err != nil {
		log.Fatalf("Error opening ports.json file, error : %s", err)
	}
	defer jsonFile.Close()

	// Read all the data
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &result)

	log.Print("Data file loaded ....")
	return nil

}
