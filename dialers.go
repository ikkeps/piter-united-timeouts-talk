
// MYSQL OMIT
mysql.RegisterDial("tcp", myDialFunc)

// ENDMYSQL OMIT

// HTTP OMIT
func clientWithTimeout(timeout time.Duration) *http.Client {
    tr := &http.Transport{
        ResponseHeaderTimeout: timeout,
        Dial: (&net.Dialer{
            Timeout: timeout,
            KeepAlive: time.Second,
        }).Dial,
        TLSHandshakeTimeout: timeout,
        MaxIdleConnsPerHost: 100,
    }
    return &http.Client{
        Timeout: timeout,
        Transport: tr,
        }
}

//ENDHTTP OMIT
