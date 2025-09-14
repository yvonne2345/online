package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//实现滞留文件上传工具
//录制文件目录 /data/recording/
//每个录制任务会创建一个子目录，比如 /date/recording/xxxx 和 /date/recording/yyyy
//每个录制目录下生成两类文件 *.ts (文件很多，大小MB级别) *.mp4 (文件较少，大小GB级别)
//每个录制任务完成后在对应文件夹下生成一个标识文件 recording_done.txt

// 1. 扫描滞留文件，完成上传
// 2. 在1的基础上支持并发上传
// 3. 在2的基础上思考如何对上传带宽进行控制

type MockAliyun struct {
}

func (a *MockAliyun) Write(b []byte) (n int, err error) {
	return len(b), nil
}

// 控制并发数量
const maxConcurrentUploads = 5

// 上传接口
func uploadFile(filePath string) error {
	// 模拟文件上传

	//readFile, err := ioutil.ReadFile(filePath)
	//if err != nil {
	//	errors.New("打开")
	//}
	//_, err = MockAliyun.Write(readFile)
	//if err != nil {
	//	return err
	//}
	fmt.Printf("Uploading file: %s\n", filePath)
	time.Sleep(2 * time.Second) // 模拟上传时间
	return nil
}

// 扫描指定目录下的滞留文件
func scanRecordingDirs(baseDir string) ([]string, error) {
	var fileList []string
	// 读取目录
	dirs, err := os.ReadDir(baseDir)
	if err != nil {
		return nil, err
	}

	// 遍历子目录
	for _, dir := range dirs {
		if dir.IsDir() {
			dirPath := filepath.Join(dir.Name())
			files, err := os.ReadDir(dirPath)
			if err != nil {
				return nil, err
			}

			// 遍历文件
			for _, file := range files {
				if filepath.Ext(file.Name()) == ".ts" || filepath.Ext(file.Name()) == ".mp4" {
					fileList = append(fileList, filepath.Join(dirPath, file.Name()))
				}
			}
		}
	}
	return fileList, nil
}

// 执行并发上传
func uploadFilesConcurrently(files []string, maxConcurrentUploads int) {
	var wg sync.WaitGroup
	// 使用通道控制并发数
	flag := make(chan struct{}, maxConcurrentUploads)

	// 遍历文件进行上传
	for _, file := range files {
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()

			// 获取信号量
			flag <- struct{}{}

			// 上传文件
			if err := uploadFile(filePath); err != nil {
				fmt.Printf("Failed to upload file %s: %v\n", filePath, err)
			}

			// 释放信号量
			<-flag
		}(file)
	}

	// 等待所有上传任务完成
	wg.Wait()
}

func main() {
	// 扫描录制目录
	baseDir := "/data/recording"
	files, err := scanRecordingDirs(baseDir)
	if err != nil {
		fmt.Printf("Error scanning recording directories: %v\n", err)
		return
	}

	// 如果有文件需要上传
	if len(files) > 0 {
		fmt.Printf("Found %d files to upload.\n", len(files))
		// 并发上传文件
		uploadFilesConcurrently(files, maxConcurrentUploads)
	} else {
		fmt.Println("No files to upload.")
	}
}
