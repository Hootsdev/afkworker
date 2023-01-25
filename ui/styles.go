package ui

import "github.com/charmbracelet/lipgloss"

const (
	listHeight   = 20
	defaultWidth = 200
)
const (
	header = "<|!|> AFK Worker v0.1 <|!|>"
)
const (
	hotPink     = lipgloss.Color("#FF06B7")
	black       = lipgloss.Color("0")
	white       = lipgloss.Color("#FFFFFF")
	darkGray    = lipgloss.Color("#767676")
	purple      = lipgloss.Color("99")
	brightGreen = lipgloss.Color("#00FF00")
	bloodRed    = lipgloss.Color("#FF0000")
	someG       = lipgloss.Color("#00FFa0")
	someR       = lipgloss.Color("#FFa000")
	orange      = lipgloss.Color("#faa805")
	sep         = "\n>>>"
)

var (
	//////////////
	/// LEFT /////
	// Panel ////
	////////////

	// Title
	tStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(white).
		ColorWhitespace(true).
		Align(lipgloss.Right)

	// list Title Bar
	tbStyle = lipgloss.NewStyle().
		ColorWhitespace(true).
		MarginBottom(1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(hotPink)

	menulistStyle = lipgloss.NewStyle().
			ColorWhitespace(true).
			Align(lipgloss.Left).
			Margin(0, 0, 0, 2).
			Width(50)

	///////////////
	/// RIGHT ////
	// Panel ////
	////////////
	statusStyle = lipgloss.NewStyle().
			MarginLeft(1).
			Border(lipgloss.RoundedBorder()).
			Bold(true).
			Width(50).
			PaddingLeft(3).
			PaddingRight(5).
			Align(lipgloss.Right, lipgloss.Top).
			BorderForeground(bloodRed)

	spinnerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("69"))

	runnunTaskStyle = statusStyle.Copy().
			MarginTop(2).
			UnsetPaddingLeft().
			Width(55).
			UnsetBorderStyle().Align(lipgloss.Left, lipgloss.Bottom)

	//	happyClr = colorful.FastHappyColor()
	helpStyle = lipgloss.NewStyle().
			MarginLeft(1).
			Align(lipgloss.Bottom).
			Foreground(darkGray)

	quitStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF0000")).
			MarginBackground(lipgloss.Color("#00FF00")).
			Margin(10)
	// hz
	// execRespStyle = lipgloss.NewStyle().
	// 		Border(lipgloss.RoundedBorder()).
	// 		BorderForeground(hotPink).
	// 		Foreground(lipgloss.Color("#77DE77")).
	// 		Align(lipgloss.Bottom).
	// 		MarginLeft(30)

	taskName = lipgloss.NewStyle().Foreground(orange)
)

// ////////////////
// / settings ////
// // input /////
// /////////////
var (
	// MultiText Input Form
	topInputStyle = menulistStyle.Copy().
		// Margin(10, 0, 10, 2).
		// Width(45)
		MarginTop(10)
	focusedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))

	cursorStyle = focusedStyle.Copy()

	noStyle = lipgloss.NewStyle()

	cursorModeHelpStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Copy().
			Render("[ Submit ]")

	blurredButton = f("[ %s ]", blurredStyle.Render("Submit"))
)

var (
	itemStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			MarginLeft(1).
			MarginBottom(3)
	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("170"))
)
