# Rainbow Runner

Private server + tools for Dungeon Runners.

This is very much a work in progress and is not close to completion and definitely not user-friendly.

## Requirements

* Golang >=1.19 - Server + CLI tools
* [ANTLR](https://www.antlr.org/index.html) - Generating config parser from language
  file([antlr/DRConfig.g4](antlr/DRConfig.g4))

## How to run

1. Clone repo
2. Copy `config.example.yaml` to `config.yaml`
3. Run `go get`
4. Run `go run .` - Runs auth server on port `2110`, gameserver on port `2603` (configurable
   in [config.yaml](./config.example.yaml))
5. Modify `config/DungeonRunners.cfg` and make sure `[AuthServer]`->`Address` is `localhost`
6. Start Dungeon Runners

## Features

All features are implemented in a static/basic way currently and are only meant to serve as a way of getting things
running.
This means that the game is in no way playable and is not meant to be until a significant number of systems have been
fully understood.

### Server

* Login
* World list
* Character creation
* Loading into zones (hardcoded values used to choose zone)
* Player movement
* Player equipment (only visuals)
* Player inventory
* Chat
* NPC Spawning (static only)
* Various bits of GUI information can be modified
* Skills (only visuals)

#### Additional non-game related server features

* GraphQL API - Can be used to retrieve player/zone/entities etc. See: [internal/api](internal/api)
* Command API - Can be used to send raw packets directly to clients through the server
  See: [internal/admin](internal/admin)

### Docs

There is some documentation, but it is not updated frequently: [Dungeon Runners](resources/Docs/v2/DungeonRunners.md)

### Tools

All tools are found under the [cmd](cmd) directory, they currently have lots of hardcoded values in them so for the time
being
they will have to be modified directly to change parameters.

* [cmd/hash](cmd/hash) - Generate GCObject hash for single string
* [cmd/generate_hashes.go](cmd/generate_hashes.go) - Generate hashes for all objects
  in [resources/Dumps/GCDictionary.txt](resources/Dumps/GCDictionary.txt)
* [cmd/scan_pkg.go](cmd/scan_pkg.go) - Scans and extracts `game.pkg` - requires the decompressed `game.pki` (
  use `cmd/uncompress_pki.go`)
* [cmd/uncompress_pki.go](cmd/uncompress_pki.go) - Decompresses `game.pki`
* [cmd/file_extensioniser.go](cmd/file_extensioniser.go) - Adds appropriate file extensions to extracted files
  from `game.pkg`
* [cmd/staticworldtounity](cmd/staticworldtounity) - Generates metadata for all 3D models in a format that can be used
  within Unity for testing

#### [cmd/rrcli](cmd/rrcli) - Slightly more final CLI interface

To run all commands in here use: `go run cmd/rrcli`

#### Config `rrcli config`

Commands for working with DR game configuration.

* `get` - Retrieve configuration data for specific GCObjects with text/regexp support, can also limit to categories
  e.g. "Armor"
* `list` - List all GCObjects with simple filter and depth options

#### Config Extract `rrcli config extract`

Commands for extracting data from the configuration files.

* `dump` - Parses all Dungeon Runners configuration files in a directory and outputs a compiled JSON that is readable by
  RainbowRunner, use after you run `cmd/scan_pkg.go` to extract the config files.
* `categorise` - Parse the dumped configuration from `dump` and generate the category config
  file ([resources/Dumps/generated/drcategories.json](resources/Dumps/generated/drcategories.json)).

#### Config Category `rrcli config category`

Commands which attempt to group GCObjects into categories for helping with discovery of objects.
These categories are not always the clearest or most useful but can be in very specific circumstances.

* `list` - Retrieve a list of all categories that can be searched, use with `-d $NUMBER` to increase
  depth, `-1` for unlimited
* `get` - Get a specific category details

#### Models `rrcli models`

Commands for working with DR format models

* `convert` - Convert models from `.3dnode` to `.obj` and `.mtl`

## Contributing

Contributions are welcome, the focus right now is to get everything working, don't try to make major refactor PRs, I'm
aware the codebase is a mess.

1. Fork
2. Open PR
3. Get approved and merge!

Before merging a PR please ensure that you have run `go fmt ./...` and `go generate ./...`.
Automated checks to come later if there are any contributions to this project.

## License

[GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/)
