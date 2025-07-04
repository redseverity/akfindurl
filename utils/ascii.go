package utils

import (
	"fmt"
	"strings"
)

func ToolName() {
	fmt.Println(strings.Repeat(BoldText+RedText+"="+DefaultText, 73))
	fmt.Print("\n")
	fmt.Println(RedText + ` ▄▄▄       ██ ▄█▀  █████▒██▓ ███▄    █ ▓█████▄  █    ██  ██▀███   ██▓    
▒████▄     ██▄█▒ ▓██   ▒▓██▒ ██ ▀█   █ ▒██▀ ██▌ ██  ▓██▒▓██ ▒ ██▒▓██▒    
▒██  ▀█▄  ▓███▄░ ▒████ ░▒██▒▓██  ▀█ ██▒░██   █▌▓██  ▒██░▓██ ░▄█ ▒▒██░    
░██▄▄▄▄██ ▓██ █▄ ░▓█▒  ░░██░▓██▒  ▐▌██▒░▓█▄   ▌▓▓█  ░██░▒██▀▀█▄  ▒██░    
 ▓█   ▓██▒▒██▒ █▄░▒█░   ░██░▒██░   ▓██░░▒████▓ ▒▒█████▓ ░██▓ ▒██▒░██████▒
 ▒▒   ▓▒█░▒ ▒▒ ▓▒ ▒ ░   ░▓  ░ ▒░   ▒ ▒  ▒▒▓  ▒ ░▒▓▒ ▒ ▒ ░ ▒▓ ░▒▓░░ ▒░▓  ░
  ▒   ▒▒ ░░ ░▒ ▒░ ░      ▒ ░░ ░░   ░ ▒░ ░ ▒  ▒ ░░▒░ ░ ░   ░▒ ░ ▒░░ ░ ▒  ░
  ░   ▒   ░ ░░ ░  ░ ░    ▒ ░   ░   ░ ░  ░ ░  ░  ░░░ ░ ░   ░░   ░   ░ ░   
      ░  ░░  ░           ░           ░    ░       ░        ░         ░  ░
                                        ░` + DefaultText)
	fmt.Println(strings.Repeat(BoldText+RedText+"="+DefaultText, 73))
}

const Counter = "[00:00:00] █████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░"
