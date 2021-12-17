package main

import (
	"fmt"
	"strconv"
)

func day16Part1(filePath string) {
	puzzleInput := readFileLines(filePath)
	fmt.Println("Hexadecimal:", puzzleInput[0])
	binaryString := convertHexToBinary(puzzleInput[0])
	fmt.Println("Converted binary:", binaryString)
	parsePackets(binaryString, 0)
}

func parsePackets(binaryString string, versionSum int) {
	// all numbers encoded in any packet are represented as binary with the most significant bit first
	convertToNum := map[string]int{
		"000": 0,
		"001": 1,
		"010": 2,
		"011": 3,
		"100": 4,
		"101": 5,
		"110": 6,
		"111": 7,
	}
	// no header or length to parse
	if len(binaryString) < 8 {
		fmt.Println("Current ending version total:", versionSum)
		return
	}
	version := binaryString[:3]
	versionNum := convertToNum[version]
	fmt.Println(version, "version converts to =>", versionNum)
	versionSum += versionNum
	binaryString = binaryString[3:]
	typeID := binaryString[:3]
	typeIDNum := convertToNum[typeID]
	fmt.Println(typeID, "type ID converts to:", typeIDNum)
	binaryString = binaryString[3:]
	if typeIDNum == 4 {
		// packet is literal value
		var litValue string
		currBits := binaryString[:5]
		prefix := string(currBits[0])
		binaryString = binaryString[5:]
		if prefix == "1" {
			for {
				litValue = litValue + currBits[1:]
				currBits = binaryString[:5]
				prefix = string(currBits[0])
				binaryString = binaryString[5:]
				if prefix == "0" {
					break
				}
			}
		}
		litValue = litValue + currBits[1:]
		fmt.Println(litValue)
		resultLiteral, _ := strconv.ParseInt(litValue, 2, 64)
		fmt.Println("Resulting literal:", resultLiteral)
		parsePackets(binaryString, versionSum)
	} else {
		// packet is an operator
		lengthTypeID := string(binaryString[0])
		binaryString = binaryString[1:]
		if lengthTypeID == "0" {
			totalLengthBits := binaryString[:15]
			binaryString = binaryString[15:]
			fmt.Println("Total length in bits:", totalLengthBits)
			lenSubPackets, _ := strconv.ParseInt(totalLengthBits, 2, 64)
			fmt.Println("Length of sub-packets:", lenSubPackets)
			parsePackets(binaryString, versionSum)
		} else {
			numSubPackets := binaryString[:11]
			fmt.Println("Number of sub-packets binary:", numSubPackets)
			binaryString = binaryString[11:]
			subPackets, _ := strconv.ParseInt(numSubPackets, 2, 64)
			fmt.Println("Number of sub-packets:", subPackets)
			parsePackets(binaryString, versionSum)
		}
	}
}

func convertHexToBinary(hex string) string {
	var binaryString string
	conversions := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}
	for _, c := range hex {
		converted := conversions[string(c)]
		binaryString = binaryString + converted
	}
	return binaryString
}