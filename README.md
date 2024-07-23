# jq

jq is a small command line utility to help with getting values from an arbitrary json file, for scripts and other purposes.

# usage
```jq [filename] [optional node]```

Running without specifying a node will give a list of the top level keys. When giving the node the next level of keys or values will be returned. Nodes are separated by "."
