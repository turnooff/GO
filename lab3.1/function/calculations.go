package function

import (
	log "github.com/sirupsen/logrus"
)

func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	var mylty int64 = n
	n--
	for n > 0 {
		mylty *= n
		n--
	}
	return mylty
}

func Calculate(n int64, flag bool) int64 {
	if flag {
		log.Info("Start calculations...")
		log.Infof("Calculate <%d>!", n)
		log.Info("Calculations complete!")
	}
	return factorial(n)
}
