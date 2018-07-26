package main

import (
	"log"
	"os"
	"os/exec"
)

var cmdChain = []*exec.Cmd{
	// synonymsを動かすには
	// Big Huge Thesaurus(https://words.bighugelabs.com/)のAPIキーを
	// 環境変数にセットする必要があります
	// exec.Command("lib/synonyms"),
	exec.Command("lib/sprinkle"),
	exec.Command("lib/coolify"),
	exec.Command("lib/domainify"),
	exec.Command("lib/avaliable"),
}

func main() {
	cmdChain[0].Stdin = os.Stdin
	cmdChain[len(cmdChain)-1].Stdout = os.Stdout

	for i := 0; i < len(cmdChain)-1; i++ {
		thisCmd := cmdChain[i]
		nextCmd := cmdChain[i+1]
		stdout, err := thisCmd.StdoutPipe()
		if err != nil {
			log.Panicln(err)
		}
		nextCmd.Stdin = stdout
	}

	for _, cmd := range cmdChain {
		if err := cmd.Start(); err != nil {
			log.Panicln(err)
		} else {
			defer cmd.Process.Kill()
		}
	}

	for _, cmd := range cmdChain {
		if err := cmd.Wait(); err != nil {
			log.Panicln(err)
		}
	}
}
