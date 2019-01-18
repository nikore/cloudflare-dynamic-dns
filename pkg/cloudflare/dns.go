package cloudflare

import (
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"log"
)

type cfDNSUpdater struct {
	client   *cloudflare.API
	zoneName string
	records  []string
}

func NewDNSUpdater(apiKey string, email string) (*cfDNSUpdater, error) {
	api, err := cloudflare.New(apiKey, email)
	if err != nil {
		return nil, err
	}

	return &cfDNSUpdater{
		client: api,
	}, nil
}

func (c *cfDNSUpdater) ZoneName(zoneName string) *cfDNSUpdater {
	c.zoneName = zoneName

	return c
}

func (c *cfDNSUpdater) RecordList(records []string) *cfDNSUpdater {
	c.records = records

	return c
}

func (c *cfDNSUpdater) checkAndUpdate(currentIp string, name string, zoneID string) error {
	records, err := c.client.DNSRecords(zoneID, cloudflare.DNSRecord{Type: "A", Name: name})
	if err != nil {
		return err
	}

	if len(records) > 1 || len(records) == 0 {
		return fmt.Errorf("got the wrong number of records in our query: %v\n", records)
	}

	record := records[0]

	if record.Content != currentIp {
		record.Content = currentIp
		err := c.client.UpdateDNSRecord(zoneID, record.ID, record)
		if err != nil {
			return err
		}
		log.Printf("Updated Record %s with IP %s\n", name, currentIp)
	} else {
		log.Printf("Record %s is already using current IP\n", name)
	}

	return nil
}

func (c *cfDNSUpdater) Run(currentIp string) error {
	id, err := c.client.ZoneIDByName(c.zoneName)
	if err != nil {
		return err
	}

	err = c.checkAndUpdate(currentIp, c.zoneName, id)
	if err != nil {
		return err
	}

	for _, r := range c.records {
		name := fmt.Sprintf("%s.%s", r, c.zoneName)
		err = c.checkAndUpdate(currentIp, name, id)
		if err != nil {
			return err
		}
	}

	return nil
}
