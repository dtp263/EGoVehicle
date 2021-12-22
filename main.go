package main

import (
	"github.com/sirupsen/logrus"

	"github.com/dtp263/bessie/src/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
