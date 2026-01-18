import { createSecureServer, constants } from 'node:http2';
const {
  HTTP2_HEADER_METHOD,
  HTTP2_HEADER_PATH,
  HTTP2_HEADER_STATUS,
  HTTP2_HEADER_CONTENT_TYPE,
} = constants;
import fs from 'fs'

const options = {
  key: fs.readFileSync('localhost-privkey.pem'),
  cert: fs.readFileSync('localhost-cert.pem'),
};

const server = createSecureServer(options);

server.on('stream', (stream, headers) => {
  const method = headers[HTTP2_HEADER_METHOD];
  const path = headers[HTTP2_HEADER_PATH];
  console.log(method);
  console.log(path);

  stream.respond({
    [HTTP2_HEADER_STATUS]: 200,
    [HTTP2_HEADER_CONTENT_TYPE]: 'text/plain; charset=utf-8',
  });
  stream.write('some\n');
  stream.end('any')
})

server.listen(3001, () => { console.log('up on 3001') })