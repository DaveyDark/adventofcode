package day23

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/emirpasic/gods/sets/hashset"
)

func init() {
	registry.Registry["day23/part2"] = solve2
}

func solve2(inputString string) (int64, error) {
	input, err := os.Open(inputString)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(input)

	parties := []*hashset.Set{}
	connections := map[string]*hashset.Set{} // Adjaency list
	universalSet := hashset.New()

	// Construct the graph
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "-")
		left := split[0]
		right := split[1]

		// All connections are parties by themselves
		set := hashset.New()
		set.Add(left)
		set.Add(right)
		parties = append(parties, set)

		// Add the connection to the list of connections
		leftSet, ok := connections[left]
		if !ok {
			leftSet = hashset.New()
			connections[left] = leftSet
		}
		leftSet.Add(right)

		rightSet, ok := connections[right]
		if !ok {
			rightSet = hashset.New()
			connections[right] = rightSet
		}
		rightSet.Add(left)

		universalSet.Add(left)
		universalSet.Add(right)
	}

	// Start merging parties by finding common connections until no more merges can be done
	success := true
	for {
		parties, success = mergeParties(parties, connections, universalSet)
		if !success {
			break
		}
		parties = cleanParties(parties)
	}

	// Find largest party
	largest := hashset.New()
	for _, party := range parties {
		if party.Size() > largest.Size() {
			largest = party
		}
	}

	// Display password for largest party
	fmt.Println(calculatePassword(largest))

	return int64(largest.Size()), nil
}

func cleanParties(parties []*hashset.Set) []*hashset.Set {
	uniqueParties := hashset.New()
	newParties := []*hashset.Set{}

	for _, party := range parties {
		password := calculatePassword(party)
		if !uniqueParties.Contains(password) {
			uniqueParties.Add(password)
			newParties = append(newParties, party)
		}
	}

	return newParties
}

func calculatePassword(party *hashset.Set) string {
	// Get all computers in the party
	computers := []string{}
	for _, v := range party.Values() {
		computers = append(computers, v.(string))
	}

	// Sort the computers
	sort.Strings(computers)

	return strings.Join(computers, ",")
}

func mergeParties(parties []*hashset.Set, connections map[string]*hashset.Set, universalSet *hashset.Set) ([]*hashset.Set, bool) {
	newParties := []*hashset.Set{}
	success := false
	for _, party := range parties {
		// Find intersection of connections of the party
		set := hashset.New(universalSet.Values()...)
		for _, val := range party.Values() {
			set = set.Intersection(connections[val.(string)])
		}

		// Check if any common connections were found
		if len(set.Values()) == 0 {
			newParties = append(newParties, party)
			continue
		}
		success = true

		// Create new parties with the common connections
		for _, common := range set.Values() {
			newParty := hashset.New()
			newParty.Add(common)
			newParty = newParty.Union(party)
			newParties = append(newParties, newParty)
		}
	}

	return newParties, success
}
