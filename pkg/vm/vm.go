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
		}
	case 0x1000: // 0x1NNN jumps to address nnn
		vm._0x1000()
		break
	case 0x2000 // 0x2NNN calls subroutine at NNN
		vm._0x2000()
		break
	case 0x3000:  // 0x3XNN skips the next instruction if vx equals nn
		vm._0x3000()
	case 0x4000: // 0x4XNN skips the next instruction if vx doesn't equal NN
		vm._0x4000()
	case 0x5000: // 0x5XY0 skips the next instruction if vx equals vy
		vm._0x5000()
	case 0x6000: // 0x6XNN sets vx to nn
		vm._0x6000()
	case 0x7000: // 0x7NN adds NN to vy
		vm._0x7000()
	}
}

func (vm *VM) fetchOpcode() {
	vm.opcode = uint16(vm.mem[vm.pc])<<8 | uint16(vm.mem[vm.pc+1])
}
