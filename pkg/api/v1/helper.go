package v1

import (
	"context"
	"fmt"

	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/pico"
)

func getPicoHandlerFromCtx(ctx context.Context) (*pico.Handler, error) {
	h := ctx.Value("pico")
	if h == nil {
		return nil, fmt.Errorf("Server context value for \"pico\" is missing (*pico.Handler)")
	}
	return h.(*pico.Handler), nil
}
