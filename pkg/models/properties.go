package models

import "time"

type (
	Properties struct {
		IsMachineOn        bool `json:"isMachineOn"`
		TelemetryFrequency int  `json:"telemetryFrequency"`
		ShiftDurationHours int  `json:"shiftDurationHours"`
		BatchDurationHours int  `json:"batchDurationHours"`
	}

	ReportedProperties struct {
		HostName  string    `json:"hostName"`
		IPAddress string    `json:"ipAddress"`
		HostTime  time.Time `json:"hostTime"`
	}

	BoltMachineTelemetryMessage struct {
		PlantName             string    `json:"plantName"`
		ProductionLine        string    `json:"productionLine"`
		ShiftNumber           int       `json:"shiftNumber"`
		BatchNumber           int       `json:"batchNumber"`
		MessageTimestamp      time.Time `json:"messageTimestamp"`
		TotalPartsMade        int       `json:"totalPartsMade"`
		DefectivePartsMade    int       `json:"defectivePartsMade"`
		MachineHealth         string    `json:"machineHealth"`
		OilLevel              float64   `json:"oilLevel"`
		Temperature           float64   `json:"temperature"`
		CpuLoad               float64   `json:"cpuLoad"`
		SystemDiskUsedPercent int       `json:"systemDiskUsedPercent"`
		SystemDiskFreePercent int       `json:"systemDiskFreePercent"`
		MemoryUsed            int64     `json:"memoryUsed"`
		MemoryFree            int64     `json:"memoryFree"`
	}

	BoltMachine struct {
		PlantName             string  `json:"plantName"`
		ProductionLine        string  `json:"productionLine"`
		ShiftNumber           int     `json:"shiftNumber"`
		BatchNumber           int     `json:"batchNumber"`
		TotalPartsMade        int     `json:"totalPartsMade"`
		DefectivePartsMade    int     `json:"defectivePartsMade"`
		MachineHealth         string  `json:"machineHealth"`
		OilLevel              float64 `json:"oilLevel"`
		Temperature           float64 `json:"temperature"`
		CpuLoad               float64 `json:"cpuLoad"`
		SystemDiskUsedPercent int     `json:"systemDiskUsedPercent"`
		SystemDiskFreePercent int     `json:"systemDiskFreePercent"`
		MemoryUsed            int64   `json:"memoryUsed"`
		MemoryFree            int64   `json:"memoryFree"`
		Format                string  `json:"format"`
	}
)
