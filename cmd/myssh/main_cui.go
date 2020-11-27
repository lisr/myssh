package main

import (
	"fmt"
	"os"

	"github.com/lisr/myssh/internal/conf"
	"github.com/lisr/myssh/internal/cui"
	"github.com/lisr/myssh/internal/rsession"
)

func cuiMain() {
	cfg := conf.LoadMySSHConfig()

	var items []*cui.ListViewItem

	for _, item := range cfg.Servers {
		items = append(items, &cui.ListViewItem{
			Key:   item.Host,
			Value: fmt.Sprintf("[%s](fg:blue) \t%s", item.Host, item.Desc),
		})
	}

	lv := cui.NewListView(items)

	selection, err := lv.Run()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	if selection == nil {
		fmt.Println("no selection")
		os.Exit(1)
	}

	var server conf.ServerCfg
	for _, item := range cfg.Servers {
		if item.Host == selection.Key {
			server = item
		}
	}

	fmt.Println("Connecting to ...")
	fmt.Println("👉 ", server.Host, " 🐷 ", server.Desc)
	fmt.Println("--------------------------------------------")

	// fmt.Println(pad.Right(selection.Key, 50, "~"))

	cred := rsession.ParseCred(server.Cred)
	if cred == nil {
		fmt.Println("invalid credential")
		os.Exit(1)
	}

	rsession.ConnectSSH(server.Host, cred)
}
