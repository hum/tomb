# tomb

🤖 Turn your Discord server into a CSV dataset. Requires for your bot to be in the server.

## Usage

You can run this as a CLI tool. Note: `OUTPUT_PATH` should be a folder destination, not a file.

```bash
> tomb -t [BOT_TOKEN] -g [GUILD_ID] -o [OUTPUT_PATH]
```

## Install

### From source

```bash
> git clone git@github.com:hum/tomb.git
> cd tomb
> make build
> ./bin/tomb
```

### Go Install

```bash
> go install github.com/hum/tomb@latest
> tomb
```
