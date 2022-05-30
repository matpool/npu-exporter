# Huawei Atlas NPU Prometheus Exporter

This exporter used the npu-smi to query information
about the installed Huawei NPUs.

## Requirements

The npu-smi need to be executable. When running
in a container it must be either baked in or mounted from the host.

## Example

```
0# TYPE atlas_ai_core gauge
atlas_ai_core{deviceid="0"} 0
atlas_ai_core{deviceid="1"} 0
atlas_ai_core{deviceid="2"} 0
atlas_ai_core{deviceid="3"} 0
atlas_ai_core{deviceid="4"} 0
atlas_ai_core{deviceid="5"} 0
atlas_ai_core{deviceid="6"} 0
atlas_ai_core{deviceid="7"} 0
atlas_ai_core{deviceid="8"} 0
atlas_ai_core{deviceid="9"} 0
0# HELP npu_device_count Count of found NPU devices
# TYPE npu_device_count gauge
-npu_device_count 16
-:# HELP atlas_driver_info NPU Info
-# TYPE atlas_driver_info gauge
-atlas_driver_info{version="21.0.3.1"} 1
:# HELP atlas_info Info as reported by the device
-# TYPE atlas_info gauge
-atlas_info{chipid="0",deviceid="0",name="310",npuid="1"} 1
-atlas_info{chipid="0",deviceid="4",name="310",npuid="3"} 1
:atlas_info{chipid="0",deviceid="8",name="310",npuid="4"} 1
--atlas_info{chipid="1",deviceid="1",name="310",npuid="1"} 1
- atlas_info{chipid="1",deviceid="5",name="310",npuid="3"} 1
--atlas_info{chipid="1",deviceid="9",name="310",npuid="4"} 1
-atlas_info{chipid="2",deviceid="2",name="310",npuid="1"} 1
:atlas_info{chipid="2",deviceid="6",name="310",npuid="3"} 1
atlas_info{chipid="3",deviceid="3",name="310",npuid="1"} 1
12atlas_info{chipid="3",deviceid="7",name="310",npuid="3"} 1
89# HELP atlas_memory_total Total memory as reported by the device
0# TYPE atlas_memory_total gauge

atlas_memory_total{deviceid="0"} 8.589934592e+09
atlas_memory_total{deviceid="1"} 8.589934592e+09
atlas_memory_total{deviceid="2"} 8.589934592e+09
atlas_memory_total{deviceid="3"} 8.589934592e+09
atlas_memory_total{deviceid="4"} 8.589934592e+09
atlas_memory_total{deviceid="5"} 8.589934592e+09
atlas_memory_total{deviceid="6"} 8.589934592e+09
atlas_memory_total{deviceid="7"} 8.589934592e+09
atlas_memory_total{deviceid="8"} 8.589934592e+09
atlas_memory_total{deviceid="9"} 8.589934592e+09
# HELP atlas_memory_used Used memory as reported by the device
# TYPE atlas_memory_used gauge
atlas_memory_used{deviceid="0"} 2.748317696e+09
atlas_memory_used{deviceid="1"} 2.834300928e+09
atlas_memory_used{deviceid="2"} 2.748317696e+09
atlas_memory_used{deviceid="3"} 2.748317696e+09
atlas_memory_used{deviceid="4"} 2.748317696e+09
atlas_memory_used{deviceid="5"} 2.748317696e+09
atlas_memory_used{deviceid="6"} 2.748317696e+09
atlas_memory_used{deviceid="7"} 2.748317696e+09
atlas_memory_used{deviceid="8"} 2.748317696e+09
atlas_memory_used{deviceid="9"} 2.748317696e+09
# HELP atlas_power_usage Power usage as reported by the device
# TYPE atlas_power_usage gauge
atlas_power_usage{deviceid="0"} 12.8
atlas_power_usage{deviceid="1"} 12.8
atlas_power_usage{deviceid="2"} 12.8
atlas_power_usage{deviceid="3"} 12.8
atlas_power_usage{deviceid="4"} 12.8
atlas_power_usage{deviceid="5"} 12.8
atlas_power_usage{deviceid="6"} 12.8
atlas_power_usage{deviceid="7"} 12.8
atlas_power_usage{deviceid="8"} 12.8
atlas_power_usage{deviceid="9"} 12.8
# HELP atlas_temperatures Temperature as reported by the device
# TYPE atlas_temperatures gauge
atlas_temperatures{deviceid="0"} 41
atlas_temperatures{deviceid="1"} 42
atlas_temperatures{deviceid="2"} 42
atlas_temperatures{deviceid="3"} 42
atlas_temperatures{deviceid="4"} 48
atlas_temperatures{deviceid="5"} 46
atlas_temperatures{deviceid="6"} 46
atlas_temperatures{deviceid="7"} 45
atlas_temperatures{deviceid="8"} 45
atlas_temperatures{deviceid="9"} 44
```
