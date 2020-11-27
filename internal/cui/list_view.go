package cui

import (
	"fmt"
	"log"
	"sort"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// ListView list view
type ListView struct {
	selection *ListViewItem
	list      []*ListViewItem
}

// ListViewItem list view item
type ListViewItem struct {
	Key   string
	Value string
}

type _ListViewItemList []*ListViewItem

func (lvl _ListViewItemList) Len() int {
	return len(lvl)
}
func (lvl _ListViewItemList) Less(i, j int) bool {
	return strings.Compare(lvl[i].Value, lvl[j].Value) < 0
}

func (lvl _ListViewItemList) Swap(i, j int) {
	lvl[i], lvl[j] = lvl[j], lvl[i]
}

// NewListView create new list view
func NewListView(list []*ListViewItem) *ListView {

	sort.Sort(_ListViewItemList(list))

	return &ListView{
		list: list,
	}
}

// Run show list view and choose one
func (lv *ListView) Run() (*ListViewItem, error) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	l := widgets.NewList()
	l.Title = "Remote Hosts"

	for idx, item := range lv.list {
		l.Rows = append(l.Rows, fmt.Sprintf("[%d] %s", idx, item.Value))
	}

	l.TextStyle = ui.NewStyle(ui.ColorClear, ui.ColorClear)
	l.SelectedRowStyle = ui.NewStyle(ui.ColorClear, ui.ColorCyan, ui.ModifierBold)
	l.WrapText = false

	termWidth, termHeight := ui.TerminalDimensions()
	l.SetRect(0, 0, termWidth, termHeight)

	ui.Render(l)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return nil, nil
		case "<Enter>":
			return lv.list[l.SelectedRow], nil
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<C-f>":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Home>":
			l.ScrollTop()
		case "G", "<End>":
			l.ScrollBottom()
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		ui.Render(l)
	}
}
