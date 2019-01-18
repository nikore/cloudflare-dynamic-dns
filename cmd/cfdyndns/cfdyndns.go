package main

import (
	"github.com/nikore/cloudflare-dynamic-dns/pkg/cloudflare"
	"github.com/nikore/cloudflare-dynamic-dns/pkg/iputils"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

var (
	Version = "undefined"
)

func main() {
	kingpin.Version(Version)
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate)
	apiKey := kingpin.Flag("api-key", "Cloudflare API Key").Required().String()
	email := kingpin.Flag("email", "Cloudflare email address").Required().String()
	zoneName := kingpin.Flag("zone", "Name of zone to update").Required().String()
	recordList := kingpin.Flag("records", "List of records to update").Strings()
	kingpin.Parse()

	log.Println("Getting Public IP Address")
	currentIP, err := iputils.GetPublicIp()
	if err != nil {
		log.Fatal("Error getting public IP", err)
	}

	log.Printf("Current Public IP is %s \n", currentIP)

	updater, err := cloudflare.NewDNSUpdater(*apiKey, *email)
	if err != nil {
		log.Fatal("Error creating cloudflare dns updater", err)
	}

	log.Println("Starting to check dns for updates")

	err = updater.ZoneName(*zoneName).RecordList(*recordList).Run(currentIP)
	if err != nil {
		log.Fatal("Error updating dns records ", err)
	}

	log.Println("All Done!")
}
