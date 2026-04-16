package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const version = "1.2.0"

// parseTime parses a string representing a time in various formats.
//
// Supported formats: "3:04pm", "3:04 PM" (with space, uppercase), "15:04" (24-hour).
func parseTime(input string) (time.Time, error) {
	input = strings.TrimSpace(input)
	for _, format := range []string{"15:04", "3:04pm", "3:04 PM"} {
		if t, err := time.Parse(format, input); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time: %s", input)
}

// returnCursor moves the cursor back to the beginning of the line and clears it.
func returnCursor() {
	fmt.Print("\033[F\r\033[K")
}

// printUsage prints the usage information for the duration program.
func printUsage() {
	fmt.Println("Formats: 12-hour 'hh:mm AM/PM' or 'hh:mmam/pm' or 24-hour 'HH:mm'.")
	fmt.Println("Examples: '09:00pm 1:30pm' or '09:00 AM 1:30 PM' or '09:00 13:30'.")
	fmt.Println("Use 'stop' to stop.")
}

// printHelp prints the help message with program name, flags, and usage.
func printHelp() {
	fmt.Println("duration - Time duration calculator")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -h, --help      Show this help message")
	fmt.Println("  -v, --version   Show version")
	fmt.Println()
	fmt.Println("Usage:")
	printUsage()
}

// printVersion prints the program version.
func printVersion() {
	fmt.Printf("duration version %s\n", version)
}

func main() {
	// Define flags
	helpFlag := flag.Bool("help", false, "Show help message")
	versionFlag := flag.Bool("version", false, "Show version")

	// Allow both -h/--help and -v/--version
	flag.BoolVar(helpFlag, "h", false, "Show help message")
	flag.BoolVar(versionFlag, "v", false, "Show version")

	flag.Parse()

	// Handle flags
	if *helpFlag {
		printHelp()
		return
	}

	if *versionFlag {
		printVersion()
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	var total time.Duration

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())

		if strings.ToLower(line) == "stop" {
			returnCursor()
			fmt.Printf("> %s -> %v\n", line, total)
			break
		}

		words := strings.Fields(line)
		if len(words) < 2 {
			returnCursor()
			fmt.Printf("> %s -> Invalid format. Provide at least 2 times.\n", line)
			continue
		}

		var ss, es string
		if len(words) == 2 {
			ss = words[0]
			es = words[1]
		} else if len(words) == 4 {
			ss = words[0] + " " + words[1]
			es = words[2] + " " + words[3]
		} else {
			returnCursor()
			fmt.Printf("> %s -> Invalid format.\n", line)
			continue
		}

		start, err1 := parseTime(ss)
		end, err2 := parseTime(es)

		if err1 != nil || err2 != nil {
			returnCursor()
			fmt.Printf("> %s -> Invalid time format\n", line)
			continue
		}

		if end.Before(start) {
			end = end.Add(24 * time.Hour)
		}

		duration := end.Sub(start)
		total += duration

		returnCursor()
		fmt.Printf("> %s -> %v\n", line, duration)
	}
}
