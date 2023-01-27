package main

import (
  "os"
  "fmt"
	"strconv"
  // "log"

  "github.com/pelletier/go-toml"
  // "github.com/BurntSushi/toml"
	// "github.com/charmbracelet/bubbles/key"
	// "github.com/charmbracelet/lipgloss"
  "github.com/spf13/cobra"

)

var (
  label       string             
  duration    string       
  cycles      = "" 
)


func init() {
  rootCmd.Flags().StringVarP(&cycles, "cycles", "c", "", "Number of Cycles")

  config, err := toml.LoadFile("config.toml")
  if err != nil {
    os.Exit(1)
  }
  cycle_lengths := config.Get("config.cycles").([]interface{})
  cycle_names := config.Get("config.cycle_names:")([]interface{})
}

var rootCmd = &cobra.Command{
  Args:         cobra.ExactArgs(0),
  RunE: func(cmd *cobra.Command, args []string) error {
    if (cycles != "") {
      cycle_num, err := strconv.ParseInt(cycles, 0, 64 )
      if (err != nil) {
        os.Exit(0)
      }
      for i := 0; i < int(cycle_num); i++ {
        for j := 0; j < cycle_lengths.len(); j++ {
          label = cycle_names[j]
          duration = cycle_lengths[j]
          timerA(duration, label)
          notification(label+" ended", "")
          interlude()
        }
      }
    } else {
        for {
          for j := 0; j < cycle_lengths.len(); j++ {
            label = cycle_names[j]
            duration = cycle_lengths[j]
            timerA(duration, label)
            notification(label+" ended", "")
            interlude()
          }
        }
      }
    }
    return nil
  },
}

func main() {
  rootCmd.Execute() 
  }
