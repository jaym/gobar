package main

import "github.com/Sirupsen/logrus"
import model "github.com/jaym/gobar/model"

func main() {
	bar := model.BarStruct{
		Baz: "thing",
	}
	logrus.Infof("Ima bar: %+v", bar)
}
