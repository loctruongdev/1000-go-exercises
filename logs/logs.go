package logs

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Logs() {
	// var sum map[string]int
	sum := map[string]int{}

	var (
		domains []string
		total   int
		lines   int
	)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		lines++

		fields := strings.Fields(in.Text())

		if len(fields) != 2 {
			fmt.Printf("wrong input: %v (lines #%d)\n", fields, lines)
			return
		}

		domain := fields[0]

		visits, err := strconv.Atoi(fields[1])
		if err != nil || visits < 0 {
			fmt.Printf("wrong input: %v (lines #%d)\n", fields, lines)
			return
		}

		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}
		total += visits
		sum[domain] += visits
		// fmt.Printf("domain: %s - visits: %s\n", fields[0], fields[1])

	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	sort.Strings(domains)

	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}

	fmt.Printf("\n%-30s %10d\n", "TOTAL", total)

	if err := in.Err(); err != nil {
		fmt.Println("> Err: ", err)
	}
}

func LogsB() {
	// var sum map[string]int
	sum := map[string]int{}

	// var domains []string

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		fields := strings.Fields(in.Text())

		domain := fields[0]

		visits, _ := strconv.Atoi(fields[1])

		// if _, ok := sum[domain]; !ok {
		// 	domains = append(domains, domain)
		// }

		if in.Scan() {
			sum[domain] += visits
		}

	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	var domains []string
	for domain := range sum {
		domains = append(domains, domain)
		sort.Strings(domains)
	}

	for domain, visits := range sum {
		fmt.Printf("%-30s %10d\n", domain, visits)
	}

}
