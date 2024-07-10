package writingingo

import (
	"fmt"
	"os"
)

func Writeinfile(investmentAmount int) {
	balancetext := fmt.Sprint(investmentAmount)
	os.WriteFile("balance.txt", []byte(balancetext), 0644)
}
