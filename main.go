package main

import (
	"SerialTest/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tarm/serial"
)

func main() {
	log.Println("串口初步调试")
	// 打开串口连接
	port, err := serial.OpenPort(&serial.Config{
		Name:        "COM2",      // 请根据实际情况设置COM口名称
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
	scanner := bufio.NewScanner(os.Stdin)
	for {
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
}

func listenSerialData(port *serial.Port) {
	buffer := make([]byte, 1024) //避免资源浪费阻塞等待数据落入缓冲区
	for {
		n, err := port.Read(buffer)
		if err != nil {
			log.Fatalf("读取数据时出错：%v", err)
		}
		if n > 0 {
			log.Printf("已读取 %d 字节数据: %s\n", n, buffer[:n])
			utils.Log.Println("read <-", fmt.Sprintf("ASCII %c", buffer[:n]))
			utils.Log.Println("read <-", fmt.Sprintf("HEX %x", buffer[:n]))
		}
	}
}
