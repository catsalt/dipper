package main

import (
	"fmt"
	"strings"

	// "strconv"

	"ziwei/tanlang"
	//"github.com/catsalt/zwds"

	"fyne.io/fyne"
	"fyne.io/fyne/app"

	// "fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type Person struct {
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Calendar string `json:"calendar"`
	Birthday
}
type Birthday struct {
	Year   string `json:"year"`
	Month  string `json:"month"`
	Day    string `json:"day"`
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
}
type Persons []Person

var demo = []string{"Example", "Male", "Solar",
	"1", "9", "9", "9", "0", "9", "0", "9", "0", "9", "0", "9"}
var infoJson = "data/infoes.txt"

func aGnewForm(w fyne.Window, demo []string, infoes [][]string) [][]string {
	if len(demo) != 15 {
		fmt.Println("!!! format wrong.")
		return infoes
	}
	labels := []string{"Name:", "Gender:", "Calendar:", "Year:", "Mon:", "Day:", "Hour:", "Min:"}
	nums := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	name := &widget.Entry{Text: demo[0]}
	gender := &widget.Radio{Options: []string{"Male", "Female"}, Selected: demo[1], Horizontal: true}
	calendar := &widget.Radio{Options: []string{"Solar", "Lunar"}, Selected: demo[2], Horizontal: true}

	date := []*widget.Select{
		&widget.Select{Options: nums[0:3], Selected: demo[3]},
		&widget.Select{Options: nums, Selected: demo[4]},
		&widget.Select{Options: nums, Selected: demo[5]},
		&widget.Select{Options: nums, Selected: demo[6]},
		&widget.Select{Options: nums[0:2], Selected: demo[7]},
		&widget.Select{Options: nums, Selected: demo[8]},
		&widget.Select{Options: nums[0:4], Selected: demo[9]},
		&widget.Select{Options: nums, Selected: demo[10]},
		&widget.Select{Options: nums[0:3], Selected: demo[11]},
		&widget.Select{Options: nums, Selected: demo[12]},
		&widget.Select{Options: nums[0:6], Selected: demo[13]},
		&widget.Select{Options: nums, Selected: demo[14]}}
	submit := widget.NewButton("SUBMIT", func() {
		info := make([]string, 15)
		info[0], info[1], info[2] = name.Text, gender.Selected, calendar.Selected
		for k, v := range date {
			info[k+3] = v.Selected
		}

		infoes = append(infoes, info)
		fmt.Println(infoes)
		return
	})
	exit := widget.NewButton("EXIT", func() {
		w.Hide()
		return
	})
	n := widget.NewHBox(widget.NewLabel(labels[0]), name)
	g := widget.NewHBox(widget.NewLabel(labels[1]), gender)
	c := widget.NewHBox(widget.NewLabel(labels[2]), calendar)
	year := widget.NewHBox(
		widget.NewLabel(labels[3]), date[0], date[1], date[2], date[3])
	monthDay := widget.NewHBox(
		widget.NewLabel(labels[4]), date[4], date[5],
		widget.NewLabel(labels[5]), date[6], date[7])
	hourMinute := widget.NewHBox(
		widget.NewLabel(labels[6]), date[8], date[9],
		widget.NewLabel(labels[7]), date[10], date[11])
	w.SetContent(fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), n),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), g),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), c),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), year),
		// year,
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), monthDay),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), hourMinute),
		fyne.NewContainerWithLayout(layout.NewGridLayout(2), submit, exit),
	))
	w.Show()
	fmt.Println("\r\nRuning...")
	return infoes
}
func bGfromInfoes(infoes [][]string) (str string) {
	for _, v := range infoes {
		str += strings.Join(v, " ") + "\r\n"
	}
	return str
}
func cGinit() {
	a := app.New()
	w := a.NewWindow("Add Personal Info")
	u := a.NewWindow("Persons")
	a.Settings().SetTheme(theme.LightTheme())
	var infoes [][]string
	tanlang.ZfreadJson(&infoes, infoJson)
	listInfo := widget.NewMultiLineEntry()
	listInfo.SetText(bGfromInfoes(infoes))

	addInfo := widget.NewButton("ADD", func() {
		infoes = aGnewForm(w, demo, infoes)
		return
	})
	quit := widget.NewButton("QUIT", func() { a.Quit() })
	u.SetContent(fyne.NewContainerWithLayout(
		layout.NewGridLayout(1), listInfo,
		fyne.NewContainerWithLayout(layout.NewGridLayout(2), addInfo, quit)))
	u.ShowAndRun()
	tanlang.ZfwriteJson(infoes, infoJson)
	u.SetOnClosed(func() { a.Quit() })
	return
}

func dGpersonFrom(info []string) Person {
	return Person{Name: info[0], Gender: info[1], Calendar: info[2], Birthday: Birthday{
		Year: info[3] + info[4] + info[5] + info[6], Month: info[7] + info[8],
		Day: info[9] + info[10], Hour: info[11] + info[12], Minute: info[13] + info[14]}}
}
func main() {
	//	var Hero ziwei.zwds.Hero
	// tanlang.ZfzxPan()
	// cGinit()
	test()
}
func test() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("test")

	g := widget.NewGroup("Test", widget.NewLabel("Hello"), widget.NewLabel("world"))

	// w.SetContent(g)
	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(1), g, g, g))
	w.ShowAndRun()
	return
}
