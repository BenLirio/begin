package input

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func Listen() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input listening")
	for {
		fmt.Print("-")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare(text, "exit") == 0 {
			os.Exit(0)
		}
	}
}
