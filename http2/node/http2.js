const http2 = require('http2');
const fs = require('fs');

const server2 = http2.createSecureServer({
  key: fs.readFileSync('localhost-privkey.pem'),
  cert: fs.readFileSync('localhost-cert.pem'),
}, (req, res) => {
  res.writeHead(200, { 'Content-Type': 'text/plain' });
  res.end('Hello from HTTP/2!');
});

server2.listen(8443, () => {
  console.log('HTTP/2 server is listening on port 8443');
});
