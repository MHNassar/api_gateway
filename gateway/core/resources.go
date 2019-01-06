package core

import "time"

type JsonRoot struct {
	Router Router
}

type Router struct {
	Port     string
	Settings Settings
	Services []Services
}

type Settings struct {
	TimeOut time.Duration
}

type Services struct {
	ServicePrefix string
	TargetPath    TargetPath
}

type TargetPath struct {
	Path string
	Auth bool
}
