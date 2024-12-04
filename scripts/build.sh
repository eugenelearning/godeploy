#!/usr/bin/env bash
targets=("windows/amd64" "windows/386" "darwin/amd64" "darwin/arm64" "linux/amd64" "linux/arm64")

for target in "${targets[@]}"; do
	parts=(${target//\// })

	GOOS=${parts[0]}
	GOARCH=${parts[1]}

	output_name='build/godeploy-'$GOOS'-'$GOARCH;

	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi	

    echo $output_name

	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name

	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
done