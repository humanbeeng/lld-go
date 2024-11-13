package main

type Notifier interface {
	Notify() error
}
