PORT="${PORT:-10000}"

# setup iptables to redirect 443 and 80 to envoy
if ! systemctl is-enabled netfilter-persistent; then
    # netfilter-persistent needs to be installed first for it to work
    tries=10
    while ! sudo apt-get -yq install netfilter-persistent; do
        tries=$((tries - 1))
        if (( tries == 0 )); then
            log "Unable to install netfilter-persistent after 10 tries"
        else
            log "Failed to install netfilter-persistent, sleeping"
            sleep $(( 15 - tries ))
        fi
    done
fi
if ! systemctl is-enabled iptables-persistent; then
    tries=10
    while ! sudo DEBIAN_FRONTEND=noninteractive apt-get -yq install iptables-persistent; do
        tries=$((tries - 1))
        if (( tries == 0 )); then
            log "Unable to install iptables-persistent after 10 tries"
        else
            log "Failed to install iptables-persistent, sleeping"
            sleep $(( 15 - tries ))
        fi
    done
fi

sudo iptables -t nat -L | grep -q "${PORT}" || {
    log "Port 10000 redirect not setup, adding iptables rules"
    sudo iptables -A PREROUTING -t nat -p tcp --dport 443 -j REDIRECT --to-port "$PORT"
    sudo iptables -A PREROUTING -t nat -p tcp --dport 80 -j REDIRECT --to-port "$((PORT + 1))"
    sudo mkdir -p /etc/iptables
    sudo /sbin/iptables-save | sudo tee /etc/iptables/rules.v4 > /dev/null
    sudo ip6tables-save | sudo tee /etc/iptables/rules.v6 > /dev/null
    sudo netfilter-persistent save
}

if [ "$AVOID_INIT" == false ]; then
    # when we are rebuilding a node, we'll set AVOID_INIT=false to make sure we don't quit early.
    # so, in this case, we need to explicitly knock docker-compose over, so any updated images will be used.
    sudo systemctl restart dh
fi

log "System setup of $(hostname) complete!"
# Do not change this message, it must be the very last thing in the log file
log "InitialDeephavenSetupComplete"
