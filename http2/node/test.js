const http = require('http');
const http2 = require('http2');
const { PerformanceObserver, performance } = require('perf_hooks');

async function testHttp1() {
  const requests = 100; // Количество запросов
  const url = 'http://localhost:3000';

  for (let i = 0; i < requests; i++) {
    performance.mark(`start-${i}`);
    await new Promise((resolve) => {
      http.get(url, (res) => {
        res.on('data', () => { });
        res.on('end', () => {
          performance.mark(`end-${i}`);
          resolve();
        });
      });
    });
  }
}

async function testHttp2() {
  const requests = 100;
  const client = http2.connect('https://localhost:8443', {
    rejectUnauthorized: false // Отключаем проверку сертификата для тестирования
  });

  for (let i = 0; i < requests; i++) {
    performance.mark(`start-${i}`);
    await new Promise((resolve) => {
      const req = client.request({ ':path': '/' });
      req.on('response', (headers, flags) => {
        req.on('data', () => { });
        req.on('end', () => {
          performance.mark(`end-${i}`);
          resolve();
        });
      });
      req.end();
    });
  }
  client.close();
}

async function runTests() {
  console.log('Testing HTTP/1.1...');
  await testHttp1();
  console.log('Testing HTTP/2...');
  await testHttp2();

  const obs = new PerformanceObserver((list) => {
    list.getEntries().forEach((entry) => {
      console.log(`${entry.name}: ${entry.duration}`);
    });
  });
  obs.observe({ entryTypes: ['mark', 'measure'] });

  // Measure time for HTTP/1.1
  for (let i = 0; i < 100; i++) {
    performance.measure(`http1-response-${i}`, `start-${i}`, `end-${i}`);
  }

  // Measure time for HTTP/2
  for (let i = 0; i < 100; i++) {
    performance.measure(`http2-response-${i}`, `start-${i}`, `end-${i}`);
  }
}

runTests();
