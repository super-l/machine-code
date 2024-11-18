package types

type MachineInformation struct {
	UUID              string `json:"uuid"`
	BoardSerialNumber string `json:"boardSerialNumber"` // 主板序列号
	CpuSerialNumber   string `json:"cpuSerialNumber"`   // cpu序列号
	DiskSerialNumber  string `json:"diskSerialNumber"`  // 硬盘序列号
	Mac               string `json:"mac"`               // 本地网卡mac相关信息
}
