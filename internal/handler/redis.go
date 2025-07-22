package handler

import (
	"fmt"
	"net/http"
	"redisorement/internal/redis"
) 


func SetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	val := r.URL.Query().Get("val")

	if key == "" || val == "" {
		http.Error(w, "key and val query params required", http.StatusBadRequest)
		return
	}

	err := redis.Rdb.Set(redis.Ctx, key, val, 0).Err()
	if err != nil {
		http.Error(w, "redis error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Set key '%s' to '%s'\n", key, val)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	if key == "" {
		http.Error(w, "key is necessary", http.StatusBadRequest)
	}

	val, err := redis.Rdb.Get(redis.Ctx, key).Result()
	if err != nil {
		http.Error(w, "redis error: "+err.Error(), http.StatusInternalServerError)
		return 
	}

	fmt.Fprintf(w, "Value for key '%s' is '%s'\n", key, val)
}