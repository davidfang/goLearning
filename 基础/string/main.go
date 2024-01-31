package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
// 对比a源目录和目标目录的文件内容是否一致
func main() {
	// srcDir := "src"
	// destDir := "desrc"
	srcDir := os.Args[1]  // 第一个参数为第一个目录
	destDir := os.Args[2] // 第二个参数为第二个目录
	// 记录程序开始时间
	start := time.Now()

	defer func() {
		fmt.Printf("程序运行时间: %s\n", time.Since(start))
		// 记录程序结束时间
		end := time.Now()

		// 计算程序执行时间
		elapsed := end.Sub(start)

		// 记录程序执行时长日志
		fmt.Printf("程序开始时间：%s 结束时间：%s 执行时间：%s\n", start, end, elapsed)
	}()
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if file is a regular file
		if info.Mode().IsRegular() {
			srcPath := path
			destPath := filepath.Join(destDir, strings.TrimPrefix(path, srcDir))

			// Check if file exists in destination directory
			if _, err := os.Stat(destPath); os.IsNotExist(err) {
				fmt.Printf("文件不存在%s  %s\n", srcPath, destPath)

				// Copy file to destination directory
				// err := copyFile(srcPath, destPath)
				// if err != nil {
				// 	fmt.Printf("复制失败 %s 到 %s 错误: %v\n", srcPath, destPath, err)
				// }
			}
			// else {
			// 	fmt.Printf("%s 文件已存在于 %s\n", path, destDir)
			// }
		}

		if info.IsDir() {
			srcPath := path
			destPath := filepath.Join(destDir, strings.TrimPrefix(path, srcDir))

			// Check if file exists in destination directory
			if _, err := os.Stat(destPath); os.IsNotExist(err) {
				fmt.Printf("文件夹不存在%s  %s\n", srcPath, destPath)

				// Copy file to destination directory
				// err := copyDir(srcPath, destPath)
				// if err != nil {
				// 	fmt.Printf("复制文件夹失败 %s 到 %s 错误: %v\n", srcPath, destPath, err)
				// }
			}
			// else {
			// 	fmt.Printf("%s 文件夹已存在于 %s\n", path, destDir)
			// }
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the directory:", err)
	}
}

func copyFile(src, dest string) (err error) {
	out, err := os.Create(dest)
	if err != nil {
		return
	}
	defer out.Close()

	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	return os.Chmod(dest, os.ModePerm)
}

// 新增 copyDir 函数
func copyDir(srcDir, destDir string) error {
	err := os.MkdirAll(destDir, 0777)
	if err != nil {
		return err
	}
	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		srcFile := filepath.Join(srcDir, file.Name())
		destFile := filepath.Join(destDir, file.Name())
		if file.IsDir() {
			err = copyDir(srcFile, destFile)
		} else {
			err = copyFile(srcFile, destDir)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
