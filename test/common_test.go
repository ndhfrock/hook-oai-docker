package test

import (
	"log"
	"oai-snap-in-docker/internal/pkg/common"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestCfg(t *testing.T) {
	newFile, err := os.Create(logPath)
	if err != nil {
		t.Error(err) // to indicate test failed
		return
	}
	logger := log.New(newFile, "[Debug] "+time.Now().Format("2006-01-02 15:04:05")+" ", log.Lshortfile)
	confStruct := common.Cfg{}
	err = confStruct.GetConf(logger, confPath)
	if err != nil {
		t.Error(err) // to indicate test failed
		return
	}
	v := reflect.ValueOf(confStruct)
	vn := reflect.ValueOf(confStruct)
	for i := 0; i < v.NumField(); i++ {
		logger.Println(vn.Type().Field(i).Name, " is ", v.Field(i).Interface())
	}
}
