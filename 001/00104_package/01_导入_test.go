package sample

/*
	Go 语言中的包和其他语言的库或模块的概念类似，目的都是为了支持模块化、封装、单独编译和代码重用。

	一个包的源代码保存在一个或多个以.go 为文件后缀名的源文件中。

	包的声明在目录内的 .go 文件首行 `如上 package sample`，通常包名与包路径的最后一个字段相同，包内每个 .go 文件声明的包名必须一致，`e.g runtime 包`
*/

import (
	"fmt"
	"testing"

	// 导入的包必须在 gopath中，可使用命令 `go get package` 来下载包
	// $GOPATH/pkg/mod/gitlab.mvalley.com/surefive1/gardener/golang/sample/constants
	"gitlab.mvalley.com/surefive1/gardener/golang/sample/constants"
	"gitlab.mvalley.com/surefive1/gardener/golang/sample/variables"

	// 路径最后名称相同的两个包，可以使用别名来区分
	logger "gitlab.mvalley.com/surefive1/gardener/golang/sample/log"
	"log"
)

func TestImport(t *testing.T) {
	// 导入成功后，可以通过包的最后路径命或者别名来引用其中的变量 & 函数
	log.Println("[1]", constants.Global)
	logger.Logger.Println("[2]", constants.Global)

	// 每个包都对应一个独立的命名空间, 虽然都叫 Global, 但是是属于不同包的，所以互相独立
	log.Println("[3]", variables.Global)
	log.Println("[4]", constants.Global)

	// 通过变量名首字母大小写来区分变量能否被其他包引用
	fmt.Println("[5]", constants.privateName) // 这里的代码有什么问题？
	fmt.Println("[6]", constants.PublicName)
}

/*
课堂实践：
- 创建一个使用 go modules 的项目

  `mkdir gardener`

  `cd gardener/`

  `go mod init gitlab.mvalley.com/your/name/gardener` ，问：go.mod 文件发生了什么变化？

- 新建一个 main.go
  ```go
  package main

  import (
      "github.com/sirupsen/logrus"
      "gitlab.mvalley.com/surefive1/gardener/golang/sample/constants"
  )

  func main() {
      logger := logrus.New()
      logger.Println(constants.Global)
  }
  ```

- 运行命令 `go mod tidy`， 问：go.mod 文件发生了什么变化？为什么能引入新的包？引入的包符合什么规则？
  1. Semver 语义化版本递增规则(https://semver.org/lang/zh-CN/)

  2. Semver 版本校验比较工具：[Go Version](https://github.com/hashicorp/go-version)
*/
