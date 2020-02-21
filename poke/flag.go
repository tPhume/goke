package poke

import "github.com/urfave/cli/v2"
// flags for poke cli
func defineGetPokeFlags() []cli.Flag {
	pokemonFlag := &cli.StringFlag{
		Name: "pokemon",
		Aliases: []string{"p"},
		Usage: "retrieve `POKEMON` information - must be either id or name",
		Required: true,
	}

	outputFlag := &cli.PathFlag{
		Name: "output",
		Aliases: []string{"o"},
		Usage: "output result to `FILE` - directory must already exist",
	}

	return defineFlags(pokemonFlag, outputFlag)
}

// helper function to create flags
func defineFlags(flagArgs ...cli.Flag) []cli.Flag {
	flags := make([]cli.Flag, len(flagArgs))
	for index, f := range flagArgs {
		flags[index] = f
	}

	return flags
}
