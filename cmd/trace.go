package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "IP tracing",
	Long:  `IP tracing.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}
		} else {
			c := color.New(color.FgRed)
			c.Println("Provide the IP to trace.")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

type Ip struct {
	IP       string `json::"ip"`
	City     string `json::"city"`
	Region   string `json::"region"`
	Country  string `json::"country"`
	Loc      string `json::"loc"`
	Timezone string `json::"timezone"`
	Postal   string `json::"postal"`
}

func showData(ip string) {
	url := "https://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)

	data := Ip{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		c := color.New(color.FgRed)
		c.Println("Unable to unmarshal response")
	}

	c := color.New(color.FgGreen).Add(color.Underline).Add(color.Bold)
	c.Println("DATA FOUND")

	fmt.Printf("IP : %s\nCITY : %s\nREGION : %s\nCOUNTRY : %s\nLOC : %s\nTIMEZONE : %s\nPOSTAL : %s\n", data.IP, data.City, data.Region, data.Country, data.Loc, data.Timezone, data.Postal)
}

func getData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		c := color.New(color.FgRed)
		c.Println("Unable to get response")
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c := color.New(color.FgRed)
		c.Println("Unable to get response")
	}

	return responseByte
}
