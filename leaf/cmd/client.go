package cmd

import (
	"net/http"

	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		client := http.Client{}
		client.Get("http://localhost:8080/v1/test?name=lupe")
		//go func() {
		//	select {
		//	case <-ctx.Done():
		//		fmt.Printf("到达时间\n")
		//		cancel()
		//	}
		//}()
		//request(ctx)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
