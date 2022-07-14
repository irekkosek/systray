package main

import (
	"SysTray/icon"
	"github.com/getlantern/systray"
)

func main() {
	systray.Register(onReady, nil)
	configureWebview("Webview example", 1024, 768)
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Webview example")
	mShowGoogle := systray.AddMenuItem("Show Google", "")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		for {
			select {
			case <-mShowGoogle.ClickedCh:
				showWebview("https://www.google.com")
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()

}
