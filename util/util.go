package util

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

//生成uuid      32位
func GetUUID() string {
	uuId := uuid.NewV4()
	sep := "-"
	s := strings.Split(uuId.String(), sep)
	var result string
	for i := 0; i < len(s); i++ {
		result += s[i]
	}
	return result
}

func GetTime() time.Time{
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	timeWrong, _ := time.Parse("2006-01-02 15:04:05",timeStr)//先解析为time类型，这样直接存数据库时间不对，可以尝试一下
	year := timeWrong.Year()
	month := timeWrong.Month()
	day := timeWrong.Day()
	hour := timeWrong.Hour()
	min := timeWrong.Minute()
	sec := timeWrong.Second()
	timeCorrect := time.Date(year, month, day, hour, min, sec, 0, time.Local)//这样存数据库就对了。相当于给timeWrong加上了一个时区
	return timeCorrect
}

// 上传文件
func Upload(name string) (string, error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	file,err := os.Open(name)
	if err != nil {
		log.Println("open file failed, err:", err)
		return "",err
	}
	bin, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return "",err
	}
	fw, err := w.CreateFormFile("file", file.Name())
	if err != nil {
		fmt.Println(err)
		return "",err
	}
	_, err = fw.Write(bin)
	if err != nil {
		fmt.Println(err)
		return "",err
	}
	w.Close()
	req, err := http.NewRequest("POST", "https://xiaofa-lawyer.aegis-info.com/xiaofa-manager/fileManagementApi/pic/upload", buf)
	if err != nil {
		fmt.Println("req err: ", err)
		return "",err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("resp err: ", err)
		return "",err
	}
	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(ret))
	var m map[string]interface{}
	err = json.Unmarshal(ret, &m)
	if err != nil {
		fmt.Println("json unmarshal failed, err:", err)
		return "", err
	}
	return m["data"].(string), nil
}


// 把图片变base64
func ImagesToBase64(str_images string) string {
	//读原图片
	ff, _ := os.Open(str_images)
	defer ff.Close()
	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Read(sourcebuffer)
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	return sourcestring
}
// 图片获取，通过网址进行图片下载，并保存在本地
func GetFile(imagePath string) string{

	resp, err := http.Get(imagePath)
	if err != nil {
		fmt.Println("get pic failed, err:", err)
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read failed, err:", err)
		return ""
	}
	name := imagePath[len(imagePath)-10:]
	out, err := os.Create(name)
	if err != nil {
		fmt.Println("create file failed, err:", err)
		return ""
	}
	defer out.Close()
	io.Copy(out, bytes.NewReader(body))
	return out.Name()
}

// 删除文件
func DeleteFile(name string) {
	os.Remove(name)
}


func CallJava(){
	cmd := exec.Command("java","Hello","tom")

	out,err := cmd.Output()

	if err != nil {
		println("error:",err.Error())

	}

	fmt.Printf("call java class Hello: %q\n", string(out))



}