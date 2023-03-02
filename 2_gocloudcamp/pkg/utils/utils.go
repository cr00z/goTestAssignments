package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInputString(out string) string {
	fmt.Print(out)
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.Trim(str, "\n")
}

func GetInputInt(out string) int {
	fmt.Print(out)
	var num int
	fmt.Scanf("%d\n", &num)
	return num
}
