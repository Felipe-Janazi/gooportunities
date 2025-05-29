package main

// Para fazer a instalação de todos pacotes que estão sendo utilizados e deltar os que não estão sendo podemos usar o:
// go mod tidy
import (
	"github.com/Felipe-Janazi/gopportunities/config"
	"github.com/Felipe-Janazi/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
