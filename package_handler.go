package main

type packageHandler interface {
	isPackage(filename string) bool
	getPackageInfo(filePath string) (packageInfo, error)
	changePackageVersion(file packageInfo, newVersion string) error
	applyVersion(byteValue []byte, newVersion string) ([]byte, error)
}
