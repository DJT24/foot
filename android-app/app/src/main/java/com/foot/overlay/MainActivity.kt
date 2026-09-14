package com.foot.overlay

import android.content.Intent
import android.net.Uri
import android.os.Build
import android.os.Bundle
import android.provider.Settings
import android.widget.Button
import android.widget.EditText
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity

class MainActivity : AppCompatActivity() {

    private lateinit var urlInput: EditText
    private lateinit var startButton: Button
    private lateinit var stopButton: Button

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        urlInput = findViewById(R.id.urlInput)
        startButton = findViewById(R.id.startButton)
        stopButton = findViewById(R.id.stopButton)

        // Standard-URL für Emulator (Backend läuft auf Host-Maschine)
        urlInput.setText("http://10.0.2.2:8080/v1/matches/current/overlay")

        startButton.setOnClickListener {
            if (Settings.canDrawOverlays(this)) {
                val intent = Intent(this, OverlayService::class.java)
                intent.putExtra("OVERLAY_URL", urlInput.text.toString())
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
                    startForegroundService(intent)
                } else {
                    startService(intent)
                }
                Toast.makeText(this, "Overlay gestartet", Toast.LENGTH_SHORT).show()
            } else {
                val intent = Intent(
                    Settings.ACTION_MANAGE_OVERLAY_PERMISSION,
                    Uri.parse("package:$packageName")
                )
                startActivity(intent)
                Toast.makeText(this, "Overlay-Permission benötigt", Toast.LENGTH_LONG).show()
            }
        }

        stopButton.setOnClickListener {
            val intent = Intent(this, OverlayService::class.java)
            stopService(intent)
            Toast.makeText(this, "Overlay gestoppt", Toast.LENGTH_SHORT).show()
        }
    }
}
