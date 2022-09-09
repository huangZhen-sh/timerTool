<h3>简介</h3>

基于go的定时器，实现类似js的setTimeout、setInterval函数功能

<h3>安装</h3>

````
go get github.com/huangZhen-sh/timerTool
````

<h3>使用方法</h3>

````
type TestTimerAction struct {
}
func (a TestTimerAction) Action() {
	fmt.Printf("时间到了....\r\n")
}
func main() {
	a := TestTimerAction{}
	timer := timerTool.SetInterval(a, 1*time.Second)
	time.Sleep(5 * time.Second)
	timer.Stop()
	for {
	}
}
````