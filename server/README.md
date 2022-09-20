# Mock server for http request on large response

Default host: `localhost:3333`

Endpoints:
1. `/mirror` - [POST] Response with request body
2. `/json` - [POST] Response with large json body from fs (26MB)
3. `/json/cached` - [POST] Response with large json body from memory (26MB)

## Setup new GCP instance

1. Update and Upgrade all dependencies

```bash
sudo apt update && 
  sudo apt upgrade -y
```

2. Add Docker GPG public key

```bash
# Download public key
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/debian/gpg |
  sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

# Add to apt source list
echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" | 
  sudo tee /etc/apt/sources.list.d/docker.list >/dev/null
```

3. Install Docker

```bash
sudo apt update && 
  sudo apt install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin
```

4. Download JVSG server

```bash
git clone "https://github.com/kamontat/jvsg.git"
```

5. Start docker on background

```bash
cd jvsg/docker &&
  sudo docker compose up -d

# Verify docker start without problem
sudo docker ps -a
```
