package hostAgents

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

/**
 1) Should run on multiple operating systems.
 2) The correct values should be uploaded to aggregation service.
 3) Keep sending it periodically instead of just sending once.
 */
func GetCPUSample() {

	sum := 0
	for {
		percentage, err := cpu.Percent(0, false)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(percentage)
		sum++
		fmt.Println("Wait 2 seconds")
		time.Sleep(time.Duration(2)*time.Second)
	}
	fmt.Println(sum)

}

