package iputils

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

const (
	EXTERNAL_IP_URL = "http://myexternalip.com/raw"
)

func GetPublicIp() (string, error) {
	resp, err := http.Get(EXTERNAL_IP_URL)
	if err != nil {
		return "", nil
	}

	if resp.StatusCode > 299 && resp.StatusCode < 200 {
		return "", fmt.Errorf("receieved invalid status code %v", resp.StatusCode)
	}

	buf := new(bytes.Buffer)
	read, err := buf.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}

	if read == 0 {
		return "", errors.New("no data read, possibly no IP found")
	}

	return buf.String(), nil
}
