package utils

import (
	"fmt"

	"gocv.io/x/gocv"
)

func ConvertVideoToFrames() {
	currentFrame := 0
	video, _ := gocv.VideoCaptureFile("./assets/video.mp4")
	window := gocv.NewWindow("Uma janela (:")
	img := gocv.NewMat()
	var sucess bool = true
	var writeSucess bool = true

	for sucess || writeSucess {
		sucess = video.Read(&img)

		writeSucess = gocv.IMWrite(fmt.Sprintf("./assets/frames/%d.png", currentFrame), img)

		window.SetWindowTitle(fmt.Sprintf("Processing video... Current Frame: %v", currentFrame))
		window.IMShow(img)

		window.WaitKey(1)
		currentFrame++
	}
}
