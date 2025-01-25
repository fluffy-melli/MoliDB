## 🐍 python

```py
import os
import json
import requests
import gzip
import base64
from io import BytesIO
from Crypto.Cipher import AES
from Crypto.Util.Padding import pad, unpad
from dotenv import load_dotenv

load_dotenv()
SERVER_URL = "http://127.0.0.1:17233"

SECRET_KEY = "ThisIs32byteAESkeyForThisExample"
API_TOKEN  = 'ThisIsExampleAPIKey'

headers = {
    'Authorization': f"{API_TOKEN}",
    'Content-Type': 'application/json'
}

def aes_encrypt(data: bytes) -> bytes:
    cipher = AES.new(SECRET_KEY.encode(), AES.MODE_CBC)
    ct_bytes = cipher.encrypt(pad(data, AES.block_size))
    return cipher.iv + ct_bytes

def aes_decrypt(data: bytes) -> bytes:
    iv = data[:16]
    ct = data[16:]
    cipher = AES.new(SECRET_KEY.encode(), AES.MODE_CBC, iv=iv)
    decrypted = unpad(cipher.decrypt(ct), AES.block_size)
    return decrypted

def gzip_compress(data: bytes) -> bytes:
    buf = BytesIO()
    with gzip.GzipFile(fileobj=buf, mode='wb') as f:
        f.write(data)
    return buf.getvalue()

def gzip_decompress(data: bytes) -> bytes:
    buf = BytesIO(data)
    with gzip.GzipFile(fileobj=buf, mode='rb') as f:
        return f.read()

def create_collection(data):
    json_data = json.dumps(data).encode()
    compressed_data = gzip_compress(json_data)
    encrypted_data = aes_encrypt(compressed_data)
    encrypted_data_base64 = base64.b64encode(encrypted_data).decode('utf-8')
    headers['body'] = encrypted_data_base64
    response = requests.post(f"{SERVER_URL}/collection", headers=headers)
    if response.status_code == 200:
        encrypted_data_base64 = response.json().get("data")
        encrypted_data = base64.b64decode(encrypted_data_base64)
        aes_decrypt_data = aes_decrypt(encrypted_data)
        decompressed_data = gzip_decompress(aes_decrypt_data)
        data = json.loads(decompressed_data)
        return data
    else:
        raise f"HTTP FAIL : {response.json()}"

def get_collection(id):
    response = requests.get(f"{SERVER_URL}/collection/{id}", headers=headers)
    if response.status_code == 200:
        encrypted_data_base64 = response.json().get("data")
        encrypted_data = base64.b64decode(encrypted_data_base64)
        decrypted_data = aes_decrypt(encrypted_data)
        decompressed_data = gzip_decompress(decrypted_data)
        data = json.loads(decompressed_data)
        return data
    else:
        raise f"HTTP FAIL : {response.json()}"

def update_collection(id, data):
    json_data = json.dumps(data).encode()
    compressed_data = gzip_compress(json_data)
    encrypted_data = aes_encrypt(compressed_data)
    encrypted_data_base64 = base64.b64encode(encrypted_data).decode('utf-8')
    headers['body'] = encrypted_data_base64
    response = requests.put(f"{SERVER_URL}/collection/{id}", headers=headers)
    if response.status_code == 200:
        encrypted_data_base64 = response.json().get("data")
        encrypted_data = base64.b64decode(encrypted_data_base64)
        aes_decrypt_data = aes_decrypt(encrypted_data)
        decompressed_data = gzip_decompress(aes_decrypt_data)
        data = json.loads(decompressed_data)
        return data
    else:
        raise f"HTTP FAIL : {response.json()}"

def delete_collection(id):
    response = requests.delete(f"{SERVER_URL}/collection/{id}", headers=headers)
    if response.status_code != 200:
        raise f"HTTP FAIL : {response.json()}"
```