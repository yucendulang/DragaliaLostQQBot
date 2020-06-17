package stickerBot

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"iotqq-plugins-demo/Go/util"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

//表情包来源
//https://bbs.nga.cn/read.php?tid=21466080

type sticker struct {
	url     []string
	title   string
	keyword string
}

var stickerMap map[string]*sticker

//<>符号代表没有找到图源
func init() {
	stickerMap = make(map[string]*sticker)
	new("不行", "不行", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-gpubK1jT1kSe8-e8.jpg")
	new("你好", "你好", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-92uyK1mT1kSe8-e8.jpg")
	new("辛苦", "辛苦", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-kwnK1oT1kSe8-e8.jpg")
	new("谢谢", "谢谢", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-a0mvK1kT1kSe8-e8.jpg")
	new("抱歉", "抱歉", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-26gcK1qT1kSe8-e8.jpg")
	new("好的", "好的", "https://img.nga.178.com/attachments/mon_201904/27/hbQ5-gvq4KuToS3w-3w.png")
	new("集合", "集合", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-9sytK1lT1kSe8-e8.jpg")
	new("很棒", "很棒", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-g8odK1lT1kSe8-e8.jpg")
	new("别在意", "别在意", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-988yK1vT1kSe8-e8.jpg")
	new("等下", "等下", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-a8stK1jT1kSe8-e8.jpg")
	new("救命", "救命", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-ffyiK1vT1kSe8-e8.jpg")
	new("上了哦", "上了哦", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-bo40K1qT1kSe8-e8.jpg")
	new("好机会", "好机会", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-gohgK1sT1kSe8-e8.jpg")
	new("回复", "回复", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-hljqK1qT1kSe8-e8.jpg")
	new("太好", "太好", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-6qeeK1qT1kSe8-e8.jpg")
	new("呜", "呜", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-jsetK1jT1kSe8-e8.jpg")
	new("坏笑", "坏笑", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-dvxcK1iT1kSe8-e8.jpg")
	new("危险", "危险", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-clwwK1zT1kSe8-e8.jpg")
	new("再来一次", "再来一次", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-7184K1oT1kSe8-e8.jpg")
	new("交给我", "交给我", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-i4x1K1uT1kSe8-e8.jpg")
	new("马上就来", "马上就来", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-egmbK1rT1kSe8-e8.jpg")
	new("跳舞", "跳舞", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-8m0wK1jT1kSe8-e8.jpg")
	new("加油", "加油", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-61ofK1oT1kSe8-e8.jpg")
	new("恐怖", "恐怖", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-13tnK1mT1kSe8-e8.jpg")
	new("惊讶", "惊讶", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-j559K1hT1kSe8-e8.jpg")
	new("真的", "真的", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-j559K1hT1kSe8-e8.jpg")
	new("安心", "安心", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-9p04K1nT1kSe8-e8.jpg")
	new("嚯", "嚯", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-9p04K1nT1kSe8-e8.jpg")
	new("我来攻击", "我来攻击", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-5xfeK1oT1kSe8-e8.jpg")
	new("我来防御", "我来防御", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-47ijK1lT1kSe8-e8.jpg")
	new("绝对要赢", "绝对要赢", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-etzyK1mT1kSe8-e8.jpg")
	new("求回复", "求回复", "https://img.nga.178.com/attachments/mon_201904/27/hbQ5-2t36KwToS3w-3w.png")
	new("防御就交给你", "防御就交给你", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-etzyK1mT1kSe8-e8.jpg")
	new("攻击就靠你", "攻击就靠你", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-we4K1uT1kSe8-e8.jpg")
	new("没错没错", "没错没错", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-5pptK1rT1kSe8-e8.jpg")
	new("呀", "呀", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-13i5K1gT1kSe8-e8.jpg")
	new("得意", "得意", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-a3etK1aT1kSe8-e8.jpg")
	new("doya", "doya", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-a3etK1aT1kSe8-e8.jpg")
	new("恭喜", "恭喜", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-hpszK1jT1kSe8-e8.jpg")
	new("你好", "你好", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-9orpK1mT1kSe8-e8.jpg")
	new("快夸我", "快夸我", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-evxnK1mT1kSe8-e8.jpg")
	new("沉默", "沉默", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-23wiK1eT1kSe8-e8.jpg")
	new("得意", "得意", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-aghwK1dT1kSe8-e8.jpg")
	new("doya", "doya", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-aghwK1dT1kSe8-e8.jpg")
	new("拜托", "拜托", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-iu58K1uT1kSe8-e8.jpg")
	new("交给我吧", "交给我吧", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-d558K1uT1kSe8-e8.jpg")
	new("陪你玩玩好", "陪你玩玩好", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-9m0mK22T1kSe8-e8.jpg")
	new("出发吧", "出发吧", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-is5oK1vT1kSe8-e8.jpg")
	new("齐心协力", "齐心协力", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-4y97K1sT1kSe8-e8.jpg")
	new("太好", "太好", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-15lyK1uT1kSe8-e8.jpg")
	new("软绵绵", "软绵绵", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-hej9KvT1kSe8-e8.jpg")
	new("有意思", "有意思", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-4domK1lT1kSe8-e8.jpg")
	new("出事", "出事", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-6aymK1iT1kSe8-e8.jpg")
	new("叫我", "叫我", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-1n5qK1hT1kSe8-e8.jpg")
	new("洛克人登场", "洛克人登场", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-2aliK18T1kSe8-e8.jpg")
	new("E罐", "E罐", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-f41tK17T1kSe8-e8.jpg")
	new("GG", "GG", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-7eanKpT1kSe8-e8.jpg")
	new("GAMEOVER", "GAMEOVER", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-7eanKpT1kSe8-e8.jpg")
	new("磕头", "磕头", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-l2jjK13T1kSe8-e8.jpg")
	new("值得夸奖", "值得夸奖", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-8n7rK1oT1kSe8-e8.jpg")
	new("你真温柔", "你真温柔", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-jrclK1lT1kSe8-e8.jpg")
	new("锵锵", "锵锵", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-i3wgK14T1kSe8-e8.jpg")
	new("新年快乐", "新年快乐", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-b69mK1lT1kSe8-e8.jpg")
	new("新年好", "新年好", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-32hmK1rT1kSe8-e8.jpg")
	new("今年也请多多关照", "今年也请多多关照", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-ffmxK1uT1kSe8-e8.jpg")
	new("狩猎走起", "狩猎走起", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-lczlK1nT1kSe8-e8.jpg")
	new("烤的非常棒", "烤的非常棒", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-g4ciK1jT1kSe8-e8.jpg")
	new("我跟你一起去", "我跟你一起去", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-h4dxK1lT1kSe8-e8.jpg")
	new("艾露猫", "艾露猫", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-3oygK1bT1kSe8-e8.jpg")
	new("嚯", "嚯", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-dr2iK18T1kSe8-e8.jpg")
	new("小狗", "小狗", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-f62uK19T1kSe8-e8.jpg")
	new("祝你做个好梦", "祝你做个好梦", "https://img.nga.178.com/attachments/mon_202004/30/hbQ5-h19pKpToS3w-3w.png")
	new("大哥哥", "大哥哥", "https://img.nga.178.com/attachments/mon_202005/06/hbQ5-gylxKoToS3w-3w.png")
	new("我要改变命运", "我要改变命运", "https://img.nga.178.com/attachments/mon_202004/30/hbQ5-5mf2KnToS3w-3w.png")
	new("我想和你成为朋友", "我想和你成为朋友", "https://img.nga.178.com/attachments/mon_202004/30/hbQ5-bawaKoToS3w-3w.png")
	new("菲", "菲", "https://img.nga.178.com/attachments/mon_202005/09/hbQ5-271tKgToS3w-3w.png")
	new("要来我的同盟", "要来我的同盟", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-2t35K1qT1kSe8-e8.jpg")
	new("成为我的战友", "成为我的战友", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-l1iaK1uT1kSe8-e8.jpg")
	new("趁现在龙化", "趁现在龙化", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-8em2K1wT1kSe8-e8.jpg")
	new("晚安", "晚安", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-d98wK1gT1kSe8-e8.jpg")
	new("待会见", "待会见", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-807tK1nT1kSe8-e8.jpg")
	new("休息一下", "休息一下", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-cu0fK1oT1kSe8-e8.jpg")
	new("左", "左", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-lescKjT1kSe8-e8.jpg")
	new("左上", "左上", "https://img.nga.178.com/attachments/mon_202001/29/hbQ5-4nznK5ToS3w-3w.png")
	new("上", "上", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-7hf5KjT1kSe8-e8.jpg")
	new("右上", "右上", "https://img.nga.178.com/attachments/mon_202001/29/hbQ5-ejbuK5ToS3w-3w.png")
	new("右", "右", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-7749KjT1kSe8-e8.jpg")
	new("右下", "右下", "https://img.nga.178.com/attachments/mon_202001/29/hbQ5-2pv0K5ToS3w-3w.png")
	new("下", "下", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-eaxoKkT1kSe8-e8.jpg")
	new("左下", "左下", "https://img.nga.178.com/attachments/mon_202001/29/hbQ5-ddddK5ToS3w-3w.png")
	new("就差一点", "就差一点", "https://img.nga.178.com/attachments/mon_202002/14/hbQ5-l17sKoToS3w-3w.png")
	new("能行", "能行", "https://img.nga.178.com/attachments/mon_202006/12/hbQ5-f88vKkToS3w-3w.png")
	new("永远好伙计", "永远好伙计", "https://img.nga.178.com/attachments/mon_202006/12/hbQ5-ew8KnToS3w-3w.png")

	//恶搞
	new("大哥哥", "大哥哥(小狗)", "https://img.nga.178.com/attachments/mon_202005/08/hbQ5-3ni6KiToS3w-3w.jpg")
	new("大哥哥", "大哥哥(兰扎卜)", "https://img.nga.178.com/attachments/mon_202005/16/hbQ5-2aigKuToS3w-3w.png")
	//DIY部分
	new("大佬", "大佬", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-aucK18T1kSe8-e8.jpg")
	new("没毛", "没毛", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-bqj3K16T1kSe8-e8.jpg")
	new("饿", "饿", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-ft2fK1gT1kSe8-e8.jpg")
	new("没体力", "没体力", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-6sesK1gT1kSe8-e8.jpg")
	new("ng", "ng", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-g2yvK1iT1kSe8-e8.jpg")
	new("NG", "ng", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-g2yvK1iT1kSe8-e8.jpg")
	new("欢迎", "欢迎", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-7ndcK1jT1kSe8-e8.jpg")
	new("送客", "送客", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-kkdeK1jT1kSe8-e8.jpg")
	new("尴尬", "尴尬", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-c3erK1jT1kSe8-e8.jpg")
	new("萌新", "萌新", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-1x3oK1jT1kSe8-e8.jpg")
	new("硬核", "硬核", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-ea0xK1jT1kSe8-e8.jpg")
	new("肝", "肝", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-5v9hK1hT1kSe8-e8.jpg")
	new("肝帝", "肝帝", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-ivwnK1jT1kSe8-e8.jpg")
	new("肝报废", "肝报废", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-9zncK1iT1kSe8-e8.jpg")
	new("没体", "没体", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-kd55K1jT1kSe8-e8.jpg")
	new("丢人", "丢人", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-bxw6K1iT1kSe8-e8.jpg")
	new("雷", "雷", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-2y1aK1hT1kSe8-e8.jpg")
	new("课", "课", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-fz5rK1iT1kSe8-e8.jpg")
	new("课长", "课长", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-7ixbK1jT1kSe8-e8.jpg")
	new("完", "完", "https://img.nga.178.com/attachments/mon_202004/25/hbQ5-kj5oK1hT1kSe8-e8.jpg")
}

func new(keyword, title, url string) {
	if value, ok := stickerMap[keyword]; !ok {
		stickerMap[keyword] = &sticker{
			keyword: keyword,
			title:   title,
			url:     []string{url},
		}
	} else {
		value.url = append(value.url, url)
	}
}

func IsStickerKey(key string) (string, bool) {
	if len(key) > 20 {
		return "", false
	}
	if res, ok := stickerMap[key]; ok {
		url := res.url[rand.Intn(len(res.url))]
		if localUrl, ok1 := IsStickerFileLocal(url); ok1 {
			return localUrl, true
		} else {
			go CacheStickerFile(url)
			return res.url[rand.Intn(len(res.url))], ok
		}
	} else {
		return "", ok
	}
}

func IsStickerFileLocal(srcUrl string) (localUrl string, ok bool) {
	hashUrl := util.HashURL(srcUrl)
	var suffix string
	if strings.HasSuffix(srcUrl, "jpg") {
		suffix = "jpg"
	} else {
		suffix = "png"
	}
	if _, err := os.Stat(fmt.Sprintf("./asset/stickers/%s.%s", hashUrl, suffix)); err != nil {
		return "", false
	} else {
		return fmt.Sprintf("http://localhost:12345/asset/stickers/%s.%s", hashUrl, suffix), true
	}
}

func CacheStickerFile(url string) {
	resp, _ := http.Get(url)
	//body,_:=ioutil.ReadAll(resp.Body)

	// Decoding gives you an Image.
	// If you have an io.Reader already, you can give that to Decode
	// without reading it into a []byte.
	var img image.Image
	var isJPG = strings.HasSuffix(url, "jpg")
	if isJPG {
		img, _ = jpeg.Decode(resp.Body)
	} else {
		img, _ = png.Decode(resp.Body)
	}

	// check err
	fmt.Println(url, img.Bounds().Dx())
	if img.Bounds().Dx() > 140 {
		img = resize.Resize(140, 140, img, resize.Lanczos3)
	} else {
		fmt.Println(img.Bounds().Dx())
	}

	var out *os.File
	if isJPG {
		out, _ = os.Create(fmt.Sprintf("./asset/stickers/%s.jpg", util.HashURL(url)))
	} else {
		out, _ = os.Create(fmt.Sprintf("./asset/stickers/%s.png", util.HashURL(url)))
	}

	if strings.HasSuffix(url, "jpg") {
		jpeg.Encode(out, img, nil)
	} else {
		png.Encode(out, img)
	}

	//io.Copy(out,bytes.NewReader(body))
}
