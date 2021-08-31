package settings

import (
	"fmt"
	"github.com/TicketsBot/common/permission"
	"github.com/TicketsBot/worker/bot/command"
	"github.com/TicketsBot/worker/bot/command/registry"
	"github.com/TicketsBot/worker/bot/dbclient"
	"github.com/TicketsBot/worker/bot/utils"
	"github.com/TicketsBot/worker/i18n"
	"github.com/rxdn/gdl/objects/channel/embed"
	"github.com/rxdn/gdl/objects/interaction"
	"github.com/schollz/progressbar/v3"
	"io/ioutil"
	"strings"
)

type LanguageCommand struct {
}

func (LanguageCommand) Properties() registry.Properties {
	return registry.Properties{
		Name:            "language",
		Description:     i18n.HelpLanguage,
		PermissionLevel: permission.Admin,
		Category:        command.Settings,
		Arguments: command.Arguments(
			command.NewRequiredArgument("language", "The country-code of the language to switch to", interaction.OptionTypeString, i18n.MessageLanguageInvalidLanguage),
		),
	}
}

func (c LanguageCommand) GetExecutor() interface{} {
	return c.Execute
}

// TODO: Show options properly
func (c LanguageCommand) Execute(ctx registry.CommandContext, newLanguage string) {
	var valid bool
	var newFlag string
	for language, flag := range i18n.Flags {
		if newLanguage == string(language) || newLanguage == flag {
			if err := dbclient.Client.ActiveLanguage.Set(ctx.GuildId(), language.String()); err != nil { // TODO: Don't wrap
				ctx.HandleError(err)
				return
			}

			newFlag = flag
			valid = true
			break
		}
	}

	if !valid {
		c.sendInvalidMessage(ctx)
		return
	}

	ctx.ReplyRaw(utils.Green, "Language", fmt.Sprintf("Server language has been changed to %s", newFlag))
}

func (LanguageCommand) sendInvalidMessage(ctx registry.CommandContext) {
	example := embed.EmbedField{
		Name:   "Example",
		Value:  fmt.Sprintf("`%slanguage en`\n`%slanguage fr`\n`%slanguage de`", utils.DEFAULT_PREFIX, utils.DEFAULT_PREFIX, utils.DEFAULT_PREFIX),
		Inline: false,
	}

	var list string
	for _, language := range i18n.LanguagesAlphabetical {
		coverage := i18n.GetCoverage(language)
		if coverage == 0 {
			continue
		}

		flag := i18n.Flags[language]

		bar := progressbar.NewOptions(100,
			progressbar.OptionSetWriter(ioutil.Discard),
			progressbar.OptionSetWidth(15),
			progressbar.OptionSetPredictTime(false),
			progressbar.OptionSetRenderBlankState(true),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "=",
				SaucerHead:    ">",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}),
		)
		_ = bar.Set(coverage)

		list += fmt.Sprintf("%s %s `%s`\n", flag, language, strings.TrimSpace(bar.String()))
	}

	list = strings.TrimSuffix(list, "\n")

	ctx.ReplyWithFields(utils.Red, "Error", i18n.MessageLanguageInvalidLanguage, utils.FieldsToSlice(example), list)
	ctx.Accept()
}
