package msgno

import (
	"context"
	"sync"
	"testing"
)

func BenchmarkMsgNoGen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateMsgNo()
	}
}

func TestMsgNoMiddleware(t *testing.T) {
	testEndpoint := func(ctx context.Context, req interface{}, resp interface{}) (err error) {
		t.Log(ctx.Value("msgno"))
		return nil
	}

	mw := MsgNoMiddleware(testEndpoint)
	ctx := context.Background()
	req := "test"
	resp := "OK"
	err := mw(ctx, req, resp)
	if err != nil {
		t.Errorf("MsgNoMiddleware failed: %+v", err)
	}
}

func TestGenerateMsgNo(t *testing.T) {
	var (
		wg  = &sync.WaitGroup{}
		nos = make(map[string]int)
		mu  = &sync.Mutex{}
	)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			no := generateMsgNo()
			mu.Lock()
			nos[no] += 1
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	for s, i := range nos {
		t.Logf("%s: %d", s, i)
	}
}
