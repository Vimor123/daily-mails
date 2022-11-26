package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	silverMailFile, _ := os.ReadFile("./quotes/silvermails.txt")
	wardMailFile, _ := os.ReadFile("./quotes/wardmails.txt")

	silverMailString := string(silverMailFile)
	wardMailString := string(wardMailFile)

	silverMails := strings.Split(silverMailString, "\n")
	wardMails := strings.Split(wardMailString, "\n")

	mails := make([]string, 1)
	mails = append(mails, "")

	mailIndex := 0

	for _, line := range silverMails {
		if len(line) > 1 {
			concatedMail := mails[mailIndex] + "\n" + line
			mails[mailIndex] = concatedMail
		} else {
			mails = append(mails, "")
			mailIndex++
		}
	}

	mails = append(mails, "")
	mailIndex++

	for _, line := range wardMails {
		if len(line) > 1 {
			concatedMail := mails[mailIndex] + "\n" + line
			mails[mailIndex] = concatedMail
		} else {
			mails = append(mails, "")
			mailIndex++
		}
	}

	mails = mails[:len(mails)-1]

	rand.Seed(time.Now().UnixNano())
	randomMailIndex := rand.Intn(len(mails))

	fmt.Println(mails[randomMailIndex])
}
