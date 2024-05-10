package app

import (
	"fmt"
	"metaballs/internal"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	
)


func Run() {
	model := initialModel()

	t := tea.NewProgram(model)
	if _, err := t.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	ballSize := Map_BallSize[METABALL_SELECTIONS["Ball Size"].GetSelected()]
	ballSpeed := Map_BallSpeed[METABALL_SELECTIONS["Ball Speed"].GetSelected()]

	screenSize := Map_ScreenSize[METABALL_SELECTIONS["Screen Size"].GetSelected()]
	resolution := Map_Resolution[METABALL_SELECTIONS["Resolution"].GetSelected()]
	fps := Map_FPS[METABALL_SELECTIONS["FPS"].GetSelected()]
	
	ballCount := Map_BallCount[METABALL_SELECTIONS["Ball Count"].GetSelected()]
	ballColor := Map_BallColor[METABALL_SELECTIONS["Ball Color"].GetSelected()]

	metaBallApp := internal.NewMetaBallApp(ballCount, ballSpeed, ballSize, ballColor, fps, screenSize, resolution)
	metaBallApp.Run()
}
