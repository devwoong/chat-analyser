package parser

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

// Kakao is
type Kakao struct {
	Date     string `json:"date"`
	Author   string `json:"author"`
	Contents string `json:"contents"`
}

func (kakao *Kakao) parse(logLine string) error {
	if logLine == "" {
		return errors.New("empty line")
	}
	if strings.Contains(logLine, ",") == true {
		firstSplice := strings.Split(logLine, ",")
		kakao.Date = dateFormat(firstSplice[0])
		log.Println(kakao.Date)
		excludeDateLog := strings.Join(firstSplice[1:], ",")
		if strings.Contains(excludeDateLog, ":") {
			secondSplice := strings.Split(excludeDateLog, ":")
			kakao.Author = strings.TrimSpace(secondSplice[0])
			kakao.Contents = strings.Join(secondSplice[1:], ":")
		} else {
			return errors.New("not found autor and contents")
		}
	} else {
		return errors.New("no log")
	}
	return nil
}

func dateFormat(date string) string {
	date = strings.TrimSpace(date)
	temp := strings.Split(date, "년")
	year := temp[0]
	temp = strings.Split(strings.Join(temp[1:], ""), "월")
	month := strings.TrimSpace(temp[0])
	integer, _ := strconv.Atoi(month)
	if integer < 10 {
		month = "0" + month
	}
	temp = strings.Split(strings.Join(temp[1:], ""), "일")
	day := strings.TrimSpace(temp[0])
	integer, _ = strconv.Atoi(day)
	if integer < 10 {
		day = "0" + day
	}

	var hour string
	var minute string
	t := strings.Join(temp[1:], "")
	if strings.Contains(t, "오후") == true {
		times := strings.Replace(t, "오후", "", 1)
		integer, _ := strconv.Atoi(strings.TrimSpace(strings.Split(times, ":")[0]))
		integer += 12
		hour = strconv.Itoa(integer)
		minute = strings.Split(times, ":")[1]
	} else if strings.Contains(t, "오전") == true {
		times := strings.Replace(t, "오전", "", 1)
		hour = strings.TrimSpace(strings.Split(times, ":")[0])
		integer, _ = strconv.Atoi(hour)
		if integer < 10 {
			hour = "0" + hour
		}
		minute = strings.Split(times, ":")[1]
	}

	return year + "-" + month + "-" + day + " " + hour + ":" + minute
}
