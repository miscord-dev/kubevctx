_kubevctx_hook() {
  trap -- '' SIGINT;
  eval "$("kubevctx" export zsh)";
  trap - SIGINT;
}
typeset -ag precmd_functions;
if [[ -z "${precmd_functions[(r)_kubevctx_hook]+1}" ]]; then
  precmd_functions=( _kubevctx_hook ${precmd_functions[@]} )
fi
