package censys

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Location struct {
	Continent             string     `json:"continent"`
	Country               string     `json:"country"`
	CountryCode           string     `json:"country_code"`
	PostalCode            string     `json:"postal_code"`
	Timezone              string     `json:"timezone"`
	Coordinates           Coordinate `json:"coordinates"`
	RegisteredCountry     string     `json:"registered_country"`
	RegisteredCountryCode string     `json:"registered_country_code"`
}

type AutonomousSystem struct {
	ASN         int    `json:"asn"`
	Description string `json:"description"`
	BGPPrefix   string `json:"bgp_prefix"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
}

type Service struct {
	Port              int    `json:"port"`
	ServiceName       string `json:"service_name"`
	TransportProtocol string `json:"transport_protocol"`
	Certificate       string `json:"certificate,omitempty"`
}

type Hit struct {
	Name             string           `json:"name,omitempty"`
	IP               string           `json:"ip"`
	Services         []Service        `json:"services"`
	Location         Location         `json:"location"`
	AutonomousSystem AutonomousSystem `json:"autonomous_system"`
}

type Result struct {
	Query string `json:"query"`
	Total int    `json:"total"`
	Hits  []Hit  `json:"hits"`
}

type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result Result `json:"result"`
	Links  struct {
		Prev string `json:"prev"`
		Next string `json:"next"`
	} `json:"links"`
}

// Your JSON data goes here
func (client *NewClient) HostSearch(serv_name string) (*Response, error) {
	user_ID := client.id
	user_pass := client.secret
	//res, err := http.Get(fmt.Sprintf(""))
	htt_client := &http.Client{
		Timeout: time.Second * 10,
	}
	method := "GET"
	req, err := http.NewRequest(method, BaseURL+"/v2/hosts/search?q=services.service_name: "+serv_name+"&per_page=100&virtual_hosts=EXCLUDE", nil)
	if err != nil {
		log.Fatalln("Request Unsuccessful: %s", err)
	}
	req.SetBasicAuth(user_ID, user_pass)
	response, err := htt_client.Do(req)
	if err != nil {
		log.Fatalln("Request Authorization Unsucessful: %s", err)
	}
	defer response.Body.Close()
	var ret1 Response
	if err := json.NewDecoder(response.Body).Decode(&ret1); err != nil {
		return nil, err
	}
	return &ret1, nil
}
