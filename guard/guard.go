package guard

import (
	"log"
	"time"
	"os"
	"os/exec"
	"bytes"
	"strings"
	"fmt"
	"lib"
	"strconv"
	"runtime"
)

func Start() {
	initStore()
	config := loadConfig()
	var d, e = time.ParseDuration(fmt.Sprintf("%ds", config.CheckFrequency))
	if e != nil {
		panic(e)
	}
	pros:=loadPrograms()
	prosMap:=func()map[string]*Program{
		 m:=make(map[string]*Program)
		for _,pro:=range pros{
			m[pro.Name]=pro
		}
		return m
	}()
	go watch(d, prosMap)
	done:=make(chan bool)
	for range time.Tick(d){
		fmt.Println(stackTrace(true))
	}
	<-done
}

func watch(d time.Duration, pros map[string]*Program) {
	watchingPid := make(map[int]*Program)
	watchProcess := func(pid int, pro *Program) {
		p, e := os.FindProcess(pid)
		pro.processStart(p)
		ps, e := p.Wait()
		if e != nil {
			log.Fatal(e)
		}
		if ps.Exited() {
			delete(watchingPid, pid)
			pro.processExited(p)
		}
	}
	for range time.Tick(d) {
		cmd := exec.Command("tasklist")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		for {
			line, err := out.ReadString('\n')
			if err != nil {
				break;
			}
			segments := strings.Fields(line)
			if len(segments) == 6 {
				name := segments[0]
				pid, _ := strconv.Atoi(segments[1])
				if lib.In(name, pros) && !lib.In(pid, watchingPid) {
					pro := pros[name]
					watchingPid[pid] = pro
					go watchProcess(pid, pro)
				}
			}
		}
	}
}
func stackTrace(all bool) string {
	// Reserve 10K buffer at first
	buf := make([]byte, 10240)

	for {
		size := runtime.Stack(buf, all)
		// The size of the buffer may be not enough to hold the stacktrace,
		// so double the buffer size
		if size == len(buf) {
			buf = make([]byte, len(buf)<<1)
			continue
		}
		break
	}

	return string(buf)
}    // Reserve 10K buffer at first