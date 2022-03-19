package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

var count, sixs, fives, 六星未出计数 = 0, 0, 0, 0
var 变化概率, 六星变化概率 = 0, 0
var 保底 = false

func init() {
	六星up("银灰")
	六星up("塞雷亚")
	六星("推进之王")
	六星("能天使")
	六星("伊芙利特")
	六星("艾雅法拉")
	六星("安洁丽娜")
	六星("闪灵")
	六星("夜莺")
	六星("星熊")
	六星("斯卡蒂")
	六星("陈")
	六星("黑")
	六星("赫拉格")
	六星("麦哲伦")
	六星("莫斯提马")
	六星("煌")
	六星("阿")
	六星("刻俄柏")
	五星("白面鸮")
	五星("凛冬")
	五星("德克萨斯")
	五星("芙兰卡")
	五星("拉普兰德")
	五星("幽灵鲨")
	五星up("蓝毒")
	五星("白金")
	五星("陨星")
	五星("天火")
	五星("梅尔")
	五星("赫默")
	五星("华法琳")
	五星("临光")
	五星("红")
	五星("雷蛇")
	五星("可颂")
	五星("普罗旺斯")
	五星("守林人")
	五星up("崖心")
	五星("初雪")
	五星("真理")
	五星("空")
	五星("惊蛰")
	五星("狮蝎")
	五星("食铁兽")
	五星("夜魔")
	五星("诗怀雅")
	五星("格劳克斯")
	五星("星极")
	五星("送葬人")
	五星("槐琥")
	五星up("苇草")
	五星("布洛卡")
	五星("灰喉")
	五星("哞")
	四星("夜烟")
	四星("远山")
	四星("杰西卡")
	四星("白雪")
	四星("清道夫")
	四星("红豆")
	四星("杜宾")
	四星("缠丸")
	四星("霜叶")
	四星("慕斯")
	四星("砾")
	四星("暗索")
	四星("末药")
	四星("调香师")
	四星("角峰")
	四星("蛇屠箱")
	四星("古米")
	四星("深海色")
	四星("地灵")
	四星("阿消")
	四星("猎蜂")
	四星("格雷伊")
	四星("苏苏洛")
	四星("桃金娘")
	四星("红云")
	四星("梅")
	四星("安比尔")
	三星("芬")
	三星("香草")
	三星("翎羽")
	三星("玫兰莎")
	三星("卡缇")
	三星("米格鲁")
	三星("克洛丝")
	三星("炎熔")
	三星("芙蓉")
	三星("安塞尔")
	三星("史都华德")
	三星("梓兰")
	三星("空爆")
	三星("月见夜")
	三星("斑点")
	三星("泡普卡")
	for _, 干员 := range 干员们 {
		总干员概率 += 干员.概率
	}
}
func exit(input *bufio.Scanner) {
	if input.Text() == "exit" {
		warning("寻访模拟结束！感谢您的使用")
		colorreset()
		os.Exit(0)
	}
}

func warning(msg string) {
	colorout("\n=============================\n\n", 字体颜色.亮青色)
	colorout(msg, 字体颜色.亮红色)
	colorout("\n\n=============================\n\n", 字体颜色.亮青色)
}
func main() {
	colorout(time.Now().Format("2006-01-02 15:04:05\n"), 字体颜色.亮蓝色)
	input := bufio.NewScanner(os.Stdin)
	for {
		sixs = 0
		fives = 0
		colorout("请输入要抽的数量，", 字体颜色.亮蓝色)
		colorout("输入 “exit” 表示退出程序\n", 字体颜色.亮蓝色)
		colorreset()
		input.Scan()
		exit(input)
		if !regexp.MustCompile(`^\d+$`).MatchString(input.Text()) {
			warning("输入非法!")
			colorreset()
		}
		num, _ := strconv.Atoi(input.Text())

		for i := 1; i <= num; i++ {
			开抽(保底)
			count++
			if count%10 == 0 {
				println()
				println()
			}
			if 六星未出计数 >= 50 {
				变化概率++
				六星变化概率 += 2
			}
		}
		if num < 10 {
			println()
		}
		colorout("共计：", 字体颜色.亮蓝色)
		println()
		colorreset()
		colorout("六星数量:", 字体颜色.亮黄色)
		println(sixs)
		colorout("五星数量:", 字体颜色.黄色)
		println(fives)
	}

}

func 开抽(a bool) {
	抽中的范围 := 0
	for _, 干员 := range 干员们 {
		rand.Seed(time.Now().UnixNano())
		抽中的干员 := rand.Intn(总干员概率) //必须为100，所以需要动态概率，尤其是五十发之后
		抽中的范围 += 干员.概率
		if 抽中的干员 <= 抽中的范围 {
			//println("你抽中了", 干员.星级, "星的", 干员.名字)
			if 干员.星级 < 5 && 保底 == false {
				干员.星级 = 5
			}
			switch 干员.星级 {
			case 6:
				colorout(干员.名字, 字体颜色.亮黄色)
				fmt.Printf(" ")
				colorreset()
				sixs++
				六星未出计数 = 0
				保底 = true
			case 5:
				colorout(干员.名字, 字体颜色.黄色)
				fmt.Printf(" ")
				colorreset()
				fives++
				保底 = true
			case 4:
				colorout(干员.名字, 字体颜色.亮紫色)
				fmt.Printf(" ")
				colorreset()
			case 3:
				colorout(干员.名字, 字体颜色.白色)
				fmt.Printf(" ")
				colorreset()
			}
			break
		}
	}
}
