package main

import (
	"github.com/lichtwellenreiter/reaktor/config"
	"github.com/spf13/viper"
	"log"
)

func main() {
	printHeader()
	log.Println("Starting Reaktor")
	readConfiguration()
}

func readConfiguration() {
	viper.SetConfigName("reaktor")
	viper.AddConfigPath(".")

	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		log.Fatalf("Unable to unmarshal config, %s", err)
	}

	log.Printf("Name is %s\n", configuration.Name)
	log.Printf("port is %d\n", configuration.Port)
}

func printHeader() {

	log.Printf(`
__________               __      __                
\______   \ ____ _____  |  | ___/  |_  ___________ 
 |       _// __ \\__  \ |  |/ /\   __\/  _ \_  __ \
 |    |   \  ___/ / __ \|    <  |  | (  <_> )  | \/
 |____|_  /\___  >____  /__|_ \ |__|  \____/|__|   
        \/     \/     \/     \/
Version 0.1

`)

}

/*
	go func() {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			log.Println("Config file changed, reread configuration")
			if err := viper.Unmarshal(&configuration); err != nil {
				log.Fatalf("Unable to unmarshal config, %s", err)
			}
		})
	}()
*/
