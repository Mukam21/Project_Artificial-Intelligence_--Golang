News Feed Bot

  Bot for Telegram that gets and posts news to a channel.

Features

  Fetching articles from RSS feeds
  Article summaries powered by GPT-3.5
  Admin commands for managing sources

Environment variables
  NFB_TELEGRAM_BOT_TOKEN — token for Telegram Bot API
  NFB_TELEGRAM_CHANNEL_ID — ID of the channel to post to, can be obtained via @JsonDumpBot
  NFB_DATABASE_DSN — PostgreSQL connection string
  NFB_FETCH_INTERVAL — the interval of checking for new articles, default 10m
  NFB_NOTIFICATION_INTERVAL — the interval of delivering new articles to Telegram channel, default 1m
  NFB_FILTER_KEYWORDS — comma separated list of words to skip articles containing these words
  NFB_OPENAI_KEY — token for OpenAI API
  NFB_OPENAI_PROMPT — prompt for GPT-3.5 Turbo to generate summary

HCL
  News Feed Bot can be configured with HCL config file. The service is looking for config file in following locations:

  ./config.hcl
  ./config.local.hcl
  $HOME/.config/news-feed-bot/config.hcl
  The names of parameters are the same except that there is no prefix and names are in lower case instead of upper case.