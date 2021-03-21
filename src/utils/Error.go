package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

//If err != nil close the program.
func CheckError(err error) {
	if err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}
}
