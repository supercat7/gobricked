package main

import (
	"flag"
	"fmt"
	"gobricked/pkg/comms"
	"gobricked/pkg/stats"
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
	stats.UpTimeInit()

	fmt.Println("Loading server configurations...")
	config, err := util.LoadServerConfig(configPath)
	if err != nil {
		panic(err)
	}
	var port string = util.GetServerPort(config)

	fmt.Println("Launching Client Server on port:", port)
	var ServerInstance *comms.Listener = comms.NewListener(port)
	var ServerChannel chan struct{} = make(chan struct{})
	ServerInstance.Start(ServerChannel, config)
}
