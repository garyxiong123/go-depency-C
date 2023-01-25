package B

import (
	"github.com/garyxiong123/go-depency-A/A2"
	"github.com/garyxiong123/go-depency-B/B"
)

func Go_dependency_C_V1() {
	B.Go_dependency_B_V1()

	A2.Go_dependency_A_V2()
	println("go_dependency_C_V1")
}
