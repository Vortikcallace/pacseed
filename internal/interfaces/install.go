package interfaces

type PackageManager interface {
	Install(name string) error
}
