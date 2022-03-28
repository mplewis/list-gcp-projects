package main

import (
	"context"
	"fmt"

	crm "google.golang.org/api/cloudresourcemanager/v1"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	crmSvc, err := crm.NewService(context.Background())
	check(err)

	resp, err := crmSvc.Projects.List().Do()
	check(err)
	fmt.Println(resp)
}
