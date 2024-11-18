# 简介

machine-code，是忘忧草安全团队出品的一款使用GO语言实现的跨平台机器码(硬件信息)获取程序。

能获取的信息包括系统UUID、主板序列号、CPU序列号、MAC网卡信息、硬盘序列号等。主要用于设备特征识别、商业软件授权CDKEY生成等等。

同时支持windows、Linux、mac等系统，引入本库后，可与您的项目代码无缝集成，并且同时支持跨平台编译。

# 更新


## 2024-11-18  v1.2

1. 优化linux下获取硬盘ID序列码。避免序列码太多导致拼接起来太长，使用MD5算法二次加密。


## 2024-11-18  v1.1 

1. 项目重写与优化,调整项目文件结构；
2. 添加编译时变量，可直接无需调整任何代码，编译生成多个系统的目标文件不会出现编译错误；
3. linux下获取磁盘系列号方案更新；
4. 本着精简，能不用其他库就不用的原则，本项目不再使用其他三方库如wmi。
5. 迁移与合并码友对老版本的改进代码。
 

## 2021-07-20  v1.0 
1. 只是简单的做个分享


# 安装说明

```
go get github.com/super-l/machine-code
```


# 使用方法

具体可查看项目的demo目录。并且提供编译好的demo测试程序，以及运行截图。

使用代码如：

```
package main

import (
	"encoding/json"
	"fmt"
	machine_code "github.com/super-l/machine-code"
	"strings"
)

func main() {
	if machine_code.MachineErr != nil {
		fmt.Println("获取机器码信息错误:" + machine_code.MachineErr.Error())
		return
	}

	machineJson, _ := json.Marshal(machine_code.Machine)
	fmt.Println("机器码信息汇总:" + string(machineJson))

	trafficIp, _ := machine_code.GetIpAddr()
	fmt.Println("当前活跃IP：" + trafficIp)

	allIp, _ := machine_code.GetIpAddrAll()
	fmt.Println("所有IP：" + strings.Join(allIp, " "))
}

```

# 运行截图

![image1](https://github.com/super-l/machine-code/blob/master/demo/run1.png)

![image2](https://github.com/super-l/machine-code/blob/master/demo/run2.png)

# 注意事项

windows下使用的是wmi相关命令来获取系统信息。理论上99.99%成功率。但不排除部分电脑的wmi关闭或损坏等。

linux下主要是使用的基于特定系统文件读取信息。


# 其他

如果能帮到您，如果您觉得不错，请给项目点个赞（star）吧，您的支持是我前进的动力，谢谢！

```html
技术交流QQ群：50246933
技术博客： http://www.xiao6.net
```

感谢 @lazyphp 和 @bestK 的改进提交！