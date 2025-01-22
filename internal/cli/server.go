package cli

import (
	"gosolve/internal/index"
	"gosolve/internal/server"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// newServerCmd creates a new command for the server.
func newServerCmd() *cobra.Command {
	// Create a new command for the server.
	serverCmd := &cobra.Command{
		Use:   cmdServer,
		Short: "Starts the server",
		// Run is called when the command is executed.
		Run: func(cmd *cobra.Command, args []string) {
			// Init File manager.
			// The file manager is responsible for loading the index and
			// updating it every hour.
			imFile, err := index.NewIndexManager()
			if err != nil {
				log.Error("File manager Loaded with errors (missing file?)")
			} else {
				log.Info("File manager loaded")
			}

			// Init Webserver with configurations.
			// The webserver is responsible for starting the server and
			// registering handlers for it. The server is started with the
			// provided host and port.
			webserver := server.NewWebServer(imFile, &GlobalOpts.LogLevel)
			// Start the server.
			webserver.Start(&GlobalOpts.Host, &GlobalOpts.Port)
		},
	}

	return serverCmd
}
