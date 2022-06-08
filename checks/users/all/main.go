package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"regexp"
	"strings"

	"github.com/pwnedshell/quimera-lab/utils"
)

var (
	table   = utils.Table{}
	decoder Decoder
	path    = "/etc/passwd"
	re      = regexp.MustCompile(`USERS=\/proc\/\$\$\/fd\/([0-9]+)`)
)

type Decoder struct {
	Groups []Group `json:"groups"`
}

type Group struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}

func main() {
	table.Header = []string{"User", "Name", "Gid", "Groups"}
	match := re.FindStringSubmatch(os.Getenv("IMPORTS"))
	if len(match) != 2 {
		fmt.Println("USERS env variable not found")
		os.Exit(2)
	}
	memfd := fmt.Sprintf("/proc/%d/fd/%s", os.Getpid(), match[1])

	file, err := ioutil.ReadFile(memfd)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	if err := json.Unmarshal(file, &decoder); err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	passwd, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	defer passwd.Close()
	scanner := bufio.NewScanner(passwd)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}
		retrieveInfo(strings.Split(line, ":")[0])
	}
	json, err := json.Marshal(table)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	fmt.Println(string(json))
	os.Exit(0)
}

func retrieveInfo(username string) {
	u, err := user.Lookup(username)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	uid, err := user.LookupId(u.Uid)
	if err != nil {
		return
	}
	gid, err := user.LookupGroupId(u.Gid)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	groupsIds, err := u.GroupIds()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	var userGroups []any
	for _, id := range groupsIds {
		group, err := user.LookupGroupId(id)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}
		userGroups = append(userGroups, unit(group.Name))
	}
	table.Body = append(table.Body, []any{username, uid.Name, gid.Name, userGroups})
}

func unit(group string) any {
	for _, g := range decoder.Groups {
		if g.Name == group {
			return utils.Unit{Text: g.Name, Level: g.Level}
		}
	}
	return group
}
