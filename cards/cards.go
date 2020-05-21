package cards

import (
	"log"
	"math/rand"
)

/*
Card是一张抽的卡
CardSet是卡的集合 代表了某一种概率下的卡的集合
CardCollection是CardSet的集合 代表了一个池子
不同的池子 就是不同的CardCollection
*/

var Cards []Card

var PickupCards5StarCharacter CardSet

type cardCollection map[int]*CardSet

var CardCollection cardCollection

type CardCollectionV2 struct {
	cardSets     []*CardSet
	TopBannerUrl string
}

type CardSet struct {
	star     int
	cardType int //1是character 2是Dragon
	rareType int //1代表普池 2代表特选
	Prob     int //10000进制 1%填100
	cards    []Card
}

func initCards() {
	Cards = []Card{
		{Title: "占位",
			IconUrl: ""},
		{ID: 1, Star: 3, Title: "影龙", Water: 150, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-afzoKbT8S28-28.png"},
		{ID: 2, Star: 4, Title: "希露姬", Water: 2200, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-aewKcT8S28-28.png"},
		{ID: 3, Star: 5, Title: "纳杰夫", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-l95xKdT8S28-28.png"},
		{ID: 4, Star: 5, Title: "丽雅", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201909/12/hbQ5-75k0KfT8S28-28.png"},
		{ID: 5, Star: 5, Title: "米科特", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-l138KdT8S28-28.png"},
		{ID: 6, Star: 5, Title: "雷吉娜", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201907/11/hbQ5-8imtKdT8S28-28.png"},
		{ID: 7, Star: 5, Title: "艾塞莉特", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-kspKdT8S28-28.png"},
		{ID: 8, Star: 5, Title: "罗吉娜", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201907/11/hbQ5-2xwfKdT8S28-28.png"},
		{ID: 9, Star: 5, Title: "切尔茜", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201911/13/hbQ5-8v81KfT8S28-28.png"},
		{ID: 10, Star: 5, Title: "梅莉贝尔(花开Vers.)", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201905/31/hbQ5-29dsKfT8S28-28.png"},
		{ID: 11, Star: 5, Title: "希尔德加德(情人节Vers.)", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-9i15KdT8S28-28.png"},
		{ID: 12, Star: 5, Title: "阿莱克西斯", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-jwl0KdT8S28-28.png"},
		{ID: 13, Star: 5, Title: "赛丽艾拉(夏日Vers.)", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201908/01/hbQ5-7tr5KdT8S28-28.png"},
		{ID: 14, Star: 5, Title: "瓦菜里奥", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202002/28/hbQ5-f08hKeT8S28-28.png"},
		{ID: 15, Star: 5, Title: "三叶", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202003/03/hbQ5-kn7sKfT8S28-28.png"},
		{ID: 16, Star: 5, Title: "朱丽叶(夏日Vers.)", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201908/01/hbQ5-26nrKcT8S28-28.png"},
		{ID: 17, Star: 5, Title: "拉兹莉", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202004/14/hbQ5-ien9KdT8S28-28.png"},
		{ID: 18, Star: 5, Title: "扎因弗拉德", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-bjkeKeT8S28-28.png"},
		{ID: 19, Star: 5, Title: "拉拉诺亚", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201903/31/hbQ5-514wKdT8S28-28.png"},
		{ID: 20, Star: 5, Title: "莉莉", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-2wlmKeT8S28-28.png"},
		{ID: 21, Star: 5, Title: "太公望", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201912/02/hbQ5-93ahKdT8S28-28.png"},
		{ID: 22, Star: 5, Title: "埃尔菲利斯(新娘Vers.)", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201906/20/hbQ5-3555KeT8S28-28.png"},
		{ID: 23, Star: 5, Title: "维克托", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201908/31/hbQ5-7pr2KeT8S28-28.png"},
		{ID: 24, Star: 5, Title: "德莱茨", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202003/12/hbQ5-8ytiKdT8S28-28.png"},
		{ID: 25, Star: 5, Title: "葵(新娘Vers.)", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201906/20/hbQ5-f4c8KfT8S28-28.png"},
		{ID: 26, Star: 5, Title: "林佑", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-1d8cKdT8S28-28.png"},
		{ID: 27, Star: 5, Title: "加斯特", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/19/hbQ5-7py6KdT8S28-28.png"},
		{ID: 28, Star: 5, Title: "梅利贝尔", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-66c2KdT8S28-28.png"},
		{ID: 29, Star: 5, Title: "科斯蒂", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202001/15/hbQ5-743fKdT8S28-28.png"},
		{ID: 30, Star: 5, Title: "洛伊泽", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-bmojKdT8S28-28.png"},
		{ID: 31, Star: 5, Title: "霍克", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-i0jsKcT8S28-28.png"},
		{ID: 32, Star: 5, Title: "阿加莎", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201911/05/hbQ5-792vKfT8S28-28.png"},
		{ID: 33, Star: 5, Title: "阿尔贝尔", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201903/20/-mjkdpQ5-xchKcT8S28-28.png"},
		{ID: 34, Star: 5, Title: "夏迪(美型Vers.)", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201905/14/hbQ5-k3c5KfT8S28-28.png"},
		{ID: 35, Star: 5, Title: "朱丽叶", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-doxxKdT8S28-28.png"},
		{ID: 36, Star: 5, Title: "安妮丽艾", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-6ncoKdT8S28-28.png"},
		{ID: 37, Star: 5, Title: "库拉乌(夏日Vers.)", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201908/13/hbQ5-b8a0KfT8S28-28.png"},
		{ID: 38, Star: 5, Title: "卢克雷齐娅", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-9xbkKdT8S28-28.png"},
		{ID: 39, Star: 5, Title: "希尔德加德", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-eqyhKdT8S28-28.png"},
		{ID: 40, Star: 5, Title: "夜天", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201906/30/hbQ5-c34eKeT8S28-28.png"},
		{ID: 41, Star: 5, Title: "娜茨", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201906/30/hbQ5-6fw3KfT8S28-28.png"},
		{ID: 42, Star: 5, Title: "德尔菲", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201909/30/hbQ5-84w4KfT8S28-28.png"},
		{ID: 43, Star: 5, Title: "贝利娜", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202004/02/hbQ5-eky9KdT8S28-28.png"},
		{ID: 44, Star: 5, Title: "库格尔", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201903/20/hbQ5-9kezKfT8S28-28.png"},
		{ID: 45, Star: 5, Title: "拉托尼", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201910/31/hbQ5-lab7KfT8S28-28.png"},
		{ID: 46, Star: 5, Title: "奈法利耶", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-7tumKcT8S28-28.png"},
		{ID: 47, Star: 5, Title: "卡珊德拉", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201909/30/hbQ5-2mdvKfT8S28-28.png"},
		{ID: 48, Star: 5, Title: "海因瓦德尔", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201903/20/hbQ5-f8s8KfT8S28-28.png"},
		{ID: 49, Star: 5, Title: "维尔莎拉(夏日Vers.)", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201908/13/hbQ5-66lKfT8S28-28.png"},
		{ID: 50, Star: 5, Title: "格蕾丝", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202004/06/hbQ5-gpd0KdT8S28-28.png"},
		{ID: 51, Star: 4, Title: "卡尔", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-5airKdT8S28-28.png"},
		{ID: 52, Star: 4, Title: "塞雷娜", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201905/31/hbQ5-hvigKfT8S28-28.png"},
		{ID: 53, Star: 4, Title: "勇也", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201911/13/hbQ5-3bz6KfT8S28-28.png"},
		{ID: 54, Star: 4, Title: "凡妮莎", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-aqkoKdT8S28-28.png"},
		{ID: 55, Star: 4, Title: "悦", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/19/hbQ5-20w3KcT8S28-28.png"},
		{ID: 56, Star: 4, Title: "艾玛", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201909/12/hbQ5-pgoKfT8S28-28.png"},
		{ID: 57, Star: 4, Title: "艾塞莉特(情人节Vers.)", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-627nKdT8S28-28.png"},
		{ID: 58, Star: 4, Title: "希诺亚", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-fmeiKeT8S28-28.png"},
		{ID: 59, Star: 4, Title: "维尔沙拉", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-f92fKdT8S28-28.png"},
		{ID: 60, Star: 4, Title: "兰扎卜(夏日Vers.)", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201908/01/hbQ5-hy2yKcT8S28-28.png"},
		{ID: 61, Star: 4, Title: "卢塔", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-itapKdT8S28-28.png"},
		{ID: 62, Star: 4, Title: "奥尔森", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-8vgzKdT8S28-28.png"},
		{ID: 63, Star: 4, Title: "卢吉娜", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201907/11/hbQ5-itinKdT8S28-28.png"},
		{ID: 64, Star: 4, Title: "卡丽娜", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-3f1sKeT8S28-28.png"},
		{ID: 65, Star: 4, Title: "戈尔德", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202004/14/hbQ5-8pyqKcT8S28-28.png"},
		{ID: 66, Star: 4, Title: "啵噜", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202002/28/hbQ5-9cvaKeT8S28-28.png"},
		{ID: 67, Star: 4, Title: "塞纳", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-2hr0KdT8S28-28.png"},
		{ID: 68, Star: 4, Title: "艾露(花开Vers.)", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201905/31/hbQ5-bnqzKfT8S28-28.png"},
		{ID: 69, Star: 4, Title: "苦海", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-cax9KcT8S28-28.png"},
		{ID: 70, Star: 4, Title: "霍普(教会骑士Vers.)", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202003/12/hbQ5-ejbbKdT8S28-28.png"},
		{ID: 71, Star: 4, Title: "武藏", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-fuz1KdT8S28-28.png"},
		{ID: 72, Star: 4, Title: "艾露", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-dz2fKeT8S28-28.png"},
		{ID: 73, Star: 4, Title: "皮雅茜", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-4aq0KdT8S28-28.png"},
		{ID: 74, Star: 4, Title: "艾蕾欧诺拉", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-jogyKdT8S28-28.png"},
		{ID: 75, Star: 4, Title: "诺艾尔", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201908/31/hbQ5-254nKgT8S28-28.png"},
		{ID: 76, Star: 4, Title: "卢恩", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-h2ntKdT8S28-28.png"},
		{ID: 77, Star: 4, Title: "桑尼亚(新娘Vers.)", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201906/20/hbQ5-99xyKfT8S28-28.png"},
		{ID: 78, Star: 4, Title: "欧蒂塔", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-hf6dKdT8S28-28.png"},
		{ID: 79, Star: 4, Title: "八千代", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201905/14/hbQ5-em51KgT8S28-28.png"},
		{ID: 80, Star: 4, Title: "弗里茨", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-16jjKdT8S28-28.png"},
		{ID: 81, Star: 4, Title: "芙露露", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201903/31/hbQ5-kscfKdT8S28-28.png"},
		{ID: 82, Star: 4, Title: "卢卡(夏日Vers.)", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201908/13/hbQ5-5py3KfT8S28-28.png"},
		{ID: 83, Star: 4, Title: "梁泉", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-q0jKeT8S28-28.png"},
		{ID: 84, Star: 4, Title: "乌尔嘉", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-k7yuKdT8S28-28.png"},
		{ID: 85, Star: 4, Title: "天音", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-9z37KdT8S28-28.png"},
		{ID: 86, Star: 4, Title: "维克塞尔", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-k470KdT8S28-28.png"},
		{ID: 87, Star: 4, Title: "贝尔扎克", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-en0gKdT8S28-28.png"},
		{ID: 88, Star: 4, Title: "杜拉尔德", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201912/02/hbQ5-eke5KdT8S28-28.png"},
		{ID: 89, Star: 4, Title: "俄里翁", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-96ayKdT8S28-28.png"},
		{ID: 90, Star: 4, Title: "帕蒂亚", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201909/30/hbQ5-dq16KfT8S28-28.png"},
		{ID: 91, Star: 4, Title: "诺斯通", Water: 2200, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201905/14/hbQ5-3n9yKfT8S28-28.png"},
		{ID: 92, Star: 3, Title: "马蒂", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-afh1KdT8S28-28.png"},
		{ID: 93, Star: 3, Title: "俄里翁(情人节Vers.)", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-gfi2KdT8S28-28.png"},
		{ID: 94, Star: 3, Title: "葵", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-la73KdT8S28-28.png"},
		{ID: 95, Star: 3, Title: "拉辛", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-fuowKcT8S28-28.png"},
		{ID: 96, Star: 3, Title: "阿兰", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-cj9hKdT8S28-28.png"},
		{ID: 97, Star: 3, Title: "乔", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-4qc8KcT8S28-28.png"},
		{ID: 98, Star: 3, Title: "桑尼亚", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-kqm3KcT8S28-28.png"},
		{ID: 99, Star: 3, Title: "欧雷因", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-a6bqKcT8S28-28.png"},
		{ID: 100, Star: 3, Title: "夏迪", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-1c6vKdT8S28-28.png"},
		{ID: 101, Star: 3, Title: "十太郎", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-csprKdT8S28-28.png"},
		{ID: 102, Star: 3, Title: "雷克斯", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-dz7tKdT8S28-28.png"},
		{ID: 103, Star: 3, Title: "皮多特", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-7xb3KdT8S28-28.png"},
		{ID: 104, Star: 3, Title: "修贝尔", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-h108KeT8S28-28.png"},
		{ID: 105, Star: 3, Title: "让", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-anihKdT8S28-28.png"},
		{ID: 106, Star: 3, Title: "威克", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-73hbKdT8S28-28.png"},
		{ID: 107, Star: 3, Title: "里卡多", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-8j8iKdT8S28-28.png"},
		{ID: 108, Star: 3, Title: "米罗蒂", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-pkgKdT8S28-28.png"},
		{ID: 109, Star: 3, Title: "弗兰切斯卡", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-2jccKdT8S28-28.png"},
		{ID: 110, Star: 3, Title: "伊汉娜", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-a5fhKcT8S28-28.png"},
		{ID: 111, Star: 3, Title: "菲丽雅", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-8i9cKdT8S28-28.png"},
		{ID: 112, Star: 3, Title: "尼古拉", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-gojhKdT8S28-28.png"},
		{ID: 113, Star: 3, Title: "索菲", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-ld1xKcT8S28-28.png"},
		{ID: 114, Star: 3, Title: "莱蒙德", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-boouKdT8S28-28.png"},
		{ID: 115, Star: 3, Title: "伊尔凡", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-38xsKdT8S28-28.png"},
		{ID: 116, Star: 3, Title: "莱纳斯", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-675uKeT8S28-28.png"},
		{ID: 117, Star: 3, Title: "马尔卡", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-5rvxKdT8S28-28.png"},
		{ID: 118, Star: 3, Title: "马萝拉", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-8d3KcT8S28-28.png"},
		{ID: 119, Star: 3, Title: "小蕾", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-9aktKdT8S28-28.png"},
		{ID: 120, Star: 3, Title: "霍普", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-h5hyKdT8S28-28.png"},
		{ID: 121, Star: 3, Title: "艾斯蒂尔", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-inffKdT8S28-28.png"},
		{ID: 122, Star: 3, Title: "罗德里戈", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-ircmKdT8S28-28.png"},
		{ID: 123, Star: 3, Title: "太郎", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-ic04KdT8S28-28.png"},
		{ID: 124, Star: 3, Title: "华兹", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-3pp7KdT8S28-28.png"},
		{ID: 125, Star: 3, Title: "维特", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-cdx7KdT8S28-28.png"},
		{ID: 126, Star: 3, Title: "埃里克", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-459lKdT8S28-28.png"},
		{ID: 127, Star: 3, Title: "吉斯", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-jcdqKcT8S28-28.png"},
		{ID: 128, Star: 3, Title: "伊露提米娅", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-6w2zKdT8S28-28.png"},
		{ID: 129, Star: 3, Title: "爱德华", Water: 150, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-k4fsKcT8S28-28.png"},
		{ID: 130, Star: 5, Title: "阿格尼", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-hjlnKcT8S28-28.png"},
		{ID: 131, Star: 5, Title: "坎贝萝丝", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-gpoqKdT8S28-28.png"},
		{ID: 132, Star: 5, Title: "普罗米修斯", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-5f28KeT8S28-28.png"},
		{ID: 133, Star: 5, Title: "木开花耶", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201905/31/hbQ5-5xc2KfT8S28-28.png"},
		{ID: 134, Star: 5, Title: "阿尔库特斯", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201907/11/hbQ5-cy1mKdT8S28-28.png"},
		{ID: 135, Star: 5, Title: "阿波罗", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201909/12/hbQ5-g2pwKfT8S28-28.png"},
		{ID: 136, Star: 5, Title: "伽具土", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201911/13/hbQ5-ede6KgT8S28-28.png"},
		{ID: 137, Star: 5, Title: "波塞冬", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-chywKdT8S28-28.png"},
		{ID: 138, Star: 5, Title: "利维坦", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-1h49KdT8S28-28.png"},
		{ID: 139, Star: 5, Title: "塞壬", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201908/01/hbQ5-10ptKdT8S28-28.png"},
		{ID: 140, Star: 5, Title: "思摩夫", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-6987KdT8S28-28.png"},
		{ID: 141, Star: 5, Title: "神威", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201912/02/hbQ5-k3c4KeT8S28-28.png"},
		{ID: 142, Star: 5, Title: "古里曼", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202002/28/hbQ5-3px3KeT8S28-28.png"},
		{ID: 143, Star: 5, Title: "格布纽&克雷纽", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202004/14/hbQ5-1zfcKbT8S28-28.png"},
		{ID: 144, Star: 5, Title: "瓦基扬", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-be6mKdT8S28-28.png"},
		{ID: 145, Star: 5, Title: "迦楼罗", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-c3jpKcT8S28-28.png"},
		{ID: 146, Star: 5, Title: "龙龙", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-6kjmKeT8S28-28.png"},
		{ID: 147, Star: 5, Title: "帕祖祖", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/19/hbQ5-df1cKeT8S28-28.png"},
		{ID: 148, Star: 5, Title: "芙蕾雅", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201906/20/hbQ5-ksrxKfT8S28-28.png"},
		{ID: 149, Star: 5, Title: "伐由", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201908/31/hbQ5-hwj9KfT8S28-28.png"},
		{ID: 150, Star: 5, Title: "哈斯塔", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201911/05/hbQ5-csyzKfT8S28-28.png"},
		{ID: 151, Star: 5, Title: "AC-011 加兰德", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202001/15/hbQ5-c6dmKeT8S28-28.png"},
		{ID: 152, Star: 5, Title: "亚列", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202003/12/hbQ5-k6g5KdT8S28-28.png"},
		{ID: 153, Star: 5, Title: "贞德", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-hjsgKdT8S28-28.png"},
		{ID: 154, Star: 5, Title: "神圣菲尼克斯", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201912/12/hbQ5-41l8KdT8S28-28.png"},
		{ID: 155, Star: 5, Title: "吉尔伽美什", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-557mKdT8S28-28.png"},
		{ID: 156, Star: 5, Title: "丘比特", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-8v4tKdT8S28-28.png"},
		{ID: 157, Star: 5, Title: "雷牙", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-kzfiKeT8S28-28.png"},
		{ID: 158, Star: 5, Title: "建遇雷", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201905/14/hbQ5-951nKeT8S28-28.png"},
		{ID: 159, Star: 5, Title: "塞壬(演出Vers.)", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201908/13/hbQ5-g348KfT8S28-28.png"},
		{ID: 160, Star: 5, Title: "尼德霍格", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-bhb9KdT8S28-28.png"},
		{ID: 161, Star: 5, Title: "奈亚拉托提普", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-h2c7KdT8S28-28.png"},
		{ID: 162, Star: 5, Title: "忍", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201906/30/hbQ5-gevKeT8S28-28.png"},
		{ID: 163, Star: 5, Title: "普鲁托", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201909/30/hbQ5-iijzKfT8S28-28.png"},
		{ID: 164, Star: 5, Title: "埃庇米修斯", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202002/14/hbQ5-d8zmKfT8S28-28.png"},
		{ID: 165, Star: 5, Title: "安德洛墨达", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202004/02/hbQ5-2j54KcT8S28-28.png"},
		{ID: 166, Star: 4, Title: "菲尼克斯", Water: 2200, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-6e29KdT8S28-28.png"},
		{ID: 167, Star: 4, Title: "伊芙利特", Water: 2200, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-foqmKdT8S28-28.png"},
		{ID: 168, Star: 4, Title: "波利亚芙", Water: 2200, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-i0azKbT8S28-28.png"},
		{ID: 169, Star: 4, Title: "沃家诺伊", Water: 2200, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-p0iKdT8S28-28.png"},
		{ID: 170, Star: 4, Title: "沃格尔", Water: 2200, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-5t4xKcT8S28-28.png"},
		{ID: 171, Star: 4, Title: "独角兽", Water: 2200, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-aqyeKcT8S28-28.png"},
		{ID: 172, Star: 4, Title: "林德巴鲁姆", Water: 2200, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-few0KcT8S28-28.png"},
		{ID: 173, Star: 4, Title: "扎格纳特", Water: 2200, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-klrjKdT8S28-28.png"},
		{ID: 174, Star: 3, Title: "火龙", Water: 150, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-u98KbT8S28-28.png"},
		{ID: 175, Star: 3, Title: "焰龙", Water: 150, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-l9plKbT8S28-28.png"},
		{ID: 176, Star: 3, Title: "烛魔", Water: 150, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-4h9dKcT8S28-28.png"},
		{ID: 177, Star: 3, Title: "雪龙", Water: 150, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-gkfbKaT8S28-28.png"},
		{ID: 178, Star: 3, Title: "冰龙", Water: 150, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-23p6KbT8S28-28.png"},
		{ID: 179, Star: 3, Title: "井魔", Water: 150, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-6zejKcT8S28-28.png"},
		{ID: 180, Star: 3, Title: "旋龙", Water: 150, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-ffbgKbT8S28-28.png"},
		{ID: 181, Star: 3, Title: "风龙", Water: 150, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-hmjyKbT8S28-28.png"},
		{ID: 182, Star: 3, Title: "气魔", Water: 150, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-112cKbT8S28-28.png"},
		{ID: 183, Star: 3, Title: "阳龙", Water: 150, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-7tvgKaT8S28-28.png"},
		{ID: 184, Star: 3, Title: "月龙", Water: 150, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-1wrtKbT8S28-28.png"},
		{ID: 185, Star: 3, Title: "星魔", Water: 150, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-dg0sKbT8S28-28.png"},
		{ID: 186, Star: 3, Title: "暗龙", Water: 150, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-i4ajKaT8S28-28.png"},
		{ID: 187, Star: 3, Title: "墓魔", Water: 150, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-18cmKbT8S28-28.png"},
		{ID: 188, Star: 4, Title: "斯特丽伯格", Water: 2200, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-gx9iKcT8S28-28.png"},
		{ID: 189, Star: 5, Title: "尤里乌斯", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202005/12/hbQ5-5d4iKdT8S28-28.png"},
		{ID: 190, Star: 5, Title: "法尔提", Water: 8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202005/12/hbQ5-bcuwKdT8S28-28.png"},
		{ID: 191, Star: 5, Title: "阿撒兹勒", Water: 8500, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202005/12/hbQ5-h8u0KdT8S28-28.png"},
		{ID: 192, Star: 5, Title: "思摩夫(天堂Ver.)", Water: 85000, CardType: 2, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202005/15/hbQ5-bn3dK9T8S28-28.png"},
		{ID: 193, Star: 5, Title: "库拉乌(琴酒Ver.)", Water: -8500, CardType: 1, rareType: 1,
			IconUrl: "https://img.nga.178.com/attachments/mon_202005/15/hbQ5-c912K8T8S28-28.png"},
		{ID: 194, Star: 5, Title: "琪姬", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_202005/04/hbQ5-1dkgKdT8S28-28.png"},
		{ID: 195, Star: 5, Title: "库洛姆", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_202004/30/hbQ5-5eqxKdT8S28-28.png"},
		{ID: 196, Star: 5, Title: "皮亚尼", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_202004/30/hbQ5-fp0zKdT8S28-28.png"},
		{ID: 197, Star: 5, Title: "马尔斯", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/27/hbQ5-j4djKdT8S28-28.png"},
		{ID: 198, Star: 5, Title: "菲约尔姆", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/27/hbQ5-a2mwKdT8S28-28.png"},
		{ID: 199, Star: 5, Title: "维罗妮卡", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/27/hbQ5-kibyKdT8S28-28.png"},
		{ID: 200, Star: 5, Title: "阿迪斯(情人节Ver.)", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_202002/14/hbQ5-itx3KeT8S28-28.png"},
		{ID: 201, Star: 5, Title: "米罗蒂(情人节Ver.)", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_202002/14/hbQ5-2yhgKfT8S28-28.png"},
		{ID: 202, Star: 5, Title: "希里丝(猎人Ver.)", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_202002/04/hbQ5-4grbKfT8S28-28.png"},
		{ID: 203, Star: 5, Title: "黑炎王", Water: 8500, CardType: 2, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_202002/04/hbQ5-k969KfT8S28-28.png"},
		{ID: 204, Star: 5, Title: "贝尔扎克(猎人Ver.)", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_202001/29/hbQ5-bt6kKfT8S28-28.png"},
		{ID: 205, Star: 5, Title: "凡妮莎(猎人Ver.)", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_202001/29/hbQ5-hvwiKfT8S28-28.png"},
		{ID: 206, Star: 5, Title: "黑龙", Water: 8500, CardType: 2, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_202001/29/hbQ5-atmKeT8S28-28.png"},
		{ID: 207, Star: 5, Title: "光秀", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201912/31/hbQ5-l2ykKfT8S28-28.png"},
		{ID: 208, Star: 5, Title: "信长", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201912/31/hbQ5-3qclKeT8S28-28.png"},
		{ID: 209, Star: 4, Title: "千岁", Water: 2200, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201912/31/hbQ5-9iaiKfT8S28-28.png"},
		{ID: 210, Star: 5, Title: "家康", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-drhpKdT8S28-28.png"},
		{ID: 211, Star: 4, Title: "阿迪斯", Water: 2200, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-3q2zKcT8S28-28.png"},
		{ID: 212, Star: 4, Title: "山茶花", Water: 2200, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-2dh9KdT8S28-28.png"},
		{ID: 213, Star: 5, Title: "大黑天", Water: 8500, CardType: 2, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201912/31/hbQ5-5ektKeT8S28-28.png"},
		{ID: 214, Star: 5, Title: "摩利支天", Water: 8500, CardType: 2, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-6ujqKcT8S28-28.png"},
		{ID: 215, Star: 5, Title: "马萝拉(星龙祭Ver.)", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201912/12/hbQ5-eiknKdT8S28-28.png"},
		{ID: 216, Star: 5, Title: "库拉乌(星龙祭Ver.)", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-jeoeKdT8S28-28.png"},
		{ID: 217, Star: 5, Title: "贞德(星龙祭Ver.)", Water: 8500, CardType: 2, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-b12qKdT8S28-28.png"},
		{ID: 218, Star: 4, Title: "扎因弗拉德(星龙祭Ver.)", Water: 2200, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201912/12/hbQ5-jyk8KdT8S28-28.png"},
		{ID: 219, Star: 4, Title: "奈法利耶(星龙祭Ver.)", Water: 2200, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-ddc7KcT8S28-28.png"},
		{ID: 220, Star: 4, Title: "阿莱克西斯(星龙祭Ver.)", Water: 2200, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-l3wsKdT8S28-28.png"},
		{ID: 221, Star: 5, Title: "穆穆(万圣节Ver.)", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201910/19/hbQ5-6edzKdT8S28-28.png"},
		{ID: 222, Star: 5, Title: "埃尔菲利斯(万圣节Ver.)", Water: 8500, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-860bKdT8S28-28.png"},
		{ID: 223, Star: 5, Title: "玛利蒂姆斯(万圣节Ver.)", Water: 8500, CardType: 2, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201910/19/hbQ5-20e3K9T8S28-28.png"},
		{ID: 224, Star: 4, Title: "欧蒂塔(万圣节Ver.)", Water: 2200, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201910/19/hbQ5-hk4xKdT8S28-28.png"},
		{ID: 225, Star: 4, Title: "伊露提米娅(万圣节Ver.)", Water: 2200, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-j7ipKdT8S28-28.png"},
		{ID: 226, Star: 4, Title: "希露姬(万圣节Ver.)", Water: 2200, CardType: 2, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201904/04/hbQ5-j2trKdT8S28-28.png"},
		{ID: 227, Star: 3, Title: "爱德华(万圣节Ver.)", Water: 150, CardType: 1, rareType: 2,
			IconUrl: "https://img.nga.178.com/attachments/mon_201902/23/-mjkdpQ5-c43dKcT8S28-28.png"},
	}
	CardCollection = map[int]*CardSet{}
	CardCollection[512] = &CardSet{star: 5, cardType: 1, rareType: 2, Prob: 100}
	CardCollection[522] = &CardSet{star: 5, cardType: 2, rareType: 2, Prob: 80}
	CardCollection[511] = &CardSet{star: 5, cardType: 1, rareType: 1, Prob: 100}
	CardCollection[521] = &CardSet{star: 5, cardType: 2, rareType: 1, Prob: 120}

	CardCollection[412] = &CardSet{star: 4, cardType: 1, rareType: 2, Prob: 350}
	CardCollection[422] = &CardSet{star: 4, cardType: 2, rareType: 2, Prob: 350}
	CardCollection[411] = &CardSet{star: 4, cardType: 1, rareType: 1, Prob: 505}
	CardCollection[421] = &CardSet{star: 4, cardType: 2, rareType: 1, Prob: 395}

	CardCollection[311] = &CardSet{star: 3, cardType: 1, rareType: 1, Prob: 4800}
	CardCollection[321] = &CardSet{star: 3, cardType: 2, rareType: 1, Prob: 3200}

	for _, card := range Cards {
		key := 100*card.Star + 10*card.CardType + card.rareType
		if _, ok := CardCollection[key]; ok {
			CardCollection[key].cards = append(CardCollection[key].cards, card)
		}
	}
}

func SummonOne(key int) Card {
	return getCardPool(key).PickOne()
}

func (c *CardSet) PickOne() Card {
	r := rand.Intn(len(c.cards))
	return c.cards[r]
}

func FindCardIndex(name []string) []int {
	res := make([]int, len(name))
	for i := range res {
		res[i] = -1
	}
	for _, card := range Cards {
		for i := range name {
			if card.Title == name[i] {
				res[i] = card.ID
			}
		}
	}
	return res
}

func GetCardsNum(key int) int {
	return len(getCardPool(key).cards)
}
func GetCardsNumByStarType(star, cardType int) int {
	res := 0
	for _, card := range Cards {
		if card.Star == star && card.CardType == cardType {
			res++
		}
	}
	return res
}

// 返回中 0代表5星人 1代表4星人 2代表3星人 3代表5星龙一次类推
func GetCardsAnalysis(cardIndex []int) []int {
	res := make([]int, 6)
	for _, i := range cardIndex {
		switch Cards[i].Star*10 + Cards[i].CardType {
		case 51:
			res[0]++
		case 41:
			res[1]++
		case 31:
			res[2]++
		case 52:
			res[3]++
		case 42:
			res[4]++
		case 32:
			res[5]++
		}
	}
	return res
}

func getCardPool(key int) *CardSet {
	if value, ok := CardCollection[key]; ok {
		return value
	} else {
		log.Panic("could not find")
	}
	return nil
}

func (c CardCollectionV2) PickUpByStar(star int) []*CardSet {
	var res []*CardSet
	for _, set := range c.cardSets {
		if set.star == star {
			res = append(res, set)
		}
	}
	return res
}