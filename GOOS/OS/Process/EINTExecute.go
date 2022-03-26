package Process
//处理外部中断
import (
	"GOOS/Global"
	"GOOS/Interrupt"
	"errors"
	"fmt"
)

func EINTExecute(interrupt *Interrupt.Interrupt) (err error){
	//保存所有寄存器信息 到指定内存空间Global.RegisterState
	saveState()
	//处理中断
	switch interrupt.Type {
	case Interrupt.E_STDIN:
		estdin()
	case Interrupt.E_STDOUT:
		estdout()
	case Interrupt.E_MEMIN:

	case Interrupt.E_MEMOUT:

	default:
		err = errors.New("eint error")
	}
	//恢复环境，
	loadState()
	return
}

func saveState() {
	Global.RegisterState = append(Global.RegisterState, []int{1, 2, 3})
}

func loadState(){
	//寄存器=Global.RegisterState
	Global.RegisterState = Global.RegisterState[:len(Global.RegisterState) - 1]
}

func estdin() { //标准输入返回
	temp := Global.PCBs[Global.RunningPID]
	err := Awake(Global.RunningPID, Interrupt.E_STDIN)
	if err != nil {
		fmt.Println(err)
		err = temp.Load()
	}
}

func estdout() { //标准输出返回
	temp := Global.PCBs[Global.RunningPID]
	err := Awake(Global.RunningPID, Interrupt.E_STDOUT)
	if err != nil {
		fmt.Println(err)
		err = temp.Load()
	}
}
