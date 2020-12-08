package main

import (
	"encoding/json"
	"fmt"

	"github.com/ga4gh/drs-reference-server/internal/drsconstants"

	"github.com/ga4gh/drs-reference-server/internal/drsmodel"
)

func main() {
	fmt.Println("Starting Program")

	checksum := drsmodel.NewChecksum()
	checksum.SetChecksum("1234abab")
	checksum.SetType(drsconstants.ChecksumTypeSHA256)

	bytes, _ := json.Marshal(checksum)
	fmt.Println(string(bytes))

	// fmt.Println("Server started on port 3000")
	// http.ListenAndServe(":3000", nil)
}
