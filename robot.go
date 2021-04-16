package main

import (
	"log"
	"os/exec"
	"time"
)

func main() {
 var host string = "安夏"
 var sentence = []string{
  "'欢迎来到夏夏520趴🔥 点关注不迷路✨ 今天凑齐百人榜💞动动小手榜上住🏠'",
  `'夏夏520趴，夏夏正在参加PK排位赛，有礼物的小耳朵们🎁支持下，感谢❤️'`,
  `'╭┈┈✨🏅✨┈┈╮
  🔥 安夏520狂欢趴🔥
 ╰┈┈🌸 ️🏆🌸 ️┈'`,
  `'☀️Good afternoon☀️
  安夏✨今日520🎏
  感恩遇见❤'`,
  "'✨安夏520狂欢趴🌸✨春风十里不如你🌸✨我在等风也等你🌸✨请点主播关注'",
  "'浮世万千，吾爱有三，一为日二为月三为卿，日为朝，月为暮，" + host + "陪你朝朝暮暮'",
 }
 var interval = 13 * time.Second
 for i := 1; i < 10000; i++ {
  log.Println("第", i, "次场控")
  ck(sentence, i, interval)
 }
}
func ck(sentence []string, index int, interval time.Duration) {
 sentenceIndex := index % len(sentence)
 arg1 := []string{
  "shell", "input", "tap", "85", "2000",
 }
 arg3 := []string{
  "shell", "input", "tap", "1015", "1950",
 }
 arg2 := []string{
  "shell", "am", "broadcast", "-a", "ADB_INPUT_TEXT", "--es", "msg",
 }
 arg2 = append(arg2, sentence[sentenceIndex])
 cmd1 := exec.Command(`D:\P\platform-tools\adb.exe`, arg1...)
 cmd2 := exec.Command(`D:\P\platform-tools\adb.exe`, arg2...)
 cmd3 := exec.Command(`D:\P\platform-tools\adb.exe`, arg3...)
 cmd1.Run()
 time.Sleep(interval)
 cmd2.Run()
 time.Sleep(interval)
 cmd3.Run()
 time.Sleep(interval)
}
