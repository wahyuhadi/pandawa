package cli

import (
	"fmt"

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

	switch in {
	case "github":
		fmt.Println("[+] Collect data github")
		return
	default:
		fmt.Println("[!] Not in services")
		return

	}
	LivePrefixState.LivePrefix = in + "> "
	LivePrefixState.IsEnable = true
}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "github", Description: "Get All data from github"},
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
