# talosctx

Faster context switching for `talosctl`, inspired by [@ahmetb/kubectx](https://github.com/ahmetb/kubectx).

Switching the configuration context in `talosctl` is already quite ergonomic when compared to `kubectl`. However, if you find yourself managing more than a handful of clusters, `talosctx` can help you search and switch between them very quickly.

## Examples

```
# launch the built-in fuzzy finder for interactive selection
$> talosctx
Switched to context "i-love-fzf".

# switch to a known context
$> talosctx hyper
Switched to context "hyper".

# switch to the previous context
$> talosctx -
Switched to context "i-love-fzf".
```

## Installing

`talosctx` is available in Homebrew:

```
brew install hugginsio/tap/talosctx
```

Alternatively, you can download binaries for your platform from [the latest release](https://github.com/hugginsio/talosctx/releases/latest).

## Configuration

Currently, there is no configuration. `talosctx` uses the same logic as `talosctl` to determine where your configuration file is located. At the time of writing, the following filepaths are checked:

> Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.

When a new context is selected, the previously selected context will be written to `$XDG_DATA_HOME/talosctx` as plain text. This is used to inform the `talosctx -` command and to pre-select your previous context in the fuzzy finder.

## Legal

"Talos" is a trademark of Sidero Labs. The `talosctx` tool is not endorsed by, sponsored by, or affiliated with Sidero Labs. My use of this name (and brand) is for identification purposes only, and does not imply any such endorsement, sponsorship, or affiliation.
