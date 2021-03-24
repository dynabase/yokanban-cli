package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	jose "github.com/dvsekhvalnov/jose2go"
	"io/ioutil"
	"log"
	"os"
	"yokanban-cli/internal/config"
)

type ServiceAccountCredentials struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	PublicKey   string `json:"publicKey"`
	PrivateKey  string `json:"privateKey"`
}

type Jwt struct {
	Iss   string `json:"iss"`
	Scope string `json:"scope"`
	Aud   string `json:"aud"`
}

/**
Create a JWT in order to retrieve access token from yokanban API.
*/
func GetServiceAccountJwt() string {
	serviceAccountCredentials, err := getServiceAccountCredentials()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Create JWT for serviceAccount %s\n", serviceAccountCredentials.Id)
	jwt := &Jwt{
		Iss:   serviceAccountCredentials.Id,
		Aud:   config.GetApiUrl() + "/auth/oauth2/token",
		Scope: "test user board",
	}

	payload, err := json.Marshal(jwt)
	if err != nil {
		log.Fatal(err)
	}

	token, err := jose.Sign(string(payload), jose.RS256, getRsaPrivateKey(serviceAccountCredentials))
	if err != nil {
		log.Fatal(err)
	}

	return token
}

func getServiceAccountCredentials() (ServiceAccountCredentials, error) {
	apiKeyPath, err := config.GetApiKeysPath()
	if err != nil {
		return ServiceAccountCredentials{}, err
	}

	jsonFile, err := os.Open(apiKeyPath)
	if err != nil {
		return ServiceAccountCredentials{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var serviceAccountCredentials ServiceAccountCredentials

	err = json.Unmarshal(byteValue, &serviceAccountCredentials)
	if err != nil {
		return ServiceAccountCredentials{}, err
	}

	return serviceAccountCredentials, nil
}

func getRsaPrivateKey(serviceAccountCredentials ServiceAccountCredentials) *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(serviceAccountCredentials.PrivateKey))
	parseResult, _ := x509.ParsePKCS8PrivateKey(block.Bytes)
	key := parseResult.(*rsa.PrivateKey)

	return key
}
