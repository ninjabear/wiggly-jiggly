package main

import "fmt"
import "bufio"
import "os"
import "time"
import "github.com/go-vgo/robotgo"

func jiggly() {
	time.Sleep(1 * time.Second)
	robotgo.MoveRelative(-10, 0)
	time.Sleep(250 * time.Millisecond)
	robotgo.MoveRelative(10, 0)
}

func doMouseStuff(tx chan bool) {
	for {
		select {
		case <-tx:
			return
		default:
			jiggly()
		}
	}
}

func main() {
	ch := make(chan bool)
	go doMouseStuff(ch)
	fmt.Println("========== wiggly-jiggyly ==============")
	fmt.Println("======== press enter to exit ===========")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	ch <- true
	fmt.Println("=============== bye! ===================")
}
