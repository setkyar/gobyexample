#!/bin/bash

public="public"
gh_public="public-gh"

copy_files() {
	cp -r $public/* $gh_public
}

set_file_extension() {
	for file in $gh_public/*; do
		if [[ $file != *.* ]]; then
			mv $file $file.html
		fi
	done
}

copy_files
set_file_extension