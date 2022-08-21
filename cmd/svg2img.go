package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/TaceyWong/svg2img"
)

var (
	help bool

	w int64
	h int64

	o string

	q int
)

func init() {
	flag.BoolVar(&help, "help", false, "显示当前帮助信息")

	flag.Int64Var(&w, "w", 512, "指定输出图片的`宽度`")
	flag.Int64Var(&h, "h", 512, "指定输出图片的`高度`")

	flag.StringVar(&o, "o", "out.png", "指定输出图片的`文件名`")

	flag.IntVar(&q, "q", 100, "输出图片的`质量`")

	flag.Usage = usage
}
func usage() {
	fmt.Fprintf(os.Stderr, `svg2img: 将svg转换为png、jpg、webp格式图片

使用: svg2img [-o 输出文件] [-w 宽] [-h 高] [-q 质量] SVG文件

选项:
`)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}
	if flag.NArg() < 1 {
		log.Fatalln(errors.New("缺少svg输入"))
	}

	err := svg2img.Svg2Img(flag.Arg(0), o, w, h, q)
	if err != nil {
		log.Fatalln(err)
	}
}
