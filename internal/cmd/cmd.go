package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/XiXiangFiles/gitcralwer/internal/model"
)

func ParseMessage(dir, scopePrefix string) string {
	commitments := []*model.Commitment{}
	commitsBytes, _ := ioutil.ReadFile(dir)
	_ = json.Unmarshal(commitsBytes, &commitments)

	tickets := make(map[string][]string)
	for _, cmt := range commitments {
		matched, _ := regexp.MatchString(`[a-z]{1,9}\([A-Z]{1,10}\-[0-9]{1,6}\)`, cmt.Message)
		if matched {
			tickets[cmt.Author] = append(tickets[cmt.Author], scopePrefix+extractTicket(cmt.Message))
		}
	}
	memo := TicketToTxt(tickets)
	return memo
}

func extractTicket(str string) string {
	start := 0
	end := 0
	for i, c := range str {
		c := string(c)
		if c == "(" {
			start = i + 1
		}
		if c == ")" {
			end = i
		}
	}
	return str[start:end]
}

func reduceTicket(tickets []string) []string {
	reducer := make(map[string]int)
	for _, v := range tickets {
		v := v
		reducer[v] += 1
	}
	dst := []string{}
	for title, _ := range reducer {
		dst = append(dst, title)
	}
	return dst
}

func TicketToTxt(tickets map[string][]string) string {
	str := ""
	for author, ticket := range tickets {
		ts := reduceTicket(ticket)
		str += fmt.Sprintf("author = %s\ntickets = \n%s\n\n", author, strings.Join(ts, "\n"))
	}
	return str
}
