/*
Copyright © 2024 Vercello

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
	"fmt"
	microgen "github.com/Verce11o/microgen/internal/cmd"
	config "github.com/Verce11o/microgen/internal/config"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"strings"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generates a boilerplate application structure",
	Long:  `This command generates a boilerplate application structure in new folder.`,
	Run: func(cmd *cobra.Command, args []string) {
		var moduleName string
		var serviceName string
		var additions []string
		var storages []string

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Value(&moduleName).
					Title("✨ Module name").
					Placeholder("👤github.com/john/myproject").
					Validate(func(s string) error {
						if len(s) < 2 {
							return errors.New("module name too short")
						}
						return nil
					}),
			),
			huh.NewGroup(
				huh.NewInput().
					Value(&serviceName).
					Title("📝 Project name").
					Description("This name will be used to create the project folder").
					Placeholder("📦cool-project").
					Validate(func(s string) error {
						if len(s) < 2 {
							return errors.New("project name too short")
						}
						return nil
					}),
			),
			//huh.NewGroup(
			//	huh.NewSelect[string]().
			//		Title("What framework would you 🩵 to use?").
			//		Options(
			//			huh.NewOption("🍸 Gin", "gin").Selected(true),
			//			huh.NewOption("⚡️ Fiber", "fiber"),
			//			huh.NewOption("💠 Echo", "echo"),
			//			huh.NewOption("🧩 Chi", "chi"),
			//		).
			//		Value(&framework),
			//),

			huh.NewGroup(
				huh.NewMultiSelect[string]().
					Title("Additions").
					Description("⚡️ Choose what to add into microservice").
					Options(
						huh.NewOption("🦫 Jaeger", "jaeger"),
						huh.NewOption("📦 AWS (Minio)", "aws"),
					).
					Value(&additions).
					Limit(1).
					Filterable(true),
			),
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					Title("Storage").
					Description("📦 Choose your storage").
					Options(
						huh.NewOption("🐘 Postgres", "postgres"),
						huh.NewOption("☘️  MongoDB", "mongo"),
					).
					Value(&storages).
					Filterable(true),
			),
		)

		err := form.Run()
		if err != nil {
			panic(err)
		}

		var sb strings.Builder
		keyword := func(s []string) string {
			if len(s) > 1 {
				return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(strings.Join(s, ", "))
			}
			return "-"
		}

		fmt.Fprintf(&sb,
			"🚀 Module name: %s\n\n⚡ Additions: %s\n\n📦Storages: %s\n\nMade by Vercello with <3",
			lipgloss.NewStyle().Bold(true).Render(moduleName),
			keyword(additions),
			keyword(storages),
		)

		fmt.Println(
			lipgloss.NewStyle().
				Width(60).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Margin(1, 2).
				Render(sb.String()),
		)

		confStorages := make([]config.Storage, 0, len(storages))
		for _, storage := range storages {
			confStorages = append(confStorages, config.Storage(storage))
		}

		config := config.NewConfig(config.App{Module: moduleName, Storages: confStorages}, serviceName)
		microgen.InitApp(config)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
