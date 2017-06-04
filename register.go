package main

import (
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"log"
	"os"
)

type Settings struct {
	CloudFlare CloudFlareSettings
}

type CloudFlareSettings struct {
	ApiKey  string
	Email   string
	Website string
}

func getSettings() Settings {
	ready := true
	settings := Settings{
		CloudFlare: CloudFlareSettings{},
	}
	settings.CloudFlare.ApiKey = os.Getenv("CLOUDFLARE_API_KEY")
	if len(settings.CloudFlare.ApiKey) <= 0 {
		fmt.Println("Missing 'CLOUDFLARE_API_KEY'")
		ready = false
	}
	settings.CloudFlare.Email = os.Getenv("CLOUDFLARE_EMAIL")
	if len(settings.CloudFlare.Email) <= 0 {
		fmt.Println("Missing 'CLOUDFLARE_EMAIL'")
		ready = false
	}
	settings.CloudFlare.Website = os.Getenv("CLOUDFLARE_WEBSITE")
	if len(settings.CloudFlare.Website) <= 0 {
		fmt.Println("Missing 'CLOUDFLARE_WEBSITE'")
		ready = false
	}
	if !ready {
		os.Exit(1)
	}
	return settings
}

func main() {
	settings := getSettings()
	api, err := cloudflare.New(settings.CloudFlare.ApiKey, settings.CloudFlare.Email)
	if err != nil {
		log.Fatal(err)
	}
	zoneId, err := api.ZoneIDByName(settings.CloudFlare.Website)
	if err != nil {
		log.Fatal(err)
	}
	dnsRecord := cloudflare.DNSRecord{
		Type:    "A",
		Name:    "testing",
		Content: "1.2.3.4",
	}
	api.CreateDNSRecord(zoneId, dnsRecord)
	fmt.Printf("%v", api)
}
