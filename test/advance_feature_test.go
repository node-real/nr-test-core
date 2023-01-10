package test

import (
	"github.com/node-real/nr-test-core/src/core/nrsuite"
)

type AdvanceSuiteTest struct {
	nrsuite.NRBaseSuite
}

//func TestAdvanceSuite(t *testing.T) {
//	nrsuite.Run(t, new(AdvanceSuiteTest))
//}

//func (t *AdvanceSuiteTest) Test_Func_Retry() {
//	t.RunFunWithRetry(func() error {
//		a := []int{}
//		b := rand.Intn(15)
//		fmt.Println(b)
//		if b > 3 {
//			fmt.Println(a[1])
//		}
//		return nil
//	}, 4)
//}

//// Tags:: $RetryCount:2
//func (t *AdvanceSuiteTest) Test_Case_Retry() {
//	a := []int{}
//	b := rand.Intn(15)
//	fmt.Println(b)
//	if b > 3 {
//		fmt.Println(a[1])
//	}
//}
