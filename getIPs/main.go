package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"sort"
)

func main() {
	readFile, err := os.Open("file.log")
	check(err)

	defer readFile.Close()

	var ips []string

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanWords)

	for fileScanner.Scan() {
		word := fileScanner.Text()
		if checkIPAddress(word) {
			ips = append(ips, word)
		}
	}

	sort.Strings(ips)

	writeFile, err := os.Create("data.log")
	check(err)

	defer writeFile.Close()

	fileWriter := bufio.NewWriter(writeFile)

	defer fileWriter.Flush()

	for _, ip := range ips {
		_, err := fileWriter.WriteString(ip)
		check(err)
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func checkIPAddress(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	} else {
		return true
	}
}