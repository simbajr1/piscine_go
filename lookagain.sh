# Find all files ending with .sh, strip the extension, remove './' and directories
find . -type f -name "*.sh" | sed 's|^\./||' | sed 's/\.sh$//' | awk -F/ '{print $NF}' | sort -r
