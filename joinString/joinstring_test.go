package joinString

import (
	"fmt"
	"os"
	"testing"
)

// TestMain会在下面所有测试方法执行开始前先执行，一般用于初始化资源和执行完后释放资源
func TestMain(m *testing.M) {
	fmt.Println("初始化资源")
	result := m.Run() //运行go的测试，相当于调用main方法
	fmt.Println("释放资源")
	os.Exit(result) //退出程序
}

func TestMergerAlternately(t *testing.T) {
	result := mergerAlternately("this", "word")
	if result == "twhoirsd" {
		t.Log("successful")
	} else {
		t.Error("failed")
	}
}
