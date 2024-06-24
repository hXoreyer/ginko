package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type step int

const (
	stepProjectName step = iota
	stepCreateDirs
	stepCreateFiles
	stepFinished
	stepError
)

type model struct {
	textInput  textinput.Model
	spinner    spinner.Model
	quitting   bool
	finalValue string
	current    step
	err        error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "gin-project"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{
		textInput: ti,
		spinner:   sp,
		current:   stepProjectName,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			if m.current == stepProjectName {
				if m.textInput.Value() == "" {
					m.finalValue = m.textInput.Placeholder
				} else {
					m.finalValue = m.textInput.Value()
				}
				m.current = stepCreateDirs
				return m, tea.Batch(m.spinner.Tick, createDirsCmd(m.finalValue))
			} else if m.current == stepFinished || m.current == stepError {
				m.quitting = true
				return m, tea.Quit
			}
		case "q":
			if m.current == stepFinished || m.current == stepError {
				m.quitting = true
				return m, tea.Quit
			}
		}

	case spinner.TickMsg:
		if m.current == stepCreateDirs || m.current == stepCreateFiles {
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}

	case dirsCreatedMsg:
		m.current = stepCreateFiles
		return m, createFilesCmd(m.finalValue)

	case dirsErrorMsg:
		m.err = msg.err
		m.current = stepError
		return m, nil

	case fileCreatedMsg:
		m.current = stepFinished
		return m, nil
	}

	var textCmd tea.Cmd
	m.textInput, textCmd = m.textInput.Update(msg)

	return m, textCmd
}

func (m model) View() string {
	var text string
	switch m.current {
	case stepProjectName:
		text = fmt.Sprintf(
			"🌭 Please enter your project name:\n\n%s\n\nPress Enter to start creating directories.",
			m.textInput.View(),
		)
	case stepCreateDirs:
		text = fmt.Sprintf("%s Creating directories...", m.spinner.View())
	case stepCreateFiles:
		text = fmt.Sprintf("%s Creating files...", m.spinner.View())
	case stepFinished:
		text = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Render(fmt.Sprintf("Project %q created successfully!\n\nswag init and go run .\n\nPress q or esc to quit.", m.finalValue))
	case stepError:
		text = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(fmt.Sprintf("Error creating project: %v\n\nPress q or esc to quit.", m.err))
	}
	return text
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting program: %v\n", err)
		os.Exit(1)
	}
}

type fileCreatedMsg struct{}
type dirsCreatedMsg struct{}
type dirsErrorMsg struct {
	err error
}

func createDirsCmd(projectName string) tea.Cmd {
	return func() tea.Msg {
		err := createDirs(projectName)
		if err != nil {
			return dirsErrorMsg{err}
		}
		return dirsCreatedMsg{}
	}
}

func createFilesCmd(projectName string) tea.Cmd {
	return func() tea.Msg {
		err := createFiles(projectName)
		if err != nil {
			return dirsErrorMsg{err}
		}
		return fileCreatedMsg{}
	}
}

func createFiles(projectName string) error {
	Create(projectName)
	return nil
}

func createDirs(projectName string) error {
	dirs := []string{
		"api", "api/user", "bootstrap", "common", "common/code", "common/request", "config", "global", "internal", "internal/user", "middlewares", "models", "routes", "utils",
	}
	base := "./" + projectName
	err := os.Mkdir(base, os.ModePerm)
	if err != nil {
		return err
	}
	for _, v := range dirs {
		err := os.MkdirAll(filepath.Join(base, v), os.ModePerm)
		if err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond) // Simulate folder creation delay
	}
	return nil
}
