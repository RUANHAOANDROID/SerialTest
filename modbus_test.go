package main

import (
	"fmt"
	"testing"
)

func TestModbusPacket_Build(t *testing.T) {
	//03 05 00 01 ff 00 dc 18 -> [3 5 0 1 255 0 220 24]
	cmd1 := make([]byte, 8)
	cmd1[0] = 03
	cmd1[1] = 05
	cmd1[2] = 00
	cmd1[3] = 01
	cmd1[4] = 0xff
	cmd1[5] = 0x00
	packet := ModbusPacket{
		address:  03,
		function: 05,
		starting: 0x0001,
		quantity: 0xff00,
	}
	modbusPacket := packet.Build()
	fmt.Println(modbusPacket)
	fmt.Printf("write len=%d\n", cmd1)
}
