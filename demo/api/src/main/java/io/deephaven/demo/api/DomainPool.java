package io.deephaven.demo.api;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

/**
 * DomainPool:
 * <p>
 * <p>
 * <p> A pool which maps from a String domain name to a {@link DomainMapping} object.
 * <p>
 */
public class DomainPool {

    private final Map<String, DomainMapping> allDomains = new ConcurrentHashMap<>();

    public long markAll() {
        long mark = System.currentTimeMillis();
        allDomains.values().forEach(d->d.setMark(mark - 1));
        return mark;
    }

    public void sweep(final long mark) {
        allDomains.entrySet().removeIf(next -> next.getValue().getMark() < mark);
    }

    public DomainMapping getOrCreate(final String subdomain, final String domainRoot) {
        return allDomains.computeIfAbsent(subdomain + "." + domainRoot, missing->
                new DomainMapping(subdomain, domainRoot));
    }
}
