# Plugin

This package is a very simple wrapper loading a golang plugin with the goal
of making it easier to load in a work with a plugin.

### Motivation

Using the `plugin` golang package directly requires you to manually lookup
each property of the plugin you want to load and cast it to something useful.
The code to do this just to be able to get at the plugin's properties
becomes somewhat tedious boilerplate code. The goal of this package is to
allow you to just define an "interface" (through the use of a struct) that
you expect from the loaded plugin and this library will "hydrate" that struct
with the properties of the plugin. This let's you use the struct in your
progress and not have to deal the loading and validating the plugin.

### Testing

You can run the example test with the `make` command.

```bash
make test
```

It will build the example plugin and run the example that loads it.
You can run those commands manually if you want like this:

```bash
go generate
go test ./...
```
