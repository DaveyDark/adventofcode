package day23

import (
	"bufio"
	"os"
	"sort"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/emirpasic/gods/sets/hashset"
)

func init() {
	registry.Registry["day23/part1"] = solve
}

func solve(inputString string) (int64, error) {
	input, err := os.Open(inputString)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(input)

	parties := map[string][]string{} // Adjaency list

	// Construct the graph
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "-")
		left := split[0]
		right := split[1]
		parties[left] = append(parties[left], right)
		parties[right] = append(parties[right], left)
	}

	combinations := hashset.New()

	// Find cycles of length 3 starting with t*
	for party, conns := range parties {
		if !strings.HasPrefix(party, "t") {
			continue
		}

		// Create a set of connections
		connSet := hashset.New()
		for _, conn := range conns {
			connSet.Add(conn)
		}

		// Find all connections of the party
		for _, conn := range conns {
			// Find all connections of the connection
			for _, conn2 := range parties[conn] {
				// If the connection is in the set of connections of the initial party
				if connSet.Contains(conn2) {
					// Create a combination
					combo := []string{party, conn, conn2}
					// Sort the combo, so we can compare it to other combos to avoid duplicates
					sort.Strings(combo)
					// Join the combo to a string to add it to the set
					combinations.Add(strings.Join(combo, ""))
				}
			}
		}
	}

	return int64(combinations.Size()), nil
}
