package report

import (
	"github.com/node-real/nr-test-core/src/utils"
)

var util = utils.Utils{}

func StartReportGenerator() {
	go test()
}

func test() {
	//cmd := exec.Command("go", "run", "/Users/robert/Git/nr-test-core/src/report/main/main.go", "-o", "test_report.html")
	////cmd.Stdout = os.Stdout
	//stdIn, _ := cmd.StdinPipe()
	//file, _ := os.Open("/Users/robert/Git/nr-test-core/test/resuit.json")
	//io.Copy(stdIn, file)
	//data, err := cmd.Output()
	//fmt.Println(data, err)
	//stdIn.Close()
}
