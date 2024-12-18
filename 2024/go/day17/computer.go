package day17

type Registers struct {
	A int
	B int
	C int
}

type Computer struct {
	registers Registers
}

func NewComputer(a, b, c int) *Computer {
	// Return a new instance of Computer
	return &Computer{Registers{a, b, c}}
}

func (this Computer) comboOperand(op int) int {
	// Resolve operand into a register value if applicable
	if op == 4 {
		return this.registers.A
	} else if op == 5 {
		return this.registers.B
	} else if op == 6 {
		return this.registers.C
	}
	return op
}

func (this *Computer) execute(instruction, operand int, output *[]int) int {
	// Execute the given instruction with given operand
	switch instruction {
	case 0:
		this.adv(operand)
	case 1:
		this.bxl(operand)
	case 2:
		this.bst(operand)
	case 3:
		return this.jnz(operand)
	case 4:
		this.bxc(operand)
	case 5:
		*output = append(*output, this.out(operand))
	case 6:
		this.bdv(operand)
	case 7:
		this.cdv(operand)
	default:
		panic("Invalid instruction")
	}
	return -1
}

func (this *Computer) adv(operand int) {
	// Divides A by 2 ^ operand (combo)
	operand = this.comboOperand(operand)
	denominator := 1
	for range operand {
		denominator *= 2
	}
	this.registers.A = this.registers.A / denominator
}

func (this *Computer) bxl(operand int) {
	// XOR of B and operand(literal)
	this.registers.B = this.registers.B ^ operand
}

func (this *Computer) bst(operand int) {
	// store value to B register
	val := this.comboOperand(operand) % 8
	this.registers.B = val
}

func (this *Computer) jnz(operand int) int {
	// Jump if A not zero
	if this.registers.A == 0 {
		return -1
	}
	return operand * 2
}

func (this *Computer) bxc(_ int) {
	// XOR of B and C
	this.registers.B = this.registers.B ^ this.registers.C
}

func (this *Computer) out(operand int) int {
	// Outputs operand(combo)
	return this.comboOperand(operand) % 8
}

func (this *Computer) bdv(operand int) {
	// Divide A by operand(combo) but store in B
	operand = this.comboOperand(operand)
	denominator := 1
	for range operand {
		denominator *= 2
	}
	this.registers.B = this.registers.A / denominator
}

func (this *Computer) cdv(operand int) {
	// Divide A by operand(combo) but store in C
	operand = this.comboOperand(operand)
	denominator := 1
	for range operand {
		denominator *= 2
	}
	this.registers.C = this.registers.A / denominator
}
