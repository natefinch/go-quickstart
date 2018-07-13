# Go Quickstart

Each top level subdirectory of this repo contains a stub of a common project
layout - server, CLI app, etc.

They are laid out using what I consider to be best practices with the best tools
and libraries.  If you wish to suggest a replacement for a tool or library I'm
using here, feel free to file an issue, but without significant changes, I'm not
likely to change my mind.

Here's some of my choices of tools & libraries.  No tools are perfect, but these
are popular and well-tested, and in my opinion, will serve you well.

## Vendoring - [dep](github.com/golang/dep)

For now, dep is the standard in the go ecosystem.  This will change when/if
another tool becomes the standard (likely when vgo is integrated into the go
tool).

## CLI - [viper](github.com/spf13/viper) & [cobra](github.com/spf13/cobra) While

I have a few issues with the way these libraries work, but they're the best
  supported ones on the market, they make it easy to spin up a complex CLI, and
  you can mostly work around their wonkiness.
  
## Logging - [logrus](github.com/sirupsen/logrus) 

Structured logging is a good thing, and logrus is well supported and easy to
  use.  Other logging systems trade an awkward API for fewer allocations, and
  unless you're google or a trading system you shouldn't care about the speed of
  your logging system 

## Build tool - [mage](github.com/magefile/mage)

Users of make long ago realized the power of canonicalizing dev tasks in code.
But why use another language when you can use Go for that, too?  Mage makes it
easy to put common dev tasks at your fingertips.

## License

You may copy any of the code in this repo and relicense it in any way you want.

