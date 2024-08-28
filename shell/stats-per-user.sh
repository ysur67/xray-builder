#!/bin/env bash

xray api statsquery | jq 'reduce .stat[] as $item ({};
  if $item.name | startswith("user>>>") then
    ($item.name | split(">>>")) as $keys |
    .[$keys[1]] += { ($keys[3] + "-mb"): ($item.value | tonumber / 1024 / 1024) | round }
  else
    .
  end
)'
