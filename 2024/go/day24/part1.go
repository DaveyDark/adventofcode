package day24

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
)

type Gate struct {
	gateType string
	in1      string
	in2      string
	out      string
}

func (g Gate) output(wires *map[string]*Wire) (bool, error) {
	inwire1, ok := (*wires)[g.in1]
	if !ok || inwire1.value == nil {
		return false, fmt.Errorf("Wire %s not found", g.in1)
	}
	inwire2, ok := (*wires)[g.in2]
	if !ok || inwire2.value == nil {
		return false, fmt.Errorf("Wire %s not found", g.in2)
	}
	switch g.gateType {
	case "AND":
		return and(*inwire1.value, *inwire2.value), nil
	case "OR":
		return or(*inwire1.value, *inwire2.value), nil
	case "XOR":
		return xor(*inwire1.value, *inwire2.value), nil
	default:
		return false, fmt.Errorf("Unknown gate type %s", g.gateType)
	}
}

type Wire struct {
	value *bool
}

func NewWire() *Wire {
	return &Wire{
		value: nil,
	}
}

func NewWireWithValue(value bool) *Wire {
	return &Wire{
		value: &value,
	}
}

func (w *Wire) Set(value bool) {
	w.value = &value
}

func init() {
	registry.Registry["day24/part1"] = solve
}

func solve(inputString string) (int64, error) {
	input, err := os.Open(inputString)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(input)

	wires := map[string]*Wire{}

	// Parse section 1: initial config of wires
	wireRegex := regexp.MustCompile(`(\w+): (\d)`)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			// End of section
			break
		}

		matches := wireRegex.FindStringSubmatch(line)
		wire := matches[1]
		val := false
		if matches[2] == "1" {
			val = true
		}
		wires[wire] = NewWireWithValue(val)
	}

	// Parse Section 2: Gates
	gateRegex := regexp.MustCompile(`(\w+) (OR|XOR|AND) (\w+) -> (\w+)`)
	gates := []Gate{}

	for scanner.Scan() {
		line := scanner.Text()
		matches := gateRegex.FindStringSubmatch(line)
		in1 := matches[1]
		gateType := matches[2]
		in2 := matches[3]
		op := matches[4]
		_, ok := wires[in1]
		if !ok {
			wires[in1] = NewWire()
		}
		_, ok = wires[in2]
		if !ok {
			wires[in2] = NewWire()
		}
		_, ok = wires[op]
		if !ok {
			wires[op] = NewWire()
		}
		gate := Gate{
			gateType: gateType,
			in1:      in1,
			in2:      in2,
			out:      op,
		}
		gates = append(gates, gate)
	}

	// Process gates
	for len(gates) > 0 {
		newGates := []Gate{}
		for _, gate := range gates {
			out, err := gate.output(&wires)
			if err != nil {
				newGates = append(newGates, gate)
			} else {
				wires[gate.out].Set(out)
			}
		}
		gates = newGates
	}

	// Calculate output
	res := strings.Builder{}
	for i := 0; ; i++ {
		_wire := fmt.Sprintf("z%02d", i)
		wire, ok := wires[_wire]
		if !ok || wire.value == nil {
			break
		}
		if *wire.value {
			res.WriteString("1")
		} else {
			res.WriteString("0")
		}
	}
	resStr := reverse(res.String())
	resInt, err := strconv.ParseInt(resStr, 2, 64)
	if err != nil {
		return 0, err
	}

	return resInt, nil
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func xor(b1, b2 bool) bool {
	return (b1 && !b2) || (!b1 && b2)
}

func or(b1, b2 bool) bool {
	return b1 || b2
}

func and(b1, b2 bool) bool {
	return b1 && b2
}
