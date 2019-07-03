package test

import (
    "testing"
    "oai-snap-in-docker/internal/pkg/common"
    "os"
    "log"
)

func TestCfg(t *testing.T) {
    newFile, err := os.Create(logPath)
	if err != nil {
        t.Error(err) // to indicate test failed
		return 
	}
	logger := log.New(newFile, "[Debug] ", log.Lshortfile)
	confStruct := common.Cfg{}
	err = confStruct.GetConf(logger, confPath)
	if err != nil {
		t.Error(err) // to indicate test failed
		return
	}
}