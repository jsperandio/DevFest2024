UID := $(id -u)
GID := $(id -g)
# to = pdf,pptx
# iferr = sudo chown $LOGNAME 'Go – Do básico ao desafio técnico.'
slide:
	docker run --rm --init -v $$PWD:/home/marp/app/ -e MARP_USER=$(UID):$(GID) -e LANG=$$LANG marpteam/marp-cli 'Go – Do básico ao desafio técnico.md' --$(to)
slide-html:
	docker run --rm --init -v $$PWD:/home/marp/app/ -e MARP_USER=$(UID):$(GID) -e LANG=$$LANG marpteam/marp-cli 'Go – Do básico ao desafio técnico.md'
