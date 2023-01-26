package main

import (
	"time"
  "os"
  "fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
  "github.com/spf13/cobra"

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
			return m, tea.Quit
		}
		if key.Matches(msg, intKeys) {
			m.interrupting = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting || m.interrupting {
		return "\n"
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

var (
	name                string
	altscreen           bool
	winHeight, winWidth int
	quitKeys            = key.NewBinding(key.WithKeys("esc", "q"))
	intKeys             = key.NewBinding(key.WithKeys("ctrl+c"))
	boldStyle           = lipgloss.NewStyle().Bold(true)
	italicStyle         = lipgloss.NewStyle().Italic(true)
)

const (
	padding  = 2
	maxWidth = 80
)

var rootCmd = &cobra.Command{
	// Use:          "timer",
	// SilenceUsage: true,
	// Args:         cobra.ExactArgs(0),
	RunE: func() error {
    duration, _:= time.ParseDuration("3s")
		var opts []tea.ProgramOption
		if altscreen {
			opts = append(opts, tea.WithAltScreen())
		}
		m, err := tea.NewProgram(model{
			duration:  duration,
			timer:     timer.NewWithInterval(duration, time.Second),
			progress:  progress.New(progress.WithDefaultGradient()),
			name:      name,
			altscreen: altscreen,
			start:     time.Now(),
		}, opts...).Run()
		if err != nil {
			return err
		}
		if m.(model).interrupting {
			return fmt.Errorf("interrupted")
		}
		if name != "" {
			cmd.Printf("%s ", name)
		}
		cmd.Printf("finished!\n")
		return nil
	},
}

// func init() {
	// rootCmd.Flags().StringVarP(&name, "name", "n", "", "timer name")
	// rootCmd.Flags().BoolVarP(&altscreen, "fullscreen", "f", false, "fullscreen")

	// rootCmd.AddCommand(manCmd)
// }

  func wrapperidk() {
    
    duration, _:= time.ParseDuration("3s")

		var opts []tea.ProgramOption
		if altscreen {
			opts = append(opts, tea.WithAltScreen())
		}
		m := tea.NewProgram(model{
			duration:  duration,
			timer:     timer.NewWithInterval(duration, time.Second),
			progress:  progress.New(progress.WithDefaultGradient()),
			name:      name,
			altscreen: altscreen,
			start:     time.Now(),
		}, opts...)
    m.Run()
    fmt.Fprint(os.Stdout, "complted")
    m.Run()
		// if name != "" {
		// 	cmd.Printf("%s ", name)
		// }
		// cmd.Printf("finished!\n")
  }
func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
  // wrapperidk()
		// duration, err := time.ParseDuration(args[0])
}
