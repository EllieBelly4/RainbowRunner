# Resources

Game resources are stored as packages in game.pkg. Packages are indexed in the game.pki(package index).

There is a resource manager config in game directory [resourcemanager.cfg](../Files/resourcemanager.cfg)

Index is read in `PackageIndex::Read()`

## Index
File size is (0x4215 * 5) * 8

## Compression 

Compression for packages is Zlib INFLATE v1.2.3

## Caching

Resources are cached based on [resourcemanager.cfg](../Files/resourcemanager.cfg)

* SimpleResourceCache

Resources seem to go through a cache and can be retrieved with a string ID. Resource keys are case-insensitive

## Types

### Zones

### Fonts

* Found string: %!PS-Adobe-3.0 Resource-CIDFont
