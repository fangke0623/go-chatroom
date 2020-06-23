package fileHandle

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"wechat/src/common/util"
)

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening fileHandle")
		return err
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(respBody))
	return nil
}
func UploadBase64(param []byte) {
	entity := FileEntity{}
	util.HandleParamsToStruct(param, &entity)
	decodeData, err := base64.StdEncoding.DecodeString(entity.FileContent)
	if err != nil {
		log.Println(err)
	}
	f, err := os.OpenFile("/go/file/"+entity.FileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if f == nil {
		return
	}
	if err != nil {
		log.Println(err.Error())
	}
	defer f.Close()
	_, err = f.Write(decodeData)
	if err != nil {
		log.Println(err.Error())
	}
}
