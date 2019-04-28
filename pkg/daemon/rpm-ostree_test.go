package daemon

type GetBootedOSImageURLReturn struct {
	OsImageURL	string
	Version		string
	Error		error
}
type RpmOstreeClientMock struct {
	GetBootedOSImageURLReturns	[]GetBootedOSImageURLReturn
	RunPivotReturns			[]error
}

func (r RpmOstreeClientMock) GetBootedOSImageURL(string) (string, string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	returnValues := r.GetBootedOSImageURLReturns[0]
	if len(r.GetBootedOSImageURLReturns) > 1 {
		r.GetBootedOSImageURLReturns = r.GetBootedOSImageURLReturns[1:]
	}
	return returnValues.OsImageURL, returnValues.Version, returnValues.Error
}
func (r RpmOstreeClientMock) RunPivot(string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := r.RunPivotReturns[0]
	if len(r.RunPivotReturns) > 1 {
		r.RunPivotReturns = r.RunPivotReturns[1:]
	}
	return err
}
func (r RpmOstreeClientMock) GetStatus() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "rpm-ostree mock: blah blah some status here", nil
}
