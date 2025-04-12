# gpio/camera.py

import cv2
from PIL import Image
import numpy as np
import threading

class Camera:
    def __init__(self):
        self.cap = cv2.VideoCapture(0)
        if not self.cap.isOpened():
            raise RuntimeError("Error opening camera")
        self.frame = None
        self.lock = threading.Lock()
        self.running = True

        # Start background frame reader
        self.thread = threading.Thread(target=self._update_frame, daemon=True)
        self.thread.start()

    def _update_frame(self):
        while self.running:
            ret, frame = self.cap.read()
            if ret:
                with self.lock:
                    self.frame = self._mat_to_rgba(frame)

    def _mat_to_rgba(self, frame):
        rgba = cv2.cvtColor(frame, cv2.COLOR_BGR2RGBA)
        return Image.fromarray(rgba)

    def get_frame(self):
        with self.lock:
            return self.frame.copy() if self.frame else None

    def close(self):
        self.running = False
        self.cap.release()
