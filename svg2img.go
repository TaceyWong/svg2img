package svg2img

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/chai2010/webp"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

func Svg2Img(src, dst string, width, height int64, quality int) error {
	dstE := strings.Split(dst, ".")
	if len(dstE) < 2 {
		return errors.New("请提供webp,png,jpg,jpeg后缀")
	}
	format := ""
	switch strings.ToLower(dstE[len(dstE)-1]) {
	case "jpg", "jpeg":
		format = "jpg"
	case "png":
		format = "png"
	case "webp":
		format = "webp"
	default:
		return errors.New("目标格式不支持")
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	icon, _ := oksvg.ReadIconStream(in)
	icon.SetTarget(0, 0, float64(width), float64(height))
	rgba := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	icon.Draw(rasterx.NewDasher(int(width), int(height),
		rasterx.NewScannerGV(int(width), int(height), rgba, rgba.Bounds())), 1)

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	switch format {
	case "jpg":
		return jpeg.Encode(out, rgba, &jpeg.Options{Quality: quality})
	case "png":
		return png.Encode(out, rgba)

	case "webp":
		return webp.Encode(out, rgba, &webp.Options{Lossless: false, Quality: float32(quality)})
	}
	return nil

}
