func getDataCached(key string) string{
    cache.lock.RLock()
    data, ok := cache.items[key]
    cache.lock.RUnlock()
    
    if ok {
        return data
    }
    cache.lock.Lock()
    defer cache.lock.Unlock()
    data, ok := cache.items[key] // double check
    if ok {
        return data
    }
    
    data = doRequest(key) // HL
    cache.items[key] = data
    return data
}
