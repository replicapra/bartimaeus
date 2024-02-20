package add

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/replicapra/bartimaeus/internal/constants"
	"github.com/replicapra/bartimaeus/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// bookCmd represents the book command
var BookCmd = &cobra.Command{
	Use:   "book",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, relPath := range args {
			addRelPathToBookSpace(relPath)
		}
	},
}

func init() {
	BookCmd.Flags().BoolP("rename", "r", false, "Rename the book before adding it to the collection")
	viper.BindPFlag("rename", BookCmd.Flags().Lookup("rename"))
}

func addRelPathToBookSpace(relPath string) {

	absPath, err := filepath.Abs(relPath)
	util.CheckErr(err)

	if _, err := os.Stat(absPath); err != nil {
		util.CheckErr(err)
	}

	fileName := strings.TrimSuffix(filepath.Base(absPath), filepath.Ext(absPath))

	if viper.GetBool("rename") {
		p := tea.NewProgram(initialModel(fileName))
		m, err := p.Run()
		if err != nil {
			log.Fatal(err)
		}

		if m, ok := m.(model); ok {
			fileName = m.textInput.Value()
		}
	}

	obsidianDirPath := viper.GetString("obsidian.path")
	newBookDir := path.Join(obsidianDirPath, constants.BookSpacePath, fileName)
	newBookPath := path.Join(newBookDir, fileName+".pdf")
	newMdPath := path.Join(newBookDir, fileName+".md")

	if _, err := os.Stat(newBookPath); err == nil {
		log.Errorf("Book %s already exists", fileName)
		return
	}

	os.Mkdir(newBookDir, os.ModePerm)

	source, err := os.Open(absPath)
	util.CheckErr(err)
	defer source.Close()

	destination, err := os.Create(newBookPath)
	util.CheckErr(err)
	defer destination.Close()

	_, err = os.Create(newMdPath)
	util.CheckErr(err)

	_, err = io.Copy(destination, source)
	util.CheckErr(err)

	log.Infof("Book %s added", fileName)

}

type model struct {
	textInput textinput.Model
	err       error
}

func initialModel(name string) model {
	ti := textinput.New()
	ti.SetValue(name)
	ti.Focus()
	ti.CharLimit = 255
	ti.Width = 80

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.textInput.View() + "\n"
}
