package agent

import (
	"log"
	"net"

	types "k0s.io/pkg/agent"
	"k0s.io/pkg/agent/tty"
	asciitransport "k0s.io/pkg/asciitransport/v2"
)

func StartTerminalV2Server(c types.Config) chan net.Conn {
	var (
		cmd              []string = c.GetCmd()
		ro               bool     = c.GetReadOnly()
		fac                       = tty.NewFactory(cmd)
		terminalListener          = NewLys()
	)
	_ = ro
	go serveTerminalV2(terminalListener, fac)
	return terminalListener.Conns
}

func serveTerminalV2(ln net.Listener, fac types.TtyFactory) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			term, err := fac.MakeTty()
			if err != nil {
				log.Println(err)
				return
			}

			opts := []asciitransport.Opt{
				asciitransport.WithReader(term),
				asciitransport.WithWriter(term),
				asciitransport.WithResizeHook(func(w, h uint16) {
					err := term.Resize(int(w), int(h))
					if err != nil {
						log.Println(err)
					}
				}),
			}
			server := asciitransport.Server(conn, opts...)
			println("new server")
			<-server.Done()
			println("done")
		}()
	}
}
