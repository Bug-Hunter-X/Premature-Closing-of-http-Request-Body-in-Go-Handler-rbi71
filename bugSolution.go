func handleRequest(w http.ResponseWriter, r *http.Request) {
    // ... other code ...
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Process body here...
    // r.Body.Close() is not needed here, ioutil.ReadAll handles it
    // ... more code ...
} 