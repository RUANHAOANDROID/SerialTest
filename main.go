package main

import (
	"SerialTest/utils"
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tarm/serial"
)

func pack() {
	originalPacket := ModbusPacket{
		address:     0xFF,
		function:    0x03,
		starting:    0x00A0,
		quantity:    0x0001,
		transaction: 0x0000,
	}

	data := originalPacket.Build()
	fmt.Printf("Original Packet: %X\n", data)

	receivedPacket := ParsePacket(data)
	fmt.Printf("Received Packet:\nAddress: 0x%X\nFunction: 0x%X\nStarting: 0x%X\nQuantity: 0x%X\nTransaction: 0x%X\n",
		receivedPacket.address, receivedPacket.function, receivedPacket.starting,
		receivedPacket.quantity, receivedPacket.transaction)
}
func main() {
	//pack()
	log.Println("RS485-modbus-variant tools")
	log.Println("Powered by hao88.cloud")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("输入COM口名称: ")
	scanner.Scan()
	inputCom := scanner.Text()
	inputCom = strings.ReplaceAll(inputCom, " ", "")

	fmt.Print("输入波特率: ")
	scanner.Scan()
	inputBaud := scanner.Text()
	inputBaud = strings.ReplaceAll(inputBaud, " ", "")
	baud, err := strconv.Atoi(inputBaud)
	if err != nil {
		fmt.Println("转换失败:", baud)
	} else {
		fmt.Println("转换结果:", baud)
	}
	// 打开串口连接
	port, err := serial.OpenPort(&serial.Config{
		Name:        inputCom,    // 请根据实际情况设置COM口名称
		Baud:        baud,        // 波特率
		ReadTimeout: time.Second, // 读取超时时间
	})
	log.Printf("串口=%s 波特率=%s 读超时时间=%s\n",
		inputCom, inputBaud, time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	// 主线程用于发送命令到串口

	for {
		inputCMD(scanner, port)
		//inputAutoTest(port)
	}
}
func inputAutoTest(port *serial.Port) {
	//01 03 00 00 00 0A C5 CD
	cmd1 := make([]byte, 8)
	cmd1[0] = 01
	cmd1[1] = 03
	cmd1[2] = 00
	cmd1[3] = 00
	cmd1[4] = 00
	cmd1[5] = 0x0A
	cmd1[6] = 0xC5
	cmd1[7] = 0xCD
	status, err := port.Write(cmd1)
	fmt.Printf("write len=%d err=%v \n", status, err)
	readData(port)
}
func inputCMD(scanner *bufio.Scanner, port *serial.Port) {
	fmt.Print("输入命令: ")
	scanner.Scan()
	input := scanner.Text()
	input = strings.ReplaceAll(input, " ", "")
	cmd, err := hex.DecodeString(input)
	utils.Log.Println("input cmd->", input)
	writeLen, err := port.Write(cmd)
	if err != nil {
		log.Fatal(err)
		utils.Log.Println(err.Error())
	}
	utils.Log.Println("write->", writeLen, cmd)
	time.Sleep(time.Second)
	readData(port)
}

func readData(port *serial.Port) {
	buffer := make([]byte, 1024) //避免资源浪费阻塞等待数据落入缓冲区
	n, err := port.Read(buffer)
	if err != nil {
		log.Fatalf("读取数据时出错：%v", err)
	}
	data := buffer[:n] // 实际读取到的数据
	if n > 0 {
		utils.Log.Println("read(byte) <-", data)
		hexString := hexSp16(data)
		utils.Log.Println("byte to hex:", hexString)
	}
}

// hexSp16 转换为带空格的16进制
func hexSp16(data []byte) string {
	hexString := ""
	for _, b := range data {
		hexString += fmt.Sprintf("%02X ", b)
	}
	return hexString
}
