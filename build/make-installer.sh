#!/bin/bash

set -e

PWDDIR=`dirname "${0}"`
cd "$PWDDIR"

archive=$1
script=$2

tmp=__tmp__$RANDOM

printf "#!/bin/bash
if [[ \"\$HOSTTYPE\" != x86_64 ]]; then
	echo \"Unsupported OS architecture\"
	exit 1
fi

DATA_LINE=\`awk '/^__DATA_END__/ {print NR + 1; exit 0; }' \$0\`
tail -n+\$DATA_LINE \$0 | tar -xvz

# Run as sudo or root
if [ \"\$(id -u)\" != \"0\" ]; then
   exec sudo \"\$0\" \"\$@\"
fi

mv -f oneandone /usr/local/bin/oneandone
chmod +x /usr/local/bin/oneandone

if [[ \"\$OSTYPE\" != linux* ]]; then
	echo
	echo \"To enable bash auto-completion for 1&1 CLI, add the following command in your .bashrc file.\"
	echo \"PROG=oneandone source bash_autocomplete\"
	echo \"Make sure that bash_autocomplete file is in your PATH or specify the full path to the file.\"
else
	# move bash_autocomplete file and rename it to match binary name
	mv -f bash_autocomplete /etc/bash_completion.d/oneandone
	echo
	echo \"To enable bash auto-completion for 1&1 CLI, run the following command:\"
	echo \"source /etc/bash_completion.d/oneandone\"
fi

exit 0
__DATA_END__\n" > "$tmp"

cat "$tmp" "$archive" > "$script" && rm "$tmp"
chmod +x "$script"
