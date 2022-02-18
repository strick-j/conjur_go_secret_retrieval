package main

import (
	"flag"
	"fmt"

	conjurIamClient "github.com/strick-j/conjur-authn-iam-go-client"
)

func main() {
	methodPtr := flag.String("method", "", "IAM Method (e.g. static, iamrole, assumerole)")
	profilePtr := flag.String("profile", "", "IAM Profile to use (e.g. Default)")
	rolePtr := flag.String("role", "", "AWS Role ARN")
	accessKeyPtr := flag.String("accesskey", "", "AWS AKID")
	secretKeyIdPtr := flag.String("secretkey", "", "AWS Secret Key Id")
	sessionTokenPtr := flag.String("sessiontoken", "", "AWS Session Token")
	variableIdPtr := flag.String("variableid", "", "Variable Id (e.g. policy/path/variable-id)")
	flag.Parse()

	params := &conjurIamClient.ConjurIamParams{
		IamAuthMethod: *methodPtr,
		Profile:       *profilePtr,
		RoleArn:       *rolePtr,
		AccessKey:     *accessKeyPtr,
		SecretKey:     *secretKeyIdPtr,
		SessionToken:  *sessionTokenPtr,
	}

	// Retrieve Conjur Client based on IAM Role
	conjurClient, err := params.NewConjurIamClient()
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
