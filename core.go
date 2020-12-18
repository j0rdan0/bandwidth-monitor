package bandwidth

import (
	"fmt"
	"io/ioutil"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getInterfaces() []string {
	interfacesNames := make([]string, 0)
	interfaces, err := net.Interfaces()
	checkErr(err)
	for i := 0; i < len(interfaces); i++ {
		// || strings.HasPrefix(interfaces[i].Name, "et") || strings.HasPrefix(interfaces[i].Name, "wl")
		if strings.HasPrefix(interfaces[i].Name, "wl") {
			interfacesNames = append(interfacesNames, interfaces[i].Name)
		}
	}
	return interfacesNames
}

func getData(interf string) (float64, float64, time.Time) {
	fileSuffix := []string{"rx_bytes", "tx_bytes"}
	var rx, tx int
	var gbConv = 9.31 * math.Pow(10, -10) // 1 Bytes = 9.31×10-10 Gigabytes
	for _, suffix := range fileSuffix {
		filename := path + interf + "/statistics/" + suffix
		buff, err := ioutil.ReadFile(filename)
		checkErr(err)
		if suffix == "rx_bytes" {
			rx, _ = strconv.Atoi(strings.Trim(string(buff), "\n"))
		} else {
			tx, _ = strconv.Atoi(strings.Trim(string(buff), "\n"))
		}
	}
	rxBytes := float64(rx) * gbConv
	txBytes := float64(tx) * gbConv

	return rxBytes, txBytes, time.Now()
}

func ComposeData() []bwData {
	interfaces := getInterfaces()
	for i := 0; i < 30; i++ {
		for _, intf := range interfaces {
			rx, tx, t := getData(intf)
			Stats = append(Stats, bwData{rx, tx, t})
			time.Sleep(time.Minute * 1)
		}
	}
	return Stats
}

func PrintData(s []bwData) {
	fmt.Printf("RX\tTX\tDATE\n")
	for _, data := range s {
		fmt.Printf("%.2f\t %.2f\t %s \n", data.rx, data.tx, strings.Split(data.timestamp.String(), ".")[0])
	}
	writeFile(s)
}
func writeFile(s []bwData) {
	fRX, err := os.Create("output/rx")
	defer fRX.Close()
	checkErr(err)
	fTX, err := os.Create("output/tx")
	checkErr(err)
	defer fRX.Close()
	for _, data := range s {
		fRX.WriteString(fmt.Sprintln(data.rx, strings.Split(data.timestamp.String(), ".")[0]))
		fTX.WriteString(fmt.Sprintln(data.tx, strings.Split(data.timestamp.String(), ".")[0]))
	}

}
