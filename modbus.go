package main

import (
	"encoding/binary"
)

// ModbusPacket
// -----------寻址广播---------------
// 从站地址	功能码	起始(高)	起始(低)	数量(高)	数量(低)	CRC16 校验(低)	CRC16 校验(高)
// 0xFF		0x03	0x00	0xA0	0x00	0x01	0x91			0xF6
// 0xFF 发送给所有设备。
// 0x03 读取寄存器的值。
// 起始地址为 0xA0，高字节为 0x00，低字节为 0xA0，表示要读取的寄存器的起始地址。
// 寄存器数量为 0x01，高字节为 0x00，低字节为 0x01，表示要读取的寄存器数量。
// CRC16 校验值为 0x91 0xF6。
// --------命令回应-------------------
// 从站地址	功能码	字节计数	字节1	字节2(uid高位)	字节3(uid低位)	CRC16 校验(低)	CRC16 校验(高)
// 0x01		0x03	0x00	0xA0	0x00			0x01			0x84			0x28
// 设备地址为 0x01，表示回应的设备地址。
// 功能码为 0x03，表示回应的功能码与请求相同，表明是对读取寄存器的请求的回应。
// 起始地址为 0xA0，高字节为 0x00，低字节为 0xA0，与请求相同，表示回应的寄存器的起始地址。
// 寄存器数量为 0x01，高字节为 0x00，低字节为 0x01，与请求相同，表示回应的寄存器数量。
// 数据为 0x84，表示读取到的寄存器的值。
// CRC16 校验值为 0x28 0x84。
// ---------------------------------
type ModbusPacket struct {
	address     byte   //地址
	function    byte   //功能码
	starting    uint16 //[起始高]-[起始低]
	quantity    uint16 //[数量高]-[数量低]
	transaction uint16 //[校验位低]-[校验位高]
}

func (p *ModbusPacket) Build() []byte {
	packet := make([]byte, 8)
	packet[0] = p.address
	packet[1] = p.function
	binary.BigEndian.PutUint16(packet[2:4], p.starting)
	binary.BigEndian.PutUint16(packet[4:6], p.quantity)
	crc := calculateCRC(packet[:6])
	packet[6] = crc[0]
	packet[7] = crc[1]
	return packet
}

// 计算CRC16校验
func calculateCRC(data []byte) []byte {
	crc := uint16(0xFFFF)
	polynomial := uint16(0xA001)

	for _, b := range data {
		crc ^= uint16(b)

		for i := 0; i < 8; i++ {
			lsb := crc & 0x0001
			crc >>= 1

			if lsb == 1 {
				crc ^= polynomial
			}
		}
	}

	result := make([]byte, 2)
	binary.LittleEndian.PutUint16(result, crc)

	return result
}

func ParsePacket(data []byte) *ModbusPacket {
	packet := &ModbusPacket{
		address:     data[0],
		function:    data[1],
		starting:    binary.BigEndian.Uint16(data[2:4]),
		quantity:    binary.BigEndian.Uint16(data[4:6]),
		transaction: binary.BigEndian.Uint16(data[6:8]),
	}

	return packet
}
