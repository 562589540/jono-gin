package server_info

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"os"
	"runtime"
	"time"
)

var serverInfoService service.IServerInfoService

type Service struct{}

func New() service.IServerInfoService {
	if serverInfoService == nil {
		serverInfoService = &Service{}
	}
	return serverInfoService
}

func (s *Service) ServerInfo() (*dto.ServerInfoRes, error) {
	memoryInfo, err := s.getMemoryInfo()
	if err != nil {
		return nil, err
	}
	cpuInfo, err := s.getCPUInfo()
	if err != nil {
		return nil, err
	}

	diskInfos, err := s.getDiskUsage()
	if err != nil {
		return nil, err
	}
	return &dto.ServerInfoRes{
		RunIngInfo: s.getRunIngInfo(),
		MemoryInfo: memoryInfo,
		CPUInfo:    cpuInfo,
		DiskInfos:  diskInfos,
	}, nil
}

func (s *Service) getRunIngInfo() *dto.RunIngInfo {
	ip, err := gutils.GetLocalIP()
	if err != nil {
		ip = "127.0.0.1"
	}
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "未知"
	}
	return &dto.RunIngInfo{
		Goos:         runtime.GOOS,
		Goarch:       runtime.GOARCH,
		StartTime:    ghub.StartTime.Format("2006-01-02 15:04:05"),
		RunTime:      gutils.FormatDuration(time.Since(ghub.StartTime)),
		GoVersion:    runtime.Version(),
		ServerIp:     ip,
		ServerName:   hostname,
		NumGoroutine: runtime.NumGoroutine(),
	}
}

// 获取系统内存信息和当前进程的内存使用情况
func (s *Service) getMemoryInfo() (*dto.MemoryInfo, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	//获取当前程序
	p, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		return nil, err
	}
	memInfo, err := p.MemoryInfo()
	if err != nil {
		return nil, err
	}

	totalGB := float64(vmStat.Total) / (1024 * 1024 * 1024)            // 转换为 GB
	usedGB := float64(vmStat.Used) / (1024 * 1024 * 1024)              // 转换为 GB
	freeGB := float64(vmStat.Available) / (1024 * 1024 * 1024)         // 转换为 GB
	usagePercent := float64(vmStat.Used) / float64(vmStat.Total) * 100 // 计算使用率

	return &dto.MemoryInfo{
		Total:        fmt.Sprintf("%.2f GB", totalGB),
		Used:         fmt.Sprintf("%.2f GB", usedGB),
		Free:         fmt.Sprintf("%.2f GB", freeGB),
		UsagePercent: fmt.Sprintf("%.2f", usagePercent),
		ProcessMB:    fmt.Sprintf("%.2f MB", float64(memInfo.RSS)/(1024*1024)),
	}, nil
}

// getCPUInfo 获取系统 CPU 信息
func (s *Service) getCPUInfo() (*dto.CPUInfo, error) {
	cores, err := cpu.Counts(true)
	if err != nil {
		return nil, err
	}

	// 获取 CPU 使用率
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}

	// 获取系统平均负载
	avg, err := load.Avg()
	if err != nil {
		return nil, err
	}

	return &dto.CPUInfo{
		Cores:     cores,
		Usage:     fmt.Sprintf("%.1f", percentages[0]), // 取单个 CPU 使用率
		LoadAvg1:  fmt.Sprintf("%.2f", avg.Load1),
		LoadAvg5:  fmt.Sprintf("%.2f", avg.Load5),
		LoadAvg15: fmt.Sprintf("%.2f", avg.Load15),
	}, nil
}

// getDiskUsage 获取所有磁盘的使用信息，并以 GB 为单位返回
func (s *Service) getDiskUsage() ([]dto.DiskInfo, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}

	var diskInfos []dto.DiskInfo
	for _, p := range partitions {
		usageStat, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue // 或处理错误
		}
		diskInfo := dto.DiskInfo{
			MountPath:    p.Mountpoint,
			FileSystem:   p.Fstype,
			Total:        fmt.Sprintf("%.2f GB", gutils.BytesToGB(usageStat.Total)), // 总大小转换为 GB
			Used:         fmt.Sprintf("%.2f GB", gutils.BytesToGB(usageStat.Used)),  // 已使用大小转换为 GB
			Available:    fmt.Sprintf("%.2f GB", gutils.BytesToGB(usageStat.Free)),  // 可用大小转换为 GB
			UsagePercent: fmt.Sprintf("%.2f%%", usageStat.UsedPercent),              // 使用率保持不变
		}
		diskInfos = append(diskInfos, diskInfo)
	}
	return diskInfos, nil
}
