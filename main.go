package main

import (
	"fmt"
	"os"

	"github.com/cheyang/numa-utils/numa"
	nvidia "github.com/cheyang/nv-tools/helper"
	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use:          os.Args[0],
	Short:        "Run gpu display",
	SilenceUsage: false,
	RunE: func(cmd *cobra.Command, args []string) error {
		// log.SetOutput(os.Stderr)

		if flag, err := cmd.Flags().GetBool("show"); err != nil {
			return err
		} else if !flag {
			return nil
		}

		if nv, err := nvidia.NewNvHelper(); err != nil {
			return err
		} else {
			devs := nv.Devices
			for i, dev := range devs {
				fmt.Printf("dev %d: %+v", i, dev)
				node := *(dev.NVMLDevice.CPUAffinity)
				cpus, err := numa.CPUsOfNode(int(node))
				if err != nil {
					return err
				}
				fmt.Printf("node %d cpus: %+v\n", node, cpus)
				all, free := numa.MemoryOfNode(int(node))
				fmt.Printf("node %d size: %d MB\n", node, numa.MemInMB(all))
				fmt.Printf("node %d free: %d MB\n", node, numa.MemInMB(free))
			}
		}

		return nil
	},
}

func main() {
	if err := mainCmd.Execute(); err != nil {
		fmt.Printf("Err is %v", err)
		os.Exit(-1)
	}
}

func init() {
	mainCmd.Flags().BoolP("show", "S", true, "Display GPU info")
}
