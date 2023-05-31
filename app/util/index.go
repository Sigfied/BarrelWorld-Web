package util

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"regexp"
)

// FormatFileSize 字节的单位转换 保留两位小数
func FormatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2f B", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2f KB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f MB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f GB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f TB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else {
		//if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2f PB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

// FuzzySearch 模糊搜索
func FuzzySearch(substring string, array []minio.ObjectInfo) []minio.ObjectInfo {
	var result []minio.ObjectInfo
	pattern := fmt.Sprintf(".*%s.*", substring)
	reg := regexp.MustCompile(pattern)
	for _, str := range array {
		if reg.MatchString(str.Key) {
			result = append(result, str)
		}
	}
	return result
}
