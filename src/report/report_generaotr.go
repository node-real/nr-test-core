package report

import (
	"encoding/json"
	"fmt"
	"github.com/node-real/nr-test-core/src/invokers/rpc"
	"github.com/node-real/nr-test-core/src/log"
)

type ReportOperator struct {
}

func (t *ReportOperator) BuildRpcError(msg *rpc.RpcMessage, actual, exp interface{}) string {
	m, _ := json.Marshal(msg)
	info := fmt.Sprintf("req: %s\nexp: %v\nactul: %v", m, exp, actual)
	log.Debug(info)
	return info
}

func (t *ReportOperator) BuildRpcErrorWithTrace(msg *rpc.RpcMessage, actual, exp interface{}, traceID string) string {
	m, _ := json.Marshal(msg)
	info := fmt.Sprintf("req: %s\nexp: %v\nactul: %v\ntraceId: %s", m, exp, actual, traceID)
	log.Debug(info)
	return info
}

func (t *ReportOperator) BuildRpcBatchError(msg []*rpc.RpcMessage, actual, exp interface{}) string {
	m, _ := json.Marshal(msg)
	info := fmt.Sprintf("req: %s\nexp: %v\nactul: %v", m, exp, actual)
	log.Debug(fmt.Sprintf("req: %s\n===exp: %v\n===actul: %v", m, exp, actual))
	return info
}

func (t *ReportOperator) BuildRpcBatchErrorWithTrace(msg []*rpc.RpcMessage, actual, exp interface{}, traceID string) string {
	m, _ := json.Marshal(msg)
	info := fmt.Sprintf("req: %s\nexp: %v\nactul: %vtraceId: %s", m, exp, actual, traceID)
	log.Debug(fmt.Sprintf("req: %s\n===exp: %v\n===actul: %v", m, exp, actual))
	return info
}

//func (t *ReportOperator) AlertToSlack(hookStr string) {
//	//TODO:
//}
