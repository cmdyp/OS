package OS

import (
	"GOOS/Global"
	"GOOS/Interrupt"
	"GOOS/OS/Process"
)

//中断处理程序，也是操作系统的核心，这两个可以写成一个函数其实


func INT(interrupt *Interrupt.Interrupt) (err error){ //内部中断处理，处理系统调用，错误和中止
	Global.PCSave = append(Global.PCSave, Global.PC) //PC立即保存
	err = Process.INTExecute(interrupt) //处理中断
	Global.PC = Global.PCSave[len(Global.PCSave) - 1] //中断返回
	Global.PCSave = Global.PCSave[:len(Global.PCSave) - 1]
	return
}

func UserInterface(command string) (err error) { //用户接口

	return
}

func EINT(interrupt *Interrupt.Interrupt) (err error) { //外部中断响应
	Global.PCSave = append(Global.PCSave, Global.PC) //PC立即保存
	err = Process.EINTExecute(interrupt) //处理中断
	Global.PC = Global.PCSave[len(Global.PCSave) - 1] //中断返回
	Global.PCSave = Global.PCSave[:len(Global.PCSave) - 1]
	return
}

