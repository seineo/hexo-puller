# Hexo Puller

<img align="right" width="159px" src="https://oss.seineo.cn/images/202311151624249.png">

Hexo Puller is a simple git repository synchronization tool that provides a one-time solution for private cloud servers to synchronize hexo blogs.

Compared to existing methods like `hexo-deployer-git` and `rsync`, Hexo Puller has the following advantages:

- Requires configuration only once.
- Can synchronize blog markdown source files and theme configurations.
- No additional commands are necessary apart from pushing the blog to GitHub.





## Getting Started

1. Clone hexo puller:

```shell
git clone https://github.com/seineo/hexo-puller.git
```

2. Configure the paths for SSL certificate and private key in the `config/config.yaml`:

```yaml
tls:
  crt: xxx
  key: xxx
```

3. Compile and run hex puller(recommended in the background or as a daemon process):

```shell
go build
nohup ./hexo-puller > log.log 2>&1 &
```

3. Configure reverse proxy, for instance, in NGINX:

```nginx
server {
  listen 443 ssl;
  #...
  location / {
    root /home/xxx/blog/public; # point to your blog
    index index.html;
  }
  location /hexo-puller { 
    proxy_pass https://127.0.0.1:33333;  # pass to hexo puller
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header Content-Type "application/json"; 
  }
}
```

4. Set up a GitHub Action for your blog repository. There are only three things that need to be changed in the YAML file:
   1. `{{YOUR_SERVER}}`: your server IP or domain name;
   2. `{{YOUR_REPO_URL}}`: the address of your blog repository;
   3. `{{YOUR_TARGET_FOLDER}}`: the path to the local directory where your blog will be placed.

```yaml
name: Hexo Pull Notifier
run-name: Notify the server to pull the repository to update the blog content.
on: [push]
jobs:
  Notify:
    runs-on: ubuntu-latest
    steps:
      - name: Post data to backend server.
        run: >
          status_code=$(curl -s -o /dev/null -w "%{http_code}" https://{{YOUR_SERVER}}/hexo-puller --header 'Content-Type: application/json' --data '{"repoUrl": "{{YOUR_REPO_URL}}", "targetDir":"{{YOUR_TARGET_FOLDER}}"}');

          if [ "$status_code" -ne "200" ]; then
            echo "Error: status code is ${status_code}"
            exit 1
          else 
            echo "Success."
          fi
      - run: echo "🍏 This job's status is ${{ job.status }}."
```

Afterward, each push operation to the blog repository will prompt the server to pull the latest blog content into the target directory (it will clone if the content doesn't exist locally).

## How does Hexo Puller work？

![截屏2023-11-15 16.26.43](https://oss.seineo.cn/images/202311151626017.png)

The workflow of Hexo puller is straightforward. When a user pushes their local blog to the remote GitHub repository, it triggers a GitHub Action. This action sends an HTTP POST request to the Hexo Puller deployed on the server, carrying JSON-formatted data. Upon receiving the request, Hexo Puller, based on the repository address and target path provided in the data, will either clone or pull the repository to the specified local target path.
