package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Config struct {
	Source string   `json:"source"`
	Column int      `json:"column"`
	Format []string `json:"format"`
	Output string   `json:"output"`
}

func printUsage() {
	fmt.Printf(`
Usage:
    %s config.json

Configure file, json type
    {
      "source": "data.txt",
      "output": "output.sql",
      "column": 3,
      "format": [
	    "update users set name='#2#', phone=#3#",
		"where uid='#1#';"
	  ]
    }

Source file, text type, separatd by TAB
    60001	Tony	18612345678
	60002	Alex	15812345678
	`, os.Args[0])
}

func ReplaceColumns(f []byte, n int, cols [][]byte) []byte {
	b := f
	for i := 0; i < n; i++ {
		b = bytes.Replace(b, []byte(fmt.Sprintf("#%d#", i+1)), cols[i], -1)
	}
	return b
}

func Cross(args []string) error {
	d, e := ioutil.ReadFile(args[1])
	if e != nil {
		fmt.Printf("错误: %s\n", e.Error())
		return e
	}

	conf := Config{}
	e = json.Unmarshal(d, &conf)
	if e != nil {
		fmt.Printf("错误: %s\n", e.Error())
		return e
	}

	f, e := os.Open(conf.Source)
	if e != nil {
		fmt.Printf("错误: %s\n", e.Error())
		return e
	}
	defer f.Close()

	t := []byte("\t")

	format := [][]byte{}
	for _, f := range conf.Format {
		format = append(format, []byte(f))
	}

	ss := [][]byte{}

	rb := bufio.NewReader(f)
	i := 0
	for {
		i++
		line, _, e := rb.ReadLine()
		if e != nil {
			if e == io.EOF {
				break
			}
			fmt.Printf("第%d行有错误: %s\n", i, e.Error())
			fmt.Printf("错误行内容: %s\n", string(line))
			fmt.Printf("忽略第%d行\n", i)
			continue
		}
		cols := bytes.Split(line, t)
		if conf.Column > len(cols) {
			fmt.Printf("错误: 第%d行数量不足\n", i)
			fmt.Printf("忽略第%d行\n", i)
			continue
		}
		for _, f := range format {
			s := ReplaceColumns(f, conf.Column, cols)
			ss = append(ss, s)
		}
	}

	out := bytes.Join(ss, []byte("\n"))
	e = ioutil.WriteFile(conf.Output, out, os.ModePerm)
	if e != nil {
		fmt.Printf("错误: %s\n", e.Error())
		return e
	}
	fmt.Printf("数据生成至: %s\n", conf.Output)

	return nil
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	Cross(os.Args)
}
