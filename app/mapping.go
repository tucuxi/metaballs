package app

import "metaballs/internal"

var (
	Map_BallSize = map[string]internal.BallSize{
		"Small":  internal.SmallSizeBall,
		"Medium": internal.MediumSizeBall,
		"Large":  internal.LargeSizeBall,
	}

	Map_BallSpeed = map[string]internal.BallSpeed{
		"Slow":   internal.SlowSpeedBall,
		"Medium": internal.NormalSpeedBall,
		"Fast":   internal.FastSpeedBall,
	}

	Map_ScreenSize = map[string]internal.ScreenSize{
		"Small":  internal.ScreenSmall,
		"Medium": internal.ScreenMedium,
		"Large":  internal.ScreenLarge,
	}

	Map_Resolution = map[string]internal.Resolution{
		"Low":    internal.Resolution512,
		"Medium": internal.Resolution1024,
		"High":   internal.Resolution2048,
	}

	Map_FPS = map[string]internal.FPS{
		"30": internal.FPS30,
		"45": internal.FPS45,
		"60": internal.FPS60,

	}

	Map_BallCount = map[string]uint8{
		"4":  4,
		"8":  8,
		"20": 20,
	}

	Map_BallColor = map[string]internal.BallColor{
		"Pink":    internal.Pink,
		"Cyan":  internal.Cyan,
		"Gray":   internal.Gray,
	}
		
)
