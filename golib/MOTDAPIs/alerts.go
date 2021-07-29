package main

import (
	"bytes"
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
			Name:  "CONNECTION, cn",
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
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "BODY, b",
					Usage: "HTTP body",
				},
			},
			Action: func(c *cli.Context) error {
				return alertPut(c)
			},
		},

		cli.Command{
			Name:        "DELETE",
			Aliases:     []string{"delete", "D", "d"},
			Usage:       "alert delete",
			Description: "alert delete calls",
			Action: func(c *cli.Context) error {
				return alertDelete(c)
			},
		},
	},
}

func alertGet(c *cli.Context) error {
	return httpCall(c, http.MethodGet, "")
}

func alertPut(c *cli.Context) error {

	if body := c.String("BODY"); len(body) != 0 {
		return httpCall(c, http.MethodPut, body)
	} else {
		return errors.New("the BODY flag is not set probably")
	}
}

func alertDelete(c *cli.Context) error {
	return httpCall(c, http.MethodDelete, "")
}

func httpCall(c *cli.Context, action, body string) error {
	token := os.Getenv("AIJWTTOKEN")
	if len(token) == 0 {
		return errors.New("the JWT Token is not set at environment variable AIJWTTOKEN")
	}

	serverRoot := c.Parent().String("CONNECTION")
	if len(serverRoot) == 0 {
		return errors.New("the CONNECTION flag is not set probably")
	}
	url := "https://" + serverRoot + ADMINPATH + ALERTPATH
	req, err := http.NewRequest(action, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return errors.New("creating HTTP Request failed")
	}

	req.Header.Add("Accept","application/json")
	req.Header.Add("Authorization", "Bearer " + token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("alert get call failed")
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(respBody))

	return nil
}

