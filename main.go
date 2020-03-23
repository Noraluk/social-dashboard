package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

const layout = "2006-01-02 15:04:05"
const layout2 = "2006-01-02"

var daily = make(map[string]int)

func main() {
	start := time.Now()
	lines, err := readLines("./rawdata.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(lines))
	log.Println("time diff : ", time.Now().Sub(start).String())
}

func readLines(filename string) ([][]string, error) {
	var lines [][]string
	var data []string
	var isFirstLine bool
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(file)
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		datas := strings.Split(line, ",")
		if isFirstLine {
			for _, val := range datas {
				switch len(data) {
				case 0, 1, 2, 4, 5, 6:
					data = append(data, val)
				case 3:
					t, err := time.Parse(layout, val)
					if err == nil {
						data = append(data, t.String())
					} else {
						data[len(data)-1] = fmt.Sprintf("%s %s", data[len(data)-1], val)
					}
				case 7:
					lines = append(lines, append(data, val))
					data = nil
				}
			}
		} else {
			lines = append(lines, datas)
			isFirstLine = true
		}
	}
	return lines, nil
}
