package example

import (
	"fmt"
	"github.com/ketianlin/ktools"
	"testing"
)

// Verify 检查是否含有可能产生注入的非法字符
// 返回值为true时表示含有非法字符，同时返回的字符串值为匹配到的非法字符
func TestSQL(t *testing.T) {
	sql := ktools.SQL
	verify, s := sql.Verify("select * from user where id = 1")
	fmt.Println("Verify:", verify, s)

	verify, s = sql.Verify("select id,name,age from user where id = 1")
	fmt.Println("Verify2:", verify, s)
}
