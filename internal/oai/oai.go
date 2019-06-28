package oai

import (
	"log"
	"oai-snap-in-docker/internal/pkg/common"
	"os"
)

// Oai stores the log and conf
type Oai struct {
	logFile *os.File    // File for log to write something
	Logger  *log.Logger // Collect log
	Conf    *common.Cfg // config files

}

// Init the Oai with log and conf
func (me *Oai) Init(logPath string, confPath string) error {
	newFile, err := os.Create(logPath)
	if err != nil {
		return err
	}
	me.logFile = newFile
	me.Logger = log.New(me.logFile, "[Debug] ", log.Lshortfile)
	me.Conf = new(common.Cfg)
	err = me.Conf.GetConf(me.Logger, confPath)
	if err != nil {
		return err
	}
	return nil
}

// Clean will Close the logFile and clean up Obj
func (me *Oai) Clean() {
	me.logFile.Close()
}
