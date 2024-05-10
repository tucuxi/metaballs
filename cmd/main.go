// package internal

// // import "github.com/pkg/profile"

// func main() {
// 	// defer profile.Start(profile.CPUProfile,profile.ProfilePath("./profile")).Stop()
// 	// defer profile.Start(profile.TraceProfile, profile.ProfilePath("./profile"), profile.NoShutdownHook,).Stop()
// 	NewMetaBallApp(
// 		5,
// 		SlowSpeedBall,
// 		MediumSizeBall,
// 		FPS(45),
// 		ScreenSmall,
// 		Resolution512).
// 		Run()

// }

package main

import "metaballs/app"

func main(){

    app.Run()


}