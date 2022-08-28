package twentyfive

import (
	"fmt"
	"time"
)

func Sleep(s int) {
	currentTime := time.Now().Unix()
	fmt.Printf("Sleeping for %d seconds", s)
	for {
		if currentTime+int64(s) <= time.Now().Unix() {
			fmt.Println("\nEnd sleep")
			return
		}
	}
}

func Sleep2(duration time.Duration) {
	<-time.NewTimer(duration * time.Second).C
}

func main() {
	//Sleep(5)
	Sleep2(5)
}
