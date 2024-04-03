package main

import (
	"context"
)

func handler(ctx context.Context, event map[string]any) error {
	a, err := arrange(event)
	if err != nil {
		return err
	}

	res, err := act(a)
	if err != nil {
		return err
	}

	return assert(res)

}
