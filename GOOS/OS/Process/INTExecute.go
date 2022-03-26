package Process
//处理内部中断
import (
	"GOOS/Global"
	"GOOS/Interrupt"
	"GOOS/OS"
	"errors"
	"fmt"
)

func INTExecute(interrupt *Interrupt.Interrupt) (err error){
	//保存所有寄存器信息 到指定空间
	saveState()
	//处理中断
	switch interrupt.Type {
	case Interrupt.STDIN:
		stdin()
	case Interrupt.STDOUT:
		stdout()
	case Interrupt.NEWPROCESS:
		newProcess()
	case Interrupt.MEMORYAPPLY:
		memoryApply(0)
	case Interrupt.READFILE:
		readFile("")
	case Interrupt.WRITEFILE:
		writeFile("")
	case Interrupt.PAGEFAULT:
		swap(0)
	case Interrupt.MEMIN:

	case Interrupt.MEMOUT:

	case Interrupt.ABORTPROCESS:

	case Interrupt.UNDEFINE1://用户中断，该中断可以被中断
		run(interrupt)
	default:
		err = errors.New("int error")
	}
	return
}

func stdin() { //请求标准输入
	temp := Global.PCBs[Global.RunningPID]
	err := Waiting(Interrupt.STDIN);
	if err != nil {
		fmt.Println(err)
		err = temp.Load()
	}
}

func stdout() { //请求标准输出
	temp := Global.PCBs[Global.RunningPID]
	err := Waiting(Interrupt.STDOUT);
	if err != nil {
		fmt.Println(err)
		err = temp.Load()
	}
}

func newProcess() { //新建作业/程序
	temp := Global.PCBs[Global.RunningPID]
	u := 1024 //程序一开始只有部分被装入到内存，之后可以发生缺页中断进行调用，假定分配1024页空间
	err := New(u)
	if err != nil {
		fmt.Println(err)
		err = temp.Load()
	}
}

func memoryApply(size int){ //当前进程申请内存资源

}

func readFile(name string) { //当前进程读文件中断

}

func writeFile(name string) { //当前进程写文件中断

}

func swap(loGicPageNum int) { //当前进程缺页中断，中断中给出中断页，执行置换算法，调用PCB内部的Swap

}

func run(interrupt *Interrupt.Interrupt){
	for i := 0 ; i < 10 ; i ++ {
		select {
		case inInterrupt := <-Global.BUS:
			if (inInterrupt.Priority < interrupt.Priority) {
				if interrupt.Type > 0 { //内中断
					err := OS.INT(&inInterrupt)
					if err != nil {
						fmt.Println(err)
					}
				} else if interrupt.Type < 0 { //外中断
					err := OS.EINT(&inInterrupt)
					if err != nil {
						fmt.Println(err)
					}
				} else { //用户接口调用
					err := OS.UserInterface(interrupt.Command)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
}