package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// part1
	// movePointerLinear()
	// part 2
	// movePointerCircular()
	// part 3
	swapPositionCircular()
}

func swapPositionCircular() {
	nameInput := "Braeorath,Malpyros,Kazphor,Gareldrin,Cragzar,Aeorxelor,Wyrynn,Ballith,Laziradarin,Ascallyr,Vaeldra,Marallyr,Shaelxelor,Fyrzyph,Tirzal,Drazaris,Hyrapyros,Arvynn,Tharnural,Cynvarzorin,Harnimar,Kazdin,Lazgryph,Selthel,Pylarris,Axaladir,Ilmarmir,Brelaxar,Norakvoran,Valquin"
	// nameInput := "Vyrdax,Drakzyph,Fyrryn,Elarzris"
	// stepInput := "R3,L2,R3,L1"
	stepInput := "L36,R39,L9,R47,L34,R24,L27,R22,L44,R28,L18,R18,L8,R43,L26,R12,L30,R48,L29,R10,L5,R45,L5,R13,L5,R28,L5,R24,L5,R39,L5,R7,L5,R36,L5,R8,L5,R28,L5,R42,L32,R48,L23,R11,L48,R7,L26,R31,L46,R25,L12,R21,L8,R32,L20,R14,L10,R35,L15"
	names := strings.Split(nameInput, ",")
	steps := strings.Split(stepInput, ",")

	for _, step := range steps {
		currentIdx := 0
		isLeft := string(step[0]) == "L"
		// simple atoi logic
		num, _ := strconv.Atoi(step[1:])

		if isLeft {
			currentIdx = goLeftCircular(currentIdx, num, len(names)-1)
		} else {
			currentIdx = goRightCircular(currentIdx, num, len(names)-1)
		}

		temp := names[0]
		names[0] = names[currentIdx]
		names[currentIdx] = temp

	}

	fmt.Printf("final part 3: current index: %d, name: %s\n", 0, names[0])
}

func movePointerCircular() {
	nameInput := "Sarnnix,Zormal,Quirzrak,Zyrgyth,Briveldrin,Ascaleon,Skarfeth,Durnpyros,Calasis,Vyrkris,Arakwyris,Shaemdren,Krynnalar,Sillith,Zynardith,Draithnar,Tarlparth,Kazaxar,Ilmarnar,Nylthyris"
	// nameInput := "Vyrdax,Drakzyph,Fyrryn,Elarzris"
	// stepInput := "R3,L2,R3,L1"
	stepInput := "L5,R7,L14,R11,L15,R13,L11,R14,L19,R7,L5,R12,L5,R19,L5,R13,L5,R11,L5,R18,L7,R11,L11,R10,L16,R6,L9,R6,L15"

	names := strings.Split(nameInput, ",")
	steps := strings.Split(stepInput, ",")

	currentIdx := 0
	for _, step := range steps {
		isLeft := string(step[0]) == "L"
		// simple atoi logic
		num, _ := strconv.Atoi(step[1:])

		if isLeft {
			currentIdx = goLeftCircular(currentIdx, num, len(names)-1)
		} else {
			currentIdx = goRightCircular(currentIdx, num, len(names)-1)
		}
	}
}

func movePointerLinear() {
	part1NameInput := "Cyndindor,Erassor,Mornverax,Krynnzor,Gorathindor,Myndthyris,Harnrex,Fyrcion,Naldnix,Sildar"
	part1StepInput := "L1,R3,L2,R6,L1,R5,L1,R5,L2,R7,L7"

	names := strings.Split(part1NameInput, ",")
	steps := strings.Split(part1StepInput, ",")

	currentIdx := 0
	for _, step := range steps {
		isLeft := string(step[0]) == "L"
		// simple atoi logic
		num, _ := strconv.Atoi(step[1:])
		if isLeft {
			currentIdx = goLeftWithCap(currentIdx, num)
		} else {
			currentIdx = goRightWithoutCap(currentIdx, num, len(names)-1)
		}
	}
}

func goLeftWithCap(currentIdx, stepToMove int) int {
	result := currentIdx - stepToMove
	if result < 0 {
		return 0
	}

	return result
}

func goRightWithoutCap(currentIdx, stepToMove, endIdx int) int {
	result := currentIdx + stepToMove
	if result > endIdx {
		return endIdx
	}
	return result
}

func goLeftCircular(currentIdx, stepToMove, endIdx int) int {
	stepToMove = stepToMove % (endIdx + 1)
	result := currentIdx - stepToMove

	if result < 0 {
		return endIdx + 1 + result
	}

	return result
}

func goRightCircular(currentIdx, stepToMove, endIdx int) int {
	stepToMove = stepToMove % (endIdx + 1)
	result := currentIdx + stepToMove
	if result > endIdx {
		return result - endIdx - 1
	}
	return result
}
