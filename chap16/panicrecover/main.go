package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			// 此处无法处理 ConfigValue() 引发的panic
		}
	}()

	go ConfigValue("a.config")
	//ApplyConfig(SafeConfigValue()) // 应用配置
}

// SafeConfigValue 尝试从 a.config 读取配置值
// 如果读取失败, 则返回默认值.
func SafeConfigValue() (v string) {
	defer func() {
		if r := recover(); r != nil { // // 发生了 panic
			v = defConfigValue() // 2: 使用默认值
		}
	}()
	return ConfigValue("a.config") // 尝试从配置文件中读取
}

// ConfigValue 从 configFile 中读取一个特定的配置值
// 如果无法打开此文件, 将触发 panic
func ConfigValue(configFile string) string {
	f, err := os.Open(configFile) // 打开配置文件
	if err != nil {
		panic(err) // 无法打开文件, 前置条件不满足
	}
	defer f.Close()
	return readConfigValue(f) // 读取并返回配置值
}

func readConfigValue(*os.File) string {
	return ""
}

func defConfigValue() string { return "def value" }

func ApplyConfig(v string) {
	fmt.Printf("apply %#v\n", v)
}

var maxID uintptr

// NextID 返回下一个全局递增 ID
func NextID() (id uintptr) {
	id = maxID
	if maxID == ^uintptr(0) {
		panic("id overflow")
	}
	maxID++
	return
}
