// a terminal command that returns current time
// an option of choosing time zone exists
// For example,
// ```
// $ go run now.go "America/Halifax"
// ```
// default TZ is GMT

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	var prefix string
	var timeZone string

	// room for imporvment: a better sanity check for could be implemented to check
	// the arguments passed
	argsWithProg := os.Args
	if len(argsWithProg) == 1 {
		timeZone = "GMT"
	} else {
		timeZone = os.Args[1]
	}

	// if the timezone provided incorrectly
	location, err := time.LoadLocation(timeZone)
	if err != nil {
		log.Fatalf("Unkown timezone given: `%v`. Refer to https://en.wikipedia.org/wiki/List_of_tz_database_time_zones \n", timeZone)
	}

	// retrieves time from the NTP
	currentTime, err := ntp.Time("0.pool.ntp.org")
	if err != nil {
		// if the network problems
		currentTime = time.Now()
		prefix = "NTP is unavailable. Maybe network issues. Yet the system time:"
	} else {
		prefix = "Network time:"
	}

	formattedTime := currentTime.In(location).Format(time.UnixDate)
	fmt.Printf("%v: %v\n", prefix, formattedTime)
}
