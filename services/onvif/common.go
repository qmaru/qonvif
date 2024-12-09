package onvif

import (
	"io"

	"github.com/beevik/etree"
	goonvif "github.com/use-go/onvif"
)

func handleMethod(dev *goonvif.Device, method any) (*etree.Element, error) {
	response, err := dev.CallMethod(method)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	doc := etree.NewDocument()
	err = doc.ReadFromBytes(body)
	if err != nil {
		return nil, err
	}

	xmlRoot := doc.SelectElement("SOAP-ENV:Envelope")
	xmlBody := xmlRoot.SelectElement("SOAP-ENV:Body")

	return xmlBody, nil
}
