package service

import (
	"fmt"
	"testing"
	"time"
)

func TestCrawData(t *testing.T) {
	//publishAt := "Mar 2, 2024 by Danny Vena"
	//tmpTime := strings.Split(publishAt, " by")[0]
	////tmpTime := strings.ReplaceAll(strings.Split(publishAt, "by")[0], " ", "")
	//layout := "Jan 2, 2006"
	//publishTime, _ := time.Parse(layout, tmpTime)
	//fmt.Println(publishTime)

	// 定义时间字符串和格式
	timeStr := "Mar 2, 2024 4:04 PM EST"
	layout := "Jan 2, 2006 3:04 PM MST" // 使用Go语言的基准时间

	// 将时间字符串解析为time.Time类型
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	// 输出转换后的时间
	fmt.Println("Parsed time:", parsedTime)

}
