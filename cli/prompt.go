package cli

import (
	"fmt"
	"strings"

	db "pandawa/database"

	prompt "github.com/c-bata/go-prompt"
)

var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

func executor(in string) {
	if in == "" {
		LivePrefixState.IsEnable = false
		LivePrefixState.LivePrefix = in
		return
	}

	initcommand := strings.Split(in, " ")

	if len(initcommand) == 1 {
		fmt.Println("[!] please set the operation name")
		fmt.Println("[!] example >>> github pandawa")

		return
	}
	if initcommand[1] == "" {
		fmt.Println("[!] please set the operation name")
		fmt.Println("[!] example >>> github pandawa")

		return
	}
	switch initcommand[0] {
	case "github":
		fmt.Println("[+] Collect data github")
		db.CollectData(initcommand[1])
		return
	default:
		fmt.Println("[!] Not in services")
		return

	}
	LivePrefixState.LivePrefix = initcommand[1] + "ops > "
	LivePrefixState.IsEnable = false
	return
}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "github", Description: "Get All data from github with operation name"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func changeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

func Cli() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle("live-prefix-example"),
	)
	p.Run()
}
