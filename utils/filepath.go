package utils

import (
	"errors"
	"os"
	"path/filepath"
)

// EnsureAbsolutePath 确保路径为绝对路径
// 如果传入的路径已经是绝对路径，直接返回
// 如果是相对路径，则基于当前工作目录转换为绝对路径
func EnsureAbsolutePath(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}
	
	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		return "", errors.New("获取当前工作目录失败")
	}
	
	// 将相对路径转换为绝对路径
	return filepath.Join(currentDir, path), nil
}

// EnsureAbsolutePathWithBase 基于指定基础目录确保路径为绝对路径
// 如果传入的路径已经是绝对路径，直接返回
// 如果是相对路径，则基于指定的基础目录转换为绝对路径
func EnsureAbsolutePathWithBase(path, baseDir string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}
	
	// 确保基础目录是绝对路径
	absBaseDir, err := EnsureAbsolutePath(baseDir)
	if err != nil {
		return "", errors.New("基础目录路径处理失败")
	}
	
	// 将相对路径转换为绝对路径
	return filepath.Join(absBaseDir, path), nil
}

// GetExecutableDir 获取可执行文件所在目录的绝对路径
func GetExecutableDir() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", errors.New("获取可执行文件路径失败")
	}
	return filepath.Dir(execPath), nil
}

// EnsureAbsolutePathFromExecutable 基于可执行文件目录确保路径为绝对路径
// 如果传入的路径已经是绝对路径，直接返回
// 如果是相对路径，则基于可执行文件所在目录转换为绝对路径
func EnsureAbsolutePathFromExecutable(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}
	
	// 获取可执行文件目录
	execDir, err := GetExecutableDir()
	if err != nil {
		// 如果获取可执行文件目录失败，回退到当前工作目录
		return EnsureAbsolutePath(path)
	}
	
	// 将相对路径转换为绝对路径
	return filepath.Join(execDir, path), nil
}