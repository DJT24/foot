# Android Overlay App

## Emulator-Setup

1. **Android Studio installieren**
   - Download: https://developer.android.com/studio

2. **AVD (Android Virtual Device) erstellen**
   - Tools → Device Manager → Create Device
   - Empfohlen: Pixel 6, API 34 (Android 14)

3. **Overlay-Permission im Emulator aktivieren**
   - Settings → Apps → Special app access → Display over other apps → Foot Overlay → Allow

4. **Backend auf Host-Maschine starten**
   ```powershell
   cd backend
   go run cmd\server\main.go
   ```

5. **App im Emulator testen**
   - Run → Run 'app'
   - URL: `http://10.0.2.2:8080/v1/matches/current/overlay`
   - "Overlay starten" klicken

## Wichtige Hinweise

- **Emulator-Netzwerk**: `10.0.2.2` = Host-Maschine (dein PC)
- **Cleartext-Traffic**: Erlaubt in `AndroidManifest.xml` für lokales Testing
- **Foreground-Service**: Benötigt für Overlay ab Android 8+

## Build & Run

```bash
cd android-app
./gradlew assembleDebug
```

APK liegt in `app/build/outputs/apk/debug/app-debug.apk`
