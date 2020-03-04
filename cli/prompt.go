package cli

import (
	"fmt"
	"os"
	"strings"

	db "pandawa/database"
	git "pandawa/git"
	shodan "pandawa/shodan"
	spider "pandawa/spider"

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
		if initcommand[0] == "exit" {
			fmt.Println("[+] Pandawa shutdown services ...")
			os.Exit(1)
			return
		}
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
	case "set-operation":
		fmt.Println("[+] Generating operation database ..")
		dbname := db.GenerateDb(initcommand[1])
		fmt.Println("[+] Success generate db " + dbname)
		return

	case "github":
		fmt.Println("[+] Collect data github")
		db.CollectData(initcommand[1])
		return

	case "shodan":
		fmt.Println("[+] Collect data shodan")
		db.CollectDataShodan(initcommand[1])
		return

	case "spider-js":
		spider.InitJs(initcommand[1])
		return

	case "spider-page":
		spider.InitPage(initcommand[1])
		return

	case "spider-github":
		git.InitialSearch(initcommand[1])
		return

	case "git-org":
		git.GetUserFromOrg(initcommand[1])
		return

	case "spider-shodan":
		shodankey := os.Getenv("shodan")
		if shodankey == "" {
			fmt.Println("[!] shodan key not found please export shodan key in .bashrc or .zshrc")
			os.Exit(1)
		}
		dbname := "pandawa-output/" + initcommand[2] + ".db"
		shodan.PreSearch(shodankey, initcommand[1], dbname)
		return

	case "exit":
		os.Exit(1)

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
		{Text: "exit", Description: "Exit."},
		{Text: "set-operation", Description: "Set operation name"},
		{Text: "github", Description: "Get All data from github with operation name"},
		{Text: "git-org", Description: "Get All users from org"},
		{Text: "shodan", Description: "Get All data from shodan with operation name"},
		{Text: "spider-js", Description: "Get js file from web "},
		{Text: "spider-page", Description: "Get page file from web "},
		{Text: "spider-github", Description: "Get data from github and commit"},
		{Text: "spider-shodan", Description: "Get data from shodna with mm3"},
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
