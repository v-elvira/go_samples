package main

import "fmt"

// CANCEL
func work(ctx context.Context) error {
	done := make(chan struct{})

	go func() {
		defer close(done)
		fmt.Println("do stuff")
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	work(ctx)
}

//--------------------------------------------------
// TIMER
ctx := context.Background()
ctx, cancel := context.WithCancel(ctx)
ctx, cancel = context.WithTimeout(ctx, 3*time.Second)
defer cancel()
work(ctx)

//--------------------------------------------------
//DEADLINE
ctx := context.Background()
ctx, cancel := context.WithCancel(ctx)
deadline := time.Now().Add(3 * time.Second)
ctx, cancel = context.WithDeadline(ctx, deadline)
defer cancel()
work(ctx)

//--------------------------------------------------
// error CAUSE
ctx, cancel := context.WithCancelCause(context.Background())
cancel(errors.New("the night is dark"))
fmt.Println(context.Cause(ctx))
// the night is dark


//--------------------------------------------------
// AFTER (callback)
ctx, cancel := context.WithCancel(context.Background())
cleanup := func() { fmt.Println("cleanup") }
context.AfterFunc(ctx, cleanup)
cancel()
// cleanup
