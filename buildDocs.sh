echo "Building documentation page"
aglio --theme-variables streak -i docs/main.md -o docs/index.html
printf "Done\n"
