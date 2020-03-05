// +build darwin linux

package main

import (
	"encoding/json"
	"fmt"
	"github.com/andrewarrow/feedback/api"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/exp/gl/glutil"
	//	"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/gl"
)

var display = "jjaa.me"
var displayIndex = 0
var displayItems = []string{"jjaa.me", "A", "B", "C", "D", "E", "F", "G"}
var flavor = 1

func main() {
	rand.Seed(time.Now().UnixNano())
	//go hitApi()

	app.Main(func(a app.App) {
		var glctx gl.Context
		var sz size.Event
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					glctx, _ = e.DrawContext.(gl.Context)
					onStart(glctx)
					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					onStop()
					glctx = nil
				}
			case size.Event:
				sz = e
			case paint.Event:
				if glctx == nil || e.External {
					continue
				}
				if rand.Intn(100) > 50 {
					onPaint(glctx, sz)
					a.Publish()
					a.Send(paint.Event{})
				}
			case touch.Event:
				if down := e.Type == touch.TypeBegin; down || e.Type == touch.TypeEnd {
					game.Touch(down, e.X, e.Y, sz)
				}
			case key.Event:
				if e.Code != key.CodeSpacebar {
					break
				}
				if down := e.Direction == key.DirPress; down || e.Direction == key.DirRelease {
					game.Touch(down, 0, 0, sz)
				}
			}
		}
	})
}

func hitApi() {
	host := "jjaa.me"
	data, err := http.Get(fmt.Sprintf("https://%s/api/version", host))
	if err != nil {
		return
	}
	all, _ := ioutil.ReadAll(data.Body)
	data.Body.Close()

	var ar api.ApiResponse
	json.Unmarshal(all, &ar)
	display = fmt.Sprintf("%v", ar.SentAt)
}

var (
	startTime = time.Now()
	images    *glutil.Images
	game      *Game
)

func onStart(glctx gl.Context) {
	images = glutil.NewImages(glctx)
	game = NewGame()
}

func onStop() {
	images.Release()
	game = nil
}

func onPaint(glctx gl.Context, sz size.Event) {
	if flavor == 1 {
		glctx.ClearColor(0, rand.Float32(), 0, 1)
	} else if flavor == 2 {
		glctx.ClearColor(0, 0, rand.Float32(), 1)
	} else if flavor == 3 {
		glctx.ClearColor(rand.Float32(), 0, 0, 1)
	} else if flavor == 4 {
		glctx.ClearColor(rand.Float32(), rand.Float32(), 0, 1)
	} else if flavor == 5 {
		glctx.ClearColor(rand.Float32(), 0, rand.Float32(), 1)
	} else if flavor == 6 {
		glctx.ClearColor(0, rand.Float32(), rand.Float32(), 1)
	} else if flavor == 7 {
		glctx.ClearColor(rand.Float32(), rand.Float32(), rand.Float32(), 1)
	}
	glctx.Clear(gl.COLOR_BUFFER_BIT)
	game.Render(sz, glctx, images)
	/*
		glctx.ClearColor(1, 1, 1, 1)
		glctx.Clear(gl.COLOR_BUFFER_BIT)
		now := clock.Time(time.Since(startTime) * 60 / time.Second)
		game.Update(now)
		game.Render(sz, glctx, images)*/
}
