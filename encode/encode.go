package encode

import "encoding/binary"

// GenerateFrame 生成通信帧
func GenerateFrame(slaveAddress, functionCode, startAddressHigh, startAddressLow, coilCountHigh, coilCountLow byte) []byte {
	// 计算CRC16校验
	crc := calculateCRC([]byte{slaveAddress, functionCode, startAddressHigh, startAddressLow, coilCountHigh, coilCountLow})

	// 构建通信帧
	frame := []byte{slaveAddress, functionCode, startAddressHigh, startAddressLow, coilCountHigh, coilCountLow, crc[1], crc[0]}

	return frame
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
