package utils

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	// "path/filepath"
	"regexp"
	// "runtime"
	"context"
	"strconv"
	"strings"
	"sync"
	"testing"
)

//array 到 json str
// arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
// lang, err := json.Marshal(arr)
// if err == nil {
//     fmt.Println("================array 到 json str==")
//     fmt.Println(string(lang))
// }

func TestLog(t *testing.T) { //测试函数，参数必须是t *Testing.T
	f, err := os.Open("D:/file.log") //日志文件路径
	if err != nil {
		log.Fatalf("Open file error,%s\n", err.Error())
	}
	defer f.Close()
	var resultSlice []int                                          //存储最后的结果切片，因为是要找出每一个时间值，所以用切片来存储
	var exp = regexp.MustCompile(`[\d]+_[\d]+(executeTime:)[\d]+`) //正则表达式，是为了从日志文件里获23_3executeTime:1234这种数据结果。
	buf := bufio.NewReader(f)                                      //读取日志文件里的字符流
	for {                                                          //逐行读取日志文件
		line, err := buf.ReadString('\n')
		expArr := exp.FindAllString(line, -1)
		if len(expArr) > 0 {
			arr := strings.Split(expArr[0], ":")
			value, _ := strconv.Atoi(arr[1])
			resultSlice = append(resultSlice, value)
		}
		if err != nil {
			if err == io.EOF {
				break //表示文件读取完了
			}
		}
	}
	fmt.Println(len(resultSlice)) //打印出结果的总条数
	// bubbleSort(resultSlice)       //对结果排序
	amount := 0
	for i := 0; i < len(resultSlice); i++ {
		amount += resultSlice[i]
	}
	average := amount / len(resultSlice)
	fmt.Println("average", average, "amount", amount)
	for k, v := range resultSlice {
		fmt.Println(k, v) //逐条逐条打印出时间。
	}
}

//注：在这里省略了冒泡排序函数。

//Go1.9按行读取日志文件并处理
// package main

// import (
//     "bufio"
//     "bytes"
//     "context"
//     "log"
//     "os"
//     "sync"
// )

const (
	logname    = "log/engineercms.log"
	concurrent = 5                //并发处理数,可以根据物理内存调整
	maxsize    = 10 * 1024 * 1024 //每次读取的大小,可以根据物理内存调整
)

var bufs = make([][]byte, concurrent)

func main() {
	var (
		chs         = make(map[int]chan int)
		wait        = new(sync.WaitGroup)
		chct        = make(chan int, concurrent)
		ctx, cancel = context.WithCancel(context.Background())
	)

	File, err := os.Open(logname)
	if err != nil {
		log.Fatalf("Open file error,%s\n", err.Error())
	}

	for i := 0; i < concurrent; i++ {
		chct <- i
		chs[i] = make(chan int)
		bufs[i] = make([]byte, maxsize)
		go resolvectx(ctx, wait, i, chct, chs[i])
	}

	var i, n, l int
	for i = range chct {
		n, err = File.Read(bufs[i])
		if err != nil {
			wait.Wait() //等待数据全部处理完毕,然后返回
			break
		}
		for s := 1; s < n; s++ { //如果行过长,那么效率会变低
			if bufs[i][n-s] == '\n' {
				n = n - s + 1
				File.Seek(int64(l+n), 0)
				break
			}
		}
		l += n
		wait.Add(1)
		chs[i] <- n
	}
	cancel()
	close(chct)
	File.Close()
}

func resolvectx(ctx context.Context, wait *sync.WaitGroup, index int, chct, ch chan int) {
	var (
		err    error
		line   []byte
		length int
		buf    = bufio.NewReader(nil)
	)

	for {
		select {
		case <-ctx.Done():
			return
		case length = <-ch:
			buf.Reset(bytes.NewBuffer(bufs[index][:length]))
			for {
				line, _, err = buf.ReadLine()
				if err != nil {
					break
				}
				_ = line
			}
			chct <- index
			wait.Done()
		}
	}
}

//从尾部读取文件
const readbuf = 1 << 10 //1024

type flags struct {
	line   int
	lines  int
	size   string
	offset string
	output string
	file   string
}

type Command struct {
	UsageLine string
	Run       func(*Command, []string) bool
	Short     string
	Long      string
}

var args flags

func init() {
	flag.StringVar(&args.output, "o", "", "-o 指定输出的路径,不指定则输出到标准输出")
	flag.IntVar(&args.line, "l", 0, "-l 指定从倒数第几行开始读取")
	flag.IntVar(&args.lines, "n", 0, "-n 指定读取的行数")
	flag.StringVar(&args.offset, "i", "", "-i 指定开始读取的位置,单位:b,kb,mb,默认单位:b")
	flag.StringVar(&args.size, "s", "", "-s 指定读取的大小,单位:b,kb,mb,默认单位:b")
	flag.StringVar(&args.file, "f", "", "-f 指定要查看的文件路径")
}

var Tail = &Command{
	UsageLine: `tail -f main.go -l 10 -n 5 -o tmp.txt`,
	Run:       tail,
	Short:     "从文件结尾或指定位置读取内容",
	Long: `从文件结尾或指定位置读取内容,可以按行读取,也可以按大小读取,-i 和 -l同时使用的话-i生效,-s 与 -n 
同时使用的话-s生效`,
}

func tail(cmd *Command, arg []string) bool {
	File, err := os.OpenFile(args.file, os.O_RDONLY, 0644)
	if err != nil {
		log.Printf("打开文件失败:%s\n", err.Error())
		return true
	}

	defer File.Close() //20191108添加close

	var w io.Writer
	if args.output != "" {
		w, err = os.Create(args.output)
		if err != nil {
			log.Printf("创建输出文件失败:%s\n", err.Error())
			return true
		}
	} else {
		w = os.Stdout
	}

	f := NewTail(File)
	defer f.Close()

	if args.offset != "" {
		offset := parseCompany(args.offset)
		if offset == 0 {
			return false
		}
		size := parseCompany(args.size)
		if err = f.Read(offset, size, int64(args.lines), w); err != nil {
			log.Printf("读取内容错误:%s\n", err.Error())
			return true
		}
	}
	if args.line > 0 {
		if err = f.TailLine(args.line, args.lines, w); err != nil {
			log.Printf("读取内容错误:%s\n", err.Error())
			return true
		}
	}
	return false
}

func parseCompany(c string) int64 {
	if len(c) < 1 {
		return 0
	}

	if c[len(c)-1] != 'b' {
		c += "b"
	}

	index := len(c) - 2
	var company int64 = 1
	switch c[index] {
	case 'k':
		company = 1 << 10 //1024
	case 'm':
		company = 1 << 20 //1024*1024
	default:
		index++
	}
	size, err := strconv.ParseInt(c[:index], 10, 0)
	if err != nil {
		return -1
	}
	return size * company
}

func NewTail(File *os.File) *TailFile {
	offset, _ := File.Seek(0, 2)
	return &TailFile{file: File, size: offset, offset: offset}
}

type TailFile struct {
	mu     sync.Mutex
	file   *os.File
	size   int64
	offset int64
}

func (f *TailFile) read(p []byte) (n int, err error) {
	if f.offset == 0 {
		return 0, io.EOF
	}
	var (
		offset int
		length = int64(len(p))
	)
	if f.offset >= length {
		f.offset -= length
	} else {
		offset = int(f.offset)
		f.offset = 0
	}
	_, err = f.file.Seek(f.offset, 0)
	if err == nil {
		n, err = f.file.Read(p)
		if offset != 0 && offset < n {
			n = offset
		}
	}
	return
}

func (f *TailFile) Read(offset, size, lines int64, w io.Writer) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.size < offset {
		return io.EOF
	}
	f.file.Seek(offset, 0)
	switch {
	case size > 0:
		if _, err := io.CopyN(w, f.file, size); err != io.EOF {
			return err
		}
	case lines > 0:
		buf := bufio.NewReader(f.file)
		for i := 0; i < int(lines); i++ {
			line, err := buf.ReadBytes('\n')
			if err != nil {
				return err
			}
			if _, err = w.Write(line); err != nil {
				return err
			}
		}
	}
	return nil
}

//暂时用不到
func (f *TailFile) ReadLine() ([]byte, error) {
	var list []*[]byte
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.offset == 0 {
		return nil, io.EOF
	}

	for {
		var buf = make([]byte, 256)
		n, err := f.read(buf[:])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		index := bytes.LastIndexByte(buf[:n], '\n')
		if index >= 0 {
			f.offset += int64(index)
			f.file.Seek(f.offset, 0)
			buf = buf[index:n]
			list = append(list, &buf)
			break
		} else {
			buf = buf[:n]
			list = append(list, &buf)
		}
	}

	var (
		last int
		line = make([]byte, len(list)*256)
	)

	for i := len(list) - 1; i >= 0; i-- {
		copy(line[last:], *list[i])
		last += len(*list[i])
	}

	return line[:last], nil
}

func (f *TailFile) TailLine(line, lines int, w io.Writer) (err error) {
	buf := make([]byte, readbuf)

	f.mu.Lock()
	defer f.mu.Unlock()
	var (
		n, l int
		sep  = []byte("\n")
	)
	for {
		n, err = f.read(buf)
		if err != nil {
			if err == io.EOF {
				f.file.Seek(0, 0)
				break
			}
		}
		l += bytes.Count(buf[:n], sep)
		if l >= line {
			var seek, i int
			buf = buf[:n]
			for l > line {
				i = bytes.Index(buf, sep) + 1
				buf = buf[i:]
				seek += i
				l--
			}
			f.file.Seek(f.offset+int64(seek), 0)
			break
		}
	}
	if lines == 0 {
		_, err = io.Copy(w, f.file)
	} else {
		buf := bufio.NewReader(f.file)
		var line []byte
		for i := 0; i < lines; i++ {
			line, err = buf.ReadBytes('\n')
			if err != nil {
				return
			}
			if _, err = w.Write(line); err != nil {
				return
			}
		}
	}
	return
}

func (f *TailFile) Close() error {
	return f.file.Close()
}

// golang文件读取-按行读取

// a.txt文件内容:

// ABCDEFGHI
// HELLO GOLANG

// package main

// import (
//     "fmt"
//     "os"
//     "io"
//     "bufio"
//     "strings"
// )

func mainttr() {
	fileName := "log/engineercms.log"
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}
}
