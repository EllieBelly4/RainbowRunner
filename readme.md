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
4. Run `go run .` - Runs auth server on port `2110`, gameserver on port `2603` (not configurable)
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
* Player movement (limited to static areas)
* Player equipment (only visuals)
* Player inventory
* Chat
* NPC Spawning (static only)
* Various bits of GUI information can be modified

#### Additional non-game related server features

* GraphQL API - Can be used to retrieve player/zone/entities etc. See: [internal/api](internal/api)
* Command API - Can be used to send raw packets directly to clients through the server
  See: [internal/admin](internal/admin)

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

To run all in here commands use: `go run cmd/rrcli`

* `dump` - Parses all Dungeon Runners configuration files in a directory and outputs a compiled JSON that is readable by
  RainbowRunner, use after you run `cmd/scan_pkg.go` to extract the config files.
* `get` - Retrieve configuration data for specific GCObjects with text/regexp support, can also limit to categories
  e.g. "Armor", See for all
  categories: [resources/Dumps/generated/drcategories.json](resources/Dumps/generated/drcategories.json)
* `categorise` - Parse the dumped configuration from `dump` and generate the category config
  file ([resources/Dumps/generated/drcategories.json](resources/Dumps/generated/drcategories.json)).

## Contributing

Contributions are welcome, the focus right now is to get everything working, don't try to make major refactor PRs, I'm
aware the codebase is a mess.

1. Fork
2. Open PR
3. Get approved and merge!

## License

[GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/)
