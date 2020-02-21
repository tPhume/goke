package poke

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"net/http"
)

func RunApp(args []string) error {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "get",
				Usage: "Retrieve details of Pokemon from PokéAPI.",
				Action: getPokemon,
				Flags: defineGetPokeFlags(),
			},
		},
	}

	return app.Run(args)
}

func getPokemon(c *cli.Context) error {
	// calling the api endpoint
	resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", c.Value("pokemon")))
	if err != nil {
		return err
	} else if resp.StatusCode == http.StatusNotFound {
		return errors.New("invalid pokemon name or id")
	}

	// reading the response body
	resBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// making it readable
	prettyBytes := bytes.Buffer{}
	if err =json.Indent(&prettyBytes, resBytes, "", "  "); err != nil {
		return err
	}

	// display the JSON body
	fmt.Printf("%s\n", prettyBytes.String())

	// save file if output flag is given
	output := c.Value("output")
	if output != "" {
		if err = ioutil.WriteFile(output.(string), prettyBytes.Bytes(), 0644); err != nil {
			return err
		}

		fmt.Printf("[JSON written to %s]\n", output)
	}

	return nil
}