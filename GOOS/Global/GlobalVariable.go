package Global
//全局变量放在这
import (
	"GOOS/Interrupt"
	"GOOS/OS/Process"
)

type PII struct {
	First int
	Second int
}

//页号、【==内存块号、状态位（是否在内存）、访问字段（记录这个页面被访问几次或最近访问时间）、修改位（页面调入内存后是否被修改过）、外存地址（页面在外存中的位置）==】
type PageAtom struct {
	Mem int
	Stat bool
	Frequency int
	IsChange bool
	Disk int
}

var PCBs []Process.PCB //PCB常驻内存

var ReadyQueue []int //就绪队列

var RunningPID int //-1则表示没有在运行中的进程

var WaitingQueue []PII //阻塞队列

var BUS = make(chan Interrupt.Interrupt, 100000) //模拟中断信号总线，往里面放中断信号

var PC int //程序计数器，存储实际内存位置

var PCSave []int //程序计数器暂存，中断处理程序时候立即触发，多级中断这里使用栈实现

var RegisterState [][]int //寄存器保留值，多级中断这里使用栈实现

var SuspendReadyQueue []PII //这个阻塞几乎不会发生，如果磁盘空间也不够，触发等待（不实现）

var OpenFileTable []interface{} //打开文件表常驻内存

var FileTable []interface{} //一级文件索引表常驻内存