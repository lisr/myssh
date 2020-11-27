package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/lisr/myssh/internal/conf"
)

func mainLs() {
	cfg := conf.LoadMySSHConfig()
	if cfg == nil || len(cfg.Servers) == 0 {
		fmt.Println("no server found")
	}

	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 48, 8, 0, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "%s\t%s\t", "Host", "Description")
	fmt.Fprintf(w, "\n%s\t%s\t", "----", "----")

	for _, item := range cfg.Servers {
		fmt.Fprintf(w, "\n%s\t%s\t", item.Host, item.Desc)
	}
	fmt.Fprintln(w, "\n.")
}
