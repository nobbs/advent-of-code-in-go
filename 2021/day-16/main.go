package main

import (
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type Packet struct {
	version      uint64
	typeID       uint64
	lengthTypeID uint64
	literalValue uint64
	subpackets   []Packet
}

func hexToBin(h string) string {
	bytes, _ := hex.DecodeString(h)
	builder := strings.Builder{}

	for _, b := range bytes {
		builder.WriteString(fmt.Sprintf("%08b", b))
	}

	return builder.String()
}

func binToUint64(b string) uint64 {
	val, _ := strconv.ParseUint(b, 2, 64)
	return val
}

func parsePacket(bin string) (*Packet, string) {
	packet := &Packet{
		version: binToUint64(bin[0:3]),
		typeID:  binToUint64(bin[3:6]),
	}
	ptr := 6

	// handle different types
	switch packet.typeID {
	case 4:
		// handle literal values
		blocksize := 5
		value := ""
		for i := 6; true; i += blocksize {
			value += bin[i+1 : i+blocksize]

			if bin[i] == '0' {
				ptr = i + blocksize
				break
			}
		}
		packet.literalValue = binToUint64(value)
	default:
		// every other kind of typeID is an operator, so let's first handle all the subpackets
		// parse the lengthTypeID
		packet.lengthTypeID = binToUint64(string(bin[ptr]))
		ptr++

		switch packet.lengthTypeID {
		case 0:
			totalLength := int(binToUint64(bin[ptr : ptr+15]))
			ptr += 15

			subString := bin[ptr : ptr+totalLength]
			subpackets := []Packet{}
			for len(subString) != 0 {
				child, remainder := parsePacket(subString)
				subpackets = append(subpackets, *child)
				subString = remainder
			}

			packet.subpackets = subpackets
			ptr += totalLength
		case 1:
			numSubpackets := binToUint64(bin[7:18])

			subString := bin[18:]
			subpackets := []Packet{}
			for len(subpackets) != int(numSubpackets) {
				child, remainder := parsePacket(subString)
				subpackets = append(subpackets, *child)
				subString = remainder
			}

			packet.subpackets = subpackets
			bin = subString
			ptr = 0
		}
	}

	// and now let's handle specific operator types
	switch packet.typeID {
	case 0:
		// sum type: sum of the values of their subpackets
		var sum uint64
		for _, c := range packet.subpackets {
			sum += c.literalValue
		}
		packet.literalValue = sum
	case 1:
		// product type: product of the values of their subpackets
		var prod uint64 = 1
		for _, c := range packet.subpackets {
			prod *= c.literalValue
		}
		packet.literalValue = prod
	case 2:
		// minimum type
		var min uint64 = math.MaxUint64
		for _, c := range packet.subpackets {
			if c.literalValue < min {
				min = c.literalValue
			}
		}
		packet.literalValue = min
	case 3:
		// maximum type
		var max uint64 = 0
		for _, c := range packet.subpackets {
			if c.literalValue > max {
				max = c.literalValue
			}
		}
		packet.literalValue = max
	case 5:
		// greater than type: 1 if value of first subpacket > val. of second subpacket, 0 otherwise
		if packet.subpackets[0].literalValue > packet.subpackets[1].literalValue {
			packet.literalValue = 1
		} else {
			packet.literalValue = 0
		}
	case 6:
		// less than type, as before
		if packet.subpackets[0].literalValue < packet.subpackets[1].literalValue {
			packet.literalValue = 1
		} else {
			packet.literalValue = 0
		}
	case 7:
		// equal to type, as before
		if packet.subpackets[0].literalValue == packet.subpackets[1].literalValue {
			packet.literalValue = 1
		} else {
			packet.literalValue = 0
		}
	}

	return packet, bin[ptr:]
}

func (p *Packet) versionSum() uint64 {
	var sum uint64 = p.version

	for _, c := range p.subpackets {
		sum += c.versionSum()
	}

	return sum
}

func partOne(lines []string) int {
	bin := hexToBin(lines[0])
	packet, _ := parsePacket(bin)

	return int(packet.versionSum())
}

func partTwo(lines []string) int {
	bin := hexToBin(lines[0])
	packet, _ := parsePacket(bin)

	return int(packet.literalValue)
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
