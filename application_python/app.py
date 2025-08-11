from flask import Flask

app = Flask(__name__)

@app.route("/")
def home():
    return "Welcome!"

@app.route("/how-are-you")
def how_are_you():
    return "I am good, how about you?"

if __name__ == "__main__":
    print("Server starting on http://0.0.0.0:8085")
    app.run(host="0.0.0.0", port=8085)
