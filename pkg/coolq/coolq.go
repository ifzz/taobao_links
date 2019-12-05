package coolq

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/wppzxc/taobao_links/pkg/coolq/app"
	"strings"
)

type CoolQ struct {
	MainPage     *TabPage
	WebSocketUrl *walk.LineEdit
	Groups       *walk.TextEdit
	Users        *walk.TextEdit
	AutoImport   *walk.PushButton
	Start        *walk.PushButton
	Stop         *walk.PushButton
	StopCh       chan struct{}
}

func GetCoolQPage() *CoolQ {
	coolq := &CoolQ{
		WebSocketUrl: &walk.LineEdit{},
		Groups:       &walk.TextEdit{},
		Users:        &walk.TextEdit{},
		AutoImport:   &walk.PushButton{},
		Start:        &walk.PushButton{},
		Stop:         &walk.PushButton{},
	}
	coolq.MainPage = &TabPage{
		Title:  "QQ消息转发微信",
		Layout: VBox{},
		DataBinder: DataBinder{
			DataSource: coolq,
			AutoSubmit: true,
		},
		Children: []Widget{
			Composite{
				Layout: VBox{},
				Children: []Widget{
					HSpacer{},
					TextLabel{
						Text: "酷Q websocket 地址（默认使用本地）：",
					},
					LineEdit{
						AssignTo: &coolq.WebSocketUrl,
					},
				},
			},
			Composite{
				Layout: VBox{},
				Children: []Widget{
					HSpacer{},
					TextLabel{
						Text: "接收群号（多个群组用 / 分隔）：",
					},
					TextEdit{
						AssignTo: &coolq.Groups,
					},
				},
			},
			Composite{
				Layout: VBox{},
				Children: []Widget{
					HSpacer{},
					Composite{
						Layout: HBox{},
						Children: []Widget{
							TextLabel{
								Text: "转发列表（多个用户用 / 分隔）：",
							},
							PushButton{
								Text:      "自动获取",
								AssignTo:  &coolq.AutoImport,
								OnClicked: coolq.AutoImportUsers,
							},
						},
					},
					TextEdit{
						AssignTo: &coolq.Users,
						VScroll:  true,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text:      "开始",
						AssignTo:  &coolq.Start,
						OnClicked: coolq.StartWork,
					},
					PushButton{
						Text:      "停止",
						AssignTo:  &coolq.Stop,
						OnClicked: coolq.StopWork,
						Enabled:   false,
					},
				},
			},
		},
	}
	return coolq
}

func (c *CoolQ) AutoImportUsers() {
	fmt.Println("auto import users")
}

func (c *CoolQ) StartWork() {
	fmt.Println("start coolq work!")
	wsUrl := c.WebSocketUrl.Text()
	if len(wsUrl) == 0 {
		fmt.Println("Must provide websocket url!")
		return
	}
	groups := c.GetGroups()
	if len(groups) == 0{
		fmt.Println("Must provide groups!")
		return
	}
	users := c.GetUsers()
	if len(users) == 0 {
		fmt.Println("Must provide user!")
		return
	}
	c.StopCh = make(chan struct{})
	c.SetUIEnable(false)
	app.Start(wsUrl, groups, users, c.StopCh)
	fmt.Println("coolq started!")
}

func (c *CoolQ) StopWork() {
	fmt.Println("stop coolq work")
	close(c.StopCh)
	c.SetUIEnable(true)
}

func (c *CoolQ) GetGroups() []string {
	data := c.Groups.Text()
	if len(data) == 0 {
		return nil
	}
	groups := strings.Split(data, "/")
	return groups
}

func (c *CoolQ) GetUsers() []string {
	data := c.Users.Text()
	if len(data) == 0 {
		return nil
	}
	users := strings.Split(data, "/")
	return users
}

func (c *CoolQ) SetUIEnable(enable bool) {
	c.Groups.SetEnabled(enable)
	c.Users.SetEnabled(enable)
	c.Start.SetEnabled(enable)
	c.WebSocketUrl.SetEnabled(enable)
	c.AutoImport.SetEnabled(enable)
	// stop is unEnable with others
	c.Stop.SetEnabled(!enable)
}
