Plutabot
========

Plutabot is minimal irc-bot running on pluta our chaostreff gateway pc. It
is used for informing people that chaos treff has opened with it joining
our channel.

# Setup
1. Build plutabot with go
2. Copy binary to pluta via scp
3. Start plutabot -> use systemd service file (example located in extras dir) to start at boot

# Options
Add a fifo that if written to will trigger a plutabot posting to the channel.
-fifo <path to fifo>


