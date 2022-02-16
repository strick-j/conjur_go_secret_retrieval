package main

import (
	"flag"
	"fmt"

	conjurIamClient "github.com/strick-j/conjur-authn-iam-go-client"
)

func main() {
	hostIdPtr := flag.String("hostid", "", "Conjur Host Id (e.g. host/policy/prefix/id")
	serviceIdPtr := flag.String("serviceid", "", "Conjur Service ID (e.g. prod)")
	variableIdPtr := flag.String("variableid", "", "Variable Id (e.g. policy/path/variable-id)")
	flag.Parse()

	details := &conjurIamClient.ConjurDetails{
		HostId:    *hostIdPtr,
		ServiceId: *serviceIdPtr,
	}

	// Retrieve Conjur Client based on IAM Role
	conjurClient, err := conjurIamClient.NewClientFromRole(*details)
	if err != nil {
		fmt.Printf("error creating client : %s", err)
	}

	// Retrieve Secret using Conjur Client
	secretValue, err := conjurClient.RetrieveSecret(*variableIdPtr)
	if err != nil {
		fmt.Printf("error retriveing secret : %s", err)
	}
	fmt.Printf("Secret Value: %s", string(secretValue))
}
