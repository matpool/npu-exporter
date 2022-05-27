package main

import (
	"testing"
)

var (
	testStr = `+------------------------------------------------------------------------------+
| npu-smi 21.0.3.1                     Version: 21.0.3.1                       |
+-------------------+-----------------+----------------------------------------+
| NPU     Name      | Health          | Power(W)          Temp(C)              |
| Chip    Device    | Bus-Id          | AICore(%)         Memory-Usage(MB)     |
+===================+=================+========================================+
| 1       310       | OK              | 12.8              42                   |
| 0       0         | 0000:01:00.0    | 0                 2621 / 8192          |
+-------------------+-----------------+----------------------------------------+
| 1       310       | OK              | 12.8              43                   |
| 1       1         | 0000:02:00.0    | 0                 2703 / 8192          |
+-------------------+-----------------+----------------------------------------+
| 1       310       | OK              | 12.8              42                   |
| 2       2         | 0000:03:00.0    | 0                 2621 / 8192          |
+-------------------+-----------------+----------------------------------------+
| 1       310       | OK              | 12.8              42                   |
| 3       3         | 0000:04:00.0    | 0                 2621 / 8192          |
+===================+=================+========================================+
| 3       310       | OK              | 12.8              48                   |
| 0       4         | 0000:0A:00.0    | 0                 2621 / 8192          |
+-------------------+-----------------+----------------------------------------+
| 3       310       | OK              | 12.8              46                   |
| 1       5         | 0000:0B:00.0    | 0                 2621 / 8192          |
+-------------------+-----------------+----------------------------------------+
| 3       310       | OK              | 12.8              46                   |
| 2       6         | 0000:0C:00.0    | 0                 2621 / 8192          |
+-------------------+-----------------+----------------------------------------+
| 3       310       | OK              | 12.8              45                   |
| 3       7         | 0000:0D:00.0    | 0                 2621 / 8192          |
+===================+=================+========================================+
| 4       310       | OK              | 12.8              45                   |
| 0       8         | 0000:81:00.0    | 0                 2621 / 8192          |
+-------------------+-----------------+----------------------------------------+
| 4       310       | OK              | 12.8              44                   |
| 1       9         | 0000:82:00.0    | 0                 2621 / 8192          |
+-------------------+-----------------+----------------------------------------+
| 4       310       | OK              | 12.8              43                   |
| 2       10        | 0000:83:00.0    | 0                 2621 / 8192          |
+-------------------+-----------------+----------------------------------------+
| 4       310       | OK              | 12.8              47                   |
| 3       11        | 0000:84:00.0    | 0                 2621 / 8192          |
+===================+=================+========================================+
| 6       310       | OK              | 12.8              47                   |
| 0       12        | 0000:8B:00.0    | 0                 2621 / 8192          |
+-------------------+-----------------+----------------------------------------+
| 6       310       | OK              | 12.8              47                   |
| 1       13        | 0000:8C:00.0    | 0                 2621 / 8192          |
+-------------------+-----------------+----------------------------------------+
| 6       310       | OK              | 12.8              47                   |
| 2       14        | 0000:8D:00.0    | 0                 2621 / 8192          |
+-------------------+-----------------+----------------------------------------+
| 6       310       | OK              | 12.8              47                   |
| 3       15        | 0000:8E:00.0    | 0                 2621 / 8192          |
+===================+=================+========================================+`

	deviceCount = 16
)

func TestParseNpuDevices(t *testing.T) {
	devices := parseNpuDevices(testStr)

	if len(devices) != deviceCount {
		t.Errorf("devices count error, wanted:%d, accutally:%d", deviceCount, len(devices))
		return
	}

	if devices[0].NpuID != "1" && devices[1].NpuID != "310" &&
		devices[2].Health != "OK" && devices[3].PowerUsage != "12.8" &&
		devices[4].Temperature != "48" && devices[5].ChipID != "1" &&
		devices[6].Device != "6" && devices[7].BusID != "0000:0D:00.0" &&
		devices[8].AICore != "0" && devices[9].MemoryUsageMB != "2621 / 8192" {
		t.Error("parse failed")
	}
}