package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"golang.org/x/sys/unix"

	config "du_alerts/components/config"
)

// ToDo: implement alert routine
func usageTresholdExceed() {

}

func main() {

	appConfig := config.GetConfiguration(os.Args[1:])
	var stat unix.Statfs_t

	for {
		unix.Statfs(appConfig.WatchDir, &stat)
		totalStorageBytes := stat.Blocks * uint64(stat.Bsize)
		usedStorageBytes := totalStorageBytes - stat.Bavail*uint64(stat.Bsize)

		usage := float64(100*usedStorageBytes) / float64(totalStorageBytes)
		if math.IsNaN(usage) {
			fmt.Println("Unable to ditermine storage.. exiting")
			os.Exit(1)
		}

		if usage >= float64(appConfig.Treshold) {
			usageTresholdExceed()
		}

		fmt.Printf("Used %.2f %% of storage\n", usage)
		time.Sleep(time.Duration(appConfig.Refresh) * time.Minute)
	}

}
