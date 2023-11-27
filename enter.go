package ktools

import (
	"github.com/ketianlin/ktools/crypto"
	"github.com/ketianlin/ktools/excel"
	fileT "github.com/ketianlin/ktools/file"
	jsonT "github.com/ketianlin/ktools/json"
	"github.com/ketianlin/ktools/jwt"
	mapT "github.com/ketianlin/ktools/map"
	"github.com/ketianlin/ktools/math"
	"github.com/ketianlin/ktools/slice"
	sqlT "github.com/ketianlin/ktools/sql"
	stringT "github.com/ketianlin/ktools/string"
	structT "github.com/ketianlin/ktools/struct"
	timeT "github.com/ketianlin/ktools/time"
	uuidT "github.com/ketianlin/ktools/uuid"
)

var (
	Crypto = new(crypto.Enter) // 加密工具
	Excel  = new(excel.Excel)
	File   = new(fileT.Enter)        // 文件工具
	Json   = new(jsonT.Enter)        // json工具
	Jwt    = new(jwt.Enter)          // jwt工具
	Map    = new(mapT.Enter)         // map工具
	Math   = new(math.Enter)         // 计算工具
	Slice  = new(slice.Enter)        // 切片工具
	SQL    = new(sqlT.Enter)         // sql工具
	String = new(stringT.Enter)      // 字符串工具
	Struct = new(structT.Enter[any]) // 结构体工具
	Time   = new(timeT.Enter)        // 时间工具
	UUID   = new(uuidT.Enter)        // uuid工具
)
