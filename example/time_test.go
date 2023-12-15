package example

import (
	"fmt"
	"github.com/ketianlin/ktools"
	"math/rand"
	"testing"
	"time"
)

func TestTimeUtil(t *testing.T) {
	// string转time
	str1 := "2023-11-11 15:00:10"
	tt, _ := ktools.Time.TUtil.StringToTime(str1)
	fmt.Printf("tt1: %T\t%#v\n", tt, tt.Format("2006-01-02 15:04:05"))

	// string转time
	str2 := "2023/11/11 15:00:10"
	tt2, _ := ktools.Time.TUtil.StringToTime(str2, "2006/01/02 15:04:05")
	fmt.Printf("tt2: %T\t%#v\n", tt2, tt2.Format("2006-01-02 15:04:05"))

	// int64转time
	var sou int64 = 1700805792
	st := ktools.Time.TUtil.Int64ToTime(sou)
	fmt.Printf("tt2: %T\t%#v\n", st, st.Format("2006-01-02 15:04:05"))

	var sou2 int64 = 1700807375839
	st2 := ktools.Time.TUtil.Int64ToTime(sou2)
	fmt.Printf("tt2: %T\t%#v\n", st2, st2.Format("2006-01-02 15:04:05"))

	curTime := ktools.Time.TFmt.GetNowDateTime()
	fmt.Println(curTime)
}

func TestTime5(tt *testing.T) {
	t1 := "2023-09-27T08:00:54+08:00"
	formatString, _ := ktools.Time.TUtil.RFC3339TimeStringToFormatString(t1)
	fmt.Println(formatString)
	formatString2, _ := ktools.Time.TUtil.RFC3339TimeStringToFormatString(t1, "2006/01/02 15:04:05")
	fmt.Println(formatString2)
}

func TestTime3(tt *testing.T) {
	str1 := "2023-11-11T15:00:10Z"
	ttt, _ := ktools.Time.TUtil.StringToTime2(str1)
	fmt.Printf("tt1: %T\t%v\n", ttt, ttt)
}

func TestTime2(tt *testing.T) {
	//fmt.Println(time.Now())
	t := "2019-10-10 10:10:10"
	t1, _ := time.Parse("2006-01-02 15:04:05", t)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t1.Equal(t2))
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	fmt.Println("SH-now : ", time.Now().In(cstSh))
	fmt.Println("SH-now2 : ", time.Now())
	fmt.Println("SH : ", time.Now().In(cstSh).Format("2006-01-02 15:04:05"))
	fmt.Println("SH-t1 : ", t1.In(cstSh).Format("2006-01-02 15:04:05"))
	fmt.Println("SH-t2 : ", t2.In(cstSh).Format("2006-01-02 15:04:05"))

	//时区转换
	fmt.Println("***************")
	t = "2021-01-11T23:46:05Z"
	t1, _ = time.Parse("2006-01-02T15:04:05Z", t)
	fmt.Println(t)
	fmt.Println("SH : ", t1.In(cstSh).Format("2006-01-02 15:04:05"))
}

func TestTime6(tt *testing.T) {
	curTime := time.Now()
	end := curTime.Format("2006-01-02 15:04:05")
	start := ktools.Time.TCalc.SubtractMinutes(curTime, 30).Format("2006-01-02 15:04:05")
	fmt.Printf("start: %T\t%v\n", start, start)
	fmt.Printf("end: %T\t%v\n", end, end)
}

func TestTime7(tt *testing.T) {
	a := fmt.Sprintf("%s%d", ktools.Time.TUtil.GetDate("20060102"), rand.Intn(9000)+1000)
	fmt.Println(a)
	for i := 0; i < 300; i++ {
		num := ktools.Math.Random.GenerateSpecifyIntervalNumber(1000, 9000)
		fmt.Println(num)
	}
}
