# stream.py
from gpio.camera import Camera
import time

camera = Camera()

def main_loop():
    try:
        while True:
            frame = camera.get_frame()
            if frame:
                # Feed this frame to your Fyne app update method
                # For example: update_ui_with_frame(frame)
                print("Frame received")
            time.sleep(1 / 30)  # ~30 FPS
    except KeyboardInterrupt:
        print("Shutting down...")
        camera.close()

if __name__ == "__main__":
    main_loop()
