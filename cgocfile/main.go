package main

import (
	"fmt"

	"github.com/panyingyun/cgoguide/cgocfile/packet"
	"github.com/panyingyun/cgoguide/cgocfile/utils"
)

func main() {
	fmt.Println("111")
	sum := packet.Add(1, 2)
	fmt.Println("sum = ", sum)
	sum2 := utils.GoSum(4, 5)
	fmt.Println("sum2 = ", sum2)
	fmt.Println("222")
}
