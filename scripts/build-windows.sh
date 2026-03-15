#/bin/bash
BINARY_NAME="photofiler"
echo "📦 Building for Windows..."
gogio -target windows -icon icon.png -o ${BINARY_NAME}.exe .
echo "🗑️  Removing temporary files..."
rm ${BINARY_NAME}_windows_amd64.syso
echo "🎉 Done!"
echo "🔢 Binary size: $(du -k ${BINARY_NAME}.exe | cut -f1) KB"