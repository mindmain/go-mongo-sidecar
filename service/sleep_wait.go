package service

import (
	"log"
	"time"

	"github.com/mindmain/go-mongo-sidecar/types"
)

func (s *sidecarService) initDuration() {

	sleep := time.Second * time.Duration(types.SIDECAR_TIME_SLEEP.Int64())
	wait := time.Second * time.Duration(types.SIDECAR_TIME_TO_WAIT.Int64())

	if sleep <= 0 {
		log.Println("[INFO] sleep duration is 0, set default value 5 seconds")
		sleep = time.Second * 5
	}

	if wait <= 0 {
		log.Println("[INFO] wait duration is 0, set default value 10 seconds")
		wait = time.Second * 10
	}

}

func (s *sidecarService) wait() {
	time.Sleep(s.waitDuration)
}

func (s *sidecarService) sleep() {
	time.Sleep(s.sleepDuration)
}
