package main

import (
  "time"
  "github.com/shirou/gopsutil/cpu"
  "github.com/shirou/gopsutil/mem"
  "encoding/json"
  "net/http"
)

type CPUStats struct {
  UsedPercentPerCore []float64 `json:"used_percent_per_core"`
}

type SystemStats struct {
  MemoryStats *mem.VirtualMemoryStat `json:"memory"`
  CPUStats    CPUStats               `json:"cpu"`
}

func main() {
  http.HandleFunc("/stats", stats)
  http.ListenAndServe(":8071", nil)
}

func stats(w http.ResponseWriter, r *http.Request) {
  memory, _ := mem.VirtualMemory()
  c, _      := cpu.CPUPercent(time.Second, true)
  cpu       := CPUStats{UsedPercentPerCore: c}

  stats := SystemStats{
    MemoryStats: memory,
    CPUStats: cpu,
  }

  js, err := json.Marshal(stats)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}
