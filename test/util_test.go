package test

import (
	"log"
	"oai-snap-in-docker/internal/pkg/util"
	"os"
	"testing"
	"time"
)

func TestUtil(t *testing.T) {
	newFile, err := os.Create(logPath)
	if err != nil {
		t.Error(err) // to indicate test failed
		return
	}
	logger := log.New(newFile, "[Debug] "+time.Now().Format("2006-01-02 15:04:05")+" ", log.Lshortfile)

	ret, err := util.GetIPFromDomain(logger, "google.com")
	if ret == "" || err != nil {
		t.Error("GetIPFromDomain failed") // to indicate test failed
		return
	}
	t.Log("google.com IP is ", ret)

	ret = util.GetOutboundIP()
	t.Log("Outbound IP is ", ret)
	if ret != "192.168.12.78" {
		t.Error("GetOutboundIP failed") // to indicate test failed
		return
	}

	ret, err = util.GetInterfaceByIP("192.168.12.78")
	if ret != "enx00249b154d0a" || err != nil {
		t.Error("GetInterfaceByIP failed") // to indicate test failed
		return
	}
	t.Log("Outbound interface is ", ret)

	ret, err = util.GetInterfaceIP(logger, "lo")
	if ret != "127.0.0.1" || err != nil {
		t.Error("GetInterfaceIP failed") // to indicate test failed
		return
	}
	t.Log("localhost is ", ret)
}
