test:
	@# delete if uuidgen is widely installed
	@# printf '%04x%04x-%04x-4%03x-%04x-%04x%04x%04x' $(od -An -N2 -tu2 /dev/urandom) $(od -An -N2 -tu2 /dev/urandom) $(od -An -N2 -tu2 /dev/urandom) $(od -An -N2 -tu2 /dev/urandom) $(od -An -N2 -tu2 /dev/urandom) $(od -An -N2 -tu2 /dev/urandom)
	UUID_ENV_VAR=$$(uuidgen) go test -trimpath ./...
	# UUID_ENV_VAR=1867BF9B-E510-400D-9597-08A88006177E go test -trimpath ./...
