package main

import (
	"flag"
	"fmt"
)

func main() {
	var cin string
	const tangshi = "1"
	const songci = "2"
	const help = `
程序名称：main2
描述:输入1或2，输入1时输出唐诗，输入2时输出宋词
`
	flag.Usage = func() {
		fmt.Printf(help)
	}
	flag.Parse()
	fmt.Println("输入1或2:")
	if _, err := fmt.Scanf("%s", &cin); err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	if tangshi == cin {
		fmt.Printf("秦时明月汉时关，万里长征人未还。\n但是龙城飞将在，不教胡马度阴山。")
		return
	}
	if songci == cin {
		fmt.Printf("薄雾浓云愁永昼，瑞脑消金兽。佳节又重阳，玉枕纱厨，半夜凉初透。\n东篱把酒黄昏后，有暗香盈袖。莫道不销魂，帘卷西风，人比黄花瘦。")
		return
	}
}
