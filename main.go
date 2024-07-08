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
)

func main() {
	// 初始化数据库连接
	err := models.InitDB()
	if err != nil {
		panic(err)
	}

	// testTransferPNG()

	// testRenderTemplate()

	http.ListenAndServe(":8080", apis.Router())
}

// func testRenderTemplate() {
// 	templateID := 1

// 	// 获取模板和模块
// 	resume, err := getTemplate(db, templateID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 渲染简历
// 	htmlContent, err := renderResume(resume)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(htmlContent)
// }

// func renderResume(resume *models.Resume) (string, error) {
// 	t, err := template.ParseFiles("resumeTemplate.html")
// 	if err != nil {
// 		return "", err
// 	}

// 	var tpl bytes.Buffer
// 	if err := t.Execute(&tpl, resume); err != nil {
// 		return "", err
// 	}

// 	return tpl.String(), nil
// }

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
