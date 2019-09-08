// Sergey Olisov (c) 2019
// Lesson 1 - NTP time reader

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

func main() {

	// selected the first server from the Russian NTP pool for test.
	// 0.ru.pool.ntp.org
	// 1.ru.pool.ntp.org
	// 2.ru.pool.ntp.org
	// 3.ru.pool.ntp.org

	if time, err := ntp.Time("0.ru.pool.ntp.org"); err != nil {
		log.Println("Error reading NTP time :", err)
		fmt.Println("Local time: ", time)
		os.Exit(1)
	} else {
		fmt.Println("Exact NTP Time: ", time)
		os.Exit(0)
	}
}
