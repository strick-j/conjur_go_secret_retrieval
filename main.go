package main

import (
	"flag"
	"fmt"

	conjurIamClient "github.com/strick-j/conjur-authn-iam-go-client"
)

func main() {
	methodPtr := flag.String("methodid", "", "IAM Method (e.g. static, iamrole, assumerole)")
	profilePtr := flag.String("profile", "", "IAM Profile to use (e.g. Default)")
	sessionIdPtr := flag.String("sessionid", "", "Session ID for Role Assumption (e.g. Default)")
	accessKeyPtr := flag.String("accesskeyid", "", "AWS AKID")
	secretKeyIdPtr := flag.String("secretkeyid", "", "AWS Secret Key Id")
	sessionTokenPtr := flag.String("sessiontoken", "", "AWS Session Token")
	hostIdPtr := flag.String("hostid", "", "Conjur Host Id (e.g. host/policy/prefix/id)")
	serviceIdPtr := flag.String("serviceid", "", "Conjur Service ID (e.g. prod)")
	variableIdPtr := flag.String("variableid", "", "Variable Id (e.g. policy/path/variable-id)")
	flag.Parse()

	params := &conjurIamClient.ConjurParams{
		IamAuthMethod: *methodPtr,
		Profile:       *profilePtr,
		Session:       *sessionIdPtr,
		AccessKey:     *accessKeyPtr,
		SecretKey:     *secretKeyIdPtr,
		SessionToken:  *sessionTokenPtr,
		HostId:        *hostIdPtr,
		ServiceId:     *serviceIdPtr,
	}

	// Retrieve Conjur Client based on IAM Role
	conjurClient, err := conjurIamClient.NewConjurIamClient(*params)
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
