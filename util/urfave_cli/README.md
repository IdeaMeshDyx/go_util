# cli 使用的 cmd 接入的插件
https://www.bwangel.me/2019/04/21/go-cli/
https://cli.urfave.org/v1/getting-started/


urfave/cli是一个简单、快速且有趣的包，用于在 Go 中构建命令行应用程序。目标是使开发人员能够以富有表现力的方式编写快速且可分发的命令行应用程序。

# 环境安装

```
$ go get github.com/urfave/cli/v2
```

``` go
import (
  "github.com/urfave/cli/v2" // imports as package "cli"
)
```

ok 了，接下来只需要使用他的架构就可以了

# 使用步骤教程

## 接收传入的参数并存入 c.Args
下面的代码没有添加 command，只有一个默认的动作。
所有命令后跟着的参数都会被放到 c.Args 中
``` go
package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "watch"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Printf("Hello %q!", c.Args().Get(0))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
```


## 添加全局 flag
下面的代码中添加了全局的__flag__: --lang
__flag__的Value既可以通过Destination属性存储到自定义变量中，也可以使用c.String("lang")的方式来获取
使用的指令形式为： ./step2 --lang chinese


``` go
package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	var language string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
			Destination: &language,
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "bwangel"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		// if c.String("lang") == "chinese" {
		if language == "chinese" {
			fmt.Println("你好", name)
		} else {
			fmt.Println("Hello", name)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
```

## Flag 以及在 Flag 接收的 usage 参数
通过 `` 设置 Flag 接收的参数，
解析 Name 时候他写的架子应该是按照 ， 来区分的

``` go
func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "test, t",
			Usage: "上述的test和t会按照 , 解析分开,然后填充反引号里面的内容 `CONTENTIN s`",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
```

输出结果为：
``` shell
[root@VM-243-88-tencentos /data/home/yixinda/learn/package/urfave_cli/step3]# ./step3 --help
NAME:
   step3 - A new cli application

USAGE:
   step3 [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --test CONTENTIN s, -t CONTENTIN s  上述的test和t会按照 , 解析分开,然后填充反引号里面的内容 CONTENTIN s
   --help, -h                          show help
   --version, -v                       print the version
```

## flag 与 command
flag： -v XXX --> 含义是对 v 接收 XXX 的 参数
command： 直接接收的动作参数比如： c XXXX,c 行为接收 XXXX 参数

```go
package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "Language for the greeting",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println(c.Command.Name)
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				fmt.Println(c.Command.Name)
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
```