package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"os"
	"yokanban-cli/internal/config"

	jose "github.com/dvsekhvalnov/jose2go"
	log "github.com/sirupsen/logrus"
)

// ServiceAccountCredentials represent credentials of a service account to create access token with.
type ServiceAccountCredentials struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	PublicKey   string `json:"publicKey"`
	PrivateKey  string `json:"privateKey"`
}

// JWT structure of a json web token.
type JWT struct {
	Iss   string `json:"iss"`
	Scope string `json:"scope"`
	Aud   string `json:"aud"`
}

// Auth the basic struct.
type Auth struct{}

// Authenticator is an auth interface.
type Authenticator interface {
	GetServiceAccountJWT() string
}

// NewAuthenticator creates a new instance of the Auth.
func NewAuthenticator() Authenticator {
	return &Auth{}
}

// GetServiceAccountJWT creates a JWT in order to retrieve access token from yokanban API.
func (a *Auth) GetServiceAccountJWT() string {
	serviceAccountCredentials, err := getServiceAccountCredentials()
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("Create JWT for serviceAccount: " + serviceAccountCredentials.ID)
	jwt := &JWT{
		Iss:   serviceAccountCredentials.ID,
		Aud:   config.GetAPIURL() + "/auth/oauth2/token",
		Scope: config.GetAPIScope(),
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
