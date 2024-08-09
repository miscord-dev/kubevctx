package hook

import (
	_ "embed"
)

//go:embed zsh.sh
var zsh string

func ZSH() string {
	return zsh
}
