package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kenshaw/imdb"
	"encoding/json"
)

func main() {

	var IMDBIdFile, Str string
	var IMDBId []string

	IMDBIdFile = "IMDBIdList.txt"
	f, err := os.Open(IMDBIdFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer f.Close()
	r := bufio.NewReader(f)
	line, err := r.ReadString(10)
	for err != io.EOF {
		IMDBId = append(IMDBId, strings.TrimSuffix(line, "\n"))
		line, err = r.ReadString(10)
	}

	cl := imdb.New("9ac5feef")

	for s := range IMDBId {
		Str = IMDBId[s]
		res, err := cl.MovieByImdbID(Str)
		if err != nil {
			fmt.Println("Noe er feil")
		} else {
			data, err := json.Marshal(res)
			if err == nil {
				fmt.Printf("%s\n\n",data)
				}
		}
	}

}
