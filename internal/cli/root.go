package cli

import (
	"context"
	"net/http"
	_ "net/http/pprof"
	"strconv"

	"github.com/spf13/cobra"
)

var GlobalOpts struct {
	Host         string
	Port         int
	Json         bool
	GenCmpl      bool
	BashCmplFile string
	PprofPort    int
	LogLevel     string
}

var ctx context.Context

// NewRootCmd creates a new root command for the cli.
func NewRootCmd() *cobra.Command {
	cobra.EnablePrefixMatching = true
	var cancel context.CancelFunc
	rootCmd := &cobra.Command{
		Use:   "rest-api",
		Short: "A REST API for a given file",
		// PersistentPreRun is called before the command's Run function.
		// It can be used to start services before the command is executed.
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Start pprof to check cpu usage, memory usage and goroutines
			if GlobalOpts.PprofPort > 0 {
				go func() {
					address := "localhost:" + strconv.Itoa(GlobalOpts.PprofPort)
					if err := http.ListenAndServe(address, nil); err != nil {
						exitWithError(err)
					}
				}()
			}
		},
		// Run is called when the command is executed.
		// If the --gen-bash-cmpl flag is set, the command will generate a
		// bash completion file.
		Run: func(cmd *cobra.Command, args []string) {
			if GlobalOpts.GenCmpl {
				cmd.GenBashCompletionFile(GlobalOpts.BashCmplFile)
			} else {
				cmd.HelpFunc()(cmd, args)
			}
		},
		// PersistentPostRun is called after the command's Run function.
		// It can be used to stop services started in PersistentPreRun.
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if cancel != nil {
				cancel()
			}
		},
	}

	// Add flags to the root command
	rootCmd.PersistentFlags().StringVarP(&GlobalOpts.Host, "host", "u", "127.0.0.1", "host")
	rootCmd.PersistentFlags().IntVarP(&GlobalOpts.Port, "port", "p", 1323, "port")
	rootCmd.PersistentFlags().BoolVarP(&GlobalOpts.Json, "json", "j", false, "use json format to output format")
	rootCmd.PersistentFlags().IntVarP(&GlobalOpts.PprofPort, "pprof-port", "r", 0, "pprof port")
	rootCmd.PersistentFlags().StringVarP(&GlobalOpts.LogLevel, "log-level", "l", "ERROR", "log level (ERROR, WARNING, INFO, DEBUG)")

	// Add the server command to the root command
	serverCmd := newServerCmd()
	rootCmd.AddCommand(serverCmd)

	return rootCmd
}
