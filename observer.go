package monolith

import "github.com/amaury95/monolith/observer"

type IObserver[T comparable] interface {
	Subscribe(t T, p observer.Processor[T]) observer.SourceID
	Unsubscribe(s observer.SourceID)
	Emit(e *observer.Event[T]) error
}

//go:generate mockgen --destination=./mocks/observer.go --package=mocks github.com/amaury95/monolith Observer
type Observer = IObserver[string]
