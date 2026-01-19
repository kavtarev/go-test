import fs from 'fs'
import { execSync } from 'child_process'

// Читаем сертификат
const certPem = fs.readFileSync('localhost-cert.pem', 'utf-8');

// Декодируем сертификат
console.log(certPem);

// Если нужен вывод в формате JSON
const certInfo = execSync(`openssl x509 -in localhost-cert.pem -text -noout`).toString();
console.log(certInfo);
