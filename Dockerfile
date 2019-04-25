FROM alpine:latest
MAINTAINER Portworx Inc. <support@portworx.com>

RUN apk add tar python3
RUN apk add ca-certificates
RUN apk upgrade
RUN pip3 install --upgrade pip

RUN wget -O /usr/local/bin/aws-iam-authenticator https://amazon-eks.s3-us-west-2.amazonaws.com/1.10.3/2018-07-26/bin/linux/amd64/aws-iam-authenticator
RUN chmod +x /usr/local/bin/aws-iam-authenticator

ARG GCLOUD_SDK=google-cloud-sdk-243.0.0-linux-x86_64.tar.gz
RUN wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/$GCLOUD_SDK
RUN tar xf $GCLOUD_SDK
RUN rm -rf $GCLOUD_SDK
# Remove the test directories
RUN rm -rf /google-cloud-sdk/platform/gsutil/third_party/oauth2client/tests
RUN rm -rf /google-cloud-sdk/platform/gsutil/third_party/rsa/tests
RUN rm -rf /google-cloud-sdk/platform/gsutil/third_party/httplib2/python2/httplib2/test
RUN pip3 install pyyaml==5.1 rsa==4.0 --upgrade -t /google-cloud-sdk/lib/third_party
# Don't need gsutil
RUN rm -rf /google-cloud-sdk/platform/gsutil

WORKDIR /

COPY ./bin/stork /
COPY ./bin/linux/storkctl /storkctl/linux/
COPY ./bin/darwin/storkctl /storkctl/darwin/
COPY ./bin/windows/storkctl.exe /storkctl/windows/
