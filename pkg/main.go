package main

import (
	"flag"
	"fmt"
	"gobricked/pkg/comms"
	"gobricked/pkg/util"
	"os"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "Server Yaml config path (required)")
	flag.Parse()

	if configPath == "" {
		fmt.Println("Error: the config file path is required")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("Initalizing Server statistics and data...")
	util.UpTimeInit()

	fmt.Println("Loading server configurations...")
	config, err := util.LoadServerConfig(configPath)
	if err != nil {
		panic(err)
	}
	var port string = config.Teamserver.Port

	fmt.Println("Launching Operator Server on port:", port)
	var ServerInstance *comms.OperatorServer = comms.NewOperatorServer(port, config)
	var ServerChannel chan struct{} = make(chan struct{})
	ServerInstance.Start(ServerChannel)
}
