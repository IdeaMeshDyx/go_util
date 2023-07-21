package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	// 创建 CLI 结构
	// 并写入对应的参数
	// 这个程序会接受可执行文件后面输入的文字并把他重新输出出来
	app := cli.NewApp()
	app.Name = "watch"
	app.Usage = "fight the loneliness!"
	app.Version = "v1.0"
	app.Action = func(c *cli.Context) error {
		// 程序后面的字符串会被全部放入到 c.Args当中, c.Args().Get(1)表示获取字符串第二个字符
		fmt.Printf("Hello %q!", c.Args().Get(1))
		return nil
	}

	// 启动进程
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
