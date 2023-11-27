package example

import (
	"fmt"
	"github.com/ketianlin/ktools"
	"testing"
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
