package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	KEYFILESUFFIX = "/runtime/Tomcat-7/webapps/ROOT/WEB-INF/classes/com/adaptiveplanning/system/ticket/AuthKey.parameters.properties"
)

func main() {
	app := cli.NewApp()

	app.Name = os.Args[0]
	app.Usage = "A tool to generate AI-JWT token using Planning security key (AuthKey.parameters.properties). Please refer the detailed help manual of each sub-command"
	app.UsageText = "For example, create a JWT token like ./jwttoken create"
	app.Version = "1.0.0" // after local debugging, it should work for others.
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose, vb",
			Usage: "log verbose messages",
		},
	}

	app.Commands = []cli.Command{
		jtwcomms,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


var jtwcomms = cli.Command{
	Name:    "jwt-creater",
	Aliases: []string{"create", "C", "c"},
	Usage:   "create ",
	Description: "Create a JWT token ",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "JWT_CONTENT_TYPE, ct",
			Usage: "JWT_CONTENT_TYPE",
			Value: "text/plain",
		},
		cli.StringFlag{
			Name:  "JWT_ALGORITHM, alg",
			Usage: "JWT_ALGORITHM",
			Value: "HS256",
		},
		cli.StringFlag{
			Name:  "JWT_LOGIN, lgin",
			Usage: "JWT_LOGIN",
			Value: "systemapiuser",
		},
		cli.StringFlag{
			Name:  "JWT_ISSUER, iss",
			Usage: "JWT_ISSUER",
			Value: "ap_build_install",
		},

		cli.StringFlag{
			Name:  "JWT_AUDIENCE, aud",
			Usage: "JWT_AUDIENCE",
			Value: "EnforcerService",
		},

		cli.StringFlag{
			Name:  "JWT_ID, jid",
			Usage: "JWT_ID",
			Value: "854FC13E-1AC5-4535-874D-A9E3B9E9219F",
		},

		cli.Uint64Flag{
			Name:  "JWT_DURATION, during",
			Usage: "JWT_DURATION",
			Value: 120000,
		},

		cli.StringFlag{
			Name:  "JWT_MULTI_USE, mu",
			Usage: "JWT_MULTI_USE",
			Value: "0",
		},

		cli.StringFlag{
			Name:  "JWT_IS_AUTH, aut",
			Usage: "JWT_IS_AUTH",
			Value: "1",
		},
		cli.BoolFlag{
			Name:   "verbose, v",
			Hidden: true,
			Usage:  "details, details, I need the details!",
		},
	},

	Before: checkDeps,
	Action: func(c *cli.Context) error {
		return generateJWTToken(c)
	},
}

var authkeyfile, jwtkey string

func checkDeps(c *cli.Context) error {
	planningPath := os.Getenv("PLANHOME")

	if len(planningPath) == 0 {
		authkeyfile = "~/git/planning" + KEYFILESUFFIX
	} else {
		authkeyfile = planningPath + KEYFILESUFFIX
	}

	if _, err := os.Stat(authkeyfile); err != nil {
		return errors.New(fmt.Sprintf("Environment variable PLANHOME: %s .AuthKey.parameters.properties doesn't exist at %s", planningPath, authkeyfile))
	}

	file, err := os.Open(authkeyfile)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		msg := scanner.Text()

		if strings.HasPrefix(msg,"key=") {
			tks := strings.Split(msg, "=")
			if len(tks) < 2 {
				return errors.New("AuthKey.parameters.properties doesn't have the key value")
			}
			jwtkey = tks[1]
		}
	}

	_, err = exec.LookPath("openssl")
	if err != nil {
		return errors.New("cannot find openssl")
	}
	if c.Bool("verbose") {
		log.Println("JWT Key: ", jwtkey)
	}

	return nil
}

func generateJWTToken(c *cli.Context) error {
	if len(jwtkey) == 0 {
		return errors.New(fmt.Sprintf("Empty JWT Key"))
	}

	issuedAt := time.Now()

	//header := &jwtHeader{c.String("JWT_CONTENT_TYPE"), c.String("JWT_ALGORITHM"), issuedAt.Unix(), expiration.Unix()}
	//payload := &jwtPayload{c.String("JWT_LOGIN"), c.String("JWT_ISSUER"), c.String("JWT_AUDIENCE"), c.String("JWT_ID"),
	//		c.Int("JWT_MULTI_USE"), c.Int("JWT_IS_AUTH"), issuedAt.Unix(), expiration.Unix(), expiration.Unix()}
	//header := &jwtHeader{c.String("JWT_CONTENT_TYPE"), c.String("JWT_ALGORITHM")}
	//payload := &jwtPayload{c.String("JWT_LOGIN"), c.String("JWT_ISSUER"), c.String("JWT_AUDIENCE"),
	//	c.String("JWT_MULTI_USE"), c.String("JWT_IS_AUTH"), 1626456263, 1626456263 + 12000000, 1626456263, c.String("JWT_ID")}


	header := &jwtHeader{c.String("JWT_CONTENT_TYPE"), c.String("JWT_ALGORITHM")}
	payload := &jwtPayload{c.String("JWT_LOGIN"), c.String("JWT_ISSUER"), c.String("JWT_AUDIENCE"),
		c.String("JWT_MULTI_USE"), c.String("JWT_IS_AUTH"), issuedAt.Unix(),
		issuedAt.Unix() + c.Int64("during"), issuedAt.Unix(), c.String("JWT_ID")}

	hjson, err := json.Marshal(header)
	if err != nil {
		log.Fatal("error:", err)
		return err
	}

	pljson, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("error:", err)
		return err
	}

	b64 := base64.RawStdEncoding.EncodeToString([]byte(hjson)) + "." + base64.RawStdEncoding.EncodeToString([]byte(pljson))

	sslCmd := exec.Command("openssl", "dgst",  "-binary",  "-sha256",  "-hmac", jwtkey)
	sslCmd.Stdin = strings.NewReader(b64)
	var digest bytes.Buffer
	sslCmd.Stdout = &digest
	if err := sslCmd.Run(); err != nil {
		log.Fatal(err)
		return err
	}

	b64 = b64 + "." + base64.RawStdEncoding.EncodeToString(digest.Bytes())

	b64 = strings.ReplaceAll(b64, "/", "_")
	b64 = strings.ReplaceAll(b64, "+", "-")
	b64 = strings.ReplaceAll(b64, "=", "")
	b64 = strings.ReplaceAll(b64, "\n", "")

	if c.Bool("verbose") {
		log.Println("JWT Header:", *header)
		log.Println("JWT Payload:", *payload)
		log.Println("Header JSON:", string(hjson))
		log.Println("Payload JSON:",string(pljson))
		log.Println("Base64 header & Payload:",b64)
		log.Println("Token:", b64)
	}

	log.Println("JWT Token:\n", b64)

	return nil
}



type jwtHeader struct {
	Cty  string `json:"cty"`
	Alg  string `json:"alg"`
}

type jwtPayload struct {
	Login_id string `json:"login_id"`
	Iss string `json:"iss"`
	Aud string `json:"aud"`
	Multi_use string  `json:"multi_use"`
	Is_auth	string  `json:"is_auth"`
	Iat	int64 `json:"iat"`
	Exp int64 `json:"exp"`
	Nbf	int64 `json:"nbf"`
	Jti	string	 `json:"jti"`
}
