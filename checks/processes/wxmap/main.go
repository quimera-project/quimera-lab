package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/pwnedshell/quimera-lab/utils"
)

var (
	procfs = "/proc"
	re     = regexp.MustCompile(`^[0-9]*$`)
	wg     = &sync.WaitGroup{}
	table  = &utils.Table{}
	mutex  = sync.Mutex{}
)

func main() {
	table.Header = []string{"PID", "Name", "Line nº", "Address", "Perms", "Offset", "Dev", "Inode", "Pathname"}
	files, err := ioutil.ReadDir(procfs)
	if err != nil {
		os.Exit(1)
	}
	for _, f := range files {
		if f.IsDir() && re.MatchString(f.Name()) {
			wg.Add(1)
			go searchWX(f)
		}
	}
	wg.Wait()
	json, err := json.Marshal(table)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(json))
	os.Exit(0)
}

func searchWX(dir fs.FileInfo) {
	var (
		pid = dir.Name()
	)
	defer wg.Done()
	file, err := os.Open(filepath.Join(procfs, pid, "maps"))
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var i int
	for scanner.Scan() {
		i++
		line := strings.Fields(scanner.Text())
		if line[1][1:3] == "wx" {
			var wx []any
			wx = append(wx, pid)
			wx = append(wx, getNameFromPid(pid))
			wx = append(wx, i)
			wx = append(wx, line[0])
			wx = append(wx, &utils.Unit{Text: line[1], Level: "high"})
			wx = append(wx, line[2])
			wx = append(wx, line[3])
			wx = append(wx, line[4])
			wx = append(wx, getPathName(line))
			if err != nil {
				return
			}
			mutex.Lock()
			table.Body = append(table.Body, wx)
			mutex.Unlock()
		}
	}
}

func getNameFromPid(pid string) string {
	file, err := os.Open(filepath.Join(procfs, pid, "status"))
	if err != nil {
		return ""
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := strings.Fields(scanner.Text())
	if len(line) > 1 {
		return line[1]
	}
	return ""
}

func getPathName(line []string) string {
	if len(line) > 5 {
		return line[5]
	}
	return ""
}
