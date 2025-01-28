package service

import (
	"log"
	"strconv"
	"time"

	"github.com/mindmain/go-mongo-sidecar/types"
)

func (s *sidecarService) initDuration() {

	sleepSeconds, _ := strconv.ParseInt(types.SIDECAR_TIME_SLEEP, 10, 64)

	if sleepSeconds <= 0 {
		log.Println("[INFO] sleep duration is 0, set default value 5 seconds")
		sleepSeconds = 5
	}

	waitSeconds, _ := strconv.ParseInt(types.SIDECAR_TIME_TO_WAIT, 10, 64)
	if waitSeconds <= 0 {
		log.Println("[INFO] wait duration is 0, set default value 10 seconds")
		waitSeconds = 10
	}

	sleep := time.Second * time.Duration(sleepSeconds)
	wait := time.Second * time.Duration(waitSeconds)

	if sleep <= 0 {
		log.Println("[INFO] sleep duration is 0, set default value 5 seconds")
		sleep = time.Second * 5
	}

	if wait <= 0 {
		log.Println("[INFO] wait duration is 0, set default value 10 seconds")
		wait = time.Second * 10
	}

	s.sleepDuration = sleep
	s.waitDuration = wait

}

func (s *sidecarService) wait() {
	time.Sleep(s.waitDuration)
}

func (s *sidecarService) sleep() {
	time.Sleep(s.sleepDuration)
}
