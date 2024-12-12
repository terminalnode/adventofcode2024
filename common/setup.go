package common

import (
	"context"
	"github.com/terminalnode/adventofcode2024/common/util"
	"github.com/terminalnode/adventofcode2024/common/web"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Setup(
	day int,
	part1 util.Solution,
	part2 util.Solution,
) {
	httpServer := web.CreateHttpServer(day, part1, part2)

	// Open a signal channel, listening for SIGTERM and SIGINT
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	log.Printf("Received signal %s, shutting down...", <-signalChan)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
