package pass_through

type passRequest string

const (
	versionRequest passRequest = "/version::get"
)

var passThroughMap = map[passRequest]bool{
	versionRequest: true,
}

func IsPassThroughPath(path, verb string) bool {
	return passThroughMap[passRequest(path+"::"+verb)]
}
