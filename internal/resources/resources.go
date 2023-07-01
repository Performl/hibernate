package resources

// deployments, statefulsets
// both implement this interface
type Resources interface {
	Sleep()
	Wake()
}
