package test

import (
	"crypto/tls"
	"github.com/aiziyuer/registry/client/registry"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"testing"
)

var client *registry.Registry

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client = registry.NewClient(&http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}, &registry.Endpoint{
		Schema: os.Getenv("REGISTRY_SCHEMA"),
		Host:   os.Getenv("REGISTRY_HOST"),
	}, &registry.BasicAuth{
		UserName: os.Getenv("REGISTRY_USERNAME"),
		PassWord: os.Getenv("REGISTRY_PASSWORD"),
	})
}

func TestClient(t *testing.T) {
	_ = client.Ping()
}

func TestTags(t *testing.T) {
	output, _ := client.TagsWithPretty("aiziyuer/centos")
	logrus.Info(output)
}
