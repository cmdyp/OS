package Hardware

import (
	"GOOS/Global"
	"GOOS/Interrupt"
	"GOOS/OS"
	"fmt"
)

func execute() {
	Global.PC ++
	//取指令，执行指令
	if true/*中断发生*/{
		Global.BUS <- Interrupt.Interrupt{}
	}
}

func Run() {
	for {
		execute()
		select { //中断发生，陷入操作系统处理
		case interrupt := <- Global.BUS:
			if interrupt.Type > 0 { //内中断
				err := OS.INT(&interrupt)
				if err != nil {
					fmt.Println(err)
				}
			} else if interrupt.Type < 0 { //外中断
				err := OS.EINT(&interrupt)
				if err != nil {
					fmt.Println(err)
				}
			} else { //用户接口调用
				err := OS.UserInterface(interrupt.Command)
				if err != nil {
					fmt.Println(err)
				}
			}
		default:
			break
		}
	}
}
