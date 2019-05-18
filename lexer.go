package bf

import (
	"fmt"
	"os"
	"regexp"
)

// state is bf lexer
type state struct {
	cs            *commands
	inputStr      string
	inputCommands []string
	memory        []byte
	pointer       int
}

var r *regexp.Regexp

// InitState is setting state
func InitState(inputSrc string, c commands, maxCommandNum int) *state {
	state := new(state)
	state.cs = &c
	state.inputStr = inputSrc
	state.memory = make([]byte, maxCommandNum)
	state.pointer = 0
	state.parse()
	return state
}

// parse is inputing commands with regexp
func (s *state) parse() {
	r = regexp.MustCompile(`(` + s.cs.Ops() + `)`)
	s.inputCommands = r.FindAllString(s.inputStr, -1)
}

// Start is excuting bf file
func (s *state) Start() {
	turn := 0
	for turn < len(s.inputCommands) {
		switch s.inputCommands[turn] {
		case s.cs.NEXT:
			s.pointer++
		case s.cs.PREV:
			s.pointer--
		case s.cs.INC:
			s.memory[s.pointer]++
		case s.cs.DEC:
			s.memory[s.pointer]--
		case s.cs.WRITE:
			fmt.Print(string(s.memory[s.pointer]))
		case s.cs.READ:
			r := make([]byte, 1)
			os.Stdin.Read(r)
			s.memory[s.pointer] = r[0]
		case s.cs.OPEN:
			if s.memory[s.pointer] == 0 {
				for depth := 1; depth > 0; {
					turn++
					nCommand := s.inputCommands[turn]
					if nCommand == s.cs.OPEN {
						depth++
					}
					if nCommand == s.cs.CLOSE {
						depth--
					}
				}
			}
		case s.cs.CLOSE:
			for depth := 1; depth > 0; {
				turn--
				nCommand := s.inputCommands[turn]
				if nCommand == s.cs.OPEN {
					depth--
				}
				if nCommand == s.cs.CLOSE {
					depth++
				}
			}
			turn--
		}
		turn++
	}
}
