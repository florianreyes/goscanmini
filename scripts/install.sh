#!/usr/bin/env bash

BINARY_NAME="goscanm"
INSTALL_DIR="/usr/local/bin"
SOURCE_FILE="../cmd/main.go"

# Ensure Go is installed
if ! command -v go &> /dev/null; then
    echo "[!] Go is not installed. Please install it first: https://go.dev/dl/"
    exit 1
fi

# Ensure script is run from the correct directory
if [ ! -f "$SOURCE_FILE" ]; then
    echo "[!] Error: '$SOURCE_FILE' not found in the current directory."
    exit 1
fi

# Build the Go binary
echo "[*] Building $BINARY_NAME..."
go build -o "$BINARY_NAME" "$SOURCE_FILE"

# Verify the build was successful
if [ ! -f "$BINARY_NAME" ]; then
    echo "[!] Build failed."
    exit 1
fi

# Prompt user for installation
read -p "[?] Do you want to add '$BINARY_NAME' to '$INSTALL_DIR'? (y/n): " choice

choice=${choice,,}  # Lowercase

if [[ "$choice" == "y" || "$choice" == "-y" ]]; then
    echo "[*] Moving '$BINARY_NAME' to '$INSTALL_DIR'..."
    
    # Move the binary to /usr/local/bin
    sudo mv "$BINARY_NAME" "$INSTALL_DIR/"  
    sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"

    # Verify installation
    if command -v "$BINARY_NAME" &> /dev/null; then
        echo "[+] Installation successful! You can now run '$BINARY_NAME'."
    else
        echo "[!] Installation failed. Check permissions and try again."
        exit 1
    fi
else
    echo -e "[*] Installation skipped. \n[*] Moving to directory root"
    # Move the binary to /usr/local/bin
    mv "$BINARY_NAME" ../ 
    chmod +x "../$BINARY_NAME"  
fi
