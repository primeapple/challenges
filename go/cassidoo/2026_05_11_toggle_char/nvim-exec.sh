#!/bin/bash
# nvim-exec: pipe text through neovim and output the result
# Usage: echo "some text" | nvim-exec ":%s/foo/bar/g"
#    or: nvim-exec ":%s/foo/bar/g" < input.txt

COMMAND="${1:?Usage: nvim-exec <ex-command>}"

nvim - --clean --headless --cmd \
  + "set nonumber" \
  +"$COMMAND" \
  +'%print' \
  +'qa!'
