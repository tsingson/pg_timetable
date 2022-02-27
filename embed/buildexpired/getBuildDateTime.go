package buildexpired

/*
// C 标志io头文件，你也可以使用里面提供的函数
#include <stdio.h>

const char* build_time(void)
{
    static const char* psz_build_time = __DATE__ " " __TIME__ ;
    return psz_build_time;
}

*/
import "C" // 切勿换行再写这个

import (
	"time"
)

const longForm = "Jan 2 2006 15:04:05"
const Expired = time.Hour * 24 * 7 * 26

func BuildDateTime() (time.Time, error) {
	var buildTime = C.GoString(C.build_time())
	return time.Parse(longForm, buildTime)
}
