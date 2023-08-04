package scalinglightning

import (
	"fmt"

	sl "github.com/scaling-lightning/scaling-lightning/pkg/network"
	"github.com/spf13/cobra"
)

var pubkeyNodeName string

var pubkeyCmd = &cobra.Command{
	Use:   "pubkey",
	Short: "Get the pubkey of a node",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		slnetwork, err := sl.DiscoverStartedNetwork("")
		if err != nil {
			fmt.Printf(
				"Problem with network discovery, is there a network running? Error: %v\n",
				err.Error(),
			)
			return
		}
		for _, node := range slnetwork.LightningNodes {
			if node.GetName() == pubkeyNodeName {
				pubkey, err := node.GetPubKey()
				if err != nil {
					fmt.Printf("Problem getting pubkey: %v\n", err.Error())
					return
				}
				fmt.Println(pubkey)
				return
			}
		}
		fmt.Printf(
			"Can't find node with name %v, here are the lightning nodes that are running: %v\n",
			pubkeyNodeName,
			slnetwork.LightningNodes,
		)
	},
}

func init() {
	rootCmd.AddCommand(pubkeyCmd)

	pubkeyCmd.Flags().
		StringVarP(&pubkeyNodeName, "node", "n", "", "The name of the node to get the wallet balance of")
	pubkeyCmd.MarkFlagRequired("node")
}
