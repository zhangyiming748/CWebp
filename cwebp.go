package CWebp

import (
	"github.com/zhangyiming748/CWebp/log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func CWebp(src, pattern, dst string) {
	files := getFiles(src, pattern)
	log.Info.Println(files)
	l := len(files)
	for index, file := range files {
		cwebp_help(src, file, dst, index, l)
	}
}
func cwebp_help(src, file, dst string, index, total int) {
	in := strings.Join([]string{src, file}, "/")
	extname := path.Ext(file)
	filename := strings.Trim(file, extname)
	out := strings.Join([]string{dst, strings.Join([]string{filename, "webp"}, ".")}, "/")
	cmd := exec.Command("cwebp", in, "-o", out)
	log.Debug.Printf("开始处理文件%s\t生成的命令是:%s", file, cmd)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Debug.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		log.Debug.Printf("cmd.Run产生的错误:%v", err)
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		log.Info.Printf("正在处理第 %d/%d 个文件: %s\n", index+1, total, file)
		t := string(tmp)
		t = strings.Replace(t, "\u0000", "", -1)
		log.Info.Println(t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Debug.Println("命令执行中有错误产生", err)
	}
	log.Debug.Printf("完成当前文件的处理:源文件是%s\t目标文件是%s\n", in, out)
}

func getFiles(dir, pattern string) []string {
	files, _ := os.ReadDir(dir)
	var aim []string
	for _, f := range files {
		//fmt.Println(f.Name())
		if l := strings.Split(f.Name(), ".")[0]; len(l) != 0 {
			if strings.HasSuffix(f.Name(), pattern) {
				log.Debug.Printf("有效的目标文件:%v\n", f.Name())
				//absPath := strings.Join([]string{dir, f.Name()}, "/")
				//log.Printf("目标文件的绝对路径:%v\n", absPath)
				aim = append(aim, f.Name())
			}
		}
	}
	return aim
}
