package testdata

type DataOperator struct {
}

func (dataOp *DataOperator) GetOnlineData(key string) string {
	//TODO: Robert
	return ""
}

func (dataOp *DataOperator) UploadOnlineData(key string) bool {
	//TODO: Robert
	return false
}

func (dataOp *DataOperator) ReadCsvData() []string {
	//TODO: Robert
	return nil
}

func (dataOp *DataOperator) ReadCustomFileData() map[string]string {
	//TODO:Robert
	return nil
}

func (dataOp *DataOperator) GenerateGql() string {
	//TODO: Robert
	return ""
}
