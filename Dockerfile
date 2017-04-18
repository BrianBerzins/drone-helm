FROM alpine:latest

RUN mkdir /root/.kube

ADD https://storage.googleapis.com/kubernetes-release/release/v1.6.0/bin/linux/amd64/kubectl /usr/local/bin/kubectl
ADD https://storage.googleapis.com/kubernetes-helm/helm-v2.3.0-linux-amd64.tar.gz /tmp

RUN tar -zxvf tmp/helm-v2.3.0-linux-amd64.tar.gz -C /tmp \
    && mv /tmp/linux-amd64/helm /usr/local/bin/helm \
    && rm -rf /tmp \
    && chmod a+x /usr/local/bin/kubectl \
    && chmod a+x /usr/local/bin/helm

ADD drone-helm /bin/
ENTRYPOINT ["/bin/drone-helm"]

