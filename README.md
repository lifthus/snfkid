# snofkid

int64 Snowflake ID generator in Go

## As a command

### Installation

```bash
    go install github.com/lifthus/snofkid/cmd/snofkid@latest
```

### Usage

```bash
    # [-e] [-m]
    # just generating Snowflake based on given epoch and machine-id
    # default values of both are Twitter epoch and 0
    snofkid [-e <epoch>] [-m <machine-id>]

    # Start Snowflake generator server on given port.
    # Default epoch is Twitter epoch.
    # HTTPS is enabled if both cert and key files are provided.
    # e.g. a request to path "/:machine-id" will return a Snowflake ID of the given machine-id.
    snofkid -p <port> [-e <epoch>] [-cert <cert> -key <key>]
```

## As a library

### Installation

```bash
    go get github.com/lifthus/snofkid
```

### Usage

```go
    import "github.com/lifthus/snofkid"

    ...
    mch, err := snofkid.NewMachine(epoch, machineID)

    snf, err := mch.New()

    t := mch.ParseTime(snf)
    ...
```
