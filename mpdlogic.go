package plutabot

import (
	"log"
	"fmt"

	"github.com/fhs/gompd/mpd"
)

func MpdGetCurrent() string {
	conn, err := mpd.Dial("tcp", "pluta:6600")
	if err != nil {
		    log.Fatalln(err)
	}
	defer conn.Close()

	status, err := conn.Status()
	if err != nil {
		log.Fatalln(err)
	}
	song, err := conn.CurrentSong()
	if err != nil {
		log.Fatalln(err)
	}


	line1 := ""
	if status["state"] == "play" {
		line1 = fmt.Sprintf("%s - %s", song["Artist"], song["Title"])
	} else {
		line1 = fmt.Sprintf("State: %s", status["state"])
	}
	return line1;
}


