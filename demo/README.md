This directory contains the helm chart(s) used for the dhdemo application.

For now, this will be a localhost minikube deployment only.  
Once we are happy with this setup, we'll move on to a production setup.

If you are a Deephaven developer / contributor,  
make sure you have setup minikube on your local machine  
Ubuntu users: https://phoenixnap.com/kb/install-minikube-on-ubuntu



Notes:

Getting locally built images into minikube:

Option 1 (build in minikube):
eval $(minikube docker-env)
./gradlew prepareCompose
# now all images are built inside minikube

Option 2 (import into minikube):

./gradlew prepareCompose
docker-compose push

Then, import images to minikube:

minikube image load deephaven/server:local-build
minikube image load deephaven/grpc-proxy:local-build
minikube image load deephaven/web:local-build
minikube image load deephaven/envoy:local-build

minikube image load deephaven/server:local-build  deephaven/grpc-proxy:local-build  deephaven/web:local-build deephaven/envoy:local-build



There! Now you can reference server:local images in kubernets/minikube



# All in one Update Minikube:
./gradlew preCo && docker-compose push && minikube image load deephaven/server:local-build  deephaven/grpc-proxy:local-build  deephaven/web:local-build deephaven/envoy:local-build

# To change envoy log levels:
`<insert instructions how to shell into container through k8 pod>`
curl -X POST localhost:9090/logging?level=trace

curl -k localhost:9090/logging?router=debug -X POST && curl -k localhost:9090/logging?http=debug -X POST && curl -k localhost:9090/logging?http2=debug -X POST

cd /dh && sudo docker-compose pull && sudo systemctl start dh && sudo rm -rf /deployments && sleep 2 && sudo docker cp dh_demo-server_1:/deployments /deployments && sudo systemctl stop dh && echo "Done updating controller" && cd -

cd /deployments && sudo JAVA_OPTIONS="-Dquarkus.http.host=0.0.0.0 -Dquarkus.http.port=7117 -Djava.util.logging.manager=org.jboss.logmanager.LogManager -Dquarkus.http.access-log.enabled=true -Dquarkus.ssl.native=false" ./run-java.sh && cd -


minikube service --url dh-local



DEBUGGING:
sudo docker run --rm -it --entrypoint bash us-central1-docker.pkg.dev/deephaven-oss/deephaven/demo-server:0.5.0


created new json service key:
https://console.cloud.google.com/iam-admin/serviceaccounts/details/113408456843581678161/keys?authuser=2&project=deephaven-oss
upload to kubernetes:
mv ~/Downloads/deephaven-oss-dbdeef90ff00.json ~/Downloads/deephaven-svc-act.json
kubectl create secret generic dh-svc-act --from-file \
"~/Downloads/deephaven-svc-act.json"



Notes for customers using our helm chart (someday):

Step 1: install minikube OR enable Google Kubernetes Engine

1a (minikube): link to setup guide
1b (gke): create kubernetes cluster
-> use autopilot
-> select Public cluster (you can get private to work, but we're using public for simplicity)


PROJECT_ID=deephaven-oss
CLUSTER_NAME=dhce-auto
ZONE=us-central1
K8S_CONTEXT=gke_"$PROJECT_ID"_"$ZONE"_"$CLUSTER_NAME"
K8S_NAMESPACE=dh
DOCKER_VERSION=0.8.22

https://console.cloud.google.com/artifacts/create-repo?project=deephaven-oss



gcloud artifacts repositories create deephaven \
--repository-format=docker \
--location=$ZONE \
--description="Docker repository"

gcloud auth configure-docker ${ZONE}-docker.pkg.dev


enable artifact registry:
https://console.cloud.google.com/apis/library/artifactregistry.googleapis.com?project=deephaven-oss

docker tag deephaven/grpc-proxy:local-build ${ZONE}-docker.pkg.dev/${PROJECT_ID}/deephaven/grpc-proxy:$DOCKER_VERSION
docker tag deephaven/server:local-build ${ZONE}-docker.pkg.dev/${PROJECT_ID}/deephaven/server:$DOCKER_VERSION

docker tag deephaven/web:local-build ${ZONE}-docker.pkg.dev/${PROJECT_ID}/deephaven/web:$DOCKER_VERSION


docker push ${ZONE}-docker.pkg.dev/${PROJECT_ID}/deephaven/grpc-proxy:$DOCKER_VERSION &
docker push ${ZONE}-docker.pkg.dev/${PROJECT_ID}/deephaven/server:$DOCKER_VERSION &
docker push ${ZONE}-docker.pkg.dev/${PROJECT_ID}/deephaven/web:$DOCKER_VERSION &

docker push ${ZONE}-docker.pkg.dev/${PROJECT_ID}/deephaven/cert-wildcard-job:$DOCKER_VERSION &

docker tag deephaven/envoy:local-build ${ZONE}-docker.pkg.dev/${PROJECT_ID}/deephaven/envoy:$DOCKER_VERSION
docker push ${ZONE}-docker.pkg.dev/${PROJECT_ID}/deephaven/envoy:$DOCKER_VERSION


# OPTIONAL: add a grpcurl container, to debug grpc:

docker pull fullstorydev/grpcurl:v1.8.2
docker push ${ZONE}-docker.pkg.dev/${PROJECT_ID}/fullstorydev/grpcurl:v1.8.2 &


gcloud container clusters get-credentials "${CLUSTER_NAME}" \
--zone "${ZONE}" \
--project "${PROJECT_ID}"
kubectl config use-context "${K8S_CONTEXT}"
kubectl config get-contexts
kubectl config current-context

now, deploy the app!

kubectl apply -f demo/dh-prod.yaml

Next, setup some public DNS for your deployment
(TODO: get this moved to kubectl)

gcloud beta container clusters update $CLUSTER_NAME --project $PROJECT_ID --zone $ZONE --cluster-dns clouddns --cluster-dns-scope cluster









https://cloud.google.com/kubernetes-engine/docs/how-to/cloud-dns

gcloud beta container clusters create CLUSTER_NAME \
--cluster-dns clouddns --cluster-dns-scope cluster \
--cluster-version VERSION | --release-channel \
[--zone ZONE_NAME | --region REGION_NAME]



kubectl expose pod $pod --name dns-test --port 8080




{ kubectl delete deployment dh-local || true ; } && kubectl apply -f demo/dh-localhost.yaml



k get pods -o wide
# find node name
kubectl get nodes --output wide
# get external IP from the node you saw in pod list
gcloud compute firewall-rules create dh-api --project ${PROJECT_ID} --allow tcp:30080
gcloud compute firewall-rules create dh-admin --project ${PROJECT_ID} --allow tcp:30443


Setup DNS:

DNS_ZONE=dh-demo
DOMAIN_ROOT=deephaven.app
NODE_IP=34.149.181.117
MACHINE_NAME=demo


gcloud beta dns --project=${PROJECT_ID} managed-zones create ${DNS_ZONE} --description="DNS for Deephaven" --dns-name="${DOMAIN_ROOT}." --visibility="public" --dnssec-state="off"

gcloud dns --project=${PROJECT_ID} record-sets transaction start --zone=${DNS_ZONE}
gcloud dns --project=${PROJECT_ID} record-sets transaction add ${NODE_IP} --name=${MACHINE_NAME}.${DOMAIN_ROOT}. --ttl=300 --type=A --zone=${DNS_ZONE}
gcloud dns --project=${PROJECT_ID} record-sets transaction add ${NODE_IP} --name=*.${MACHINE_NAME}.${DOMAIN_ROOT}. --ttl=300 --type=A --zone=${DNS_ZONE}
gcloud dns --project=${PROJECT_ID} record-sets transaction execute --zone=${DNS_ZONE}




# Get certificates for your domain

https://cloud.google.com/load-balancing/docs/ssl-certificates/google-managed-certs
You are a project Owner or Editor (roles/owner or roles/editor).
You have both the Compute Security Admin role (compute.securityAdmin) and the Compute Network Admin role (compute.networkAdmin) in the project.
You have a custom role for the project that includes the compute.sslCertificates.* permissions and one or both of compute.targetHttpsProxies.* and compute.targetSslProxies.*, depending on the type of load balancer that you are using.



CERT_NAME=dh-demo-cert
CERT_DESC="Certificate used to enable deephaven https / tls"
DOMAINS_CSV="demo.deephaven.app"

gcloud compute ssl-certificates create "$CERT_NAME" \
--description="$CERT_DESC" \
--domains="$DOMAINS_CSV" \
--global \
--project "$PROJECT_ID"

# [optional] wait until this gcloud reports certificate is ACTIVE
# Note: cert can be added to load balancer while in PROVISIONING state
while ! gcloud compute ssl-certificates describe "$CERT_NAME" --project "${PROJECT_ID}" --global    --format="get(name,managed.status, managed.domainStatus)" --project "${PROJECT_ID}" | grep -q "ACTIVE"; do
    echo -n .
    sleep 1
done
echo ""
echo "Certificate $CERT_NAME is ACTIVE!"


# [optional] use a static IP address so you don't have to update DNS
# we are using --global and IPV6. You can use IPV4 and --region $ZONE if you prefer
DH_IP_ADDR=dh-ip
gcloud compute addresses create $DH_IP_ADDR \
--global \
--project "$PROJECT_ID"

# gcloud compute addresses create $DH_IP_ADDR \
# --global \
# --project "$PROJECT_ID" \
# --ip-version IPV6


Now, pass the name of your ip address to helm
--set dh.ipAddrName=$DH_IP




Progress notes:
expose web and server as separate services over separate protocols
web will be https, server as http/2, http redirects to https
server will need a sidecar to handle healthchecks (health check goes to the port of the container, not targetPort)
ditch envoy. ditch grpc-proxy. just web + server



grpc certs:

openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key \
-out server.crt -days 3650

# from https://stackoverflow.com/questions/47099664/grpc-java-ssl-server-side-authentication-certificate-generation
openssl pkcs8 -topk8 -nocrypt -in server.key -out server.key2



# maybe optional
openssl req -new -sha256 -key server.key -out server.csr
openssl x509 -req -sha256 -in server.csr -signkey server.key \
-out server.crt -days 3650



# alternative:
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365



CERT_ROOT=/tmp/certs

CERT_ROOT=/dh/ws0/deephaven-core/demo/certs

mkdir -p $CERT_ROOT
openssl req -x509 -nodes -newkey rsa:4096 \
    -keyout "$CERT_ROOT/tls.key" \
    -out "$CERT_ROOT/tls.crt" \
    -days 36500 \
    -subj '/CN=cluster.local/O=Company' \
    -addext 'extendedKeyUsage=serverAuth,clientAuth' \
    -addext "subjectAltName=DNS:cluster.local,DNS:${DOMAINS_CSV:-demo.deephaven.app},IP:127.0.0.1"

kubectl create secret tls dh-grpc-secret \
--cert="$CERT_ROOT/tls.crt" \
--key="$CERT_ROOT/tls.key"



# Pull lets encrypt certs and ca into gcloud
CERT_ROOT=/dh/ws0/deephaven-core/demo/certs/wildcards

# We need to create a gcloud ssl-certificates, so we can reference it as a pre-shared-cert
# TODO: make sure our service account binding for workers doesn't allow pulling this secret from gcloud
gcloud compute ssl-certificates create dh-wildcard-cert \
    --certificate "$CERT_ROOT/fullchain.pem" \
    --private-key "$CERT_ROOT/privkey.pem"

# We also upload the CA as a generic secret, so clients can use as truststore
kubectl create secret generic dh-wildcard-ca --from-file \
    "$CERT_ROOT/chain.pem"

# We're putting the wildcard cert and key into kubernetes as a secret.
# We really shouldn't do that, as a kubernetes savvy user might read it
# (TODO: make sure workers run without any kubernetes read permissions)
kubectl create secret tls dh-wildcard-cert \
    --cert "$CERT_ROOT/fullchain.pem" \
    --key "$CERT_ROOT/privkey.pem"

# this kubectl command to read secrets can be `curl`ed by a savvy user:
secret_raw="$(kubectl get secret dh-wildcard-cert -o go-template='{{range $k,$v := .data}}{{\"### \"}}{{$k}}{{\"\n\"}}{{$v|base64decode}}{{\"\n\n\"}}{{end}}')"


kubectl get secret dh-wildcard-cert -o go-template='{{range $k,$v := .data}}{{if eq $k "tls.key" }}{{$v|base64decode}}{{end}}{{end}}' > "$DH_SSL_DIR/tls.key"

kubectl get secret dh-wildcard-cert -o go-template='{{range $k,$v := .data}}{{if eq $k "tls.crt" }}{{$v|base64decode}}{{end}}{{end}}' > "$DH_SSL_DIR/tls.crt"





# check if the gateway is done (will fail if no address exposed yet)
kubectl get gateway dh-gateway -o=jsonpath="{.status.addresses[0].value}"


curl -s -L https://github.com/fullstorydev/grpcurl/releases/download/v1.8.2/grpcurl_1.8.2_linux_x86_64.tar.gz | tar -xvzf - -C /tmp

WS=/dh/ws0/deephaven-core
PROTOS="$WS/proto/proto-backplane-grpc/src/main/proto"
curl_args="-cacert $WS/demo/certs/tls.crt  -import-path $PROTOS -proto $PROTOS/grpc/health/v1/health.proto"
./grpcurl $curl_args demo.deephaven.app:8888 grpc.health.v1.Health/Check


PROTOS=/deployments/proto
curl_args="-cacert /etc/ssl/dh/ca.crt -import-path $PROTOS -proto $PROTOS/grpc/health/v1/health.proto"
grpcurl $curl_args 127.0.0.1:8888 grpc.health.v1.Health/Check



kubectl patch gateway/dh-gateway \
--type json \
--patch='[ { "op": "remove", "path": "/metadata/finalizers" } ]'






## Cluster setup
#1) Enable IAM
https://console.cloud.google.com/apis/api/iamcredentials.googleapis.com/overview

# 2) Set vars
CLUSTER_NAME="${CLUSTER_NAME:-dh-demo}"
PROJECT_ID="${PROJECT_ID:-deephaven-oss}"
ZONE="${ZONE:-us-central1}"
K8NS="${K8NS:-dh}"
K8_SRV_ACT="${K8_SRV_ACT:-dhadmin}"
GCE_SRV_ACT="${GCE_SRV_ACT:-dhadmin}"

# 3) Prepare environment
gcloud config set compute/region "$ZONE"
gcloud config set project "$PROJECT_ID"
gcloud components update

# 4a) Create cluster
gcloud container clusters create "$CLUSTER_NAME" \
    --workload-pool="${PROJECT_ID}.svc.id.goog" \
    --zone "$ZONE"
# OR: 4b) Update cluster
gcloud container clusters update "$CLUSTER_NAME" \
    --workload-pool="${PROJECT_ID}.svc.id.goog" \
    --zone "$ZONE"

kubectl create namespace "${K8NS}"
kubectl create serviceaccount --namespace "$K8NS" "$K8_SRV_ACT"
gcloud iam service-accounts create "$GCE_SRV_ACT"



gcloud iam service-accounts add-iam-policy-binding \
    --role roles/iam.workloadIdentityUser \
    --member "serviceAccount:${PROJECT_ID}.svc.id.goog[$K8NS/$K8_SRV_ACT]" \
    "${GCE_SRV_ACT}@${PROJECT_ID}.iam.gserviceaccount.com"

gcloud projects add-iam-policy-binding "$PROJECT_ID" \
    --role "roles/compute.securityAdmin" \
    --member "serviceAccount:${PROJECT_ID}.svc.id.goog[$K8NS/$K8_SRV_ACT]"

gcloud projects add-iam-policy-binding "$PROJECT_ID" \
    --role "roles/compute.networkAdmin" \
    --member "serviceAccount:${PROJECT_ID}.svc.id.goog[$K8NS/$K8_SRV_ACT]"



# TODO: reduce this to a minimum needed for certbot
#gcloud projects add-iam-policy-binding "$PROJECT_ID" \
#    --role "roles/dns.admin" \
#    --member "serviceAccount:${PROJECT_ID}.svc.id.goog[$K8NS/$K8_SRV_ACT]"
gcloud projects add-iam-policy-binding "$PROJECT_ID" \
    --role "roles/dns.changes.create" \
    --member "serviceAccount:${PROJECT_ID}.svc.id.goog[$K8NS/$K8_SRV_ACT]"



kubectl annotate serviceaccount \
    --namespace "$K8NS" \
    "$K8_SRV_ACT" \
    "iam.gke.io/gcp-service-account=${GCE_SRV_ACT}@${PROJECT_ID}.iam.gserviceaccount.com"



# CREATE A CLUSTER
CLUSTER_NAME="${CLUSTER_NAME:-dh-demo}"
gcloud container  clusters create "$CLUSTER_NAME" \
    --machine-type "n1-standard-4" \
    --region us-central1  --num-nodes 2 --enable-ip-alias  \
    --cluster-version "1.20"  -q

# INSTALL GATEWAY API INTO CLUSTER:
kubectl kustomize "github.com/kubernetes-sigs/gateway-api/config/crd?ref=v0.3.0" \
| kubectl apply -f -


PERMISSSIONS:
iam.serviceAccounts.create
container.clusters.get
container.clusters.update
RBAC needed:
https://v1-19.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#-strong-write-operations-serviceaccount-v1-core-strong-

https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fcloud-platform
https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fcloud-platform.read-only
https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fndev.clouddns.readonly
https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fndev.clouddns.readwrite








CLUSTER_NAME=dhce-auto
PROJECT_ID=deephaven-oss
ZONE=us-central1
K8S_CONTEXT=gke_"$PROJECT_ID"_"$ZONE"_"$CLUSTER_NAME"
K8S_NAMESPACE=dh
DOCKER_VERSION=0.0.4


cd demo/certs
docker build . -t ${ZONE}-docker.pkg.dev/${PROJECT_ID}/deephaven/cert-wildcard-job:$DOCKER_VERSION
docker push ${ZONE}-docker.pkg.dev/${PROJECT_ID}/deephaven/cert-wildcard-job:$DOCKER_VERSION


blech.
Go to https://console.cloud.google.com/iam-admin/serviceaccounts?authuser=2&orgonly=true&project=deephaven-oss&supportedpurview=organizationId
Create/download a key for your service account (json)
Rename to google-svc.json and place next to the Dockerfile in demo/certs

...ughhhhh, real instructions should be:
On a trusted machine, get your svc account json, run script, get certs, sed patch file

kubectl patch secret dh-wildcard-cert --type='strategic' --patch "$(cat secret-patch.json)"

docker build -t cert-wildcard-job .
docker run -d cert-wildcard-job

# debug script
docker run --entrypoint /bin/sh -it --rm cert-wildcard-job

docker ps 
# find container id


deephaven.app setup:

create a role (named DhControl in this example):
https://console.cloud.google.com/iam-admin/roles/create?project=deephaven-oss

grant the following:
artifactregistry.repositories.downloadArtifacts
container.clusters.get
container.clusters.getCredentials
container.clusters.list
container.secrets.get

create a service account named dh-controller:
https://console.cloud.google.com/iam-admin/serviceaccounts?project=deephaven-oss

grant that service account the IAM role you created:
https://console.cloud.google.com/iam-admin/iam?project=deephaven-oss
Note: you need to search for the role by the "nice name" given, not the id

create a firewall rule:
Logs: off
Targets: specified target tags
Tags: dh-demo
Source filter: IP ranges
Source IP ranges: 0.0.0.0/0  (or a different source if you are using VPN, which is preferred)
Specified protocols and ports:
tcp: 80,443

Run GoogleDeploymentManagerTest.testMachineSetup with demo.gradle having uncommented: systemProperty("noClean", "true")

This will create the machine.
Shell into it, and run prepare-snapshot.sh




gcloud compute images create deephaven-app-0-0-4 --source-disk=snapshot-root     --source-disk-zone=us-central1-a



# Make sure to reserve enough instances that we don't run out prematurely
gcloud compute reservations create demo-reservation --machine-type=n2d-standard-4 --zone=us-central1-f --vm-count=15 --project=deephaven-oss



TO ADD NEW USERS WHO CAN DEPLOY MACHINES:

Have them send a .pub ssh key, add it here:
https://console.cloud.google.com/compute/metadata?project=deephaven-oss&tab=sshkeys

Next, go to https://console.cloud.google.com/iam-admin/iam?project=deephaven-oss and edit/create a principal for their illumon.com email address. Give them the following permissions:
Demo Admin Role
Demo Controller Role

hit save.

Have user run `gcloud auth login` from their shell, to login with their illumon.com email.

Now, have the user run
./gradlew deployMachine -Phostname=any-machine-name


This should succeed.


# Deployment

Once ./gradlew deployDemo succeeds, and you have tested the new controller,  
your next step is to update the DNS record for the main site to point to the  
new controller.  The console will print the new controller's IP address,  
or you can find it using `dig +short controller-1-2-3.demo.deephaven.app`.

The DNS record is updated using the url:  
[https://console.cloud.google.com/net-services/dns/zones/deephaven-app/rrsets/demo.deephaven.app./A/view?project=deephaven-oss]

Once the old controller realizes it is no longer the leader,  
it will stop modifying the cluster; once the new controller notices the  
DNS change, it will also start modifying the cluster.

You may then turn off the old controller in the web console.  
[https://console.cloud.google.com/compute/instances?project=deephaven-oss]  
Currently, you will also need to manually delete the old controller instances.  
You can use the filter bar, searching for `Labels.dh-version:<oldversion>`,  
then select all matching machines and click Delete.