package main

import (
	"log"
	"net/url"

	"github.com/zhoukk/krtmpd"
)

func main() {
	rs := krtmpd.NewRtmpServer(":1935")
	rs.SetHook(krtmpd.RtmpdHook{
		PushStart: func(s string, u *url.URL) bool {
			log.Printf("%s push %s start\n", s, u.String())
			return true
		},
		PushEnd: func(s string, u *url.URL) {
			log.Printf("%s push %s end\n", s, u.String())
		},
		PullStart: func(s string, u *url.URL) bool {
			log.Printf("%s pull %s start\n", s, u.String())
			return true
		},
		PullEnd: func(s string, u *url.URL) {
			log.Printf("%s pull %s end\n", s, u.String())
		},
	})
	rs.Start()
}
