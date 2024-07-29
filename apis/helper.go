package apis

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gofrs/uuid"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func RespondWith(w http.ResponseWriter, r *http.Request, route string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		RespondWithError(w, r, route, err, http.StatusInternalServerError)
		return
	}
}

func RespondWithError(w http.ResponseWriter, r *http.Request, route string, err error, responseCode int) {
	type response struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
	}

	if errors.Is(err, context.DeadlineExceeded) {
		responseCode = http.StatusGatewayTimeout
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)

	resp := response{
		Success: false,
		Error:   err.Error(),
	}

	_ = json.NewEncoder(w).Encode(resp) //nolint:errchkjson
}

func GenerateUUID() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}

type TemplateData struct {
	InfoName       string `json:"InfoName"`
	SchoolName     string `json:"SchoolName"`
	ExperienceName string `json:"ExperienceName"`
	SkillName      string `json:"SkillName"`
}

func TestRenderSimpleTemplate(w http.ResponseWriter, r *http.Request) {
	// 循环读取templates/html目录下模板文件
	files, err := ioutil.ReadDir("templates/html")
	if err != nil {
		panic(err)
	}
	// 遍历模板文件
	for _, file := range files {
		// 读取模板文件内容
		temp, err := ioutil.ReadFile("templates/html/" + file.Name())
		if err != nil {
			panic(err)
		}

		// 解析模板
		tmpl, err := template.New(file.Name()).Parse(string(temp))
		if err != nil {
			panic(err)
		}

		// 渲染模板
		name := strings.Split(file.Name(), ".")[0]
		file, err := os.Open("templates/json/" + name + ".json")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		var data TemplateData
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&data); err != nil {
			panic(err)
		}

		err = tmpl.Execute(w, data)
		var output bytes.Buffer
		if err != nil {
			panic(err)
		}

		// 输出渲染后的HTML
		fmt.Println(output.String())
	}
}
