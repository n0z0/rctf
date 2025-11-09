package main

import (
	"bytes"
	_ "embed"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/gen2brain/beeep"
)

//go:embed art/tot2wuk2.wav
var notifWav []byte

//go:embed art/indo.png
var indoPng []byte

func notifikasi(pesan string) error {
	// 2) buat file sementara
	tmpDir := os.TempDir()
	iconPath := filepath.Join(tmpDir, "indo.png")

	// tulis isi embed ke file
	if err := os.WriteFile(iconPath, indoPng, 0o644); err != nil {
		return err
	}

	// 3) pakai di beeep
	if err := beeep.Notify("Judul Notif", pesan, iconPath); err != nil {
		return err
	}

	return nil
}

func playNotificationSound() {
	// decode WAV dari byte yang di-embed
	streamer, format, err := wav.Decode(bytes.NewReader(notifWav))
	if err != nil {
		log.Fatal("gagal decode wav:", err)
	}
	defer streamer.Close()

	// init speaker sesuai format audio
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal("gagal init speaker:", err)
	}

	done := make(chan struct{})

	// play sampai selesai
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		close(done)
	})))

	<-done
}
