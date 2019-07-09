# dwapi-go

[![GoDoc](https://godoc.org/github.com/datadotworld/dwapi-go/dwapi?status.svg)](https://godoc.org/github.com/datadotworld/dwapi-go/dwapi)

This package makes it easy to use [data.world's REST API](https://apidocs.data.world/api) with Go.

Users can:
* Create and update datasets, projects, metadata, and files
* Query datasets using SQL and SPARQL
* Download files and entire datasets

## Installation

```bash
go get github.com/datadotworld/dwapi-go/dwapi
```

## Usage
The full package documentation is available at https://godoc.org/github.com/datadotworld/dwapi-go/dwapi.

You can also check out the API documentation at https://apidocs.data.world/api for specifics on the endpoints.

```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/datadotworld/dwapi-go/dwapi"
)

func main() {
	// new client
	owner := "my-username"
	token := os.Getenv("DW_AUTH_TOKEN") // token from https://data.world/settings/advanced"
	dw := dwapi.NewClient(owner, token)

	// get info on the current user
	user, err := dw.User.Self()
	if err != nil {
		fmt.Fprintln(os.Stderr, "User.Self() returned an error:", err)
		os.Exit(1)
	}
	fmt.Println("Name:", user.DisplayName)
	fmt.Println("Creation Date:", user.Created)
	fmt.Println("-----")
	/* output:
	Name: My Display Name
	Creation Date: 2016-07-13T23:38:44.026Z
	-----
	*/

	// create a new dataset
	request := dwapi.DatasetCreateRequest{
		Title:       "My Awesome Dataset",
		Description: "A short description",
		Summary:     "A long description",
		Tags:        []string{"first", "puppies and kittens"},
		License:     "PDDL",
		Visibility:  "PRIVATE",
	}
	createResp, err := dw.Dataset.Create(owner, &request)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Dataset.Create() returned an error:", err)
		os.Exit(1)
	}
	fmt.Println("Create Response Message:", createResp.Message)
	fmt.Println("Dataset URI:", createResp.URI)
	fmt.Println("-----")
	/* output:
	Response Message: Dataset created successfully.
	Dataset URI: https://data.world/my-username/my-awesome-dataset
	-----
	*/

	// retrieve the metadata for a dataset
	pattern := regexp.MustCompile(`https://data.world/(?:.*)/(.*)`)
	datasetid := pattern.FindStringSubmatch(createResp.URI)[1]

	retrieveResp, err := dw.Dataset.Retrieve(owner, datasetid)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Dataset.Retrieve() returned an error:", err)
		os.Exit(1)
	}
	fmt.Println("Title:", retrieveResp.Title)
	fmt.Println("Description:", retrieveResp.Description)
	fmt.Println("Access Level:", retrieveResp.AccessLevel)
	fmt.Println("Creation Date:", retrieveResp.Created)
	fmt.Println("Last Updated Date:", retrieveResp.Updated)
	fmt.Println("Dataset Status:", retrieveResp.Status)
	fmt.Println("-----")
	/* output:
	Title: My Awesome Dataset
	Description: A short description
	Access Level: ADMIN
	Creation Date: 2018-11-21T17:32:40.057Z
	Last Updated Date: 2018-11-21T17:32:40.057Z
	Dataset Status: NEW
	-----
	*/

	// upload a file
	s := []string{"first_name,last_name", "Abe,Marcos", "Abby,Johnson"}
	sj := strings.Join(s, "\n")
	if err = ioutil.WriteFile("/tmp/test-file.csv", []byte(sj), 0644); err != nil {
		fmt.Fprintln(os.Stderr, "Dataset.UploadFile() returned an error while creating a file:", err)
		os.Exit(1)
	}

	uploadResp, err := dw.Dataset.UploadFile(owner, datasetid, "test-file.csv", "/tmp/test-file.csv", false)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Dataset.UploadFile() returned an error:", err)
		os.Exit(1)
	}
	fmt.Println("Upload Response Message:", uploadResp.Message)
	fmt.Println("-----")
	/* output:
	Response Message: File uploaded.
	-----
	*/

	// delete a dataset
	deleteResp, err := dw.Dataset.Delete(owner, datasetid)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Dataset.Delete() returned an error:", err)
		os.Exit(1)
	}
	fmt.Println("Delete Response Message:", deleteResp.Message)
	fmt.Println("-----")
	/* output:
	Delete Response Message: Dataset has been successfully deleted.
	-----
	*/
}
```
