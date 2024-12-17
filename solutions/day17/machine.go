package main

import (
	"fmt"
	"log"
)

type machine struct {
	a     int64
	b     int64
	c     int64
	seq   []int64
	out   []int64
	pos   int
	count int
}

func (m *machine) run(maxItt int) {
	for m.pos < len(m.seq) && (maxItt <= 0 || m.count < maxItt) {
		m.next()
		m.count++
	}
}

func (m *machine) next() {
	switch m.seq[m.pos] {
	// Division operators
	case 0:
		// log.Println("ADV")
		m.a = m.a / (1 << m.comboOperand())
		m.incrementPos(2)
	case 6:
		// log.Println("BDV")
		m.b = m.a / (1 << m.comboOperand())
		m.incrementPos(2)
	case 7:
		// log.Println("CDV")
		m.c = m.a / (1 << m.comboOperand())
		m.incrementPos(2)

	// Bitwise operators
	case 1:
		// log.Println("BXL")
		m.b = m.b ^ m.literalOperand()
		m.incrementPos(2)
	case 4:
		// log.Println("BXC")
		m.b = m.b ^ m.c
		m.incrementPos(2)

	// Modulo operators
	case 5:
		// log.Println("OUT")
		m.out = append(m.out, m.comboOperand()%8)
		m.incrementPos(2)
	case 2:
		// log.Println("BST")
		m.b = m.comboOperand() % 8
		m.incrementPos(2)

	// Other
	case 3:
		// log.Println("JNZ")
		if m.a == 0 {
			m.incrementPos(2)
		} else {
			m.pos = int(m.seq[m.pos+1])
		}
	}
}

func (m *machine) incrementPos(n int) {
	m.pos += n
}

func (m *machine) comboOperand() int64 {
	raw := m.literalOperand()
	switch {
	case raw < 0:
		panic("operand can not be < 0")
	case raw <= 3:
		return raw
	case raw == 4:
		return m.a
	case raw == 5:
		return m.b
	case raw == 6:
		return m.c
	default:
		log.Println(m)
		panic("failed to get operand (can not be >= 7)")
	}
}

func (m *machine) literalOperand() int64 {
	return m.seq[m.pos+1]
}

func (m *machine) strOut() string {
	if len(m.out) == 0 {
		return ""
	}

	s := fmt.Sprintf("%d", m.out[0])
	for _, o := range m.out[1:] {
		s += fmt.Sprintf(",%d", o)
	}
	return s
}

func (m *machine) String() string {
	curr := "out of bounds"
	if m.pos > 0 && m.pos < len(m.seq) {
		curr = fmt.Sprintf("%d", m.seq[m.pos])
	}
	return fmt.Sprintf("machine{a:%d, b:%d, c:%d, seq:%v, out:%v, pos:%d (%s)}", m.a, m.b, m.c, m.seq, m.out, m.pos, curr)
}
