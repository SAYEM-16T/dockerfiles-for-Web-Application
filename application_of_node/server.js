// Load express
const express = require('express');
const app = express();
const port = 8085;

// Route: "/"
app.get('/', (req, res) => {
    res.send('Welcome!');
});

// Route: "/how-are-you"
app.get('/how-are-you', (req, res) => {
    res.send('I am good, how about you?');
});

// Start the server
app.listen(port, '0.0.0.0', () => {
    console.log(`Server starting on http://0.0.0.0:${port}`);
});
