package Process
//原语，不可中断
import (
	"GOOS/Global"
	"GOOS/Interrupt"
	"errors"
)

//只修改队列情况，内部不做处理

func Abort(pid int) (err error){
	if pid == Global.RunningPID {
		err = Global.PCBs[pid].Delete()
		Global.RunningPID = -1
		if len(Global.ReadyQueue) > 0 {
			SortReady()
			Global.RunningPID = Global.ReadyQueue[0]
			Global.ReadyQueue = Global.ReadyQueue[1:]
			err = Global.PCBs[Global.RunningPID].Load() //加载新的PCB
		}
	} else if true/*判断是否在ready队列中*/ {

	} else if true/*判断是否在waiting队列中*/ {

	} else if true/*判断是否在suspend队列中*/ {
		
	} else {
		err = errors.New("process doesn't exist")
	}
	return
}

func Switch() (err error){ //切换原语，当前running进程进入ready并从ready中拿出一个进程
	pid := Global.RunningPID
	err = Global.PCBs[pid].Save() //保存PCB
	Global.ReadyQueue = append(Global.ReadyQueue, pid)
	Global.RunningPID = -1
	if len(Global.ReadyQueue) > 0 {
		SortReady()
		Global.RunningPID = Global.ReadyQueue[0]
		Global.ReadyQueue = Global.ReadyQueue[1:]
		err = Global.PCBs[Global.RunningPID].Load() //加载新的PCB
	}
	return
}

func Waiting(waitWhat int) (err error){ //阻塞原语，当前running进程发生阻塞（请求设备），此时把它放到阻塞队列中,选择一个ready中的执行
	pid := Global.RunningPID
	err = Global.PCBs[pid].Save() //保存PCB
	Global.WaitingQueue = append(Global.WaitingQueue, Global.PII{pid, waitWhat})
	Global.RunningPID = -1
	if len(Global.ReadyQueue) > 0 {
		SortReady()
		Global.RunningPID = Global.ReadyQueue[0]
		Global.ReadyQueue = Global.ReadyQueue[1:]
		err = Global.PCBs[Global.RunningPID].Load() //加载新的PCB
	}
	return
}

func Awake(pid int, waitWhat int) (err error) { //唤醒某个waiting中的进程到ready中，再查看是否发生switch
	for i := 0 ; i < len(Global.WaitingQueue) ; i ++ {
		if Global.WaitingQueue[i].First == pid && Global.WaitingQueue[i].Second == waitWhat {
			Global.ReadyQueue = append(Global.ReadyQueue, pid)
			Global.WaitingQueue = append(Global.WaitingQueue[:i],Global.WaitingQueue[i+1:]...)
			SortReady()
			if CompareWithRunning(Global.ReadyQueue[0]) { //awake后造成切换
				err = Switch()
			}
			return
		}
	}
	err = errors.New("there is no process in waiting queue")
	return
}

func New(size int) (err error) { //新建一个进程，为其分配空间，再查看是否发生switch
	pid , err := Create(size)
	if err != nil {
		Global.SuspendReadyQueue = append(Global.SuspendReadyQueue, Global.PII{pid, Interrupt.NOMEM})
	} else {
		Global.ReadyQueue = append(Global.ReadyQueue, pid)
		SortReady()
		if CompareWithRunning(Global.ReadyQueue[0]) { //awake后造成切换,和当前运行进程比一比
			err = Switch()
		}
	}
	return
}
