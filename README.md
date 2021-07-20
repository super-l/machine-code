### 一、程序简介

machine-code，是GO语言实现的跨平台机器码(硬件信息)获取程序，包括PlatformUUID、SerialNumber、MAC网卡信息、CPUID信息等。同时支持windows、Linux、mac等系统！


#### 支持的系统：
    windows
    linux
    mac

#### 目前可获取的信息
    PlatformUUID
    SerialNumber
    MAC网卡信息
    CPUID

### 二、安装说明

```
$ go get github.com/super-l/machine-code/machine
```


### 三、实例代码

#### 1：单独获取信息
```
serialNumber, err := machine.GetSerialNumber()
if err != nil {
    fmt.Println(err.Error())
}
fmt.Println("serialNumber = ", serialNumber)

uuid, err := machine.GetPlatformUUID()
if err != nil {
    fmt.Println(err.Error())
}
fmt.Println("uuid = ", uuid)

cpuid, err := machine.GetCpuId()
if err != nil {
    fmt.Println(err.Error())
}
fmt.Println("cpuid = ", cpuid)

macInfo, err := machine.GetMACAddress()
if err != nil {
    fmt.Println(err.Error())
}
fmt.Println("mac = ", macInfo)
```

运行结果：

```
serialNumber =  C02P7676XXXX
uuid =  748869E5-06B6-5855-A0A4-7E19497XXXXX
cpuid =  0x00040000
mac =  34:36:3b:XX:XX:XX

```

#### 2：获取全部信息

实例代码：
```
machineData := machine.GetMachineData()
result, err := json.Marshal(machineData)
if err != nil {
    fmt.Println(err.Error())
}
fmt.Println(string(result))
```

运行结果：

```
{
    "platformUUID": "748869E5-06B6-5855-A0A4-7E19497XXXXX",
    "serialNumber": "C02P767XXXXX",
    "cpuId": "0x00040000",
    "mac": "34:36:3b:XX:XX:XX"
}

```