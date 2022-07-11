all: html view

html:
	marp -I . --allow=local-files

pdf:
	marp -I . --allow=local-files --pdf

view:
	browse README.html

.PHONY: all pdf html