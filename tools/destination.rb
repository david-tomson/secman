#!/usr/bin/ruby -W
require 'colorize'

$l = `verx abdfnx/secman -l`
$c = `verx -c`

def _n()
    ly = $l.cyan.bold
    nr = "there's a new release of secman is avalaible:".yellow
    up = "to update it run".yellow
    smu = "secman upd".blue
    puts new_r = "#{nr} #{ly}#{up} #{smu}"
end

def check()
    if $l != $c
        _n
    end
end
