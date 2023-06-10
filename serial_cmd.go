package main

// AssembleCmd 组装数据作为命令
func AssembleCmd(data string) []byte {
	// 进行命令的组装操作，例如添加起始标志、校验位等
	// 这里只是简单示例，将数据转换为字节数组并添加换行符
	command := []byte(data + "\n")
	return command
}

// PackCmd 组装数据作为命令
func PackCmd(data string) []byte {
	// 进行命令的组装操作，例如添加起始标志、校验位等
	// 这里只是简单示例，将数据转换为字节数组并添加换行符
	command := []byte(data + "\n")
	return command
}
