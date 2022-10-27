package data

const (
	CaseDescFlag = "CaseDesc"
	CaseEnd      = "CaseEnd"
)

type CustomCaseData struct {
	CaseDesc  string
	CaseInfos map[string]string
}
