package app

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLog()  {
	G.Log = logrus.New()
	G.Log.SetLevel(logrus.TraceLevel)
	G.Log.SetOutput(os.Stdout)
	G.Log.SetReportCaller(true)
	G.Log.SetFormatter(&logrus.JSONFormatter{})
}