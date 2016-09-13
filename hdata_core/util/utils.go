package util

import (
	"os"
	"regexp"
	"strings"
	"time"
)

/**
 * 线程休眠
 *
 * @param millis
 */
func Sleep(millis int64) {
	time.Sleep(time.Duration(millis) * time.Millisecond)
}

func GetColumns(columns, excludeColumns []string) []string {
	if excludeColumns == nil || len(excludeColumns) < 1 {
		if columns == nil {
			return nil
		} else {
			return columns
		}
	}
	cls := make([]string, 0)
	for _, column := range columns {
		for _, ecls := range excludeColumns {
			if ecls != column {
				cls = append(cls, column)
			}
		}
	}
	return cls
}

/**
 * 修复HDFS路径（将主机名改成IP）
 *
 * @param srcLocaltion
 * @param metastoreUris
 * @return
 */
func FixLocaltion(srcLocaltion, metastoreUris string) string {
	matche := regexp.MustCompile("(\\d+\\.){3}\\d+")
	matcheString := matche.FindAllString(strings.TrimSpace(strings.Split(metastoreUris, ",")[0]), -1)
	for _, masterIP := range matcheString {
		srcLocaltion = strings.Replace(srcLocaltion, "^hdfs://\\w+:", "hdfs://"+masterIP+":", 1)
	}
	return srcLocaltion
}

/**
 * 解析分区值
 *
 * @param partitions
 * @return
 */
func ParsePartitionValue(partitions string) []string {
	partitionValues := make([]string, 0)
	partitionKeyValue := strings.Split(partitions, "\\s*,\\s*")
	for _, v := range partitionKeyValue {
		tokens := strings.Split(v, "=")
		partitionValues = append(partitionValues, tokens[1])
	}
	return partitionValues
}

/**
 * 获取配置目录
 *
 * @return
 */
func GetConfigDir() string {
	return os.Getenv("hdata.conf.dir") + os.Getenv("file.separator")
}
