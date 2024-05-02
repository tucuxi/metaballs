package internal

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

type BallSpeed struct {
	minVelocity float32
	maxVelocity float32
}

var (
	SlowSpeedBall   = BallSpeed{0.003, 0.007}
	NormalSpeedBall = BallSpeed{0.007, 0.01}
	FastSpeedBall   = BallSpeed{0.01, 0.02}
)

type FPS time.Duration

var (
	FPS30  = FPS(time.Second / 30)
	FPS45  = FPS(time.Second / 45)
	FPS60  = FPS(time.Second / 60)
	FPS120 = FPS(time.Second / 120)
)

type BallSize struct {
	minRadius float32
	maxRadius float32
}

var (
	SmallSizeBall  = BallSize{0.025, 0.075}
	MediumSizeBall = BallSize{0.075, 0.1}
	LargeSizeBall  = BallSize{0.1, 0.15}
)

type ScreenSize struct {
	width  int
	height int
}

var (
	ScreenSmall  = ScreenSize{300, 300}
	ScreenMedium = ScreenSize{512, 512}
	ScreenLarge  = ScreenSize{768, 768}
)

type Resolution int

var (
	Resolution128  = Resolution(64)
	Resolution256  = Resolution(256)
	Resolution512  = Resolution(512)
	Resolution1024 = Resolution(900)
	Resolution2048 = Resolution(2048)
)

type MetaBallApp struct {
	fyneApp fyne.App
	screen  *Screen

	ballCount uint8
	ballSpeed BallSpeed
	ballSize  BallSize
	ballColor BallColor
	fps FPS

	screenSize ScreenSize
	resolution Resolution
}

func NewDefaultMetaBallApp() *MetaBallApp {

	return &MetaBallApp{
		ballCount: 8,
		ballSpeed: NormalSpeedBall,
		ballSize:  MediumSizeBall,

		fps: FPS30,

		screenSize: ScreenSmall,
		resolution: Resolution256,
	}
}

func NewMetaBallApp(ballCount uint8, ballSpeed BallSpeed, ballSize BallSize,ballColor BallColor, fps FPS, screenSize ScreenSize, resolution Resolution) *MetaBallApp {

	return &MetaBallApp{
		ballCount: ballCount,
		ballSpeed: ballSpeed,
		ballSize:  ballSize,
		ballColor: ballColor,

		fps: fps,

		screenSize: screenSize,
		resolution: resolution,
	}

}

func (m *MetaBallApp) Run() {
	m.fyneApp = app.New()
	g := newRandomGroup(int(m.ballCount), m.ballSpeed, m.ballSize)
	m.screen = NewScreen(g, m.resolution,m.ballColor)

	w := m.fyneApp.NewWindow("Metaballs")

	m.screen.raster = canvas.NewRaster(m.screen.draw)
	m.screen.ExtendBaseWidget(m.screen)
	w.SetContent(m.screen)
	w.Resize(fyne.NewSize(float32(m.screenSize.width), float32(m.screenSize.height)))

	m.Animate()
	w.ShowAndRun()
}

func (m *MetaBallApp) Animate() {
	// render the screen
	go func() {

		frames := time.NewTicker(time.Second / time.Duration(m.fps))
		for range frames.C {
			go m.screen.Refresh()
		}
	}()
	// move the balls
	go func() {

		frames := time.NewTicker(time.Millisecond * 40)
		for range frames.C {
			go m.screen.group.move()
		}
	}()
}
