package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/areias03/metagen/api/tui"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func main() {
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for an item in all DBs",
	Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ip := tea.NewProgram(tui.InitialInputModel())
		if _, err := ip.Run(); err != nil {
			log.Fatal(err)
		}
		items := []list.Item{}
		for _, v := range tui.DBs.Databases {
			items = append(items, tui.Item{Name: v.Name, Desc: v.Match})

		}

		m := tui.ListModel{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}
		m.List.Title = "Found Items"

		lp := tea.NewProgram(m, tea.WithAltScreen())

		if _, err := lp.Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
