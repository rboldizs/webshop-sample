package main

import (
	"flag"

	"git.appagile.io/consumption/consumption-go-common/rest"
	"git.appagile.io/services-paas/bb/vSweets/shop-backend-rest/info"
	"git.appagile.io/services-paas/bb/vSweets/shop-backend-rest/server/factories"

	"github.com/golang/glog"
)

var configFile = "config/config.json"

func main() {
	flag.Lookup("logtostderr").Value.Set("true")
	flag.Lookup("log_dir").Value.Set("./log")
	flag.Parse()

	glog.Infoln("Version: ", info.Version)
	glog.Infoln("Build Date: ", info.BuildDate)
	glog.Infoln("Git commit: ", info.GitCommit)

	factory := factories.GetFactory()
	factory.GetConfigurationObj().InitConf(configFile)

	restServer := rest.RestServer{
		Endpoints: []rest.Endpoint{
			factory.GetStatusEPObj(),
		},
		Unsecured: true,
		Port:      factory.GetConfigurationObj().GetServerPort(),
		APIToken:  factory.GetConfigurationObj().GetServerToken(),
	}

	glog.Info("Starting REST server.")
	if err := restServer.StartServer(); err != nil {
		glog.Fatalf("Rest server cannot start. Error: %s", err)
	}

	glog.Flush()
}
