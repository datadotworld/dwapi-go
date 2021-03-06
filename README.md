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
	"path/filepath"
	"regexp"
	"strings"

	"github.com/datadotworld/dwapi-go/dwapi"
)

func main() {
	// new client
	token := os.Getenv("DW_AUTH_TOKEN")
	dw := dwapi.NewClient(token)

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
	owner := "my-username"
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
	testFilePath := filepath.Join(os.TempDir(), "test-file.csv")
	if err = ioutil.WriteFile(testFilePath, []byte(sj), 0644); err != nil {
		fmt.Fprintln(os.Stderr, "Dataset.UploadFile() returned an error while creating a file:", err)
		os.Exit(1)
	}
	defer os.Remove(testFilePath)

	uploadResp, err := dw.Dataset.UploadFile(owner, datasetid, "test-file.csv", testFilePath, false)
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

## Changing the hostname

The API calls are made to `https://api.data.world` by default, but the URL can be changed by setting the `DW_API_HOST` environment variable.

For customers in a single-tenant environment, you can also use the `DW_ENVIRONMENT` variable to alter the default URL. For example, for the customer `customer`, setting it will alter the URL to `https://api.customer.data.world`.

Additionally, the hostname can also be changed by explicitly setting the `BaseURL` property of the client, i.e.:
```
dw = dwapi.NewClient("token")
dw.BaseURL = "http://localhost:1010/v0"
```
_Notice that the stage also needs to be set if going down this path._
