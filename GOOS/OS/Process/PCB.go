package Process
//维护PCB以及恢复进程状态
import (
	"GOOS/Global"
	"GOOS/OS/MemmorySystem"
	"errors"
)

//*要求全局性的修改,包括修改PCB内部信息和PC，外层不会再提供任何维护*//


type PCB struct { //不需要状态字，因为系统已经通过全局变量维护了,pid就是数组索引，也不维护
	Allocated int
	Ppid int
	PC int //执行到代码的逻辑位置，当save的时候会更新
	RegisterState []int //寄存器值，当save的时候会更新
	File []int //文件系统（未明确指定）
	PageTable []Global.PageAtom //这里把页表存下来，真实的操作系统只会存一个指针,页表常驻内存，当swap的时候会更新，当前PC位置禁止被换出，防止影响load的原子操作
	PageSiz int //len(PageTable)
	Disk int //程序存放的外存地址，一经确认不在改变，虽然UNIX系统分了交换区和文件区
}

func LogicPC(PC int, pageTable *[]Global.PageAtom) (myPC int){ //得到逻辑位置
	return 0
}

func PhysicPC(PC int, pageTable *[]Global.PageAtom) (myPC int){ //得到物理位置
	return 0
}

func (this *PCB)Save() (err error) {
	this.PC = LogicPC(Global.PCSave[len(Global.PCSave) - 1], &this.PageTable)
	//保存寄存器值
	this.RegisterState = Global.RegisterState[len(Global.RegisterState)- 1]
	return
}

func (this *PCB)Load() (err error) {
	Global.PCSave[len(Global.PCSave) - 1] = LogicPC(this.PC, &this.PageTable)
	//修改寄存器
	//各个寄存器的值等于 this.RegisterState
	return
}

func (this *PCB)Delete() (err error) {
	this.Allocated = 0
	err = MemmorySystem.MemoryRelease(&this.PageTable) //内存清除
	return
}

func SortReady() {

}  //这里给出进程排序算法，FCFS SJF/SRTF HRRN RR queue

func SwapPage(pidin int, pidout int, PhysicPageNum int, logicPageNum int) (err error) {//将PhysicPageNum处的pidout进程的页给pidin的logicPageNum

	return
}

func CompareWithRunning(pid int) (flag bool) {
	return false
}

func Create(size int) (pid int, err error) { //返回一个可用PCB的编号，整个过程是New
	pid = len(Global.PCBs)
	for i := 1 ; i < len(Global.PCBs) ; i ++ {
		if Global.PCBs[i].Allocated == 0 {
			pid = i
			goto _BREAK
		}
	}
_BREAK:
	Global.PCBs[pid].Allocated = 1
	Global.PCBs[pid].PageTable, err = MemmorySystem.MemoryAllocate(size)  //在分配内存的时候，只会分配程序需要的一部分内存，也就是程序起始模块必须被装入
	if err != nil { //一般不会发生错误，只有磁盘容量都不满足的时候发生
		return
	}
	Global.PCBs[pid].PageSiz = size
	Global.PCBs[pid].PC = 0
	err = errors.New("there is no empty PCBs")
	return
}