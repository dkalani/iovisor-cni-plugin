package main

import (
       "fmt"
       "github.com/appc/cni/pkg/skel"
       "github.com/dkalani/iomodules/policy/client"
 )

func cmdAdd(args *skel.CmdArgs) error {
  client := client.NewClient()
  result, err := client.CreateEndpoint(args)

  if (err != nil) {
     return fmt.Errorf("Create endpoint failed: %s", err)
  }
  Error.Print("received add event %s", "policy-plugin")

  return result.Print()
}

func cmdDel(args *skel.CmdArgs) error {
     return nil
}

func main() {
     skel.PluginMain(cmdAdd, cmdDel)
}
