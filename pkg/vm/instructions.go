package vm

import "math/rand"

// _0xOOE0 clears the display
// TODO: Update
func (vm *VM) _0x00E0() {

}

// _0x00EE returns from subroutine
func (vm *VM) _0x00EE() {
	vm.sp--                 // 16 levels of stach, decrease stack pointer to prevent overwrite
	vm.pc = vm.stack[vm.sp] // put the stored return address from the stack back into the pc
	vm.pc += 2              // increase pc
}

// _0x1NNN jumps to address nnn
func (vm *VM) _0x1NNN() {
	vm.pc = vm.opcode & 0x0FFF
}

// _0x2NNN calls subroutine at nnn
func (vm *VM) _0x2NNN() {
	vm.stack[vm.sp] = vm.pc    // store current address in stack
	vm.sp++                    // increment the stack pointer
	vm.pc = vm.opcode & 0x0FFF // set the pc to address at nnn
}

// _0x3XNN skips the next instruction if vx equals NN
func (vm *VM) _0x3XNN() {

	// decode x and nn
	x := (vm.opcode & 0x0F00) >> 8
	nn := byte(vm.opcode & 0x00FF)

	// skip next instruction
	if vm.v[x] == nn {
		vm.pc += 4
		return
	}

	// continue to next instruction
	vm.pc += 2
}

// _0x4XNN skips the next instruction if vx doesn't equals NN
func (vm *VM) _0x4XNN() {

	// decode x and nn
	x := (vm.opcode & 0x0F00) >> 8
	nn := byte(vm.opcode & 0x00FF)

	// skip next instruction
	if vm.v[x] != nn {
		vm.pc += 4
		return
	}

	// continue to next instruction
	vm.pc += 2
}

// _0x5XY0 skips the next instruction if vx equals vy
func (vm *VM) _0x5XY0() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	y := (vm.opcode & 0x00F0) >> 4

	// skip next instruction
	if vm.v[x] != vm.v[y] {
		vm.pc += 4
		return
	}

	// continue to next instruction
	vm.pc += 2
}

// _0x6XNN sets vx to nn
func (vm *VM) _0x6XNN() {

	// decode x and nn
	x := (vm.opcode & 0x0F00) >> 8
	nn := byte(vm.opcode & 0x00FF)

	// set vx to nn
	vm.v[x] = nn

	// continue to next instruction
	vm.pc += 2
}

// _0x7XNN adds nn to vx
func (vm *VM) _0x7XNN() {

	// decode x and nn
	x := (vm.opcode & 0x0F00) >> 8
	nn := byte(vm.opcode & 0x00FF)

	// set vx to nn
	vm.v[x] += nn

	// continue to next instruction
	vm.pc += 2
}

// _0x8XY0 sets vx to the value of vy
func (vm *VM) _0x8XY0() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	y := (vm.opcode & 0x00F0) >> 4

	// set vx to value of vy
	vm.v[x] = vm.v[y]

	// continue to next instruction
	vm.pc += 2
}

// _0x8XY1 sets vx to the value of "vx or vy"
func (vm *VM) _0x8XY1() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	y := (vm.opcode & 0x00F0) >> 4

	// set vx to value of vy
	vm.v[x] |= vm.v[y]

	// continue to next instruction
	vm.pc += 2
}

// _0x8XY2 sets vx to the value of "vx and vy"
func (vm *VM) _0x8XY2() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	y := (vm.opcode & 0x00F0) >> 4

	// set vx to value of vy
	vm.v[x] &= vm.v[y]

	// continue to next instruction
	vm.pc += 2
}

// _0x8XY3 sets vx to the value of "vx xor vy"
func (vm *VM) _0x8XY3() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	y := (vm.opcode & 0x00F0) >> 4

	// set vx to value of vy
	vm.v[x] ^= vm.v[y]

	// continue to next instruction
	vm.pc += 2
}

// _0x8XY4 adds vy to vx. vf is set to 1 when there's a carry
func (vm *VM) _0x8XY4() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	y := (vm.opcode & 0x00F0) >> 4

	// carry bit
	if vm.v[y] > (vm.v[x] - 0xFF) {
		vm.v[0xF] = 1
	} else {
		vm.v[0xF] = 0
	}

	// add vx to vy
	vm.v[x] += vm.v[y]

	// continue to next instruction
	vm.pc += 2
}

// _0x8XY5 vy is subtracted from vx. vf is set to 0 when there's a borrow
func (vm *VM) _0x8XY5() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	y := (vm.opcode & 0x00F0) >> 4

	if vm.v[y] > vm.v[x] {
		vm.v[0xF] = 0 // there is a borrow
	} else {
		vm.v[0xF] = 1
	}

	// substract vy from vx
	vm.v[x] -= vm.v[y]

	// continue to next instruction
	vm.pc += 2
}

// _0x8XY6 shifts vx right by one.
// vf is set to the value of the least significant bit of vx before the shift
func (vm *VM) _0x8XY6() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	// y := (vm.opcode & 0x00F0) >> 4

	// set vf
	vm.v[0xF] = vm.v[x] & 0x1
	// shift vx
	vm.v[x] >>= 1

	// continue to next instruction
	vm.pc += 2
}

// _0x8XY7 sets vx to vy - vx.
// vf is set to 0 when there is a borrow
func (vm *VM) _0x8XY7() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	y := (vm.opcode & 0x00F0) >> 4

	// set vf
	if vm.v[y] > vm.v[x] {
		vm.v[0xF] = 0 // there is a borrow
	} else {
		vm.v[0xF] = 1
	}

	// update vx
	vm.v[x] = vm.v[y] - vm.v[x]

	// continue to next instruction
	vm.pc += 2
}

// _0x8XYE shifts VX left by one.
// vf is set to the value of the most significant bit of vx before the shift
func (vm *VM) _0x8XYE() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	//y := (vm.opcode & 0x00F0) >> 4

	// set vf
	vm.v[0xF] = vm.v[x] >> 7
	// update vx
	vm.v[x] <<= 1

	// continue to next instruction
	vm.pc += 2
}

// _0x9XY0 skips the next instruction if VX doesn't equal VY
func (vm *VM) _0x9XY0() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	y := (vm.opcode & 0x00F0) >> 4

	// skip next instruction
	if vm.v[x] == vm.v[y] {
		vm.pc += 4
		return
	}

	// continue to next instruction
	vm.pc += 2
}

// _0xANNN sets I to the address NNN
func (vm *VM) _0xANNN() {

	// update index register
	vm.i = vm.opcode & 0x0FFF

	// continue to next instruction
	vm.pc += 2
}

// _0xBNNN sets I to the address NNN
func (vm *VM) _0xBNNN() {

	// update the program counter
	vm.pc = (vm.opcode & 0x0FFF) + uint16(vm.v[0])
}

// _0xCXNN sets VX to a random number and NN
func (vm *VM) _0xCXNN() {

	// decode x and nn
	x := (vm.opcode & 0x0F00) >> 8
	nn := byte(vm.opcode & 0x00FF)

	// update vx
	vm.v[x] = byte(rand.Float32()*255) & nn

	// continue to next instruction
	vm.pc += 2
}

// _0xDXYN draws a sprite at coordinate (VX, VY) that has a width of 8 pixels and a height of N pixels.
// Each row of 8 pixels is read as bit-coded starting from memory location I;
// I value doesn't change after the execution of this instruction.
// VF is set to 1 if any screen pixels are flipped from set to unset when the sprite is drawn,
// and to 0 if that doesn't happen
func (vm *VM) _0xDXYN() {
	// TODO: IMPLEMENT
}

// _0xEX9E skips the next instruction if the key stored in VX is pressed
func (vm *VM) _0xEX9E() {

	// decode x
	x := (vm.opcode & 0x0F00) >> 8

	// skip next instruction
	if vm.keypad[vm.v[x]] != 0 {
		vm.pc += 4
		return
	}

	// continue to next instruction
	vm.pc += 2
}

// _0xEXA1 skips the next instruction if the key stored in VX isn't pressed
func (vm *VM) _0xEXA1() {

	// decode x
	x := (vm.opcode & 0x0F00) >> 8

	// skip next instruction
	if vm.keypad[vm.v[x]] == 0 {
		vm.pc += 4
		return
	}

	// continue to next instruction
	vm.pc += 2
}

// _0xFX07 sets VX to the value of the delay timer
func (vm *VM) _0xFX07() {

	// decode x
	x := (vm.opcode & 0x0F00) >> 8

	// set vx
	vm.v[x] = vm.delayTimer

	// continue to next instruction
	vm.pc += 2
}

// _0xFX0A a key press is awaited, and then stored in VX
func (vm *VM) _0xFX0A() {

	// decode x
	x := (vm.opcode & 0x0F00) >> 8

	var keyPress bool

	for i, k := range vm.keypad {
		if k != 0 {
			vm.v[x] = byte(i)
			keyPress = true
		}
	}

	// if we didn't received a keypress, skip this cycle and try again.
	if !keyPress {
		return
	}

	// continue to next instruction
	vm.pc += 2
}

// _0xFX15 sets the delay timer to VX
func (vm *VM) _0xFX15() {

	// decode x
	x := (vm.opcode & 0x0F00) >> 8

	// set delay timer
	vm.delayTimer = vm.v[x]

	// continue to next instruction
	vm.pc += 2
}

// _0xFX18 sets the sound timer to VX
func (vm *VM) _0xFX18() {

	// decode x
	x := (vm.opcode & 0x0F00) >> 8

	// set sound timer
	vm.soundTimer = vm.v[x]

	// continue to next instruction
	vm.pc += 2
}

// _0xFX1E adds VX to I
func (vm *VM) _0xFX1E() {

	// decode x
	x := (vm.opcode & 0x0F00) >> 8

	// VF is set to 1 when range overflow (I+VX>0xFFF)
	if vm.i+uint16(vm.v[x]) > 0xFFF {
		vm.v[0xF] = 1
	} else {
		vm.v[0xF] = 0
	}

	// set i
	vm.i += uint16(vm.v[x])

	// continue to next instruction
	vm.pc += 2
}

// _0xFX29 sets I to the location of the sprite for the character in VX.
// Characters 0-F (in hexadecimal) are represented by a 4x5 font
func (vm *VM) _0xFX29() {

	// decode x
	x := (vm.opcode & 0x0F00) >> 8

	// set i
	vm.i = uint16(vm.v[x]) * 0x5

	// continue to next instruction
	vm.pc += 2
}

// _0xFX33 stores the Binary-coded decimal representation of VX
// at the addresses I, I plus 1, and I plus 2
func (vm *VM) _0xFX33() {

	// decode x
	x := (vm.opcode & 0x0F00) >> 8

	// set i
	vm.mem[vm.i] = vm.v[x] / 100
	vm.mem[vm.i+1] = (vm.v[x] / 10) % 10
	vm.mem[vm.i+2] = (vm.v[x] % 100) % 10

	// continue to next instruction
	vm.pc += 2
}

// _0xFX55 stores V0 to VX in memory starting at address I
func (vm *VM) _0xFX55() {

	// decode x
	x := (vm.opcode & 0x0F00) >> 8

	for i := uint16(0); i <= x; i++ {
		vm.mem[vm.i+i] = vm.v[i]
	}

	// on the original interpreter, when the operation is done, I = I + X + 1.
	vm.i += x + 1

	// continue to next instruction
	vm.pc += 2
}

// _0xFX65 fills V0 to VX with values from memory starting at address I
func (vm *VM) _0xFX65() {

	// decode x
	x := (vm.opcode & 0x0F00) >> 8

	for i := uint16(0); i <= x; i++ {
		vm.v[i] = vm.mem[vm.i+i]
	}

	// on the original interpreter, when the operation is done, I = I + X + 1.
	vm.i += x + 1

	// continue to next instruction
	vm.pc += 2
}
