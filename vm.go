package vm

const (
	Load  = 0x01
	Store = 0x02
	Add   = 0x03
	Sub   = 0x04
	Halt  = 0xff
)

// Stretch goals
const (
	Addi = 0x05
	Subi = 0x06
	Jump = 0x07
	Beqz = 0x08
)

// Given a 256 byte array of "memory", run the stored program
// to completion, modifying the data in place to reflect the result
//
// The memory format is:
//
// 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f ... ff
// __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ ... __
// ^==DATA===============^ ^==INSTRUCTIONS==============^
func compute(memory []byte) {

	registers := [3]byte{8, 0, 0} // PC, R1 and R2

	// Keep looping, like a physical computer's clock
	for {
		pc := registers[0]
		op := memory[pc]
		registers[0] = pc + 1
		switch op {
		case Load:
			pc = registers[0]
			reg := memory[pc]
			registers[0] = pc + 1

			pc = registers[0]
			memPtr := memory[pc]
			registers[0] = pc + 1

			registers[reg] = memory[memPtr]

		case Store:
			pc = registers[0]
			reg := memory[pc]
			registers[0] = pc + 1

			pc = registers[0]
			memPtr := memory[pc]
			registers[0] = pc + 1

			memory[memPtr] = registers[reg]

		case Add:
			pc = registers[0]
			reg1 := memory[pc]
			registers[0] = pc + 1

			pc = registers[0]
			reg2 := memory[pc]
			registers[0] = pc + 1

			registers[reg1] = registers[reg1] + registers[reg2]

		case Sub:
			pc = registers[0]
			reg1 := memory[pc]
			registers[0] = pc + 1

			pc = registers[0]
			reg2 := memory[pc]
			registers[0] = pc + 1

			registers[reg1] = registers[reg1] - registers[reg2]

		case Addi:
			pc = registers[0]
			reg := memory[pc]
			registers[0] = pc + 1

			pc = registers[0]
			addend := memory[pc]
			registers[0] = pc + 1

			registers[reg] = registers[reg] + addend

		case Subi:
			pc = registers[0]
			reg := memory[pc]
			registers[0] = pc + 1

			pc = registers[0]
			subtrahend := memory[pc]
			registers[0] = pc + 1

			registers[reg] = registers[reg] - subtrahend

		case Jump:
			pc = registers[0]
			newPc := memory[pc]
			registers[0] = newPc

		case Beqz:
			pc = registers[0]
			reg := memory[pc]
			registers[0] = pc + 1

			pc = registers[0]
			offset := memory[pc]
			registers[0] = pc + 1

			if registers[reg] == 0 {
				pc = registers[0]
				registers[0] = pc + offset
			}

		case Halt:
			return

		default:
			return
		}
	}
}
