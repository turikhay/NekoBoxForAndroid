package libcore

var intfBox BoxPlatformInterface
var intfNB4A NB4AInterface

var useProcfs bool

type NB4AInterface interface {
	UseOfficialAssets() bool
	Selector_OnProxySelected(tag string)
}

type LocalResolver interface {
	LookupIP(network string, domain string) (string, error)
}

var localResolver LocalResolver // Android: passed from java (only when VPNService)

type BoxPlatformInterface interface {
	AutoDetectInterfaceControl(fd int32) error
	OpenTun(singTunOptionsJson, tunPlatformOptionsJson string) (int, error)
	UseProcFS() bool
	FindConnectionOwner(ipProtocol int32, sourceAddress string, sourcePort int32, destinationAddress string, destinationPort int32) (int32, error)
	PackageNameByUid(uid int32) (string, error)
	UIDByPackageName(packageName string) (int32, error)
}

func SetBoxPlatformInterface(iif BoxPlatformInterface) {
	intfBox = iif
	useProcfs = intfBox.UseProcFS()
}
