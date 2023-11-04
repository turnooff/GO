package main

import (
	"flag"
	"os"
	"strconv"

	fun "main/function"

	log "github.com/sirupsen/logrus"
)

func main() {
	helpFlag := flag.Bool("b", false, "mast be bool: true or false!")
	flag.Parse()

	helpN, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Error(err)
		return
	}

	argLen := len(os.Args[1:])
	if argLen != 2 {
		log.Error("Only one argument is needed")
		return
	}

	n := int64(helpN)

	log.Info(fun.Calculate(n, *helpFlag))
}
