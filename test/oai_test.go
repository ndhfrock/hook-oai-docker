package test

import (
	"oai-snap-in-docker/internal/oai"
	"testing"
)

func TestOAI(t *testing.T) {
	// Initialize oai struct
	OaiObj := oai.Oai{}
	err := OaiObj.Init(logPath, confPath)
	if err != nil {
		t.Error("Init Oaiobj failed") // to indicate test failed
		return
	}
	oai.StartCN(OaiObj)

}
