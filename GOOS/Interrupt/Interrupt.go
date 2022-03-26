package Interrupt

const (
	//用户接口
	USERINTERFACE = 0
//内中断
	STDIN = 1 //标准输入
	STDOUT = 2 //标准输出
	TIMEOUT = 3 //超时中断
	NEWPROCESS = 4 //新建进程
	ABORTPROCESS = 5 //进程终止
	PAGEFAULT = 6 //缺页
	DIVIDEBY0 = 7 //除0
	READFILE = 8 //读文件
	WRITEFILE = 9 //写文件
	MEMORYAPPLY = 10 //内存申请
	NOMEM = 11 //存储不足，虚拟内存不足（磁盘空间不足）
	MEMIN = 12 //读内存
	MEMOUT = 13 //写内存
	UNDEFINE1 = 16 //保留
	UNDEFINE2 = 17

//外中断
	E_STDIN = -1 //标准输入完成
	E_STDOUT = -2 //标准输出完成
	E_HARDERROR = -3 //硬件错误
	E_MEMIN = -4 //读内存
	E_MEMOUT = -5 //写内存
	E_UNDEFINE1 = -6 //保留
	E_UNDEFINE2 = -7
)

type Interrupt struct {
	Type int //中断类型
	Priority int //优先级
	Pid int //请求/返回中断的进程id
	FilePath string
	Device string
	Command string
	others interface{} //可以不断完善interrupt
}



