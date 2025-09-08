# talosctx completion for fish

function __fish_talosctx_get_contexts
    talosctl config contexts 2>/dev/null | string match -v -r '^(CURRENT|NAME|\-\-\-|\*?\s*CURRENT|Context)' | string replace -r '^\*?\s*' '' | string match -r '^\S+' | string trim
end

# Complete talosctx with context names for the first argument only
complete -c talosctx -f -n 'test (count (commandline -opc)) -eq 1' -a '(__fish_talosctx_get_contexts)' -d 'Talos context'

# Add the special '-' option to switch to previous context
complete -c talosctx -f -n 'test (count (commandline -opc)) -eq 1' -a '-' -d 'Switch to previous context'
