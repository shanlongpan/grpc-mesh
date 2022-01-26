/**
* @Author:Tristan
* @Date: 2022/1/18 3:16 下午
 */

package lib

import (
	"math/rand"
	"time"
)

func RandomInt(i int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(i)
}

func RandomInt64(i int64) int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(i)
}
