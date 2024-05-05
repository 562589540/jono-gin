package dto

type ServerInfoRes struct {
	RunIngInfo *RunIngInfo `json:"runIngInfo"`
	MemoryInfo *MemoryInfo `json:"memoryInfo"`
	CPUInfo    *CPUInfo    `json:"cpuInfo"`
	DiskInfos  []DiskInfo  `json:"diskInfos"`
}

type RunIngInfo struct {
	Goos         string `json:"goos"`         //操作系统
	Goarch       string `json:"goarch"`       //系统架构
	StartTime    string `json:"startTime"`    //开始运行时间
	RunTime      string `json:"runTime"`      //运行时长
	GoVersion    string `json:"goVersion"`    //go语言版本
	ServerIp     string `json:"serverIp"`     //服务器ip
	ServerName   string `json:"serverName"`   //服务器名称
	NumGoroutine int    `json:"numGoroutine"` //Goroutine数量
}

// MemoryInfo 封装系统内存信息和当前进程的内存使用情况
type MemoryInfo struct {
	Total        string `json:"total"`        // 总内存 GB
	Used         string `json:"used"`         // 已使用内存 GB
	Free         string `json:"free"`         // 剩余内存 GB
	UsagePercent string `json:"usagePercent"` // 内存使用率
	ProcessMB    string `json:"processMB"`    // 当前进程使用内存 MB
}

// CPUInfo 封装 CPU 相关信息
type CPUInfo struct {
	Cores     int    `json:"cores"`
	Usage     string `json:"usage"`     // CPU 使用率
	LoadAvg1  string `json:"loadAvg1"`  // 1 分钟平均负载
	LoadAvg5  string `json:"loadAvg5"`  // 5 分钟平均负载
	LoadAvg15 string `json:"loadAvg15"` // 15 分钟平均负载
}

type DiskInfo struct {
	MountPath    string `json:"mountPath"`    // 挂载路径，显示磁盘挂载在系统中的位置
	FileSystem   string `json:"fileSystem"`   // 文件系统类型，如 ext4、NTFS
	Total        string `json:"total"`        // 磁盘总大小，以字节为单位
	Used         string `json:"used"`         // 已使用的磁盘空间，以字节为单位
	Available    string `json:"available"`    // 可用的磁盘空间，以字节为单位
	UsagePercent string `json:"usagePercent"` // 磁盘使用百分比
}
