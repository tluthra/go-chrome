package CacheStorage

/*
CacheID is the unique identifier of the Cache object.
*/
type CacheID string

/*
CacheDataEntry is a data entry.
*/
type CacheDataEntry struct {
	// Request URL.
	RequestURL string `json:"requestURL"`

	// Request method.
	RequestMethod string `json:"requestMethod"`

	// Request headers.
	RequestHeaders []*Header `json:"requestHeaders"`

	// Number of seconds since epoch.
	ResponseTime int `json:"responseTime"`

	// HTTP response status code.
	ResponseStatus int `json:"responseStatus"`

	// HTTP response status text.
	ResponseStatusText string `json:"responseStatusText"`

	// Response headers.
	ResponseHeaders []*Header `json:"responseHeaders"`
}

/*
Cache is a cache identifier.
*/
type Cache struct {
	// An opaque unique ID of the cache.
	CacheID CacheID `json:"cacheId"`

	// Security origin of the cache.
	SecurityOrigin string `json:"securityOrigin"`

	// The name of the cache.
	CacheName string `json:"cacheName"`
}

/*
Header is a single header value
*/
type Header struct {
	// Header name
	Name string `json:"name"`

	// Header value
	Value string `json:"value"`
}

/*
CachedResponse represents a cached response
*/
type CachedResponse struct {
	// Entry content, base64-encoded.
	Body string `json:"body"`
}

/*
RequestCacheNamesParams represents CacheStorage.requestCacheNames parameters.
*/
type RequestCacheNamesParams struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
}

/*
RequestEntriesParams represents CacheStorage.requestEntries parameters.
*/
type RequestEntriesParams struct {
	// ID of cache to get entries from.
	CacheID CacheID `json:"cacheId"`

	// Number of records to skip.
	SkipCount int `json:"skipCount"`

	// Number of records to fetch.
	PageSize int `json:"pageSize"`
}

/*
DeleteCacheParams represents CacheStorage.deleteCache parameters.
*/
type DeleteCacheParams struct {
	// Id of cache for deletion.
	CacheID CacheID `json:"cacheId"`
}

/*
DeleteEntryParams represents CacheStorage.deleteEntry parameters.
*/
type DeleteEntryParams struct {
	// ID of cache where the entry will be deleted.
	CacheID CacheID `json:"cacheId"`

	// URL spec of the request.
	Request string `json:"request"`
}

/*
RequestCachedResponseParams represents CacheStorage.requestCachedResponse parameters.
*/
type RequestCachedResponseParams struct {
	// ID of cache that contains the enty.
	CacheID CacheID `json:"cacheId"`

	// URL spec of the request.
	RequestURL string `json:"requestURL"`
}