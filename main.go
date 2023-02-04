package main

import (
  "os"
  "fmt"
	"strconv"

  "github.com/spf13/cobra"
	mcobra "github.com/muesli/mango-cobra"
  "github.com/spf13/viper"
  "github.com/muesli/roff"


)

var (
  label       string             
  duration    string       
  cycles      = "" 
  cycle_ctr   = 0
  source      ="."
)

var manCmd = &cobra.Command{
	Use:                   "man",
	Short:                 "Generates man pages",
	SilenceUsage:          true,
	DisableFlagsInUseLine: true,
	Hidden:                true,
	Args:                  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		manPage, err := mcobra.NewManPage(1, rootCmd)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, manPage.Build(roff.NewDocument()))
		return err
	},
}

func getConf(source string) ([]string, []string, bool) {
  viper.SetConfigName("pomodoro-config") // name of config file (without extension)
  viper.AddConfigPath(source)
  viper.SetConfigType("yaml")
  err := viper.ReadInConfig() // Find and read the config file
  if err != nil { // Handle errors reading the config file
    panic(fmt.Errorf("fatal error config file: %w", err))
  }
  cycle_lengths := viper.GetStringSlice("cycles_lengths")
  cycle_names := viper.GetStringSlice("cycle_names")
  confirm_new := viper.GetBool("confirm_new")

  return cycle_lengths, cycle_names, confirm_new
}

var rootCmd = &cobra.Command{
  Args:         cobra.ExactArgs(0),
  RunE: func(cmd *cobra.Command, args []string) error {
    if (source == "") {
      source = "~/"
    }
    cycle_lengths, cycle_names, confirm_new := getConf(source)
    if (cycles != "") {
      cycle_num, err := strconv.ParseInt(cycles, 0, 64 )
      if (err != nil) {
        os.Exit(0)
      }
      for i := 0; i < int(cycle_num); i++ {
        for j := 0; j < len(cycle_lengths); j++ {
          label = cycle_names[j]
          duration = cycle_lengths[j]
          timerA(duration, label + " | cycle " + strconv.Itoa(i) + "/" + cycles)
          notification(label+" ended", "")
          interlude(confirm_new)
        }
      }
    } else {
        for {
          for j := 0; j < len(cycle_lengths); j++ {
            label = cycle_names[j]
            duration = cycle_lengths[j]
            timerA(duration, "[ " + label + " cycle " + strconv.Itoa(cycle_ctr) + " ]")
            notification(label+" ended", "")
            interlude(confirm_new)
  cycle_ctr += 1;
  fmt.Print(cycle_ctr)
          }
        }
      }
    return nil
    },
}
func init() {
  rootCmd.Flags().StringVarP(&cycles, "cycles", "c", "", "Number of Cycles")
  rootCmd.Flags().StringVarP(&source, "config path", "s", "", "Path to dir containing config.yaml file")
	rootCmd.AddCommand(manCmd)
}


func main() {
  rootCmd.Execute() 
  }
