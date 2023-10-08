package censys

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type APIInfo_Censys struct {
	Email      string `json:"email"`
	Login      string `json:"login"`
	FirstLogin string `json:"first_login"`
	LastLogin  string `json:"last_login"`
	Quota      struct {
		QuotaUsed      int    `json:"used"`
		QuotaAllowance int    `json:"allowance"`
		ResetsAt       string `json:"resets_at"`
	} `json:"quota"`
}

func (client *NewClient) APIInfo_Censys() (*APIInfo_Censys, error) {
	user_ID := client.id
	user_pass := client.secret
	//res, err := http.Get(fmt.Sprintf(""))
	htt_client := &http.Client{
		Timeout: time.Second * 10,
	}
	method := "GET"
	req, err := http.NewRequest(method, BaseURL+"/v1/account", nil)
	if err != nil {
		log.Fatalln("Request Unsuccessful: %s", err)
	}
	req.SetBasicAuth(user_ID, user_pass)
	response, err := htt_client.Do(req)
	if err != nil {
		log.Fatalln("Request Authorization Unsucessful: %s", err)
	}
	defer response.Body.Close()

	var ret APIInfo_Censys
	if err := json.NewDecoder(response.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
