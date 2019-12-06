module github.com/wppzxc/taobao_links

go 1.13

require (
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/atotto/clipboard v0.1.2
	github.com/go-vgo/robotgo v0.0.0-20191128163956-6b94d024dc37
	github.com/lxn/walk v0.0.0-20191128110447-55ccb3a9f5c1
	github.com/lxn/win v0.0.0-20191128105842-2da648fda5b4
	github.com/mvdan/xurls v1.1.0
	github.com/shirou/w32 v0.0.0-20160930032740-bb4de0191aa4
	github.com/tealeg/xlsx v1.0.5
	github.com/zserge/lorca v0.1.8
	github.com/zserge/webview v0.0.0
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
	gopkg.in/Knetic/govaluate.v3 v3.0.0 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
	sigs.k8s.io/yaml v1.1.0
)

replace github.com/zserge/webview v0.0.0 => ./libs/webview