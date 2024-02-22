package gterr

import (
	"log"
)

const (
	ErrPrefixFormat = "gtcli: (%s) %v"

	ErrBadParam    = "参数错误"
	ErrJsonParse   = "json解析失败"
	ErrMakeDir     = "创建文件夹失败"
	ErrCreateFile  = "创建文件失败"
	ErrProjExisted = "项目已存在"
	ErrModNotExist = "模块不存在"
	ErrProjCheck   = "项目检测失败"
)

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalf(ErrPrefixFormat, msg, err)
	}
}

func LogFatalf(msg, desc any) {
	log.Fatalf(ErrPrefixFormat, msg, desc)
}
