package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	method string
)

// allow methods
var methods = map[string]bool{"GET": true, "POST": true}
var headers = []string{"Content-Type:application/json"}
var data string

var rootCmd = &cobra.Command{
	Use:   "curl",
	Short: "like curl",
	Long:  "curl developed by golang",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		url := args[0]
		// check methods
		if !methods[strings.ToUpper(method)] {
			return errors.New("method not allow")
		}

		client := &http.Client{}
		req, err := http.NewRequest(method, url, strings.NewReader(data))
		if err != nil {
			return
		}

		// header 处理
		for _, v := range headers {
			arr := strings.Split(v, ":")
			if len(arr) == 2 {
				req.Header.Add(arr[0], arr[1])
			}
		}

		res, err := client.Do(req)

		if err != nil {
			return
		}

		fmt.Fprintln(os.Stdout, "[request]")
		fmt.Fprintln(os.Stdout, req.URL)
		printHeader(req.Header)
		fmt.Println()

		fmt.Fprintln(os.Stdout, "[response]")
		printHeader(res.Header)
		fmt.Println()
		printBody(res)

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&method, "method", "X", "GET", "HTTP method type")
	rootCmd.PersistentFlags().StringVarP(&data, "data", "d", "", "POST method body data")
	rootCmd.PersistentFlags().StringSliceVarP(&headers, "header", "H", headers, "headers")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func printHeader(header http.Header) {
	for key, value := range header {
		fmt.Printf("%v:%v\n", key, strings.Join(value, ";"))
	}
}

func printBody(res *http.Response) {
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	fmt.Fprintln(os.Stdout, string(bodyBytes))
}
