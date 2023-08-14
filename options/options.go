package options

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type Options func()

type FallbackHandlerInterface interface {
	func(ctx *gin.Context) | func(ctx echo.Context) error
}

type FallBackHandlerType interface {
	*gin.Context | *echo.Context
}

// WhiteLists
// A list of whitelisting ip addresses
var (
	WhiteLists                 []string
	FallBackHandler            interface{}
	SuccessRedirectionCallback interface{}
)

// set the WhiteLists
func WithWhiteLists(ips []string) Options {
	return func() {
		WhiteLists = ips
	}
}

// set the FallbackHandler
func WithFallbackHandler[T FallbackHandlerInterface](handler T) Options {
	return func() {
		FallBackHandler = handler
	}
}

func TriggerFallbackHandler[T FallbackHandlerInterface, C FallBackHandlerType](ctx C) {
	if fn, ok := FallBackHandler.(T); ok {
		fmt.Printf("fn--- %T", fn)
	}
}
