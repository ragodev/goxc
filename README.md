# goxc

Go cross platform compilation utility

## Usage

```
$ ./goxc linux_amd64.json windows_x86.json...
```

## Configuration file

[Possible combinations](https://golang.org/doc/install/source#environment)

Example config file
```json
{
    "os": "linux",
    "arch": "amd64",
    "out": "file"
}
```

You can include multiple configs in one file in an array
```json
[
    {
        "os": "linux",
        "arch": "amd64",
        "out": "file"
    },
    {
        ...
    }
]
```

## Included configs

- `common.json` - Compile for all common OS combinations
    - `linux x86`
    - `linux amd64`
    - `macOS amd64`
    - `windows x86`
    - `windows amd64`
