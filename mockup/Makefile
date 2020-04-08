.POSIX:

all: main.html style.css newjob.html login.html

refresh: all
	reload-browser Firefox

dev: current
	entr make refresh < current

current: Makefile
	echo main.lml > current
	echo style.css >> current
	echo newjob.lml >> current
	echo login.lml >> current

%.html: %.lml
	lml $<

style.css: tailwind.css
	tailwind build $< -o $@

.PHONY: all refresh
