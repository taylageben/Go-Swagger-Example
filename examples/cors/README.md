# Adding CORS Middleware

The generated API provides two methods for providing your own custom middleware in the configure file:
* setupMiddlewares
* setupGlobalMiddleware

To add CORS middleware, add the following to the global method:

```bash
corsHandler := cors.New(cors.Options{
		AllowedMethods: []string{"HEAD", "PUT", "GET", "POST", "DELETE", "OPTIONS"},
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Content-Type", "X-Requested-With"},
	})
	
return corsHandler.Handler(handler)
```