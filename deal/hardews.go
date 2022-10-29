/**
 * @Author: Hardews
 * @Date: 2022/10/30 1:37
**/

// 这个包就用来声明一些处理信息返回的函数，文件命名就自己定。
// 比如hardews.go就是处理我发送的信息的文件,
// 防止冲突，每个函数前应带文件名前三个字符

package deal

import "qq-bot/model"

func HarDealWithMsg(msg model.Message) string {
	return "test"
}
