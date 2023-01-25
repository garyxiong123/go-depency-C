package B

import (
	A "github.com/garyxiong123/go-depency-A"
	"github.com/garyxiong123/go-depency-B/B"
)

func Go_dependency_C_V1() {
	B.Go_dependency_B_V1()

	A.Go_dependency_A_V2()
	println("go_dependency_C_V1")
}
