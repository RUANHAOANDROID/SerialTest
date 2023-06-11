package main

import (
	"SerialTest/utils"
	"bufio"
	"fmt"
	"log"
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
	log.Println("串口初步调试")
	// 打开串口连接
	port, err := serial.OpenPort(&serial.Config{
		Name:        "COM1",      // 请根据实际情况设置COM口名称
		Baud:        9600,        // 波特率
		ReadTimeout: time.Second, // 读取超时时间
	})
	log.Printf("串口=COM1 波特率=9600 读超时时间=%s\n", time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	// 启动一个goroutine来监听串口数据
	go listenSerialData(port)

	// 主线程用于发送命令到串口
	//scanner := bufio.NewScanner(os.Stdin)
	for {
		//inputCMD(scanner, port)
		inputAutoTest(port)
	}
}
func inputAutoTest(port *serial.Port) {
	cmd1 := make([]byte, 8)
	cmd1[0] = 01
	cmd1[1] = 03
	cmd1[2] = 00
	cmd1[3] = 01
	cmd1[4] = 01
	cmd1[5] = 01
	cmd1[6] = 01
	cmd1[7] = uint8(0xCB)
	status, err := port.Write(cmd1)
	fmt.Printf("write %d %v \n", status, err)
}
func inputCMD(scanner *bufio.Scanner, port *serial.Port) {
	fmt.Print("输入命令: ")
	scanner.Scan()
	cmd := scanner.Text()

	writeLen, err := port.Write([]byte(cmd + "\n"))
	if err != nil {
		log.Fatal(err)
		utils.Log.Println(err.Error())
	}
	utils.Log.Println("write->", writeLen, cmd)
}

func listenSerialData(port *serial.Port) {
	buffer := make([]byte, 8) //避免资源浪费阻塞等待数据落入缓冲区
	for {
		n, err := port.Read(buffer)
		if err != nil {
			log.Fatalf("读取数据时出错：%v", err)
		}
		if n > 0 {
			log.Printf("已读取 %d 字节数据: %s\n", n, buffer[:n])
			utils.Log.Println("read <-", fmt.Sprintf("ASCII %c", buffer[:n]))
			utils.Log.Println("read <-", fmt.Sprintf("HEX %x", buffer[:n]))
			port.Flush()
		}
	}
}
