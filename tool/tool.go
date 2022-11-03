/**
 * @Author: Hardews
 * @Date: 2022/10/31 20:23
**/

package tool

import (
	"math/rand"
	"strconv"
	"time"
)

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func RandNum(n int) int {
	rand.Seed(time.Now().UnixMilli())
	return rand.Intn(n)
}
