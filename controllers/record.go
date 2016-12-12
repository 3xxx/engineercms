package controllers

//识别档案号，对档案号进行正则判别，如果符合正则表达式，则进行提取
import (
	// "github.com/nfnt/resize"
	// "image"
	// "image/draw"
	// "image/jpeg"
	// "image/png"
	// "io/ioutil"
	// "log"
	// "math/rand"
	// "os"
	"github.com/astaxie/beego"
	"path"
	"regexp"
	// "strconv"
	// "fmt"
	"strings"
	// "time"
)

//分离图号图名
func Record(filenameWithSuffix string) (Suffix, FileNumber, FileName, ProNumber, ProJiduan, ProLeixing, ProZhuanye string) {
	FileSuffix := path.Ext(filenameWithSuffix) //只留下后缀名
	LengthSuffix := len([]rune(FileSuffix))
	Suffix = SubString(FileSuffix, 1, LengthSuffix-1)
	// fmt.Println("扩展名", Suffix)

	var filenameOnly string
	// var FileNumber string
	// var FileName string
	//	var ProJiduan string
	//	var ProZhuanye string
	//	var ProLeixing string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, FileSuffix) //只留下文件名，无后缀
	// fmt.Println("文件全名：", filenameOnly)                                //filenameOnly= mai
	//这个测试一个字符串是否符合一个表达式。
	//    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	//    fmt.Println(match)
	//上面我们是直接使用字符串，但是对于一些其他的正则任务，你需要使用 Compile 一个优化的 Regexp 结构体。
	r, _ := regexp.Compile(`[[:upper:]]{2}[0-9]+[[:upper:]\.0-9]+[-][0-9]+[-][0-9]+[\p{Han} \(\)\/~]`)
	//这个结构体有很多方法。这里是类似我们前面看到的一个匹配测试。
	// fmt.Println(r.MatchString(filenameOnly))
	lengthname := len([]rune(filenameOnly))
	if r.MatchString(filenameOnly) { //如果符合正则表达式

		// 查找连续2个的大写字母
		// reg := regexp.MustCompile(`[[:upper:]]{2}`)
		// fmt.Printf("大写字母%q\n", reg.FindAllString(filenameOnly, -1))
		// ["H" "G"]
		blankloc := UnicodeIndex(filenameOnly, " ") // 查找空格这个字符的位置
		if blankloc == 0 {                          //如果没有空格,                                                   //如果没有空格，则用正则表达式获取编号
			re, _ := regexp.Compile("[^a-zA-Z0-9-.~]")
			loc := re.FindStringIndex(filenameOnly)
			if loc != nil { //如果有编号——如果没文件名？？？？？

				FileNumber = SubString(filenameOnly, 0, loc[0])
				// fmt.Println("文件编号：", FileNumber)
				FileName = SubString(filenameOnly, loc[0], lengthname-loc[0])
				// fmt.Println("文件名：", FileName)
			} else { //如果没有编号
				FileNumber = filenameOnly
				// fmt.Println("文件编号：", FileNumber)
				FileName = filenameOnly
				// fmt.Println("文件名：", filenameOnly)
			}
		} else { //如果有空格
			re, _ := regexp.Compile("[^a-zA-Z0-9-.~]")
			loc := re.FindStringIndex(filenameOnly)
			if loc != nil { //如果有编号
				FileNumber = SubString(filenameOnly, 0, loc[0])
				// fmt.Println("文件编号：", FileNumber)
				FileName = SubString(filenameOnly, loc[0], lengthname-loc[0])
				// fmt.Println("文件名：", FileName)
			} else { //如果没有编号
				FileNumber = filenameOnly
				// fmt.Println("文件编号：", FileNumber)
				FileName = filenameOnly
				// fmt.Println("文件名：", filenameOnly)
			}
		}
		//这里继续提取项目号-阶段-成果类别-专业
		//首先判断有无.号，如果有，则是旧图号
		dianhaoloc := UnicodeIndex(FileNumber, ".")
		//	fmt.Println("第一个“.”位置：", dianhaoloc) //
		if dianhaoloc != 0 { //如果是旧编号
			//项目编号
			ProNumber = SubString(FileNumber, 0, dianhaoloc-1)
			// fmt.Println("项目编号：", SubString(FileNumber, 0, dianhaoloc-1))
			//阶段
			ProJiduan = SubString(FileNumber, dianhaoloc-1, 1)
			// switch ProJiduan {
			// case "A":
			// 	fmt.Println("规划阶段：" + ProJiduan)
			// case "B":
			// 	fmt.Println("项目建议书阶段：" + ProJiduan)
			// case "C":
			// 	fmt.Println("可行性阶段：" + ProJiduan)
			// case "D":
			// 	fmt.Println("初步设计阶段：" + ProJiduan)
			// case "E":
			// 	fmt.Println("招标设计阶段：" + ProJiduan)
			// case "F":
			// 	fmt.Println("施工图设计阶段：" + ProJiduan)
			// case "G":
			// 	fmt.Println("竣工图阶段：" + ProJiduan)
			// case "L":
			// 	fmt.Println("专题：" + ProJiduan)
			// }
			//专业
			ProZhuanye = SubString(FileNumber, dianhaoloc+1, 1)
			// switch ProZhuanye {
			// case "1":
			// 	fmt.Println("综合：" + ProZhuanye)
			// case "2":
			// 	fmt.Println("规划：" + ProZhuanye)
			// case "3":
			// 	fmt.Println("测量：" + ProZhuanye)
			// case "4":
			// 	fmt.Println("地质：" + ProZhuanye)
			// case "5":
			// 	fmt.Println("水工：" + ProZhuanye)
			// case "6":
			// 	fmt.Println("建筑：" + ProZhuanye)
			// case "7":
			// 	fmt.Println("机电：" + ProZhuanye)
			// case "8":
			// 	fmt.Println("资环：" + ProZhuanye)
			// case "9":
			// 	fmt.Println("施工：" + ProZhuanye)
			// }
			//二级专业代码
		} else { //新编号
			jianhaoloc := UnicodeIndex(FileNumber, "-")
			//项目编号
			ProNumber = SubString(FileNumber, 0, jianhaoloc-2)
			// fmt.Println("项目编号：", SubString(FileNumber, 0, jianhaoloc-2))
			//阶段
			ProJiduan = SubString(FileNumber, jianhaoloc-2, 1)
			// switch ProJiduan {
			// case "A":
			// 	fmt.Println("规划阶段：" + ProJiduan)
			// case "B":
			// 	fmt.Println("项目建议书阶段：" + ProJiduan)
			// case "C":
			// 	fmt.Println("可行性阶段：" + ProJiduan)
			// case "D":
			// 	fmt.Println("初步设计阶段：" + ProJiduan)
			// case "E":
			// 	fmt.Println("招标设计阶段：" + ProJiduan)
			// case "F":
			// 	fmt.Println("施工图设计阶段：" + ProJiduan)
			// case "G":
			// 	fmt.Println("竣工图阶段：" + ProJiduan)
			// case "L":
			// 	fmt.Println("专题：" + ProJiduan)
			// }
			//文件类型
			ProLeixing = SubString(FileNumber, jianhaoloc-1, 1)
			switch ProLeixing {
			case "B":
				ProLeixing = "FB"
				// fmt.Println("技术报告：" + ProLeixing)
			case "D":
				ProLeixing = "FD"
				// fmt.Println("设计大纲：" + ProLeixing)
			case "G":
				ProLeixing = "FG"
				// fmt.Println("设计/修改通知单：" + ProLeixing)
			case "T":
				ProLeixing = "FT"
				// fmt.Println("工程图纸：" + ProLeixing)
			case "J":
				ProLeixing = "FJ"
				// fmt.Println("计算书：" + ProLeixing)
			}
			//专业
			ProZhuanye = SubString(FileNumber, jianhaoloc+1, 1)
			// switch ProZhuanye {
			// case "1":
			// 	fmt.Println("综合：" + ProZhuanye)
			// case "2":
			// 	fmt.Println("规划：" + ProZhuanye)
			// case "3":
			// 	fmt.Println("测量：" + ProZhuanye)
			// case "4":
			// 	fmt.Println("地质：" + ProZhuanye)
			// case "5":
			// 	fmt.Println("水工：" + ProZhuanye)
			// case "6":
			// 	fmt.Println("建筑：" + ProZhuanye)
			// case "7":
			// 	fmt.Println("机电：" + ProZhuanye)
			// case "8":
			// 	fmt.Println("资环：" + ProZhuanye)
			// case "9":
			// 	fmt.Println("施工：" + ProZhuanye)
			// }
			//二级专业代码
		}
	} else { //2016-6-29增加，如果不符合表达式，比如：05水利科技.pdf
		blankloc := UnicodeIndex(filenameOnly, " ") // 查找空格这个字符的位置
		if blankloc == 0 {                          //如果没有空格,                                                   //如果没有空格，则用正则表达式获取编号
			re, _ := regexp.Compile("[^a-zA-Z0-9-.~]") //应该用查找连续的数字
			loc := re.FindStringIndex(filenameOnly)
			if loc != nil { //如果有编号——如果没文件名？？？？？
				FileNumber = SubString(filenameOnly, 0, loc[0])
				// fmt.Println("文件编号：", FileNumber)
				FileName = SubString(filenameOnly, loc[0], lengthname-loc[0])
				// fmt.Println("文件名：", FileName)
			} else { //如果没有编号
				FileNumber = filenameOnly
				// fmt.Println("文件编号：", FileNumber)
				FileName = filenameOnly
				// fmt.Println("文件名：", filenameOnly)
			}
		} else { //如果有空格
			re, _ := regexp.Compile("[^a-zA-Z0-9-.~]")
			loc := re.FindStringIndex(filenameOnly)
			if loc != nil { //如果有编号
				FileNumber = SubString(filenameOnly, 0, loc[0])
				// fmt.Println("文件编号：", FileNumber)
				FileName = SubString(filenameOnly, loc[0], lengthname-loc[0])
				// fmt.Println("文件名：", FileName)
			} else { //如果没有编号
				FileNumber = filenameOnly
				// fmt.Println("文件编号：", FileNumber)
				FileName = filenameOnly
				// fmt.Println("文件名：", filenameOnly)
			}
		}
	}
	return Suffix, FileNumber, FileName, ProNumber, ProJiduan, ProLeixing, ProZhuanye
}

//分离规范名称为分类，编号，年代和名称，用于规范上传
func SplitStandardName(filenameWithSuffix string) (Category, Categoryname, FileNumber, Year, FileName, Suffix string) {
	FileSuffix := path.Ext(filenameWithSuffix) //只留下后缀名
	LengthSuffix := len([]rune(FileSuffix))
	Suffix = SubString(FileSuffix, 1, LengthSuffix-1)
	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, FileSuffix) //只留下文件名，无后缀
	lengthname := len([]rune(filenameOnly))
	// fmt.Println("文件全名：", filenameOnly)                                //filenameOnly= mai
	//这个测试一个字符串是否符合一个表达式。
	//    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	//    fmt.Println(match)
	//上面我们是直接使用字符串，但是对于一些其他的正则任务，你需要使用 Compile 一个优化的 Regexp 结构体。
	//	r, _ := regexp.Compile(`[\P{Han}+`)
	//采用数字和非数字来进行分离
	//如果符合标准的格式，才进行分离，否则不分离
	//标准格式正则表达式为：英文 空格 数字或英文 减号 数字 英文或汉字或其他
	//r, _ := regexp.Compile(`[[:upper:]]{2}[0-9]+[[:upper:]\.0-9]+[-][0-9]+[-][0-9]+[\p{Han} \(\)\/~]`)
	r, _ := regexp.Compile(`[a-zA-Z]+\s[0-9A-Za-z\.]+[-][0-9]+[\p{Han}a-zA-Z0-9 \(\)\/~]`)
	//这个结构体有很多方法。这里是类似我们前面看到的一个匹配测试。
	// fmt.Println(r.MatchString(filenameOnly))
	if r.MatchString(filenameOnly) { //如果符合正则表达式
		//		fmt.Printf("%q\n", "ok")
		//减号后的数字，2个或4个，然后就是名称
		jianhao := UnicodeIndex(filenameOnly, "-")

		FileName = SubString(filenameOnly, jianhao+1, lengthname)
		//		fmt.Printf("%q\n", FileName)
		reg := regexp.MustCompile(`[0-9]+`)
		Year1 := reg.FindAllString(FileName, -1)
		Year = Year1[0]
		// fmt.Printf("%q\n", Year)
		lengthyear := len([]rune(Year))
		lengthname1 := len([]rune(FileName))
		FileName = SubString(FileName, lengthyear, lengthname1)
		// fmt.Printf("%q\n", FileName)
		//bianhao
		//kongge he jianhao zhijian
		kongge := UnicodeIndex(filenameOnly, " ")
		jianhao = UnicodeIndex(filenameOnly, "-")
		//		fmt.Printf("%q\n", jianhao)
		FileNumber = SubString(filenameOnly, kongge+1, jianhao-kongge-1)
		// fmt.Printf("%q\n", FileNumber)
		//获取分类
		Categoryname = SubString(filenameOnly, 0, kongge) //2016-7-16将kongge-1改为kongge
		Category = SubString(filenameOnly, 0, 2)
		switch Category {
		case "GB":
			Category = "GB"
		case "SL":
			Category = "SL"
		case "DL":
			Category = "DL"
		case "JT":
			Category = "JT"
		case "JG":
			Category = "JG"
		case "JC":
			Category = "JC"
		case "DB":
			Category = "DB"
		case "TB":
			Category = "TB"
		case "CE":
			Category = "CECS"
		case "CJ":
			Category = "CJ"
		case "DG":
			Category = "DG"
		case "AW":
			Category = "AWWA"
		case "EN":
			Category = "EN"
		case "DI":
			Category = "DIN"
		case "JB":
			Category = "JB"
		case "BS":
			Category = "BS"
		case "HG":
			Category = "HG"
		case "HJ":
			Category = "HJ"
		case "HY":
			Category = "HY"
		case "SY":
			Category = "SY"
		case "IS":
			Category = "ISO"
		case "SH":
			Category = "SH"
		case "AS":
			Category = "ASME"
		case "AN":
			Category = "ANSI"
		case "JI":
			Category = "JIS"
		case "NA":
			Category = "NACE"
		case "RC":
			Category = "RCC"
		default: //图集
			Category = "Atlas"
		}
		//		reg = regexp.MustCompile(`[\P{Han}]+`)
		//		FileNumber1 := reg.FindAllString(filenameOnly, -1)
		//		FileNumber = FileNumber1[0]
		//		//	fmt.Printf("%q\n", FileNumber)
		//		lengthnumber := len([]rune(FileNumber))
		//		//		lengthname := len([]rune(filenameOnly))
		//		FileName = SubString(filenameOnly, lengthnumber, lengthname)
		//		//chazhaolianxushuzi
		//		fmt.Printf("%q\n", FileName)
		//		if SubString(FileNumber, lengthnumber-1, lengthnumber) == " " { //如果最后一个字符是空格,                                                   //如果没有空格，则用正则表达式获取编号
		//			FileNumber = SubString(FileNumber, 0, lengthnumber-1)
		//		}
		//	fmt.Printf("%q\nFileNumber：", FileNumber)
		//这里继续提取年代
		//2016-4-20减号后连续数字，不超过4位
		//		jianhao := UnicodeIndex(FileNumber, "-")
		//		lengthnumber = len([]rune(FileNumber))
		//		Year = SubString(FileNumber, jianhao+1, lengthnumber)
		//	fmt.Printf("%q\nYear：", Year)
		// fmt.Printf("%q\n", Category)
	} else {
		//如果不符合正则，则对字母和汉字，数字和汉字，空格前后进行分割
		blankloc := UnicodeIndex(filenameOnly, " ") // 查找空格这个字符的位置
		if blankloc == 0 {                          //如果没有空格, 找汉字
			reg := regexp.MustCompile(`[\P{Han}]+`)
			FileNumber1 := reg.FindAllString(filenameOnly, -1)
			FileNumber = FileNumber1[0]
			// fmt.Printf("%q\nNumber：", FileNumber)
			//		reg = regexp.MustCompile(`[\P{Han}]+`)
			// FileNumber = FileNumber
			lengthNumber := len([]rune(FileNumber))
			FileName = SubString(filenameOnly, lengthNumber, lengthname)
		} else { //有空格
			FileNumber = SubString(filenameOnly, 0, blankloc)
			FileName = SubString(filenameOnly, blankloc+1, lengthname)
		}
		Category = "Atlas"
		//		 fmt.Printf("%q\n", FileName)
	}
	return Category, Categoryname, FileNumber, Year, FileName, Suffix
}

//分离上面结果中FileNumber的分类GB和编号50268
//用于搜索
func SplitStandardFileNumber(filenumber string) (Category, Categoryname, Number string) {
	r, _ := regexp.Compile(`[a-zA-Z]+\s[0-9A-Za-z\.]+[-][0-9]+`)
	if r.MatchString(filenumber) { //如果符合正则表达式
		blankloc := UnicodeIndex(filenumber, " ") // 查找空格这个字符的位置
		jianhao := UnicodeIndex(filenumber, "-")
		Number = SubString(filenumber, blankloc+1, jianhao-blankloc-1)
		Categoryname = SubString(filenumber, 0, blankloc-1) //如果不够2位，返回全部范围

		Category = SubString(filenumber, 0, 2)
		switch Category {
		case "GB":
			Category = "GB"
		case "SL":
			Category = "SL"
		case "DL":
			Category = "DL"
		case "JT":
			Category = "JT"
		case "JG":
			Category = "JG"
		case "JC":
			Category = "JC"
		case "DB":
			Category = "DB"
		case "TB":
			Category = "TB"
		case "CE":
			Category = "CECS"
		case "CJ":
			Category = "CJ"
		case "DG":
			Category = "DG"
		case "AW":
			Category = "AWWA"
		case "EN":
			Category = "EN"
		case "DI":
			Category = "DIN"
		case "JB":
			Category = "JB"
		case "BS":
			Category = "BS"
		case "HG":
			Category = "HG"
		case "HJ":
			Category = "HJ"
		case "HY":
			Category = "HY"
		case "SY":
			Category = "SY"
		case "IS":
			Category = "ISO"
		case "SH":
			Category = "SH"
		case "AS":
			Category = "ASME"
		case "AN":
			Category = "ANSI"
		case "JI":
			Category = "JIS"
		case "NA":
			Category = "NACE"
		case "RC":
			Category = "RCC"
		default: //图集
			Category = "Atlas"
		}
	} else {
		Category = "Atlas"
	}
	// fmt.Printf("%q\n", Category)
	// fmt.Printf("%q\n", Number)
	return Category, Categoryname, Number
}

//下面这个没什么用了吧，用record代替
//对于01水电院企业标准.pdf如何办呢，所以最简单是取得第一个汉字的位置即可
func SubStrings(filenameWithSuffix string) (substr1, substr2 string) {
	fileSuffix := path.Ext(filenameWithSuffix) //只留下后缀名
	//	fmt.Println("fileSuffix=", fileSuffix)     //fileSuffix= .go
	var filenameOnly string
	var fulleFilename1 string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix) //只留下文件名，无后缀
	//	fmt.Println("filenameOnly=", filenameOnly)                        //filenameOnly= mai
	end := UnicodeIndex(filenameOnly, " ")
	//	fmt.Println(fulleFilename1)
	//	rs := []rune("SL8888CT-500-88 泵站厂房布置图")
	rl := len([]rune(filenameOnly))
	if end == 0 {
		// end = -1
		//如果没有空格，则用正则表达式获取
		re, _ := regexp.Compile("[^a-zA-Z0-9-~]") //2016-1-11日拟修改DZ122D.5-10-15~15.dwg
		loc := re.FindStringIndex(filenameOnly)
		// fmt.Println(str[loc[0]:loc[1]])
		beego.Info(loc[0])
		if loc != nil {
			end = loc[0]
			fulleFilename1 = SubString(filenameOnly, 0, end)
			beego.Info(fulleFilename1)
			end = end - 1
		} else {
			fulleFilename1 = filenameOnly
			end = -1
		}
	} else {
		fulleFilename1 = SubString(filenameOnly, 0, end) //这里不能用fullfilename，因为前面赋值后当做了int类型
	}
	end = end + 1
	fulleFilename2 := SubString(filenameOnly, end, rl) //这里不能用fullfilename，因为前面赋值后当做了int类型
	//	fmt.Println(fulleFilename1)
	return fulleFilename1, fulleFilename2
}

func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	} else {
		result = 0 //如果没有空格就返回0
	}
	return result
}

//如果不够length，返回全部长度范围
func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)
	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	// 返回子串
	return string(rs[begin:end])
}

//	fmt.Println(FileNumber)
//	rs := []rune("SL8888CT-500-88 泵站厂房布置图")
//	lengthname := len([]rune(filenameOnly))
//	if end == 0 { //如果没有空格，则用正则表达式获取编号
//		re, _ := regexp.Compile("[^a-zA-Z0-9-.~]")
//		loc := re.FindStringIndex(filenameOnly)
//		fmt.Println("loc=", filenameWithSuffix[loc[0]:loc[1]]) //国
//		fmt.Println("loc[0]=", loc[0])                         //loc[0]= 20
//		if loc != nil {
//			end = loc[0]
//			FileNumber = SubString(filenameOnly, 0, end)
//			fmt.Println("文件编号：", FileNumber)
//			// 查找连续的汉字——改成查找第一个汉字？
//			re = regexp.MustCompile(`[\p{Han}]`)
//			//			lochan := re.FindStringIndex(filenameOnly)
//			//			fmt.Println("第一个汉字位置：", lochan[0])
//			FileName = SubString(filenameOnly, lochan[0], lengthname-lochan[0])
//			fmt.Println("文件名称：", FileName)

//			re, _ = regexp.Compile("[-]")
//			loc = re.FindStringIndex(FileNumber)
//			if loc != nil {
//				end = loc[0]
//			} //7
//			fmt.Println("第一个“-”位置：", loc) //[7 8]

//			loc1 := re.FindAllStringIndex(FileNumber, -1) //[[7 8] [11 12]]
//			if loc1 != nil {
//				fmt.Println("取多个-位置的第一个", loc1[0][0])                        //11
//				fmt.Println("多个的-位置", re.FindAllStringIndex(FileNumber, -1)) //-1表示所有，1表示显示一个，2表示2个
//			}
//			//n换行
//			//查找连续的数字
//			re, _ := regexp.Compile(`[0-9]+`)
//			fmt.Println("连续的数字:", re.FindAllString(filenameOnly, -1))
//			//loc := re.FindStringIndex(filenameOnly)
//			//查找.和-之间的字符
//			re, _ = regexp.Compile(`\..*-`) //  .AT-500-
//			fmt.Println(".和-之间：", re.FindAllString(filenameOnly, -1))
//			//查找第二个字母之前
//			fulleFilename3 := SubString(FileNumber, 0, end-2) //SL888
//			fmt.Println("项目编号：", fulleFilename3)
//			fulleFilename3 = SubString(FileNumber, end-2, 1) //F
//			fulleFilename3 = SubString(FileNumber, end-1, 1) //T
//			fulleFilename3 = SubString(FileNumber, end+1, 1) //5
//			//fmt.Println(fulleFilename3)
//			fulleFilename3 = SubString(FileNumber, end+1, 3) //500
//			fmt.Println("专业中间编号：" + fulleFilename3)
//			fulleFilename3 = SubString(FileNumber, end+5, 3) //500
//			fmt.Println("专业编号：" + fulleFilename3)
//			end = end - 1
//		} else {
//			FileNumber = filenameOnly
//			end = -1
//		}
//	} else { //如果有空格
//		FileNumber = SubString(filenameOnly, 0, end) //这里不能用fullfilename，因为前面赋值后当做了int类型
//	}
//	end = end + 1
//	fulleFilename2 := SubString(filenameOnly, end, lengthname) //这里不能用fullfilename，因为前面赋值后当做了int类型
//	fmt.Println(FileNumber)
