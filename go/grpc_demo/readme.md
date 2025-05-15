## 备注

---

#### 生成自签 TLS 证书（仅测试用）
```
mkdir certs
# 生成 CA 证书
openssl req -x509 -newkey rsa:2048 -days 365 -nodes \
  -keyout certs/ca-key.pem -out certs/ca.pem \
  -subj "/CN=My Root CA"

# 生成服务器证书（包含 SAN）
openssl req -newkey rsa:2048 -nodes -keyout certs/server-key.pem \
  -out certs/server.csr -subj "/CN=localhost"
openssl x509 -req -in certs/server.csr -days 365 -CA certs/ca.pem \
  -CAkey certs/ca-key.pem -CAcreateserial -out certs/server.pem \
  -extfile <(echo "subjectAltName=DNS:localhost")
```

#### nginx
```
nginx -p nginx -c nginx.conf
nginx -p nginx -c nginx.conf -s stop
```