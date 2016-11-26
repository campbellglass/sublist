package main

import (
	"bytes"
	"encoding/json"
)

// PrettyPrint prints JSON in a pretty format
func PrettyPrint(JSON []byte) (string, error) {
	var pretty bytes.Buffer
	err := json.Indent(&pretty, JSON, "", "\t")
	if err != nil {
		return "", err
	}
	return pretty.String(), nil

}
