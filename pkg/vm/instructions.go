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
func (vm *VM) _0x1000() {
	vm.pc = vm.opcode & 0x0FFF
}

// _0x2NNN calls subroutine at nnn
func (vm *VM) _0x2000() {
	vm.stack[vm.sp] = vm.pc    // store current address in stack
	vm.sp++                    // increment the stack pointer
	vm.pc = vm.opcode & 0x0FFF // set the pc to address at nnn
}

// _0x3XNN skips the next instruction if vx equals NN
func (vm *VM) _0x3000() {

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
func (vm *VM) _0x4000() {

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
func (vm *VM) _0x5000() {

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
func (vm *VM) _0x6000() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	nn := vm.opcode & 0x00FF

	// set vx to nn
	vm.v[x] = nn

	// continue to next instruction
	vm.pc += 2
}

// _0x7XNN adds nn to vx
func (vm *VM) _0x7000() {

	// decode x and y
	x := (vm.opcode & 0x0F00) >> 8
	nn := vm.opcode & 0x00FF

	// set vx to nn
	vm.v[x] += nn

	// continue to next instruction
	vm.pc += 2
}
