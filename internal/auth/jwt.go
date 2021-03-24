package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	jose "github.com/dvsekhvalnov/jose2go"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"yokanban-cli/internal/config"
)

// ServiceAccountCredentials represent credentials of a service account to create access token with.
type ServiceAccountCredentials struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	PublicKey   string `json:"publicKey"`
	PrivateKey  string `json:"privateKey"`
}

// Jwt structure of a json web token.
type Jwt struct {
	Iss   string `json:"iss"`
	Scope string `json:"scope"`
	Aud   string `json:"aud"`
}

// GetServiceAccountJwt creates a JWT in order to retrieve access token from yokanban API.
func GetServiceAccountJwt() string {
	serviceAccountCredentials, err := getServiceAccountCredentials()
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("Create JWT for serviceAccount: " + serviceAccountCredentials.ID)
	jwt := &Jwt{
		Iss:   serviceAccountCredentials.ID,
		Aud:   config.GetAPIURL() + "/auth/oauth2/token",
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
	apiKeyPath, err := config.GetAPIKeysPath()
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
