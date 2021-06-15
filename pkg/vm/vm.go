package vm

const (
	windowSize = 64 * 32
)

// VM defines the actual Virtual Machine
type VM struct {
	// current opcode
	opcode uint16

	// main memory
	mem [4096]byte

	// general purpose registers from v0 to vf
	v [16]byte

	// index register (from 0x000 to 0xFFF)
	i uint16

	// program counter (from 0x000 to 0xFFF)
	pc uint16

	// main program stack
	stack [16]uint16

	// stack pointer
	sp uint16

	// window
	gfx [windowSize]byte

	// delay timer
	delayTimer byte

	// sound timer
	soundTimer byte

	// keypad
	keypad [16]byte

	// we don't want to draw screen on every cycle,
	// so we update the drawFlag when we need to update screen
	drawFlag bool

	// audio channel
	audioCh chan struct{}
}

// New creates a new CHIP-8 VM
func New() (vm *VM, err error) {
	// init vm
	vm.init()

	return vm, err
}

// init initializes the VM
func (vm *VM) init() {

	// set program counter to beginning of the memory
	vm.pc = 0x200

	// init audiochan
	vm.audioCh = make(chan struct{})

	// TODO: load fontset
	// vm.loadFontset()

	// TODO: init keypad

}

// Reset resets the VM
func (vm *VM) Reset() {

	// reset current opcode
	vm.opcode = 0

	// reset memory
	vm.mem = [4096]byte{}

	// reset registers
	vm.v = [16]byte{}

	// reset index register
	vm.i = 0

	// reset program counter to beginning of the memory
	vm.pc = 0x200

	// reset stack pointer
	vm.sp = 0

	// reset gfx
	vm.gfx = [windowSize]byte{}

	// reset delay timer
	vm.delayTimer = 0

	// reset sound timer
	vm.soundTimer = 0

	// reset draw flag
	vm.drawFlag = false

	// TODO: load fontset
	// vm.loadFontset()

	// TODO: init keypad

}

func (vm *VM) emulateCycle() {
	// fetch opcode to vm.opcode
	vm.fetchOpcode()

	switch vm.opcode & 0xF000 {
	case 0x0000:
		switch vm.opcode & 0x000F {
		case 0x0000: // 0x00E0 clear the display
			vm._0x00E0()
			break
		case 0x000E: // 0x00EE returns from subroutine
			vm._0x00EE()
			break
		default:
			// TODO: throw an error
		}
	case 0x1000: // 0x1NNN jumps to address nnn
		vm._0x1NNN()
		break
	case 0x2000 // 0x2NNN calls subroutine at NNN
		vm._0x2NNN()
		break
	case 0x3000:  // 0x3XNN skips the next instruction if vx equals nn
		vm._0x3XNN()
		break
	case 0x4000: // 0x4XNN skips the next instruction if vx doesn't equal NN
		vm._0x4XNN()
		break
	case 0x5000: // 0x5XY0 skips the next instruction if vx equals vy
		vm._0x5XY0()
		break
	case 0x6000: // 0x6XNN sets vx to nn
		vm._0x6XNN()
		break
	case 0x7000: // 0x7NN adds NN to vy
		vm._0x7NN()
		break
	case 0x8000:
		switch vm.opcode & 0x000F {
		case 0x0000: // 0x8XY0 sets vx to the value of vy
			vm._0x8XY0()
			break
		case 0x0001: // 0x8XY1 sets vx to the value of "vx or vy" 
			vm._0x8XY1()
			break
		case 0x0002: // 0x8XY2 sets vx to the value of "vx and vy" 
			vm._0x8XY2()
			break
		case 0x0003: // 0x8XY3 sets vx to the value of "vx xor vy" 
			vm._0x8XY3()
			break
		case 0x0004:
			vm._0x8XY4() // 0x8XY4 adds vy to vx. vf is set to 1 when there's a carry
			break
		case 0x0005:
			vm._0x8XY5() // 0x8XY5 vy is subtracted from vx. vf is set to 0 when there's a borrow
			break
		case 0x0006:
			vm._0x8XY6() // 0x8XY6 shifts vx right by one. vf is set to the value of the least significant bit of vx before the shift
			break
		case 0x0007:
			vm._0x8XY7() // 0x8XY7 sets vx to vy - vx. vf is set to 0 when there is a borrow
			break
		case 0x000E:
			vm._0x8XYE() // 0x8XYE shifts vx left by one
			break
		}
	case 0x9000: // 0x9XY0 skips the next instruction if VX doesn't equal VY
		vm._0x9XY0()
		break
	case 0xA000: // 0xANNN sets I to the address NNN
		vm._0xANNN()
		break
	case 0xB000: // 0xBNNN jumps to the address NNN plus V0
		vm._0xBNNN()
		break
	case 0xC000: // 0xCXNN sets VX to a random number and NN
		vm._0xCXNN()
		break
	case 0xD000: // 0xDXYN: Draws a sprite at coordinate (VX, VY) that has a width of 8 pixels and a height of N pixels
		vm._0xDXYN()
		break 
	}
}

func (vm *VM) fetchOpcode() {
	vm.opcode = uint16(vm.mem[vm.pc])<<8 | uint16(vm.mem[vm.pc+1])
}
