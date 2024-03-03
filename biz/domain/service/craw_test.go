package service

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestCrawData(t *testing.T) {
	publishAt := "Mar 2, 2024 by Danny Vena"
	tmpTime := strings.Split(publishAt, " by")[0]
	//tmpTime := strings.ReplaceAll(strings.Split(publishAt, "by")[0], " ", "")
	layout := "Jan 2, 2006"
	publishTime, _ := time.Parse(layout, tmpTime)
	fmt.Println(publishTime)
}
