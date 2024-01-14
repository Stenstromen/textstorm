# TextStorm

![TextStorm icon](icon.png)

```bash
# Build for amd64
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o textstorm_amd64
# Build for arm64
GOOS=darwin GOARCH=arm64 go build -o textstorm_arm64

# Create a universal binary
lipo -create -output textstorm textstorm_amd64 textstorm_arm64

fyne package -os darwin -name "Textstorm" -icon big_icon.png --executable textstorm --appID se.stenstromen.textstorm --appVersion 1.0.0 --release

fyne-cross darwin -name "Textstorm" -icon big_icon.png -app-id se.stenstromen.textstorm -app-version 1.0.0 -release
```
