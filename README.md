# trash-reminder

A golang program that sends a Discord message to a  user the day before collection.

## Installation

```
go install github.com/kklopfenstein/trash-reminder@latest
```

## Usage

```
trash-reminder --place <RECOLLECT PLACE> --service <RECOLLECT SERVICE> --discordUserId <DISCORD USER ID> --discordToken <DISCORD TOKEN>
```

This program is best used as a daily cron job.

### Note on Recollect API

I found the `RECOLLECT PLACE` and `RECOLLECT SERVICE` by looking at my city's calendar URL's.
[This section](https://github.com/bachya/aiorecollect#place-and-service-ids) provides some more information on how you might find this.

See the [Discord Developer Portal](https://discord.com/developers/docs/intro) for more information on registering a bot.
Your Discord user ID should be viewable via [developer mode](https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID-) in Discord.