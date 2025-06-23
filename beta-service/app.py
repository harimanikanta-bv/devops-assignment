from flask import Flask
app = Flask(__name__)

@app.route('/beta')
def beta():
    return "Hello from Beta Service!"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001)
