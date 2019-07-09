package common

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var MY_USER_AGENT = []string{
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; AcooBrowser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506)",
	"Mozilla/4.0 (compatible; MSIE 7.0; AOL 9.5; AOLBuild 4337.35; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
	"Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Media Center PC 6.0)",
	"Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
	"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.2; .NET CLR 3.0.04506.30)",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN) AppleWebKit/523.15 (KHTML, like Gecko, Safari/419.3) Arora/0.3 (Change: 287 c9dfb30)",
	"Mozilla/5.0 (X11; U; Linux; en-US) AppleWebKit/527+ (KHTML, like Gecko, Safari/419.3) Arora/0.6",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.2pre) Gecko/20070215 K-Ninja/2.1.1",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9) Gecko/20080705 Firefox/3.0 Kapiko/3.0",
	"Mozilla/5.0 (X11; Linux i686; U;) Gecko/20070322 Kazehakase/0.4.5",
	"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.8) Gecko Fedora/1.9.0.8-1.fc10 Kazehakase/0.5.6",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.20 (KHTML, like Gecko) Chrome/19.0.1036.7 Safari/535.20",
	"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; fr) Presto/2.9.168 Version/11.52",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.11 TaoBrowser/2.0 Safari/536.11",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.71 Safari/537.1 LBBROWSER",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; LBBROWSER)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E; LBBROWSER)",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.84 Safari/535.11 LBBROWSER",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E)",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; QQBrowser/7.0.3698.400)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; SV1; QQDownload 732; .NET4.0C; .NET4.0E; 360SE)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E)",
	"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.89 Safari/537.1",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.89 Safari/537.1",
	"Mozilla/5.0 (iPad; U; CPU OS 4_2_1 like Mac OS X; zh-cn) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8C148 Safari/6533.18.5",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:2.0b13pre) Gecko/20110307 Firefox/4.0b13pre",
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:16.0) Gecko/20100101 Firefox/16.0",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11",
	"Mozilla/5.0 (X11; U; Linux x86_64; zh-CN; rv:1.9.2.10) Gecko/20100922 Ubuntu/10.10 (maverick) Firefox/3.6.10",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
}

//写文件
func Wf(fname string, str string) {
	userFile := fname
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}

	fout.WriteString(str)
}

//追加文件
func Wfa(fname string, str string) {
	fd, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	str += "\r\n"
	fd.WriteString(str)
}

//读取文件
func Rf(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}

//移动文件
func Rv(path string, newpath string) {
	reg := regexp.MustCompile(`(.*)(\\[^\\]+)`)
	dir := reg.FindStringSubmatch(newpath)[1]
	//fmt.Println(dir)
	os.MkdirAll(dir, 0777)
	err := os.Rename(path, newpath)
	if err != nil {
		fmt.Println(err)
	}
}

//删除文件
func Df(path string) {
	//删除文件
	del := os.Remove(path)
	if del != nil {
		fmt.Println(del)
	}
}

//获取内容
func GetContent(url string) (string, string) {
	var b bytes.Buffer

	rand.Seed(time.Now().UnixNano()) //利用当前时间的UNIX时间戳初始化rand包
	//num_h := rand.Intn(len(MY_USER_AGENT) - 1)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	//req.Header.Add("User-Agent", MY_USER_AGENT[num_h])
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return "", "网络异常"
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("网络异常:" + strconv.Itoa(resp.StatusCode))
		return "", "网络异常"
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024)

	for {
		n, _ := resp.Body.Read(buf)
		if 0 == n {
			break
		}
		b.WriteString(string(buf[:n]))
	}
	return b.String(), ""
}

// 判断进程是否启动
func ProcExsit(tmpDir string) (err error) {
	filePid := Rf(tmpDir)
	if len(filePid) > 0 {
		pidStr := fmt.Sprintf("%s", filePid)
		fmt.Println(pidStr)
		pid, _ := strconv.Atoi(pidStr)
		_, err := os.FindProcess(pid)
		if err == nil {
			return errors.New("已启动.")
		}
	}
	return nil
}

//判断最后更新时间是否小于两分钟
func Procheck(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	if _, err := os.Stat("over.txt"); err != nil {
		return true
	}

	timestamp := fileInfo.ModTime().Unix()
	tm := time.Unix(timestamp, 0)
	fmt.Println(tm.Format("2006-01-02 03:04:05"))
	t := time.Now().Unix()
	if t-timestamp < 120 {
		return true
	}

	return false
}

//终止进程
func Prostop(path string) {
	c := Rf(path)
	pid, err := strconv.ParseInt(c, 10, 32)
	pro, err := os.FindProcess(int(pid))
	if err != nil {
		return
	}
	pro.Kill()
	KillAll(int(pid))

	os.Remove(path)
}

//设置进程
func Prosset(tmpDir string) {
	iManPid := fmt.Sprint(os.Getpid())
	pidFile, _ := os.Create(tmpDir)
	defer pidFile.Close()
	pidFile.WriteString(iManPid)
}

//转换编码
// str := "乱码的字符串变量"
// str  = ConvertToString(str, "gbk", "utf-8")
// fmt.Println(str)
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

//判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func HttpPost(url string, opt string) (string, string) {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(opt))
	if err != nil {
		return "", "网络异常"
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "网络异常"
	}

	return string(body), ""
}

func HttpGet(url string) (string, string) {
	time.Sleep(time.Millisecond * 3000)

	resp, err := http.Get(url)
	if err != nil {
		return "", "网络异常"
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "网络异常"
	}

	return string(body), ""
}

//替换sql字符串
func SqlReplace(str string) string {
	str = strings.Replace(str, `\`, `\\`, -1)
	str = strings.Replace(str, `'`, `\'`, -1)
	str = strings.Replace(str, `\r\n`, ``, -1)
	str = strings.TrimSpace(str)
	return str
}
func SqlReplacePg(str string) string {
	str = strings.Replace(str, `\`, `\\`, -1)
	str = strings.Replace(str, `'`, `''`, -1)
	str = strings.Replace(str, `\r\n`, ``, -1)
	str = strings.TrimSpace(str)
	return str
}

//处理电话
func Chulihaoma(str string) string {
	str_t := regexp.MustCompile(`[;；/]+`).ReplaceAllString(str, ";")
	str_t = regexp.MustCompile(`[;]+`).ReplaceAllString(str_t, ";")
	str_t = regexp.MustCompile(`[^\d;；]+`).ReplaceAllString(str_t, "")
	str_t = strings.TrimSpace(str_t)
	ls := strings.Split(str, ";")
	ls_t := []string{}
	for i := 0; i < len(ls); i++ {
		if res, _ := regexp.Match(`^(((13[0-9]{1})|(14[0-9]{1})|(16[0-9]{1})|(15[0-9]{1})|(17[0-9]{1})|(18[0-9]{1})|(19[0-9]{1}))+\d{8})$`, []byte(ls[i])); res {
			ls_t = append(ls_t, ls[i])
		} else if res, _ := regexp.Match(`^([0-9]{3,4}-)?[0-9]{7,8}$`, []byte(ls[i])); res {
			ls_t = append(ls_t, ls[i])
		}
	}

	return strings.Join(ls_t, ";")
}
//使用变量
func Use(vals ...interface{}) {
    for _, val := range vals {
        _ = val
    }
}