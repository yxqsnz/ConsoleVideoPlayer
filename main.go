package main

import (
	"cvp/src/utils"
	"fmt"
	"image/color"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/dustin/go-humanize"
	"github.com/eliukblau/pixterm/pkg/ansimage"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("Console video player by yxqsnz.")
	if _, err := os.Stat("./assets/"); os.IsNotExist(err) {

		utils.CheckError(os.Mkdir("./assets", 0777))
		utils.CheckError(os.Mkdir("./assets/frames", 0777))
	}
	if _, err := os.Stat("./assets/video.mp4"); os.IsNotExist(err) {
		logrus.Fatalln("please put a file in `./assets/video.mp4` ")
		return
	}
	files, err := ioutil.ReadDir("./assets/frames/")
	if len(files) <= 1 || err != nil {
		utils.ConvertVideoToFrames()

	}
	ansimage.ClearTerminal()
	//Ok

	utils.CheckError(err)
	fss := len(files)
	i := 0
	for i < fss {
		img, err := ansimage.NewScaledFromFile(fmt.Sprintf("./assets/frames/%d.png", i), 120, 100, color.Transparent, ansimage.ScaleModeFit, ansimage.NoDithering)
		utils.CheckError(err)

		img.Draw()
		i++
		m := runtime.MemStats{}
		runtime.ReadMemStats(&m)
		//m.HeapObjects/1024/1024
		logrus.Printf("[%d%%] %d of %d RAM: %v \n", (100*i)/fss, i, fss, humanize.Bytes(m.Alloc))
	}
}
