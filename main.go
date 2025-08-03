package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func parseTime(input string) (time.Time, error) {
	layout := "3:04 PM"
	return time.Parse(layout, input)
}

func returnCursor() {
	fmt.Print("\033[F\r\033[K")
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var total time.Duration

	fmt.Println("Format 'HH:MM AM/PM'. Use 'stop' to stop.")

	for {
		fmt.Print("> ")
		if !s.Scan() {
			break
		}
		l := strings.TrimSpace(s.Text())

		if strings.ToLower(l) == "stop" {
			returnCursor()
			fmt.Printf("> %s -> %v\n", l, total)
			break
		}

		ps := strings.Fields(l)
		if len(ps) < 4 {
			returnCursor()
			fmt.Printf("> %s -> Invalid format. Use: 'HH:MM AM/PM HH:MM PM'\n", l)
			continue
		}

		start, err1 := parseTime(ps[0] + " " + ps[1])
		end, err2 := parseTime(ps[2] + " " + ps[3])

		if err1 != nil || err2 != nil {
			returnCursor()
			fmt.Printf("> %s -> Invalid time\n", l)
			continue
		}

		if end.Before(start) {
			end = end.Add(24 * time.Hour)
		}

		duration := end.Sub(start)
		total += duration

		returnCursor()
		fmt.Printf("> %s -> %v\n", l, duration)
	}
}
