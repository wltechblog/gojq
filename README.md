# jq

jq is a small command line utility to help with getting values from an arbitrary json file, for scripts and other purposes.

# usage
```jq [filename] [optional node]```

Running without specifying a node will give a list of the top level keys. When giving the node the next level of keys or values will be returned. Nodes are separated by "."

You can also use this as a filter in your shell pipeline like so:
```cat filename | jq .node```

# example
```jq test.json```
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
```jq test.json key1.subkey1```
```true```
```jq test.json key2.0.subkey3```
```null```
```jq test.json key2.0.subkey4```
```[]```

