package CloudFont

import (
	// "fmt"
	// "github.com/xuchengzhi/Library/Transfar"
	"github.com/xuchengzhi/apimonitor/gojieba"
	"github.com/xuchengzhi/wordcloud"
	"image/color"
	"log"
	"regexp"
)

func RenderNow(TTF, b_png, new_png string, textList []string) {
	//TTF, b_png string, textList []string
	//生成云词图片
	angles := []int{0, 15, -15, 90}
	colors := []*color.RGBA{
		{0x0, 0x60, 0x30, 0xff},
		{0x60, 0x0, 0x0, 0xff},
		// &color.RGBA{0x73, 0x73, 0x0, 0xff},
	}
	// log.Println(colors)
	log.Println("图片保存到", new_png)
	render := wordcloud_go.NewWordCloudRender(60, 8,
		TTF,
		b_png, textList, angles, colors, new_png)
	render.Render()
}

func Fenci(msg string) []string {
	//使用jieba分词返回list
	var s string
	var words []string
	use_hmm := true
	x := gojieba.NewJieba()
	defer x.Free()

	chiReg := regexp.MustCompile("[^\u4e00-\u9fa5]")
	s = chiReg.ReplaceAllString(msg, "")

	// s = "我来到北京清华大学"
	// words = x.CutAll(s)
	// fmt.Println(s)
	// fmt.Println("全模式:", strings.Join(words, " "))

	words = x.Cut(s, use_hmm)
	// fmt.Println(s)
	// fmt.Println("精确模式:", strings.Join(words, " "))

	// s = "他来到了网易杭研大厦"
	// words = x.Cut(s, use_hmm)
	// fmt.Println(s)
	// fmt.Println("新词识别:", strings.Join(words, "/"))

	// words = x.CutForSearch(s, use_hmm)
	// fmt.Println(words)
	// fmt.Println("搜索引擎模式:", reflect.TypeOf(strings.Join(words, ",")))

	// s = "长春市长春药店"
	// words = x.Tag(s)
	// fmt.Println(s)
	// fmt.Println("词性标注:", strings.Join(words, ","))
	return words
}
