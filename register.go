package main

import (
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Settings struct {
	CloudFlare CloudFlareSettings
	Subdomain  string
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
	settings.Subdomain = os.Getenv("SUBDOMAIN")
	if len(settings.Subdomain) <= 0 {
		fmt.Println("Missing 'SUBDOMAIN'")
		ready = false
	}
	if !ready {
		os.Exit(1)
	}
	return settings
}

func getPublicIP() (string, error) {
	res, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	ip, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(ip), nil
}

func registerCloudFlare(api *cloudflare.API, zoneId string, publicIP string, settings Settings) (string, error) {
	dnsRecord := cloudflare.DNSRecord{
		Type:    "A",
		Name:    settings.Subdomain,
		Content: publicIP,
	}
	_, err := api.CreateDNSRecord(zoneId, dnsRecord)
	if err != nil {
		return "", err
	}
	return "Registered 'A " + settings.Subdomain + "." + settings.CloudFlare.Website + " " + publicIP + "'", nil
}

func unregisterCloudFlare(api *cloudflare.API, zoneId string, publicIP string, settings Settings) (string, error) {
	dnsRecords, err := api.DNSRecords(zoneId, cloudflare.DNSRecord{Type: "A"})
	if err != nil {
		return "", err
	}
	for i := 0; i < len(dnsRecords); i++ {
		dnsRecord := dnsRecords[i]
		if dnsRecord.Content == publicIP {
			err := api.DeleteDNSRecord(zoneId, dnsRecord.ID)
			if err != nil {
				return "", err
			}
			break
		}
	}
	return "Unregistered 'A " + settings.Subdomain + "." + settings.CloudFlare.Website + " " + publicIP + "'", nil
}

func main() {
	settings := getSettings()
	command := os.Args[len(os.Args)-1]
	response := ""
	publicIP, err := getPublicIP()
	if err != nil {
		log.Fatal(err)
	}
	cloudFlareAPI, err := cloudflare.New(settings.CloudFlare.ApiKey, settings.CloudFlare.Email)
	if err != nil {
		log.Fatal(err)
	}
	cloudFlareZoneId, err := cloudFlareAPI.ZoneIDByName(settings.CloudFlare.Website)
	if err != nil {
		log.Fatal(err)
	}
	if command == "register" {
		response, err = registerCloudFlare(cloudFlareAPI, cloudFlareZoneId, publicIP, settings)
		if err != nil {
			log.Fatal(err)
		}
	} else if command == "unregister" {
		response, err = unregisterCloudFlare(cloudFlareAPI, cloudFlareZoneId, publicIP, settings)
		if err != nil {
			log.Fatal(err)
		}
	}
	if response == "" {
		fmt.Println("Command not found")
		os.Exit(1)
	}
	fmt.Println(response)
}
