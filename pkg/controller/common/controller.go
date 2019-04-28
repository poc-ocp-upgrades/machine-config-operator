package common

type Controller interface {
	Run(workers int, stopCh <-chan struct{})
}
