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

// rm -rf ~/go/src/github.com/ErosZy/AndroidLibV2ray
// go get github.com/ErosZy/AndroidLibV2ray
//rm -rf ~/go/src/v2ray.com && rm -rf ~/go/src/ErosZy/v2ray-core &&  make all