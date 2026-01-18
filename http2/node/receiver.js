import http2 from 'http2'
import fs from 'fs'

const options = {
  key: fs.readFileSync('localhost-privkey.pem'),
  cert: fs.readFileSync('localhost-cert.pem'),
};

const server = http2.createSecureServer(options)

server.on('stream', (stream, headers) => {
  console.log(123123);
  stream.respond({
    'content-type': 'text/html; charset=utf-8',
    ':status': 200,
  });
  stream.end('hui pizda')
})

server.listen(3000, () => console.log('up on 3000'))