package middleware

import "adoptGolang/internal/engine/controller"

type Middleware struct {
	*controller.Controller
}

func NewMiddleware(_controller *controller.Controller) *Middleware {
	return &Middleware{ Controller: _controller }
}
