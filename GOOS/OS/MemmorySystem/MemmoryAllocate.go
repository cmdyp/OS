package MemmorySystem

import "GOOS/Global"

//内存申请，采用虚拟页式存储，这个err是总存储量不足的时候发生错误
//给出需要多少页，返回一个page表，或者根据page表回收，需要在这里写出内存置换算法LRU FIFO CLOCK等
func MemoryAllocate(size int) (pageTable []Global.PageAtom, err error){

	return
}

func MemoryRelease(pageTable *[]Global.PageAtom) (err error) {

	return
}


func MemoryRead(physicPageNum int) (err error){ //请求写内存,给出位置

	return
}

func MemoryWrite(physicPageNum int) (err error){ //请求读内存

	return
}

func SwapInOut(physicPageNumIn int, physicPageNumOut int) (err error) { //内存换入换出

	return
}