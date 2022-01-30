package response

func Email(email string) {
	// // 简单设置 log 参数
	// log.SetFlags(log.Lshortfile | log.LstdFlags)

	// em := email.NewEmail()
	// // 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	// em.From = "xx <2695009886@qq.com>"

	// // 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	// em.To = []string{}

	// // 设置主题
	// em.Subject = "黄雪朋写的物件给你发个验证码, 你一会填一下哈"

	// // 简单设置文件发送的内容，暂时设置成纯文本
	// em.Text = []byte("hello world， 我是黄雪朋, 给你发邮件了！仅仅是测试,, 不要紧张")
	// em.Text = []byte()
	// // em.HTML = ``
	// //设置服务器相关的配置
	// err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "2695009886@qq.com", "pndftuyzutbcdhff", "smtp.qq.com"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("send successfully ... ")
}
