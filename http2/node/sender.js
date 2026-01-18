import http2 from 'http2'
import { readFileSync } from 'fs';

const session = http2.connect('https://localhost:3001', { ca: readFileSync('localhost-cert.pem') })

const req = session.request({ ':path': '/' })

req.on('error', (e) => {
  console.log(1);
  console.log(e);

})

req.on('response', (headers, flags) => {
  console.log(99999);
  console.log(headers);
})

req.on('headers', h => {
  console.log(222);
  console.log(h);
})

let data = ''
req.on('data', c => {
  data += c
})

req.on('end', () => {
  console.log(data);
})

req.end()