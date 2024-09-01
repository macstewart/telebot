package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_http "github.com/macstewart/telebot/internal/http"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var (
	serverCmd     = &cobra.Command{Use: "server", Run: start}
	port          *int
	hostname      string
	tokenCallback = make(chan string)
)

func init() {
	port = serverCmd.Flags().IntP("port", "p", 3334, "Port to run the server on")
	rootCmd.AddCommand(serverCmd)
}

func start(cmd *cobra.Command, args []string) {
	hostname = fmt.Sprintf("http://localhost:%d", *port)
	rtr := router()
	slog.Info("Server listening started", "host", hostname)
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), rtr); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}

func router() *gin.Engine {
	router := gin.Default()
	router.POST("/notify", _http.NotifyHandler())
	return router
}
