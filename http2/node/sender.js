const http2 = require('http2')
const fs = require('fs')

const session = http2.connect('https://localhost:3000', { ca: fs.readFileSync('localhost-cert.pem'), })
session.on('error', (err) => console.error(err));
const req = session.request({ ':path': '/' });

req.on('response', (headers) => {
  for (const name in headers) {
    console.log(`${name}: ${headers[name]}`);
  }
});

req.setTimeout()

req.setEncoding('utf8');
let data = '';
req.on('data', (chunk) => { data += chunk; });
req.on('end', () => {
  console.log(`\n${data}`);
  session.close();
});
req.end();