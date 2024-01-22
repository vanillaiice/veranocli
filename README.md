# Verano CLI

CLI tool to manage activities in a project.
This package uses under the hood the [verano](https://github.com/vanillaiice/verano)
package to do database transactions, parse files, and sort activities.

# Installation

```sh
$ go install github.com/vanillaiice/veranocli/cmd/veranocli@latest
```

# Features

- Sort activities in a project based on their relationships
(only start to finish relationships supported for now).
- Compute the start and finish times of all activities.
- Render a graph (with graphviz) image file showing the activities
and their relationships.
- Parse and process lists of activities in JSON, CSV, and XLSX formats
- Storage of the activities in a SQLite database.

# Usage

```sh
NAME:
   verano cli - manage activities in a project

USAGE:
   verano cli [global options] command [command options] 

VERSION:
   v0.0.7

AUTHOR:
   Vanillaiice <vanillaiice1@proton.me>

COMMANDS:
   parse, p  parse activities in json, csv, and xlsx formats
   db, d     make database transactions
   sort, s   topologically sort activities
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

# Example

Parsing a list of activities in xlsx format and generating a graph:

```sh
$ veranocli parse -f activities.xlsx -g activities.png
Id  Description   Duration  Start             Finish            Cost   
1   Tip landlord  6h0m0s    1 Jan 0001 00:00  1 Jan 0001 00:00  10000  
2   Buy eggs      20m0s     1 Jan 0001 00:00  1 Jan 0001 00:00  20     
3   Buy pan       15m0s     1 Jan 0001 00:00  1 Jan 0001 00:00  30     
4   Eat eggs      5m0s      1 Jan 0001 00:00  1 Jan 0001 00:00  0      
5   Cook eggs     10m0s     1 Jan 0001 00:00  1 Jan 0001 00:00  0      
6   Get money     36h0m0s   1 Jan 0001 00:00  1 Jan 0001 00:00  0      
```

![activities](https://github.com/vanillaiice/veranocli/assets/120596571/2c55d72a-40f2-4293-bdc2-58de09bb91f8)

# Author

Vanillaiice

# Licence

GPLv3
