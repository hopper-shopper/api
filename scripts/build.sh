realpath() {
    [[ $1 = /* ]] && echo "$1" || echo "$PWD/${1#./}"
}

SCRIPT_PATH=$(realpath "$0")
SCRIPTS_DIR=$(dirname "$SCRIPT_PATH")

echo "Scripts dir: $SCRIPTS_DIR"
echo "Script path: $SCRIPT_PATH"

cd $SCRIPTS_DIR

./build-macos.sh
./build-alpine.sh