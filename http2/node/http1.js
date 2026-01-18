const http = require('http');

const server1 = http.createServer((req, res) => {
  res.writeHead(200, { 'Content-Type': 'text/plain' });
  res.end('Hello from HTTP/1.1!');
});

server1.listen(3000, () => {
  console.log('HTTP/1.1 server is listening on port 3000');
});
