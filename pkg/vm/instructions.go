package vm

// _0xOOE0 clears the display
// TODO: Update
func (vm *VM) _0x00E0() {

}

// _0x00EE returns from subroutine
func (vm *VM) _0xOOEE() {
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
	nn := vm.opcode & 0x00FF

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
	nn := vm.opcode & 0x00FF

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
	nn := vm.opcode & 0x00FF

	// set vx to nn
	vm.v[x] = nn

	// continue to next instruction
	vm.pc += 2
}

// _0x7XNN adds nn to vx
func (vm *VM) _0x7XNN() {

	// decode x and nn
	x := (vm.opcode & 0x0F00) >> 8
	nn := vm.opcode & 0x00FF

	// set vx to nn
	vm.v[x] += nn

	// continue to next instruction
	vm.pc += 2
}

// 0x8XY0 sets vx to the value of vy
func (vm *VM) _0x8XY0() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	y := (vm.opcode & 0x00F0) >> 4

	// set vx to value of vy
	vm.v[x] = vm.v[y]

	// continue to next instruction
	vm.pc += 2
}

// 0x8XY1 sets vx to the value of "vx or vy"
func (vm *VM) _0x8XY1() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	y := (vm.opcode & 0x00F0) >> 4

	// set vx to value of vy
	vm.v[x] |= vm.v[y]

	// continue to next instruction
	vm.pc += 2
}

// 0x8XY2 sets vx to the value of "vx and vy"
func (vm *VM) _0x8XY2() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	y := (vm.opcode & 0x00F0) >> 4

	// set vx to value of vy
	vm.v[x] &= vm.v[y]

	// continue to next instruction
	vm.pc += 2
}

// 0x8XY3 sets vx to the value of "vx xor vy"
func (vm *VM) _0x8XY3() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	y := (vm.opcode & 0x00F0) >> 4

	// set vx to value of vy
	vm.v[x] ^= vm.v[y]

	// continue to next instruction
	vm.pc += 2
}

// 0x8XY4 adds vy to vx. vf is set to 1 when there's a carry
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

// 0x8XY5 vy is subtracted from vx. vf is set to 0 when there's a borrow
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

// 0x8XY6 shifts vx right by one.
// vf is set to the value of the least significant bit of vx before the shift
func (vm *VM) _0x8XY6() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	// y := (vm.opcode & 0x00F0) >> 4

	// set vf
	vm.v[0xF] = x & 0x1
	// shift vx
	vm.v[x] >>= 1

	// continue to next instruction
	vm.pc += 2
}

// 0x8XY7 sets vx to vy - vx.
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

// 0x8XYE shifts VX left by one.
// vf is set to the value of the most significant bit of vx before the shift
func (vm *VM) _0x8XYE() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	//y := (vm.opcode & 0x00F0) >> 4

	// set vf
	vm.v[0xF] = x >> 7
	// update vx
	vm.v[x] <<= 1

	// continue to next instruction
	vm.pc += 2
}

// 0x9XY0 skips the next instruction if VX doesn't equal VY
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

// 0xANNN sets I to the address NNN
func (vm *VM) _0xANNN() {

	// update index register
	vm.i = vm.opcode & 0x0FFF

	// continue to next instruction
	vm.pc += 2
}
