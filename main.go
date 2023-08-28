package main

import (
	"censys-layout/censys"
	"fmt"
	"log"
	"os"
)

func main() {
	id := os.Getenv("CENSYS_ID")
	secret := os.Getenv("CENSYS_SECRET")
	client := censys.New(id, secret)

	info, err := client.APIInfo_Censys()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("CENSYS API INFO \nEmail: %s\nFirst Login:%s\nLast Login:%s\nQuota Used:%d\nQuota Allowance:%d\nResets At:%s", info.Email, info.FirstLogin, info.LastLogin, info.Quota.QuotaUsed, info.Quota.QuotaAllowance, info.Quota.ResetsAt)
	//Working upto here
}
