package main

import (
	. "github.com/lein007/goproject/common"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"regexp"
	"encoding/base64"
)

func init() {
	//tmpDir := os.TempDir()
	tmpDir := `imanPack.txt`
	if err := ProcExsit(tmpDir); err == nil {
		Prosset(tmpDir)
	} else {
		fmt.Println("进程已启动")
		os.Exit(1)
	}
}
func main() {
	str := Rf(`run.txt`)
	ls := strings.Split(str, "\r\n")
	timeout := 3600 * 24 * 365 * time.Second
	r := New(timeout)

	for i := 0; i < len(ls); i++ {
		ls[i] = regexp.MustCompile(`"([^"]+)"`).ReplaceAllStringFunc(ls[i], func(s string) string {
			input :=""
			if res := regexp.MustCompile(`"([^"]+)"`).FindStringSubmatch(s); len(res) > 1 {
				input = res[1]
			}
			
			return `"`+base64.StdEncoding.EncodeToString([]byte(input))+`"`
		})
		rs := strings.Split(ls[i], " ")
		for k, v := range rs {
			if res := regexp.MustCompile(`"([^"]+)"`).FindStringSubmatch(v); len(res) > 1 {
				uDec, _ := base64.StdEncoding.DecodeString(res[1])
				rs[k] = string(uDec)
			}
		}
		
		if res := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`).FindStringSubmatch(rs[0]); len(res) > 0 {
			n := res[0]
			p := rs[1:]
			r.Add(createTask_t(n, p))
		} else {
			n, _ := strconv.Atoi(rs[0])
			p := rs[1:]
			r.Add(createTask(n, p))
		}
		
	}
	if err := r.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
func chuli(limit int, p[]string) {
	i := 1
	for {
		Prosset(`imanPack.txt`)
		go func(p[]string, i int) {
			fmt.Println(p, time.Now().Format("2006-01-02 15:04:05"))
			cmd := exec.Command(p[0],p[1:]...)
			//将其他命令传入生成出的进程
			cmd.Stdin = os.Stdin
			//给新进程设置文件描述符，可以重定向到文件中
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			//开始执行新进程，不等待新进程退出
			cmd.Run()
		}(p, i)
		time.Sleep(time.Second * time.Duration(limit))
		i++
	}
}
func chuli_t(limit string, p[]string) {
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", limit, loc)
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), loc)
	sr:= endTime.Sub(startTime)
	fmt.Println(p, "间隔：",sr)
	if sr>0 {
		time.Sleep(sr)
		fmt.Println(p, time.Now().Format("2006-01-02 15:04:05"))
		cmd := exec.Command(p[0],p[1:]...)
		//将其他命令传入生成出的进程
		cmd.Stdin = os.Stdin
		//给新进程设置文件描述符，可以重定向到文件中
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		//开始执行新进程，不等待新进程退出
		cmd.Run()
	}
}
func createTask(limit int, p[]string) func(int) {
	return func(id int) {
		fmt.Println("正在执行任务", id)
		chuli(limit, p)
	}
}
func createTask_t(limit string, p[]string) func(int) {
	return func(id int) {
		fmt.Println("正在执行任务", id)
		chuli_t(limit, p)
	}
}
