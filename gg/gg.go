package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"github.com/siddontang/go/log"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	_ = imageMaker("gg/resource/test2.png")
}

func drawCircle() {
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(500, 500, 400)
	dc.SetRGB255(255, 255, 255)
	dc.Fill()
	dc.SavePNG("out.png")
}

//filePath 封面图
func imageMaker(filePath string) error {
	//load cover
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var imgTmp image.Image
	// decode jpeg into image.Image
	imgTmp, err = png.Decode(file)
	if err != nil {
		_ = file.Close()
		fi, _ := os.Open(filePath)
		imgTmp, err = jpeg.Decode(fi)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	_ = file.Close()
	w := imgTmp.Bounds().Dx()
	h := imgTmp.Bounds().Dy()
	if w != h {
		return fmt.Errorf("不是正方形")
	}

	m := resize.Resize(500, 500, imgTmp, resize.Lanczos3)
	out, err := os.Create("tmp_image.jpeg")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer out.Close()
	_ = jpeg.Encode(out, m, nil)

	//load background
	imgBkg, err := gg.LoadImage("test.png")
	if err != nil {
		log.Fatal(err)
		return err
	}
	//draw image
	dc := gg.NewContext(526, 800)
	dc.DrawImage(imgBkg, 0, 0)
	imgCover, err := gg.LoadImage("tmp_image.jpeg")
	dc.DrawImage(imgCover, 13, 13)

	//保存的制作图片
	err = dc.SavePNG("test.png")
	if err != nil {
		log.Error()
		return err
	}

	_ = os.Remove("tmp_image.jpeg")
	return nil
}

/*
methods:
image:
DrawImage(im image.Image, x, y int)
DrawImageAnchored(im image.Image, x, y int, ax, ay float64)

string:
SetFontFace(fontFace font.Face)
DrawString(s string, x, y float64)
DrawStringAnchored(s string, x, y, ax, ay float64)

color:
SetRGB(r, g, b float64)
SetRGBA(r, g, b, a float64)
SetRGB255(r, g, b int)

NewSurfacePattern(im image.Image, op RepeatOp)

Load:
LoadImage(path string) (image.Image, error)
LoadPNG(path string) (image.Image, error)
SavePNG(path string, im image.Image) error

func (a Matrix ) Scale(x, y float64 )矩阵
*/
