package requests

// Discord API Endpoints
const (
	EndpointBaseURL = "https://discord.com/api/v9/"
	slash           = "/"

	// slugs
	applications = "applications"
	interactions = "interactions"
	webhooks     = "webhooks"
	guilds       = "guilds"
	channels     = "channels"
	invites      = "invites"
	stickers     = "stickers"
	users        = "users"

	// resources
	commands        = "commands"
	permissions     = "permissions"
	callback        = "callback"
	messages        = "messages"
	auditlogs       = "audit-logs"
	crosspost       = "crosspost"
	bulkdelete      = "bulk-delete"
	reactions       = "reactions"
	emoji           = "emoji"
	emojis          = "emojis"
	followers       = "followers"
	typing          = "typing"
	pins            = "pins"
	recipients      = "recipients"
	threads         = "threads"
	members         = "members"
	threadmembers   = "thread-members"
	active          = "active"
	archived        = "archived"
	public          = "public"
	private         = "private"
	scheduledevents = "scheduled-events"
	templates       = "templates"
	preview         = "preview"
	search          = "search"
	nick            = "nick"
	roles           = "roles"
	bans            = "bans"
	prune           = "prune"
	regions         = "regions"
	integrations    = "integrations"
	widget          = "widget"
	widgetjson      = "widget.json"
	widgetpng       = "widget.png"
	vanityurl       = "vanity-url"
	welcomescreen   = "welcome-screen"
	voicestates     = "voice-states"
	member          = "member"
	slack           = "slack"
	github          = "github"

	// parameters
	original = "@original"
	me       = "@me"
)

// EndpointGetGlobalApplicationCommands builds a query for an HTTP request.
func EndpointGetGlobalApplicationCommands(applicationid string) string {
	return EndpointBaseURL + applications + slash + applicationid + slash + commands
}

// EndpointCreateGlobalApplicationCommand builds a query for an HTTP request.
func EndpointCreateGlobalApplicationCommand(applicationid string) string {
	return EndpointBaseURL + applications + slash + applicationid + slash + commands
}

// EndpointGetGlobalApplicationCommand builds a query for an HTTP request.
func EndpointGetGlobalApplicationCommand(applicationid, commandid string) string {
	return EndpointBaseURL + applications + slash + applicationid + slash + commands + slash + commandid
}

// EndpointEditGlobalApplicationCommand builds a query for an HTTP request.
func EndpointEditGlobalApplicationCommand(applicationid, commandid string) string {
	return EndpointBaseURL + applications + slash + applicationid + slash + commands + slash + commandid
}

// EndpointDeleteGlobalApplicationCommand builds a query for an HTTP request.
func EndpointDeleteGlobalApplicationCommand(applicationid, commandid string) string {
	return EndpointBaseURL + applications + slash + applicationid + slash + commands + slash + commandid
}

// EndpointBulkOverwriteGlobalApplicationCommands builds a query for an HTTP request.
func EndpointBulkOverwriteGlobalApplicationCommands(applicationid string) string {
	return EndpointBaseURL + applications + slash + applicationid + slash + commands
}

// EndpointGetGuildApplicationCommands builds a query for an HTTP request.
func EndpointGetGuildApplicationCommands(applicationid, guildid string) string {
	return EndpointBaseURL + applications + slash + applicationid + slash + guilds + slash + guildid + slash + commands
}

// EndpointCreateGuildApplicationCommand builds a query for an HTTP request.
func EndpointCreateGuildApplicationCommand(applicationid, guildid string) string {
	return EndpointBaseURL + applications + slash + applicationid + slash + guilds + slash + guildid + slash + commands
}

// EndpointGetGuildApplicationCommand builds a query for an HTTP request.
func EndpointGetGuildApplicationCommand(applicationid, guildid, commandid string) string {
	return EndpointBaseURL + applications + slash + applicationid + slash + guilds + slash + guildid + slash + commands + slash + commandid
}

// EndpointEditGuildApplicationCommand builds a query for an HTTP request.
func EndpointEditGuildApplicationCommand(applicationid, guildid, commandid string) string {
	return EndpointBaseURL + applications + slash + applicationid + slash + guilds + slash + guildid + slash + commands + slash + commandid
}

// EndpointDeleteGuildApplicationCommand builds a query for an HTTP request.
func EndpointDeleteGuildApplicationCommand(applicationid, guildid, commandid string) string {
	return EndpointBaseURL + applications + slash + applicationid + slash + guilds + slash + guildid + slash + commands + slash + commandid
}

// EndpointBulkOverwriteGuildApplicationCommands builds a query for an HTTP request.
func EndpointBulkOverwriteGuildApplicationCommands(applicationid, guildid string) string {
	return EndpointBaseURL + applications + slash + applicationid + slash + guilds + slash + guildid + slash + commands
}

// EndpointGetGuildApplicationCommandPermissions builds a query for an HTTP request.
func EndpointGetGuildApplicationCommandPermissions(applicationid, guildid string) string {
	return EndpointBaseURL + applications + slash + applicationid + slash + guilds + slash + guildid + slash + commands + slash + permissions
}

// EndpointGetApplicationCommandPermissions builds a query for an HTTP request.
func EndpointGetApplicationCommandPermissions(applicationid, guildid, commandid string) string {
	return EndpointBaseURL + applications + slash + applicationid + slash + guilds + slash + guildid + slash + commands + slash + commandid + slash + permissions
}

// EndpointEditApplicationCommandPermissions builds a query for an HTTP request.
func EndpointEditApplicationCommandPermissions(applicationid, guildid, commandid string) string {
	return EndpointBaseURL + applications + slash + applicationid + slash + guilds + slash + guildid + slash + commands + slash + commandid + slash + permissions
}

// EndpointBatchEditApplicationCommandPermissions builds a query for an HTTP request.
func EndpointBatchEditApplicationCommandPermissions(applicationid, guildid string) string {
	return EndpointBaseURL + applications + slash + applicationid + slash + guilds + slash + guildid + slash + commands + slash + permissions
}

// EndpointCreateInteractionResponse builds a query for an HTTP request.
func EndpointCreateInteractionResponse(interactionid, interactiontoken string) string {
	return EndpointBaseURL + interactions + slash + interactionid + slash + interactiontoken + slash + callback
}

// EndpointGetOriginalInteractionResponse builds a query for an HTTP request.
func EndpointGetOriginalInteractionResponse(applicationid, interactiontoken string) string {
	return EndpointBaseURL + webhooks + slash + applicationid + slash + interactiontoken + slash + messages + slash + original
}

// EndpointEditOriginalInteractionResponse builds a query for an HTTP request.
func EndpointEditOriginalInteractionResponse(applicationid, interactiontoken string) string {
	return EndpointBaseURL + webhooks + slash + applicationid + slash + interactiontoken + slash + messages + slash + original
}

// EndpointDeleteOriginalInteractionResponse builds a query for an HTTP request.
func EndpointDeleteOriginalInteractionResponse(applicationid, interactiontoken string) string {
	return EndpointBaseURL + webhooks + slash + applicationid + slash + interactiontoken + slash + messages + slash + original
}

// EndpointCreateFollowupMessage builds a query for an HTTP request.
func EndpointCreateFollowupMessage(applicationid, interactiontoken string) string {
	return EndpointBaseURL + webhooks + slash + applicationid + slash + interactiontoken
}

// EndpointGetFollowupMessage builds a query for an HTTP request.
func EndpointGetFollowupMessage(applicationid, interactiontoken, messageid string) string {
	return EndpointBaseURL + webhooks + slash + applicationid + slash + interactiontoken + slash + messages + slash + messageid
}

// EndpointEditFollowupMessage builds a query for an HTTP request.
func EndpointEditFollowupMessage(applicationid, interactiontoken, messageid string) string {
	return EndpointBaseURL + webhooks + slash + applicationid + slash + interactiontoken + slash + messages + slash + messageid
}

// EndpointDeleteFollowupMessage builds a query for an HTTP request.
func EndpointDeleteFollowupMessage(applicationid, interactiontoken, messageid string) string {
	return EndpointBaseURL + webhooks + slash + applicationid + slash + interactiontoken + slash + messages + slash + messageid
}

// EndpointGetGuildAuditLog builds a query for an HTTP request.
func EndpointGetGuildAuditLog(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + auditlogs
}

// EndpointGetChannel builds a query for an HTTP request.
func EndpointGetChannel(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid
}

// EndpointModifyChannel builds a query for an HTTP request.
func EndpointModifyChannel(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid
}

// EndpointDeleteCloseChannel builds a query for an HTTP request.
func EndpointDeleteCloseChannel(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid
}

// EndpointGetChannelMessages builds a query for an HTTP request.
func EndpointGetChannelMessages(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + messages
}

// EndpointGetChannelMessage builds a query for an HTTP request.
func EndpointGetChannelMessage(channelid, messageid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + messages + slash + messageid
}

// EndpointCreateMessage builds a query for an HTTP request.
func EndpointCreateMessage(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + messages
}

// EndpointCrosspostMessage builds a query for an HTTP request.
func EndpointCrosspostMessage(channelid, messageid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + messages + slash + messageid + slash + crosspost
}

// EndpointCreateReaction builds a query for an HTTP request.
func EndpointCreateReaction(channelid, messageid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + messages + slash + messageid + slash + reactions + slash + emoji + slash + me
}

// EndpointDeleteOwnReaction builds a query for an HTTP request.
func EndpointDeleteOwnReaction(channelid, messageid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + messages + slash + messageid + slash + reactions + slash + emoji + slash + me
}

// EndpointDeleteUserReaction builds a query for an HTTP request.
func EndpointDeleteUserReaction(channelid, messageid, userid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + messages + slash + messageid + slash + reactions + slash + emoji + slash + userid
}

// EndpointGetReactions builds a query for an HTTP request.
func EndpointGetReactions(channelid, messageid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + messages + slash + messageid + slash + reactions + slash + emoji
}

// EndpointDeleteAllReactions builds a query for an HTTP request.
func EndpointDeleteAllReactions(channelid, messageid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + messages + slash + messageid + slash + reactions
}

// EndpointDeleteAllReactionsforEmoji builds a query for an HTTP request.
func EndpointDeleteAllReactionsforEmoji(channelid, messageid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + messages + slash + messageid + slash + reactions + slash + emoji
}

// EndpointEditMessage builds a query for an HTTP request.
func EndpointEditMessage(channelid, messageid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + messages + slash + messageid
}

// EndpointDeleteMessage builds a query for an HTTP request.
func EndpointDeleteMessage(channelid, messageid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + messages + slash + messageid
}

// EndpointBulkDeleteMessages builds a query for an HTTP request.
func EndpointBulkDeleteMessages(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + messages + slash + bulkdelete
}

// EndpointEditChannelPermissions builds a query for an HTTP request.
func EndpointEditChannelPermissions(channelid, overwriteid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + permissions + slash + overwriteid
}

// EndpointGetChannelInvites builds a query for an HTTP request.
func EndpointGetChannelInvites(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + invites
}

// EndpointCreateChannelInvite builds a query for an HTTP request.
func EndpointCreateChannelInvite(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + invites
}

// EndpointDeleteChannelPermission builds a query for an HTTP request.
func EndpointDeleteChannelPermission(channelid, overwriteid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + permissions + slash + overwriteid
}

// EndpointFollowNewsChannel builds a query for an HTTP request.
func EndpointFollowNewsChannel(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + followers
}

// EndpointTriggerTypingIndicator builds a query for an HTTP request.
func EndpointTriggerTypingIndicator(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + typing
}

// EndpointGetPinnedMessages builds a query for an HTTP request.
func EndpointGetPinnedMessages(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + pins
}

// EndpointPinMessage builds a query for an HTTP request.
func EndpointPinMessage(channelid, messageid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + pins + slash + messageid
}

// EndpointUnpinMessage builds a query for an HTTP request.
func EndpointUnpinMessage(channelid, messageid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + pins + slash + messageid
}

// EndpointGroupDMAddRecipient builds a query for an HTTP request.
func EndpointGroupDMAddRecipient(channelid, userid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + recipients + slash + userid
}

// EndpointGroupDMRemoveRecipient builds a query for an HTTP request.
func EndpointGroupDMRemoveRecipient(channelid, userid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + recipients + slash + userid
}

// EndpointStartThreadfromMessage builds a query for an HTTP request.
func EndpointStartThreadfromMessage(channelid, messageid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + messages + slash + messageid + slash + threads
}

// EndpointStartThreadwithoutMessage builds a query for an HTTP request.
func EndpointStartThreadwithoutMessage(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + threads
}

// EndpointStartThreadinForumChannel builds a query for an HTTP request.
func EndpointStartThreadinForumChannel(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + threads
}

// EndpointJoinThread builds a query for an HTTP request.
func EndpointJoinThread(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + threadmembers + slash + me
}

// EndpointAddThreadMember builds a query for an HTTP request.
func EndpointAddThreadMember(channelid, userid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + threadmembers + slash + userid
}

// EndpointLeaveThread builds a query for an HTTP request.
func EndpointLeaveThread(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + threadmembers + slash + me
}

// EndpointRemoveThreadMember builds a query for an HTTP request.
func EndpointRemoveThreadMember(channelid, userid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + threadmembers + slash + userid
}

// EndpointGetThreadMember builds a query for an HTTP request.
func EndpointGetThreadMember(channelid, userid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + threadmembers + slash + userid
}

// EndpointListThreadMembers builds a query for an HTTP request.
func EndpointListThreadMembers(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + threadmembers
}

// EndpointListActiveChannelThreads builds a query for an HTTP request.
func EndpointListActiveChannelThreads(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + threads + slash + active
}

// EndpointListPublicArchivedThreads builds a query for an HTTP request.
func EndpointListPublicArchivedThreads(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + threads + slash + archived + slash + public
}

// EndpointListPrivateArchivedThreads builds a query for an HTTP request.
func EndpointListPrivateArchivedThreads(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + threads + slash + archived + slash + private
}

// EndpointListJoinedPrivateArchivedThreads builds a query for an HTTP request.
func EndpointListJoinedPrivateArchivedThreads(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + users + slash + me + slash + threads + slash + archived + slash + private
}

// EndpointListGuildEmojis builds a query for an HTTP request.
func EndpointListGuildEmojis(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + emojis
}

// EndpointGetGuildEmoji builds a query for an HTTP request.
func EndpointGetGuildEmoji(guildid, emojiid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + emojis + slash + emojiid
}

// EndpointCreateGuildEmoji builds a query for an HTTP request.
func EndpointCreateGuildEmoji(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + emojis
}

// EndpointModifyGuildEmoji builds a query for an HTTP request.
func EndpointModifyGuildEmoji(guildid, emojiid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + emojis + slash + emojiid
}

// EndpointDeleteGuildEmoji builds a query for an HTTP request.
func EndpointDeleteGuildEmoji(guildid, emojiid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + emojis + slash + emojiid
}

// EndpointListScheduledEventsforGuild builds a query for an HTTP request.
func EndpointListScheduledEventsforGuild(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + scheduledevents
}

// EndpointCreateGuildScheduledEvent builds a query for an HTTP request.
func EndpointCreateGuildScheduledEvent(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + scheduledevents
}

// EndpointGetGuildScheduledEvent builds a query for an HTTP request.
func EndpointGetGuildScheduledEvent(guildid, guild_scheduled_eventid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + scheduledevents + slash + guild_scheduled_eventid
}

// EndpointModifyGuildScheduledEvent builds a query for an HTTP request.
func EndpointModifyGuildScheduledEvent(guildid, guild_scheduled_eventid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + scheduledevents + slash + guild_scheduled_eventid
}

// EndpointDeleteGuildScheduledEvent builds a query for an HTTP request.
func EndpointDeleteGuildScheduledEvent(guildid, guild_scheduled_eventid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + scheduledevents + slash + guild_scheduled_eventid
}

// EndpointGetGuildScheduledEventUsers builds a query for an HTTP request.
func EndpointGetGuildScheduledEventUsers(guildid, guild_scheduled_eventid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + scheduledevents + slash + guild_scheduled_eventid + slash + users
}

// EndpointGetGuildTemplate builds a query for an HTTP request.
func EndpointGetGuildTemplate(templatecode string) string {
	return EndpointBaseURL + guilds + slash + templates + slash + templatecode
}

// EndpointCreateGuildfromGuildTemplate builds a query for an HTTP request.
func EndpointCreateGuildfromGuildTemplate(templatecode string) string {
	return EndpointBaseURL + guilds + slash + templates + slash + templatecode
}

// EndpointGetGuildTemplates builds a query for an HTTP request.
func EndpointGetGuildTemplates(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + templates
}

// EndpointCreateGuildTemplate builds a query for an HTTP request.
func EndpointCreateGuildTemplate(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + templates
}

// EndpointSyncGuildTemplate builds a query for an HTTP request.
func EndpointSyncGuildTemplate(guildid, templatecode string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + templates + slash + templatecode
}

// EndpointModifyGuildTemplate builds a query for an HTTP request.
func EndpointModifyGuildTemplate(guildid, templatecode string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + templates + slash + templatecode
}

// EndpointDeleteGuildTemplate builds a query for an HTTP request.
func EndpointDeleteGuildTemplate(guildid, templatecode string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + templates + slash + templatecode
}

// EndpointGetGuild builds a query for an HTTP request.
func EndpointGetGuild(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid
}

// EndpointGetGuildPreview builds a query for an HTTP request.
func EndpointGetGuildPreview(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + preview
}

// EndpointModifyGuild builds a query for an HTTP request.
func EndpointModifyGuild(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid
}

// EndpointDeleteGuild builds a query for an HTTP request.
func EndpointDeleteGuild(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid
}

// EndpointGetGuildChannels builds a query for an HTTP request.
func EndpointGetGuildChannels(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + channels
}

// EndpointCreateGuildChannel builds a query for an HTTP request.
func EndpointCreateGuildChannel(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + channels
}

// EndpointModifyGuildChannelPositions builds a query for an HTTP request.
func EndpointModifyGuildChannelPositions(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + channels
}

// EndpointListActiveGuildThreads builds a query for an HTTP request.
func EndpointListActiveGuildThreads(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + threads + slash + active
}

// EndpointGetGuildMember builds a query for an HTTP request.
func EndpointGetGuildMember(guildid, userid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + members + slash + userid
}

// EndpointListGuildMembers builds a query for an HTTP request.
func EndpointListGuildMembers(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + members
}

// EndpointSearchGuildMembers builds a query for an HTTP request.
func EndpointSearchGuildMembers(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + members + slash + search
}

// EndpointAddGuildMember builds a query for an HTTP request.
func EndpointAddGuildMember(guildid, userid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + members + slash + userid
}

// EndpointModifyGuildMember builds a query for an HTTP request.
func EndpointModifyGuildMember(guildid, userid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + members + slash + userid
}

// EndpointModifyCurrentMember builds a query for an HTTP request.
func EndpointModifyCurrentMember(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + members + slash + me
}

// EndpointModifyCurrentUserNick builds a query for an HTTP request.
func EndpointModifyCurrentUserNick(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + members + slash + me + slash + nick
}

// EndpointAddGuildMemberRole builds a query for an HTTP request.
func EndpointAddGuildMemberRole(guildid, userid, roleid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + members + slash + userid + slash + roles + slash + roleid
}

// EndpointRemoveGuildMemberRole builds a query for an HTTP request.
func EndpointRemoveGuildMemberRole(guildid, userid, roleid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + members + slash + userid + slash + roles + slash + roleid
}

// EndpointRemoveGuildMember builds a query for an HTTP request.
func EndpointRemoveGuildMember(guildid, userid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + members + slash + userid
}

// EndpointGetGuildBans builds a query for an HTTP request.
func EndpointGetGuildBans(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + bans
}

// EndpointGetGuildBan builds a query for an HTTP request.
func EndpointGetGuildBan(guildid, userid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + bans + slash + userid
}

// EndpointCreateGuildBan builds a query for an HTTP request.
func EndpointCreateGuildBan(guildid, userid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + bans + slash + userid
}

// EndpointRemoveGuildBan builds a query for an HTTP request.
func EndpointRemoveGuildBan(guildid, userid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + bans + slash + userid
}

// EndpointGetGuildRoles builds a query for an HTTP request.
func EndpointGetGuildRoles(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + roles
}

// EndpointCreateGuildRole builds a query for an HTTP request.
func EndpointCreateGuildRole(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + roles
}

// EndpointModifyGuildRolePositions builds a query for an HTTP request.
func EndpointModifyGuildRolePositions(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + roles
}

// EndpointModifyGuildRole builds a query for an HTTP request.
func EndpointModifyGuildRole(guildid, roleid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + roles + slash + roleid
}

// EndpointDeleteGuildRole builds a query for an HTTP request.
func EndpointDeleteGuildRole(guildid, roleid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + roles + slash + roleid
}

// EndpointGetGuildPruneCount builds a query for an HTTP request.
func EndpointGetGuildPruneCount(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + prune
}

// EndpointBeginGuildPrune builds a query for an HTTP request.
func EndpointBeginGuildPrune(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + prune
}

// EndpointGetGuildVoiceRegions builds a query for an HTTP request.
func EndpointGetGuildVoiceRegions(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + regions
}

// EndpointGetGuildInvites builds a query for an HTTP request.
func EndpointGetGuildInvites(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + invites
}

// EndpointGetGuildIntegrations builds a query for an HTTP request.
func EndpointGetGuildIntegrations(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + integrations
}

// EndpointDeleteGuildIntegration builds a query for an HTTP request.
func EndpointDeleteGuildIntegration(guildid, integrationid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + integrations + slash + integrationid
}

// EndpointGetGuildWidgetSettings builds a query for an HTTP request.
func EndpointGetGuildWidgetSettings(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + widget
}

// EndpointModifyGuildWidget builds a query for an HTTP request.
func EndpointModifyGuildWidget(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + widget
}

// EndpointGetGuildWidget builds a query for an HTTP request.
func EndpointGetGuildWidget(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + widgetjson
}

// EndpointGetGuildVanityURL builds a query for an HTTP request.
func EndpointGetGuildVanityURL(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + vanityurl
}

// EndpointGetGuildWidgetImage builds a query for an HTTP request.
func EndpointGetGuildWidgetImage(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + widgetpng
}

// EndpointGetGuildWelcomeScreen builds a query for an HTTP request.
func EndpointGetGuildWelcomeScreen(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + welcomescreen
}

// EndpointModifyGuildWelcomeScreen builds a query for an HTTP request.
func EndpointModifyGuildWelcomeScreen(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + welcomescreen
}

// EndpointModifyCurrentUserVoiceState builds a query for an HTTP request.
func EndpointModifyCurrentUserVoiceState(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + voicestates + slash + me
}

// EndpointModifyUserVoiceState builds a query for an HTTP request.
func EndpointModifyUserVoiceState(guildid, userid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + voicestates + slash + userid
}

// EndpointGetInvite builds a query for an HTTP request.
func EndpointGetInvite(invitecode string) string {
	return EndpointBaseURL + invites + slash + invitecode
}

// EndpointDeleteInvite builds a query for an HTTP request.
func EndpointDeleteInvite(invitecode string) string {
	return EndpointBaseURL + invites + slash + invitecode
}

// EndpointGetSticker builds a query for an HTTP request.
func EndpointGetSticker(stickerid string) string {
	return EndpointBaseURL + stickers + slash + stickerid
}

// EndpointListGuildStickers builds a query for an HTTP request.
func EndpointListGuildStickers(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + stickers
}

// EndpointGetGuildSticker builds a query for an HTTP request.
func EndpointGetGuildSticker(guildid, stickerid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + stickers + slash + stickerid
}

// EndpointCreateGuildSticker builds a query for an HTTP request.
func EndpointCreateGuildSticker(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + stickers
}

// EndpointModifyGuildSticker builds a query for an HTTP request.
func EndpointModifyGuildSticker(guildid, stickerid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + stickers + slash + stickerid
}

// EndpointDeleteGuildSticker builds a query for an HTTP request.
func EndpointDeleteGuildSticker(guildid, stickerid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + stickers + slash + stickerid
}

// EndpointGetUser builds a query for an HTTP request.
func EndpointGetUser(userid string) string {
	return EndpointBaseURL + users + slash + userid
}

// EndpointGetCurrentUserGuildMember builds a query for an HTTP request.
func EndpointGetCurrentUserGuildMember(guildid string) string {
	return EndpointBaseURL + users + slash + me + slash + guilds + slash + guildid + slash + member
}

// EndpointLeaveGuild builds a query for an HTTP request.
func EndpointLeaveGuild(guildid string) string {
	return EndpointBaseURL + users + slash + me + slash + guilds + slash + guildid
}

// EndpointCreateWebhook builds a query for an HTTP request.
func EndpointCreateWebhook(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + webhooks
}

// EndpointGetChannelWebhooks builds a query for an HTTP request.
func EndpointGetChannelWebhooks(channelid string) string {
	return EndpointBaseURL + channels + slash + channelid + slash + webhooks
}

// EndpointGetGuildWebhooks builds a query for an HTTP request.
func EndpointGetGuildWebhooks(guildid string) string {
	return EndpointBaseURL + guilds + slash + guildid + slash + webhooks
}

// EndpointGetWebhook builds a query for an HTTP request.
func EndpointGetWebhook(webhookid string) string {
	return EndpointBaseURL + webhooks + slash + webhookid
}

// EndpointGetWebhookwithToken builds a query for an HTTP request.
func EndpointGetWebhookwithToken(webhookid, webhooktoken string) string {
	return EndpointBaseURL + webhooks + slash + webhookid + slash + webhooktoken
}

// EndpointModifyWebhook builds a query for an HTTP request.
func EndpointModifyWebhook(webhookid string) string {
	return EndpointBaseURL + webhooks + slash + webhookid
}

// EndpointModifyWebhookwithToken builds a query for an HTTP request.
func EndpointModifyWebhookwithToken(webhookid, webhooktoken string) string {
	return EndpointBaseURL + webhooks + slash + webhookid + slash + webhooktoken
}

// EndpointDeleteWebhook builds a query for an HTTP request.
func EndpointDeleteWebhook(webhookid string) string {
	return EndpointBaseURL + webhooks + slash + webhookid
}

// EndpointDeleteWebhookwithToken builds a query for an HTTP request.
func EndpointDeleteWebhookwithToken(webhookid, webhooktoken string) string {
	return EndpointBaseURL + webhooks + slash + webhookid + slash + webhooktoken
}

// EndpointExecuteWebhook builds a query for an HTTP request.
func EndpointExecuteWebhook(webhookid, webhooktoken string) string {
	return EndpointBaseURL + webhooks + slash + webhookid + slash + webhooktoken
}

// EndpointExecuteSlackCompatibleWebhook builds a query for an HTTP request.
func EndpointExecuteSlackCompatibleWebhook(webhookid, webhooktoken string) string {
	return EndpointBaseURL + webhooks + slash + webhookid + slash + webhooktoken + slash + slack
}

// EndpointExecuteGitHubCompatibleWebhook builds a query for an HTTP request.
func EndpointExecuteGitHubCompatibleWebhook(webhookid, webhooktoken string) string {
	return EndpointBaseURL + webhooks + slash + webhookid + slash + webhooktoken + slash + github
}

// EndpointGetWebhookMessage builds a query for an HTTP request.
func EndpointGetWebhookMessage(webhookid, webhooktoken, messageid string) string {
	return EndpointBaseURL + webhooks + slash + webhookid + slash + webhooktoken + slash + messages + slash + messageid
}

// EndpointEditWebhookMessage builds a query for an HTTP request.
func EndpointEditWebhookMessage(webhookid, webhooktoken, messageid string) string {
	return EndpointBaseURL + webhooks + slash + webhookid + slash + webhooktoken + slash + messages + slash + messageid
}

// EndpointDeleteWebhookMessage builds a query for an HTTP request.
func EndpointDeleteWebhookMessage(webhookid, webhooktoken, messageid string) string {
	return EndpointBaseURL + webhooks + slash + webhookid + slash + webhooktoken + slash + messages + slash + messageid
}