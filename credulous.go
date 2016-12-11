package main

import (
	"fmt"
	"os"
        "github.com/urfave/cli"
)


func main() {
  app := cli.NewApp()
  app.Name = "credulous"
  app.Usage = "Secure AWS Credential Management"
  app.Version = "0.2.2-dnabic"

  app.Run(os.Args)

  fmt.Println("All fine")
}
