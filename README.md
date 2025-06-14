# gojq

gojq is a small command line utility to help with getting values from an arbitrary json file, for scripts and other purposes. it is designed to be both powerful for complex tasks and simple for easy use.

# installation

## From Source
```bash
git clone https://github.com/wltechblog/gojq.git
cd gojq
go build -o gojq
```

## Using Go Install
```bash
go install github.com/wltechblog/gojq@latest
```

After installation, make sure the binary is in your PATH or copy it to a directory that's in your PATH (e.g., `/usr/local/bin`).

# usage
```gojq [filename] [optional node]```

Running without specifying a node will give a list of the top level keys. When giving the node the next level of keys or values will be returned. Nodes are separated by "."

You can also use this as a filter in your shell pipeline like so:
```cat filename | gojq .node```

# example
```gojq test.json```
```json
{
  "key1": {
    "subkey1": true,
    "subkey2": false
  },
  "key2": [
    {
      "subkey3": null,
      "subkey4": []
    }
  ]
}
```
```gojq test.json key1.subkey1```
```true```
```gojq test.json key2.0.subkey3```
```null```
```gojq test.json key2.0.subkey4```
```[]```

