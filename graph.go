package bandwidth

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/fatih/color"
)

func CreateFile(s []bwData) {
	received := make([]string, 0)
	transmitted := make([]string, 0)
	for _, data := range s {
		entry := fmt.Sprintf("[ %d, %f ]", data.timestamp.Unix(), data.rx)
		received = append(received, entry)
		entry = fmt.Sprintf("[ %d,%f] ", data.timestamp.Unix(), data.tx)
		transmitted = append(transmitted, entry)
	}
	recv := fmt.Sprintf("[" + strings.Join(received, ",") + "]")
	err := ioutil.WriteFile("../web-ui/src/app/received/recv.json", []byte(recv), 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		color.Set(color.FgRed)
		fmt.Print("[*]")
		color.Unset()
		fmt.Println(" Wrote recv.json")
	}

	sent := fmt.Sprintf("[" + strings.Join(transmitted, ",") + "]")
	err = ioutil.WriteFile("../web-ui/src/app/sent/sent.json", []byte(sent), 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		color.Set(color.FgRed)
		fmt.Print("[*]")
		color.Unset()
		fmt.Println(" Wrote sent.json")
	}

}
