const express = require('express');
const graylog2 = require('graylog2');

const app = express();

// Middleware to parse request bodies
app.use(express.json());

// Graylog configuration
const graylogConfig = {
  servers: [{ host: 'localhost', port: 12201 }],
  hostname: 'sample-node', // Optional
};

const graylog = new graylog2.graylog(graylogConfig);

// Express route
app.post('/sample/test', (req, res) => {
  const message = req.body.message;

  // Log the message to Graylog
  graylog.log(message, 'Sample log');

  res.sendStatus(200);
});

// Start the server
const port = 3000;
app.listen(port, () => {
  console.log(`Server is listening on port ${port}`);
});
