package main

import (
        "runtime"
        "syscall"
        "runtime/debug"
        "time"
        //"unsafe"
)

func main() {
        debug.SetGCPercent(-1)


        var stats runtime.MemStats
        buf := make([]byte, 0, 64)

        for {
                runtime.GC()
                runtime.ReadMemStats(&stats)
                buf = buf[:0]
                buf = append(buf, "HeapObjects: "...)
                buf = itoa(buf, stats.HeapObjects)
                buf = append(buf, '\n')
                syscall.Write(syscall.Stdout, buf)
                time.Sleep(time.Second)
        }
}

func itoa(buf []byte, i uint64) []byte {
        if i == 0 {
                return append(buf, '0')
        }

        var b [20]byte
        bp := len(b)
        for i > 0 {
                bp--
                b[bp] = byte(i%10) + '0'
                i /= 10
        }
        return append(buf, b[bp:]...)
}
