curl -sfL https://get.k3s.io | INSTALL_K3S_VERSION="v1.24.10+k3s1" sh -s - server --cluster-init


docker run -d --restart=unless-stopped \
    -p 8080:80 -p 4443:443 \
    -v /opt/rancher:/var/lib/rancher \
    rancher/rancher:latest