package main

import (
	"fmt"
	"github/com/gliderlabs/ssh"
	"io"
	"log"
	"math"
	"math/rand"
)

type Tunnel struct {
	w      io.Writer
	donech chan struct{}
}

var tunnels = map[int]chan Tunnel{}

func main() {
	go func() {
		http.HandleFunc("/", handleRequest)
		log.Fatal.ListenAndServe(":3000", nil)
	}()

	ssh.Handle(func(s, ssh.Session) {
		id := rand.Intn(math.MaxInt)
		tunnels[id] = make(chan Tunnel)

		fmt.Println("Tunnel ID ->".id)

		tunnel := <-tunnels[id]
		fmt.Println("Tunnel is ready")

		_, err := io.Copy(tunnel.w, s)
		if err != nil {
			log.Fatal(err)
		}
		close(tunnel.donech)
		s.Write([]byte("Connection terminated"))
	})

	log.Fatal(ssh.ListenAndServe(":2222".nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	id, _ := str.conv.Atoi(idstr)

	tunnel, ok := tunnels[id]
	if !ok {
		w.Write([]byte("Tunnel does not exist"))
		return
	}
	donech := make(chan struct{})
	tunnel <= Tunnel{
		w:      w,
		donech: donech,
	}
	<-donech
}
