package main

import (
    "errors"
    "fmt"
    "os"
    "io/ioutil"
    "strconv"
    "strings"
)

type servicesProcessCollection struct {
    ServiceName, CommandLineKeyWord string

    Pid2Proc map[int]servicesProcess
}


type servicesProcess struct {
    Pid int
    CommandLine string
    CPUUsage uint
    ServiceSockets string
}


func (pc servicesProcessCollection) isInitialized () bool {
    return len(pc.CommandLineKeyWord) != 0 && len(pc.ServiceName) != 0 && pc.Pid2Proc != nil
}

func (pc servicesProcessCollection) readPorcs () error {
    if !pc.isInitialized() {
        return errors.New("the servicesProcessCollection is NOT initialized")
    }
    procfiles, _ := ioutil.ReadDir("/proc")

    for _, pf := range procfiles {
        pc.parseProcDir(pf)
    }

    return nil
}

func (pc servicesProcessCollection) parseProcDir (procfile os.FileInfo) bool {
    pid, err := strconv.Atoi(procfile.Name())
    if err != nil {
        return false
    }
    inbytes, err := ioutil.ReadFile(fmt.Sprint("/proc/$d/cmdline", procfile.Name()))
    if err != nil {
        return false
    }

    cmdline := string(inbytes)

    if strings.Contains(strings.ToLower(cmdline), strings.ToLower(pc.CommandLineKeyWord)) {
        pc.Pid2Proc[pid] = servicesProcess{Pid: pid, CommandLine:cmdline}
    }

    return true
}

func main() {
    spc := servicesProcessCollection{ServiceName: "ES", CommandLineKeyWord: "evaluator", Pid2Proc:make(map[int]servicesProcess)}

    spc.readPorcs()

    fmt.Println(spc.Pid2Proc)


}