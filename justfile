alias c := commit

commit msg="update" mode="chore":
	@git add .
	@git commit -m "{{mode}}: {{msg}}"

alias p := push

push msg="update" mode="chore":
    @just commit "{{msg}}" "{{mode}}"
    @git push
