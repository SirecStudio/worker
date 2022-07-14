package manager

import (
	"github.com/TicketsBot/worker/bot/command/impl/admin"
	"github.com/TicketsBot/worker/bot/command/impl/general"
	"github.com/TicketsBot/worker/bot/command/impl/settings"
	"github.com/TicketsBot/worker/bot/command/impl/settings/setup"
	"github.com/TicketsBot/worker/bot/command/impl/statistics"
	"github.com/TicketsBot/worker/bot/command/impl/tags"
	"github.com/TicketsBot/worker/bot/command/impl/tickets"
	"github.com/TicketsBot/worker/bot/command/registry"
)

type CommandManager struct {
	registry registry.Registry
}

func (cm *CommandManager) GetCommands() map[string]registry.Command {
	return cm.registry
}

func (cm *CommandManager) RegisterCommands() {
	cm.registry = make(map[string]registry.Command)

	cm.registry["help"] = general.HelpCommand{Registry: cm.registry}

	cm.registry["admin"] = admin.AdminCommand{}
	cm.registry["registercommands"] = admin.RegisterCommandsCommand{Registry: cm.registry}

	cm.registry["about"] = general.AboutCommand{}
	cm.registry["invite"] = general.InviteCommand{}
	cm.registry["jumptotop"] = general.JumpToTopCommand{}
	cm.registry["vote"] = general.VoteCommand{}

	cm.registry["addadmin"] = settings.AddAdminCommand{}
	cm.registry["addsupport"] = settings.AddSupportCommand{}
	cm.registry["autoclose"] = settings.AutoCloseCommand{}
	cm.registry["blacklist"] = settings.BlacklistCommand{}
	cm.registry["language"] = settings.LanguageCommand{}
	cm.registry["panel"] = settings.PanelCommand{}
	cm.registry["premium"] = settings.PremiumCommand{}
	cm.registry["removeadmin"] = settings.RemoveAdminCommand{}
	cm.registry["removesupport"] = settings.RemoveSupportCommand{}
	cm.registry["premium"] = settings.PremiumCommand{}
	cm.registry["setup"] = setup.SetupCommand{}
	cm.registry["viewstaff"] = settings.ViewStaffCommand{}

	//cm.registry["sync"] = settings.SyncCommand{}
	cm.registry["stats"] = statistics.StatsCommand{}

	cm.registry["managetags"] = tags.ManageTagsCommand{}
	cm.registry["tag"] = tags.TagCommand{}

	cm.registry["add"] = tickets.AddCommand{}
	cm.registry["claim"] = tickets.ClaimCommand{}
	cm.registry["close"] = tickets.CloseCommand{}
	cm.registry["closerequest"] = tickets.CloseRequestCommand{}
	cm.registry["open"] = tickets.OpenCommand{}
	cm.registry["Start Ticket"] = tickets.StartTicketCommand{}
	cm.registry["remove"] = tickets.RemoveCommand{}
	cm.registry["rename"] = tickets.RenameCommand{}
	cm.registry["switchpanel"] = tickets.SwitchPanelCommand{}
	cm.registry["transfer"] = tickets.TransferCommand{}
	cm.registry["unclaim"] = tickets.UnclaimCommand{}
}
