package com.foot.overlay

import android.os.Bundle
import android.widget.SeekBar
import android.widget.TextView
import androidx.appcompat.app.AppCompatActivity

class SettingsActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_settings)

        // Hier können Overlay-Einstellungen (Transparenz, Position, etc.) konfiguriert werden
        // und an das Backend gesendet werden.
        
        val opacitySeekbar = findViewById<SeekBar>(R.id.opacitySeekbar)
        val opacityValue = findViewById<TextView>(R.id.opacityValue)
        
        opacitySeekbar?.setOnSeekBarChangeListener(object : SeekBar.OnSeekBarChangeListener {
            override fun onProgressChanged(seekBar: SeekBar?, progress: Int, fromUser: Boolean) {
                opacityValue?.text = "$progress%"
                // Hier könnte die Config an das Backend gesendet werden
            }
            override fun onStartTrackingTouch(seekBar: SeekBar?) {}
            override fun onStopTrackingTouch(seekBar: SeekBar?) {}
        })
    }
}
