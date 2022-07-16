package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func listwifissid() {
	list := []string{}
	cmd := exec.Command("netsh", "wlan", "show", "profiles")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("system command not worked")
	}
	m := regexp.MustCompile(`All User Profile     :(.*)`)
	for _, match := range m.FindAllString(string(out), -1) {
		res := strings.SplitN(match, ":", 2)
		list = append(list, strings.TrimSpace(res[1]))

	}
	lastmatch := regexp.MustCompile(`Key Content            : (.*)`)
	for _, val := range list {
		nextcmd := exec.Command("netsh", "wlan", "show", "profiles", val)
		cmd, err := nextcmd.CombinedOutput()
		if err != nil {
			log.Fatal("system command not worked")
		}
		matcher, _ := regexp.MatchString(`Security key           : Absent`, string(cmd))
		if matcher == true {
			continue
		} else {
			finalcmd, err := exec.Command("netsh", "wlan", "show", "profiles", val, "key=clear").CombinedOutput()
			if err != nil {
				log.Fatal("System command not worked")
			}
			result := lastmatch.FindAllString(string(finalcmd), -1)
			for _, match := range result {
				res := strings.SplitN(match, ":", 2)
				output := strings.TrimSpace(res[1])
				fmt.Printf("%s : %s\n", val, output)

			}
		}

	}
}

func main() {
	listwifissid()
}
