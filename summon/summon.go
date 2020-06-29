package summon

import (
	"bytes"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"iotqq-plugins-demo/Go/cards"
	"iotqq-plugins-demo/Go/common"
	"iotqq-plugins-demo/Go/userData"
	"iotqq-plugins-demo/Go/util"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

//type GailvConfig struct {
//	FiveStar                 int
//	FiveStarCharacterPickUp  int
//	FiveStarDragonPickUp     int
//	FiveStarCharacterCommon  int
//	FiveStarDragonCommon     int
//	FourStar                 int
//	FourStarCharacterPickUp  int
//	FourStarDragonPickUp     int
//	FourStarCharacterCommon  int
//	FourStarDragonCommon     int
//	ThreeStar                int
//	ThreeStarCharacterCommon int
//	ThreeStarDragonCommon    int
//}
//
//var config = GailvConfig{
//	FiveStar:                 400,
//	FiveStarCharacterPickUp:  100,
//	FiveStarDragonPickUp:     80,
//	FiveStarCharacterCommon:  100,
//	FiveStarDragonCommon:     120,
//	FourStar:                 1600,
//	FourStarCharacterPickUp:  350,
//	FourStarDragonPickUp:     350,
//	FourStarCharacterCommon:  505,
//	FourStarDragonCommon:     395,
//	ThreeStar:                8000,
//	ThreeStarCharacterCommon: 4800,
//	ThreeStarDragonCommon:    3200,
//}
type SummonPool struct {
}

type SummonRecord struct {
	Desc         string
	Card         []SummonCard
	TopBannerUrl string
}

type SummonCard struct {
	cards.Card
	New      bool
	StackNum int
}

var (
	waterIconImage, waterSmall, voucherIconImage, voucherSmall image.Image
)

func InitImageSource() {
	waterIconImage = GetImage("water")
	waterSmall = resize.Resize(20, 20, waterIconImage, resize.Lanczos3)
	voucherIconImage = GetImage("voucher")
	voucherSmall = resize.Resize(20, 20, voucherIconImage, resize.Lanczos3)
}

func OneSummon(user *userData.User) (res *SummonRecord) {
	res = SingleSummon(user, 0)
	res.RaiseUnHitNumber(user)
	return
}

func SingleSummon(user *userData.User, index int) (res *SummonRecord) {
	return singleSummonByCollection(user, index, cards.CardMgr.PickUpOne())
	//if user.UnHitNumber >= 100 && index == 0 {
	//	var cardPoolsKey = []int{511, 512, 521, 522}
	//	res.Card = append(res.Card, *splitSummon(cardPoolsKey))
	//} else if ran < 400+user.UnHitNumber/10*50 {
	//	var cardPoolsKey = []int{511, 512, 521, 522}
	//	res.Card = append(res.Card, *splitSummon(cardPoolsKey))
	//} else if ran < 2000+user.UnHitNumber/10*50 || index == 9 {
	//	var cardPoolsKey = []int{411, 412, 421, 422}
	//	res.Card = append(res.Card, *splitSummon(cardPoolsKey))
	//} else {
	//	var cardPoolsKey = []int{311, 321}
	//	res.Card = append(res.Card, *splitSummon(cardPoolsKey))
	//}

	//ran := rand.Intn(10000)
	//if ran < 400+(user.UnHitNumber/10)*50 {
	//	var cardPoolsKey = []int{511, 512, 521, 522}
	//	res.Card = append(res.Card, *splitSummon(cardPoolsKey))
	//} else if ran < 2000+(user.UnHitNumber/10)*50 || index == 9 {
	//	var cardPoolsKey = []int{411, 412, 421, 422}
	//	res.Card = append(res.Card, *splitSummon(cardPoolsKey))
	//} else {
	//	var cardPoolsKey = []int{311, 321}
	//	res.Card = append(res.Card, *splitSummon(cardPoolsKey))
	//}

	//ran := rand.Intn(10000)
	//if ran < 400 {
	//	var cardPoolsKey = []int{511, 512, 521, 522}
	//	res.Card = append(res.Card, *splitSummon(cardPoolsKey))
	//} else if ran < 2000 {
	//	var cardPoolsKey = []int{411, 412, 421, 422}
	//	res.Card = append(res.Card, *splitSummon(cardPoolsKey))
	//} else {
	//	var cardPoolsKey = []int{311, 321}
	//	res.Card = append(res.Card, *splitSummon(cardPoolsKey))
	//}

	//return res
}

func singleSummonByCollection(user *userData.User, index int, cardCollectionV2 *cards.CardCollectionV2) *SummonRecord {
	//if user.Udid == 570966274 {
	//	user.UnHitNumber = 10000
	//}

	res := new(SummonRecord)
	defer res.CheckAndAddNew(user)

	ran := rand.Intn(10000)

	res.TopBannerUrl = cardCollectionV2.TopBannerUrl
	baseSSRProb := common.BaseSSRProbality*10 + user.UnHitNumber/10*50 + user.Static.RebornCount*50
	//fmt.Println("baseSSRProb:", baseSSRProb, "ran:", ran)
	if user.UnHitNumber >= 100 && index == 0 {
		cardSets := cardCollectionV2.PickUpByStar(5)
		res.Card = append(res.Card, *splitSummonV2(cardSets))
	} else if ran < baseSSRProb+cardCollectionV2.ProbFix*10 {
		cardSets := cardCollectionV2.PickUpByStar(5)
		res.Card = append(res.Card, *splitSummonV2(cardSets))
	} else if ran < 2000+user.UnHitNumber/10*50 || index == 9 {
		cardSets := cardCollectionV2.PickUpByStar(4)
		res.Card = append(res.Card, *splitSummonV2(cardSets))
	} else {
		cardSets := cardCollectionV2.PickUpByStar(3)
		res.Card = append(res.Card, *splitSummonV2(cardSets))
	}
	return res
}

//func splitSummon(keys []int) *SummonCard {
//	sum := 0
//	for _, key := range keys {
//		sum += cards.CardCollection[key].Prob
//	}
//	ranSSR := rand.Intn(sum)
//	gailv := 0
//	for _, key := range keys {
//		gailv += cards.CardCollection[key].Prob
//		if ranSSR < gailv {
//			return &SummonCard{cards.SummonOne(key), false}
//		}
//	}
//	panic("概率溢出")
//}

func splitSummonV2(cardSets []*cards.CardSet) *SummonCard {
	sum := 0
	for _, cardSet := range cardSets {
		sum += cardSet.Prob
	}
	ranSSR := rand.Intn(sum)
	gailv := 0
	for _, cardSet := range cardSets {
		gailv += cardSet.Prob
		if ranSSR < gailv {
			return &SummonCard{cardSet.PickOne(), false, 1}
		}
	}
	panic("概率溢出")
}

func TenSummon(user *userData.User) (res SummonRecord) {
	cardColl := cards.CardMgr.PickUpOne()
	for i := 0; i < 10; i++ {
		summon := singleSummonByCollection(user, i, cardColl)
		res.Card = append(res.Card, summon.Card...)
	}
	res.TopBannerUrl = cardColl.TopBannerUrl
	res.RaiseUnHitNumber(user)
	return res
}

func TenSummonByCollection(user *userData.User, cardCollection *cards.CardCollectionV2) (res SummonRecord) {
	for i := 0; i < 10; i++ {
		summon := singleSummonByCollection(user, i, cardCollection)
		res.Card = append(res.Card, summon.Card...)
	}
	res.RaiseUnHitNumber(user)
	return res
}

func GetMultiSummon(num int) func(user *userData.User) (res SummonRecord) {
	return func(user *userData.User) (res SummonRecord) {
		cardColl := cards.CardMgr.PickUpOne()
		res.TopBannerUrl = cardColl.TopBannerUrl
		for i := 0; i < num/10; i++ {
			tenSummon := TenSummonByCollection(user, cardColl)
			res.Card = append(res.Card, tenSummon.Card...)
			res.Desc = tenSummon.Desc
		}
		return
	}
}

func (s *SummonRecord) Format() string {
	var isTenSummonAbove bool
	var res string
	res += fmt.Sprintf("使用%d张召唤卷进行了召唤\n", len(s.Card))
	if len(s.Card) > 10 {
		res += "10连以上不展示四星三星召唤结果\n"
		isTenSummonAbove = true
	}
	res += fmt.Sprintf("此次您的召唤结果为\n")
	for i := range s.Card {
		if isTenSummonAbove && s.Card[i].Star != 5 {
			continue
		}
		res += fmt.Sprintf("%s,%s,%s\n", s.new(i), s.Card[i].PrintStar(), s.Card[i].Title)
	}
	res += "bang bang bang bang love&die"
	return res
}

//func (s *SummonRecord) ImageFormat(volunterNum, water int) (url string) {
//	bgPng := s.ImageFormatV2(volunterNum, water)
//
//	hash, _ := hashstructure.Hash(s, nil)
//
//	path := "/asset/summon/cache/"
//	out, _ := os.Create(fmt.Sprintf(".%s%d.jpg", path, hash))
//
//	_ = jpeg.Encode(out, bgPng, nil)
//
//	return fmt.Sprintf("http://localhost:12345%s%d.jpg", path, hash)
//}

func (s *SummonRecord) ImageFormatV2(volunterNum int, water int) image.Image {
	bgPng := GetImage("background")
	//merge banner to bg
	//fmt.Println(s.TopBannerUrl)
	banner := GetCardImage(s.TopBannerUrl)
	if banner == nil {
		panic(s.Card[0].Title + s.TopBannerUrl)
	}
	mergeTopBannerToBG(banner, bgPng)

	if len(s.Card) > 1 {
		heightbase, height := 80, 100
		twoColbase, twoColWidth := 70, 120
		threeColbase, threeColWidth := 30, 100
		mergeCardToBG(s.Card[0], bgPng, image.Point{X: twoColbase, Y: heightbase})
		mergeCardToBG(s.Card[1], bgPng, image.Point{X: twoColbase + twoColWidth, Y: heightbase})
		mergeCardToBG(s.Card[2], bgPng, image.Point{X: threeColbase, Y: heightbase + height*1})
		mergeCardToBG(s.Card[3], bgPng, image.Point{X: threeColbase + threeColWidth*1, Y: heightbase + height*1})
		mergeCardToBG(s.Card[4], bgPng, image.Point{X: threeColbase + threeColWidth*2, Y: heightbase + height*1})
		mergeCardToBG(s.Card[5], bgPng, image.Point{X: threeColbase, Y: heightbase + height*2})
		mergeCardToBG(s.Card[6], bgPng, image.Point{X: threeColbase + threeColWidth*1, Y: heightbase + height*2})
		mergeCardToBG(s.Card[7], bgPng, image.Point{X: threeColbase + threeColWidth*2, Y: heightbase + height*2})
		mergeCardToBG(s.Card[8], bgPng, image.Point{X: twoColbase, Y: heightbase + height*3})
		mergeCardToBG(s.Card[9], bgPng, image.Point{X: twoColbase + twoColWidth, Y: heightbase + height*3})
	} else {
		mergeCardToBG(s.Card[0], bgPng, image.Point{X: 130, Y: 230})
	}
	buf := new(bytes.Buffer)
	err := png.Encode(buf, bgPng)
	if err != nil {
		fmt.Println(err)
	}

	board := 70
	waterStatusPoint := image.Point{X: bgPng.Bounds().Dx() - 80 - board, Y: 500}
	drawStatus(water, waterStatusPoint, waterSmall, bgPng)

	volunterStatusPoint := image.Point{X: board, Y: 500}
	drawStatus(volunterNum, volunterStatusPoint, voucherSmall, bgPng)
	return bgPng
}

func drawStatus(num int, point image.Point, icon, bgPng image.Image) {
	blackBar := GetBlackMask(100, 20)
	drawStatusWithBlackBar(strconv.Itoa(num), point, icon, bgPng, blackBar)
}

func drawStatusWithBlackBar(text string, point image.Point, icon image.Image, bgPng image.Image, blackBar *image.RGBA) {
	face := truetype.NewFace(common.Font, &truetype.Options{Size: 13})
	dc := gg.NewContextForImage(blackBar)
	dc.SetFontFace(face)
	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(text, float64(dc.Width()), float64(dc.Height()/2), 1.05, 0.5)
	if icon != nil {
		dc.DrawImage(icon, 0, 0)
	}
	blackBar = dc.Image().(*image.RGBA)

	rectBB := image.Rectangle{
		Min: point,
		Max: point.Add(blackBar.Bounds().Max),
	}

	switch bgPng.(type) {
	case *image.NRGBA:
		draw.Draw(bgPng.(*image.NRGBA), rectBB, blackBar, image.Point{}, draw.Over)
	case *image.RGBA64:
		draw.Draw(bgPng.(*image.RGBA64), rectBB, blackBar, image.Point{}, draw.Over)
	}

}

func mergeTopBannerToBG(banner image.Image, bgPng image.Image) {
	//r := image.Rectangle{Min: dp, Max: dp.Add(dp2)} // 获得更换区域
	if banner.Bounds().Max.X > bgPng.Bounds().Max.X {
		newY := uint(banner.Bounds().Max.Y * bgPng.Bounds().Max.X / banner.Bounds().Max.X)
		banner = resize.Resize(uint(bgPng.Bounds().Max.X), newY, banner, resize.Lanczos3)
	}
	dp1 := image.Point{X: (bgPng.Bounds().Max.X - banner.Bounds().Max.X) / 2}
	dp2 := dp1.Add(banner.Bounds().Max)
	draw.Draw(bgPng.(*image.NRGBA), image.Rectangle{Min: dp1, Max: dp2}, banner, image.Point{}, draw.Over)
}

func mergeCardToBG(card SummonCard, bgPng image.Image, dp1 image.Point) {
	cardPng := productCardPng(card)
	if card.StackNum > 1 {
		blackBar := GetBlackMask(30, 20)
		drawStatusWithBlackBar(fmt.Sprintf("X %d", card.StackNum), image.Point{X: 50, Y: 20}, nil, cardPng, blackBar)
	}
	//r := image.Rectangle{Min: dp, Max: dp.Add(dp2)} // 获得更换区域
	dp2 := dp1.Add(cardPng.Bounds().Max)
	draw.Draw(bgPng.(*image.NRGBA), image.Rectangle{Min: dp1, Max: dp2}, cardPng, image.Point{}, draw.Over)
}

func productCardPng(card SummonCard) image.Image {
	cardPng := GetCardImage(card.IconUrl)
	if cardPng == nil {
		panic(strconv.Itoa(card.ID) + card.IconUrl)
	}
	cardPng = resize.Resize(80, 80, cardPng, resize.Lanczos3)

	dp := image.Point{}
	if card.New {
		newPng := GetImage("new")
		dpNewMin := image.Point{X: 33, Y: 0}
		dpNewMax := dpNewMin.Add(newPng.Bounds().Max)

		draw.Draw(cardPng.(draw.Image), image.Rectangle{Min: dpNewMin, Max: dpNewMax}, newPng, dp, draw.Over)

	} else {
		blackBar := GetBlackMask(65, 15)
		dpBBMin := image.Point{X: 10, Y: 60}
		rectBB := image.Rectangle{
			Min: dpBBMin,
			Max: dpBBMin.Add(blackBar.Bounds().Max),
		}
		draw.Draw(cardPng.(draw.Image), rectBB, blackBar, dp, draw.Over)

		waterFileName := ""
		var width uint
		var start int
		switch card.Water {
		case 85000:
			waterFileName = "85000"
			width = 41
			start = 25
		case 8500:
			waterFileName = "8500"
			width = 41
			start = 25
		case 2200:
			waterFileName = "2200"
			width = 41
			start = 25
		case 150:
			waterFileName = "150"
			width = 30
			start = 30
		case -8500:
			waterFileName = "n8500"
			width = 41
			start = 25
		}
		waterImg := GetImage(waterFileName)
		waterImg = resize.Resize(width, 13, waterImg, resize.Lanczos3)
		dpWaterMin := image.Point{X: start, Y: 60}
		rectWater := image.Rectangle{
			Min: dpWaterMin,
			Max: dpWaterMin.Add(blackBar.Bounds().Max),
		}
		draw.Draw(cardPng.(draw.Image), rectWater, waterImg, dp, draw.Over)

		dpNewMin := image.Point{X: 0, Y: 46}
		dpNewMax := dpNewMin.Add(waterIconImage.Bounds().Max)
		draw.Draw(cardPng.(draw.Image), image.Rectangle{Min: dpNewMin, Max: dpNewMax}, waterIconImage, image.Point{X: 3}, draw.Over)
	}
	return cardPng
}

func GetBlackMask(width, height int) *image.RGBA {
	blackBar := image.NewRGBA(image.Rect(0, 0, width, height))
	util.Clear(blackBar, color.RGBA{R: 0, G: 0, B: 0, A: 128})
	return blackBar
}

func GetCardImage(url string) image.Image {
	if img, ok := IsFileCache(url); ok {
		return img
	} else {
		resp, _ := http.Get(url)
		img, _ = png.Decode(resp.Body)
		if resp == nil {
			return nil
		}
		hashUrl := util.HashURL(url)
		pathPattern := "./asset/cache/%s.png"
		path := fmt.Sprintf(pathPattern, hashUrl)
		out, _ := os.Create(path)
		_ = png.Encode(out, img)
		return img
	}
}

func IsFileCache(srcUrl string) (img image.Image, ok bool) {
	hashUrl := util.HashURL(srcUrl)
	pathPattern := "./asset/cache/%s.png"
	path := fmt.Sprintf(pathPattern, hashUrl)
	if _, err := os.Stat(path); err != nil {
		return nil, false
	} else {
		img = GetImageByPath(path)
		return img, true
	}
}

func GetImage(name string) image.Image {
	prefix := "./asset/summon/"
	return GetImageByPath(prefix + name + ".png")
}

func GetImageByPath(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Panic("could not find ", path, err)
	}
	defer file.Close()
	img, _ := png.Decode(file)
	return img
}

func (s *SummonRecord) new(index int) string {
	if s.Card[index].New {
		return "New!"
	} else {
		return strconv.Itoa(s.Card[index].Water) + "💧"
	}
}

func (s *SummonRecord) GetTotalWater() int {
	res := 0
	for _, card := range s.Card {
		res += card.Water
	}
	return res
}

func (s *SummonRecord) CheckAndAddNew(user *userData.User) {
	for i, card := range s.Card {
		if !util.IntContain(card.ID, user.CardIndex) {
			user.CardIndex = append(user.CardIndex, card.ID)
			s.Card[i].New = true
		} else {
			user.Water += card.Water
		}
	}
}

func (s *SummonRecord) RaiseUnHitNumber(user *userData.User) {
	if s.ContainSSR() {
		//fmt.Println("概率清零")
		user.UnHitNumber = 0
	} else {
		user.UnHitNumber += len(s.Card)
		//fmt.Println("UnHitNumber", user.UnHitNumber)
	}
}

func (s *SummonRecord) ContainSSR() bool {
	for _, card := range s.Card {
		if card.Star == 5 {
			return true
		}
	}
	return false
}

func (s *SummonRecord) CountSSR() int {
	res := 0
	for _, card := range s.Card {
		if card.Star == 5 {
			res += card.StackNum
		}
	}
	return res
}

func (s *SummonRecord) StackCard() {
	m := map[int]*SummonCard{}
	for i, card := range s.Card {
		if value, ok := m[card.ID]; ok {
			value.StackNum++
		} else {
			m[card.ID] = &s.Card[i]
		}
	}
	var summonCards []SummonCard
	for _, value := range m {
		summonCards = append(summonCards, *value)
	}
	s.Card = summonCards
}
