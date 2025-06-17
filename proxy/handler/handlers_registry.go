package handler

var HandlersRegistry = map[string]func([]byte) (Handler, error){}
