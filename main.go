package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"ivai-api/apis"
	"ivai-api/models"
	"net/http"
	"os"
	"os/exec"

	"text/template"
)

func main() {
	// 初始化数据库连接
	err := models.InitDB()
	if err != nil {
		panic(err)
	}

	// testTransferPNG()

	testRenderTemplate()

	http.ListenAndServe(":8080", apis.Router())
}

func testTransferPNG() {
	htmlContent := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>My Resume</title>
		</head>
		<body>
			<h1>John Doe</h1>
			<p>Software Developer at Example Corp</p>
			<!-- 更多简历内容 -->
		</body>
		</html>
	`
	tempFile, err := ioutil.TempFile("", "resume-*.html")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	defer os.Remove(tempFile.Name()) // 确保临时文件被删除

	_, err = tempFile.Write([]byte(htmlContent))
	if err != nil {
		fmt.Println("Error writing to temp file:", err)
		return
	}
	tempFile.Close()

	// 调用wkhtmltoimage
	cmd := exec.Command("wkhtmltoimage", tempFile.Name(), "output.png")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running wkhtmltoimage:", err)
		return
	}

	fmt.Println("Image created successfully.")
}

func testRenderTemplate() {
	// 假设从数据库获取了module_html
	personalInfoTemplate := `<div class="personal-info">
			<h1>{{ .Name }}</h1>
			<p>Phone: {{ .Phone }}</p>
			<p>Email: {{ .Email }}</p>
	</div>`

	// 定义数据
	personalInfoData := struct {
		Name  string
		Phone string
		Email string
	}{
		Name:  "John Doe",
		Phone: "123-456-7890",
		Email: "john.doe@example.com",
	}

	// 解析模板
	tmpl, err := template.New("personalInfo").Parse(personalInfoTemplate)
	if err != nil {
		panic(err)
	}

	// 渲染模板
	var output bytes.Buffer
	err = tmpl.Execute(&output, personalInfoData)
	if err != nil {
		panic(err)
	}

	// 输出渲染后的HTML
	fmt.Println(output.String())
}
