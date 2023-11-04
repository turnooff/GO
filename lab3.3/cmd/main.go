package main

import (
	"bufio"
	"bytes"
	"chech_result/internal/cmd"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/go-yaml/yaml"
)

type instanceConfig struct {
	NameTextFile string `yaml:"Name"`
}

func Parse(in []byte, out *[]instanceConfig) error {
	r := bytes.NewReader(in)
	decoder := yaml.NewDecoder(r)
	for {
		var bo instanceConfig
		if err := decoder.Decode(&bo); err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		*out = append(*out, bo)
	}
	return nil
}

func main() {
	data, err := os.ReadFile("file.yml")
	if err != nil {
		log.Fatal(err)
		return
	}
	var config []instanceConfig
	if err := Parse(data, &config); err != nil {
		log.Fatal(err)
		return
	}
	for _, order := range config {
		a := cmd.NewApp()
		fmt.Print("Введите подстроку, наличие которой хотите проверить: ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		line = strings.TrimSpace(line)
		flag, err := a.Search(order.NameTextFile, line)
		fmt.Println("Результат: ", flag)
	}
}
