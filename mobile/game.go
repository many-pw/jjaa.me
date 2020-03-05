// +build darwin linux

package main

import (
	"image"
	"image/color"
	"log"
	"strings"
	"time"

	"github.com/golang/freetype/truetype"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/geom"
	"golang.org/x/mobile/gl"
)

type Game struct {
	lastCalc   clock.Time
	touchCount uint64
	font       *truetype.Font
}

func NewGame() *Game {
	var g Game
	g.reset()
	return &g
}

func (g *Game) reset() {
	var err error
	g.font, err = LoadCustomFont()
	if err != nil {
		log.Fatalf("error parsing font: %v", err)
	}
}

func (g *Game) Touch(down bool) {
	if down {
		g.touchCount++
	}
}

func (g *Game) Update(now clock.Time) {
	for ; g.lastCalc < now; g.lastCalc++ {
		g.calcFrame()
	}
}

func (g *Game) calcFrame() {

}

func (g *Game) Render(sz size.Event, glctx gl.Context, images *glutil.Images) {
	loading := &TextSprite{
		placeholder:     "feedback",
		text:            display,
		font:            g.font,
		widthPx:         sz.WidthPx,
		heightPx:        sz.HeightPx,
		textColor:       image.White,
		backgroundColor: image.NewUniform(color.Transparent),
		fontSize:        96,
		xPt:             0,
		yPt:             PxToPt(sz, 0),
	}
	loading.Render(sz)
}

func PxToPt(sz size.Event, sizePx int) geom.Pt {
	return geom.Pt(float32(sizePx) / sz.PixelsPerPt)
}
