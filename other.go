package main 
import (
	"time"
  "os"
  "fmt"
  "bufio"
  "strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
  "github.com/gen2brain/beeep"
)

var (
	name                string
	altscreen           bool
	winHeight, winWidth int
	quitKeys            = key.NewBinding(key.WithKeys("esc", "q"))
	intKeys             = key.NewBinding(key.WithKeys("ctrl+c"))
  skipKey             = key.NewBinding(key.WithKeys("s"))
	boldStyle           = lipgloss.NewStyle().Bold(true)
	italicStyle         = lipgloss.NewStyle().Italic(true)
)
const (
	padding  = 2
	maxWidth = 80
)

type model struct {
	name         string
	altscreen    bool
	duration     time.Duration
	start        time.Time
	timer        timer.Model
	progress     progress.Model
	quitting     bool
	interrupting bool
}

func (m model) Init() tea.Cmd {
	return m.timer.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmds []tea.Cmd
		var cmd tea.Cmd

		step := 100.0 / (m.duration).Seconds()
		cmds = append(cmds, m.progress.IncrPercent(step/100.0))

		m.timer, cmd = m.timer.Update(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		winHeight, winWidth = msg.Height, msg.Width
		if !m.altscreen && m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case timer.StartStopMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case timer.TimeoutMsg:
		m.quitting = true
		return m, tea.Quit

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	case tea.KeyMsg:
		if key.Matches(msg, quitKeys) {
			m.quitting = true
      os.Exit(1)
			// return m, tea.Quit
		}
		if key.Matches(msg, intKeys) {
			m.interrupting = true
      os.Exit(1)
			// return m, tea.Quit
		}
    if key.Matches(msg, skipKey) {
			m.interrupting = true
      return m, tea.Quit
    }
}

	return m, nil
}

func (m model) View() string {
	if m.quitting || m.interrupting {
		return ""
	}

	result := boldStyle.Render(m.start.Format(time.Kitchen))
	if m.name != "" {
		result += ": " + italicStyle.Render(m.name)
	}
	result += " - " + boldStyle.Render(m.timer.View()) + "\n" + m.progress.View()
	if m.altscreen {
		textWidth, textHeight := lipgloss.Size(result)
		return lipgloss.NewStyle().Margin((winHeight-textHeight)/2, (winWidth-textWidth)/2).Render(result)
	}
	return result
}

func notification(title string, body string) {
  beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
  beeep.Alert(title, body, "assets/warning.png")
}

func timerA(duration_str string, name string){
  duration, _:= time.ParseDuration(duration_str)
        var opts []tea.ProgramOption
        if altscreen {
          opts = append(opts, tea.WithAltScreen())
        }
        tea.NewProgram(model{
          duration:  duration,
          timer:     timer.NewWithInterval(duration, time.Second),
          progress:  progress.New(progress.WithDefaultGradient()),
          name:      name,
          altscreen: altscreen,
          start:     time.Now(),
        }, opts...).Run()
}

func interlude(confirm_new bool) {
  reader := bufio.NewReader(os.Stdin)
  if confirm_new {
    for {
      fmt.Println("Press Enter to Continue or type 'q' to quit")
      text, _ := reader.ReadString('\n')
      if (strings.ToLower(strings.TrimRight(text, "\n")) == "q"){
        os.Exit(0)
      } else if (string(text[0]) == "\n") {
        return
      }
    }

  }
  return
}


