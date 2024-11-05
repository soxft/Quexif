package main

import (
	"fmt"
	"log"
	"photo_exif_do/fg"
	"photo_exif_do/qumagie"
	"photo_exif_do/x_exif"
)

func main() {
	fg.Parse()

	// 安全 QA
	if fg.SkipSafeQA {
		log.Println("请确保已经设置了快照，程序将会直接修改文件的 exif 元数据, 建议在使用前选择少量照片进行测试后再使用")
		log.Println("您确认要继续吗? (y/n)")
		var confirm string

		if _, err := fmt.Scanln(&confirm); err != nil {
			return
		} else if confirm != "y" {
			log.Fatal("已取消")
		}
	}

	// 读取目录
	switch fg.Mode {
	case "dir": // 指定文件夹批量修改
		log.Println("Not implemented yet")
		break
	case "dirDate": //按照上级文件夹名称修改
		log.Println("Not implemented yet")
		break
	case "test":
		// log.Println(x_exif.SetDate("pics/qumagie/2006-01-02 15.04.05.jpg", time.Now(), true))

		log.Println(x_exif.ReadExif("pics/qumagie/2022-06-15 10.13.50.png"))
	default:
		qumagie.Run(fg.Path)
	}

}
