package client

import (
	"bytes"
	"image/jpeg"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (c *Client) streamRoomBtn() *widget.Button {

	btn := widget.NewButton("Stream", func() {
		c.Window.SetContent(c.StreamRoom())
	})

	return btn
}
func (c *Client) StreamRoom() *fyne.Container {
	img := canvas.NewImageFromImage(nil)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(640, 480))

	URL := "http://127.0.0.1:5000/video_feed"

	go func() {
		for {
			resp, err := http.Get(URL)
			if err != nil {
				log.Println("Stream error:", err)
				time.Sleep(time.Second)
				continue
			}

			contentType := resp.Header.Get("Content-Type")
			_, params, err := mime.ParseMediaType(contentType)
			if err != nil {
				log.Println("Failed to parse media type:", err)
				resp.Body.Close()
				continue
			}

			boundary := params["boundary"]
			if boundary == "" {
				log.Println("Boundary not found in Content-Type")
				resp.Body.Close()
				continue
			}
			log.Printf("%v\n", resp.Body)
			mr := multipart.NewReader(resp.Body, boundary)

			for {
				part, err := mr.NextPart()
				if err != nil {
					log.Println("Read part error:", err)
					break
				}

				// Skip part headers and extract only the JPEG bytes
				buf := new(bytes.Buffer)
				_, err = buf.ReadFrom(part)
				if err != nil {
					log.Println("Buffer read error:", err)
					continue
				}

				// Decode only the JPEG image bytes
				imgData, err := jpeg.Decode(buf)
				if err != nil {
					log.Println("JPEG decode error:", err)
					continue
				}

				// Update the image on the canvas
				img.Image = imgData
				img.Refresh()

				// For ~30fps
				time.Sleep(time.Second / 30)
			}

			resp.Body.Close()
		}
	}()

	return container.NewCenter(container.NewVBox(
		widget.NewLabel("Streaming here :3"),
		img,
	))
}
