package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"net/http"
	"os"
)

var alertComms = cli.Command{
	Name:        "alert",
	Aliases:     []string{"alert", "A", "a"},
	Usage:       "alert commands",
	Description: "alert rest calls",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "connection, cn",
			Usage: "connection ip and port",
			Value: "localhost:1712",
		},
	},
	Subcommands: []cli.Command {
		cli.Command{
			Name:        "GET",
			Aliases:     []string{"get", "G", "g"},
			Usage:       "alert get",
			Description: "alert get calls",
			Action: func(c *cli.Context) error {
				return alertGet(c)
			},
		},

		cli.Command{
			Name:        "PUT",
			Aliases:     []string{"put", "P", "p"},
			Usage:       "alert put",
			Description: "alert put calls",
			Action: func(c *cli.Context) error {
				return generateJWTToken(c)
			},
		},

		cli.Command{
			Name:        "DELETE",
			Aliases:     []string{"delete", "D", "d"},
			Usage:       "alert delete",
			Description: "alert delete calls",
			Action: func(c *cli.Context) error {
				return generateJWTToken(c)
			},
		},
	},
}

func alertGet(c *cli.Context) error {
	token := os.Getenv("AIJWTTOKEN")
	if len(token) == 0 {
		return errors.New("the JWT Token is not set at environment variable AIJWTTOKEN")
	}

	url := "https://" + c.String("connection") + ADMINPATH
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {

	}

	req.Header.Add("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("alert get call failed")
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil
}