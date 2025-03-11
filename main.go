package main

import (
	"github.com/fabioaacarneiro/gopportunities/config"
	"github.com/fabioaacarneiro/gopportunities/router"
)

var (
    logger config.Logger
)

func main() {
    logger = *config.GetLogger("main")
    // initialize configs
    err := config.Init()
    if err != nil {
        logger.Errorf("Config initialize error: %v", err)
        return
    }

    // initialize router
    router.Initialize()
}
