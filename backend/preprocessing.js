const express = require('express');
const app = express();
app.use(express.json());

app.post('/preprocess', (req, res) => {
    const data = req.body;
    // Perform preprocessing here
    res.json({status: 'success'});
});

app.listen(3000, () => {
    console.log('Preprocessing service listening on port 3000');
});