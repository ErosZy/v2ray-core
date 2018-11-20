package export

type UDPBufferExport interface {
	OnData(data []byte) int
}

var exportor UDPBufferExport = nil

func SetUDPBufferExport(e UDPBufferExport) {
	exportor = e
}

func GetUDPBufferExport() UDPBufferExport {
	return exportor
}