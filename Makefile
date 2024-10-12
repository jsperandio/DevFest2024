# to -> pdf,pptx
slide:
	docker run --rm --init -v $PWD:/home/marp/app/ -e MARP_USER="$(id -u):$(id -g)" -e LANG=$LANG marpteam/marp-cli 'Go – Do básico ao desafio técnico.md' --$(to)
