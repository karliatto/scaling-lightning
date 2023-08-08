package scalinglightning

import (
	"fmt"

	sl "github.com/scaling-lightning/scaling-lightning/pkg/network"
	"github.com/scaling-lightning/scaling-lightning/pkg/types"
	"github.com/spf13/cobra"
)

var openchannelCmd = &cobra.Command{
	Use:   "openchannel",
	Short: "Open a channel between two nodes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		openchannelFromName := cmd.Flag("from").Value.String()
		openchannelToName := cmd.Flag("to").Value.String()
		openchannelLocalAmt, err := cmd.Flags().GetUint64("localamt")
		if err != nil {
			fmt.Println("Amount must be a valid number")
			return
		}

		slnetwork, err := sl.DiscoverStartedNetwork("")
		if err != nil {
			fmt.Printf(
				"Problem with network discovery, is there a network running? Error: %v\n",
				err.Error(),
			)
			return
		}
		var openchannelFromNode sl.LightningNode
		var openchannelToNode sl.LightningNode
		for _, node := range slnetwork.LightningNodes {
			if node.GetName() == openchannelFromName {
				openchannelFromNode = node
				continue
			}
			if node.GetName() == openchannelToName {
				openchannelToNode = node
			}
		}
		allNames := []string{}
		for _, node := range slnetwork.LightningNodes {
			allNames = append(allNames, node.GetName())
		}
		if openchannelFromNode.Name == "" {
			fmt.Printf(
				"Can't find node with name %v, here are the lightnign nodes that are running: %v\n",
				openchannelFromName,
				allNames,
			)
		}
		if openchannelToNode.Name == "" {
			fmt.Printf(
				"Can't find node with name %v, here are the lightning nodes that are running: %v\n",
				openchannelToName,
				allNames,
			)
		}

		chanPoint, err := openchannelFromNode.OpenChannel(
			&openchannelToNode,
			types.NewAmountSats(openchannelLocalAmt),
		)
		if err != nil {
			fmt.Printf("Problem opening channel: %v\n", err.Error())
			return
		}

		fmt.Printf(
			"Open channel command received. Txid: %v OutputIndex: %d",
			chanPoint.FundingTxid,
			chanPoint.OutputIndex,
		)
	},
}

func init() {
	rootCmd.AddCommand(openchannelCmd)

	openchannelCmd.Flags().
		StringP("from", "f", "", "Name of node to open channel from")
	openchannelCmd.MarkFlagRequired("from")

	openchannelCmd.Flags().
		StringP("to", "t", "", "Name of node to open channel to")
	openchannelCmd.MarkFlagRequired("to")

	openchannelCmd.Flags().
		Uint64P("amount", "a", 0, "Amount of satoshis to put into channel from the opening side")
	openchannelCmd.MarkFlagRequired("amount")

}
