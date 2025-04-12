from flask import Flask, Response
import cv2

app = Flask(__name__)

# Use Raspberry Pi camera (0) or adjust index
camera = cv2.VideoCapture(0)

def generate_frames():
    while True:
        success, frame = camera.read()
        if not success:
            continue

        # Encode as JPEG
        ret, buffer = cv2.imencode('.jpg', frame)
        if not ret:
            continue

        frame_bytes = buffer.tobytes()

        # Yield MJPEG formatted frame
        yield (b'--frame\r\n'
               b'Content-Type: image/jpeg\r\n\r\n' + frame_bytes + b'\r\n')

@app.route('/video_feed')
def video_feed():
    return Response(generate_frames(),
                    mimetype='multipart/x-mixed-replace; boundary=frame')


@app.route("/")
def hi():
    return Response("Hi")
if __name__ == '__main__':
    app.run(port=5000)
