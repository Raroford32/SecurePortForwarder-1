package portforward

import (
	"io"
	"net"
	"sync"

	"ipsec-port-forward/internal/utils"
)

func Forward(src, dst net.Conn) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if _, err := io.Copy(dst, src); err != nil {
			utils.LogError("Error forwarding src->dst", err)
		}
		dst.Close()
	}()

	go func() {
		defer wg.Done()
		if _, err := io.Copy(src, dst); err != nil {
			utils.LogError("Error forwarding dst->src", err)
		}
		src.Close()
	}()

	wg.Wait()
}