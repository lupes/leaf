package cmd

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/urfave/negroni"

	"leaf/server/routes"
)

func serverRun(cmd *cobra.Command, args []string) {
	neg := negroni.New(negroni.NewRecovery(), negroni.NewLogger())

	router := httprouter.New()
	routes.InitRoutes(router)
	neg.UseHandler(router)

	logrus.Info("start server")
	neg.Run(":8000")
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  `run server`,
	Run:   serverRun,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
